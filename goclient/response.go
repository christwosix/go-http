package goclient

import (
	"encoding/json"
	"net/http"
)

// Response represents the objects returned by a web service in response to a
// client request.
type Response struct {
	Body            []byte
	Status          string
	StatusCode      int
	ResponseHeaders http.Header
}

// BytesBody returns the byte slice of a response body.
func (r *Response) BytesBody() []byte {
	return r.Body
}

// StringBody converts the byte slice of a response body to a string and
// returns it.
func (r *Response) StringBody() string {
	return string(r.Body)
}

// UnmarshalJson uses the json package from the standard library to return
// the JSON-decoded data which is stored in the value pointed to by target.
//
// It abstracts the need for the calling application to manage package
// dependencies, resulting in concise code.
func (r *Response) UnmarshalJson(target any) error {
	return json.Unmarshal(r.BytesBody(), target)
}
