// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v1

import (
	"context"

	client "github.com/hashicorp/nomad-openapi/clients/go/v1"
)

// Jobs encapsulates and extends the generated JobsApiService with convenience methods.
type Allocations struct {
	client *Client
}

func (c *Client) Allocations() *Allocations {
	return &Allocations{client: c}
}

func (a *Allocations) AllocationsApi() *client.AllocationsApiService {
	return a.client.apiClient.AllocationsApi
}

func (a *Allocations) GetAllocations(ctx context.Context) (*[]client.AllocationListStub, *QueryMeta, OpenAPIError) {
	request := a.AllocationsApi().GetAllocations(a.client.Ctx)

	result, meta, err := a.client.ExecQuery(ctx, request)
	if err != nil {
		return nil, nil, err
	}

	final := result.([]client.AllocationListStub)
	return &final, meta, nil
}

func (a *Allocations) GetAllocation(ctx context.Context, allocID string) (*client.Allocation, *QueryMeta, OpenAPIError) {
	request := a.AllocationsApi().GetAllocation(a.client.Ctx, allocID)

	result, meta, err := a.client.ExecQuery(ctx, request)
	if err != nil {
		return nil, nil, err
	}

	final := result.(client.Allocation)
	return &final, meta, nil
}

func (a *Allocations) StopAllocation(ctx context.Context, allocID string, noShutdownDelay bool) (*client.AllocStopResponse, *WriteMeta, OpenAPIError) {
	request := a.AllocationsApi().PostAllocationStop(a.client.Ctx, allocID)

	request = request.NoShutdownDelay(noShutdownDelay)
	result, meta, err := a.client.ExecWrite(ctx, request)
	if err != nil {
		return nil, nil, err
	}

	final := result.(client.AllocStopResponse)
	return &final, meta, nil
}


func (a *Allocations) GetAllocationServices(ctx context.Context, allocID string) (*[]client.ServiceRegistration, *QueryMeta, OpenAPIError) {
	request := a.AllocationsApi().GetAllocationServices(a.client.Ctx, allocID)

	result, meta, err := a.client.ExecQuery(ctx, request)
	if err != nil {
		return nil, nil, err
	}

	final := result.([]client.ServiceRegistration)
	return &final, meta, nil
}
