APP=admin

DOCKER_COMPOSE=COMPOSE_BAKE=true docker compose

define help
Usage: make <command>
Commands:
   help:                      Show this help information
   clean:                     To clean code
   docker-up:                 Start docker containers
   docker-down:               Stop docker containers
   docker-ps:                 To watch all docker containers
   docker-exec                To entry into water system container
   encryptVault:              Encrypt vault secret file
   decryptVault:              Decrypt vault secret file
   build:                     Compile the project
   deploy:                    Deploy the code to raspberry
endef
export help

.PHONY: help
help:
	@echo "$$help"

.PHONY: clean
clean:
	go fmt ./...

.PHONY: docker-logs
docker-logs:
	docker logs -f $(APP)

.PHONY: docker-up
docker-up:
	${DOCKER_COMPOSE} up -d --build $(APP)

.PHONY: docker-down
docker-down:
	${DOCKER_COMPOSE} down

.PHONY: docker-ps
docker-ps:
	${DOCKER_COMPOSE} ps

.PHONY: docker-exec
docker-exec:
	docker exec -it $(APP) bash

.PHONY: lint
lint:
	go tool golangci-lint run

.PHONY: fumpt
fumpt:
	go tool gofumpt -w -l .

.PHONY: encryptVault
encryptVault:
	ansible-vault encrypt --vault-id raspberry_water_system_admin@devops/ansible/password devops/ansible/inventories/production/group_vars/raspberry_water_system_admin/vault.yml

.PHONY: decryptVault
decryptVault:
	ansible-vault decrypt --vault-id raspberry_water_system_admin@devops/ansible/password devops/ansible/inventories/production/group_vars/raspberry_water_system_admin/vault.yml

.PHONY: build
build:
	@make clean
	GOOS=linux GOARCH=arm GOARM=6 go build -a -ldflags "-s -w" -tags prod -buildvcs=false -o devops/ansible/assets/server ./cmd/server/

.PHONY: deploy
deploy: build decryptVault
	ansible-playbook -i devops/ansible/inventories/production/hosts devops/ansible/deploy.yml
	@make encryptVault

