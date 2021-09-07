package v1

import (
	"context"

	"github.com/hashicorp/nomad-openapi/v1/client"
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
	request = m.client.setQueryOptions(ctx, request).(client.ApiGetMetricsSummaryRequest)

	result, _, err := request.Execute()
	if err != nil {
		return nil, err
	}

	return &result, nil
}
