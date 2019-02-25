BASE_DIR=$(shell echo $$GOPATH)

GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test

all: build

build:
	@echo 'Building project...'
	$(GOCMD) build

test:
	$(GOTEST) -v --tags mock

test-api:
	$(GOTEST) -v