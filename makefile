goose_start = source .env && GOOSE_DRIVER=$$GOOSE_DRIVER GOOSE_DBSTRING=$$GOOSE_DBSTRING GOOSE_MIGRATION_DIR=$$GOOSE_MIGRATION_DIR goose -dir migrations

prepare:
	bash ./scripts/prepare.sh
dev:
	air .
dev-templ:
	templ generate --watch --proxy="http://localhost:8080" --cmd="go run ."
test-unit:
	go test -v ./tests/unit
test-integration:
	go test -v ./tests/integration
services-up:
	docker compose up -d
services-down:
	docker compose down -d
migration-create:
	$(goose_start) create $(name) sql
migration-up:
	$(goose_start) up
migration-down:
	$(goose_start) down
migration-status:
	$(goose_start) status

