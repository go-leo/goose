package goose

import (
	"net/http"
	"net/http/pprof"
)

func AppendPProf(router *http.ServeMux) *http.ServeMux {
	http.HandleFunc("GET /debug/pprof/", pprof.Index)
	http.HandleFunc("GET /debug/pprof/cmdline", pprof.Cmdline)
	http.HandleFunc("GET /debug/pprof/profile", pprof.Profile)
	http.HandleFunc("GET /debug/pprof/symbol", pprof.Symbol)
	http.HandleFunc("GET /debug/pprof/trace", pprof.Trace)
	return router
}

func AppendHealth(router *http.ServeMux) *http.ServeMux {
	router.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {})
	return router
}
