package v1

import (
	"testing"

	"github.com/hashicorp/nomad/command/agent"
	"github.com/stretchr/testify/require"
)

func TestPutSystemGC(t *testing.T) {
	t.Parallel()
	httpTest(t, nil, func(s *agent.TestAgent) {
		testClient, err := NewTestClient(s)
		require.NoError(t, err)

		result, err := testClient.System().GarbageCollection(writeOpts.Ctx())
		require.NoError(t, err)
		require.NotNil(t, result)
	})
}

func TestPutSystemReconcileSummaries(t *testing.T) {
	t.Parallel()
	httpTest(t, nil, func(s *agent.TestAgent) {
		testClient, err := NewTestClient(s)
		require.NoError(t, err)

		result, err := testClient.System().Reconcile(writeOpts.Ctx())
		require.NoError(t, err)
		require.NotNil(t, result)
	})
}
