dev:
	templ generate --watch --proxy="http://localhost:8080" --cmd="go run ."
run:
	go run .
test:
	echo "Started testing"
