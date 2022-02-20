// Exploring cmdline flags
package main

import (
	"flag"
	"fmt"
	"net/url"
	"time"
)

type urlValue struct {
	url *url.URL
}

func (u urlValue) String() string {
	if u.url != nil {
		return fmt.Sprintf("%v", u.url)
	}
	return ""
}

func (u urlValue) Set(urlString string) error {
	if v, err := url.Parse(urlString); err != nil {
		return err
	} else {
		*u.url = *v
	}
	return nil
}

func main() {

	// flags are typed
	// Int, Bool, String, Duration, Float64 are some of the default ones available in flag package.
	// The general flow goes like this
	//	Define the flags
	// 	call Parse
	// 	access the flags

	// from the cmdline,
	//	flags can be specified with -flag or --flag
	//	-h/--help shows the help information based on the flags given
	//	flags are listed in lexicographical order

	// flag.Int() returns a pointer to a variable holding the value of the flag
	// flag.IntVar() takes the address of the value to store the flag
	// Defining some basic flags
	var host = flag.String("hostname", "localhost", "hostname of the server")
	var port = flag.Int("port", 8080, "port to connect to")
	var verbose = flag.Bool("v", false, "turn on verbose logging")

	var nRetry int
	var nThreads int
	maxTimeout, _ := time.ParseDuration("10s")
	var timeout = flag.Duration("timeout", maxTimeout, "max timeout")

	// note: if the default value is 0, flag doesn't seem to consider it
	// as explicit default.

	// -h/--help in this case will show
	//   -retry int
	//			number of retry attempts
	flag.IntVar(&nRetry, "retry", 0, "number of retry attempts")

	// -h/--help i this case will show
	// -num-threads int
	// 				number of threads to use (default 1)
	flag.IntVar(&nThreads, "num-threads", 1, "number of threads to use")

	// Adding a custom value
	var u = &url.URL{}
	flag.Var(&urlValue{u}, "url", "link to the server")

	flag.Parse()

	// Number of flags specified in the command line
	fmt.Printf("flag.NFlag(): %v\n", flag.NFlag())

	fmt.Printf("host: %v\n", *host)
	fmt.Printf("port: %v\n", *port)
	fmt.Printf("verbose: %v\n", *verbose)
	fmt.Printf("nThreads: %v\n", nThreads)
	fmt.Printf("timeout: %v\n", timeout)
	fmt.Printf("host:%v, scheme %v\n", u.Host, u.Scheme)

	// Number of args remaining after the flags are parsed
	fmt.Printf("flag.NArg(): %v\n", flag.NArg())

	// Get all the args as a string slice
	args := flag.Args()
	for _, arg := range args {
		fmt.Printf("arg: %v\n", arg)
	}

	// Get specific args using flag.Arg()
	if flag.NArg() > 0 {
		fmt.Printf("flag.Arg(1): %v\n", flag.Arg(1))
	}

	// Prints the default message listing out the flags and their description
	flag.PrintDefaults()

	// we can't seem to have alias for the flag name
	// for e.g. -v / --verbose
	// we can't one flag that can take either -v or --verbose
	// have to check it out whether such option is available or not
}
