package main

import (
	"net/http"

	"github.com/hashicorp/nomad/api"
)

func (v *v1api) getEnterprisePaths() []*apiPath {
	tags := []string{"Enterprise"}

	// TODO: add to nomad-enterprise repo
	return []*apiPath{
		//s.mux.HandleFunc("/v1/sentinel/policies", s.wrap(s.entOnly))
		//s.mux.HandleFunc("/v1/sentinel/policy/", s.wrap(s.entOnly))
		//
		//s.mux.HandleFunc("/v1/quotas", s.wrap(s.entOnly))
		{
			Template: "/quotas",
			Operations: []*operation{
				newOperation(http.MethodGet, "QuotasRequests", tags, "GetQuotas",
					nil,
					defaultQueryOpts,
					newResponseConfig(200, arraySchema, api.Quotas{}, defaultQueryMeta,
						"GetQuotasResponse"),
				),
			},
		},
		//s.mux.HandleFunc("/v1/quota-usages", s.wrap(s.entOnly))
		//s.mux.HandleFunc("/v1/quota/", s.wrap(s.entOnly))
		{
			Template: "/quota/{specName}",
			Operations: []*operation{
				newOperation(http.MethodGet, "quotaSpecQuery", tags, "GetQuotaSpec",
					nil,
					appendParams(defaultQueryOpts, &quotaSpecNameParam),
					newResponseConfig(200, objectSchema, api.QuotaSpec{}, defaultQueryMeta,
						"GetQuotaSpecResponse"),
				),
				newOperation(http.MethodPost, "quotaSpecUpdate", tags, "PostQuotaSpec",
					newRequestBody(objectSchema, api.QuotaSpec{}),
					appendParams(defaultWriteOpts, &quotaSpecNameParam),
					newResponseConfig(200, nilSchema, nil, defaultWriteMeta,
						"PostQuotaSpecResponse"),
				),
				newOperation(http.MethodDelete, "quotaSpecDelete", tags, "DeleteQuotaSpec",
					nil,
					appendParams(defaultWriteOpts, &quotaSpecNameParam),
					newResponseConfig(200, nilSchema, nil, defaultWriteMeta,
						"DeleteNamespaceResponse"),
				),
			},
		},
		//s.mux.HandleFunc("/v1/quota", s.wrap(s.entOnly))
		{
			Template: "/quota",
			Operations: []*operation{
				newOperation(http.MethodPost, "quotaSpecUpdate", tags, "CreateQuotaSpec",
					newRequestBody(objectSchema, api.QuotaSpec{}),
					defaultWriteOpts,
					newResponseConfig(200, nilSchema, nil, defaultWriteMeta,
						"CreateQuotaSpecResponse"),
				),
			},
		},
		//s.mux.HandleFunc("/v1/recommendation", s.wrap(s.entOnly))
		//s.mux.HandleFunc("/v1/recommendations", s.wrap(s.entOnly))
		//s.mux.HandleFunc("/v1/recommendations/apply", s.wrap(s.entOnly))
		//s.mux.HandleFunc("/v1/recommendation/", s.wrap(s.entOnly))
	}
}
