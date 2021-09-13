package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hashicorp/nomad/acl"
	"github.com/hashicorp/nomad/command/agent"
	"github.com/hashicorp/nomad/nomad/mock"
	"github.com/hashicorp/nomad/nomad/structs"
	"github.com/stretchr/testify/require"
)

func TestAgentSelf(t *testing.T) {
	t.Parallel()
	httpTest(t, nil, func(s *agent.TestAgent) {
		rpcRegister(t, s, mock.Job())

		testClient, err := NewTestClient(s)
		require.NoError(t, err)

		result, meta, err := testClient.TODO().TODO(TODOOpts.Ctx())
		require.NoError(t, err)
		require.NotNil(t, result)
		require.NotNil(t, meta)

		req, err := http.NewRequest("GET", "/v1/agent/self", nil)
		require.NoError(err)
		respW := httptest.NewRecorder()

		// Make the request
		obj, err := s.Server.AgentSelfRequest(respW, req)
		require.NoError(err)

		// Check the job
		self := obj.(agentSelf)
		require.NotNil(self.Config)
		require.NotNil(self.Config.ACL)
		require.NotEmpty(self.Stats)

		// Check the Vault config
		require.Empty(self.Config.Vault.Token)

		// Assign a Vault token and require it is redacted.
		s.Config.Vault.Token = "badc0deb-adc0-deba-dc0d-ebadc0debadc"
		respW = httptest.NewRecorder()
		obj, err = s.Server.AgentSelfRequest(respW, req)
		require.NoError(err)
		self = obj.(agentSelf)
		require.Equal("<redacted>", self.Config.Vault.Token)

		// Assign a ReplicationToken token and require it is redacted.
		s.Config.ACL.ReplicationToken = "badc0deb-adc0-deba-dc0d-ebadc0debadc"
		respW = httptest.NewRecorder()
		obj, err = s.Server.AgentSelfRequest(respW, req)
		require.NoError(err)
		self = obj.(agentSelf)
		require.Equal("<redacted>", self.Config.ACL.ReplicationToken)

		// Check the Consul config
		require.Empty(self.Config.Consul.Token)

		// Assign a Consul token and require it is redacted.
		s.Config.Consul.Token = "badc0deb-adc0-deba-dc0d-ebadc0debadc"
		respW = httptest.NewRecorder()
		obj, err = s.Server.AgentSelfRequest(respW, req)
		require.NoError(err)
		self = obj.(agentSelf)
		require.Equal("<redacted>", self.Config.Consul.Token)

		// Check the Circonus config
		require.Empty(self.Config.Telemetry.CirconusAPIToken)

		// Assign a Consul token and require it is redacted.
		s.Config.Telemetry.CirconusAPIToken = "badc0deb-adc0-deba-dc0d-ebadc0debadc"
		respW = httptest.NewRecorder()
		obj, err = s.Server.AgentSelfRequest(respW, req)
		require.NoError(err)
		self = obj.(agentSelf)
		require.Equal("<redacted>", self.Config.Telemetry.CirconusAPIToken)
	})
}

func TestAgentJoin(t *testing.T) {
	t.Parallel()
	httpTest(t, nil, func(s *agent.TestAgent) {
		rpcRegister(t, s, mock.Job())

		testClient, err := NewTestClient(s)
		require.NoError(t, err)

		result, meta, err := testClient.TODO().TODO(TODOOpts.Ctx())
		require.NoError(t, err)
		require.NotNil(t, result)
		require.NotNil(t, meta)

		// Determine the join address
		member := s.Agent.Server().LocalMember()
		addr := fmt.Sprintf("%s:%d", member.Addr, member.Port)

		// Make the HTTP request
		req, err := http.NewRequest("PUT",
			fmt.Sprintf("/v1/agent/join?address=%s&address=%s", addr, addr), nil)
		if err != nil {
			t.Fatalf("err: %v", err)
		}
		respW := httptest.NewRecorder()

		// Make the request
		obj, err := s.Server.AgentJoinRequest(respW, req)
		if err != nil {
			t.Fatalf("err: %v", err)
		}

		// Check the job
		join := obj.(joinResult)
		if join.NumJoined != 2 {
			t.Fatalf("bad: %#v", join)
		}
		if join.Error != "" {
			t.Fatalf("bad: %#v", join)
		}
	})
}

