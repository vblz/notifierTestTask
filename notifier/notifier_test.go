package notifier

import (
	"bytes"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

func TestNotifier_Notify(t *testing.T) {
	ctx := context.Background()
	body := bytes.NewBufferString("test")
	sendMoq := &senderMock{
		SendFunc: func(ctxMoqParam context.Context, r io.Reader) error {
			assert.Equal(t, ctx, ctxMoqParam)
			assert.Equal(t, body, r)
			return nil
		},
	}
	limiterMoq := &limiterMock{
		WaitFunc: func(ctxMoqParam context.Context) error {
			assert.Equal(t, ctx, ctxMoqParam)
			return nil
		},
	}
	n := &Notifier{
		sender:  sendMoq,
		limiter: limiterMoq,
		l:       noOpLogger,
	}
	errCh := n.Notify(ctx, body)
	for range errCh {
		assert.Fail(t, "should be no err")
	}
	assert.Equal(t, 1, len(sendMoq.SendCalls()))
	assert.Equal(t, 1, len(limiterMoq.WaitCalls()))
}

func TestNotifier_NotifyWaitErr(t *testing.T) {
	ctx := context.Background()
	testErr := errors.New("test error")
	sendMoq := &senderMock{
		SendFunc: func(ctxMoqParam context.Context, reader io.Reader) error {
			return testErr
		},
	}
	limiterMoq := &limiterMock{
		WaitFunc: func(ctxMoqParam context.Context) error {
			return nil
		},
	}
	n := &Notifier{
		sender:  sendMoq,
		limiter: limiterMoq,
		l:       noOpLogger,
	}
	errCh := n.Notify(ctx, nil)
	errCount := 0
	for err := range errCh {
		errCount++
		assert.Equal(t, testErr, err)
	}
	assert.Equal(t, 1, len(sendMoq.SendCalls()))
	assert.Equal(t, 1, len(limiterMoq.WaitCalls()))
	assert.Equal(t, 1, errCount)
}

func TestNotifier_NotifySendErr(t *testing.T) {
	ctx := context.Background()
	testErr := errors.New("test error")
	sendMoq := &senderMock{}
	limiterMoq := &limiterMock{
		WaitFunc: func(ctxMoqParam context.Context) error {
			return testErr
		},
	}
	n := &Notifier{
		sender:  sendMoq,
		limiter: limiterMoq,
		l:       noOpLogger,
	}
	errCh := n.Notify(ctx, nil)
	errCount := 0
	for err := range errCh {
		errCount++
		assert.Equal(t, testErr, err)
	}
	assert.Equal(t, 1, len(limiterMoq.WaitCalls()))
	assert.Equal(t, 1, errCount)
}

func TestNotifier_NotifyMultiple(t *testing.T) {
	ctx := context.Background()
	messages := []io.Reader{
		bytes.NewBufferString("test1"),
		nil,
		bytes.NewBufferString("test2"),
	}
	sendMoq := &senderMock{
		SendFunc: func(ctxMoqParam context.Context, r io.Reader) error {
			return nil
		},
	}
	limiterMoq := &limiterMock{
		WaitFunc: func(ctxMoqParam context.Context) error {
			return nil
		},
	}

	n := &Notifier{
		sender:  sendMoq,
		limiter: limiterMoq,
		l:       noOpLogger,
	}
	errCh := n.NotifyMultiple(ctx, messages)
	for range errCh {
		assert.Fail(t, "should be no err")
	}
	assert.Equal(t, len(messages), len(sendMoq.SendCalls()))
	assert.Equal(t, len(messages), len(limiterMoq.WaitCalls()))
}

func TestNotifier_NotifyMultipleOneError(t *testing.T) {
	ctx := context.Background()
	messages := []io.Reader{
		bytes.NewBufferString("test1"),
		bytes.NewBufferString("test2"),
	}
	testErr := errors.New("test error")
	sendMoq := &senderMock{
		SendFunc: func(ctxMoqParam context.Context, r io.Reader) error {
			if r == messages[0] {
				return testErr
			}
			return nil
		},
	}
	limiterMoq := &limiterMock{
		WaitFunc: func(ctxMoqParam context.Context) error {
			return nil
		},
	}

	n := &Notifier{
		sender:  sendMoq,
		limiter: limiterMoq,
		l:       noOpLogger,
	}
	errCh := n.NotifyMultiple(ctx, messages)
	errCount := 0
	for err := range errCh {
		errCount++
		assert.Equal(t, messages[0], err.Message)
		assert.True(t, errors.Is(err, testErr))
	}
	assert.Equal(t, 1, errCount)
	assert.Equal(t, len(messages), len(sendMoq.SendCalls()))
	assert.Equal(t, len(messages), len(limiterMoq.WaitCalls()))
}
