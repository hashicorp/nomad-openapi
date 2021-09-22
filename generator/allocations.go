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
	}
}
