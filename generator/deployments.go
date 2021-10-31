package main

import (
	"net/http"

	"github.com/hashicorp/nomad/api"
)

func (v *v1api) getDeploymentsPaths() []*apiPath {
	tags := []string{"Deployments"}

	return []*apiPath{
		//s.mux.HandleFunc("/v1/deployments", s.wrap(s.DeploymentsRequest))
		{
			Template: "/deployments",
			Operations: []*operation{
				newOperation(http.MethodGet,
					httpServer.DeploymentsRequest,
					tags,
					"GetDeployments",
					nil,
					defaultQueryOpts,
					newResponseConfig(200,
						arraySchema,
						api.Deployment{},
						defaultQueryMeta,
						"GetDeploymentsResponse",
					),
				),
			},
		},
		//s.mux.HandleFunc("/v1/deployment/", s.wrap(s.DeploymentSpecificRequest))
		{
			Template: "/deployment/{deploymentID}",
			Operations: []*operation{
				newOperation(http.MethodGet,
					httpServer.DeploymentSpecificRequest,
					tags,
					"GetDeployment",
					nil,
					appendParams(defaultQueryOpts, &deploymentIDParam),
					newResponseConfig(200,
						objectSchema,
						api.Deployment{},
						defaultQueryMeta,
						"GetDeploymentResponse",
					),
				),
			},
		},
		//s.mux.HandleFunc("/v1/deployment/allocations/{deploymentID}", s.wrap(s.DeploymentSpecificRequest))
		{
			Template: "/deployment/allocations/{deploymentID}",
			Operations: []*operation{
				newOperation(http.MethodGet,
					httpServer.DeploymentSpecificRequest,
					tags,
					"GetDeploymentAllocations",
					nil,
					appendParams(defaultQueryOpts, &deploymentIDParam),
					newResponseConfig(200,
						arraySchema,
						api.AllocationListStub{},
						defaultQueryMeta,
						"GetDeploymentAllocationsResponse",
					),
				),
			},
		},
		//s.mux.HandleFunc("/v1/deployment/fail/{deploymentID}", s.wrap(s.DeploymentSpecificRequest))
		{
			Template: "/deployment/fail/{deploymentID}",
			Operations: []*operation{
				newOperation(http.MethodPost,
					httpServer.DeploymentSpecificRequest,
					tags,
					"PostDeploymentFail",
					nil,
					appendParams(defaultWriteOpts, &deploymentIDParam),
					newResponseConfig(200,
						objectSchema,
						api.DeploymentUpdateResponse{},
						nil,
						"PostDeploymentFailResponse",
					),
				),
			},
		},
		//s.mux.HandleFunc("/v1/deployment/pause/{deploymentID}", s.wrap(s.DeploymentSpecificRequest))
		{
			Template: "/deployment/pause/{deploymentID}",
			Operations: []*operation{
				newOperation(http.MethodPost,
					httpServer.DeploymentSpecificRequest,
					tags,
					"PostDeploymentPause",
					newRequestBody(objectSchema, api.DeploymentPauseRequest{}),
					appendParams(defaultWriteOpts, &deploymentIDParam),
					newResponseConfig(200,
						objectSchema,
						api.DeploymentUpdateResponse{},
						nil,
						"PostDeploymentPauseResponse",
					),
				),
			},
		},
		//s.mux.HandleFunc("/v1/deployment/promote/{deploymentID}", s.wrap(s.DeploymentSpecificRequest))
		{
			Template: "/deployment/promote/{deploymentID}",
			Operations: []*operation{
				newOperation(http.MethodPost,
					httpServer.DeploymentSpecificRequest,
					tags,
					"PostDeploymentPromote",
					newRequestBody(objectSchema, api.DeploymentPromoteRequest{}),
					appendParams(defaultWriteOpts, &deploymentIDParam),
					newResponseConfig(200,
						objectSchema,
						api.DeploymentUpdateResponse{},
						nil,
						"PostDeploymentPromoteResponse",
					),
				),
			},
		},
		//s.mux.HandleFunc("/v1/deployment/allocation-health/{deploymentID}", s.wrap(s.DeploymentSpecificRequest))
		{
			Template: "/deployment/allocation-health/{deploymentID}",
			Operations: []*operation{
				newOperation(http.MethodPost,
					httpServer.DeploymentSpecificRequest,
					tags,
					"PostDeploymentAllocationHealth",
					newRequestBody(objectSchema, api.DeploymentAllocHealthRequest{}),
					appendParams(defaultWriteOpts, &deploymentIDParam),
					newResponseConfig(200,
						objectSchema,
						api.DeploymentUpdateResponse{},
						nil,
						"PostDeploymentAllocationHealthResponse",
					),
				),
			},
		},
		//s.mux.HandleFunc("/v1/deployment/unblock/{deploymentID}", s.wrap(s.DeploymentSpecificRequest))
		// NOTE: API documention is wrong!!  There should be a request body/payload
		// https://www.nomadproject.io/api-docs/deployments#unblock-deployment
		{
			Template: "/deployment/unblock/{deploymentID}",
			Operations: []*operation{
				newOperation(http.MethodPost,
					httpServer.DeploymentSpecificRequest,
					tags,
					"PostDeploymentUnblock",
					newRequestBody(objectSchema, api.DeploymentUnblockRequest{}),
					appendParams(defaultWriteOpts, &deploymentIDParam),
					newResponseConfig(200,
						objectSchema,
						api.DeploymentUpdateResponse{},
						nil,
						"PostDeploymentUnblockResponse",
					),
				),
			},
		},
	}
}
