package middleware

import (
	"net/http"
	"quote/api/app/config"
	"quote/api/app/tools/logger"
	"regexp"

	"github.com/gin-gonic/gin"
)

var (
	healthRegex = regexp.MustCompile("^(/health)$")
)

// NewAuthenticationMiddleware returns a middleware that will check if the user is authenticated
func NewAuthenticationMiddleware(config config.AuthConfig) Middleware {
	log := logger.NewLogger("auth-log")
	return func(ctx *gin.Context) {
		username, password, hasAuth := ctx.Request.BasicAuth()

		requestUrl := ctx.Request.URL
		// requestMethod := ctx.Request.Method

		requestPath := requestUrl.Path

		if healthRegex.MatchString(requestPath) {
			ctx.Next()
			return
		}

		if hasAuth {
			if username == config.Username && password == config.Password {
				log.Debugf("User %s authenticated", username)
				ctx.Set("user", username)
				ctx.Next()
			}
		} else {
			log.Errorf("User %s not authenticated", username)
			ctx.AbortWithStatus(http.StatusUnauthorized)
			ctx.Writer.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
		}
	}
}
