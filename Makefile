GO=go
GOFLAGS=-race
BIN=bin/swerve

all: build/local

build/local:
	$(GO) build -o $(BIN) $(GOFLAGS) main.go

run/dynamo: 
	docker-compose -f example/stack/stack.yml up