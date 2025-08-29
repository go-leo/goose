package gen

import (
	"errors"
	"fmt"
	"net/url"
	"path"
	"strings"
	"unicode"

	"golang.org/x/net/http/httpguts"
)

// A Pattern is something that can be matched against an HTTP request.
// It has an optional method, an optional host, and a path.
type Pattern struct {
	str    string // original string
	method string
	host   string
	// The representation of a path differs from the surface syntax, which
	// simplifies most algorithms.
	//
	// Paths ending in '/' are represented with an anonymous "..." wildcard.
	// For example, the path "a/" is represented as a literal segment "a" followed
	// by a segment with multi==true.
	//
	// Paths ending in "{$}" are represented with the literal segment "/".
	// For example, the path "a/{$}" is represented as a literal segment "a" followed
	// by a literal segment "/".
	segments []Segment
	loc      string // source location of registering call, for helpful messages
}

func (p *Pattern) String() string { return p.str }

func (p *Pattern) lastSegment() Segment {
	return p.segments[len(p.segments)-1]
}

// A Segment is a pattern piece that matches one or more path segments, or
// a trailing slash.
//
// If wild is false, it matches a literal Segment, or, if s == "/", a trailing slash.
// Examples:
//
//	"a" => Segment{s: "a"}
//	"/{$}" => Segment{s: "/"}
//
// If wild is true and multi is false, it matches a single path Segment.
// Example:
//
//	"{x}" => Segment{s: "x", wild: true}
//
// If both wild and multi are true, it matches all remaining path segments.
// Example:
//
//	"{rest...}" => Segment{s: "rest", wild: true, multi: true}
type Segment struct {
	s     string // literal or wildcard name or "/" for "/{$}".
	wild  bool
	multi bool // "..." wildcard
}

// parsePattern parses a string into a Pattern.
// The string's syntax is
//
//	[METHOD] [HOST]/[PATH]
//
// where:
//   - METHOD is an HTTP method
//   - HOST is a hostname
//   - PATH consists of slash-separated segments, where each segment is either
//     a literal or a wildcard of the form "{name}", "{name...}", or "{$}".
//
// METHOD, HOST and PATH are all optional; that is, the string can be "/".
// If METHOD is present, it must be followed by a single space.
// Wildcard names must be valid Go identifiers.
// The "{$}" and "{name...}" wildcard must occur at the end of PATH.
// PATH may end with a '/'.
// Wildcard names in a path must be distinct.
func ParsePattern(s string) (_ *Pattern, err error) {
	if len(s) == 0 {
		return nil, errors.New("empty pattern")
	}
	off := 0 // offset into string
	defer func() {
		if err != nil {
			err = fmt.Errorf("at offset %d: %w", off, err)
		}
	}()

	method, rest, found := strings.Cut(s, " ")
	if !found {
		rest = method
		method = ""
	}
	if method != "" && !validMethod(method) {
		return nil, fmt.Errorf("invalid method %q", method)
	}
	p := &Pattern{str: s, method: method}

	if found {
		off = len(method) + 1
	}
	i := strings.IndexByte(rest, '/')
	if i < 0 {
		return nil, errors.New("host/path missing /")
	}
	p.host = rest[:i]
	rest = rest[i:]
	if j := strings.IndexByte(p.host, '{'); j >= 0 {
		off += j
		return nil, errors.New("host contains '{' (missing initial '/'?)")
	}
	// At this point, rest is the path.
	off += i

	// An unclean path with a method that is not CONNECT can never match,
	// because paths are cleaned before matching.
	if method != "" && method != "CONNECT" && rest != cleanPath(rest) {
		return nil, errors.New("non-CONNECT pattern with unclean path can never match")
	}

	seenNames := map[string]bool{} // remember wildcard names to catch dups
	for len(rest) > 0 {
		// Invariant: rest[0] == '/'.
		rest = rest[1:]
		off = len(s) - len(rest)
		if len(rest) == 0 {
			// Trailing slash.
			p.segments = append(p.segments, Segment{wild: true, multi: true})
			break
		}
		i := strings.IndexByte(rest, '/')
		if i < 0 {
			i = len(rest)
		}
		var seg string
		seg, rest = rest[:i], rest[i:]
		if i := strings.IndexByte(seg, '{'); i < 0 {
			// Literal.
			seg = pathUnescape(seg)
			p.segments = append(p.segments, Segment{s: seg})
		} else {
			// Wildcard.
			if i != 0 {
				return nil, errors.New("bad wildcard segment (must start with '{')")
			}
			if seg[len(seg)-1] != '}' {
				return nil, errors.New("bad wildcard segment (must end with '}')")
			}
			name := seg[1 : len(seg)-1]
			if name == "$" {
				if len(rest) != 0 {
					return nil, errors.New("{$} not at end")
				}
				p.segments = append(p.segments, Segment{s: "/"})
				break
			}
			name, multi := strings.CutSuffix(name, "...")
			if multi && len(rest) != 0 {
				return nil, errors.New("{...} wildcard not at end")
			}
			if name == "" {
				return nil, errors.New("empty wildcard")
			}
			if !isValidWildcardName(name) {
				return nil, fmt.Errorf("bad wildcard name %q", name)
			}
			if seenNames[name] {
				return nil, fmt.Errorf("duplicate wildcard name %q", name)
			}
			seenNames[name] = true
			p.segments = append(p.segments, Segment{s: name, wild: true, multi: multi})
		}
	}
	return p, nil
}

func validMethod(method string) bool {
	/*
	     Method         = "OPTIONS"                ; Section 9.2
	                    | "GET"                    ; Section 9.3
	                    | "HEAD"                   ; Section 9.4
	                    | "POST"                   ; Section 9.5
	                    | "PUT"                    ; Section 9.6
	                    | "DELETE"                 ; Section 9.7
	                    | "TRACE"                  ; Section 9.8
	                    | "CONNECT"                ; Section 9.9
	                    | extension-method
	   extension-method = token
	     token          = 1*<any CHAR except CTLs or separators>
	*/
	return len(method) > 0 && strings.IndexFunc(method, isNotToken) == -1
}

func isNotToken(r rune) bool {
	return !httpguts.IsTokenRune(r)
}

// cleanPath returns the canonical path for p, eliminating . and .. elements.
func cleanPath(p string) string {
	if p == "" {
		return "/"
	}
	if p[0] != '/' {
		p = "/" + p
	}
	np := path.Clean(p)
	// path.Clean removes trailing slash except for root;
	// put the trailing slash back if necessary.
	if p[len(p)-1] == '/' && np != "/" {
		// Fast path for common case of p being the string we want:
		if len(p) == len(np)+1 && strings.HasPrefix(p, np) {
			np = p
		} else {
			np += "/"
		}
	}
	return np
}

func pathUnescape(path string) string {
	u, err := url.PathUnescape(path)
	if err != nil {
		// Invalidly escaped path; use the original
		return path
	}
	return u
}

func isValidWildcardName(s string) bool {
	if s == "" {
		return false
	}
	// Valid Go identifier.
	for i, c := range s {
		if !unicode.IsLetter(c) && c != '_' && (i == 0 || !unicode.IsDigit(c)) {
			return false
		}
	}
	return true
}
