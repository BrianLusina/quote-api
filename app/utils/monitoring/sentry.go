package monitoring

import (
	"log"
	"quote/api/app/config"
	"time"

	"github.com/getsentry/sentry-go"
)

type contextKey int

const SomeContextKey = contextKey(1)

// Sentry is a wrapper for the Sentry client.
type SentryClient struct {
	Dsn string
	sentry.Client
}

// NewSentry creates a new Sentry client.
func NewSentry(config config.Monitoring) *SentryClient {
	dsn := config.Dsn

	sentrySyncTransport := sentry.NewHTTPSyncTransport()
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              dsn,
		Transport:        sentrySyncTransport,
		AttachStacktrace: true,
		BeforeSend: func(event *sentry.Event, hint *sentry.EventHint) *sentry.Event {
			if hint.Context != nil {
				// would give you stored string that now can be attached to the event
				hint.Context.Value(SomeContextKey)
			}
			return event
		},
	})

	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}

	return &SentryClient{Dsn: dsn}
}

func (s *SentryClient) CaptureException(err error) {
	sentry.CaptureException(err)
}

func (s *SentryClient) CaptureMessage(message string) {
	sentry.CaptureMessage(message)
}

func (s *SentryClient) Recover() {
	err := recover()

	if err != nil {
		sentry.CurrentHub().Recover(err)
		sentry.Flush(time.Second * 5)
	}
}
