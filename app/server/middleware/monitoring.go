package middleware

import (
	"quote/api/app/tools/logger"
	"quote/api/app/utils/monitoring"

	"github.com/gin-gonic/gin"
)

func NewMonitoringMiddleware(monitoring monitoring.Monitoring) Middleware {
	log := logger.NewLogger("monitoring-log")
	return func(ctx *gin.Context) {

		defer func() {
			err := recover()

			if err != nil {
				log.Errorf("Error encountered %s", err)

				monitoring.CaptureException(err.(error))
			}
		}()
	}
}
