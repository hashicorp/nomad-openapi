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
					defaultQueryOpts,
					newResponseConfig(200, arraySchema, api.Namespace{}, defaultQueryMeta,
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
					appendParams(defaultQueryOpts, &namespaceNameParam),
					newResponseConfig(200, objectSchema, api.Namespace{}, defaultQueryMeta,
						"GetNamespaceResponse"),
				),
				newOperation(http.MethodPost, "namespaceUpdate", tags,
					"PostNamespace",
					newRequestBody(objectSchema, api.Namespace{}),
					appendParams(defaultWriteOpts, &namespaceNameParam),
					newResponseConfig(200,
						nilSchema, nil, defaultWriteMeta,
						"PostNamespaceResponse"),
				),
				newOperation(http.MethodDelete, "namespaceDelete", tags,
					"DeleteNamespace",
					nil,
					appendParams(defaultWriteOpts, &namespaceNameParam),
					newResponseConfig(200,
						nilSchema, nil, defaultWriteMeta,
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
					defaultWriteOpts,
					newResponseConfig(200, nilSchema, nil, defaultWriteMeta,
						"CreateNamespaceResponse"),
				),
			},
		},
	}
}
