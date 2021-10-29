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
		//s.mux.HandleFunc("/v1/deployment/pause/{deploymentID}", s.wrap(s.DeploymentSpecificRequest))
		//s.mux.HandleFunc("/v1/deployment/promote/{deploymentID}", s.wrap(s.DeploymentSpecificRequest))
		//s.mux.HandleFunc("/v1/deployment/allocation-health/{deploymentID}", s.wrap(s.DeploymentSpecificRequest))
		//s.mux.HandleFunc("/v1/deployment/unblock/{deploymentID}", s.wrap(s.DeploymentSpecificRequest))
	}
}
