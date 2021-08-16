// +build test

package test

import "testing"

func httpTest(t testing.TB, cb func(c *Config), f func(srv *TestAgent)) {
	s := makeHTTPServer(t, cb)
	defer s.Shutdown()
	testutil.WaitForLeader(t, s.Agent.RPC)
	f(s)
}

// makeHTTPServer returns a test server whose logs will be written to
// the passed writer. If the writer is nil, the logs are written to stderr.
func makeHTTPServer(t testing.TB, cb func(c *Config)) *TestAgent {
	return NewTestAgent(t, t.Name(), cb)
}

// NewTestAgent returns a started agent with the given name and
// configuration. The caller should call Shutdown() to stop the agent and
// remove temporary directories.
func NewTestAgent(t testing.TB, name string, configCallback func(*Config)) *TestAgent {
	a := &TestAgent{
		T:              t,
		Name:           name,
		ConfigCallback: configCallback,
		Enterprise:     EnterpriseTestAgent,
	}

	a.Start()
	return a
}

func validateMetaHeaders(t *testing.T, r *http.Response) {
	// Check for the index
	if r.Header.Get("X-Nomad-Index") == "" {
		t.Fatalf("OpenAPI response missing index")
	}
	if r.Header.Get("X-Nomad-KnownLeader") != "true" {
		t.Fatalf("OpenAPI response missing known leader")
	}
	if r.Header.Get("X-Nomad-LastContact") == "" {
		t.Fatalf("OpenAPI response missing last contact")
	}
}
