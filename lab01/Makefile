.PHONY: build_all
build_all:
	make build_client
	make build_server

build_client:
	go build -o client/go_tcp_client client/go_tcp_client.go

build_server:
	go build -o server/go_tcp_server server/go_tcp_server.go


serve:
	./server/go_tcp_server

run_client:
	./client/go_tcp_client
