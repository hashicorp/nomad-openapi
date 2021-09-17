package main

import (
	"reflect"
)

type schemaType string

const (
	arraySchema  schemaType = "array"
	objectSchema            = "object"
	stringSchema            = "string"
	numberSchema            = "number"
	boolSchema              = "boolean"
	intSchema               = "integer"
	nilSchema               = ""
)

type parameter struct {
	Id          string
	Name        string
	SchemaType  schemaType
	In          parameterLocation
	Style       parameterStyle
	Explode     bool
	Description string
	Required    bool
}

type parameterLocation string
type parameterStyle string

const (
	inHeader parameterLocation = "header"
	inQuery                    = "query"
	inPath                     = "path"
	inCookie                   = "cookie"
)

// See for explanation https://stackoverflow.com/questions/62527254/how-does-open-api-3-0-support-a-single-query-param-key-with-multiple-values
const (
	formStyle parameterStyle = "form"
)

type responseHeader struct {
	Name        string
	SchemaType  schemaType
	Description string
}

type requestBody struct {
	SchemaType schemaType
	Model      reflect.Type
}

type content struct {
	SchemaType schemaType
	Model      reflect.Type
}

type response struct {
	Name        string
	Description string
}

type responseConfig struct {
	Code     int
	Response *response
	Content  *content
	Headers  []*responseHeader
}

type operation struct {
	Method      string
	Handler     string
	Tags        []string
	OperationId string
	Summary     string
	Description string
	RequestBody *requestBody
	Parameters  []*parameter
	Responses   []*responseConfig
}

type apiPath struct {
	Template   string
	Operations []*operation
}
