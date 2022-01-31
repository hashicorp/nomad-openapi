package v1

import (
	"testing"

	client "github.com/hashicorp/nomad-openapi/clients/go/v1"
	"github.com/hashicorp/nomad/command/agent"
	"github.com/stretchr/testify/require"
)

// boolToPtr returns the pointer to a boolean
func boolToPtr(b bool) *bool {
	return &b
}

func TestGetOperatorRaftConfiguration(t *testing.T) {
	httpTest(t, nil, func(s *agent.TestAgent) {
		testClient, err := NewTestClient(s)
		require.NoError(t, err)

		result, err := testClient.Operator().Raft(queryOpts.Ctx())
		require.NoError(t, err)
		require.NotNil(t, result)
	})
}

func TestDeleteOperatorRaftPeer(t *testing.T) {
	httpTest(t, nil, func(s *agent.TestAgent) {
		testClient, err := NewTestClient(s)
		require.NoError(t, err)

		err = testClient.Operator().Peer(writeOpts.Ctx())
		require.Error(t, err, "Must specify either ?id with the server's ID or ?address with IP:port of peer to remove")
	})
}

func TestGetAutopilotConfiguration(t *testing.T) {
	httpTest(t, nil, func(s *agent.TestAgent) {
		testClient, err := NewTestClient(s)
		require.NoError(t, err)

		result, err := testClient.Operator().Autopilot(queryOpts.Ctx())
		require.NoError(t, err)
		require.NotNil(t, result)
	})
}

func TestPutAutopilotConfiguration(t *testing.T) {
	httpTest(t, nil, func(s *agent.TestAgent) {
		testClient, err := NewTestClient(s)
		require.NoError(t, err)

		result, err := testClient.Operator().UpdateAutopilot(writeOpts.Ctx(), true, "200ms", 250, "10s", false, false, false)
		require.NoError(t, err)
		require.NotNil(t, result)
	})
}

func TestGetAutopilotHealth(t *testing.T) {
	httpTest(t, nil, func(s *agent.TestAgent) {
		testClient, err := NewTestClient(s)
		require.NoError(t, err)

		result, err := testClient.Operator().AutopilotHealth(queryOpts.Ctx())
		require.Error(t, err, "all servers must have raft_protocol set to 3 or higher to use this endpoint")
		require.Nil(t, result)
	})
}

func TestGetSchedulerConfiguration(t *testing.T) {
	httpTest(t, nil, func(s *agent.TestAgent) {
		testClient, err := NewTestClient(s)
		require.NoError(t, err)

		result, meta, err := testClient.Operator().Scheduler(queryOpts.Ctx())
		require.NoError(t, err)
		require.NotNil(t, meta)
		require.NotNil(t, result)
	})
}

func TestPostSchedulerConfiguration(t *testing.T) {
	httpTest(t, nil, func(s *agent.TestAgent) {
		testClient, err := NewTestClient(s)
		require.NoError(t, err)

		config := &client.PreemptionConfig{
			BatchSchedulerEnabled:    boolToPtr(true),
			ServiceSchedulerEnabled:  boolToPtr(true),
			SysBatchSchedulerEnabled: boolToPtr(true),
			SystemSchedulerEnabled:   boolToPtr(true),
		}

		result, meta, err := testClient.Operator().UpdateScheduler(writeOpts.Ctx(), "spread", false, config)
		require.NoError(t, err)
		require.NotNil(t, meta)
		require.NotNil(t, result)
	})
}
