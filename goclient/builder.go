package goclient

import (
	"net/http"
	"time"
)

// Builder provides the interface for custom HTTP implementations.
type Builder interface {
	Build() Client
	SetBaseURL(baseUrl string) Builder
	SetRequestHeaders(headers http.Header) Builder
	SetConnectionTimeout(timeout time.Duration) Builder
	SetResponseTimeout(timeout time.Duration) Builder
	SetUserAgent(name string) Builder
}

// builder provides configuration options for custom HTTP implementations.
type builder struct {
	baseURL             string
	userAgent           string
	headers             http.Header
	responseTimeout     time.Duration
	connectionTimeout   time.Duration
	maxIdleConnsPerHost int
}

// NewBuild provides a custom HTTP builder implementation.
func NewBuild() Builder {
	return &builder{}
}

// Build provides a custom HTTP client implementation with the desired config.
func (c *builder) Build() Client {
	return &client{builder: c}
}

// SetBaseURL sets the base URL of a web service. If specified, it is
// concatenated with the endpoint parameter value for every request. Also, the
// endpoint parameter only requires the API resource. Otherwise, the absolute
// URI must be supplied.
func (c *builder) SetBaseURL(baseUrl string) Builder {
	c.baseURL = baseUrl
	return c
}

// SetRequestHeaders sets request headers defined as part of the HTTP build.
func (c *builder) SetRequestHeaders(headers http.Header) Builder {
	c.headers = headers
	return c
}

// SetMaxIdleConnsPerHost sets the max number of idle connections per host.
func (c *builder) SetMaxIdleConnsPerHost(maxIdleConnsPerHost int) Builder {
	c.maxIdleConnsPerHost = maxIdleConnsPerHost
	return c
}

// SetConnectionTimeout sets the max duration that the HTTP client will wait
// for a connection to complete.
func (c *builder) SetConnectionTimeout(timeout time.Duration) Builder {
	c.connectionTimeout = timeout
	return c
}

// SetResponseTimeout sets the max duration that the HTTP client will wait for
// the response headers.
func (c *builder) SetResponseTimeout(timeout time.Duration) Builder {
	c.responseTimeout = timeout
	return c
}

// SetUserAgent sets the value for the User-Agent request header. This is used
// if not defined as part of the HTTP build or client request.
func (c *builder) SetUserAgent(name string) Builder {
	c.userAgent = name
	return c
}
