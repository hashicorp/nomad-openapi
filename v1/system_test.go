package v1

import (
	"testing"

	"github.com/hashicorp/nomad/command/agent"
	"github.com/stretchr/testify/require"
)

func TestPostSystemGC(t *testing.T) {
	t.Parallel()
	httpTest(t, nil, func(s *agent.TestAgent) {
		testClient, err := NewTestClient(s)
		require.NoError(t, err)

		meta, err := testClient.System().GarbageCollection(writeOpts.Ctx())
		require.NoError(t, err)
		require.NotNil(t, meta)
	})
}

func TestPostSystemReconcileSummaries(t *testing.T) {
	t.Parallel()
	httpTest(t, nil, func(s *agent.TestAgent) {
		testClient, err := NewTestClient(s)
		require.NoError(t, err)

		meta, err := testClient.System().Reconcile(writeOpts.Ctx())
		require.NoError(t, err)
		require.NotNil(t, meta)
	})
}
