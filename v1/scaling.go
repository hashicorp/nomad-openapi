package v1

import (
	"context"
	"errors"

	client "github.com/hashicorp/nomad-openapi/clients/go/v1"
)

type Scaling struct {
	client *Client
}

func (c *Client) Scaling() *Scaling {
	return &Scaling{client: c}
}

func (s *Scaling) ScalingApi() *client.ScalingApiService {
	return s.client.apiClient.ScalingApi
}

func (s *Scaling) Policies(ctx context.Context) (*[]client.ScalingPolicyListStub, *QueryMeta, error) {
	request := s.ScalingApi().GetScalingPolicies(s.client.Ctx)
	result, meta, err := s.client.ExecQuery(ctx, request)
	if err != nil {
		return nil, nil, err
	}

	final := result.([]client.ScalingPolicyListStub)
	return &final, meta, nil
}

func (s *Scaling) GetPolicy(ctx context.Context, policyID string) (*client.ScalingPolicy, *QueryMeta, error) {
	if policyID == "" {
		return nil, nil, errors.New("scaling policy id is required")
	}

	request := s.ScalingApi().GetScalingPolicy(s.client.Ctx, policyID)
	result, meta, err := s.client.ExecQuery(ctx, request)
	if err != nil {
		return nil, nil, err
	}

	final := result.(client.ScalingPolicy)
	return &final, meta, nil
}
