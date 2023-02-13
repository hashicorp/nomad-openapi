// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v1

import (
	"errors"
	"fmt"
	"path"
	"reflect"
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
	return NewClient(WithAddress(testAgent.HTTPAddr()))
}
func Test_MakeAPIError(t *testing.T) {
	testcases := []struct {
		name              string
		input             interface{}
		expectErrContains string
	}{
		{
			name:              "good error",
			input:             errors.New("test error"),
			expectErrContains: "test error",
		},
		{
			name:              "string",
			input:             "foo",
			expectErrContains: "Received non-error value",
		},
		{
			name:              "empty string",
			input:             "",
			expectErrContains: "Received non-error value",
		},
		{
			name:              "string slice",
			input:             []string{"a", "b"},
			expectErrContains: "Received non-error value",
		},
		{
			name:              "nil",
			input:             nil,
			expectErrContains: "Received invalid value",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			v := reflect.ValueOf(tc.input)
			apiErr := MakeAPIError([]reflect.Value{v})
			require.Error(t, apiErr)
			require.Contains(t, apiErr.Error(), tc.expectErrContains)
		})
	}
}

func TestSetQueryOptions(t *testing.T) {
	ctx := queryOpts.Ctx()
	qCtx := ctx.Value(contextKeyQueryOpts).(*QueryOpts)

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
	wCtx := ctx.Value(contextKeyWriteOpts).(*WriteOpts)

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
		APITestCase{"ConfigFromArgs", testConfigFromArgs},
	)
}

func testConfigFromEnv(t *testing.T, s *agent.TestAgent) {
	t.Setenv(EnvNomadCACert, mTLSFixturePath("client", "cafile"))
	t.Setenv(EnvNomadClientCert, mTLSFixturePath("client", "certfile"))
	t.Setenv(EnvNomadClientKey, mTLSFixturePath("client", "keyfile"))
	t.Setenv(EnvNomadAddr, s.HTTPAddr())
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

func testConfigFromArgs(t *testing.T, s *agent.TestAgent) {
	client, err := NewClient(
		WithTLSCerts(
			mTLSFixturePath("client", "cafile"),
			mTLSFixturePath("client", "certfile"),
			mTLSFixturePath("client", "keyfile"),
		),
		WithAddress(s.HTTPAddr()),
	)

	require.NoError(t, err)

	q := &QueryOpts{
		Region:    globalRegion,
		Namespace: defaultNamespace,
	}
	result, err := client.Status().Leader(q.Ctx())
	t.Logf("result: %q", *result)
	require.NoError(t, err)
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
