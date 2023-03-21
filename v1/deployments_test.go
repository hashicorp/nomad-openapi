package v1

import (
	"testing"

	"github.com/hashicorp/nomad/command/agent"
	"github.com/stretchr/testify/require"
)

func TestDeployments(t *testing.T) {
	httpTests(t, nil,
		APITestCase{"GetDeployments", testGetDeployments},
		APITestCase{"GetDeployment", testGetDeployment},
		APITestCase{"GetDeploymentAllocations", testGetDeploymentAllocations},
		APITestCase{"PostDeploymentFail", testPostDeploymentFail},
		APITestCase{"PostDeploymentPause", testPostDeploymentPause},
		APITestCase{"PostDeploymentPromote1", testPostDeploymentPromote1},
		APITestCase{"PostDeploymentPromote2", testPostDeploymentPromote2},
		APITestCase{"PostDeploymentAllocationHealth1", testPostDeploymentAllocationHealth1},
		APITestCase{"PostDeploymentAllocationHealth2", testPostDeploymentAllocationHealth2},
		APITestCase{"PostDeploymentUnblock", testPostDeploymentUnblock},
		APITestCase{"JobDeployment", testJobDeployment},
		APITestCase{"JobDeployments", testJobDeployments},
		APITestCase{"JobStable", testJobStable},
		APITestCase{"JobScaleStatus", testJobScaleStatus},
		APITestCase{"JobScaleTaskGroup", testJobScaleTaskGroup},
	)
}

func testGetDeployments(t *testing.T, s *agent.TestAgent) {
	testClient, err := NewTestClient(s)
	require.NoError(t, err)

	result, meta, err := testClient.Deployments().Deployments(queryOpts.Ctx())
	require.NoError(t, err)
	require.NotNil(t, meta)
	require.Len(t, *result, 0)
	require.NotNil(t, result)
}

func testGetDeployment(t *testing.T, s *agent.TestAgent) {
	testClient, err := NewTestClient(s)
	require.NoError(t, err)

	result, meta, err := testClient.Deployments().GetDeployment(queryOpts.Ctx(), "70638f62-5c19-193e-30d6-f9d6e689ab8e")
	require.Error(t, err, "deployment not found")
	require.Nil(t, meta)
	require.Nil(t, result)
}

func testGetDeploymentAllocations(t *testing.T, s *agent.TestAgent) {
	testClient, err := NewTestClient(s)
	require.NoError(t, err)

	result, meta, err := testClient.Deployments().Allocations(queryOpts.Ctx(), "5456bd7a-9fc0-c0dd-6131-cbee77f57577")
	require.NoError(t, err)
	require.NotNil(t, meta)
	require.Len(t, *result, 0)
	require.NotNil(t, result)
}

func testPostDeploymentFail(t *testing.T, s *agent.TestAgent) {
	testClient, err := NewTestClient(s)
	require.NoError(t, err)

	result, err := testClient.Deployments().Fail(writeOpts.Ctx(), "5456bd7a-9fc0-c0dd-6131-cbee77f57577")
	require.Error(t, err, "deployment not found")
	require.Nil(t, result)
}

func testPostDeploymentPause(t *testing.T, s *agent.TestAgent) {
	testClient, err := NewTestClient(s)
	require.NoError(t, err)

	result, err := testClient.Deployments().Pause(writeOpts.Ctx(), "5456bd7a-9fc0-c0dd-6131-cbee77f57577", true)
	require.Error(t, err, "deployment not found")
	require.Nil(t, result)
}

func testPostDeploymentPromote1(t *testing.T, s *agent.TestAgent) {
	testClient, err := NewTestClient(s)
	require.NoError(t, err)

	result, err := testClient.Deployments().Promote(writeOpts.Ctx(), "5456bd7a-9fc0-c0dd-6131-cbee77f57577", true, nil)
	require.Error(t, err, "deployment not found")
	require.Nil(t, result)
}

func testPostDeploymentPromote2(t *testing.T, s *agent.TestAgent) {
	testClient, err := NewTestClient(s)
	require.NoError(t, err)

	result, err := testClient.Deployments().Promote(writeOpts.Ctx(), "5456bd7a-9fc0-c0dd-6131-cbee77f57577", false, []string{"web", "api-server"})
	require.Error(t, err, "deployment not found")
	require.Nil(t, result)
}

func testPostDeploymentAllocationHealth1(t *testing.T, s *agent.TestAgent) {
	testClient, err := NewTestClient(s)
	require.NoError(t, err)

	result, err := testClient.Deployments().AllocationHealth(writeOpts.Ctx(), "5456bd7a-9fc0-c0dd-6131-cbee77f57577", nil, nil)
	require.Error(t, err, "must specify at least one healthy/unhealthy allocation ID")
	require.Nil(t, result)
}

func testPostDeploymentAllocationHealth2(t *testing.T, s *agent.TestAgent) {
	testClient, err := NewTestClient(s)
	require.NoError(t, err)

	result, err := testClient.Deployments().AllocationHealth(writeOpts.Ctx(), "5456bd7a-9fc0-c0dd-6131-cbee77f57577", []string{"eb13bc8a-7300-56f3-14c0-d4ad115ec3f5", "6584dad8-7ae3-360f-3069-0b4309711cc1"}, nil)
	require.Error(t, err, "deployment not found")
	require.Nil(t, result)
}

func testPostDeploymentUnblock(t *testing.T, s *agent.TestAgent) {
	testClient, err := NewTestClient(s)
	require.NoError(t, err)

	result, err := testClient.Deployments().Unblock(writeOpts.Ctx(), "5456bd7a-9fc0-c0dd-6131-cbee77f57577")
	require.Error(t, err, "deployment not found")
	require.Nil(t, result)
}
