package v1

import (
	"testing"

	"github.com/hashicorp/nomad/command/agent"
	"github.com/stretchr/testify/require"
)

func TestEvaluations(t *testing.T) {
	httpTests(t, nil,
		APITestCase{"GetEvaluations", testGetEvaluations},
		APITestCase{"GetEvaluation", testGetEvaluation},
		APITestCase{"GetEvaluationAllocations", testGetEvaluationAllocations},
	)
}

func testGetEvaluations(t *testing.T, s *agent.TestAgent) {
	testClient, err := NewTestClient(s)
	require.NoError(t, err)

	result, meta, err := testClient.Evaluations().Evaluations(queryOpts.Ctx())
	require.NoError(t, err)
	require.NotNil(t, meta)
	require.Len(t, *result, 0)
	require.NotNil(t, result)
}

func testGetEvaluation(t *testing.T, s *agent.TestAgent) {
	testClient, err := NewTestClient(s)
	require.NoError(t, err)

	result, meta, err := testClient.Evaluations().GetEvaluation(queryOpts.Ctx(), "5456bd7a-9fc0-c0dd-6131-cbee77f57577")
	require.Error(t, err, "eval not found")
	require.Nil(t, meta)
	require.Nil(t, result)
}

func testGetEvaluationAllocations(t *testing.T, s *agent.TestAgent) {
	testClient, err := NewTestClient(s)
	require.NoError(t, err)

	result, meta, err := testClient.Evaluations().Allocations(queryOpts.Ctx(), "5456bd7a-9fc0-c0dd-6131-cbee77f57577")
	require.NoError(t, err)
	require.NotNil(t, meta)
	require.Len(t, *result, 0)
	require.NotNil(t, result)
}
