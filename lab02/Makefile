.PHONY: all
all: clean prepare main run


.PHONY: clean
clean:
	rm ./service

.PHONY: prepare
prepare:
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux go mod download

main:
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -mod readonly ./cmd/service/service.go

.PHONY: run
run:
	./service