package goose

import "strings"

// URLPath constructs a URL path by replacing placeholders in the template path with actual values.
// Placeholders are defined using curly braces, e.g., "/users/{id}".
// The function supports both regular placeholders and spread placeholders ending with "...".
//
// Parameters:
//   - path: The URL template path containing placeholders
//   - pairs: A map where keys are placeholder names (without braces) and values are their replacements
//
// Returns:
//   - string: The constructed URL path with placeholders replaced by actual values
//
// Example:
//   URLPath("/users/{id}", map[string]string{"id": "123"}) returns "/users/123"
//   URLPath("/files/{path...}", map[string]string{"path": "dir/file.txt"}) returns "/files/dir/file.txt"
func URLPath(path string, pairs map[string]string) string {
	sections := strings.Split(path, "/")
	for i := 0; i < len(sections); i++ {
		section := sections[i]
		if section == "" {
			continue
		}
		if !strings.HasPrefix(section, "{") {
			continue
		}
		if !strings.HasSuffix(section, "}") {
			continue
		}
		if strings.HasSuffix(section, "...}") {
			section = section[1 : len(section)-4]
		} else {
			section = section[1 : len(section)-1]
		}
		sections[i] = pairs[section]
	}
	return strings.Join(sections, "/")
}
