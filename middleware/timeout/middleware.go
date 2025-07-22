package timeout

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

func Middleware(duration time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		ctx, cancel := context.WithTimeout(ctx, duration)
		defer cancel()
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
