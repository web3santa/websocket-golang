run:
    @./bin/chat

build:
    @go build -o bin/chat

test:
    @go test -v ./...