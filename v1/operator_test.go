package v1

import (
	"testing"

	"github.com/hashicorp/nomad/command/agent"
	"github.com/stretchr/testify/require"
)

// ERROR: json: cannot unmarshal object into Go value of type []client.RaftServer
// func TestGetOperatorRaftConfiguration(t *testing.T) {
// 	httpTest(t, nil, func(s *agent.TestAgent) {
// 		testClient, err := NewTestClient(s)
// 		require.NoError(t, err)

// 		result, err := testClient.Operator().Raft(queryOpts.Ctx())
// 		require.NoError(t, err)
// 		require.NotNil(t, result)
// 	})
// }

func TestDeleteOperatorRaftPeer(t *testing.T) {
	httpTest(t, nil, func(s *agent.TestAgent) {
		testClient, err := NewTestClient(s)
		require.NoError(t, err)

		err = testClient.Operator().Peer(writeOpts.Ctx())
		require.Error(t, err, "Must specify either ?id with the server's ID or ?address with IP:port of peer to remove")
	})
}

// func TestGetAutopilotConfiguration(t *testing.T) {
// 	httpTest(t, nil, func(s *agent.TestAgent) {
// 		testClient, err := NewTestClient(s)
// 		require.NoError(t, err)

// 		result, err := testClient.Operator().Autopilot(queryOpts.Ctx())
// 		require.NoError(t, err)
// 		require.NotNil(t, result)
// 	})
// }

// NOTE: payload has some values as string, but config says they should be int64...
// https://www.nomadproject.io/api-docs/operator/autopilot#update-autopilot-configuration
// func TestPutAutopilotConfiguration(t *testing.T) {
// 	httpTest(t, nil, func(s *agent.TestAgent) {
// 		testClient, err := NewTestClient(s)
// 		require.NoError(t, err)

// 		result, err := testClient.Operator().UpdateAutopilot(writeOpts.Ctx(), true, "200ms", 250, "10s", false, false, false)
// 		require.NoError(t, err)
// 		require.NotNil(t, result)
// 	})
// }

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

// QUESTION: how do I pass in values for preemptionconfig?
// func TestPostSchedulerConfiguration(t *testing.T) {
// 	httpTest(t, nil, func(s *agent.TestAgent) {
// 		testClient, err := NewTestClient(s)
// 		require.NoError(t, err)

// 		result, meta, err := testClient.Operator().UpdateScheduler(writeOpts.Ctx(), "spread", false, ??????????)
// 		require.NoError(t, err)
// 		require.NotNil(t, meta)
// 		require.NotNil(t, result)
// 	})
// }
