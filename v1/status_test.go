package v1

import (
	"testing"

	"github.com/hashicorp/nomad/command/agent"
	"github.com/stretchr/testify/require"
)

func TestGetStatusLeader(t *testing.T) {
	httpTest(t, nil, func(s *agent.TestAgent) {
		testClient, err := NewTestClient(s)
		require.NoError(t, err)

		result, err := testClient.Status().Leader(queryOpts.Ctx())
		require.NoError(t, err)
		// using Contains because the port is random
		require.Contains(t, *result, "127.0.0.1:")
	})
}

func TestGetStatusPeers(t *testing.T) {
	httpTest(t, nil, func(s *agent.TestAgent) {
		testClient, err := NewTestClient(s)
		require.NoError(t, err)

		result, err := testClient.Status().Peers(queryOpts.Ctx())
		require.NoError(t, err)
		require.Len(t, *result, 1)
		// using Contains because the port is random
		require.Contains(t, (*result)[0], "127.0.0.1:")
	})
}
