module github.com/hashicorp/nomad-openapi

go 1.16

// TODO: Test the replace idea to reference ../clients/go

require (
	github.com/hashicorp/cronexpr v1.1.1
	github.com/hashicorp/nomad v1.1.4
	github.com/hashicorp/nomad/api v0.0.0-20210816133024-fdb868400424 // indirect
	github.com/stretchr/testify v1.7.0
	golang.org/x/oauth2 v0.0.0-20210819190943-2bc19b11175f
)