func TestAgentMembers(t *testing.T) {
	t.Parallel()
	httpTest(t, nil, func(s *agent.TestAgent) {
		rpcRegister(t, s, mock.Job())

		testClient, err := NewTestClient(s)
		require.NoError(t, err)

		result, meta, err := testClient.TODO().TODO(TODOOpts.Ctx())
		require.NoError(t, err)
		require.NotNil(t, result)
		require.NotNil(t, meta)

		// Make the HTTP request
		req, err := http.NewRequest("GET", "/v1/agent/members", nil)
		if err != nil {
			t.Fatalf("err: %v", err)
		}
		respW := httptest.NewRecorder()

		// Make the request
		obj, err := s.Server.AgentMembersRequest(respW, req)
		if err != nil {
			t.Fatalf("err: %v", err)
		}

		// Check the job
		members := obj.(structs.ServerMembersResponse)
		if len(members.Members) != 1 {
			t.Fatalf("bad: %#v", members.Members)
		}

	})
}

func TestAgentForceLeave(t *testing.T) {
	t.Parallel()
	httpTest(t, nil, func(s *agent.TestAgent) {
		rpcRegister(t, s, mock.Job())

		testClient, err := NewTestClient(s)
		require.NoError(t, err)

		result, meta, err := testClient.TODO().TODO(TODOOpts.Ctx())
		require.NoError(t, err)
		require.NotNil(t, result)
		require.NotNil(t, meta)

		// Make the HTTP request
		req, err := http.NewRequest("PUT", "/v1/agent/force-leave?node=foo", nil)
		if err != nil {
			t.Fatalf("err: %v", err)
		}
		respW := httptest.NewRecorder()

		// Make the request
		_, err = s.Server.AgentForceLeaveRequest(respW, req)
		if err != nil {
			t.Fatalf("err: %v", err)
		}

	})
}

func TestAgentList(t *testing.T) {
	t.Parallel()
	httpTest(t, nil, func(s *agent.TestAgent) {
		rpcRegister(t, s, mock.Job())

		testClient, err := NewTestClient(s)
		require.NoError(t, err)

		result, meta, err := testClient.TODO().TODO(TODOOpts.Ctx())
		require.NoError(t, err)
		require.NotNil(t, result)
		require.NotNil(t, meta)

	})
}

func TestAgentServersUpdate(t *testing.T) {
	t.Parallel()
	httpTest(t, nil, func(s *agent.TestAgent) {
		rpcRegister(t, s, mock.Job())

		testClient, err := NewTestClient(s)
		require.NoError(t, err)

		result, meta, err := testClient.TODO().TODO(TODOOpts.Ctx())
		require.NoError(t, err)
		require.NotNil(t, result)
		require.NotNil(t, meta)

	})
}

func TestAgentInstallKey(t *testing.T) {
	t.Parallel()

	key1 := "HS5lJ+XuTlYKWaeGYyG+/A=="
	key2 := "wH1Bn9hlJ0emgWB1JttVRA=="

	httpTest(t, nil, func(s *agent.TestAgent) {
		rpcRegister(t, s, mock.Job())

		testClient, err := NewTestClient(s)
		require.NoError(t, err)

		result, meta, err := testClient.TODO().TODO(TODOOpts.Ctx())
		require.NoError(t, err)
		require.NotNil(t, result)
		require.NotNil(t, meta)

		b, err := json.Marshal(&structs.KeyringRequest{Key: key2})
		if err != nil {
			t.Fatalf("err: %v", err)
		}
		req, err := http.NewRequest("GET", "/v1/agent/keyring/install", bytes.NewReader(b))
		if err != nil {
			t.Fatalf("err: %s", err)
		}
		respW := httptest.NewRecorder()

		_, err = s.Server.KeyringOperationRequest(respW, req)
		if err != nil {
			t.Fatalf("err: %s", err)
		}
		req, err = http.NewRequest("GET", "/v1/agent/keyring/list", bytes.NewReader(b))
		if err != nil {
			t.Fatalf("err: %s", err)
		}
		respW = httptest.NewRecorder()

		out, err := s.Server.KeyringOperationRequest(respW, req)
		if err != nil {
			t.Fatalf("err: %s", err)
		}
		kresp := out.(structs.KeyringResponse)
		if len(kresp.Keys) != 2 {
			t.Fatalf("bad: %v", kresp)
		}
	})
}

