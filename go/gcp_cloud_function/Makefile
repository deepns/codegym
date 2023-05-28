deploy:
	gcloud functions deploy go-http-function \
		--gen2 \
		--runtime=go120 \
		--region=us-central1 \
		--source=. \
		--entry-point=HelloWorldHandler \
		--trigger-http \
		--allow-unauthenticated