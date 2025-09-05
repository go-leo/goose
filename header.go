package goose

import "net/http"

// CopyHeader copies all header values from the source header to the target header.
// It preserves all existing values in the target header and adds the source header values.
//
// Parameters:
//   - tgt: The target http.Header to which values will be copied
//   - src: The source http.Header from which values will be copied
func CopyHeader(tgt http.Header, src http.Header) {
	for key, values := range src {
		for _, value := range values {
			tgt.Add(key, value)
		}
	}
}
