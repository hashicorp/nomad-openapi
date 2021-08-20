.PHONY: spec
spec:
	(cd generator && CGO_ENABLED=0 GOOS=$(shell go env GOOS) go build -o bin/generator && ./bin/generator ../v1/openapi.yaml)

.PHONY: test
test:
	(cd ./api && go test ./... -v -count=1)

.PHONY: openapi
openapi: spec
	@echo "==> Building OpenAPI Specification and test client..."
	@docker run \
		--rm \
		--volume "$(shell pwd):/local" \
		openapitools/openapi-generator-cli:v5.2.0 batch /local/v1/config.yaml

.PHONY: update-client
update-client:
	(cd ./api && go get github.com/hashicorp/nomad-openapi/v1/client)