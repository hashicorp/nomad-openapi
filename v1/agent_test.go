package v1

import (
	"fmt"
	"testing"

	"github.com/hashicorp/nomad/command/agent"
	"github.com/hashicorp/nomad/nomad/mock"
	"github.com/stretchr/testify/require"
)

func TestAgentSelf(t *testing.T) {
	t.Parallel()
	httpTest(t, nil, func(s *agent.TestAgent) {
		testClient, err := NewTestClient(s)
		require.NoError(t, err)

		result, err := testClient.Agent().Self(queryOpts.Ctx())
		require.NoError(t, err)
		require.NotNil(t, result)

		// Check the result
		require.NotNil(t, *result.Config)
		require.NotNil(t, *result.Config.ACL)
		require.NotEmpty(t, *result.Stats)

		// Check the Vault config
		require.Empty(t, result.Config.Vault.Token)

		// Assign a Vault token and require it is redacted.
		s.Config.Vault.Token = "badc0deb-adc0-deba-dc0d-ebadc0debadc"
		result, err = testClient.Agent().Self(queryOpts.Ctx())
		require.NoError(t, err)
		require.NotNil(t, result)
		require.Equal(t, "<redacted>", *result.Config.Vault.Token)

		// Assign a ReplicationToken token and require it is redacted.
		s.Config.ACL.ReplicationToken = "badc0deb-adc0-deba-dc0d-ebadc0debadc"
		result, err = testClient.Agent().Self(queryOpts.Ctx())
		require.NoError(t, err)
		require.NotNil(t, result)
		require.Equal(t, "<redacted>", *result.Config.ACL.ReplicationToken)

		// Check the Consul config
		require.Empty(t, *result.Config.Consul.Token)

		// Assign a Consul token and require it is redacted.
		s.Config.Consul.Token = "badc0deb-adc0-deba-dc0d-ebadc0debadc"
		result, err = testClient.Agent().Self(queryOpts.Ctx())
		require.NoError(t, err)
		require.NotNil(t, result)
		require.Equal(t, "<redacted>", *result.Config.Consul.Token)

		// Check the Circonus config
		require.Empty(t, *result.Config.Telemetry.CirconusAPIToken)

		// Assign a Consul token and require it is redacted.
		s.Config.Telemetry.CirconusAPIToken = "badc0deb-adc0-deba-dc0d-ebadc0debadc"
		result, err = testClient.Agent().Self(queryOpts.Ctx())
		require.NoError(t, err)
		require.NotNil(t, result)
		require.Equal(t, "<redacted>", *result.Config.Telemetry.CirconusAPIToken)
	})
}

func TestAgentJoin(t *testing.T) {
	t.Parallel()
	httpTest(t, nil, func(s *agent.TestAgent) {
		// Determine the join address
		member := s.Agent.Server().LocalMember()
		addr := fmt.Sprintf("%s:%d", member.Addr, member.Port)

		testClient, err := NewTestClient(s)
		require.NoError(t, err)

		result, err := testClient.Agent().Join(writeOpts.Ctx(), []string{addr, addr})
		require.NoError(t, err)
		require.NotNil(t, result)
		require.NotNil(t, *result.NumJoined)
		require.NotNil(t, *result.Error)
		require.Equal(t, 2, *result.NumJoined)
		require.Empty(t, *result.Error)
	})
}

func TestAgentMembers(t *testing.T) {
	t.Parallel()
	httpTest(t, nil, func(s *agent.TestAgent) {
		testClient, err := NewTestClient(s)
		require.NoError(t, err)

		result, err := testClient.Agent().Members(queryOpts.Ctx())
		require.NoError(t, err)
		require.NotNil(t, result)
		require.NotNil(t, *result.Members)
		require.Len(t, *result.Members, 1)
	})
}

func TestAgentForceLeave(t *testing.T) {
	t.Parallel()
	httpTest(t, nil, func(s *agent.TestAgent) {
		testClient, err := NewTestClient(s)
		require.NoError(t, err)

		err = testClient.Agent().ForceLeave(writeOpts.Ctx())
		require.NoError(t, err)
	})
}

