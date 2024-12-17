.PHONY: build test run clean

BUILD_DIR := build
BINARY := $(BUILD_DIR)/fizzbuzz

build:
	@mkdir -p $(BUILD_DIR)
	go build -o $(BINARY) ./cmd/fizzbuzz

test:
	go test ./...

run: build
	$(BINARY)

clean:
	rm -rf $(BUILD_DIR)
