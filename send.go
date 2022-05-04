package main

//go:generate moq -out send_moq_test.go . scanner notifier

import (
	"bytes"
	"context"
	"github.com/go-pkgz/lgr"
	"io"
	"sync"
)

type scanner interface {
	ReadLine(context.Context) ([]byte, bool)
}

type notifier interface {
	Notify(context.Context, io.Reader) <-chan error
}

func sendMessages(ctx context.Context, n notifier, s scanner, l lgr.L) int {
	wg := sync.WaitGroup{}
	messagesProcessed := 0

	for {
		message, ok := s.ReadLine(ctx)
		if !ok {
			break
		}

		messagesProcessed++
		wg.Add(1)

		errCh := n.Notify(ctx, bytes.NewBuffer(message))

		go func(errCh <-chan error) {
			defer wg.Done()
			for err := range errCh {
				l.Logf("ERROR can't perform request: %s", err)
			}
		}(errCh)
	}

	wg.Wait()

	return messagesProcessed
}
