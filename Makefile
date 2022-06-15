OPENAPI_GEN_VERSION=v6.0.0
DOCKER_IMAGE="openapitools/openapi-generator-cli:$(OPENAPI_GEN_VERSION)"
THIS_OS := $(shell go env GOOS)

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

.PHONY: openapi
v1: spec
	@echo "==> Building v1 OpenAPI Specification and clients..."
	@docker run \
		--rm \
		--volume "$(PWD):/local" \
		 $(DOCKER_IMAGE) batch --clean /local/clients/go/v1/config.yaml
	@docker run \
		--rm \
		--volume "$(PWD):/local" \
		$(DOCKER_IMAGE) batch --clean /local/clients/java/v1/config.yaml
	@docker run \
		--rm \
		--volume "$(PWD):/local" \
		$(DOCKER_IMAGE) batch --clean /local/clients/javascript/v1/config.yaml
	@docker run \
		--rm \
		--volume "$(PWD):/local" \
		$(DOCKER_IMAGE) batch --clean /local/clients/python/v1/config.yaml
	@docker run \
		--rm \
		--volume "$(PWD):/local" \
		$(DOCKER_IMAGE) batch --clean /local/clients/rust/reqwest/v1/config.yaml
	@docker run \
		--rm \
		--volume "$(PWD):/local" \
		$(DOCKER_IMAGE) batch --clean /local/clients/rust/hyper/v1/config.yaml
	@docker run \
		--rm \
		--volume "$(shell pwd):/local" \
		$(DOCKER_IMAGE) batch --clean /local/clients/ruby/v1/config.yaml
	@docker run \
		--rm \
		--volume "$(shell pwd):/local" \
		$(DOCKER_IMAGE) batch --clean /local/clients/typescript/v1/config.yaml


.PHONY: update-nomad
update-nomad:
	(cd generator && go get github.com/hashicorp/nomad)
	(cd generator && go get github.com/hashicorp/nomad/api)
	(cd generator && go mod tidy)
	@go get github.com/hashicorp/nomad
	@go mod tidy
