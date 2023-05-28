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
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("HelloWorldHandler", HelloWorldHandler)
}

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
