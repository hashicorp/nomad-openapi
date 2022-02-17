package v1

import (
	"fmt"
	"os"
	"path"
	"testing"
	"time"

	"github.com/hashicorp/nomad/command/agent"
	"github.com/hashicorp/nomad/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var queryOpts = DefaultQueryOpts().
	WithAllowStale(true).
	WithWaitIndex(1000).
	WithWaitTime(100 * time.Second)

var writeOpts = DefaultWriteOpts()

// APITestCase is used to make parallelizable test cases against a single TestAgent
// instance to reduce port churn
type APITestCase struct {
	Name string
	Func func(*testing.T, *agent.TestAgent)
}

// makeHTTPServer returns a test server whose logs will be written to
// the passed writer. If the writer is nil, the logs are written to stderr.
func makeHTTPServer(t testing.TB, cb func(c *agent.Config)) *agent.TestAgent {
	quietByDefault := func(c *agent.Config) {
		c.LogLevel = "WARN"
		if cb != nil {
			cb(c)
		}
	}
	return agent.NewTestAgent(t, t.Name(), quietByDefault)
}

func httpTests(t *testing.T, cb func(c *agent.Config), tests ...APITestCase) {
	s := makeHTTPServer(t, cb)
	defer func() {
		err := s.Shutdown()
		assert.NoError(t, err)
	}()

	testutil.WaitForLeader(t, s.Agent.RPC)
	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			tc.Func(t, s)
		})
	}
}

func NewTestClient(testAgent *agent.TestAgent) (*Client, error) {
	os.Setenv("NOMAD_ADDR", fmt.Sprintf("http://%s:%d", testAgent.Config.BindAddr, testAgent.Config.Ports.HTTP))
	defer os.Setenv("NOMAD_ADDR", "http://127.0.0.1:4646")

	return NewClient()
}

func TestSetQueryOptions(t *testing.T) {
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
}

func TestSetWriteOptions(t *testing.T) {
	ctx := writeOpts.Ctx()
	wCtx := ctx.Value("WriteOpts").(*WriteOpts)

	require.Equal(t, wCtx.Region, writeOpts.Region)
	require.Equal(t, wCtx.Namespace, writeOpts.Namespace)
	require.Equal(t, wCtx.AuthToken, writeOpts.AuthToken)
	require.Equal(t, wCtx.IdempotencyToken, writeOpts.IdempotencyToken)
}

func TestTLSEnabled(t *testing.T) {
	enableTLS := func(c *agent.Config) {
		tC := c.TLSConfig
		tC.VerifyHTTPSClient = true
		tC.EnableHTTP = true
		tC.CAFile = mTLSFixturePath("server", "cafile")
		tC.CertFile = mTLSFixturePath("server", "certfile")
		tC.KeyFile = mTLSFixturePath("server", "keyfile")
	}
	httpTests(t, enableTLS,
		APITestCase{"ConfigFromEnv", testConfigFromEnv},
	)
}

func testConfigFromEnv(t *testing.T, s *agent.TestAgent) {
	t.Setenv("NOMAD_CACERT", mTLSFixturePath("client", "cafile"))
	t.Setenv("NOMAD_CLIENT_CERT", mTLSFixturePath("client", "certfile"))
	t.Setenv("NOMAD_CLIENT_KEY", mTLSFixturePath("client", "keyfile"))
	t.Setenv("NOMAD_ADDR", s.HTTPAddr())
	client, err := NewClient()
	require.NoError(t, err)

	q := &QueryOpts{
		Region:    globalRegion,
		Namespace: defaultNamespace,
	}
	result, err := client.Status().Leader(q.Ctx())
	require.NoError(t, err)
	t.Logf("result: %q\n", *result)
	require.NotNil(t, result)
}

func mTLSFixturePath(nodeType, pemType string) string {
	var filename string
	switch pemType {
	case "cafile":
		filename = "nomad-agent-ca.pem"
	case "certfile":
		filename = fmt.Sprintf("global-%s-nomad-0.pem", nodeType)
	case "keyfile":
		filename = fmt.Sprintf("global-%s-nomad-0-key.pem", nodeType)
	}

	return path.Join("../test_fixtures/mtls", filename)
}
