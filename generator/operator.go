package main

import (
	"net/http"

	"github.com/hashicorp/nomad/api"
)

var AutopilotSetConfiguration bool

func (v *v1api) getOperatorPaths() []*apiPath {
	tags := []string{"Operator"}

	return []*apiPath{
		//s.mux.HandleFunc("/v1/operator/license", s.wrap(s.LicenseRequest))
		// NOTE: there's no license endpoint
		// QUESTION: should this be under the enterprise-api instead...?
		//s.mux.HandleFunc("/v1/operator/raft/", s.wrap(s.OperatorRequest))
		{
			Template: "/operator/raft/",
			Operations: []*operation{
				//configuration
				newOperation(http.MethodGet,
					httpServer.OperatorRequest,
					tags,
					"GetOperatorRaft",
					nil,
					defaultQueryOpts,
					newResponseConfig(200,
						arraySchema,
						api.RaftServer{},
						nil,
						"GetOperatorRaftResponse",
					),
				),
				//peer
				newOperation(http.MethodDelete,
					httpServer.OperatorRequest,
					tags,
					"DeleteOperatorRaft",
					nil,
					defaultWriteOpts,
					newResponseConfig(200,
						nilSchema,
						nil,
						nil,
						"DeleteOperatorRaftResponse",
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
						AutopilotSetConfiguration, // ??
						nil,
						"PutOperatorAutopilotConfigurationResponse",
					),
				),
			},
		},
		//s.mux.HandleFunc("/v1/operator/autopilot/health", s.wrap(s.OperatorServerHealth))
		{
			Template: "/operator/autopilot/health",
			Operations: []*operation{
				newOperation(http.MethodGet,
					httpServer.OperatorServerHealth,
					tags,
					"GetOperatorAutopilotHealth",
					nil,
					defaultQueryOpts,
					newResponseConfig(200,
						objectSchema,
						api.OperatorHealthReply{},
						nil,
						"GetOperatorAutopilotHealthResponse",
					),
				),
			},
		},
		//s.mux.HandleFunc("/v1/operator/snapshot", s.wrap(s.SnapshotRequest))
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
		{
			Template: "/operator/scheduler/configuration",
			Operations: []*operation{
				newOperation(http.MethodGet,
					httpServer.OperatorSchedulerConfiguration,
					tags,
					"GetOperatorSchedulerConfiguration",
					nil,
					defaultQueryOpts,
					newResponseConfig(200,
						objectSchema,
						api.SchedulerConfigurationResponse{},
						defaultQueryMeta,
						"GetOperatorSchedulerConfigurationResponse",
					),
				),
				newOperation(http.MethodPost,
					httpServer.OperatorSchedulerConfiguration,
					tags,
					"PostOperatorSchedulerConfiguration",
					newRequestBody(objectSchema, api.SchedulerConfiguration{}),
					defaultWriteOpts,
					newResponseConfig(200,
						objectSchema,
						api.SchedulerSetConfigurationResponse{},
						defaultWriteMeta,
						"PostOperatorSchedulerConfigurationResponse",
					),
				),
			},
		},
	}
}
