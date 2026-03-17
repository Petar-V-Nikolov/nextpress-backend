APP_NAME := nextpress-backend
CMD_PATH := ./cmd/nextpress

.PHONY: run-local build test lint up-local down-local

# run-local starts the API using the Go toolchain directly.
run-local:
	go run $(CMD_PATH)

# build produces a production-ready binary in the ./bin directory.
build:
	mkdir -p bin
	go build -o bin/$(APP_NAME) $(CMD_PATH)

# test runs the Go test suite.
test:
	go test ./...

# lint is a placeholder for future linters (golangci-lint, etc.).
lint:
	@echo "linting not configured yet"

# up-local brings up local infrastructure (Postgres) via docker-compose.
up-local:
	docker compose -f deployments/docker/docker-compose.local.yml up -d

# down-local stops local infrastructure.
down-local:
	docker compose -f deployments/docker/docker-compose.local.yml down

APP_NAME=nextpress
CMD_PATH=cmd/api/main.go

run:
	go run $(CMD_PATH)

build:
	go build -o bin/$(APP_NAME) $(CMD_PATH)

test:
	go test ./...

tidy:
	go mod tidy

fmt:
	go fmt ./...

vet:
	go vet ./...

lint:
	golangci-lint run