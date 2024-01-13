
# Variables
APP_NAME := Boss Payback
MAIN_FILE := main.go
BINARY_NAME := boss-payback
BINARY_PATH := ./bin/$(BINARY_NAME)
BUILD_DIR := ./bin
SRC_DIR := ./cmd

# Build the application
build:
	@echo "Building $(APP_NAME)..."
	@go build -o $(BINARY_PATH) $(SRC_DIR)/$(MAIN_FILE)

# Build first then Run the application
run:
	@echo "Building $(APP_NAME)..."
	@go build -o $(BINARY_PATH) $(SRC_DIR)/$(MAIN_FILE)
	@echo "Running $(APP_NAME)..."
	@$(BINARY_PATH)

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf $(BUILD_DIR)

# Install dependencies
deps:
	@echo "Installing dependencies..."
	@go mod tidy

# Format the code
fmt:
	@echo "Formatting code..."
	@go fmt ./...
	@echo "Formatting completed."

# Help target to display available make commands
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  build       Build the application"
	@echo "  run         Run the application"
	@echo "  clean       Clean build artifacts"
	@echo "  deps        Install dependencies"
	@echo "  fmt         Format the code"
	@echo "  help        Display this help message"

# Default target
.DEFAULT_GOAL := help
