GOCMD=go
GOBUILD=$(GOCMD) build
GOGET=$(GOCMD) get
GOTEST=$(GOCMD) test

all: build

build:
	$(GOBUILD) -i

deps:
	$(GOGET) ./..

test: test-mock test-api

test-mock:
	$(GOTEST) -v --tags mock

test-api:
	$(GOTEST) -v