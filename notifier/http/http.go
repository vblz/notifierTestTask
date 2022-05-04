package http

import (
	"context"
	"fmt"
	"github.com/hashicorp/go-cleanhttp"
	"github.com/hashicorp/go-retryablehttp"
	"io"
	"net/http"
)

// Client with retry policy sending messages to fixed url.
type Client struct {
	endpoint    string
	contentType string
	client      *http.Client
}

// New creates a Client with given endpoint, contentType, logger and retry policy based on retryNumbers.
func New(endpoint, contentType string, retryNumbers int, l LeveledLogger) *Client {
	return &Client{
		endpoint:    endpoint,
		contentType: contentType,
		client:      createHTTPClient(retryNumbers, l),
	}
}

// Send a message with a given context. If response status code is not 2xx, returns an error.
func (c *Client) Send(ctx context.Context, message io.Reader) error {
	r, err := http.NewRequestWithContext(http.MethodPost, c.endpoint, message)
	if err != nil {
		return err
	}

	r.Header.Set("Content-Type", c.contentType)

	resp, err := c.client.Do(r)

	if err != nil {
		return err
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return fmt.Errorf("the response status code is not successful: %d", resp.StatusCode)
	}

	return nil
}

func createHTTPClient(retryNumbers int, l LeveledLogger) *http.Client {
	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = retryNumbers
	retryClient.Logger = l

	retryClient.HTTPClient = &http.Client{
		Transport: createHTTPTransport(),
	}

	standardClient := retryClient.StandardClient()

	return standardClient
}

func createHTTPTransport() *http.Transport {
	res := cleanhttp.DefaultPooledTransport()

	// TODO: think about customization of these values
	// by design we know that all our connections go to the single host
	res.MaxIdleConnsPerHost = res.MaxIdleConns

	return res
}
