// Learning the usage and management of environment variables
package main

import (
	"fmt"
	"os"
)

func main() {

	// Get all environment variables
	// os.Environ() returns all environment variables in a list of key=value pairs
	envvars := os.Environ()
	for _, ev := range envvars {
		fmt.Println(ev)
	}

	// Get a specific environment variable
	fmt.Printf("os.Getenv(\"SHELL\"): %v\n", os.Getenv("SHELL"))

	// Lookup environment variable
	// LookupEnv differentiates between a variable that is set to empty
	// value versus value that is not set.
	term, ok := os.LookupEnv("TERM")
	if !ok {
		fmt.Println("TERM variable is not set")
	} else {
		fmt.Printf("TERM set to %v\n", term)
	}

	// Set environment variable
	os.Setenv("DBUSER", "admin")
	os.Setenv("DBPASS", "")

	// With Getenv(), if the variable is not set, it returns a empty string
	// so we can't differentiate between a variable which is set to an empty
	// value versus variables which are not set.
	password, ok := os.LookupEnv("DBPASS")
	if !ok {
		fmt.Println("DBPASS is not set")
	} else {
		if len(password) == 0 {
			fmt.Println("DBPASS is empty")
		}
	}

	// Unset environment variable
	os.Unsetenv("DBPASS")

	os.Setenv("DBPASS", "jdi03jfj7v")

	// Expand environment variable
	// variables can be expanded in the similar way shell
	// variables are expanded.
	fmt.Println(os.ExpandEnv("Using the credentials $DBUSER:$DBPASS"))

	// os.ExpandEnv internally uses Getenv() as the mapping function
	// Getenv signature is also func(string) strings
	os.Expand("Credentials $USER $PASSWORD", func(s string) string {
		switch s {
		case "USER":
			return "root"
		case "PASSWORD":
			return "root123"
		default:
			return "-"
		}
	})

	// Clear all environment variables
	os.Clearenv()
	fmt.Printf("os.Environ(): %v\n", os.Environ())
}
