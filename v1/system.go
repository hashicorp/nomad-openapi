package v1

import (
	"context"

	client "github.com/hashicorp/nomad-openapi/clients/go/v1"
)

type System struct {
	client *Client
}

func (c *Client) System() *System {
	return &System{client: c}
}

func (s *System) SystemApi() *client.SystemApiService {
	return s.client.apiClient.SystemApi
}

func (s *System) GarbageCollection(ctx context.Context) (interface{}, error) {
	request := s.SystemApi().PutSystemGC(s.client.Ctx)

	result, err := s.client.ExecNoResponseRequest(ctx, request)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (s *System) Reconcile(ctx context.Context) (interface{}, error) {
	request := s.SystemApi().PutSystemReconcileSummaries(s.client.Ctx)

	result, err := s.client.ExecNoResponseRequest(ctx, request)
	if err != nil {
		return result, err
	}

	return result, nil
}
