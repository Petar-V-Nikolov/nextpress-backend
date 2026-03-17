# Makefile for nextpress-backend

# Default environment (used in future phases for migrations/seeds if needed)
APP_ENV ?= local

# Build variables
BINARY_NAME=server
MIGRATE_BINARY=migrate
SEED_BINARY=seed

.PHONY: all build run clean test migrate-up migrate-down migrate-down-all migrate-drop migrate-version seed seed-build db-fresh help tidy deps

## help: Display this help message
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@sed -n 's/^##//p' $(MAKEFILE_LIST) | column -t -s ':' | sed -e 's/^/ /'

## all: Build the server binary
all: build

## build: Build the API server binary (bin/server)
build:
	@echo "Building nextpress-backend server..."
	mkdir -p bin
	go build -o bin/$(BINARY_NAME) ./cmd/api
	@echo "Done."

## run: Run the API server directly with go run
run:
	go run ./cmd/api

## clean: Clean build artifacts
clean:
	rm -rf bin/
	go clean

# =============================================================================
# Database Seeding (placeholders until migrations/seeds are added)
# =============================================================================

## seed: Run seeders (not implemented yet)
seed:
	@echo "Seeders are not implemented yet for nextpress-backend."

## seed-build: Build the seed binary (not implemented yet)
seed-build:
	@echo "Seed binary is not implemented yet for nextpress-backend."

# =============================================================================
# Database Migrations (placeholders until migrations are added)
# =============================================================================

## migrate-up: Run all pending migrations (not implemented yet)
migrate-up:
	@echo "Migrations are not implemented yet for nextpress-backend."

## migrate-down: Rollback the last migration (not implemented yet)
migrate-down:
	@echo "Migrations are not implemented yet for nextpress-backend."

## migrate-version: Show current migration version (not implemented yet)
migrate-version:
	@echo "Migrations are not implemented yet for nextpress-backend."

## migrate-down-all: Rollback all migrations (not implemented yet)
migrate-down-all:
	@echo "Migrations are not implemented yet for nextpress-backend."

## migrate-drop: Drop all tables (not implemented yet)
migrate-drop:
	@echo "Migrations are not implemented yet for nextpress-backend."

## db-fresh: Drop all tables then run all migrations (not implemented yet)
db-fresh:
	@echo "Migrations are not implemented yet for nextpress-backend."

## test: Run tests
test:
	go test -v ./...

## test-coverage: Run tests with coverage
test-coverage:
	go test -cover ./...

## tidy: Tidy up dependencies
tidy:
	go mod tidy

## deps: Download dependencies
deps:
	go mod download
