goose_start = source .env && GOOSE_DRIVER=$$GOOSE_DRIVER GOOSE_DBSTRING=$$GOOSE_DBSTRING GOOSE_MIGRATION_DIR=$$GOOSE_MIGRATION_DIR goose -dir migrations

prepare:
	bash ./scripts/prepare.sh
dev:
	gow -c -v run .
templ:
	gow -c -v run ./templates
test-unit:
	gow -c -v test -v -count=1 ./tests/unit
test-integration:
	make migration-reset && make migration-up && gow -c -v test -v -count=1 ./tests/integration
test-e2e:
	gow -c -v test -v -count=1 ./tests/e2e
services-up:
	docker compose up -d
services-down:
	docker compose down
migration-create:
	$(goose_start) create $(name) sql
migration-up:
	$(goose_start) up
migration-down:
	$(goose_start) down
migration-reset:
	$(goose_start) reset
migration-status:
	$(goose_start) status

