.PHONY: spec
spec:
	(cd generator && CGO_ENABLED=0 GOOS=$(shell go env GOOS) go build -o bin/generator && ./bin/generator ../api/v1/openapi.yaml)
	#(cd asyncapi-generator && CGO_ENABLED=0 GOOS=$(shell go env GOOS) go build -o bin/asyncapi-generator && ./bin/asyncapi-generator ../api/v1/asyncapi.yaml)

.PHONY: test
test:
	(cd ./api/v1 && go test ./... -v -count=1)

.PHONY: openapi
openapi: spec
	@echo "==> Building OpenAPI Specification, test client, and proxy server..."
	@docker run \
		--rm \
		--volume "$(shell pwd):/local" \
		openapitools/openapi-generator-cli:v5.2.0 batch /local/api/v1/client-config.yaml
	@docker run \
		--rm \
		--volume "$(shell pwd):/local" \
		openapitools/openapi-generator-cli:v5.2.0 batch /local/api/v1/server-config.yaml
	(cd ./api/v1/proxyserver && goimports -w ./ && @docker build --network=host -t nomad-api-proxy .