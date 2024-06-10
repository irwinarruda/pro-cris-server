dev:
	air .
dev-templ:
	templ generate --watch --proxy="http://localhost:8080" --cmd="go run ."
test-unit:
	go test -v ./tests/unit
test-integration:
	go test -v ./tests/integration
run:
	go run .
