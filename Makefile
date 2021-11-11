VERSION=v5.2.0
DOCKER_IMAGE="openapitools/openapi-generator-cli:$(VERSION)"

.PHONY: spec
spec:
	(cd generator && GOOS=$(shell go env GOOS) go build -o bin/generator && ./bin/generator ../v1/openapi.yaml)

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
