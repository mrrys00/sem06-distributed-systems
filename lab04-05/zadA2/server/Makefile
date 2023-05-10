server_path = "server/"
grpcproject_path = "grpcproject/"
client_path = "client/"

.PHONY: prepare
prepare:
	go get google.golang.org/protobuf/cmd/protoc-gen-go
	go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go install google.golang.org/protobuf/cmd/protoc-gen-go
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
	export PATH="$$PATH:$(go env GOPATH)/bin"

proto:
	mkdir -p $(grpcproject_path)
	protoc -I ../$(grpcproject_path) --go_out=$(grpcproject_path). --go_opt=paths=source_relative   --go-grpc_out=$(grpcproject_path). --go-grpc_opt=paths=source_relative   ../$(grpcproject_path)grpcproject.proto

run:
	go run .
