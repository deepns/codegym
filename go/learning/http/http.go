package http

// Exploring net/http package in golang
// https://pkg.go.dev/net/http

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// GetAllAndDecodeJson makes a GET call to golang github URL,
// reads all data and then decodes the data using json decoder
func GetAllAndDecodeJson() {
	fmt.Println(">>> Exploring net/http. Get all data and decode them using json <<<")
	// The go repo in github
	const golangRepoURL = "https://api.github.com/repos/golang/go"

	// http.Get returns a response object and error.
	// note that completion of Get doesn't mean that all data is received.
	// response body is streamed on demand.
	resp, err := http.Get(golangRepoURL)
	if err != nil {
		panic(err)
	}

	// notes from docs of http.Response

	// The response body is streamed on demand as the Body field
	// is read. If the network connection fails or the server
	// terminates the response, Body.Read calls return an error.
	//
	// The http Client and Transport guarantee that Body is always
	// non-nil, even on responses without a body or responses with
	// a zero-length body. It is the caller's responsibility to
	// close Body. The default HTTP client's Transport may not
	// reuse HTTP/1.x "keep-alive" TCP connections if the Body is
	// not read to completion and closed.

	defer resp.Body.Close()

	// resp.Body is of type io.ReadCloser which implements the Reader interface
	// I remember using ioutil.ReadAll() to read from files using the File type.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// body is still a byte array.
	// converting to strings
	jsonData := string(body)
	fmt.Printf("data: %v\n", jsonData)

	// lets get some interesting info from the json data
	decoder := json.NewDecoder(strings.NewReader(jsonData))
	type Repo struct {
		Language                    string
		Created_at                  string
		Size, Forks                 int
		Open_issues, Watchers_count int
	}

	var golang Repo
	if err := decoder.Decode(&golang); err != nil {
		panic(err)
	}
	fmt.Printf("golang: %+v\n", golang)
}

// GetAndDecodeJson makes a GET request and runs a json
// decoder on the response body.
func GetAndDecodeJson() {
	fmt.Println(">>> Exploring net/http. Decode json in a http response body <<<")

	// The go repo in github
	const golangRepoURL = "https://api.github.com/repos/golang/go"

	// http.Get returns a response object and error.
	// note that completion of Get doesn't mean that all data is received.
	// response body is streamed on demand.
	resp, err := http.Get(golangRepoURL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// lets get some interesting info from the data contained
	// in the response body
	decoder := json.NewDecoder(resp.Body)
	type Repo struct {
		Language                    string
		Created_at                  string
		Size, Forks                 int
		Open_issues, Watchers_count int
	}

	var golang Repo
	if err := decoder.Decode(&golang); err != nil {
		panic(err)
	}
	fmt.Printf("golang: %+v\n", golang)
}
