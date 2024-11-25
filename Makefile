BUILD_NAME := debug

.PHONY: build

build:
	CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-s -w -extldflags "-static"' -o build/$(BUILD_NAME) ./cmd/nautilus/main.go ;

test: build
	./build/debug --address http://127.0.0.1:8080