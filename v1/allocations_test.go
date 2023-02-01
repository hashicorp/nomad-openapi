// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v1

import (
	"testing"

	"github.com/hashicorp/nomad/command/agent"
	"github.com/hashicorp/nomad/nomad/mock"
	"github.com/hashicorp/nomad/nomad/structs"
	"github.com/stretchr/testify/require"
)

func TestAllocations(t *testing.T) {
	httpTests(t, nil,
		APITestCase{"GetAllocations", testGetAllocations},
		APITestCase{"GetAllocation", testGetAllocation},
		APITestCase{"StopAllocation", testStopAllocation},
		APITestCase{"GetAllocationServices", testGetAllocationServices},
	)
}

func testGetAllocations(t *testing.T, s *agent.TestAgent) {

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

	_ = state.UpsertJobSummary(998, mock.JobSummary(alloc1.JobID))
	_ = state.UpsertJobSummary(999, mock.JobSummary(alloc2.JobID))
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
}

func testGetAllocation(t *testing.T, s *agent.TestAgent) {

	// Directly manipulate the state
	state := s.Agent.Server().State()
	alloc1 := mock.Alloc()

	err := state.UpsertAllocs(structs.MsgTypeTestSetup, 1000, []*structs.Allocation{alloc1})
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	// Make the HTTP request
	testClient, err := NewTestClient(s)
	require.NoError(t, err)

	result, meta, err := testClient.Allocations().GetAllocation(queryOpts.Ctx(), alloc1.ID)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, meta)

	alloc := *result
	require.Equal(t, alloc1.DesiredStatus, *alloc.DesiredStatus)
}

func testStopAllocation(t *testing.T, s *agent.TestAgent) {

	// Directly manipulate the state
	state := s.Agent.Server().State()
	alloc1 := mock.Alloc()

	err := state.UpsertAllocs(structs.MsgTypeTestSetup, 1000, []*structs.Allocation{alloc1})
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	// Make the HTTP request
	testClient, err := NewTestClient(s)
	require.NoError(t, err)

	result, meta, err := testClient.Allocations().StopAllocation(queryOpts.Ctx(), alloc1.ID, false)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, meta)

	stopResult := *result
	require.NotNil(t, stopResult.EvalID)
}

func testGetAllocationServices(t *testing.T, s *agent.TestAgent) {

	// Directly manipulate the state
	state := s.Agent.Server().State()
	alloc1 := mock.Alloc()

	err := state.UpsertAllocs(structs.MsgTypeTestSetup, 1000, []*structs.Allocation{alloc1})
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	serviceRegistrations := mock.ServiceRegistrations()
	for _, registration := range serviceRegistrations {
		registration.AllocID = alloc1.ID
	}

	err = state.UpsertServiceRegistrations(structs.MsgTypeTestSetup, 2000, serviceRegistrations)
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	// Make the HTTP request
	testClient, err := NewTestClient(s)
	require.NoError(t, err)

	result, meta, err := testClient.Allocations().GetAllocationServices(queryOpts.Ctx(), alloc1.ID)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, meta)

	services := *result
	require.NotNil(t, services)
	require.Len(t, services, 2)
}
