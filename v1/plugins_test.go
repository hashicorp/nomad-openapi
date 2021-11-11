package v1

import (
	"testing"

	"github.com/hashicorp/nomad/command/agent"
	"github.com/stretchr/testify/require"
)

func TestListPlugins(t *testing.T) {
	httpTest(t, nil, func(s *agent.TestAgent) {
		testClient, err := NewTestClient(s)
		require.NoError(t, err)

		result, err := testClient.Plugins().Get(queryOpts.Ctx())
		require.NoError(t, err)
		require.Len(t, *result, 0)
		require.NotNil(t, result)
	})
}

func TestGetCSIPlugin(t *testing.T) {
	httpTest(t, nil, func(s *agent.TestAgent) {
		testClient, err := NewTestClient(s)
		require.NoError(t, err)

		result, meta, err := testClient.Plugins().GetPlugin(queryOpts.Ctx(), "example-plugin")
		require.Error(t, err, "plugin not found")
		require.Nil(t, meta)
		require.Nil(t, result)
	})
}
