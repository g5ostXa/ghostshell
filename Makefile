.PHONY: build install clean run help

BINARY_NAME=ghostshell
GO=go
GOFLAGS=-v
VERSION?=dev
BUILD_DIR=./bin

help:
	@echo "ghostshell - Build targets:"
	@echo "  make build    - Build the binary"
	@echo "  make install  - Install to $$GOBIN or $$HOME/go/bin"
	@echo "  make run      - Build and run locally"
	@echo "  make clean    - Remove build artifacts"

build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	$(GO) build $(GOFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) .

install: build
	@echo "Installing $(BINARY_NAME)..."
	cp $(BUILD_DIR)/$(BINARY_NAME) $$($(GO) env GOBIN)

run: build
	./$(BUILD_DIR)/$(BINARY_NAME)

clean:
	@echo "Cleaning build artifacts..."
	rm -rf $(BUILD_DIR)
	$(GO) clean