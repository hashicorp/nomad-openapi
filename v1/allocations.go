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

func (a *Allocations) GetAllocations(ctx context.Context) (*[]client.AllocationListStub, *QueryMeta, error) {
	request := a.AllocationsApi().GetAllocations(a.client.Ctx)

	result, meta, err := a.client.ExecQuery(ctx, request)
	if err != nil {
		return nil, nil, err
	}

	final := result.([]client.AllocationListStub)
	return &final, meta, nil
}

func (a *Allocations) GetAllocation(ctx context.Context, allocationID string) (*client.Allocation, *QueryMeta, error) {
	request := a.AllocationsApi().GetAllocation(a.client.Ctx, allocationID)
	result, meta, err := a.client.ExecQuery(ctx, request)
	if err != nil {
		return nil, nil, err
	}

	final := result.(client.Allocation)
	return &final, meta, nil
}

func (a *Allocations) Stop(ctx context.Context, allocationID string) (*client.AllocStopResponse, *WriteMeta, error) {
	request := a.AllocationsApi().StopAllocation(a.client.Ctx, allocationID)
	result, meta, err := a.client.ExecWrite(ctx, request)
	if err != nil {
		return nil, nil, err
	}

	final := result.(client.AllocStopResponse)
	return &final, meta, nil
}
