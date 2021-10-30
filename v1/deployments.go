package v1

import (
	"context"
	"errors"

	client "github.com/hashicorp/nomad-openapi/clients/go/v1"
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

func (d *Deployments) Deployments(ctx context.Context) (*[]client.Deployment, *QueryMeta, error) {
	request := d.DeploymentsApi().GetDeployments(d.client.Ctx)
	result, meta, err := d.client.ExecQuery(ctx, request)
	if err != nil {
		return nil, nil, err
	}

	final := result.([]client.Deployment)
	return &final, meta, nil
}

func (d *Deployments) GetDeployment(ctx context.Context, deploymentID string) (*client.Deployment, *QueryMeta, error) {
	if deploymentID == "" {
		return nil, nil, errors.New("deployment id is required")
	}

	request := d.DeploymentsApi().GetDeployment(d.client.Ctx, deploymentID)
	result, meta, err := d.client.ExecQuery(ctx, request)
	if err != nil {
		return nil, nil, err
	}

	final := result.(client.Deployment)
	return &final, meta, nil
}

func (d *Deployments) Allocations(ctx context.Context, deploymentID string) (*[]client.AllocationListStub, *QueryMeta, error) {
	if deploymentID == "" {
		return nil, nil, errors.New("deployment id is required")
	}

	request := d.DeploymentsApi().GetDeploymentAllocations(d.client.Ctx, deploymentID)
	result, meta, err := d.client.ExecQuery(ctx, request)
	if err != nil {
		return nil, nil, err
	}

	final := result.([]client.AllocationListStub)
	return &final, meta, nil
}

func (d *Deployments) Fail(ctx context.Context, deploymentID string) (*client.DeploymentUpdateResponse, error) {
	if deploymentID == "" {
		return nil, errors.New("deployment id is required")
	}

	request := d.DeploymentsApi().PostDeploymentFail(d.client.Ctx, deploymentID)
	result, err := d.client.ExecRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	final := result.(client.DeploymentUpdateResponse)
	return &final, nil
}

/* func (d *Deployments) Pause(ctx context.Context, deploymentID string) (*client.DeploymentUpdateResponse, error) {
	if deploymentID == "" {
		return nil, errors.New("deployment id is required")
	}

	request := d.DeploymentsApi().PostDeploymentPause(d.client.Ctx, deploymentID)

	deploymentsRequest := client.NewDeployment()
	deploymentsRequest.SetID(deploymentID)

	result, err := d.client.ExecRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	final := result.(client.DeploymentUpdateResponse)
	return &final, nil
} */

/* func (d *Deployments) Unblock(ctx context.Context, deploymentID string) (*client.DeploymentUpdateResponse, error) {
	if deploymentID == "" {
		return nil, errors.New("deployment id is required")
	}

	request := d.DeploymentsApi().PostDeploymentUnblock(d.client.Ctx, deploymentID)

	deploymentsRequest := client.NewDeployment()
	deploymentsRequest.SetID(deploymentID)

	result, err := d.client.ExecRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	final := result.(client.DeploymentUpdateResponse)
	return &final, nil
} */
