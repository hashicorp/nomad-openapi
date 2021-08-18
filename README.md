# Overview

This repository contains the HashiCorp Nomad OpenAPI specification and related artifacts.

The OpenAPI specification defines a machine-readable schema for describing HTTP APIs.
From an OpenAPI specification, clients and servers for your project can be
_generated_ in a number of programming languages.

## Quick start

The latest version of the OpenAPI specification for the Version 1 Nomad HTTP API
can be found in this repository at [v1/openapi.yaml](https://github.com/hashicorp/nomad-openapi/blob/main/v1/openapi.yaml).
This file _is itself a generated file_ and _should not be edited directly_. You can
use this file to generate a client for the Nomad HTTP API in the language of your
choice.

The OpenAPI ecosystem has a number of existing code generators to choose from.
This repository uses the [OpenAPI Generator](https://openapi-generator.tech/)
project to generate a test client in Golang that is then used to validate the
generated specification.

To generate a `ruby` client using the OpenAPI Generator, you can use the following
docker command from the root of this repository.

```shell
$ mkdir -p /tmp/nomad-openapi && docker run \
	--rm \
	-v "$PWD:/local" \
	-v "/tmp/nomad-openapi:/output" \
	openapitools/openapi-generator-cli:v5.2.0 generate -i /local/v1/openapi.yaml -g ruby -o /output/
```

Generating in a client in a different language should be as straightforward as
changing the `-g` argument in the command above to your language of choice. Check
the [OpenAPI Generators](https://openapi-generator.tech/docs/generators) page
for a full list of supported languages. 

This repository is currently _experimental_, and, as such, there is no guarantee
of support at this time.

## Community Guidelines

The HashiCorp community is a mixture of professionals, volunteers, students, and
employees who collaborate to make HashiCorp tools better. Community members play
a variety of roles, including mentor, teacher, or connector. If you would like to
contribute to HashiCorp products, review our community guidelines which can be
found online [here](https://www.hashicorp.com/community-guidelines).

## Issues and Questions

If you have any issues or questions using this pacakge, please raise a Gitub issue
in this repository. Issues raised in the main Nomad repository will be redirected
here.

## Motivation

The OpenAPI specification is ideal if you are working in a greenfield scenario, and can
write your specification first. This _spec first_ approach is highly recommended and
widely supported by many of the tool vendors and experts in the OpenAPI space.

Sadly, this approach leaves brownfield projects with an existing API that either
predates the OpenAPI specification, or for whatever reason was built with a _code first_
approach, to their own devices in terms of generating a specification from existing
code. The code found in the `generator` directory is one such device. It is highly,
experimental, and subject to change.

## Contributing

The `v1/openapi.yaml` specification file is a generated file that should not be
edited manually. 

The [README](https://github.com/hashicorp/nomad-openapi/blob/main/generator/README.md)
file in the `generator` directory contains a detailed technical overview as well
as instructions on _how to contribute_.

## Client and API packages are not officially supported packages

The `client` and `api` found in this repository are _not officially supported packages_.
They exist here to be consumed by unit tests as a means of validating the
specification produced by the `generator` package.








