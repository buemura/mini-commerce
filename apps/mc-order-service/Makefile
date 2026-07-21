APP_NAME=mc-order-service
MAIN_PATH=cmd/grpc/main.go
MIGRATIONS_DIR=migrations
GOOSE=$(shell go env GOPATH)/bin/goose
DATABASE_URL?=postgresql://postgres:postgres@localhost:5432/edc_svc_order

.PHONY: run build test clean migrate-up migrate-down migrate-status migrate-create

run:
	go run $(MAIN_PATH)

build:
	go build -o bin/$(APP_NAME) $(MAIN_PATH)

test:
	go test ./... -v

clean:
	rm -rf bin/

migrate-up:
	$(GOOSE) -dir $(MIGRATIONS_DIR) postgres "$(DATABASE_URL)" up

migrate-down:
	$(GOOSE) -dir $(MIGRATIONS_DIR) postgres "$(DATABASE_URL)" down

migrate-status:
	$(GOOSE) -dir $(MIGRATIONS_DIR) postgres "$(DATABASE_URL)" status

migrate-create:
	$(GOOSE) -dir $(MIGRATIONS_DIR) create $(NAME) sql
