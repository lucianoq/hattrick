.PHONY: all build fmt fmt-check vet lint test cover tidy clean check

all: check

build:
	go build ./...

fmt:
	gofmt -w .

fmt-check:
	@test -z "$$(gofmt -l .)" || (echo "gofmt needs to be run on:"; gofmt -l .; exit 1)

vet:
	go vet ./...

lint:
	golangci-lint run ./...

test:
	go test ./...

cover:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html

tidy:
	go mod tidy

clean:
	rm -f coverage.out coverage.html

check: fmt-check vet lint test
