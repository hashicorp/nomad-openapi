.PHONY: spec
spec:
	(cd generator && CGO_ENABLED=0 GOOS=$(shell go env GOOS) go build -o bin/generator && ./bin/generator ../api/v1/openapi.yaml)

.PHONY: test
test:
	(cd ./api/v1 && go test ./... -v -count=1)

.PHONY: openapi
openapi: spec
	@echo "==> Building OpenAPI Specification and test client..."
	@docker run \
		--rm \
		--volume "$(shell pwd):/local" \
		openapitools/openapi-generator-cli:v5.2.0 batch /local/api/v1/config.yaml