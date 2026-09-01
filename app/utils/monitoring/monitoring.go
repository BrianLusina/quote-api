package monitoring

// Monitoring is an interface for monitoring clients to implement
type Monitoring interface {
	CaptureException(err error)
	CaptureMessage(message string)
	Recover()
}
