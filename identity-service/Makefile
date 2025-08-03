.PHONY: ensure-mod lint install-tools test run-local

install-tools:
	@echo "> Installing tools..."
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@1.59.1

lint:
	@echo "> Checking codes..."
	golangci-lint run --verbose --out-format code-climates | jq -r ''

test:
	@echo "> Running tests..."
	go test -cover -race -v ./...

run-local: ensure-mod
	cd cmd/identity-service && go run main.go

ensure-mod:
	go mod download
	go mod vendor