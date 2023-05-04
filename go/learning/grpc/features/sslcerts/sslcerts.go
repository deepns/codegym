package sslcerts

import (
	"path/filepath"
	"runtime"
)

var mydir string

func init() {
	// runtime.Caller Caller reports file and line number information about
	// function invocations on the calling goroutine's stack.  We need the
	// current file path to find the sslcerts directory relative to where it is
	// imported.
	_, currentFilePath, _, _ := runtime.Caller(0)
	mydir = filepath.Dir(currentFilePath)
}

// Path returns absolute path of the given relative path
// Used to find tls certs from the sslcerts directory
func Path(path string) string {
	if filepath.IsAbs(path) {
		return path
	}

	return filepath.Join(mydir, path)
}
