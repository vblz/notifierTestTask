package notifier

import (
	"github.com/vblz/notifierTestTask/notifier/http"
)

// Logger is interface used for logging. The package writes loglevel as the first word in capital case.
// Levels: TRACE DEBUG INFO ERROR
type Logger interface {
	Logf(format string, v ...interface{})
}

type funcLogger func(format string, v ...interface{})

func (f funcLogger) Logf(format string, v ...interface{}) { f(format, v...) }

var noOpLogger = funcLogger(func(format string, v ...interface{}) {})

func toLeveledLogger(l Logger) http.LeveledLogger {
	return leveledLogger(l.Logf)
}

type leveledLogger func(format string, v ...interface{})

func (ll leveledLogger) Debug(msg string, keysAndValues ...interface{}) {
	ll("DEBUG "+msg, keysAndValues...)
}

func (ll leveledLogger) Info(msg string, keysAndValues ...interface{}) {
	ll("INFO "+msg, keysAndValues...)
}

func (ll leveledLogger) Warn(msg string, keysAndValues ...interface{}) {
	ll("WARN "+msg, keysAndValues...)
}

func (ll leveledLogger) Error(msg string, keysAndValues ...interface{}) {
	ll("ERROR "+msg, keysAndValues...)
}
