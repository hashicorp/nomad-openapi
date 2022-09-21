VERSION=v5.2.0
DOCKER_IMAGE="openapitools/openapi-generator-cli:$(VERSION)"
THIS_OS := $(shell go env GOOS)
GENERATE_CLIENT = docker run --rm --volume "$(PWD):/local" $(DOCKER_IMAGE) batch --clean

CGO_ENABLED = 1
CGO_CFLAGS ?=
SDKROOT ?=

ifeq (darwin,$(THIS_OS))
spec: CGO_ENABLED = 0
CGO_CFLAGS=-Wno-undef-prefix
endif

.PHONY: spec
spec: CC ?= $(shell go env CC)
spec: GOOS=$(shell go env GOOS)
spec:
	(cd generator && CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) CC=$(CC) CGO_CFLAGS=$(CGO_CFLAGS) go build -o bin/generator && ./bin/generator ../v1/openapi.yaml)

.PHONY: test
test:
	(cd ./v1 && go test ./... -v -count=1)


CLIENTS = v1-go v1-java v1-javascript v1-python \
		v1-rust-reqwest v1-rust-hyper v1-ruby v1-typescript

.PHONY: openapi
v1: spec $(CLIENTS)
	@echo "==> Building v1 OpenAPI Specification and clients..."

# these have to be before the v1-% target to special case the subfolder names
.PHONY: v1-rust-%
v1-rust-%:
	@$(GENERATE_CLIENT) /local/clients/rust/$*/v1/config.yaml

.PHONY: v1-%
v1-%:
	@$(GENERATE_CLIENT) /local/clients/$*/v1/config.yaml


.PHONY: update-nomad
update-nomad:
	(cd generator && go get github.com/hashicorp/nomad)
	(cd generator && go get github.com/hashicorp/nomad/api)
	(cd generator && go mod tidy)
	@go get github.com/hashicorp/nomad
	@go mod tidy
