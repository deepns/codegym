// Setting up a https server
// Turns out to be super simple to start, if we can use the default
// listener. Register the handler with http.HandleFunc and start the
// server with http.ListenAndServeTLS(). Provide the server cert and
// key in the arguments to http.ListenAndServeTLS().
// Setting up our own listener and managing the incoming connections is
// for another day
package main

import (
	"fmt"
	"io"
	"net/http"
)

func EchoHeader(rw http.ResponseWriter, r *http.Request) {
	header := r.Header
	for k := range header {
		io.WriteString(rw, fmt.Sprintf("%v: %v\n", k, header[k]))
	}

	// panic("want to see stack trace")
	//
	// panicked the handler to see the call stack leading up to the handler
	//
	// 2022/03/03 11:34:23 http2: panic serving 127.0.0.1:36466: want to see stack trace
	// goroutine 35 [running]:
	// net/http.(*http2serverConn).runHandler.func1()
	//         /usr/local/go/src/net/http/h2_bundle.go:5825 +0x125
	// panic({0x61fe40, 0x6b72c0})
	//         /usr/local/go/src/runtime/panic.go:1038 +0x215
	// main.EchoHeader({0x6bde20, 0xc00025c000}, 0x1)
	//         /home/deepan_seeralan/github/codegarage/go/learning/https/https.go:21 +0x170
	// net/http.HandlerFunc.ServeHTTP(0x84ed80, {0x6bde20, 0xc00025c000}, 0x0)
	//         /usr/local/go/src/net/http/server.go:2046 +0x2f
	// net/http.(*ServeMux).ServeHTTP(0x0, {0x6bde20, 0xc00025c000}, 0xc000258100)
	//         /usr/local/go/src/net/http/server.go:2424 +0x149
	// net/http.serverHandler.ServeHTTP({0x0}, {0x6bde20, 0xc00025c000}, 0xc000258100)
	//         /usr/local/go/src/net/http/server.go:2878 +0x43b
	// net/http.initALPNRequest.ServeHTTP({{0x6bef00, 0xc000090000}, 0xc000109180, {0xc000156000}}, {0x6bde20, 0xc00025c000}, 0xc000258100)
	//         /usr/local/go/src/net/http/server.go:3479 +0x245
	// net/http.(*http2serverConn).runHandler(0x0, 0x0, 0x0, 0x0)
	//         /usr/local/go/src/net/http/h2_bundle.go:5832 +0x78
	// created by net/http.(*http2serverConn).processHeaders
	//         /usr/local/go/src/net/http/h2_bundle.go:5562 +0x510
}

func main() {
	// Register the endpoint handlers
	// The handlers called in separate go routines
	http.HandleFunc("/", EchoHeader)
	http.HandleFunc("/hello", func(rw http.ResponseWriter, r *http.Request) {
		io.WriteString(rw, "Hello, world!")
	})

	// Start the server with the TLS certificate and keys
	http.ListenAndServeTLS(":9000", "certs/server.crt", "certs/server.key", nil)
}
