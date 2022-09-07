VERSION=v5.4.0
DOCKER_IMAGE="openapitools/openapi-generator-cli:$(VERSION)"
THIS_OS := $(shell go env GOOS)

CGO_ENABLED = 1
CGO_CFLAGS ?=
SDKROOT ?=

CLIENTS = $(shell find clients -type f -name 'config.yaml' -exec dirname "{}" \; | sed 's_/v1__g' | sed 's_clients/__g')

.PHONY: v1 client/go client/java client/javascript/ client/python client/rust/reqwest client/rust/hyper client/ruby client/typescript

ifeq (darwin,$(THIS_OS))
spec: CGO_ENABLED = 0
CGO_CFLAGS=-Wno-undef-prefix
endif

spec: CC ?= $(shell go env CC)
spec: GOOS=$(shell go env GOOS)
spec:
	(cd generator && CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) CC=$(CC) CGO_CFLAGS=$(CGO_CFLAGS) go build -o bin/generator && ./bin/generator ../v1/openapi.yaml)

test:
	(cd ./v1 && go test ./... -v -count=1)

v1: test spec client/go client/java client/javascript/ client/python client/rust/reqwest client/rust/hyper client/ruby client/typescript

client/go:
	@echo "==> Building nomad-openapi client for go..."
	@docker run \
		--rm \
		--volume "$(PWD):/local" \
		 $(DOCKER_IMAGE) batch --clean /local/clients/go/v1/config.yaml

client/java:
	@echo "==> Building nomad-openapi client for java..."
	@docker run \
		--rm \
		--volume "$(PWD):/local" \
		$(DOCKER_IMAGE) batch --clean /local/clients/java/v1/config.yaml

client/javascript:
	@echo "==> Building nomad-openapi client for javascript..."
	@docker run \
		--rm \
		--volume "$(PWD):/local" \
		$(DOCKER_IMAGE) batch --clean /local/clients/javascript/v1/config.yaml

client/python/:
	@echo "==> Building nomad-openapi client for python using the stable generator..."
	@docker run \
		--rm \
		--volume "$(PWD):/local" \
		$(DOCKER_IMAGE) batch --clean /local/clients/python/v1/config.yaml

client/rust/reqwest:
	@echo "==> Building nomad-openapi client for rust using the reqwest generator..."
	@docker run \
		--rm \
		--volume "$(PWD):/local" \
		$(DOCKER_IMAGE) batch --clean /local/clients/rust/reqwest/v1/config.yaml

client/rust/hyper:
	@echo "==> Building nomad-openapi client for rust using the hyper generator..."
	@docker run \
		--rm \
		--volume "$(PWD):/local" \
		$(DOCKER_IMAGE) batch --clean /local/clients/rust/hyper/v1/config.yaml

client/ruby:
	@echo "==> Building nomad-openapi client for ruby..."
	@docker run \
		--rm \
		--volume "$(shell pwd):/local" \
		$(DOCKER_IMAGE) batch --clean /local/clients/ruby/v1/config.yaml

client/typescript:
	@echo "==> Building nomad-openapi client for typescript..."
	@docker run \
		--rm \
		--volume "$(shell pwd):/local" \
		$(DOCKER_IMAGE) batch --clean /local/clients/typescript/v1/config.yaml