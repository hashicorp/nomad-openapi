## What is the OpenAPI Specification?

- A standard, language-agnostic interface to HTTP APIs
  that allows both humans and computers to discover and
  understand the capabilities of the service without access
  to source code, documentation, or through network traffic
  inspection.
- What's it look like?

## Quick tour of the repo

- Generator package
- Makefile
- Output
  - Client code
  - Docs
- V1 Reference implementation

## Workflow for updating

- `make spec` - Update/Generate the specification
- `make v1` - Update/Generate clients for each language

## Workflow for adding a new language

### Add a config for your language

```yaml
# File is used by the openapi-generator cli to generate a csharp-netcore client.
generatorName: csharp-netcore
gitUserId: hashicorp
gitRepoId: nomad-openapi/clients/csharp-netcore/v1
inputSpec: /local/v1/openapi.yaml
outputDir: /local/clients/csharp-netcore/v1
packageName: Nomad.Client
packageVersion: 1.1.4
hideGenerationTimestamp: true
licenseId: Mozilla Public License Version 2.0
```
### Makefile updates

```shell
	@docker run \
		--rm \
		--volume "$(shell pwd):/local" \
		$(DOCKER_IMAGE) batch --clean /local/clients/dotnet-core/v1/config.yaml
```

### Build & view output
  - Folder structure
  - Docs

## Help wanted

- New languages & tests
- Open issues
