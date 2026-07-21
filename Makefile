COMPOSE=docker compose
INFRA=postgres rabbitmq jaeger prometheus grafana
APPS=product-service customer-service order-service payment-service api-gateway frontend
SERVICES=mc-product-service mc-customer-service mc-order-service mc-payment-service
GOOSE=$(shell go env GOPATH)/bin/goose

.PHONY: up down infra-up apps-up build logs ps clean db-init migrate-up migrate-down migrate-status install-goose ensure-goose

## Start everything: infra, databases, migrations and apps
up: infra-up db-init migrate-up apps-up
	@echo ""
	@echo "frontend    -> http://localhost:3000"
	@echo "api-gateway -> http://localhost:8080"
	@echo "rabbitmq    -> http://localhost:15672 (guest/guest)"
	@echo "jaeger      -> http://localhost:16686"
	@echo "prometheus  -> http://localhost:9090"
	@echo "grafana     -> http://localhost:3001 (admin/admin)"

## Start infra only and wait for the healthchecks to pass
infra-up:
	$(COMPOSE) up -d --wait $(INFRA)

## Start the applications (builds images when missing)
apps-up:
	$(COMPOSE) up -d --build $(APPS)

## Create the databases from scripts/init.sql on a running postgres.
## Postgres only auto-runs init.sql on a fresh volume, so this makes it repeatable.
db-init:
	$(COMPOSE) exec -T postgres psql -v ON_ERROR_STOP=1 -U postgres -d postgres \
		-f /docker-entrypoint-initdb.d/init.sql

build:
	$(COMPOSE) build

down:
	$(COMPOSE) down

## Stop everything and drop the postgres volume
clean:
	$(COMPOSE) down -v --remove-orphans

logs:
	$(COMPOSE) logs -f $(APPS)

ps:
	$(COMPOSE) ps

install-goose:
	go install github.com/pressly/goose/v3/cmd/goose@latest

ensure-goose:
	@test -x $(GOOSE) || $(MAKE) install-goose

## Run goose migrations for every service against localhost:5432
migrate-up: ensure-goose
	@for svc in $(SERVICES); do echo "==> $$svc"; $(MAKE) -C $$svc migrate-up || exit 1; done

migrate-down: ensure-goose
	@for svc in $(SERVICES); do echo "==> $$svc"; $(MAKE) -C $$svc migrate-down || exit 1; done

migrate-status: ensure-goose
	@for svc in $(SERVICES); do echo "==> $$svc"; $(MAKE) -C $$svc migrate-status || exit 1; done
