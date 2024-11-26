BIN_DIR = _bin
PROTO_DIR = proto
SERVER_DIR = beef/server
CLIENT_DIR = client
SERVER_BIN = ${BIN_DIR}/server
CLIENT_BIN = ${BIN_DIR}/client

# Define platform-specific commands
ifeq ($(OS), Windows_NT)
	SHELL := powershell.exe
	.SHELLFLAGS := -NoProfile -Command
	SERVER_BIN = ${BIN_DIR}/beef/server.exe
	CLIENT_BIN = ${BIN_DIR}/beef/client.exe
else
	SHELL := bash
	SERVER_BIN = ${BIN_DIR}/beef/server
	CLIENT_BIN = ${BIN_DIR}/beef/client
endif

.DEFAULT_GOAL := help
.PHONY: all clean build run help

# Target to build the server and client for Beef
all: build ## Build the Beef server and client

build: ## Generate Pbs and build the Beef service
	@echo "Building Beef server and client..."
	@set -e  # Exit immediately if a command fails
	protoc -I${PROTO_DIR} --go_opt=module=$(shell go list -m) --go_out=. --go-grpc_opt=module=$(shell go list -m) --go-grpc_out=. ${PROTO_DIR}/*.proto

# Target to clean generated files
clean: ## Clean up generated files
	@echo "Cleaning up..."
	rm -f ${SERVER_BIN} ${CLIENT_BIN}
	rm -f ${PROTO_DIR}/*.pb.go

# Target to run the server
run_server: ## Run the Beef server
	@echo "Running Beef server..."
	./beef/bin/server

# Target to run the client
run_client: ## Run the Beef client
	@echo "Running Beef client..."
	go run ./beef/client/.

# Run the rest api
run_restapi: ## Run the rest api
	@echo "Running REST server..."
	restapi/_bin/restapi

build_restapi:
	go  build -o restapi/_bin/restapi ./restapi/_cmd/.

build_server:
	go build -o ./beef/_bin/server ./beef/server/.


# Display basic help information
help: ## Show this help
	@echo "Available commands:"
	@echo "  all            - Build the Beef server and client"
	@echo "  build          - Generate protobuf files and build server/client"
	@echo "  clean          - Clean generated files"
	@echo "  run_server     - Run the Beef server"
	@echo "  run_client     - Run the Beef client"
	@echo "  build_decoder     - Generate the decoder protobuf files and build server/client"
