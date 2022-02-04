// Learning about temp files and temp directories
package main

import (
	"fmt"
	"os"
)

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// Get the path of temp dir used by the runtime
	fmt.Printf("os.TempDir(): %v\n", os.TempDir())

	// MkdirTemp creates a temporary directory under the source
	// dir (if empty, it uses os.TempDir() instead).
	// * in the pattern will be replaced by a random string
	// It is the caller's responsibility to remove the temporary directory
	// then what's the point of creating temp?
	newTempDir, err := os.MkdirTemp("", "tempdir_*" /* pattern */)
	handleErr(err)
	fmt.Printf("newTempDir: %v\n", newTempDir)

	// get directory of the executable
	curDir, _ := os.Executable()
	fmt.Printf("curDir: %v\n", curDir)

	// when running the file with 'go run .', the executable is created
	// in a temp directory under the go temp directory indicated by
	// os.TempDir()

	// CreateTemp() creates a new file under the given directory
	// and returns a os.File type.
	newTempFile, err := os.CreateTemp(newTempDir, "tfile_*")
	handleErr(err)
	defer newTempFile.Close()
	fmt.Printf("newTempFile: %v\n", newTempFile)

	// write some data to the temp file
	newTempFile.WriteString("Hello, world!\n")

	// can temp file be created under normal directory?. Turns out Yes.
	// Getwd() returns the current working directory
	srcDir, _ := os.Getwd()
	anotherTempFile, err := os.CreateTemp(srcDir, "tempfile_*")
	handleErr(err)
	// Remove file at the end of the run
	// Note: defer calls are run in the LIFO order.
	// so close the file handle before removing the file
	defer os.Remove(anotherTempFile.Name())
	defer anotherTempFile.Close()

	anotherTempFile.Write([]byte("Just checking\n"))
}
