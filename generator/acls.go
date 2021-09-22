package main

import (
	"net/http"

	"github.com/hashicorp/nomad/api"
)

func (v *v1api) getACLPaths() []*apiPath {
	tags := []string{"ACL"}

	return []*apiPath{
		//s.mux.HandleFunc("/v1/acl/policies", s.wrap(s.ACLPoliciesRequest))
		{
			Template: "/acl/policies",
			Operations: []*operation{
				newOperation(http.MethodGet,
					httpServer.ACLPoliciesRequest,
					tags,
					"GetACLPolicies",
					nil,
					defaultQueryOpts,
					newResponseConfig(200,
						arraySchema,
						api.ACLPolicyListStub{},
						defaultQueryMeta,
						"GetACLPoliciesResponse",
					),
				),
			},
		},
		//s.mux.HandleFunc("/v1/acl/policy/", s.wrap(s.ACLPolicySpecificRequest))
		//s.mux.HandleFunc("/v1/acl/token/onetime", s.wrap(s.UpsertOneTimeToken))
		//s.mux.HandleFunc("/v1/acl/token/onetime/exchange", s.wrap(s.ExchangeOneTimeToken))
		//s.mux.HandleFunc("/v1/acl/bootstrap", s.wrap(s.ACLTokenBootstrap))
		//s.mux.HandleFunc("/v1/acl/tokens", s.wrap(s.ACLTokensRequest))
		//s.mux.HandleFunc("/v1/acl/token", s.wrap(s.ACLTokenSpecificRequest))
		//s.mux.HandleFunc("/v1/acl/token/", s.wrap(s.ACLTokenSpecificRequest))
	}
}
