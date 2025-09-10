# Makefile for Laundry Backend

# Variables
BINARY=laundry-backend
MAIN_DIR=cmd

# Build
build:
	go build -o ${BINARY} ${MAIN_DIR}/main.go

# Run
run:
	go run ${MAIN_DIR}/main.go

# Install dependencies
deps:
	go mod tidy

# Test
test:
	go test -v ./...

# Clean
clean:
	rm -f ${BINARY}

# Run with race detector
race:
	go run -race ${MAIN_DIR}/main.go

# Build for different platforms
build-linux:
	GOOS=linux GOARCH=amd64 go build -o ${BINARY}-linux ${MAIN_DIR}/main.go

build-windows:
	GOOS=windows GOARCH=amd64 go build -o ${BINARY}-windows.exe ${MAIN_DIR}/main.go

build-mac:
	GOOS=darwin GOARCH=amd64 go build -o ${BINARY}-mac ${MAIN_DIR}/main.go

.PHONY: build run deps test clean race build-linux build-windows build-mac