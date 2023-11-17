SOURCE := cmd/main.go
BINARY_NAME := go-pdirs
BUILD_DIR := bin

GOCMD := go
GOBUILD := $(GOCMD) build
GOCLEAN := $(GOCMD) clean
GOTEST := $(GOCMD) test

.PHONY: clean test build
all: clean test build

clean:
	$(GOCLEAN) -testcache
	rm -rf "$(BUILD_DIR)/$(BINARY_NAME)"

test: clean
	$(GOTEST) -v ./...

build: clean
	mkdir -p "$(BUILD_DIR)"
	$(GOBUILD) -o "$(BUILD_DIR)/$(BINARY_NAME)" $(SOURCE)

