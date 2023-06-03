ENTRY_POINT = ServeHTTP
RUNTIME = go120
REGION = us-central1
FUNCTION_NAME = go-http-function

local:
	FUNCTION_TARGET=$(ENTRY_POINT) go run cmd/main.go
deploy:
	gcloud functions deploy $(FUNCTION_NAME) \
		--gen2 \
		--runtime=$(RUNTIME) \
		--region=$(REGION) \
		--source=. \
		--entry-point=$(ENTRY_POINT) \
		--trigger-http \
		--allow-unauthenticated