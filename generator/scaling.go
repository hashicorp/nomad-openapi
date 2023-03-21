package main

import (
	"net/http"

	"github.com/hashicorp/nomad/api"
)

func (v *v1api) getScalingPaths() []*apiPath {
	tags := []string{"Scaling"}

	return []*apiPath{
		//s.mux.HandleFunc("/v1/scaling/policies", s.wrap(s.ScalingPoliciesRequest))
		{
			Template: "/scaling/policies",
			Operations: []*operation{
				newOperation(http.MethodGet,
					httpServer.ScalingPoliciesRequest,
					tags,
					"GetScalingPolicies",
					nil,
					defaultQueryOpts,
					newResponseConfig(200,
						arraySchema,
						api.ScalingPolicyListStub{},
						defaultQueryMeta,
						"GetScalingPoliciesResponse",
					),
				),
			},
		},
		//s.mux.HandleFunc("/v1/scaling/policy/", s.wrap(s.ScalingPolicySpecificRequest))
		{
			Template: "/scaling/policy/{policyID}",
			Operations: []*operation{
				newOperation(http.MethodGet,
					httpServer.ScalingPolicySpecificRequest,
					tags,
					"GetScalingPolicy",
					nil,
					appendParams(defaultQueryOpts, &scalingPolicyIDParam),
					newResponseConfig(200,
						objectSchema,
						api.ScalingPolicy{},
						defaultQueryMeta,
						"GetScalingPolicyResponse",
					),
				),
			},
		},
	}
}
