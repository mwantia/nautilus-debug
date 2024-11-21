.PHONY: build

build:
	go build -o build/nautilus-debug ./cmd/nautilus/main.go ;

test: build
	./build/nautilus-debug --address :12345