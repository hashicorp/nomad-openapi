# Overview

This repository contains the HashiCorp Nomad OpenAPI specification and related artifacts.

The OpenAPI specification defines a machine-readable schema for defining HTTP APIs.
From an OpenAPI specification, clients and servers can be _generated_ in a number
of programming languages.

The latest version of the OpenAPI specification for the Version 1 Nomad HTTP API
can be found in this repository at [v1/openapi.yaml](https://github.com/hashicorp/nomad-openapi/v1/openapi.yaml).
This file _is itself a generated file_ and _should not be edited directly_. You can
use this file to generate a client for the Nomad HTTP API in the language of your
choice.

The OpenAPI ecosystem has a number of existing code generators to choose from.
For example, this project uses the [OpenAPI Generator](https://openapi-generator.tech/)
project to generate a test client in Golang that we then use to validate our
specification.

This repository is currently _experimental_, and, as such, there is no guarantee
of support at this time.

## Why the generator package?

The OpenAPI specification is ideal if you are in a greenfield scenario, and can
write your specification first. This _spec first_ approach is highly recommended and
widely supported by many of the tool vendors and experts in the OpenAPI space.

Sadly, this approach leaves brownfield projects with an existing API that either
predates the OpenAPI specification, or for whatever reason was built with a _code first_
approach, to their own devices in terms of generating a specification from existing
code. The code found in the `generator` directory is one such device. It is highly,
experimental, and subject to change. The [README](https://github.com/hashicorp/nomad-openapi/generator/README.md)
file in that directory contains a detailed technical overview as well as instructions
on how to contribute.

## Client package is not an officially supported package

The client found in the [v1/client](https://github.com/hashicorp/nomad-openapi/v1/client)
directory is _not an officially supported_ package. It exists here to be consumed
by unit tests as a means of validating the specification produced by the `generator`
package.






