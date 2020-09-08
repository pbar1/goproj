export CGO_ENABLED     ?= 0
export DOCKER_BUILDKIT ?= 1

BIN        := $(shell basename $(PWD))
VERSION    := $(shell git describe --tags --always --dirty)
LDFLAGS    := -ldflags="-s -w -X main.version=$(VERSION)"
IMAGE_REPO := docker.io/pbar1
IMAGE      := $(IMAGE_REPO)/$(BIN)

build:
	@go build -o bin/$(BIN) $(LDFLAGS) main.go

image:
	@docker build . --tag=$(IMAGE):latest

push: image
	@docker push $(IMAGE):latest

run: image
	@docker run --rm --publish=8080:8080

version:
	@echo $(VERSION)
