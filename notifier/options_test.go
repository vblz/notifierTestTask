package notifier

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLgr(t *testing.T) {
	o := Options{
		logger: noOpLogger,
	}
	assert.NotNil(t, o.logger)

	Lgr(nil)(&o)
	assert.Nil(t, o.logger)
}

func TestRetryNumber(t *testing.T) {
	const number = 7
	o := Options{}
	RetryNumber(number)(&o)

	assert.Equal(t, number, o.retryNumber)
}

func TestContentType(t *testing.T) {
	const contentType = "testContentType"
	o := Options{}
	ContentType(contentType)(&o)

	assert.Equal(t, contentType, o.contentType)
}

func TestRequestsPerSecond(t *testing.T) {
	const number = 155.5
	o := Options{}
	RequestsPerSecond(number)(&o)

	assert.Equal(t, number, o.requestsPerSecond)
}

func TestDefaultOptions(t *testing.T) {
	o := defaultOptions()
	assert.NotNil(t, o)
	assert.Equal(t, 3, o.retryNumber)
	assert.Equal(t, "application/octet-stream", o.contentType)
	assert.Equal(t, float64(100), o.requestsPerSecond)
}

func TestCreateOptions(t *testing.T) {
	o := createOptions([]Option{
		RequestsPerSecond(1),
	})
	assert.NotNil(t, o)
	assert.Equal(t, 3, o.retryNumber)
	assert.Equal(t, "application/octet-stream", o.contentType)
	assert.Equal(t, float64(1), o.requestsPerSecond)
}
