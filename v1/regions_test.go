package v1

import (
	"testing"

	"github.com/hashicorp/nomad/command/agent"
	"github.com/stretchr/testify/require"
)

func TestRegions_GetRegions(t *testing.T) {
	httpTest(t, nil, func(s *agent.TestAgent) {

		// Make the HTTP request
		testClient, err := NewTestClient(s)
		require.NoError(t, err)

		result, err := testClient.Regions().GetRegions(queryOpts.Ctx())
		require.NoError(t, err)
		require.Equal(t, []string{"global"}, *result)
	})
}
