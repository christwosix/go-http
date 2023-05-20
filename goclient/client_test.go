package goclient

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type mockClient struct {
	Name     string `json:"Name,omitempty"`
	Response string `json:"Response,omitempty"`
}

func NewClient() Client {
	headers := make(http.Header)
	headers.Set(HeaderContentType, ContentTypeJson)
	headers.Set(HeaderUserAgent, DefaultUserAgent)

	return NewBuild().
		SetRequestHeaders(headers).
		SetUserAgent(DefaultUserAgent).
		Build()
}

func TestGet(t *testing.T) {
	t.Run("SuccessfulGet", func(t *testing.T) {
		s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
			assert.Equal(t, "go-http", r.UserAgent())

			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{ "Response": "OK" }`))
		}))
		defer s.Close()

		c := NewClient()

		response, err := c.Get(s.URL, nil)
		var jsonData mockClient
		err = response.UnmarshalJson(&jsonData)
		assert.Equal(t, http.StatusOK, response.StatusCode)
		assert.Equal(t, "OK", jsonData.Response)
		require.NoError(t, err, "expected no errors")
	})
}

func TestPut(t *testing.T) {
	t.Run("SuccessfulPut", func(t *testing.T) {
		s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPut, r.Method)
			assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
			assert.Equal(t, "go-http", r.UserAgent())

			requestBody, err := io.ReadAll(r.Body)
			assert.Equal(t, `{"Name":"foobar"}`, string(requestBody))
			require.NoError(t, err, "expected no errors")

			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{ "Response": "OK" }`))
		}))
		defer s.Close()

		c := NewClient()

		response, err := c.Put(s.URL, mockClient{Name: "foobar"}, nil)
		var jsonData mockClient
		err = response.UnmarshalJson(&jsonData)
		assert.Equal(t, http.StatusOK, response.StatusCode)
		assert.Equal(t, "OK", jsonData.Response)
		require.NoError(t, err, "expected no errors")
	})
}

func TestPost(t *testing.T) {
	t.Run("SuccessfulPost", func(t *testing.T) {
		s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
			assert.Equal(t, "go-http", r.UserAgent())

			requestBody, err := io.ReadAll(r.Body)
			assert.Equal(t, `{"Name":"foobar"}`, string(requestBody))
			require.NoError(t, err, "expected no errors")

			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{ "Response": "OK" }`))
		}))
		defer s.Close()

		c := NewClient()

		response, err := c.Post(s.URL, mockClient{Name: "foobar"}, nil)
		var jsonData mockClient
		err = response.UnmarshalJson(&jsonData)
		assert.Equal(t, http.StatusOK, response.StatusCode)
		assert.Equal(t, "OK", jsonData.Response)
		require.NoError(t, err, "expected no errors")
	})
}

func TestPatch(t *testing.T) {
	t.Run("SuccessfulPatch", func(t *testing.T) {
		s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPatch, r.Method)
			assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
			assert.Equal(t, "go-http", r.UserAgent())

			requestBody, err := io.ReadAll(r.Body)
			assert.Equal(t, `{"Name":"foobar"}`, string(requestBody))
			require.NoError(t, err, "expected no errors")

			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{ "Response": "OK" }`))
		}))
		defer s.Close()

		c := NewClient()

		response, err := c.Patch(s.URL, mockClient{Name: "foobar"}, nil)
		var jsonData mockClient
		err = response.UnmarshalJson(&jsonData)
		assert.Equal(t, http.StatusOK, response.StatusCode)
		assert.Equal(t, "OK", jsonData.Response)
		require.NoError(t, err, "expected no errors")
	})
}

func TestDelete(t *testing.T) {
	t.Run("SuccessfulDelete", func(t *testing.T) {
		s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodDelete, r.Method)
			assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
			assert.Equal(t, "go-http", r.UserAgent())

			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{ "Response": "OK" }`))
		}))
		defer s.Close()

		c := NewClient()

		response, err := c.Delete(s.URL, nil)
		var jsonData mockClient
		err = response.UnmarshalJson(&jsonData)
		assert.Equal(t, http.StatusOK, response.StatusCode)
		assert.Equal(t, "OK", jsonData.Response)
		require.NoError(t, err, "expected no errors")
	})
}
