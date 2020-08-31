BINARY := terraform-provider-infracost
VERSION := $(shell cat version)

PLUGIN_ROOT_PATH := ~/.terraform.d/plugins
PLUGIN_INSTALL_PATH := $(PLUGIN_ROOT_PATH)/infracost.io/infracost/infracost/$(VERSION)

.PHONY: deps run build windows linux darwin build_all clean test fmt lint

deps:
	go mod download

run:
	go run $(ENTRYPOINT) $(ARGS)

build:
	go build -i -v -o build/$(BINARY)_$(VERSION)

windows:
	env GOOS=windows GOARCH=amd64 go build -i -v -o build/$(BINARY)_v$(VERSION)-windows-amd64

linux:
	env GOOS=linux GOARCH=amd64 go build -i -v -o build/$(BINARY)_v$(VERSION)-linux-amd64

darwin:
	env GOOS=darwin GOARCH=amd64 go build -i -v -o build/$(BINARY)_v$(VERSION)-darwin-amd64

build_all: build windows linux darwin

install_linux: linux
	mkdir -p $(PLUGIN_INSTALL_PATH)/linux_amd64
	cp build/$(BINARY)_v$(VERSION)-linux-amd64 $(PLUGIN_INSTALL_PATH)/linux_amd64/$(BINARY)_v$(VERSION)
	mkdir -p $(PLUGIN_ROOT_PATH)/linux_amd64
	ln -s $(PLUGIN_INSTALL_PATH)/linux_amd64/$(BINARY)_v$(VERSION) $(PLUGIN_ROOT_PATH)/linux_amd64/$(BINARY)_v$(VERSION)

install_darwin: darwin
	mkdir -p $(PLUGIN_INSTALL_PATH)/darwin_amd64
	cp build/$(BINARY)_v$(VERSION)-darwin-amd64 $(PLUGIN_INSTALL_PATH)/darwin_amd64/$(BINARY)_v$(VERSION)
	mkdir -p $(PLUGIN_ROOT_PATH)/darwin_amd64
	ln -sF $(PLUGIN_INSTALL_PATH)/darwin_amd64/$(BINARY)_v$(VERSION) $(PLUGIN_ROOT_PATH)/darwin_amd64/$(BINARY)_v$(VERSION)

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
