package middleware

import (
	"quote/api/app/config"

	"github.com/gin-gonic/gin"
)

func NewCORSMiddleware(cors config.CorsConfig) Middleware {
	return func(context *gin.Context) {
		context.Header("Access-Control-Allow-Origin", cors.AllowedOrigins)
		context.Header("Access-Control-Allow-Methods", cors.AllowedMethods)
		context.Header("Access-Control-Allow-Headers", cors.AllowedHeaders)
		context.Header("Access-Control-Allow-Credentials", cors.AllowCredentials)
		context.Header("Access-Control-Max-Age", cors.MaxAge)
		context.Header("Access-Control-Expose-Headers", cors.ExposedHeaders)
		context.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		context.Header("Access-Control-Allow-Methods", "HEAD, GET, POST, DELETE, PUT, OPTIONS")
	}
}
