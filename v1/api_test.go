package v1

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/hashicorp/nomad/command/agent"
	"github.com/hashicorp/nomad/testutil"
	"github.com/stretchr/testify/require"
)

var queryOpts = DefaultQueryOpts().
	WithAllowStale(true).
	WithWaitIndex(1000).
	WithWaitTime(100 * time.Second).
	WithAuthToken("foobar")

var writeOpts = DefaultWriteOpts().
	WithAuthToken("foobar")

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

func NewTestClient(testAgent *agent.TestAgent) (*Client, error) {
	os.Setenv("NOMAD_ADDR", fmt.Sprintf("http://%s:%d", testAgent.Config.BindAddr, testAgent.Config.Ports.HTTP))
	defer os.Setenv("NOMAD_ADDR", "http://127.0.0.1:4646")

	return NewClient()
}

func TestSetQueryOptions(t *testing.T) {
	httpTest(t, nil, func(s *agent.TestAgent) {

		ctx := queryOpts.Ctx()
		qCtx := ctx.Value("QueryOpts").(*QueryOpts)

		require.Equal(t, qCtx.Region, queryOpts.Region)
		require.Equal(t, qCtx.Namespace, queryOpts.Namespace)
		require.Equal(t, qCtx.AllowStale, queryOpts.AllowStale)
		require.Equal(t, qCtx.WaitIndex, queryOpts.WaitIndex)
		require.Equal(t, qCtx.WaitTime, queryOpts.WaitTime)
		require.Equal(t, qCtx.AuthToken, queryOpts.AuthToken)
		require.Equal(t, qCtx.PerPage, queryOpts.PerPage)
		require.Equal(t, qCtx.NextToken, queryOpts.NextToken)
		require.Equal(t, qCtx.Prefix, queryOpts.Prefix)
	})
}

func TestSetWriteOptions(t *testing.T) {
	httpTest(t, nil, func(s *agent.TestAgent) {
		ctx := writeOpts.Ctx()
		wCtx := ctx.Value("WriteOpts").(*WriteOpts)

		require.Equal(t, wCtx.Region, writeOpts.Region)
		require.Equal(t, wCtx.Namespace, writeOpts.Namespace)
		require.Equal(t, wCtx.AuthToken, writeOpts.AuthToken)
		require.Equal(t, wCtx.IdempotencyToken, writeOpts.IdempotencyToken)
	})
}
