BINARY := terraform-provider-infracost
VERSION := $(shell cat version)
PLUGIN_ROOT_PATH := ~/.terraform.d/plugins
PLUGIN_INSTALL_PATH := $(PLUGIN_ROOT_PATH)/infracost.io/infracost/infracost/$(VERSION)

ifndef $(GOOS)
	GOOS=$(shell go env GOOS)
endif

ifndef $(GOARCH)
	GOARCH=$(shell go env GOARCH)
endif

default: testacc

.PHONY: testacc build windows linux darwin install_linux install_darwin

# Run acceptance tests
testacc:
	TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 120m

build:
	go build -i -v -o build/$(BINARY)

windows:
	env GOOS=windows GOARCH=amd64 go build -i -v -o build/$(BINARY)-windows-amd64

linux:
	env GOOS=linux GOARCH=amd64 go build -i -v -o build/$(BINARY)-linux-amd64

darwin:
	env GOOS=darwin GOARCH=amd64 go build -i -v -o build/$(BINARY)-darwin-amd64

install:
	mkdir -p $(PLUGIN_INSTALL_PATH)/$(GOOS)_$(GOARCH)
	cp build/$(BINARY)-$(GOOS)-$(GOARCH) $(PLUGIN_INSTALL_PATH)/$(GOOS)_$(GOARCH)/$(BINARY)
	mkdir -p $(PLUGIN_ROOT_PATH)/$(GOOS)_$(GOARCH)
	# Add a link from the old-style plugin paths to maintain compatibility with older versions of Terraform
	ln -sf $(PLUGIN_INSTALL_PATH)/$(GOOS)_$(GOARCH)/$(BINARY) $(PLUGIN_ROOT_PATH)/$(GOOS)_$(GOARCH)/$(BINARY)
