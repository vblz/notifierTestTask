package main

import (
	"context"
	"github.com/go-pkgz/lgr"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"testing"
)

func Test_sendMessages(t *testing.T) {
	ctx := context.Background()
	messages := [][]byte{
		[]byte("test1"),
		{},
		[]byte("test2"),
		[]byte("test3"),
	}

	scannerCall := 0
	s := &scannerMock{
		ReadLineFunc: func(contextMoqParam context.Context) ([]byte, bool) {
			assert.Equal(t, ctx, contextMoqParam)
			if scannerCall >= len(messages) {
				return nil, false
			}
			scannerCall++
			return messages[scannerCall-1], true
		},
	}

	notifierCall := 0
	n := &notifierMock{
		NotifyFunc: func(contextMoqParam context.Context, reader io.Reader) <-chan error {
			assert.Equal(t, ctx, contextMoqParam)

			b, err := io.ReadAll(reader)
			require.NoError(t, err)
			assert.Equal(t, messages[notifierCall], b)
			notifierCall++

			res := make(chan error)
			close(res)
			return res
		},
	}

	messagesSent := sendMessages(ctx, n, s, lgr.Default())
	assert.Equal(t, len(messages), messagesSent)
	assert.Equal(t, len(messages), scannerCall)
	assert.Equal(t, len(messages), notifierCall)

}
