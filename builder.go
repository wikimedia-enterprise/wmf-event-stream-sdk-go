package eventstream

import (
	"net/http"
	"time"
)

// NewBuilder create new builder instance
func NewBuilder() *ClientBuilder {
	cb := new(ClientBuilder)
	cb.client = NewClient()
	return cb
}

// ClientBuilder create new client with configuration
type ClientBuilder struct {
	client *Client
}

// URL set base URL for the client
func (cb *ClientBuilder) URL(url string) *ClientBuilder {
	cb.client.url = url
	return cb
}

// HTTPClient provide custom http client
func (cb *ClientBuilder) HTTPClient(client *http.Client) *ClientBuilder {
	cb.client.httpClient = client
	return cb
}

// BackoffTime set backoff time for client
func (cb *ClientBuilder) BackoffTime(backoffTime time.Duration) *ClientBuilder {
	cb.client.backoffTime = backoffTime
	return cb
}

// Options set client urls
func (cb *ClientBuilder) Options(options *Options) *ClientBuilder {
	cb.client.options = options
	return cb
}

// Build create new client with provided configuration
func (cb *ClientBuilder) Build() *Client {
	return cb.client
}
