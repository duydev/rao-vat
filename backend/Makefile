GOCMD=go
GOTEST=$(GOCMD) test
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOFMT=$(GOCMD) fmt
GOLINT=golangci-lint run --config .golangci.yml

all: test build

build:
		${GOBUILD} -o bin/www -v .

test:
		${GOTEST} -v .

clean:
		${GOCLEAN}
		rm -rf bin

lint:
		${GOLINT} .

fmt:
		${GOFMT} .

run:
		${GOBUILD} -o bin/www -v .
		./bin/www

deps:
		go mod tidy

default: all