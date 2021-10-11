package main

import (
	"net/http"

	"github.com/hashicorp/nomad/api"
)

func (v *v1api) getAllocationPaths() []*apiPath {
	tags := []string{"Allocations"}

	return []*apiPath{
		//s.mux.HandleFunc("/v1/allocations", s.wrap(s.AllocsRequest))
		{
			Template: "/allocations",
			Operations: []*operation{
				newOperation(http.MethodGet,
					httpServer.AllocsRequest,
					tags,
					"GetAllocations",
					nil,
					appendParams(defaultQueryOpts, &allocationsResourcesParam, &allocationsTasksStatesParam),
					// TODO: Think of ways to enhance the standard responses to handle the vagaries of different error codes
					newResponseConfig(200,
						arraySchema,
						api.AllocationListStub{},
						defaultQueryMeta,
						"GetAllocationsResponse",
					),
				),
			},
		},
		//s.mux.HandleFunc("/v1/allocation/", s.wrap(s.AllocSpecificRequest))
		{
			Template: "/allocation/{allocationID}",
			Operations: []*operation{
				newOperation(http.MethodGet,
					httpServer.AllocSpecificRequest,
					tags,
					"GetAllocation",
					nil,
					appendParams(defaultQueryOpts, &allocationIDParam),
					newResponseConfig(200,
						objectSchema,
						api.Allocation{},
						defaultQueryMeta,
						"GetAllocationResponse",
					),
				),
			},
		},
		//s.mux.HandleFunc("/v1/allocation/", s.wrap(s.AllocSpecificRequest))
		{
			Template: "/allocation/{allocationID}/stop",
			Operations: []*operation{
				newOperation(http.MethodPost,
					httpServer.AllocSpecificRequest,
					tags,
					"StopAllocation",
					nil,
					appendParams(defaultWriteOpts, &allocationIDParam),
					newResponseConfig(200,
						objectSchema,
						api.AllocStopResponse{},
						defaultWriteMeta,
						"StopAllocationResponse",
					),
				),
			},
		},
	}
}
