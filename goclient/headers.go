package goclient

import (
	"net/http"
)

// These constants represent standard HTTP field names and field name values.
// This section should only be modified when introducing a new field name or
// removing an existing field name entirely. Unlike the values for field names,
// some values can be arbitrary for field name values.
const (
	HeaderAccept        = "Accept"
	HeaderAuthorization = "Authorization"
	HeaderContentLength = "Content-Length"
	HeaderContentType   = "Content-Type"
	HeaderUserAgent     = "User-Agent"
	HeaderKeepAlive     = "Connection"

	ContentTypeJson  = "application/json"
	DefaultUserAgent = "go-http"
	DefaultKeepAlive = "Keep-Alive"
)

// getRequestHeaders returns request headers that are defined as part of a
// client request. A nil return value is the equivalent of an empty map.
func getRequestHeaders(headers ...http.Header) http.Header {
	if len(headers) > 0 {
		return headers[0]
	}
	return nil
}

// joinHeaders returns a map of consolidated HTTP headers defined as part of a
// client build or a client request. If a header has been defined in both, the
// client request takes precedence. It also ensures a name for the user agent
// is set.
func (c *client) joinRequestHeaders(headers http.Header) http.Header {
	h := make(http.Header)

	// Set client build headers.
	for key, value := range c.builder.headers {
		h.Set(key, value[0])
	}

	// Set client request headers. If a header is also defined as part of a
	// client build, its value will be replaced with the one beneath.
	for key, value := range headers {
		h.Set(key, value[0])
	}

	// Set the name of the user agent. The default value is used if User-Agent
	// is not defined as part of a client build or a client request.
	if h.Get(HeaderUserAgent) == "" {
		if c.builder.userAgent == "" {
			h.Set(HeaderUserAgent, DefaultUserAgent)
			return h
		}
		h.Set(HeaderUserAgent, c.builder.userAgent)
	}
	return h
}
