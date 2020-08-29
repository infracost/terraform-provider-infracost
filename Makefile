BINARY := terraform-provider-infracost
VERSION := $(shell cat version)

.PHONY: deps run build windows linux darwin build_all clean test fmt lint

deps:
	go mod download

run:
	go run $(ENTRYPOINT) $(ARGS)

build:
	go build -i -v -o build/$(BINARY)_$(VERSION)

windows:
	env GOOS=windows GOARCH=amd64 go build -i -v -o build/$(BINARY)_$(VERSION)-windows-amd64

linux:
	env GOOS=linux GOARCH=amd64 go build -i -v -o build/$(BINARY)_$(VERSION)-linux-amd64

darwin:
	env GOOS=darwin GOARCH=amd64 go build -i -v -o build/$(BINARY)_$(VERSION)-darwin-amd64

build_all: build windows linux darwin

install: build
	cp build/$(BINARY)_$(VERSION) ~/.terraform.d/plugins/

clean:
	go clean
	rm -rf build/$(BINARY)*
	rm -rf release/$(BINARY)*

test:
	TF_ACC=true  go test ./... $(or $(ARGS), -v)

fmt:
	go fmt ./...

lint:
	golangci-lint run
