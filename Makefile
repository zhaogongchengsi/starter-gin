APP_NAME := starter
BUILD_FLAGS :=
GO_FILES := $(wildcard cmd/*/main.go)

.PHONY: all
all: clean build-windows build-linux build-macos

.PHONY: build-windows
build-windows:
	@for file in $(GO_FILES); do \
		dir=$$(dirname $$file); \
		name=$$(basename $$dir); \
		echo "Building $$file for Windows..."; \
		GOOS=windows GOARCH=amd64 go build $(BUILD_FLAGS) -o bin/$$name.exe $$file; \
	done

.PHONY: build-linux
build-linux:
	@for file in $(GO_FILES); do \
		dir=$$(dirname $$file); \
		name=$$(basename $$dir); \
		echo "Building $$file for Linux..."; \
		GOOS=linux GOARCH=amd64 go build $(BUILD_FLAGS) -o bin/$$name $$file; \
	done

.PHONY: build-macos
build-macos:
	@for file in $(GO_FILES); do \
		dir=$$(dirname $$file); \
		name=$$(basename $$dir); \
		echo "Building $$file for macOS..."; \
		GOOS=darwin GOARCH=amd64 go build $(BUILD_FLAGS) -o bin/$$name $$file; \
	done

.PHONY: clean
clean:
	rm -rf bin/*

.PHONY: test
test:
	go test ./...