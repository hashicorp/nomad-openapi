package v1

import (
	"testing"

	"github.com/hashicorp/nomad/client/allocrunner"
	"github.com/hashicorp/nomad/command/agent"
	"github.com/hashicorp/nomad/nomad/mock"
	"github.com/stretchr/testify/require"
)

func TestGetAllocations(t *testing.T) {
	t.Parallel()
	httpTest(t, nil, func(s *agent.TestAgent) {

		// Directly manipulate the state
		alloc := mock.Alloc()
		alloc.Job.ID = "foo"
		ar1, cleanup1 := allocrunner.TestAllocRunnerFromAlloc(t, alloc)
		defer cleanup1()

		alloc.Job.ID = "bar"
		ar2, cleanup2 := allocrunner.TestAllocRunnerFromAlloc(t, alloc)
		defer cleanup2()

		go ar1.Run()
		go ar2.Run()

		// Make the HTTP request
		testClient, err := NewTestClient(s)
		require.NoError(t, err)

		result, meta, err := testClient.Allocations().GetAllocations(queryOpts.Ctx())
		require.NoError(t, err)
		require.NotNil(t, result)
		require.NotNil(t, meta)

		require.Len(t, *result, 2)
	})
}

func TestGetAllocation(t *testing.T) {
	t.Parallel()
	httpTest(t, nil, func(s *agent.TestAgent) {
		alloc := mock.Alloc()
		ar, cleanup := allocrunner.TestAllocRunnerFromAlloc(t, alloc)
		defer cleanup()

		go ar.Run()

		testClient, err := NewTestClient(s)
		require.NoError(t, err)

		result, meta, err := testClient.Allocations().GetAllocation(queryOpts.Ctx(), ar.Alloc().ID)
		require.NoError(t, err)
		require.NotNil(t, result)
		require.NotNil(t, meta)

	})
}

func TestStopAllocation(t *testing.T) {
	t.Parallel()
	httpTest(t, nil, func(s *agent.TestAgent) {
		alloc := mock.Alloc()
		ar, cleanup := allocrunner.TestAllocRunnerFromAlloc(t, alloc)
		defer cleanup()

		go ar.Run()

		testClient, err := NewTestClient(s)
		require.NoError(t, err)

		result, meta, err := testClient.Allocations().Stop(writeOpts.Ctx(), ar.Alloc().ID)
		require.NoError(t, err)
		require.NotNil(t, result)
		require.NotNil(t, meta)

	})
}
