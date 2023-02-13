// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v1

import (
	"testing"

	"github.com/hashicorp/nomad/command/agent"
	"github.com/stretchr/testify/require"
)

func TestScaling(t *testing.T) {
	httpTests(t, nil,
		APITestCase{"GetScalingPolicies", testGetScalingPolicies},
		APITestCase{"GetPolicy", testGetPolicy},
	)
}

func testGetScalingPolicies(t *testing.T, s *agent.TestAgent) {
	testClient, err := NewTestClient(s)
	require.NoError(t, err)

	result, meta, err := testClient.Scaling().Policies(queryOpts.Ctx())
	require.NoError(t, err)
	require.NotNil(t, meta)
	require.Len(t, *result, 0)
	require.NotNil(t, result)
}

func testGetPolicy(t *testing.T, s *agent.TestAgent) {
	testClient, err := NewTestClient(s)
	require.NoError(t, err)

	result, meta, err := testClient.Scaling().GetPolicy(queryOpts.Ctx(), "a1b2c3d4-e5f6-g7h8-i9j1-0k11l12")
	require.Error(t, err, "policy not found")
	require.Nil(t, meta)
	require.Nil(t, result)
}
