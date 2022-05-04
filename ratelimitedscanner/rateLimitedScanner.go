package ratelimitedscanner

import (
	"bufio"
	"context"
	"golang.org/x/time/rate"
	"io"
	"time"
)

// RateLimitedScanner reads from specified reader with rate limit with given interval.
type RateLimitedScanner struct {
	limiter *rate.Limiter
	scanner *bufio.Scanner
}

// New creates a new RateLimitedScanner with given params.
func New(reader io.Reader, interval time.Duration) *RateLimitedScanner {
	return &RateLimitedScanner{
		limiter: rate.NewLimiter(rate.Every(interval), 1),
		scanner: bufio.NewScanner(reader),
	}
}

// ReadLine returns either a line from reader and state of reading. The state is false in case end of reader or context was done.
// If state is false, the byte result should be ignored. Blocks the flow with the rate limit.
func (r *RateLimitedScanner) ReadLine(ctx context.Context) ([]byte, bool) {
	if err := r.limiter.Wait(ctx); err != nil {
		return nil, false
	}

	if !r.scan(ctx) {
		return nil, false
	}

	return r.scanner.Bytes(), true
}

func (r *RateLimitedScanner) scan(ctx context.Context) bool {
	scanCh := make(chan bool)
	defer close(scanCh)

	go func() {
		scanCh <- r.scanner.Scan()
	}()
	select {
	case <-ctx.Done():
		return false
	case scanRes := <-scanCh:
		return scanRes
	}
}
