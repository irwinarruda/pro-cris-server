dev:
	air .
dev-templ:
	templ generate --watch --proxy="http://localhost:8080" --cmd="go run ."
test:
	go test -v ./tests
run:
	go run .
