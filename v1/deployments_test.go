package v1

import (
	"testing"

	"github.com/hashicorp/nomad/command/agent"
	"github.com/stretchr/testify/require"
)

func TestGetDeployments(t *testing.T) {
	t.Parallel()
	httpTest(t, nil, func(s *agent.TestAgent) {
		testClient, err := NewTestClient(s)
		require.NoError(t, err)

		result, meta, err := testClient.Deployments().Deployments(queryOpts.Ctx())
		require.NoError(t, err)
		require.NotNil(t, meta)
		require.Len(t, *result, 0)
		require.NotNil(t, result)
	})
}

func TestGetDeployment(t *testing.T) {
	t.Parallel()
	httpTest(t, nil, func(s *agent.TestAgent) {
		testClient, err := NewTestClient(s)
		require.NoError(t, err)

		result, meta, err := testClient.Deployments().GetDeployment(queryOpts.Ctx(), "70638f62-5c19-193e-30d6-f9d6e689ab8e")
		require.Error(t, err, "deployment not found")
		require.Nil(t, meta)
		require.Nil(t, result)
	})
}

func TestGetDeploymentAllocations(t *testing.T) {
	t.Parallel()
	httpTest(t, nil, func(s *agent.TestAgent) {
		testClient, err := NewTestClient(s)
		require.NoError(t, err)

		result, meta, err := testClient.Deployments().Allocations(queryOpts.Ctx(), "5456bd7a-9fc0-c0dd-6131-cbee77f57577")
		require.NoError(t, err)
		require.NotNil(t, meta)
		require.Len(t, *result, 0)
		require.NotNil(t, result)
	})
}

func TestPostDeploymentFail(t *testing.T) {
	t.Parallel()
	httpTest(t, nil, func(s *agent.TestAgent) {
		testClient, err := NewTestClient(s)
		require.NoError(t, err)

		result, err := testClient.Deployments().Fail(writeOpts.Ctx(), "5456bd7a-9fc0-c0dd-6131-cbee77f57577")
		require.Error(t, err, "deployment not found")
		require.Nil(t, result)
	})
}
