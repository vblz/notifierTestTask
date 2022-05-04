package ratelimitedscanner

import (
	"bytes"
	"context"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
)

func TestRateLimitedScanner_ReadLine(t *testing.T) {
	lines := []string{
		"asdasd",
		"",
		"test",
	}
	data := strings.Join(lines, "\n")
	interval := time.Millisecond * 50

	s := New(bytes.NewBufferString(data), interval)

	ctx := context.Background()

	received := make([]string, 0, len(lines))
	lastNow := time.Now().Add(-interval)
	for {
		line, ok := s.ReadLine(ctx)

		assert.GreaterOrEqual(t, time.Since(lastNow), interval-time.Millisecond*5)

		if !ok {
			break
		}

		lastNow = time.Now()

		received = append(received, string(line))
	}

	assert.Equal(t, lines, received)
}

func TestRateLimitedScanner_ReadLineContext(t *testing.T) {
	const cancelInterval = time.Millisecond * 10
	data := strings.Repeat("s\n", 100)
	interval := time.Second

	s := New(bytes.NewBufferString(data), interval)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(cancelInterval)
		cancel()
	}()

	start := time.Now()
	received := ""
	for {
		l, ok := s.ReadLine(ctx)
		if !ok {
			break
		}
		received += string(l)
	}

	assert.Less(t, len(received), len(data))
	assert.Less(t, time.Since(start), cancelInterval*2)
}
