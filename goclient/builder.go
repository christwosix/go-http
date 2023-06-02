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
func (b *builder) Build() Client {
	return &client{builder: b}
}

// SetBaseURL sets the base URL of a web service. If specified, it is
// concatenated with the endpoint parameter value for every request. Also, the
// endpoint parameter only requires the API resource. Otherwise, the absolute
// URI must be supplied.
func (b *builder) SetBaseURL(baseUrl string) Builder {
	b.baseURL = baseUrl
	return b
}

// SetRequestHeaders sets request headers defined as part of the HTTP build.
func (b *builder) SetRequestHeaders(headers http.Header) Builder {
	b.headers = headers
	return b
}

// SetMaxIdleConnsPerHost sets the max number of idle connections per host.
func (b *builder) SetMaxIdleConnsPerHost(maxIdleConnsPerHost int) Builder {
	b.maxIdleConnsPerHost = maxIdleConnsPerHost
	return b
}

// SetConnectionTimeout sets the max duration that the HTTP client will wait
// for a connection to complete.
func (b *builder) SetConnectionTimeout(timeout time.Duration) Builder {
	b.connectionTimeout = timeout
	return b
}

// SetResponseTimeout sets the max duration that the HTTP client will wait for
// the response headers.
func (b *builder) SetResponseTimeout(timeout time.Duration) Builder {
	b.responseTimeout = timeout
	return b
}

// SetUserAgent sets the value for the User-Agent request header. This is used
// if not defined as part of the HTTP build or client request.
func (b *builder) SetUserAgent(name string) Builder {
	b.userAgent = name
	return b
}
