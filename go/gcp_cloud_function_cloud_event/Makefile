ENTRY_POINT = helloPubSubHandler
RUNTIME = go120
REGION = us-central1
FUNCTION_NAME = go-cloud-event-function
PUBSUB_TOPIC = helloworld

local:
	FUNCTION_TARGET=$(ENTRY_POINT) go run cmd/main.go
deploy:
	gcloud functions deploy $(FUNCTION_NAME) \
		--gen2 \
		--runtime=$(RUNTIME) \
		--region=$(REGION) \
		--source=. \
		--entry-point=$(ENTRY_POINT) \
		--trigger-topic=$(PUBSUB_TOPIC) \
		--allow-unauthenticated