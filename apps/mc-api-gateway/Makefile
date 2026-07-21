APP_NAME=mc-api-gateway
MAIN_PATH=cmd/http/main.go

.PHONY: run build test clean

run:
	go run $(MAIN_PATH)

build:
	go build -o bin/$(APP_NAME) $(MAIN_PATH)

test:
	go test ./... -v

clean:
	rm -rf bin/
