proto:
	protoc -I ./reminders \
		--go_out ./reminders --go_opt=paths=source_relative \
		--go-grpc_out ./reminders --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out ./reminders --grpc-gateway_opt paths=source_relative \
		reminders/reminders/reminders.proto

server:
	go run ./reminders/server/main.go

client:
	go run ./reminders/client/main.go

# What does PHONY do?
# https://www.gnu.org/software/make/manual/html_node/Phony-Targets.html
.PHONY: proto server client