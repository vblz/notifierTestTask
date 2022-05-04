package main

import (
	"github.com/go-pkgz/lgr"
	"github.com/jessevdk/go-flags"
	notifierLib "github.com/vblz/notifierTestTask/notifier"
	"net/url"
	"os"
	"time"
)

var opts struct {
	URL string `required:"yes" short:"u" long:"url" env:"URL" description:"The URL to notify"`

	Interval time.Duration `default:"5s" short:"i" long:"interval" env:"INTERVAL" description:"Notification interval. See time.ParseDuration for format"`

	RequestsPerSecond float64 `default:"100" long:"requests-per-second" env:"RPS" description:"Maximal number of requests per seconds"`
	MaxRetries        int     `default:"3" long:"max-retries" env:"MAX_RETRIES" description:"Maximal number of retries"`

	Verbose []bool `short:"v" long:"verbose" description:"Show verbose debug information. Add more to more verbose"`
}

// parseOptions reads options from args and fills the opts var. Might exit from the app in case of an error during parsing.
func parseOptions() {
	if _, err := flags.Parse(&opts); err != nil {
		if err.(*flags.Error).Type == flags.ErrHelp {
			os.Exit(0)
		}
		os.Exit(1)
	}
}

// createLogger returns a logger set up with opts.
func createLogger() *lgr.Logger {
	loggerOpts := []lgr.Option{
		lgr.Msec,
		lgr.LevelBraces,
	}

	if len(opts.Verbose) == 1 {
		loggerOpts = append(loggerOpts, lgr.Debug)
	} else if len(opts.Verbose) > 1 {
		loggerOpts = append(loggerOpts, lgr.Trace)
	}

	return lgr.New(loggerOpts...)
}

func createNotifier(l lgr.L) *notifierLib.Notifier {
	notifierOpts := []notifierLib.Option{
		notifierLib.RequestsPerSecond(opts.RequestsPerSecond),
		notifierLib.ContentType("text/plain; charset=utf-8"),
		notifierLib.RetryNumber(opts.MaxRetries),
	}

	if len(opts.Verbose) > 2 {
		notifierOpts = append(notifierOpts, notifierLib.Lgr(l))
	}

	notifierURL, err := url.Parse(opts.URL)
	if err != nil {
		l.Logf("ERROR can't parse URL: %s", err)
		os.Exit(1)
	}

	return notifierLib.New(notifierURL, notifierOpts...)
}
