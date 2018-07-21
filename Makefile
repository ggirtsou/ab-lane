# todo: git hash, generate static binary, lint target

.PHONY: install
install:
	go get -u -d -t -v ./...

.PHONY: clean
build: clean
	go build -o broker ./cmd/broker

.PHONY: test
test:
	go test -v -cover ./...

.PHONY: clean
clean:
	rm -f ./broker

.PHONY: protos
protos:
	protoc --proto_path=proto --go_out=generated/pb proto/*.proto

.PHONY: all
all: install protos test build
