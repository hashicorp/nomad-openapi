package main

import (
	"net/http"

	"github.com/hashicorp/nomad/api"
)

func (v *v1api) getSearchPaths() []*apiPath {
	tags := []string{"Search"}

	return []*apiPath{
		//s.mux.HandleFunc("/v1/search/fuzzy", s.wrap(s.FuzzySearchRequest))
		{
			Template: "/search/fuzzy",
			Operations: []*operation{
				newOperation(http.MethodPost, "newFuzzySearchRequest", tags, "GetFuzzySearch",
					newRequestBody(objectSchema, api.FuzzySearchRequest{}),
					defaultQueryOpts,
					newResponseConfig(200, objectSchema, api.FuzzySearchResponse{}, defaultQueryMeta, "GetFuzzySearchResponse"),
				),
			},
		},
		//s.mux.HandleFunc("/v1/search", s.wrap(s.SearchRequest))
		{
			Template: "/search",
			Operations: []*operation{
				newOperation(http.MethodPost, "newSearchRequest", tags, "GetSearch",
					newRequestBody(objectSchema, api.SearchRequest{}),
					defaultQueryOpts,
					newResponseConfig(200, objectSchema, api.SearchResponse{}, defaultQueryMeta, "GetSearchResponse"),
				),
			},
		},
	}
}
