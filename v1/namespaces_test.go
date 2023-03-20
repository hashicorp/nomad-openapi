// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v1

import (
	"fmt"
	"testing"

	client "github.com/flytocolors/nomad-openapi/clients/go/v1"
	"github.com/hashicorp/nomad/command/agent"
	"github.com/hashicorp/nomad/helper/uuid"
	"github.com/hashicorp/nomad/nomad/mock"
	"github.com/hashicorp/nomad/nomad/structs"
	"github.com/stretchr/testify/require"
)

func TestNamespaces(t *testing.T) {
	httpTests(t, nil,
		APITestCase{"testGetNamespaces", testGetNamespaces},
		APITestCase{"testGetNamespace", testGetNamespace},
		APITestCase{"testPostNamespace", testPostNamespace},
		APITestCase{"testDeleteNamespace", testDeleteNamespace},
	)
}

func testGetNamespaces(t *testing.T, s *agent.TestAgent) {
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
}

func testGetNamespace(t *testing.T, s *agent.TestAgent) {
	testClient, err := NewTestClient(s)
	require.NoError(t, err)

	result, meta, err := testClient.Namespaces().GetNamespace(queryOpts.Ctx(), defaultNamespace)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, meta)

	require.Equal(t, defaultNamespace, *result.Name)
}

func testPostNamespace(t *testing.T, s *agent.TestAgent) {
	name := fmt.Sprintf("openapi-ns-%s", uuid.Generate())
	clientNS := client.Namespace{
		Name: &name,
	}

	testClient, err := NewTestClient(s)
	require.NoError(t, err)
	meta, err := testClient.Namespaces().PostNamespace(writeOpts.Ctx(), &clientNS)
	require.NoError(t, err)
	require.NotNil(t, meta)

	result, _, err := testClient.Namespaces().GetNamespace(queryOpts.Ctx(), *clientNS.Name)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, result.Name)
	require.Equal(t, *clientNS.Name, *result.Name)
}

func testDeleteNamespace(t *testing.T, s *agent.TestAgent) {
	name := fmt.Sprintf("openapi-ns-%s", uuid.Generate())
	clientNS := client.Namespace{
		Name: &name,
	}

	testClient, err := NewTestClient(s)
	require.NoError(t, err)
	meta, err := testClient.Namespaces().PostNamespace(writeOpts.Ctx(), &clientNS)
	require.NoError(t, err)
	require.NotNil(t, meta)

	result, _, err := testClient.Namespaces().GetNamespace(queryOpts.Ctx(), *clientNS.Name)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, result.Name)
	require.Equal(t, *clientNS.Name, *result.Name)

	_, err = testClient.Namespaces().DeleteNamespace(writeOpts.Ctx(), *clientNS.Name)
	require.NoError(t, err)

	_, _, err = testClient.Namespaces().GetNamespace(queryOpts.Ctx(), *clientNS.Name)
	require.Error(t, err)
}
