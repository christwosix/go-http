package goclient

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockResponse struct {
	Want string `json:"Want"`
}

func TestBytesBody(t *testing.T) {
	r := &Response{Body: []byte("foobar")}
	have := r.BytesBody()
	assert.Equal(t, []byte{102, 111, 111, 98, 97, 114}, have)
}

func TestStringBody(t *testing.T) {
	r := &Response{Body: []byte("foobar")}
	have := r.StringBody()
	assert.Equal(t, "foobar", have)
}

func TestUnmarshalJson(t *testing.T) {
	var jsonData mockResponse
	r := &Response{Body: []byte(`{"Want": "foobar"}`)}
	r.UnmarshalJson(&jsonData)
	assert.Equal(t, "foobar", jsonData.Want)
}
