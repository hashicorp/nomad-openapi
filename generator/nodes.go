package main

import (
	"net/http"

	"github.com/hashicorp/nomad/api"
)

func (v *v1api) getNodePaths() []*apiPath {
	tags := []string{"Nodes"}

	return []*apiPath{
		// "/nodes"
		{
			Template: "/nodes",
			Operations: []*operation{
				newOperation(http.MethodGet, "NodesRequest", tags, "GetNodes",
					nil,
					appendParams(defaultQueryOpts, &nodeResourcesParam),
					newResponseConfig(200, arraySchema, api.NodeListStub{}, defaultQueryMeta, "GetNodesResponse"),
				),
			},
		},
		// "/node/{nodeId}"
		{
			Template: "/node/{nodeId}",
			Operations: []*operation{
				newOperation(http.MethodGet, "nodeQuery", tags, "GetNode",
					nil,
					appendParams(defaultQueryOpts, &nodeIdParam),
					newResponseConfig(200, objectSchema, api.Node{}, defaultQueryMeta, "GetNodeResponse"),
				),
			},
		},
		// "/node/{nodeId}/allocations"
		{
			Template: "/node/{nodeId}/allocations",
			Operations: []*operation{
				newOperation(http.MethodGet, "nodeAllocations", tags, "GetNodeAllocations",
					nil,
					appendParams(defaultQueryOpts, &nodeIdParam),
					newResponseConfig(200, arraySchema, api.AllocationListStub{}, defaultQueryMeta, "GetNodeAllocationsResponse"),
				),
			},
		},
		// "/node/{nodeId}/evaluate"
		// {
		// 	Template: "/node/{nodeId}/evaluate",
		// 	Operations: []*operation{
		// 		newOperation(http.MethodPost, "nodeEvaluate", tags, "PostNodeEvaluation",
		// 			nil,
		// 			appendParams(defaultQueryOpts, &nodeIdParam),
		// 			newResponseConfig(200, objectSchema, nil, defaultQueryMeta, "PostNodeEvaluationResponse"),
		// 		),
		// 	},
		// },
		// "/node/{nodeId}/drain"
		{
			Template: "/node/{nodeId}/drain",
			Operations: []*operation{
				newOperation(http.MethodPost, "nodeToggleDrain", tags, "UpdateNodeDrain",
					newRequestBody(objectSchema, api.NodeUpdateDrainRequest{}),
					appendParams(defaultQueryOpts, &nodeIdParam),
					newResponseConfig(200, objectSchema, api.NodeDrainUpdateResponse{}, defaultWriteMeta, "NodeDrainUpdateResponse"),
				),
			},
		},
		// "/node/{nodeId}/purge"
		{
			Template: "/node/{nodeId}/purge",
			Operations: []*operation{
				newOperation(http.MethodPost, "nodePurge", tags, "UpdateNodePurge",
					nil,
					appendParams(defaultQueryOpts, &nodeIdParam),
					newResponseConfig(200, objectSchema, api.NodePurgeResponse{}, defaultWriteMeta, "NodePurgeResponse"),
				),
			},
		},
		// "/node/{nodeId}/eligibility"
		{
			Template: "/node/{nodeId}/eligibility",
			Operations: []*operation{
				newOperation(http.MethodPost, "nodeToggleEligibility", tags, "UpdateNodeEligibility",
					newRequestBody(objectSchema, api.NodeUpdateEligibilityRequest{}),
					appendParams(defaultQueryOpts, &nodeIdParam),
					newResponseConfig(200, objectSchema, api.NodeEligibilityUpdateResponse{}, defaultWriteMeta, "NodeEligibilityUpdateResponse"),
				),
			},
		},
	}
}
