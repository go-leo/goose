package goose

import "net/http"

func CopyHeader(tgt http.Header, src http.Header) {
	for key, values := range src {
		for _, value := range values {
			tgt.Add(key, value)
		}
	}
}
