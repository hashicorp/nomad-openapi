<p align="center" style="text-align:center;">
  <a href="https://nomadproject.io">
    <img alt="HashiCorp Nomad logo" src="https://www.nomadproject.io/img/logo-hashicorp.svg" width="500" />
  </a>
</p>

# Overview

This repository contains the HashiCorp Nomad OpenAPI specification and related artifacts.

The OpenAPI specification defines a machine-readable schema for describing HTTP APIs.
From an OpenAPI specification, clients and servers for your project can be
_generated_ in a number of programming languages.

## Quick start

The latest version of the OpenAPI specification for the Version 1 Nomad HTTP API
can be found in this repository at [api/v1/openapi.yaml](https://github.com/hashicorp/nomad-openapi/blob/main/api/v1/openapi.yaml).
This file _is itself a generated file_ and _should not be edited directly_. You can
use this file to generate a client for the Nomad HTTP API in the language of your
choice.

## Using the Go Client

To use the Go client add a reference to it in your `go.mod` with `go get`.

```shell
$ go get github.com/hashicorp/nomad-openapi
```

You can use it from your client applications like this.

```go
package main

import (
	"fmt"
	"os"
	
	v1 "github.com/hashicorp/nomad-openapi/v1"
)

func main() {
	client, err := v1.NewClient()
	if err != nil {
		fmt.Println(err.Error())
	}

	jobName := "example"
	opts := v1.DefaultQueryOpts()
	
	job, meta, err := client.Jobs().GetJob(opts.Ctx(), jobName)
	if err != nil {
		os.Exit(1)
	}
	
	fmt.Println(*job.ID)
	fmt.Printf("%v", &meta)
}
```

## Environmental Configuration 

This client supports the following Nomad environment variables.

- `NOMAD_ADDR` - Required to overide the default of `http://127.0.0.1:4646`.
- `NOMAD_TOKEN` - Required with ACLs enabled. Can be overidden with `QueryOpts` or `WriteOpts`.
- `NOMAD_CACERT` - Required with TLS enabled.
- `NOMAD_CLIENT_CERT` - Required with TLS enabled.
- `NOMAD_CLIENT_KEY` - Required with TLS enabled.
- `NOMAD_REGION` - Required if you are calling a server in a region other than
  `global` as the `NOMAD_CACERT` SAN will follow the convention of `server.<region>.nomad`.

## Code generation
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

Generating a client in a different language should be as straightforward as
changing the `-g` argument in the command above to your language of choice. Check
the [OpenAPI Generators](https://openapi-generator.tech/docs/generators) page
for a full list of supported languages.

If you want to generate a client in a language not yet included in the `clients`
directory and submit it for inclusion in this repo, review the `Makefile` at the
root of this repository. Also review the `config.yaml` file in each client folder.
You will need to create a similar file that is specific to your language. Reference
the [OpenAPI Generator documentation](https://openapi-generator.tech/docs/configuration)
for more information on language specific configuration options.

This repository is currently _experimental_, and, as such, there is no guarantee
of support at this time.

## Community Guidelines

The HashiCorp community is a mixture of professionals, volunteers, students, and
employees who collaborate to make HashiCorp tools better. Community members play
a variety of roles, including mentor, teacher, or connector. If you would like to
contribute to HashiCorp products, review our community guidelines which can be
found online [here](https://www.hashicorp.com/community-guidelines).

## Issues and Questions

If you have any issues or questions using this package, please raise a Github issue
in this repository. Issues raised in the main Nomad repository will be redirected
here.

## Motivation

The OpenAPI specification is ideal if you are working in a greenfield scenario, and can
write your specification first. This _spec first_ approach is highly recommended and
widely supported by many of the tool vendors and experts in the OpenAPI space.

However, this approach leaves brownfield projects with an existing API that either
predates the OpenAPI specification, or for whatever reason was built with a _code first_
approach, to their own devices in terms of generating a specification from existing
code. The code found in the `generator` directory is one such device. It is highly,
experimental, and subject to change.

## Contributing

The `/v1/openapi.yaml` specification file is a generated file that should not be
edited manually.

This [video](https://youtu.be/x1AZbXiUENU) contains an overview of how to
contribute endpoint configurations. The slide deck for this video can be found
[here](https://docs.google.com/presentation/d/1h4OOjPFOHbDJsbtuQZRYDjotyBH1YZs7V8L7qmEjRXc/edit#slide=id.gf24f7d4584_1_20).

The [README](https://github.com/hashicorp/nomad-openapi/blob/main/generator/README.md)
file in the `generator` directory contains a detailed technical overview of the
generator package as well as instructions on _how to contribute_ to that package.

## Client and API packages are not officially supported packages

The `clients` and `v1` packages found in this repository are _not officially supported packages_.








