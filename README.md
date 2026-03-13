# NextPress Backend

NextPress is a modular CMS backend written in Go.

## Stack

- Go 1.26
- Gin HTTP Framework
- PostgreSQL
- GORM
- Zap Logger
- JWT Authentication

## Architecture

The project follows:

- Clean Architecture
- Modular Monolith
- Domain Driven Design

## Project Structure

cmd/api → application entry point

internal/config → configuration system  
internal/platform → shared infrastructure  
internal/modules → domain modules

pkg → reusable utilities

## Development

Run server:
