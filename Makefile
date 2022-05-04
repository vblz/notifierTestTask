.PHONY: list test build clean

all: lint test build

lint:
	golangci-lint run
	cd notifier && golangci-lint run

build:
	go build -o ./notifier.out .

test:
	go test -cover -v -race ./...
	cd notifier && go test -cover -v -race ./...

clean:
	rm -f notifier.out
