// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"net/http"

	"github.com/hashicorp/nomad/api"
)

func (v *v1api) getVariablesPaths() []*apiPath {
	tags := []string{"Variables"}

	return []*apiPath{
		// s.mux.Handle("/v1/vars", wrapCORS(s.wrap(s.VariablesListRequest)))
		{
			Template: "/vars",
			Operations: []*operation{
				newOperation(http.MethodGet,
					"VariablesListRequest",
					tags,
					"GetVariablesListRequest",
					nil,
					defaultQueryOpts,
					newResponseConfig(200,
						arraySchema,
						api.VariableMetadata{},
						defaultQueryMeta,
						"GetVariablesListResponse",
					),
				),
			},
		},

		// s.mux.Handle("/v1/var/", wrapCORSWithAllowedMethods(s.wrap(s.VariableSpecificRequest), "HEAD", "GET", "PUT", "DELETE"))
		{
			Template: "/var/{path}",
			Operations: []*operation{
				newOperation(http.MethodGet,
					"variableQuery",
					tags,
					"GetVariableQuery",
					nil,
					appendParams(defaultQueryOpts, &variablePath),
					newResponseConfig(200,
						objectSchema,
						api.Variable{},
						defaultQueryMeta,
						"GetVariableResponse",
					),
				),
				newOperation(http.MethodPost,
					"variableUpsert",
					tags,
					"PostVariable",
					newRequestBody(objectSchema, api.Variable{}),
					appendParams(defaultWriteOpts, &variablePath, &variableCAS),
					newResponseConfig(200,
						objectSchema,
						api.Variable{},
						defaultWriteMeta,
						"PostVariableResponse",
					),
				),
				newOperation(http.MethodPut,
					"variableUpsert",
					tags,
					"PutVariable",
					newRequestBody(objectSchema, api.Variable{}),
					appendParams(defaultWriteOpts, &variablePath, &variableCAS),
					newResponseConfig(200,
						objectSchema,
						api.Variable{},
						defaultWriteMeta,
						"PutVariableResponse",
					),
				),
				newOperation(http.MethodDelete,
					"variableDelete",
					tags,
					"DeleteVariable",
					newRequestBody(objectSchema, api.Variable{}),
					appendParams(defaultWriteOpts, &variablePath, &variableCAS),
					newResponseConfig(200,
						nilSchema,
						nil, defaultWriteMeta,
						"DeleteVariableResponse",
					),
				),
			},
		},
	}
}
