#!/bin/bash

###########################################################
# Script: test_function.sh
# Description: Send a test PubSub event to a function for local testing using curl.
# Version: 1.0
###########################################################

# Function to generate a random string of alphabets and encode them in base64
generate_random_string() {
    local length=$1
    # Limit the length to a maximum of 32 characters
    if [[ $length -gt 32 ]]; then
        length=32
        echo "Warning: Length exceeds maximum of 32 characters. Truncating to 32 characters."
    fi
    local random_string=$(openssl rand -base64 $length | tr -dc 'a-zA-Z' | head -c $length)
    echo "$random_string"
}

# Send a test PubSub event to the function for local testing
PORT=${PORT:-8080}

# Generate a random string of alphabets
LENGTH=${LENGTH:-32}
MESSAGE_DATA=$(generate_random_string $LENGTH)
MESSAGE_DATA_BASE64=$(echo -n ${MESSAGE_DATA} | base64)
echo "Message data: ${MESSAGE_DATA}"
echo "Message data in base64: ${MESSAGE_DATA_BASE64}"

# Set the URL and base headers
URL="http://localhost:${PORT}"
HEADERS=(
    "-H 'Content-Type: application/json'"
    "-H 'ce-id: 123451234512345'"
    "-H 'ce-specversion: 1.0'"
    "-H 'ce-time: 2020-01-02T12:34:56.789Z'"
    "-H 'ce-type: google.cloud.pubsub.topic.v1.messagePublished'"
    "-H 'ce-source: //pubsub.googleapis.com/projects/MY-PROJECT/topics/MY-TOPIC'"
)

# Set the JSON payload
JSON_DATA=$(cat <<EOF
{
    "message": {
        "data": "${MESSAGE_DATA_BASE64}",
        "attributes": {
            "attr1": "attr1-value"
        }
    },
    "subscription": "projects/MY-PROJECT/subscriptions/MY-SUB"
}
EOF
)

# Build the curl command
CURL_CMD="curl ${URL}"
for header in "${HEADERS[@]}"; do
    CURL_CMD+=" ${header}"
done
CURL_CMD+=" -d '${JSON_DATA}'"

# Execute the curl command
eval "${CURL_CMD}"