func TestAgentServers(t *testing.T) {
	t.Parallel()
	httpTest(t, nil, func(s *agent.TestAgent) {
		testClient, err := NewTestClient(s)
		require.NoError(t, err)

		result, err := testClient.Agent().Servers(queryOpts.Ctx())
		require.NoError(t, err)
		require.NotNil(t, result)

		instance := *result
		resultName := instance[0]
		require.Equal(t, s.Name, resultName)
	})
}

func TestAgentUpdateServers(t *testing.T) {
	t.Parallel()
	httpTest(t, nil, func(s *agent.TestAgent) {
		testClient, err := NewTestClient(s)
		require.NoError(t, err)

		opts := writeOpts.WithAuthToken(s.RootToken.SecretID)

		err = testClient.Agent().UpdateServers(opts.Ctx(), []string{"10.0.0.1", "10.0.0.2"})
		require.NoError(t, err)
	})
}

func TestAgentInstallKey(t *testing.T) {
	t.Parallel()

	key1 := "HS5lJ+XuTlYKWaeGYyG+/A=="
	key2 := "wH1Bn9hlJ0emgWB1JttVRA=="

	httpTest(t, func(c *agent.Config) {
		c.Server.EncryptKey = key1
	}, func(s *agent.TestAgent) {
		testClient, err := NewTestClient(s)
		require.NoError(t, err)

		opts := writeOpts.WithAuthToken(s.RootToken.SecretID)

		result, err := testClient.Agent().Keyring(opts.Ctx(), "install", key2)
		require.NoError(t, err)
		require.NotNil(t, result)

		result, err = testClient.Agent().Keyring(opts.Ctx(), "list", key2)
		require.NoError(t, err)
		require.NotNil(t, result)
		require.NotNil(t, *result.Keys)
		require.Len(t, *result.Keys, 2)
	})
}

func TestAgentHealthOK(t *testing.T) {
	t.Parallel()
	httpTest(t, nil, func(s *agent.TestAgent) {
		testClient, err := NewTestClient(s)
		require.NoError(t, err)

		// No ?type=
		{
			result, err := testClient.Agent().Health(queryOpts.Ctx(), []string{})
			require.NoError(t, err)
			require.NotNil(t, result)
			require.NotNil(t, *result.Client)
			require.True(t, *result.Client.Ok)
			require.Equal(t, "ok", *result.Client.Message)
			require.NotNil(t, *result.Server)
			require.True(t, *result.Server.Ok)
			require.Equal(t, "ok", *result.Server.Message)
		}

		// type=client
		{
			result, err := testClient.Agent().Health(queryOpts.Ctx(), []string{})
			require.NoError(t, err)
			require.NotNil(t, result)

			require.NotNil(t, *result.Client)
			require.True(t, *result.Client.Ok)
			require.Equal(t, "ok", *result.Client.Message)
			require.Nil(t, *result.Server)
		}

		// type=server
		{
			result, err := testClient.Agent().Health(queryOpts.Ctx(), []string{"server"})
			require.NoError(t, err)
			require.NotNil(t, result)

			require.NotNil(t, *result.Server)
			require.True(t, *result.Server.Ok)
			require.Equal(t, "ok", *result.Server.Message)
			require.Nil(t, *result.Client)
		}

		// type=client&type=server
		{
			result, err := testClient.Agent().Health(queryOpts.Ctx(), []string{"client", "server"})
			require.NoError(t, err)
			require.NotNil(t, result)

			require.NotNil(t, *result.Client)
			require.True(t, *result.Client.Ok)
			require.Equal(t, "ok", *result.Client.Message)
			require.NotNil(t, *result.Server)
			require.True(t, *result.Server.Ok)
			require.Equal(t, "ok", *result.Server.Message)
		}

	})
}

func TestAgentHost(t *testing.T) {
	t.Parallel()
	httpTest(t, nil, func(s *agent.TestAgent) {
		rpcRegister(t, s, mock.Job())

		testClient, err := NewTestClient(s)
		require.NoError(t, err)

		result, err := testClient.Agent().Host(queryOpts.Ctx())
		require.NoError(t, err)
		require.NotNil(t, result)
	})
}
