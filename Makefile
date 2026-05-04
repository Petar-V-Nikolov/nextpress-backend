# NextPressKit — three public Make targets. Everything else: ./scripts/nextpresskit help

.DEFAULT_GOAL := default

.PHONY: default setup run postman-sync

default:
	@echo "NextPressKit — use:"
	@echo "  make setup         text menu (TTY) or NP_SETUP_NONINTERACTIVE=1 for linear bootstrap only"
	@echo "  make run           API in the foreground"
	@echo "  make postman-sync  Refresh gitignored postman/ from templates + .env"
	@echo "Advanced: ./scripts/nextpresskit help"

setup:
	@bash scripts/nextpresskit setup

run:
	@bash scripts/nextpresskit run

postman-sync:
	@bash scripts/nextpresskit postman-sync
