package handler

import (
	"catalog_service/internal/pkg"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// logging http requests
func LoggerMiddleware(log *pkg.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		//putting logger in context.Context
		ctx := pkg.ToContext(c.Request.Context(), log)
		c.Request = c.Request.WithContext(ctx)

		c.Next()

		//logging the result after finising request
		duration := time.Since(start)
		log.Info("HTTP Request",
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", c.Writer.Status()),
			zap.Duration("duration", duration),
		)

	}
}
