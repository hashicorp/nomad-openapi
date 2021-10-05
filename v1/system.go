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

func (s *System) GarbageCollection(ctx context.Context) (*WriteMeta, error) {
	request := s.SystemApi().PostSystemGC(s.client.Ctx)
	meta, err := s.client.ExecNoResponseWrite(ctx, request)
	if err != nil {
		return nil, err
	}

	return meta, nil
}

func (s *System) Reconcile(ctx context.Context) (*WriteMeta, error) {
	request := s.SystemApi().PostSystemReconcileSummaries(s.client.Ctx)
	meta, err := s.client.ExecNoResponseWrite(ctx, request)
	if err != nil {
		return nil, err
	}

	return meta, nil
}

