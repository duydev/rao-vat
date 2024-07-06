GOCMD=go
GOTEST=$(GOCMD) test
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOFMT=$(GOCMD) fmt
GOLINT=golangci-lint run --config .golangci.yml
ENTRYPOINT=./cmd

all: test build

build:
		${GOBUILD} -o bin/www -v ${ENTRYPOINT}

test:
		${GOTEST} -v ${ENTRYPOINT}

clean:
		${GOCLEAN}
		rm -rf bin

lint:
		${GOLINT} ${ENTRYPOINT}

fmt:
		${GOFMT} ${ENTRYPOINT}

run:
		${GOBUILD} -o bin/www -v ${ENTRYPOINT}
		./bin/www

deps:
		go mod tidy

deploy:
		docker build -t duydev.io.vn/rao-vat .
		docker tag duydev.io.vn/rao-vat asia-southeast1-docker.pkg.dev/rao-vat-428611/docker-repo/rao-vat
		docker push asia-southeast1-docker.pkg.dev/rao-vat-428611/docker-repo/rao-vat
		gcloud run deploy rao-vat --image asia-southeast1-docker.pkg.dev/rao-vat-428611/docker-repo/rao-vat --platform managed --region asia-southeast1

dev:
		gow run ${ENTRYPOINT}

default: all