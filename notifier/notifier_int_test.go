package notifier

import (
	"bytes"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"
)

func TestNotifier(t *testing.T) {
	const testContentType = "testContentType"
	times := 0
	var lastRequestTime time.Time
	dataReceived := make([][]byte, 0)
	s := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if times != 0 {
			assert.GreaterOrEqual(t, time.Since(lastRequestTime).Seconds(), 0.99)
		}
		lastRequestTime = time.Now()
		times++
		b, err := io.ReadAll(request.Body)
		assert.NoError(t, err)
		dataReceived = append(dataReceived, b)
	}))
	defer s.Close()

	u, err := url.Parse(s.URL)
	require.NoError(t, err)
	n := New(u, ContentType(testContentType), RequestsPerSecond(1))
	dataSent := [][]byte{
		[]byte("test1"),
		{},
		[]byte("test2"),
	}

	messages := make([]io.Reader, len(dataSent))
	for i, v := range dataSent {
		messages[i] = bytes.NewBuffer(v)
	}

	errCh := n.NotifyMultiple(context.Background(), messages)

	for range errCh {
		assert.Fail(t, "should not be any error")
	}

	assert.ElementsMatch(t, dataSent, dataReceived)
}
