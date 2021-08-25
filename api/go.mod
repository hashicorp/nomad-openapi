module github.com/hashicorp/nomad-openapi/api

go 1.16

replace github.com/hashicorp/nomad-openapi/v1/client => ../v1/client

require (
	github.com/hashicorp/cronexpr v1.1.1
	github.com/hashicorp/nomad v1.1.3
	github.com/hashicorp/nomad-openapi/v1/client v0.0.0-20210825113802-a22d652fa3d3 // indirect
	github.com/hashicorp/nomad/api v0.0.0-20210816133024-fdb868400424 // indirect
	github.com/stretchr/testify v1.7.0
)
