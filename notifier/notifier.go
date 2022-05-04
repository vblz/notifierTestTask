package notifier

import (
	"context"
	"github.com/vblz/notifierTestTask/notifier/http"
	"io"
	"net/url"
	"sync"
)

//go:generate moq -out notifier_moq_test.go . sender limiter

type sender interface {
	Send(context.Context, io.Reader) error
}

type limiter interface {
	Wait(context.Context) error
}

// Notifier sends notifications.
type Notifier struct {
	sender  sender
	limiter limiter
	l       Logger
}

// New returns *Notifier that sends notifications.
func New(endpoint *url.URL, opts ...Option) *Notifier {
	options := createOptions(opts)

	res := &Notifier{
		sender:  http.New(endpoint.String(), options.contentType, options.retryNumber, toLeveledLogger(options.logger)),
		l:       options.logger,
		limiter: createTokenBucketLimiter(options.requestsPerSecond),
	}

	return res
}

// NotifyMultiple message using Notify. If an error happens, it will be wrapped in MessageError containing a link to a message.
func (n *Notifier) NotifyMultiple(ctx context.Context, messages []io.Reader) <-chan *MessageError {
	res := make(chan *MessageError)
	wg := sync.WaitGroup{}
	wg.Add(len(messages))

	go func() {
		wg.Wait()
		close(res)
	}()

	for _, m := range messages {
		go func(message io.Reader) {
			defer wg.Done()

			ch := n.Notify(ctx, message)
			for err := range ch {
				res <- &MessageError{
					Message: message,
					Err:     err,
				}
			}
		}(m)
	}

	return res
}

// Notify a notification with given context and message.
// If the provided message is also an io.Closer, it will be closed after performing HTTP request.
func (n *Notifier) Notify(ctx context.Context, message io.Reader) <-chan error {
	n.l.Logf("TRACE notify starts")

	err := make(chan error)
	go n.sendInternal(ctx, message, err)

	return err
}

func (n *Notifier) sendInternal(ctx context.Context, message io.Reader, errChan chan<- error) {
	defer close(errChan)

	err := n.limiter.Wait(ctx)
	if err != nil {
		n.l.Logf("INFO request was canceled while waiting")
		errChan <- err
		return
	}

	err = n.sender.Send(ctx, message)
	if err != nil {
		errChan <- err
		return
	}
}
