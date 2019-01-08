all:
	@echo "no default"

.PHONY: start
start:
	@go run main.go

.PHONY: test
test:
	@go test -v ./...
