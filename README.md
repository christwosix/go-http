# Go-HTTP
[![go-test](https://github.com/christwosix/gohttp/actions/workflows/go-test.yml/badge.svg)](https://github.com/christwosix/gohttp/actions/workflows/go-test.yml) [![go-lint](https://github.com/christwosix/gohttp/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/christwosix/gohttp/actions/workflows/golangci-lint.yml)  

A customisable HTTP client written in Go that is lightweight and concurrent safe.

## Getting started

###### Installation
```bash
require github.com/christwosix/gohttp
```

###### Import
```go
import "github.com/christwosix/gohttp/goclient"
```

###### Configuring the client
The HTTP client can be initialised a number of ways. Refer to the GET request example below.
```go
// Define common headers to send with each request.
headers := make(http.Header)
headers.Set("Content-Type", "application/json")
headers.Set("Authorization", "Basic token")

// The client builder takes your desired settings and builds a HTTP client.
// These settings are baked into the client and are sent with each request.
c := goclient.NewBuild().
    SetBaseURL("https://foobar.com").
    SetConnectionTimeout(5 * time.Second).
    SetResponseTimeout(5 * time.Second).
    SetRequestHeaders(headers).
    SetUserAgent("go-http").
    Build()
```

###### Performing a request
The HTTP client handles low-level plumbing operations so that you only focus on the response. For example:  
* The response body is automatically closed for each request.
* Use the built-in method to unmarshal JSON responses without concern for the code constructs.

An example of a GET request:

```go
type Student struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Grade int    `json:"grade"`
}

var c = NewClient()

func NewClient() goclient.Client {
	// Standard request headers can be parsed using defined package constants.
	headers := make(http.Header)
	headers.Set(HeaderContentType, ContentTypeJson)
	headers.Set(HeaderAuthorization, "Basic token")

	return gohttp.NewBuild().
        SetBaseURL("https://foobar.com").
    	SetConnectionTimeout(5 * time.Second).
    	SetResponseTimeout(5 * time.Second).
    	SetRequestHeaders(headers).
    	SetUserAgent("go-http").
    	Build()
}

func GetStudent(id int) (*Student, error) {
	// Instead of nil, specific headers for this API can be parsed here.
	response, err := c.Get(fmt.Sprintf("/_api/student?id=%d", id), nil)
	if err != nil {
		return nil, err
	}
	var jsonData Student
	if err := response.UnmarshalJson(&jsonData); err != nil {
		return nil, err
	}
	return &Student, nil
}

func main() {
	student, err := GetStudent(1)
	if err != nil {
		return nil, err
	}
	fmt.Prinf("%v", student.ID)
	fmt.Prinf("%v", student.Name)
	fmt.Prinf("%v", student.Age)
	fmt.Prinf("%v", student.Grade)
}
```
