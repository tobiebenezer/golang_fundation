include .envrc
MIGRATION_PATH=./cmd/migration
DB_DSN="postgres://postgres:Awodumila@localhost:5432/gosocial?sslmode=disable"

.PHONY: create_migration
migration:
	@migrate create -ext sql -dir $(MIGRATION_PATH) $(filter-out $@,$(MAKECMDGOALS))

.PHONY: migrate_up
migrate_up:
	@migrate -path=$(MIGRATION_PATH) -database="$(DB_DSN)" up

migrate_up_all:
	@migrate -path=$(MIGRATION_PATH) -database="$(DB_DSN)" up

.PHONY: migrate_down
migrate_down:
	@migrate -path=$(MIGRATION_PATH) -database="$(DB_DSN)" down $(filter-out $@,$(MAKECMDGOALS))

.PHONY: load-env
load-env:
	@direnv allow



.PHONY: generate-doc
generate-doc:
	@swag init -g cmd/api/main.go -g ./cmd/api/api.go