package goclient

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewBuild(t *testing.T) {
	b := NewBuild()
	assert.IsType(t, &builder{}, b)
}

func TestBuild(t *testing.T) {
	c := NewBuild().Build()
	assert.IsType(t, &client{}, c)
}

func TestSetBaseURL(t *testing.T) {
	t.Run("ValidBaseUrl", func(t *testing.T) {
		b := &builder{}
		have := b.SetBaseURL("https://foobar.com")
		assert.Equal(t, "https://foobar.com", b.baseURL)
		assert.IsType(t, &builder{}, have)
	})
}

func TestSetRequestHeaders(t *testing.T) {
	tt := []struct {
		name    string
		headers http.Header
		expect  any
	}{
		{
			name:    "HasHeaders",
			headers: http.Header{HeaderContentType: {ContentTypeJson}},
			expect:  "application/json",
		},
		{
			name:    "NilHeaders",
			headers: nil,
			expect:  "",
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			b := &builder{}
			have := b.SetRequestHeaders(tc.headers)
			assert.Equal(t, tc.expect, b.headers.Get(HeaderContentType))
			assert.IsType(t, &builder{}, have)
		})
	}
}

func TestSetMaxIdleConnsPerHost(t *testing.T) {
	b := &builder{}
	have := b.SetMaxIdleConnsPerHost(3)
	assert.Equal(t, b.maxIdleConnsPerHost, b.maxIdleConnsPerHost)
	assert.IsType(t, &builder{}, have)
}

func TestSetConnectionTimeout(t *testing.T) {
	b := &builder{}
	have := b.SetConnectionTimeout(30 * time.Second)
	assert.Equal(t, 30*time.Second, b.connectionTimeout)
	assert.IsType(t, &builder{}, have)
}

func TestSetResponseTimeout(t *testing.T) {
	b := &builder{}
	have := b.SetResponseTimeout(12 * time.Second)
	assert.Equal(t, 12*time.Second, b.responseTimeout)
	assert.IsType(t, &builder{}, have)
}

func TestSetUserAgent(t *testing.T) {
	b := &builder{}
	have := b.SetUserAgent(DefaultUserAgent)
	assert.Equal(t, "go-http", b.userAgent)
	assert.IsType(t, &builder{}, have)
}
