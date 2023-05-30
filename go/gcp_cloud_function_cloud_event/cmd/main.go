package main

import (
	"log"
	"os"

	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"

	// empty import to trigger init() in function.go
	_ "github.com/deepns/codegym/go/gcp_functions_cloud_event"
)

// Check out the following links for testing
// https://cloud.google.com/functions/docs/running/calling#cloudevent-function-curl-tabs-pubsub
// Testing with pubsub emulator
// https://cloud.google.com/functions/docs/local-development

func main() {
	// Use PORT environment variable, or default to 8080.
	port := "8080"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}
	if err := funcframework.Start(port); err != nil {
		log.Fatalf("funcframework.Start: %v\n", err)
	}
}
