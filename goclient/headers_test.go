package goclient

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRequestHeaders(t *testing.T) {
	tt := []struct {
		name    string
		headers []http.Header
		expect  any
	}{
		{
			name:    "NilHeaders",
			headers: nil,
			expect:  nil,
		},
		{
			name:    "HasHeaders",
			headers: []http.Header{map[string][]string{HeaderContentType: {ContentTypeJson}}},
			expect:  "application/json",
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			have := getRequestHeaders(tc.headers...)
			if len(have) == 0 {
				assert.Empty(t, have)
				return
			}
			assert.Equal(t, tc.expect, have.Get(HeaderContentType))
		})
	}
}

func TestJoinHeaders(t *testing.T) {
	tt := []struct {
		name    string
		headers http.Header
		build   *builder
		expect  any
	}{
		{
			name:    "AgentDefault",
			headers: nil,
			build:   &builder{},
			expect:  "go-http",
		},
		{
			name:    "AgentConfig",
			headers: nil,
			build:   &builder{userAgent: "go-config"},
			expect:  "go-config",
		},
		{
			name:    "AgentConfigHeaders",
			headers: nil,
			build: &builder{
				userAgent: "go-config",
				headers:   http.Header{HeaderUserAgent: {"go-build"}},
			},
			expect: "go-build",
		},
		{
			name:    "AgentRequestHeader",
			headers: http.Header{HeaderUserAgent: {"go-request"}},
			build: &builder{
				userAgent: "go-config",
				headers:   http.Header{HeaderUserAgent: {"go-build"}},
			},
			expect: "go-request",
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			c := &client{builder: tc.build}
			have := c.joinRequestHeaders(tc.headers)
			assert.Equal(t, tc.expect, have.Get(HeaderUserAgent))
		})
	}
}