func TestAgentListKeys(t *testing.T) {
	t.Parallel()
	httpTest(t, nil, func(s *agent.TestAgent) {
		rpcRegister(t, s, mock.Job())

		testClient, err := NewTestClient(s)
		require.NoError(t, err)

		result, meta, err := testClient.TODO().TODO(TODOOpts.Ctx())
		require.NoError(t, err)
		require.NotNil(t, result)
		require.NotNil(t, meta)

		require := require.New(t)

		key1 := "HS5lJ+XuTlYKWaeGYyG+/A=="

		cb := func(c *Config) {
			c.Server.EncryptKey = key1
		}

		// Make the HTTP request
		req, err := http.NewRequest("GET", "/v1/agent/keyring/list", nil)
		require.Nil(err)

		// Try request without a token and expect failure
		{
			respW := httptest.NewRecorder()
			_, err := s.Server.KeyringOperationRequest(respW, req)
			require.NotNil(err)
			require.Equal(err.Error(), structs.ErrPermissionDenied.Error())
		}

		// Try request with an invalid token and expect failure
		{
			respW := httptest.NewRecorder()
			token := mock.CreatePolicyAndToken(t, state, 1005, "invalid", mock.AgentPolicy(acl.PolicyRead))
			setToken(req, token)
			_, err := s.Server.KeyringOperationRequest(respW, req)
			require.NotNil(err)
			require.Equal(err.Error(), structs.ErrPermissionDenied.Error())
		}

		// Try request with a valid token
		{
			respW := httptest.NewRecorder()
			token := mock.CreatePolicyAndToken(t, state, 1007, "valid", mock.AgentPolicy(acl.PolicyWrite))
			setToken(req, token)
			out, err := s.Server.KeyringOperationRequest(respW, req)
			require.Nil(err)
			kresp := out.(structs.KeyringResponse)
			require.Len(kresp.Keys, 1)
			require.Contains(kresp.Keys, key1)
		}

		// Try request with a root token
		{
			respW := httptest.NewRecorder()
			setToken(req, s.RootToken)
			out, err := s.Server.KeyringOperationRequest(respW, req)
			require.Nil(err)
			kresp := out.(structs.KeyringResponse)
			require.Len(kresp.Keys, 1)
			require.Contains(kresp.Keys, key1)
		}
	})
}

