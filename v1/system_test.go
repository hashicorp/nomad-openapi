package v1

import (
	"testing"

	"github.com/hashicorp/nomad/command/agent"
	"github.com/stretchr/testify/require"
)

func TestSystem(t *testing.T) {
	httpTests(t, nil,
		APITestCase{"PutSystemGC", testPutSystemGC},
		APITestCase{"PutSystemReconcileSummaries", testPutSystemReconcileSummaries},
	)
}

func testPutSystemGC(t *testing.T, s *agent.TestAgent) {
	testClient, err := NewTestClient(s)
	require.NoError(t, err)

	err = testClient.System().GarbageCollect(writeOpts.Ctx())
	require.NoError(t, err)
}

func testPutSystemReconcileSummaries(t *testing.T, s *agent.TestAgent) {
	testClient, err := NewTestClient(s)
	require.NoError(t, err)

	err = testClient.System().Reconcile(writeOpts.Ctx())
	require.NoError(t, err)
}
