package v1

import (
	"context"
	"errors"

	client "github.com/hashicorp/nomad-openapi/clients/go/v1"
)

type Evaluations struct {
	client *Client
}

func (c *Client) Evaluations() *Evaluations {
	return &Evaluations{client: c}
}

func (e *Evaluations) EvaluationsApi() *client.EvaluationsApiService {
	return e.client.apiClient.EvaluationsApi
}

func (e *Evaluations) Evaluations(ctx context.Context) (*[]client.Evaluation, *QueryMeta, error) {
	request := e.EvaluationsApi().GetEvaluations(e.client.Ctx)
	result, meta, err := e.client.ExecQuery(ctx, request)
	if err != nil {
		return nil, nil, err
	}

	final := result.([]client.Evaluation)
	return &final, meta, nil
}

func (e *Evaluations) GetEvaluation(ctx context.Context, evalID string) (*client.Evaluation, *QueryMeta, error) {
	if evalID == "" {
		return nil, nil, errors.New("evaluation id is required")
	}

	request := e.EvaluationsApi().GetEvaluation(e.client.Ctx, evalID)
	result, meta, err := e.client.ExecQuery(ctx, request)
	if err != nil {
		return nil, nil, err
	}

	final := result.(client.Evaluation)
	return &final, meta, nil
}

func (e *Evaluations) Allocations(ctx context.Context, evalID string) (*[]client.AllocationListStub, *QueryMeta, error) {
	if evalID == "" {
		return nil, nil, errors.New("evaluation id is required")
	}

	request := e.EvaluationsApi().GetEvaluationAllocations(e.client.Ctx, evalID)
	result, meta, err := e.client.ExecQuery(ctx, request)
	if err != nil {
		return nil, nil, err
	}

	final := result.([]client.AllocationListStub)
	return &final, meta, nil
}
