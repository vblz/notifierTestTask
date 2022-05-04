package notifier

// Options contains all customizable options.
type Options struct {
	logger            Logger
	retryNumber       int
	contentType       string
	requestsPerSecond float64
}

// Option allows set customizable options.
type Option func(o *Options)

// Lgr sets the logger.
func Lgr(l Logger) Option {
	return func(o *Options) {
		o.logger = l
	}
}

// RetryNumber sets the maximum number of retries.
func RetryNumber(n int) Option {
	return func(o *Options) {
		o.retryNumber = n
	}
}

// ContentType sets the Content-Type header in outgoing request.
func ContentType(contentType string) Option {
	return func(o *Options) {
		o.contentType = contentType
	}
}

// RequestsPerSecond sets the maximum amount of outgoing requests per second.
func RequestsPerSecond(n float64) Option {
	return func(o *Options) {
		o.requestsPerSecond = n
	}
}

func defaultOptions() *Options {
	return &Options{
		logger:            noOpLogger,
		retryNumber:       3,
		contentType:       "application/octet-stream",
		requestsPerSecond: 100,
	}
}

func createOptions(opts []Option) *Options {
	options := defaultOptions()

	for _, o := range opts {
		o(options)
	}

	return options
}
