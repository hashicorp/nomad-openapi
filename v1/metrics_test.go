// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v1

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hashicorp/nomad/command/agent"
	"github.com/stretchr/testify/require"
)

func TestMetrics(t *testing.T) {
	httpTests(t, nil,
		APITestCase{"GetMetrics", testMetrics},
	)
}

func testMetrics(t *testing.T, s *agent.TestAgent) {
	// make a separate HTTP request first, to ensure Nomad has written metrics
	// and prevent a race condition
	req, err := http.NewRequest("GET", "/v1/agent/self", nil)
	require.NoError(t, err)
	respW := httptest.NewRecorder()
	_, _ = s.Server.AgentSelfRequest(respW, req)

	testClient, err := NewTestClient(s)
	require.NoError(t, err)

	// now make a metrics endpoint request, which should be already initialized
	// and written to
	result, err := testClient.Metrics().GetMetricsSummary(queryOpts.Ctx())
	require.NoError(t, err)

	require.NotEqual(t, 0, *result.Gauges)
}
