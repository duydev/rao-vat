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

deploy:
		docker build -t rao-vat .
		docker tag rao-vat asia-southeast1-docker.pkg.dev/rao-vat-428611/docker-repo/rao-vat
		docker push asia-southeast1-docker.pkg.dev/rao-vat-428611/docker-repo/rao-vat
		gcloud run deploy --image asia-southeast1-docker.pkg.dev/rao-vat-428611/docker-repo/rao-vat --platform managed --region asia-southeast1

default: all