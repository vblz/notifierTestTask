// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package main

import (
	"context"
	"io"
	"sync"
)

// Ensure, that scannerMock does implement scanner.
// If this is not the case, regenerate this file with moq.
var _ scanner = &scannerMock{}

// scannerMock is a mock implementation of scanner.
//
// 	func TestSomethingThatUsesscanner(t *testing.T) {
//
// 		// make and configure a mocked scanner
// 		mockedscanner := &scannerMock{
// 			ReadLineFunc: func(contextMoqParam context.Context) ([]byte, bool) {
// 				panic("mock out the ReadLine method")
// 			},
// 		}
//
// 		// use mockedscanner in code that requires scanner
// 		// and then make assertions.
//
// 	}
type scannerMock struct {
	// ReadLineFunc mocks the ReadLine method.
	ReadLineFunc func(contextMoqParam context.Context) ([]byte, bool)

	// calls tracks calls to the methods.
	calls struct {
		// ReadLine holds details about calls to the ReadLine method.
		ReadLine []struct {
			// ContextMoqParam is the contextMoqParam argument value.
			ContextMoqParam context.Context
		}
	}
	lockReadLine sync.RWMutex
}

// ReadLine calls ReadLineFunc.
func (mock *scannerMock) ReadLine(contextMoqParam context.Context) ([]byte, bool) {
	if mock.ReadLineFunc == nil {
		panic("scannerMock.ReadLineFunc: method is nil but scanner.ReadLine was just called")
	}
	callInfo := struct {
		ContextMoqParam context.Context
	}{
		ContextMoqParam: contextMoqParam,
	}
	mock.lockReadLine.Lock()
	mock.calls.ReadLine = append(mock.calls.ReadLine, callInfo)
	mock.lockReadLine.Unlock()
	return mock.ReadLineFunc(contextMoqParam)
}

// ReadLineCalls gets all the calls that were made to ReadLine.
// Check the length with:
//     len(mockedscanner.ReadLineCalls())
func (mock *scannerMock) ReadLineCalls() []struct {
	ContextMoqParam context.Context
} {
	var calls []struct {
		ContextMoqParam context.Context
	}
	mock.lockReadLine.RLock()
	calls = mock.calls.ReadLine
	mock.lockReadLine.RUnlock()
	return calls
}

// Ensure, that notifierMock does implement notifier.
// If this is not the case, regenerate this file with moq.
var _ notifier = &notifierMock{}

// notifierMock is a mock implementation of notifier.
//
// 	func TestSomethingThatUsesnotifier(t *testing.T) {
//
// 		// make and configure a mocked notifier
// 		mockednotifier := &notifierMock{
// 			NotifyFunc: func(contextMoqParam context.Context, reader io.Reader) <-chan error {
// 				panic("mock out the Notify method")
// 			},
// 		}
//
// 		// use mockednotifier in code that requires notifier
// 		// and then make assertions.
//
// 	}
type notifierMock struct {
	// NotifyFunc mocks the Notify method.
	NotifyFunc func(contextMoqParam context.Context, reader io.Reader) <-chan error

	// calls tracks calls to the methods.
	calls struct {
		// Notify holds details about calls to the Notify method.
		Notify []struct {
			// ContextMoqParam is the contextMoqParam argument value.
			ContextMoqParam context.Context
			// Reader is the reader argument value.
			Reader io.Reader
		}
	}
	lockNotify sync.RWMutex
}

// Notify calls NotifyFunc.
func (mock *notifierMock) Notify(contextMoqParam context.Context, reader io.Reader) <-chan error {
	if mock.NotifyFunc == nil {
		panic("notifierMock.NotifyFunc: method is nil but notifier.Notify was just called")
	}
	callInfo := struct {
		ContextMoqParam context.Context
		Reader          io.Reader
	}{
		ContextMoqParam: contextMoqParam,
		Reader:          reader,
	}
	mock.lockNotify.Lock()
	mock.calls.Notify = append(mock.calls.Notify, callInfo)
	mock.lockNotify.Unlock()
	return mock.NotifyFunc(contextMoqParam, reader)
}

// NotifyCalls gets all the calls that were made to Notify.
// Check the length with:
//     len(mockednotifier.NotifyCalls())
func (mock *notifierMock) NotifyCalls() []struct {
	ContextMoqParam context.Context
	Reader          io.Reader
} {
	var calls []struct {
		ContextMoqParam context.Context
		Reader          io.Reader
	}
	mock.lockNotify.RLock()
	calls = mock.calls.Notify
	mock.lockNotify.RUnlock()
	return calls
}
