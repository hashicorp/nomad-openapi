// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v1

import (
	"context"

	client "github.com/flytocolors/nomad-openapi/clients/go/v1"
)

type Deployments struct {
	client *Client
}

func (c *Client) Deployments() *Deployments {
	return &Deployments{client: c}
}

func (d *Deployments) DeploymentsApi() *client.DeploymentsApiService {
	return d.client.apiClient.DeploymentsApi
}

func (d *Deployments) Deployments(ctx context.Context) (*[]client.Deployment, *QueryMeta, OpenAPIError) {
	request := d.DeploymentsApi().GetDeployments(d.client.Ctx)
	result, meta, err := d.client.ExecQuery(ctx, request)
	if err != nil {
		return nil, nil, err
	}

	final := result.([]client.Deployment)
	return &final, meta, nil
}

func (d *Deployments) GetDeployment(ctx context.Context, deploymentID string) (*client.Deployment, *QueryMeta, OpenAPIError) {
	if deploymentID == "" {
		return nil, nil, &APIError{error: "deployment id is required"}
	}

	request := d.DeploymentsApi().GetDeployment(d.client.Ctx, deploymentID)
	result, meta, err := d.client.ExecQuery(ctx, request)
	if err != nil {
		return nil, nil, err
	}

	final := result.(client.Deployment)
	return &final, meta, nil
}

func (d *Deployments) Allocations(ctx context.Context, deploymentID string) (*[]client.AllocationListStub, *QueryMeta, OpenAPIError) {
	if deploymentID == "" {
		return nil, nil, &APIError{error: "deployment id is required"}
	}

	request := d.DeploymentsApi().GetDeploymentAllocations(d.client.Ctx, deploymentID)
	result, meta, err := d.client.ExecQuery(ctx, request)
	if err != nil {
		return nil, nil, err
	}

	final := result.([]client.AllocationListStub)
	return &final, meta, nil
}

func (d *Deployments) Fail(ctx context.Context, deploymentID string) (*client.DeploymentUpdateResponse, OpenAPIError) {
	if deploymentID == "" {
		return nil, &APIError{error: "deployment id is required"}
	}

	request := d.DeploymentsApi().PostDeploymentFail(d.client.Ctx, deploymentID)
	result, err := d.client.ExecNoMetaWrite(ctx, request)
	if err != nil {
		return nil, &APIError{error: err.Error()}
	}

	final := result.(client.DeploymentUpdateResponse)
	return &final, nil
}

func (d *Deployments) Pause(ctx context.Context, deploymentID string, pause bool) (*client.DeploymentUpdateResponse, OpenAPIError) {
	if deploymentID == "" {
		return nil, &APIError{error: "deployment id is required"}
	}

	request := d.DeploymentsApi().PostDeploymentPause(d.client.Ctx, deploymentID)

	pauseRequest := client.NewDeploymentPauseRequest()
	pauseRequest.SetDeploymentID(deploymentID)
	pauseRequest.SetPause(pause)

	request = request.DeploymentPauseRequest(*pauseRequest)

	result, err := d.client.ExecNoMetaWrite(ctx, request)
	if err != nil {
		return nil, &APIError{error: err.Error()}
	}

	final := result.(client.DeploymentUpdateResponse)
	return &final, nil
}

func (d *Deployments) Promote(ctx context.Context, deploymentID string, all bool, groups []string) (*client.DeploymentUpdateResponse, OpenAPIError) {
	if deploymentID == "" {
		return nil, &APIError{error: "deployment id is required"}
	}

	request := d.DeploymentsApi().PostDeploymentPromote(d.client.Ctx, deploymentID)

	promoteRequest := client.NewDeploymentPromoteRequest()
	promoteRequest.SetDeploymentID(deploymentID)
	promoteRequest.SetAll(all)
	promoteRequest.SetGroups(groups)

	request = request.DeploymentPromoteRequest(*promoteRequest)

	result, err := d.client.ExecNoMetaWrite(ctx, request)
	if err != nil {
		return nil, &APIError{error: err.Error()}
	}

	final := result.(client.DeploymentUpdateResponse)
	return &final, nil
}

func (d *Deployments) AllocationHealth(ctx context.Context, deploymentID string, healthyallocs []string, unhealthyallocs []string) (*client.DeploymentUpdateResponse, OpenAPIError) {
	if deploymentID == "" {
		return nil, &APIError{error: "deployment id is required"}
	}

	request := d.DeploymentsApi().PostDeploymentAllocationHealth(d.client.Ctx, deploymentID)

	allocHealthRequest := client.NewDeploymentAllocHealthRequest()
	allocHealthRequest.SetDeploymentID(deploymentID)
	allocHealthRequest.SetHealthyAllocationIDs(healthyallocs)
	allocHealthRequest.SetUnhealthyAllocationIDs(unhealthyallocs)

	request = request.DeploymentAllocHealthRequest(*allocHealthRequest)

	result, err := d.client.ExecNoMetaWrite(ctx, request)
	if err != nil {
		return nil, &APIError{error: err.Error()}
	}

	final := result.(client.DeploymentUpdateResponse)
	return &final, nil
}

func (d *Deployments) Unblock(ctx context.Context, deploymentID string) (*client.DeploymentUpdateResponse, OpenAPIError) {
	if deploymentID == "" {
		return nil, &APIError{error: "deployment id is required"}
	}

	request := d.DeploymentsApi().PostDeploymentUnblock(d.client.Ctx, deploymentID)

	unblockRequest := client.NewDeploymentUnblockRequest()
	unblockRequest.SetDeploymentID(deploymentID)

	request = request.DeploymentUnblockRequest(*unblockRequest)

	result, err := d.client.ExecNoMetaWrite(ctx, request)
	if err != nil {
		return nil, &APIError{error: err.Error()}
	}

	final := result.(client.DeploymentUpdateResponse)
	return &final, nil
}
