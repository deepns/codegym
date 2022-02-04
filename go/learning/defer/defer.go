// Learning defer
package main

import (
	"fmt"
	"io"
	"net/http"
)

func fA() {
	fmt.Println("fA()")
}

func fB() {
	fmt.Println("fB()")
}

func fC() {
	fmt.Println("fC()")
	defer fD()
}

func fD() {
	fmt.Println("fD()")
}

func main() {
	// functions deferred are executed in LIFO order
	defer fA()
	defer fB()
	defer fC()

	// it is idiomatic in go to defer the close of a
	// resource like file, socket etc.
	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)
	fmt.Printf("data: %s\n", data)
}
