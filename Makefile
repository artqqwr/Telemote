# Set the project name
PROJECT_NAME = $(shell basename "$(CURDIR)")

# Set the Go compiler
GO=go

# Set the Go flags
GO_FLAGS=-v

# Set the directories
SRC_DIR=./cmd/app
BIN_DIR=bin

# Default target
all: clean build

# Build the binary
build:
	@echo "Building binary..."
	$(GO) build $(GO_FLAGS) -o $(BIN_DIR)/$(PROJECT_NAME) $(SRC_DIR)/*.go

# Clean up
clean:
	@echo "Cleaning up..."
	go clean
	rm -rf $(BIN_DIR)/*

# Run the application
run:
	@echo "Running application..."
	air

# Run the tests
test:
	@echo "Running tests..."
	$(GO) test $(GO_FLAGS) ./...

# Install the dependencies
deps:
	@echo "Installing dependencies..."
	$(GO) get -v ./...

# Update the dependencies
update:
	@echo "Updating dependencies..."
	$(GO) get -u -v ./...