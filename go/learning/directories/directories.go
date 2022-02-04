// Exploring directory access and management
package main

import (
	"fmt"
	"os"
	"path"
)

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// Get current working directory
	curDir, _ := os.Getwd()

	// Get user's home directory
	homeDir, err := os.UserHomeDir()
	handleErr(err)
	fmt.Printf("homeDir: %v\n", homeDir)

	// listing files under current directory
	files, err := os.ReadDir(curDir)
	handleErr(err)

	for _, f := range files {
		fmt.Printf("f: %v\n", f)
	}

	// ReadDir returns a sorted list of files/dir under the given dir
	filesOneLevelAbove, err := os.ReadDir(path.Dir(curDir))
	handleErr(err)
	for _, f := range filesOneLevelAbove {
		stat, _ := os.Stat(path.Join(path.Dir(curDir), f.Name()))
		fmt.Printf("f: %v, dir?:%v, size:%v\n", f.Name(), f.IsDir(), stat.Size())
	}

	// get all environment variables
	// env vars are returned in key=value forms
	fmt.Println("Environment variables:")
	for _, env := range os.Environ() {
		fmt.Printf("env: %v\n", env)
	}
}
