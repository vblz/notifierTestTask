package notifier

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/time/rate"
	"testing"
)

func TestCreateTokenBucketLimiter(t *testing.T) {
	tests := []struct {
		name          string
		rps           float64
		expectedBurst int
	}{
		{"regular", 155.5, 155},
		{"below zero", 0.7, 1},
		{"below zero", 0.3, 1},
		{"negative", -1.7, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := createTokenBucketLimiter(tt.rps)
			rateLimiter, ok := l.(*rate.Limiter)
			require.True(t, ok)
			assert.Equal(t, rate.Limit(tt.rps), rateLimiter.Limit())
			assert.Equal(t, tt.expectedBurst, rateLimiter.Burst())
		})
	}
}
