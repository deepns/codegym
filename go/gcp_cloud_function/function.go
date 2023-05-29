// Directory structure:
//
// ├── myfunction.go
// ├── go.mod
// └── cmd/
//     └── main.go
//
// https://cloud.google.com/functions/docs/writing#directory-structure-go
// https://cloud.google.com/functions/docs/writing/http#functions_http_go

package function

import (
	"fmt"
	"log"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("ServeHTTP", ServeHTTP)
}

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("ServeHTTP: %v", r)
	switch r.URL.Path {
	case "/":
		HelloWorldHandler(w, r)
	case "/foo":
		FooHandler(w, r)
	default:
		http.NotFound(w, r)
	}
}

func FooHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("FooHandler: %v %v %v", r.Method, r.URL.Path, r.URL.RawQuery)
	log.Printf("FooHandler: queryParams:%v", r.URL.Query())
	fmt.Fprintf(w, "Hello Foo!")
}

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
