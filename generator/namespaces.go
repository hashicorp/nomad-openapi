package main

import (
	"net/http"

	"github.com/hashicorp/nomad/api"
)

func (v *v1api) getNamespacePaths() []*apiPath {
	tags := []string{"Namespaces"}

	// TODO: add to nomad-enterprise repo
	return []*apiPath{
		//s.mux.HandleFunc("/v1/namespaces", s.wrap(s.NamespacesRequest))
		{
			Template: "/namespaces",
			Operations: []*operation{
				newOperation(http.MethodGet, "NamespaceRequest", tags,
					"GetNamespaces",
					nil,
					queryOptions,
					newResponseConfig(200, arraySchema, api.Namespace{}, queryMeta,
						"GetNamespacesResponse"),
				),
			},
		},
		//s.mux.HandleFunc("/v1/namespace/", s.wrap(s.NamespaceSpecificRequest))
		{
			Template: "/namespace/{namespaceName}",
			Operations: []*operation{
				newOperation(http.MethodGet, "namespaceQuery", tags,
					"GetNamespace",
					nil,
					appendParams(queryOptions, &namespaceNameParam),
					newResponseConfig(200, objectSchema, api.Namespace{}, queryMeta,
						"GetNamespaceResponse"),
				),
				newOperation(http.MethodPost, "namespaceUpdate", tags,
					"PostNamespace",
					newRequestBody(objectSchema, api.Namespace{}),
					appendParams(writeOptions, &namespaceNameParam),
					newResponseConfig(200,
						nilSchema, nil, writeMeta,
						"PostNamespaceResponse"),
				),
				newOperation(http.MethodDelete, "namespaceDelete", tags,
					"DeleteNamespace",
					nil,
					appendParams(writeOptions, &namespaceNameParam),
					newResponseConfig(200,
						nilSchema, nil, writeMeta,
						"DeleteNamespaceResponse"),
				),
			},
		},
		//s.mux.HandleFunc("/v1/namespace", s.wrap(s.NamespaceCreateRequest)) routes to namespaceUpdate
		{
			Template: "/namespace",
			Operations: []*operation{
				newOperation(http.MethodPost, "namespaceUpdate", tags,
					"CreateNamespace",
					nil,
					writeOptions,
					newResponseConfig(200, nilSchema, nil, writeMeta,
						"CreateNamespaceResponse"),
				),
			},
		},
	}
}