func TestAgentRemoveKey(t *testing.T) {
	t.Parallel()
	httpTest(t, nil, func(s *agent.TestAgent) {
		rpcRegister(t, s, mock.Job())

		testClient, err := NewTestClient(s)
		require.NoError(t, err)

		result, meta, err := testClient.TODO().TODO(TODOOpts.Ctx())
		require.NoError(t, err)
		require.NotNil(t, result)
		require.NotNil(t, meta)

		key1 := "HS5lJ+XuTlYKWaeGYyG+/A=="
		key2 := "wH1Bn9hlJ0emgWB1JttVRA=="

		b, err := json.Marshal(&structs.KeyringRequest{Key: key2})
		if err != nil {
			t.Fatalf("err: %v", err)
		}

		req, err := http.NewRequest("GET", "/v1/agent/keyring/install", bytes.NewReader(b))
		if err != nil {
			t.Fatalf("err: %s", err)
		}
		respW := httptest.NewRecorder()
		_, err = s.Server.KeyringOperationRequest(respW, req)
		if err != nil {
			t.Fatalf("err: %s", err)
		}

		req, err = http.NewRequest("GET", "/v1/agent/keyring/remove", bytes.NewReader(b))
		if err != nil {
			t.Fatalf("err: %s", err)
		}
		respW = httptest.NewRecorder()
		if _, err = s.Server.KeyringOperationRequest(respW, req); err != nil {
			t.Fatalf("err: %s", err)
		}

		req, err = http.NewRequest("GET", "/v1/agent/keyring/list", nil)
		if err != nil {
			t.Fatalf("err: %s", err)
		}
		respW = httptest.NewRecorder()
		out, err := s.Server.KeyringOperationRequest(respW, req)
		if err != nil {
			t.Fatalf("err: %s", err)
		}
		kresp := out.(structs.KeyringResponse)
		if len(kresp.Keys) != 1 {
			t.Fatalf("bad: %v", kresp)
		}

	})
}

func TestAgentHealthOK(t *testing.T) {
	t.Parallel()
	httpTest(t, nil, func(s *agent.TestAgent) {
		rpcRegister(t, s, mock.Job())

		testClient, err := NewTestClient(s)
		require.NoError(t, err)

		result, meta, err := testClient.TODO().TODO(TODOOpts.Ctx())
		require.NoError(t, err)
		require.NotNil(t, result)
		require.NotNil(t, meta)

		// No ?type=
		{
			req, err := http.NewRequest("GET", "/v1/agent/health", nil)
			require.Nil(err)

			respW := httptest.NewRecorder()
			healthI, err := s.Server.HealthRequest(respW, req)
			require.Nil(err)
			require.Equal(http.StatusOK, respW.Code)
			require.NotNil(healthI)
			health := healthI.(*healthResponse)
			require.NotNil(health.Client)
			require.True(health.Client.Ok)
			require.Equal("ok", health.Client.Message)
			require.NotNil(health.Server)
			require.True(health.Server.Ok)
			require.Equal("ok", health.Server.Message)
		}

		// type=client
		{
			req, err := http.NewRequest("GET", "/v1/agent/health?type=client", nil)
			require.Nil(err)

			respW := httptest.NewRecorder()
			healthI, err := s.Server.HealthRequest(respW, req)
			require.Nil(err)
			require.Equal(http.StatusOK, respW.Code)
			require.NotNil(healthI)
			health := healthI.(*healthResponse)
			require.NotNil(health.Client)
			require.True(health.Client.Ok)
			require.Equal("ok", health.Client.Message)
			require.Nil(health.Server)
		}

		// type=server
		{
			req, err := http.NewRequest("GET", "/v1/agent/health?type=server", nil)
			require.Nil(err)

			respW := httptest.NewRecorder()
			healthI, err := s.Server.HealthRequest(respW, req)
			require.Nil(err)
			require.Equal(http.StatusOK, respW.Code)
			require.NotNil(healthI)
			health := healthI.(*healthResponse)
			require.NotNil(health.Server)
			require.True(health.Server.Ok)
			require.Equal("ok", health.Server.Message)
			require.Nil(health.Client)
		}

		// type=client&type=server
		{
			req, err := http.NewRequest("GET", "/v1/agent/health?type=client&type=server", nil)
			require.Nil(err)

			respW := httptest.NewRecorder()
			healthI, err := s.Server.HealthRequest(respW, req)
			require.Nil(err)
			require.Equal(http.StatusOK, respW.Code)
			require.NotNil(healthI)
			health := healthI.(*healthResponse)
			require.NotNil(health.Client)
			require.True(health.Client.Ok)
			require.Equal("ok", health.Client.Message)
			require.NotNil(health.Server)
			require.True(health.Server.Ok)
			require.Equal("ok", health.Server.Message)
		}

	})
}

