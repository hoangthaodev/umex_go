SERVER = ./cmd/server/main.go

GOOSE_DRIVER ?= mysql
GOOSE_DBSTRING ?= "root:umex@123@456@tcp(127.0.0.1:3306)/umexdb"
GOOSE_MIGRATION_DIR ?= sql/schema

run:
	@go run $(SERVER)

docker_build:
	docker-compose up -d --build
	docker-compose ps
docker_stop:
	docker-compose down

upse:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) up
downse:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) down
resetse:
	@GOOSE_DRIVER=$(GOOSE_DRIVER) GOOSE_DBSTRING=$(GOOSE_DBSTRING) goose -dir=$(GOOSE_MIGRATION_DIR) reset

sqlgen:
	sqlc generate

.PHONY: run upse downse resetse docker_build docker_stop

.PHONY: air
