package middleware

import (
	"quote/api/app/tools/logger"
	"time"

	"github.com/gin-gonic/gin"

	sentrygin "github.com/getsentry/sentry-go/gin"
)

func NewMonitoringMiddleware() Middleware {
	log := logger.NewLogger("log-monitoring")
	return func(ctx *gin.Context) {
		log.Info("Monitoring middleware is enabled")
		sentrygin.New(sentrygin.Options{
			Repanic: true,
		})

		hub := sentrygin.GetHubFromContext(ctx)

		defer func() {
			err := recover()

			if err != nil {
				if hub != nil {
					log.Errorf("Error encountered %s", err)

					hub.CaptureException(err.(error))
					hub.Flush((time.Second * 5))
				}
			}
		}()

		// ctx.Next()
	}
}
