package v1

import (
	"testing"

	client "github.com/hashicorp/nomad-openapi/clients/go/v1"
	"github.com/hashicorp/nomad/command/agent"
	"github.com/hashicorp/nomad/helper"
	"github.com/stretchr/testify/require"
)

// NOTE: this is temporary until Int32ToPtr gets added to nomad/helper
func Int32ToPtr(i int32) *int32 {
	return &i
}

func TestOperator(t *testing.T) {
	httpTests(t, nil,
		APITestCase{"GetOperatorRaftConfiguration", testGetOperatorRaftConfiguration},
		APITestCase{"DeleteOperatorRaftPeer", testDeleteOperatorRaftPeer},
		APITestCase{"GetAutopilotConfiguration", testGetAutopilotConfiguration},
		APITestCase{"PutAutopilotConfiguration", testPutAutopilotConfiguration},
		APITestCase{"GetAutopilotHealth", testGetAutopilotHealth},
		APITestCase{"GetSchedulerConfiguration", testGetSchedulerConfiguration},
		APITestCase{"PostSchedulerConfiguration", testPostSchedulerConfiguration},
	)
}

func testGetOperatorRaftConfiguration(t *testing.T, s *agent.TestAgent) {
	testClient, err := NewTestClient(s)
	require.NoError(t, err)

	result, err := testClient.Operator().Raft(queryOpts.Ctx())
	require.NoError(t, err)
	require.NotNil(t, result)
}

func testDeleteOperatorRaftPeer(t *testing.T, s *agent.TestAgent) {
	testClient, err := NewTestClient(s)
	require.NoError(t, err)

	err = testClient.Operator().Peer(writeOpts.Ctx())
	require.Error(t, err, "Must specify either ?id with the server's ID or ?address with IP:port of peer to remove")
}

func testGetAutopilotConfiguration(t *testing.T, s *agent.TestAgent) {
	testClient, err := NewTestClient(s)
	require.NoError(t, err)

	result, err := testClient.Operator().Autopilot(queryOpts.Ctx())
	require.NoError(t, err)
	require.NotNil(t, result)
}

func testPutAutopilotConfiguration(t *testing.T, s *agent.TestAgent) {
	testClient, err := NewTestClient(s)
	require.NoError(t, err)

	autopilotOpts := &client.AutopilotConfiguration{
		CleanupDeadServers:      helper.BoolToPtr(true),
		LastContactThreshold:    helper.StringToPtr("200ms"),
		MaxTrailingLogs:         Int32ToPtr(250),
		ServerStabilizationTime: helper.StringToPtr("10s"),
		EnableRedundancyZones:   helper.BoolToPtr(false),
		DisableUpgradeMigration: helper.BoolToPtr(false),
		EnableCustomUpgrades:    helper.BoolToPtr(false),
	}

	result, err := testClient.Operator().UpdateAutopilot(writeOpts.Ctx(), autopilotOpts)
	require.NoError(t, err)
	require.NotNil(t, result)
}

func testGetAutopilotHealth(t *testing.T, s *agent.TestAgent) {
	testClient, err := NewTestClient(s)
	require.NoError(t, err)

	result, err := testClient.Operator().AutopilotHealth(queryOpts.Ctx())
	require.Error(t, err, "all servers must have raft_protocol set to 3 or higher to use this endpoint")
	require.Nil(t, result)
}

func testGetSchedulerConfiguration(t *testing.T, s *agent.TestAgent) {
	testClient, err := NewTestClient(s)
	require.NoError(t, err)

	result, meta, err := testClient.Operator().Scheduler(queryOpts.Ctx())
	require.NoError(t, err)
	require.NotNil(t, meta)
	require.NotNil(t, result)
}

func testPostSchedulerConfiguration(t *testing.T, s *agent.TestAgent) {
	testClient, err := NewTestClient(s)
	require.NoError(t, err)

	schedulerOpts := &client.SchedulerConfiguration{
		SchedulerAlgorithm:            helper.StringToPtr("spread"),
		MemoryOversubscriptionEnabled: helper.BoolToPtr(false),
		PreemptionConfig: &client.PreemptionConfig{
			BatchSchedulerEnabled:    helper.BoolToPtr(true),
			ServiceSchedulerEnabled:  helper.BoolToPtr(true),
			SysBatchSchedulerEnabled: helper.BoolToPtr(true),
			SystemSchedulerEnabled:   helper.BoolToPtr(true),
		},
	}

	result, meta, err := testClient.Operator().UpdateScheduler(writeOpts.Ctx(), schedulerOpts)
	require.NoError(t, err)
	require.NotNil(t, meta)
	require.NotNil(t, result)
}
