.PHONY: build install clean run help

BINARY_NAME := ghostshell
GO          := go
GOFLAGS     := -v
VERSION     ?= latest
BUILD_DIR   := ./bin

help:
	@echo "ghostshell - Build targets:"
	@echo "  make build    - Build the binary"
	@echo "  make install  - Install to GOBIN or GOPATH/bin"
	@echo "  make run      - Build and run locally"
	@echo "  make clean    - Remove build artifacts"

build:
	@echo ":: Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	$(GO) build $(GOFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) .

install: build
	@echo ":: Installing $(BINARY_NAME)..."
	@mkdir -p $$(shell $(GO) env GOBIN 2>/dev/null || echo $$(shell $(GO) env GOPATH)/bin)
	cp $(BUILD_DIR)/$(BINARY_NAME) $$(shell $(GO) env GOBIN 2>/dev/null || echo $$(shell $(GO) env GOPATH)/bin)/

run: build
	@echo ":: Running $(BINARY_NAME)..."
	@$(BUILD_DIR)/$(BINARY_NAME)

clean:
	@echo ":: Cleaning build artifacts..."
	rm -rf $(BUILD_DIR)
	$(GO) clean
