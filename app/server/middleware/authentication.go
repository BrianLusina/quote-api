package middleware

import (
	"net/http"
	"quote/api/app/config"
	"quote/api/app/tools/logger"

	"github.com/gin-gonic/gin"
)

// NewAuthenticationMiddleware returns a middleware that will check if the user is authenticated
func NewAuthenticationMiddleware(config config.AuthConfig) Middleware {
	log := logger.NewLogger("auth-mdl-logger")
	return func(ctx *gin.Context) {
		username, password, hasAuth := ctx.Request.BasicAuth()

		if hasAuth && username == config.Username && password == config.Password {
			log.Infof("User %s authenticated", username)
			ctx.Set("user", username)
			ctx.Next()
		} else {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			ctx.Writer.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
		}
	}
}
