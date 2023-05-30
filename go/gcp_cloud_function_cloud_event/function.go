package gcpcloudeventfunction

import (
	"context"
	"fmt"
	"log"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/cloudevents/sdk-go/v2/event"
)

// MessagePublishedData contains the full Pub/Sub message
// https://cloud.google.com/eventarc/docs/cloudevents#pubsub_1
// https://github.com/googleapis/google-cloudevents/blob/main/proto/google/events/cloud/pubsub/v1/data.proto
// talks about data format
type MessagePublishedData struct {
	Message PubSubMessage
}

// PubSubMessage is the payload of a Pub/Sub event.
// https://cloud.google.com/pubsub/docs/reference/rest/v1/PubsubMessage talks about
// Pub/Sub message format
//
//	{
//	  "data": string,
//	  "attributes": {
//	    string: string,
//	    ...
//	  },
//	  "messageId": string,
//	  "publishTime": string,
//	  "orderingKey": string
//	}
//
// Message must contain a non-empty data field or at least one attribute.
type PubSubMessage struct {
	Data []byte `json:"data"`
}

func init() {
	// Register the CloudEvent function with the Functions Framework
	functions.CloudEvent("helloPubSubHandler", helloPubSubHandler)
}

// Function helloPubSubHandler accepts and handles a CloudEvent object
// Event data is passed in the form of a CloudEvent object
// https://cloud.google.com/eventarc/docs/cloudevents
func helloPubSubHandler(ctx context.Context, e event.Event) error {
	// My code here
	// Access the CloudEvent data payload via e.Data() or e.DataAs(...)
	// https://pkg.go.dev/github.com/cloudevents/sdk-go/v2/event#Event
	var msg MessagePublishedData
	if err := e.DataAs(&msg); err != nil {
		return fmt.Errorf("error while decoding data: %s", err.Error())
	}

	// Access the CloudEvent attributes via e.Context.GetExtensions()
	// https://pkg.go.dev/github.com/cloudevents/sdk-go/v2/event#Event
	// https://pkg.go.dev/github.com/cloudevents/sdk-go/v2/types#Attributes
	// https://pkg.go.dev/github.com/cloudevents/sdk-go/v2/types#AttributesAs

	name := string(msg.Message.Data)
	if name == "" {
		name = "Foo"
	}
	log.Printf("Hello %s!", name)

	return nil
}
