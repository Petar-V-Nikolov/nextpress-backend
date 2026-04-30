#!/usr/bin/env bash
# shellcheck shell=bash
# Resolve APP_PORT from repo .env (default 9090). Expects NP_ROOT set to repo root.

np_app_port() {
  local root="${NP_ROOT:-.}"
  local p
  p="$(awk -F= '/^APP_PORT=/{print $2; exit}' "$root/.env" 2>/dev/null | tr -d '[:space:]')"
  if [[ -z "$p" ]]; then
    echo "9090"
  else
    echo "$p"
  fi
}
