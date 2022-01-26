package main

import (
	"net/http"

	"github.com/hashicorp/nomad/api"
	"github.com/hashicorp/nomad/nomad/structs"
)

func (v *v1api) getOperatorPaths() []*apiPath {
	tags := []string{"Operator"}

	return []*apiPath{
		//s.mux.HandleFunc("/v1/operator/license", s.wrap(s.LicenseRequest))
		// NOTE: there's no license endpoint
		// QUESTION: should this be under the enterprise-api instead...?
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
						structs.RaftConfigurationResponse{},
						// did not fix :(
						//api.RaftServer{},
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
						//api.AutopilotConfiguration{},
						AutopilotConfiguration{},
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
						//boolSchema,	// with boolSchema,
						objectSchema,
						true, // any bool here will work?
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

// NOTE: temp work around to override the LastContactThreshold and ServerStabilizationTime types of time.Duration
type AutopilotConfiguration struct {
	CleanupDeadServers bool

	// LastContactThreshold is the limit on the amount of time a server can go
	// without leader contact before being considered unhealthy.
	LastContactThreshold string

	MaxTrailingLogs uint64
	MinQuorum       uint

	// ServerStabilizationTime is the minimum amount of time a server must be
	// in a stable, healthy state before it can be added to the cluster. Only
	// applicable with Raft protocol version 3 or higher.
	ServerStabilizationTime string

	EnableRedundancyZones   bool
	DisableUpgradeMigration bool
	EnableCustomUpgrades    bool
	CreateIndex             uint64
	ModifyIndex             uint64
}
