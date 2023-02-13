// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"net/http"

	"github.com/hashicorp/nomad/api"
	"github.com/hashicorp/nomad/nomad/structs"
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
		//s.mux.HandleFunc("/v1/allocation/{allocID}", s.wrap(s.AllocSpecificRequest))
		{
			Template: "/allocation/{allocID}",
			Operations: []*operation{
				newOperation(http.MethodGet,
					httpServer.AllocSpecificRequest,
					tags,
					"GetAllocation",
					nil,
					appendParams(defaultQueryOpts, &allocIDParam),
					newResponseConfig(200,
						objectSchema,
						api.Allocation{},
						defaultQueryMeta,
						"GetAllocationResponse",
					),
				),
			},
		},
		//s.mux.HandleFunc("/v1/allocation/{allocID}/stop", s.wrap(s.AllocSpecificRequest))
		{
			Template: "/allocation/{allocID}/stop",
			Operations: []*operation{
				newOperation(http.MethodPost,
					httpServer.AllocSpecificRequest,
					tags,
					"PostAllocationStop",
					nil,
					appendParams(defaultQueryOpts, &allocIDParam, &allocNoShutdownDelayParam),
					newResponseConfig(200,
						objectSchema,
						structs.AllocStopResponse{},
						defaultWriteMeta,
						"PostAllocationStopResponse",
					),
				),
			},
		},
		//s.mux.HandleFunc("/v1/allocation/{allocID}/services", s.wrap(s.AllocSpecificRequest))
		{
			Template: "/allocation/{allocID}/services",
			Operations: []*operation{
				newOperation(http.MethodGet,
					httpServer.AllocSpecificRequest,
					tags,
					"GetAllocationServices",
					nil,
					appendParams(defaultQueryOpts, &allocIDParam),
					newResponseConfig(200,
						arraySchema,
						structs.ServiceRegistration{},
						defaultQueryMeta,
						"GetAllocationServicesResponse",
					),
				),
			},
		},
	}
}
