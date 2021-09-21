package v1

import (
	"context"

	client "github.com/hashicorp/nomad-openapi/clients/go/v1"
)

type Metrics struct {
	client *Client
}

func (c *Client) Metrics() *Metrics {
	return &Metrics{client: c}
}

func (m *Metrics) MetricsApi() *client.MetricsApiService {
	return m.client.apiClient.MetricsApi
}

func (m *Metrics) GetMetricsSummary(ctx context.Context) (*client.MetricsSummary, error) {
	request := m.MetricsApi().GetMetricsSummary(m.client.Ctx)

	result, err := m.client.ExecRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	final := result.(client.MetricsSummary)
	return &final, nil
}
