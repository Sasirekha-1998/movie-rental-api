APP_NAME := movie-rental-api
ENTRY_POINT := cmd/api/main.go
BUILD_DIR := bin

.PHONY: run test build clean

run:
	go run $(ENTRY_POINT)

build:
	go build -o $(BUILD_DIR)/$(APP_NAME) $(ENTRY_POINT)

test:
	go test ./tests/...

clean:
	rm -rf $(BUILD_DIR)
