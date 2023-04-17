protoc -I ./reminders \
    --go_out=./reminders --go_opt=paths=source_relative \
    --go-grpc_out=./reminders --go-grpc_opt=paths=source_relative \
    reminders/reminders/reminders.proto
