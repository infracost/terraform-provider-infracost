BINARY := terraform-provider-infracost
VERSION := $(shell cat version)
PLUGIN_ROOT_PATH := ~/.terraform.d/plugins

ifndef $(GOOS)
	GOOS=$(shell go env GOOS)
endif

ifndef $(GOARCH)
	GOARCH=$(shell go env GOARCH)
endif

INSTALL_PATH := $(PLUGIN_ROOT_PATH)/infracost.io/infracost/infracost/$(VERSION)/$(GOOS)_$(GOARCH)

default: testacc

.PHONY: testacc build windows linux darwin install

# Run acceptance tests
testacc:
	TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 120m

build: $(GOOS)

windows:
	env GOOS=windows GOARCH=amd64 go build -i -v -o build/$(BINARY)-windows-amd64

linux:
	env GOOS=linux GOARCH=amd64 go build -i -v -o build/$(BINARY)-linux-amd64

darwin:
	env GOOS=darwin GOARCH=amd64 go build -i -v -o build/$(BINARY)-darwin-amd64

clean:
	go clean
	rm -rf build/$(BINARY)*

install: build
	mkdir -p $(INSTALL_PATH)
	cp build/$(BINARY)-$(GOOS)-$(GOARCH) $(INSTALL_PATH)/$(BINARY)
	# Add a link from the old-style plugin paths to maintain compatibility with older versions of Terraform
	mkdir -p $(PLUGIN_ROOT_PATH)/$(GOOS)_$(GOARCH)
	ln -sf $(INSTALL_PATH)/$(BINARY) $(PLUGIN_ROOT_PATH)/$(GOOS)_$(GOARCH)/$(BINARY)

fmt:
	go fmt ./...

lint:
	golangci-lint run
