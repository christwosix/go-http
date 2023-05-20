package goclient

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type mockCore struct {
	A string `json:"A"`
	B string `json:"B"`
}

func TestGetBaseURL(t *testing.T) {
	tt := []struct {
		name     string
		build    *builder
		expect   any
		hasError bool
	}{
		{
			name:   "HasURL",
			build:  &builder{baseURL: "https://foobar.com"},
			expect: "https://foobar.com",
		},
		{
			name:   "NoURL",
			build:  &builder{},
			expect: "",
		},
		{
			name:     "InvalidScheme",
			build:    &builder{baseURL: "foobar.com"},
			expect:   "",
			hasError: true,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			c := &client{builder: tc.build}
			have, err := c.getBaseURL()
			if tc.hasError {
				assert.Error(t, err)
				assert.Empty(t, have, "body should be nil")
				return
			}
			assert.Equal(t, tc.expect, have)
			require.NoError(t, err, "expected no errors")
		})
	}
}

func TestGetConnectionTimeout(t *testing.T) {
	tt := []struct {
		name   string
		build  *builder
		expect any
	}{
		{
			name:   "CustomTimeout",
			build:  &builder{connectionTimeout: 30 * time.Second},
			expect: 30 * time.Second,
		},
		{
			name:   "DefaultTimeout",
			build:  &builder{},
			expect: 15 * time.Second,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			c := &client{builder: tc.build}
			assert.Equal(t, tc.expect, c.getConnectionTimeout())
		})
	}
}

func TestGetMaxIdleConnsPerHost(t *testing.T) {
	tt := []struct {
		name   string
		build  *builder
		expect any
	}{
		{
			name:   "CustomConnections",
			build:  &builder{maxIdleConnsPerHost: 5},
			expect: 5,
		},
		{
			name:   "DefaultConnections",
			build:  &builder{},
			expect: 2,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			c := &client{builder: tc.build}
			assert.Equal(t, tc.expect, c.getMaxIdleConnsPerHost())
		})
	}
}

func TestGetResponseTimeout(t *testing.T) {
	tt := []struct {
		name   string
		build  *builder
		expect any
	}{
		{
			name:   "CustomTimeout",
			build:  &builder{responseTimeout: 20 * time.Second},
			expect: 20 * time.Second,
		},
		{
			name:   "DefaultTimeout",
			build:  &builder{},
			expect: 15 * time.Second,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			c := &client{builder: tc.build}
			assert.Equal(t, tc.expect, c.getResponseTimeout())
		})
	}
}

func TestGetRequestBody(t *testing.T) {
	tt := []struct {
		name        string
		build       *builder
		contentType string
		body        any
		expect      []byte
		hasError    bool
	}{
		{
			name:        "NilBody",
			contentType: ContentTypeJson,
			body:        nil,
			expect:      nil,
		},
		{
			name:        "ApplicationJson",
			contentType: ContentTypeJson,
			body:        &mockCore{A: "foo", B: "bar"},
			expect:      []byte(`{"A":"foo","B":"bar"}`),
		},
		{
			name:        "DefaultCase",
			contentType: "",
			body:        &mockCore{A: "foo", B: "bar"},
			expect:      []byte(`{"A":"foo","B":"bar"}`),
		},
		{
			name:        "MarshalError",
			contentType: ContentTypeJson,
			body:        make(chan int),
			hasError:    true,
			expect:      nil,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			c := &client{}
			body, err := c.getRequestBody(tc.contentType, tc.body)
			if tc.hasError {
				assert.Error(t, err)
				assert.Empty(t, body, "body should be nil")
				return
			}
			assert.Equal(t, tc.expect, body)
			require.NoError(t, err, "expected no errors")
		})
	}
}

func TestGetClient(t *testing.T) {
	c := &client{builder: &builder{}}
	assert.IsType(t, &http.Client{}, c.getClient())
}

func TestDoRequest(t *testing.T) {
	tt := []struct {
		name        string
		body        any
		url         string
		method      string
		headers     http.Header
		hasError    bool
		mockServer  bool
		mockHandler func(w http.ResponseWriter, r *http.Request)
	}{
		{
			name:     "BaseURLError",
			body:     "",
			url:      "foobar.com",
			method:   "",
			hasError: true,
		},
		{
			name:     "RequestBodyError",
			body:     make(chan int),
			url:      "",
			method:   http.MethodGet,
			hasError: true,
		},
		{
			name:     "NewRequestError",
			body:     nil,
			url:      "",
			method:   "*?",
			hasError: true,
		},
		{
			name:     "ResponseError",
			body:     nil,
			url:      "",
			method:   http.MethodGet,
			hasError: true,
		},
		{
			name:       "ResponseBodyError",
			body:       nil,
			url:        "",
			method:     http.MethodGet,
			headers:    http.Header{},
			hasError:   true,
			mockServer: true,
			mockHandler: func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set(HeaderContentLength, "1")
			},
		},
		{
			name:       "SuccessfulResponse",
			body:       nil,
			url:        "",
			method:     http.MethodGet,
			headers:    http.Header{HeaderContentType: {ContentTypeJson}},
			hasError:   false,
			mockServer: true,
			mockHandler: func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, "application/json", r.Header.Get("Content-Type"))

				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"A":"foo","B":"bar"}`))
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if tc.mockServer {
				s := httptest.NewServer(http.HandlerFunc(tc.mockHandler))
				tc.url = s.URL
				defer s.Close()
			}

			c := &client{builder: &builder{baseURL: tc.url}}
			response, err := c.doRequest(tc.method, "/api", tc.headers, tc.body)

			if tc.hasError {
				assert.Error(t, err)
				assert.Empty(t, response, "response should be nil")
				return
			}

			var jsonData mockCore
			err = response.UnmarshalJson(&jsonData)
			assert.Equal(t, http.StatusOK, response.StatusCode)
			assert.Equal(t, "foo", jsonData.A)
			assert.Equal(t, "bar", jsonData.B)
			require.NoError(t, err, "expected no errors")
		})
	}
}
