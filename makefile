goose_start = dotenv -e .env -- goose -dir ./migrations
watch_start = dotenv -e .env -- node ./external/watch/src/index.js

prepare:
	zsh ./scripts/prepare.sh
dev:
	$(watch_start) go run .
templ:
	$(watch_start) go run ./templates
test-unit:
	$(watch_start) go test -v -count=1 ./tests/unit
test-integration:
	make migration-reset && make migration-up && $(watch_start) go test -v -count=1 ./tests/integration
test-e2e:
	$(watch_start) go test -v -count=1 ./tests/e2e
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

