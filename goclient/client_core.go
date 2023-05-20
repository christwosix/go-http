package goclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	defaultConnectionTimeout   = 15 * time.Second
	defaultResponseTimeout     = 15 * time.Second
	defaultMaxIdleConnsPerHost = 2
)

// getBaseURL returns the base URL of a service or an empty string.
func (c *client) getBaseURL() (string, error) {
	if c.builder.baseURL != "" {
		if _, err := url.ParseRequestURI(c.builder.baseURL); err != nil {
			return "", err
		}
		return c.builder.baseURL, nil
	}
	return "", nil
}

// getConnectionTimeout returns the desired or default connection timeout.
func (c *client) getConnectionTimeout() time.Duration {
	if c.builder.connectionTimeout > 0 {
		return c.builder.connectionTimeout
	}
	return defaultConnectionTimeout
}

// getMaxIdleConnsPerHost returns the desired or default number of maximum idle
// connections per host.
func (c *client) getMaxIdleConnsPerHost() int {
	if c.builder.maxIdleConnsPerHost > 0 {
		return c.builder.maxIdleConnsPerHost
	}
	return defaultMaxIdleConnsPerHost
}

// getResponseTimeout returns the desired or default response timeout.
func (c *client) getResponseTimeout() time.Duration {
	if c.builder.responseTimeout > 0 {
		return c.builder.responseTimeout
	}
	return defaultResponseTimeout
}

// getRequestBody returns the content type encoding of the request body.
func (c *client) getRequestBody(contentType string, body any) ([]byte, error) {
	if body == nil {
		return nil, nil
	}
	// TODO: Add support for other content types.
	switch strings.ToLower(contentType) {
	case "application/json":
		b, err := json.Marshal(body)
		return b, err
	default:
		b, err := json.Marshal(body)
		return b, err
	}
}

// getClient returns a custom HTTP client with the desired configurations. It
// is resuable making it concurrent safe with goroutines.
func (c *client) getClient() *http.Client {
	c.initOnce.Do(func() {
		c.client = &http.Client{
			Timeout: c.getConnectionTimeout() + c.getResponseTimeout(),
			Transport: &http.Transport{
				MaxIdleConnsPerHost:   c.getMaxIdleConnsPerHost(),
				ResponseHeaderTimeout: c.getResponseTimeout(),
				DialContext: (&net.Dialer{
					Timeout: c.getConnectionTimeout(),
				}).DialContext,
			},
		}
	})
	return c.client
}

// doRequest calls Do from the standard library to perform HTTP requests. It
// also handles the low-level plumbing such as building the request, using the
// custom HTTP client, and returning the response.
func (c *client) doRequest(method, endpoint string, headers http.Header, body any) (*Response, error) {
	baseURL, err := c.getBaseURL()
	if err != nil {
		return nil, err
	}

	// If baseURL is nil, it must be parsed to the endpoint parameter.
	// Otherwise, an unsupported protocol scheme error will be thrown by the
	// HTTP client when it attempts to perform the request.
	requestURL := fmt.Sprintf(baseURL + endpoint)
	requestHeaders := c.joinRequestHeaders(headers)
	requestBody, err := c.getRequestBody(requestHeaders.Get(HeaderContentType), body)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(method, requestURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	request.Header = requestHeaders

	response, err := c.getClient().Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	responseData := Response{
		Body:            responseBody,
		Status:          response.Status,
		StatusCode:      response.StatusCode,
		ResponseHeaders: response.Header,
	}
	return &responseData, nil
}
