package main

import (
	"context"
	"github.com/go-pkgz/lgr"
	"github.com/vblz/notifierTestTask/ratelimitedscanner"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	parseOptions()

	l := createLogger()
	l.Logf("DEBUG start")

	ctx := createContext(l)
	n := createNotifier(l)
	s := ratelimitedscanner.New(os.Stdin, opts.Interval)

	messagesProcessed := sendMessages(ctx, n, s, l)

	if messagesProcessed == 0 {
		l.Logf("WARN no messages were processed")
	}

	l.Logf("DEBUG %d messages processed", messagesProcessed)
}

// createContext returns a context which cancels with os exit signals.
func createContext(l lgr.L) context.Context {
	ctx, cancel := context.WithCancel(context.Background())

	go func() { // catch signal and invoke graceful termination
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
		s := <-stop
		l.Logf("WARN interrupted by signal %s", s.String())
		cancel()
	}()

	return ctx
}
