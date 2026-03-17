#!/usr/bin/env bash
set -euo pipefail

export APP_ENV=${APP_ENV:-development}

go run ./cmd/api

