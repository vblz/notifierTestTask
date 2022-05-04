package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"syscall"
	"testing"
	"time"
)

var testLines = []string{
	"test1",
	"",
	"test2",
	"test3",
}

func TestApp(t *testing.T) {
	callNum := 0

	lastCallNow := time.Now().Add(-time.Second)
	s := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		b, err := io.ReadAll(request.Body)
		require.NoError(t, err)
		assert.Equal(t, testLines[callNum], string(b))
		callNum++

		assert.Greater(t, time.Since(lastCallNow), time.Millisecond * 990)
		lastCallNow = time.Now()
	}))
	defer s.Close()

	oldArgs := os.Args
	defer func() {
		os.Args = oldArgs
	}()
	os.Args = []string{
		"executable",
		"--url",
		s.URL,
		"-i",
		"1s",
	}

	tearDown := mockStdinWithData(t, strings.Join(testLines, "\n"))
	defer tearDown()

	main()
	assert.Equal(t, len(testLines), callNum)
}

func TestApp_SigInt(t *testing.T) {
	callNum := 0
	s := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		callNum++
	}))
	defer s.Close()

	oldArgs := os.Args
	defer func() {
		os.Args = oldArgs
	}()
	os.Args = []string{
		"executable",
		"--url",
		s.URL,
		"-i",
		"1s",
	}

	tearDown := mockStdinWithData(t, strings.Join(testLines, "\n"))
	defer tearDown()

	go func() {
		time.Sleep(time.Millisecond * 20)
		err := syscall.Kill(syscall.Getpid(), syscall.SIGINT)
		assert.NoError(t, err)
	}()
	main()
	assert.Greater(t, len(testLines), callNum)
}

// mockStdinWithData changes stdin for file with data. Returns a tear down func, restoring stdIn.
func mockStdinWithData(t *testing.T, data string) func() {
	original := os.Stdin

	read, write, err := os.Pipe()
	require.NoError(t, err)
	_, err = write.WriteString(data)
	require.NoError(t, err)
	err = write.Close()
	require.NoError(t, err)

	os.Stdin = read

	return func() {
		os.Stdin = original
	}
}
