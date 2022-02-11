package v1

import (
	"testing"

	"github.com/hashicorp/nomad/command/agent"
	"github.com/stretchr/testify/require"
)

func TestACLBootstrap(t *testing.T) {
	enableACL := func(c *agent.Config) {
		c.NomadConfig.ACLEnabled = true
	}
	httpTest(t, enableACL, func(s *agent.TestAgent) {
		// if os.Getenv("NOMAD_TOKEN") == "" {
		// 	t.Skip()
		// }

		client, err := NewTestClient(s)
		require.NoError(t, err)

		q := &QueryOpts{
			Region:    globalRegion,
			Namespace: defaultNamespace,
		}
		token, wMeta, err := client.ACL().Bootstrap(q.Ctx())
		require.NoError(t, err)
		require.NotNil(t, wMeta)
		require.NotNil(t, token)

		// Test query w/o token now that bootstrapped
		_, qMeta, err := client.Jobs().GetJobs(q.Ctx())
		require.Error(t, err)
		require.Nil(t, qMeta)

		t.Setenv("NOMAD_TOKEN", *token.SecretID)
		q.WithAuthToken(*token.SecretID)

		// Test query with token now that bootstrapped
		result, qMeta, err := client.Jobs().GetJobs(q.Ctx())
		require.NoError(t, err)
		require.NotNil(t, qMeta)
		require.NotNil(t, result)
	})
}
