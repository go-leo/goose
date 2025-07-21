package gonic

import (
	"net/http/pprof"

	"github.com/gin-gonic/gin"
)

func AppendPProf[Router gin.IRoutes](router Router) Router {
	router.Any("/debug/pprof/", gin.WrapF(pprof.Index))
	router.Any("/debug/pprof/cmdline", gin.WrapF(pprof.Cmdline))
	router.Any("/debug/pprof/profile", gin.WrapF(pprof.Profile))
	router.Any("/debug/pprof/symbol", gin.WrapF(pprof.Symbol))
	router.Any("/debug/pprof/trace", gin.WrapF(pprof.Trace))
	return router
}
