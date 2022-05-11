package middleware

import (
	"quote/api/app/config"
	"quote/api/app/tools/logger"

	"github.com/gin-gonic/gin"
	// "github.com/gin-contrib/cors"
)

func NewCORSMiddleware(cors config.CorsConfig) Middleware {
	log := logger.NewLogger("log-cors")
	return func(context *gin.Context) {
		log.Warnf("CORS Header is enabled & set to: %s", cors.AllowedOrigins)
		context.Header("Access-Control-Allow-Origin", cors.AllowedOrigins)
		context.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, X-Registry-Auth, Authorization")
		context.Header("Access-Control-Allow-Methods", "HEAD, GET, POST, DELETE, PUT, OPTIONS")
	}
}
