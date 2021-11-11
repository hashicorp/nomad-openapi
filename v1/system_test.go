package v1

import (
	"testing"

	"github.com/hashicorp/nomad/command/agent"
	"github.com/stretchr/testify/require"
)

func TestPutSystemGC(t *testing.T) {
	httpTest(t, nil, func(s *agent.TestAgent) {
		testClient, err := NewTestClient(s)
		require.NoError(t, err)

		err = testClient.System().GarbageCollect(writeOpts.Ctx())
		require.NoError(t, err)
	})
}

func TestPutSystemReconcileSummaries(t *testing.T) {
	httpTest(t, nil, func(s *agent.TestAgent) {
		testClient, err := NewTestClient(s)
		require.NoError(t, err)

		err = testClient.System().Reconcile(writeOpts.Ctx())
		require.NoError(t, err)
	})
}
