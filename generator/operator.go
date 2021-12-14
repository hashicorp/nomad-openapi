package main

import (
	"net/http"

	"github.com/hashicorp/nomad/api"
)

func (v *v1api) getOperatorPaths() []*apiPath {
	// tags := []string{"Operator"}

	return []*apiPath{
		//s.mux.HandleFunc("/v1/operator/license", s.wrap(s.LicenseRequest))
		//s.mux.HandleFunc("/v1/operator/raft/", s.wrap(s.OperatorRequest))
		{
			Template: "/operator/raft/configuration",
			Operations: []*operation{
				newOperation(http.MethodGet,
					httpServer.OperatorRequest,
					tags,
					"GetOperatorRaftConfiguration",
					nil,
					defaultQueryOpts,
					newResponseConfig(200,
						arraySchema,
						api.RaftServer{},
						nil,
						"GetOperatorRaftConfigurationResponse",
					),
				),
			},
		},
		{
			Template: "/operator/raft/peer",
			Operations: []*operation{
				newOperation(http.MethodDelete,
					httpServer.OperatorRequest,
					tags,
					"DeleteOperatorRaftPeer",
					nil,
					defaultWriteOpts,
					newResponseConfig(200,
						nilSchema,
						nil,
						nil,
						"DeleteOperatorRaftPeerResponse",
					),
				),
			},
		},
		//s.mux.HandleFunc("/v1/operator/autopilot/configuration", s.wrap(s.OperatorAutopilotConfiguration))
		{
			Template: "/operator/autopilot/configuration",
			Operations: []*operation{
				newOperation(http.MethodGet,
					httpServer.OperatorAutopilotConfiguration,
					tags,
					"GetOperatorAutopilotConfiguration",
					nil,
					defaultQueryOpts,
					newResponseConfig(200,
						objectSchema,
						api.AutopilotConfiguration{},
						nil,
						"GetOperatorAutopilotConfigurationResponse",
					),
				),
				// Does not accept POST -- only PUT
				newOperation(http.MethodPut,
					httpServer.OperatorAutopilotConfiguration,
					tags,
					"PutOperatorAutopilotConfiguration",
					newRequestBody(objectSchema, api.AutopilotConfiguration{}),
					defaultWriteOpts,
					newResponseConfig(200,
						boolSchema,
						true, // any bool here will work?
						nil,
						"PutOperatorAutopilotConfigurationResponse",
					),
				),
			},
		},
		//s.mux.HandleFunc("/v1/operator/autopilot/health", s.wrap(s.OperatorServerHealth))
		//s.mux.HandleFunc("/v1/operator/snapshot", s.wrap(s.SnapshotRequest))
		// NOTE: I don't know how to handle the snapshots :(
		/* {
		    Template: "/operator/snapshot",
		    Operations: []*operation{
		        newOperation(http.MethodGet,
		            httpServer.SnapshotRequest,
		            tags,
		            "GetOperatorSnapshot",
		            nil,
		            defaultQueryOpts,
		            newResponseConfig(200,
		                nilSchema,
		                nil,
		                nil,
		                "GetOperatorSnapshotResponse",
		            ),
		        ),
		        newOperation(http.MethodPost,
		            httpServer.SnapshotRequest,
		            tags,
		            "PostOperatorSnapshot",
		            // ???
		            newRequestBody(objectSchema, api.TODORequest{}),
		            defaultWriteOpts,
		            newResponseConfig(200,
		                nilSchema,
		                nil,
		                nil,
		                "PostOperatorSnapshotResponse",
		            ),
		        ),
		    },
		}, */
		//s.mux.HandleFunc("/v1/operator/scheduler/configuration", s.wrap(s.OperatorSchedulerConfiguration))
	}
}
