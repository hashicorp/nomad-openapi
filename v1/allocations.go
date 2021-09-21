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
