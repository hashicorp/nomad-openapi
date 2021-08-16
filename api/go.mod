module github.com/hashicorp/nomad-openapi/api

go 1.16

replace github.com/hashicorp/nomad-openapi/v1/client => ../v1/client

require (
	github.com/hashicorp/cronexpr v1.1.1 // indirect
	github.com/hashicorp/nomad-openapi/v1/client v0.0.0
)
