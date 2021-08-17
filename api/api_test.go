package api

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/hashicorp/nomad/command/agent"
	"github.com/hashicorp/nomad/testutil"
	"github.com/stretchr/testify/require"
)

var queryOpts = &QueryOpts{
	Region:     "foo",
	Namespace:  "bar",
	AllowStale: true,
	WaitIndex:  1000,
	WaitTime:   100 * time.Second,
	AuthToken:  "foobar",
}

var writeOpts = &WriteOpts{
	Region:           "foo",
	Namespace:        "bar",
	AuthToken:        "foobar",
	IdempotencyToken: "",
}

// makeHTTPServer returns a test server whose logs will be written to
// the passed writer. If the writer is nil, the logs are written to stderr.
func makeHTTPServer(t testing.TB, cb func(c *agent.Config)) *agent.TestAgent {
	return agent.NewTestAgent(t, t.Name(), cb)
}

func httpTest(t testing.TB, cb func(c *agent.Config), f func(srv *agent.TestAgent)) {
	s := makeHTTPServer(t, cb)
	defer s.Shutdown()
	testutil.WaitForLeader(t, s.Agent.RPC)
	f(s)
}

func NewTestQueryClient(testAgent *agent.TestAgent, opts *QueryOpts) (*Client, error) {
	os.Setenv("NOMAD_ADDR", fmt.Sprintf("http://%s:%d", testAgent.Config.BindAddr, testAgent.Config.Ports.HTTP))
	defer os.Setenv("NOMAD_ADDR", "http://127.0.0.1:4646")

	return NewQueryClient(opts)
}

func NewTestWriteClient(testAgent *agent.TestAgent, opts *WriteOpts) (*Client, error) {
	os.Setenv("NOMAD_ADDR", fmt.Sprintf("http://%s:%d", testAgent.Config.BindAddr, testAgent.Config.Ports.HTTP))
	defer os.Setenv("NOMAD_ADDR", "http://127.0.0.1:4646")

	return NewWriteClient(opts)
}

func TestSetQueryOptions(t *testing.T) {
	t.Parallel()
	httpTest(t, nil, func(s *agent.TestAgent) {
		client, err := NewTestQueryClient(s, queryOpts)
		require.NoError(t, err)
		cfg := *client.config

		require.Equal(t, cfg.Region, queryOpts.Region)
		require.Equal(t, cfg.Namespace, queryOpts.Namespace)
		require.Equal(t, cfg.QueryOpts.AllowStale, true)
		require.Equal(t, cfg.QueryOpts.WaitIndex, queryOpts.WaitIndex)
		require.Equal(t, cfg.QueryOpts.WaitTime, queryOpts.WaitTime)
		require.Equal(t, cfg.QueryOpts.AuthToken, queryOpts.AuthToken)
	})
}

func TestSetWriteOptions(t *testing.T) {
	t.Parallel()
	httpTest(t, nil, func(s *agent.TestAgent) {
		client, err := NewTestWriteClient(s, writeOpts)
		require.NoError(t, err)
		cfg := *client.config

		require.Equal(t, cfg.Region, writeOpts.Region)
		require.Equal(t, cfg.Namespace, writeOpts.Namespace)
		require.Equal(t, cfg.WriteOpts.AuthToken, writeOpts.AuthToken)
		require.Equal(t, cfg.WriteOpts.IdempotencyToken, writeOpts.IdempotencyToken)
	})
}
