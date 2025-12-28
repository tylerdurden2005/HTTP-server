.PHONY: run test clean

run:
	go run ./cmd/server/main.go
test:
	go test ./internal/...
clean:
	go clean