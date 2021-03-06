package middleware

import (
	"quote/api/app/config"
	"quote/api/app/tools/logger"
	"time"

	"github.com/gin-gonic/gin"
)

// NewLoggingMiddleware sets up logging middleware in application
func NewLoggingMiddleware(config config.LoggingConfig) Middleware {
	log := logger.NewLogger("log-requests")
	log.EnableJSONOutput(config.EnableJSONOutput)

	return func(context *gin.Context) {
		start := time.Now()

		// before request
		context.Next()

		// get latency
		latency := time.Since(start)

		// Log request
		log.Infof("Incoming Request. Status: %d. Method: %s. Path: %s. Latency: %s", context.Writer.Status(), context.Request.Method, context.Request.URL.EscapedPath(), latency)
	}
}
