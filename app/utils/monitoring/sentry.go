package monitoring

import (
	"log"
	"quote/api/app/config"

	"github.com/getsentry/sentry-go"
)

// Sentry is a wrapper for the Sentry client.
type Sentry struct {
	Dsn string
}

// NewSentry creates a new Sentry client.
func NewSentry(config config.Monitoring) *Sentry {
	dsn := config.Dsn

	err := sentry.Init(sentry.ClientOptions{
		Dsn:              dsn,
		AttachStacktrace: true,
	})

	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}

	return &Sentry{Dsn: dsn}
}

func (s *Sentry) CaptureException(err error) {
	sentry.CaptureException(err)
}

func (s *Sentry) CaptureMessage(message string) {
	sentry.CaptureMessage(message)
}
