.PHONY: spec
spec:
	(cd generator && CGO_ENABLED=0 GOOS=$(shell go env GOOS) go build -o bin/generator && ./bin/generator ../v1/openapi.yaml)

.PHONY: test
test:
	(cd ./v1 && go test ./... -v -count=1)

docker_spec:
	docker run --rm -v "$(PWD)":/usr/src -w /usr/src/generator golang:1.17 go build -o ../bin/generator
	./bin/generator v1/openapi.yaml

docker_v1: docker_spec
	@echo "==> Building v1 OpenAPI Specification and clients..."
	@docker run --rm -v "$(PWD):/local" \
		openapitools/openapi-generator-cli:v5.2.0 batch /local/clients/go/v1/config.yaml
	@docker run --rm -v "$(PWD):/local" \
		openapitools/openapi-generator-cli:v5.2.0 batch /local/clients/java/v1/config.yaml
	@docker run --rm -v "$(PWD):/local" \
		openapitools/openapi-generator-cli:v5.2.0 batch /local/clients/javascript/v1/config.yaml
	@docker run --rm -v "$(PWD):/local" \
		openapitools/openapi-generator-cli:v5.2.0 batch /local/clients/python/v1/config.yaml
	@docker run --rm -v "$(PWD):/local" \
		openapitools/openapi-generator-cli:v5.2.0 batch /local/clients/rust/reqwest/v1/config.yaml
	@docker run --rm -v "$(PWD):/local" \
		openapitools/openapi-generator-cli:v5.2.0 batch /local/clients/rust/hyper/v1/config.yaml

.PHONY: openapi
v1: spec
	@echo "==> Building v1 OpenAPI Specification and clients..."
	@docker run \
		--rm \
		--volume "$(shell pwd):/local" \
		openapitools/openapi-generator-cli:v5.2.0 batch /local/clients/go/v1/config.yaml
	@docker run \
    		--rm \
    		--volume "$(shell pwd):/local" \
    		openapitools/openapi-generator-cli:v5.2.0 batch /local/clients/java/v1/config.yaml
	@docker run \
    		--rm \
    		--volume "$(shell pwd):/local" \
    		openapitools/openapi-generator-cli:v5.2.0 batch /local/clients/javascript/v1/config.yaml
	@docker run \
			--rm \
			--volume "$(shell pwd):/local" \
			openapitools/openapi-generator-cli:v5.2.0 batch /local/clients/python/v1/config.yaml
	@docker run \
			--rm \
			--volume "$(shell pwd):/local" \
			openapitools/openapi-generator-cli:v5.2.0 batch /local/clients/rust/reqwest/v1/config.yaml
	@docker run \
			--rm \
			--volume "$(shell pwd):/local" \
			openapitools/openapi-generator-cli:v5.2.0 batch /local/clients/rust/hyper/v1/config.yaml