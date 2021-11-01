package v1

import (
	"testing"

	"github.com/hashicorp/nomad/command/agent"
	"github.com/hashicorp/nomad/nomad/mock"
	"github.com/hashicorp/nomad/nomad/structs"
	"github.com/stretchr/testify/require"
)

func TestGetAllocations(t *testing.T) {
	httpTest(t, nil, func(s *agent.TestAgent) {

		// Directly manipulate the state
		state := s.Agent.Server().State()
		alloc1 := mock.Alloc()
		testEvent := structs.NewTaskEvent(structs.TaskSiblingFailed)
		var events1 []*structs.TaskEvent
		events1 = append(events1, testEvent)
		taskState := &structs.TaskState{Events: events1}
		alloc1.TaskStates = make(map[string]*structs.TaskState)
		alloc1.TaskStates["test"] = taskState

		alloc2 := mock.Alloc()
		alloc2.TaskStates = make(map[string]*structs.TaskState)
		alloc2.TaskStates["test"] = taskState

		state.UpsertJobSummary(998, mock.JobSummary(alloc1.JobID))
		state.UpsertJobSummary(999, mock.JobSummary(alloc2.JobID))
		err := state.UpsertAllocs(structs.MsgTypeTestSetup, 1000, []*structs.Allocation{alloc1, alloc2})
		if err != nil {
			t.Fatalf("err: %v", err)
		}

		// Make the HTTP request
		testClient, err := NewTestClient(s)
		require.NoError(t, err)

		result, meta, err := testClient.Allocations().GetAllocations(queryOpts.Ctx())
		require.NoError(t, err)
		require.NotNil(t, result)
		require.NotNil(t, meta)

		require.Len(t, *result, 2)

		expectedMsg := "Task's sibling failed"
		allocList := *result
		require.NotNil(t, allocList)
		taskStates := *allocList[0].TaskStates
		require.NotNil(t, taskStates)
		events := *taskStates["test"].Events
		require.NotNil(t, events)
		displayMsg1 := *events[0].DisplayMessage
		require.Equal(t, expectedMsg, displayMsg1, "DisplayMessage should be set")
	})
}
