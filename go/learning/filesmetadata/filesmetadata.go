// Exploring go functionalities to read file metadata
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	// Get file stats with os.Stat
	// os.Stat returns a fs.FileInfo type
	_, err := os.Stat("missingfile")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	fileinfo, err := os.Stat("frenchwords.md")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("fileinfo.Name(): %v\n", fileinfo.Name())
	fmt.Printf("fileinfo.Size(): %v\n", fileinfo.Size())
	fmt.Printf("fileinfo.ModTime(): %v\n", fileinfo.ModTime())
	mode := fileinfo.Mode()
	fmt.Printf("mode: %v\n", mode)
	fmt.Printf("mode.IsRegular(): %v\n", mode.IsRegular())
	fmt.Printf("mode.Perm(): %v\n", mode.Perm())

	fileinfoLink, err := os.Stat("fwords")
	if err != nil {
		log.Fatal(err)
	}

	// Check whether a file and its soft link are treated the same
	// Turns out yes.
	fmt.Printf("os.SameFile(fileinfo, fileinfoLink): %v\n",
		os.SameFile(fileinfo, fileinfoLink))
}
