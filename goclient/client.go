package goclient

import (
	"net/http"
	"sync"
)

// client provides the implementation of a custom HTTP client.
type client struct {
	builder  *builder
	client   *http.Client
	initOnce sync.Once
}

// Client provides the interface for a custom HTTP client.
type Client interface {
	Get(endpoint string, headers ...http.Header) (*Response, error)
	Put(endpoint string, body any, headers ...http.Header) (*Response, error)
	Post(endpoint string, body any, headers ...http.Header) (*Response, error)
	Patch(endpoint string, body any, headers ...http.Header) (*Response, error)
	Delete(endpoint string, headers ...http.Header) (*Response, error)
}

// Get issues a GET request to the specified URL.
func (c *client) Get(endpoint string, headers ...http.Header) (*Response, error) {
	return c.doRequest(http.MethodGet, endpoint, getRequestHeaders(headers...), nil)
}

// Put issues a PUT request to the specified URL.
func (c *client) Put(endpoint string, body any, headers ...http.Header) (*Response, error) {
	return c.doRequest(http.MethodPut, endpoint, getRequestHeaders(headers...), body)
}

// Post issues a POST request to the specified URL.
func (c *client) Post(endpoint string, body any, headers ...http.Header) (*Response, error) {
	return c.doRequest(http.MethodPost, endpoint, getRequestHeaders(headers...), body)
}

// Patch issues a PATCH request to the specified URL.
func (c *client) Patch(endpoint string, body any, headers ...http.Header) (*Response, error) {
	return c.doRequest(http.MethodPatch, endpoint, getRequestHeaders(headers...), body)
}

// Delete issues a DELETE request to the specified URL.
func (c *client) Delete(endpoint string, headers ...http.Header) (*Response, error) {
	return c.doRequest(http.MethodDelete, endpoint, getRequestHeaders(headers...), nil)
}
