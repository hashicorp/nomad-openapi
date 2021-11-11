package v1

import (
	"testing"

	"github.com/hashicorp/nomad/command/agent"
	"github.com/stretchr/testify/require"
)

func TestGetDeployments(t *testing.T) {
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
	httpTest(t, nil, func(s *agent.TestAgent) {
		testClient, err := NewTestClient(s)
		require.NoError(t, err)

		result, err := testClient.Deployments().Fail(writeOpts.Ctx(), "5456bd7a-9fc0-c0dd-6131-cbee77f57577")
		require.Error(t, err, "deployment not found")
		require.Nil(t, result)
	})
}

func TestPostDeploymentPause(t *testing.T) {
	httpTest(t, nil, func(s *agent.TestAgent) {
		testClient, err := NewTestClient(s)
		require.NoError(t, err)

		result, err := testClient.Deployments().Pause(writeOpts.Ctx(), "5456bd7a-9fc0-c0dd-6131-cbee77f57577", true)
		require.Error(t, err, "deployment not found")
		require.Nil(t, result)
	})
}

func TestPostDeploymentPromote1(t *testing.T) {
	httpTest(t, nil, func(s *agent.TestAgent) {
		testClient, err := NewTestClient(s)
		require.NoError(t, err)

		result, err := testClient.Deployments().Promote(writeOpts.Ctx(), "5456bd7a-9fc0-c0dd-6131-cbee77f57577", true, nil)
		require.Error(t, err, "deployment not found")
		require.Nil(t, result)
	})
}

func TestPostDeploymentPromote2(t *testing.T) {
	httpTest(t, nil, func(s *agent.TestAgent) {
		testClient, err := NewTestClient(s)
		require.NoError(t, err)

		result, err := testClient.Deployments().Promote(writeOpts.Ctx(), "5456bd7a-9fc0-c0dd-6131-cbee77f57577", false, []string{"web", "api-server"})
		require.Error(t, err, "deployment not found")
		require.Nil(t, result)
	})
}

func TestPostDeploymentAllocationHealth1(t *testing.T) {
	httpTest(t, nil, func(s *agent.TestAgent) {
		testClient, err := NewTestClient(s)
		require.NoError(t, err)

		result, err := testClient.Deployments().AllocationHealth(writeOpts.Ctx(), "5456bd7a-9fc0-c0dd-6131-cbee77f57577", nil, nil)
		require.Error(t, err, "must specify at least one healthy/unhealthy allocation ID")
		require.Nil(t, result)
	})
}

func TestPostDeploymentAllocationHealth2(t *testing.T) {
	httpTest(t, nil, func(s *agent.TestAgent) {
		testClient, err := NewTestClient(s)
		require.NoError(t, err)

		result, err := testClient.Deployments().AllocationHealth(writeOpts.Ctx(), "5456bd7a-9fc0-c0dd-6131-cbee77f57577", []string{"eb13bc8a-7300-56f3-14c0-d4ad115ec3f5", "6584dad8-7ae3-360f-3069-0b4309711cc1"}, nil)
		require.Error(t, err, "deployment not found")
		require.Nil(t, result)
	})
}

func TestPostDeploymentUnblock(t *testing.T) {
	httpTest(t, nil, func(s *agent.TestAgent) {
		testClient, err := NewTestClient(s)
		require.NoError(t, err)

		result, err := testClient.Deployments().Unblock(writeOpts.Ctx(), "5456bd7a-9fc0-c0dd-6131-cbee77f57577")
		require.Error(t, err, "deployment not found")
		require.Nil(t, result)
	})
}
