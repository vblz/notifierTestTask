package notifier

import (
	"golang.org/x/time/rate"
)

func createTokenBucketLimiter(requestsPerSecond float64) limiter {
	burst := int(requestsPerSecond)
	if burst < 1 {
		burst = 1
	}
	return rate.NewLimiter(rate.Limit(requestsPerSecond), burst)
}
