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

func (s *System) GarbageCollect(ctx context.Context) error {
	request := s.SystemApi().PutSystemGC(s.client.Ctx)

	err := s.client.ExecNoResponseRequest(ctx, request)
	if err != nil {
		return err
	}

	return nil
}

func (s *System) Reconcile(ctx context.Context) error {
	request := s.SystemApi().PutSystemReconcileSummaries(s.client.Ctx)

	err := s.client.ExecNoResponseRequest(ctx, request)
	if err != nil {
		return err
	}

	return nil
}
