# Basic go commands
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

# Binary names
BINARY_NAME=kyber

.DEFAULT_GOAL := run.all

.PHONY: build
build:
	$(GOBUILD) -o bin/$(BINARY_NAME) -v kyber-api.go

.PHONY: run
run: build
	chmod +x ./bin/$(BINARY_NAME)
	./bin/$(BINARY_NAME) run

.PHONY: run.all
run.all: build run

