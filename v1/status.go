package v1

import (
	"context"

	client "github.com/hashicorp/nomad-openapi/clients/go/v1"
)

type Status struct {
	client *Client
}

func (c *Client) Status() *Status {
	return &Status{client: c}
}

func (s *Status) StatusApi() *client.StatusApiService {
	return s.client.apiClient.StatusApi
}

func (s *Status) Leader(ctx context.Context) (*string, error) {
	request := s.StatusApi().GetStatusLeader(s.client.Ctx)

	// TODO: This endpoint also needs a helper that takes QueryOpts but returns no QueryMeta.
	result, err := s.client.ExecRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	final := result.(string)
	return &final, nil
}

func (s *Status) Peers(ctx context.Context) (*[]string, error) {
	request := s.StatusApi().GetStatusPeers(s.client.Ctx)

	// TODO: This endpoint also needs a helper that takes QueryOpts but returns no QueryMeta.
	result, err := s.client.ExecRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	final := result.([]string)
	return &final, nil
}