func TestAgentHealthBadClient(t *testing.T) {
	t.Parallel()
	// Disable client to make server unhealthy if requested
	cb := func(c *Config) {
		c.Client.Enabled = false
	}
	httpTest(t, nil, func(s *agent.TestAgent) {
		rpcRegister(t, s, mock.Job())

		testClient, err := NewTestClient(s)
		require.NoError(t, err)

		result, meta, err := testClient.TODO().TODO(TODOOpts.Ctx())
		require.NoError(t, err)
		require.NotNil(t, result)
		require.NotNil(t, meta)

		// No ?type= means client is just skipped
		{
			req, err := http.NewRequest("GET", "/v1/agent/health", nil)
			require.Nil(err)

			respW := httptest.NewRecorder()
			healthI, err := s.Server.HealthRequest(respW, req)
			require.Nil(err)
			require.Equal(http.StatusOK, respW.Code)
			require.NotNil(healthI)
			health := healthI.(*healthResponse)
			require.NotNil(health.Server)
			require.True(health.Server.Ok)
			require.Equal("ok", health.Server.Message)
			require.Nil(health.Client)
		}

		// type=client means client is considered unhealthy
		{
			req, err := http.NewRequest("GET", "/v1/agent/health?type=client", nil)
			require.Nil(err)

			respW := httptest.NewRecorder()
			_, err = s.Server.HealthRequest(respW, req)
			require.NotNil(err)
			httpErr, ok := err.(HTTPCodedError)
			require.True(ok)
			require.Equal(500, httpErr.Code())
			require.Equal(`{"client":{"ok":false,"message":"client not enabled"}}`, err.Error())
		}
	})
}

func TestAgentHealthBadServer(t *testing.T) {
	t.Parallel()

	httpTest(t, nil, func(s *agent.TestAgent) {
		rpcRegister(t, s, mock.Job())

		testClient, err := NewTestClient(s)
		require.NoError(t, err)

		result, meta, err := testClient.TODO().TODO(TODOOpts.Ctx())
		require.NoError(t, err)
		require.NotNil(t, result)
		require.NotNil(t, meta)

		serverAgent := NewTestAgent(t, "server", nil)
		defer serverAgent.Shutdown()

		s := makeHTTPServer(t, func(c *Config) {
			// Disable server to make server health unhealthy if requested
			c.Server.Enabled = false
			c.Client.Servers = []string{fmt.Sprintf("localhost:%d", serverAgent.Config.Ports.RPC)}
		})
		defer s.Shutdown()

		// No ?type= means server is just skipped
		{
			req, err := http.NewRequest("GET", "/v1/agent/health", nil)
			require.Nil(err)

			respW := httptest.NewRecorder()
			healthI, err := s.Server.HealthRequest(respW, req)
			require.Nil(err)
			require.Equal(http.StatusOK, respW.Code)
			require.NotNil(healthI)
			health := healthI.(*healthResponse)
			require.NotNil(health.Client)
			require.True(health.Client.Ok)
			require.Equal("ok", health.Client.Message)
			require.Nil(health.Server)
		}

		// type=server means server is considered unhealthy
		{
			req, err := http.NewRequest("GET", "/v1/agent/health?type=server", nil)
			require.Nil(err)

			respW := httptest.NewRecorder()
			_, err = s.Server.HealthRequest(respW, req)
			require.NotNil(err)
			httpErr, ok := err.(HTTPCodedError)
			require.True(ok)
			require.Equal(500, httpErr.Code())
			require.Equal(`{"server":{"ok":false,"message":"server not enabled"}}`, err.Error())
		}

	})
}

func TestAgentHostData(t *testing.T) {
	t.Parallel()
	httpTest(t, nil, func(s *agent.TestAgent) {
		rpcRegister(t, s, mock.Job())

		testClient, err := NewTestClient(s)
		require.NoError(t, err)

		result, meta, err := testClient.TODO().TODO(TODOOpts.Ctx())
		require.NoError(t, err)
		require.NotNil(t, result)
		require.NotNil(t, meta)

	})
}
