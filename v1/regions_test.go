package v1

import (
	"testing"

	"github.com/hashicorp/nomad/command/agent"
	"github.com/stretchr/testify/require"
)

func TestRegions(t *testing.T) {
	httpTests(t, nil,
		APITestCase{"GetRegions", testGetRegions},
	)
}

func testGetRegions(t *testing.T, s *agent.TestAgent) {
	testClient, err := NewTestClient(s)
	require.NoError(t, err)

	result, err := testClient.Regions().GetRegions(queryOpts.Ctx())
	require.NoError(t, err)
	require.Equal(t, []string{"global"}, *result)
}
