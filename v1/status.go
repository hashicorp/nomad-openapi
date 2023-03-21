package v1

import (
	"context"

	client "github.com/flytocolors/nomad-openapi/clients/go/v1"
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

func (s *Status) Leader(ctx context.Context) (*string, OpenAPIError) {
	request := s.StatusApi().GetStatusLeader(s.client.Ctx)

	result, err := s.client.ExecNoMetaQuery(ctx, request)
	if err != nil {
		return nil, &APIError{error: err.Error()}
	}

	final := result.(string)
	return &final, nil
}

func (s *Status) Peers(ctx context.Context) (*[]string, OpenAPIError) {
	request := s.StatusApi().GetStatusPeers(s.client.Ctx)

	result, err := s.client.ExecNoMetaQuery(ctx, request)
	if err != nil {
		return nil, &APIError{error: err.Error()}
	}

	final := result.([]string)
	return &final, nil
}
