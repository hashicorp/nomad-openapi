package v1

import (
	"testing"

	"github.com/hashicorp/nomad/command/agent"
	"github.com/hashicorp/nomad/nomad/mock"
	"github.com/hashicorp/nomad/nomad/structs"
	"github.com/stretchr/testify/require"
)

func TestGetNamespaces(t *testing.T) {
	t.Parallel()
	httpTest(t, nil, func(s *agent.TestAgent) {
		ns1 := mock.Namespace()
		ns2 := mock.Namespace()
		ns3 := mock.Namespace()
		args := structs.NamespaceUpsertRequest{
			Namespaces:   []*structs.Namespace{ns1, ns2, ns3},
			WriteRequest: structs.WriteRequest{Region: "global"},
		}
		var resp structs.GenericResponse
		err := s.Agent.RPC("Namespace.UpsertNamespaces", &args, &resp)
		require.NoError(t, err)

		// Make the HTTP request
		testClient, err := NewTestClient(s)
		require.NoError(t, err)

		result, meta, err := testClient.Namespaces().GetNamespaces(queryOpts.Ctx())
		require.NoError(t, err)
		require.NotNil(t, meta)

		// Check the output (the 3 we register + default)
		require.Len(t, *result, 4)
	})
}
