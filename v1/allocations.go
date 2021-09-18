package v1

import (
	"context"

	"github.com/hashicorp/nomad-openapi/v1/client"
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
	request = a.client.setQueryOptions(ctx, request).(client.ApiGetAllocationsRequest)

	result, response, err := request.Execute()
	if err != nil {
		return nil, nil, err
	}

	meta, err := parseQueryMeta(response)
	if err != nil {
		return nil, nil, err
	}

	return &result, meta, nil
}
