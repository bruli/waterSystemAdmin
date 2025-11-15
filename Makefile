SHELL := /usr/bin/env bash

APP=admin

DOCKER_COMPOSE=COMPOSE_BAKE=true docker compose

.DEFAULT_GOAL := help

.PHONY: docker-logs
docker-logs:
	@set -euo pipefail; \
	echo "👀 Showing logs for container $(APP) (CTRL+C to exit)..."; \
	docker logs -f $(APP)

.PHONY: docker-up
docker-up:
	@set -euo pipefail; \
	echo "🚀 Starting services with Docker Compose..."; \
	${DOCKER_COMPOSE} up -d --build $(APP)

.PHONY: docker-down
docker-down:
	@set -euo pipefail; \
	echo "🛑 Stopping and removing Docker Compose services..."; \
	${DOCKER_COMPOSE} down

.PHONY: docker-ps
docker-ps:
	@set -euo pipefail; \
	echo "📋 Active services:"; \
	${DOCKER_COMPOSE} ps

.PHONY: docker-exec
docker-exec:
	@set -euo pipefail; \
	echo "🔎 Opening shell inside container..."; \
	docker exec -it $(APP) bash

.PHONY: lint
lint:
	@set -euo pipefail; \
	echo "🚀 Executing golangci-lint..."; \
    go tool golangci-lint run ./...

.PHONY: fumpt
fumpt:
	@set -euo pipefail; \
	echo "👉 Formating code with gofumpt..."; \
	go tool gofumpt -w -l .

.PHONY: check
check: fumpt lint

.PHONY: edit-vault
edit-vault:
	@set -euo pipefail; \
	echo "🏗️  Editing vault file"; \
	ansible-vault edit devops/ansible/inventories/production/group_vars/raspberry_water_system_admin/vault.yml

.PHONY: build
build:
	@set -euo pipefail; \
	echo "🏗️ Building ARM64 binary for Raspberry Pi..."; \
	GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -a -ldflags "-s -w" -buildvcs=false -o devops/ansible/assets/server ./cmd/server/

.PHONY: deploy
deploy: build
	@set -euo pipefail; \
	echo "🚚 Deploying with Ansible (production inventory)..."; \
	ansible-playbook -i devops/ansible/inventories/production/hosts devops/ansible/deploy.yml --ask-vault-pass

help:
	@echo "Available commands:"
	@grep -E '^[a-zA-Z_-]+:' Makefile | awk -F':' '{print "  - " $$1}'