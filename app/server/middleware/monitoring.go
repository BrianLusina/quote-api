package middleware

import (
	"quote/api/app/tools/logger"
	"time"

	"github.com/gin-gonic/gin"

	sentrygin "github.com/getsentry/sentry-go/gin"
)

func NewMonitoringMiddleware() Middleware {
	log := logger.NewLogger("monitoring-log")
	return func(ctx *gin.Context) {
		sentrygin.New(sentrygin.Options{
			Repanic: true,
		})

		hub := sentrygin.GetHubFromContext(ctx)

		defer func() {
			err := recover()

			if err != nil {
				log.Errorf("Error encountered %s", err)

				if hub != nil {
					log.Errorf("Exception captured")

					hub.CaptureException(err.(error))
					hub.Flush((time.Second * 5))
				}
			}
		}()

		// ctx.Next()
	}
}
