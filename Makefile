run:
	go run .

install:
	rm -f go.mod go.sum
	go mod init github.com/newline-sandbox/go-chi-docgen-example
	go mod tidy

security:
	gosec ./...

lint:
	golangci-lint run -v

validate:
	gosec ./...
	golangci-lint run -v
	go test -v ./...
	go build -v ./...
