BINARY_NAME=restack-get-started
VERSION=1.0.0
BUILD_DIR=build

.PHONY: all linux macos clean

all: linux macos

linux:
	GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 main.go

macos:
	GOOS=darwin GOARCH=amd64 go build -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64 main.go

clean:
	rm -rf $(BUILD_DIR)
