// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v1

import (
	"testing"

	"github.com/hashicorp/nomad/command/agent"
	"github.com/stretchr/testify/require"
)

func TestPlugins(t *testing.T) {
	httpTests(t, nil,
		APITestCase{"ListPlugins", testListPlugins},
		APITestCase{"GetCSIPlugin", testGetCSIPlugin},
	)
}

func testListPlugins(t *testing.T, s *agent.TestAgent) {
	testClient, err := NewTestClient(s)
	require.NoError(t, err)

	result, err := testClient.Plugins().Get(queryOpts.Ctx())
	require.NoError(t, err)
	require.Len(t, *result, 0)
	require.NotNil(t, result)
}

func testGetCSIPlugin(t *testing.T, s *agent.TestAgent) {
	testClient, err := NewTestClient(s)
	require.NoError(t, err)

	result, meta, err := testClient.Plugins().GetPlugin(queryOpts.Ctx(), "example-plugin")
	require.Error(t, err, "plugin not found")
	require.Nil(t, meta)
	require.Nil(t, result)
}
