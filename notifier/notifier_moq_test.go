// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package notifier

import (
	"context"
	"io"
	"sync"
)

// Ensure, that senderMock does implement sender.
// If this is not the case, regenerate this file with moq.
var _ sender = &senderMock{}

// senderMock is a mock implementation of sender.
//
// 	func TestSomethingThatUsessender(t *testing.T) {
//
// 		// make and configure a mocked sender
// 		mockedsender := &senderMock{
// 			SendFunc: func(contextMoqParam context.Context, reader io.Reader) error {
// 				panic("mock out the Send method")
// 			},
// 		}
//
// 		// use mockedsender in code that requires sender
// 		// and then make assertions.
//
// 	}
type senderMock struct {
	// SendFunc mocks the Send method.
	SendFunc func(contextMoqParam context.Context, reader io.Reader) error

	// calls tracks calls to the methods.
	calls struct {
		// Send holds details about calls to the Send method.
		Send []struct {
			// ContextMoqParam is the contextMoqParam argument value.
			ContextMoqParam context.Context
			// Reader is the reader argument value.
			Reader io.Reader
		}
	}
	lockSend sync.RWMutex
}

// Send calls SendFunc.
func (mock *senderMock) Send(contextMoqParam context.Context, reader io.Reader) error {
	if mock.SendFunc == nil {
		panic("senderMock.SendFunc: method is nil but sender.Send was just called")
	}
	callInfo := struct {
		ContextMoqParam context.Context
		Reader          io.Reader
	}{
		ContextMoqParam: contextMoqParam,
		Reader:          reader,
	}
	mock.lockSend.Lock()
	mock.calls.Send = append(mock.calls.Send, callInfo)
	mock.lockSend.Unlock()
	return mock.SendFunc(contextMoqParam, reader)
}

// SendCalls gets all the calls that were made to Send.
// Check the length with:
//     len(mockedsender.SendCalls())
func (mock *senderMock) SendCalls() []struct {
	ContextMoqParam context.Context
	Reader          io.Reader
} {
	var calls []struct {
		ContextMoqParam context.Context
		Reader          io.Reader
	}
	mock.lockSend.RLock()
	calls = mock.calls.Send
	mock.lockSend.RUnlock()
	return calls
}

// Ensure, that limiterMock does implement limiter.
// If this is not the case, regenerate this file with moq.
var _ limiter = &limiterMock{}

// limiterMock is a mock implementation of limiter.
//
// 	func TestSomethingThatUseslimiter(t *testing.T) {
//
// 		// make and configure a mocked limiter
// 		mockedlimiter := &limiterMock{
// 			WaitFunc: func(contextMoqParam context.Context) error {
// 				panic("mock out the Wait method")
// 			},
// 		}
//
// 		// use mockedlimiter in code that requires limiter
// 		// and then make assertions.
//
// 	}
type limiterMock struct {
	// WaitFunc mocks the Wait method.
	WaitFunc func(contextMoqParam context.Context) error

	// calls tracks calls to the methods.
	calls struct {
		// Wait holds details about calls to the Wait method.
		Wait []struct {
			// ContextMoqParam is the contextMoqParam argument value.
			ContextMoqParam context.Context
		}
	}
	lockWait sync.RWMutex
}

// Wait calls WaitFunc.
func (mock *limiterMock) Wait(contextMoqParam context.Context) error {
	if mock.WaitFunc == nil {
		panic("limiterMock.WaitFunc: method is nil but limiter.Wait was just called")
	}
	callInfo := struct {
		ContextMoqParam context.Context
	}{
		ContextMoqParam: contextMoqParam,
	}
	mock.lockWait.Lock()
	mock.calls.Wait = append(mock.calls.Wait, callInfo)
	mock.lockWait.Unlock()
	return mock.WaitFunc(contextMoqParam)
}

// WaitCalls gets all the calls that were made to Wait.
// Check the length with:
//     len(mockedlimiter.WaitCalls())
func (mock *limiterMock) WaitCalls() []struct {
	ContextMoqParam context.Context
} {
	var calls []struct {
		ContextMoqParam context.Context
	}
	mock.lockWait.RLock()
	calls = mock.calls.Wait
	mock.lockWait.RUnlock()
	return calls
}