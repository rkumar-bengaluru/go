
PATHS = ./config

run:
	go run ./cmd
tests:
	@echo "Testing..."
	@go test -v -gcflags=-l -cover $(PATHS) -coverprofile=coverage.out
	@go tool cover -func=coverage.out
