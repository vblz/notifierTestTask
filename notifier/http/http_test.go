package http

import (
	"bytes"
	"context"
	"errors"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

const testContentType = "testContentType"
const testRetriesMax = 1

func TestCreateHttpTransport(t *testing.T) {
	tr := createHTTPTransport()
	assert.NotNil(t, tr)
	assert.Equal(t, tr.MaxIdleConns, tr.MaxIdleConnsPerHost)
}

func TestCreateHttpClient(t *testing.T) {
	const retriesNumber = 7

	client := createHTTPClient(retriesNumber, nil)
	assert.NotNil(t, client)
	rt, ok := client.Transport.(*retryablehttp.RoundTripper)
	require.True(t, ok)

	assert.Equal(t, retriesNumber, rt.Client.RetryMax)
}

func TestSendInvalid(t *testing.T) {
	c := createClient("http://test.invalid")
	err := c.Send(context.Background(), nil)
	assert.Error(t, err)
}

func TestSend(t *testing.T) {
	var testBytes = []byte("testMessage")
	times := 0
	s := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		times++
		assert.Equal(t, testContentType, request.Header.Get("Content-Type"))
		b, err := io.ReadAll(request.Body)
		require.NoError(t, err)
		assert.Equal(t, testBytes, b)
	}))
	defer s.Close()

	c := createClient(s.URL)
	err := c.Send(context.Background(), bytes.NewBuffer(testBytes))
	assert.NoError(t, err)

	assert.Equal(t, 1, times)
}

func TestSendRetry(t *testing.T) {
	times := 0
	s := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if times == 0 {
			writer.WriteHeader(http.StatusInternalServerError)
		}
		times++
	}))
	defer s.Close()

	c := createClient(s.URL)
	err := c.Send(context.Background(), nil)
	assert.NoError(t, err)

	assert.Equal(t, 2, times)
}

func TestSendError(t *testing.T) {
	times := 0
	s := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		times++
		writer.WriteHeader(http.StatusInternalServerError)
	}))
	defer s.Close()

	c := createClient(s.URL)
	err := c.Send(context.Background(), nil)
	assert.Error(t, err)

	assert.Equal(t, testRetriesMax+1, times)
}

func TestSendStatus400NoRetries(t *testing.T) {
	times := 0
	s := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		times++
		writer.WriteHeader(http.StatusNotFound)
	}))
	defer s.Close()

	c := createClient(s.URL)
	err := c.Send(context.Background(), nil)
	assert.Error(t, err)

	assert.Equal(t, 1, times)
}

func TestSendContext(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		time.Sleep(time.Millisecond * 100)
		writer.WriteHeader(http.StatusNotFound)
	}))
	defer s.Close()

	c := createClient(s.URL)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(time.Millisecond * 20)
		cancel()
	}()

	err := c.Send(ctx, nil)

	assert.Error(t, err)
	assert.True(t, errors.Is(err, ctx.Err()))
}

func createClient(endpoint string) *Client {
	return New(endpoint, testContentType, testRetriesMax, nil)
}
