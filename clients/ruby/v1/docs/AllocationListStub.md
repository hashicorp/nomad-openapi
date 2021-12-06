# NomadClient::AllocationListStub

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **allocated_resources** | [**AllocatedResources**](AllocatedResources.md) |  | [optional] |
| **client_description** | **String** |  | [optional] |
| **client_status** | **String** |  | [optional] |
| **create_index** | **Integer** |  | [optional] |
| **create_time** | **Integer** |  | [optional] |
| **deployment_status** | [**AllocDeploymentStatus**](AllocDeploymentStatus.md) |  | [optional] |
| **desired_description** | **String** |  | [optional] |
| **desired_status** | **String** |  | [optional] |
| **eval_id** | **String** |  | [optional] |
| **followup_eval_id** | **String** |  | [optional] |
| **id** | **String** |  | [optional] |
| **job_id** | **String** |  | [optional] |
| **job_type** | **String** |  | [optional] |
| **job_version** | **Integer** |  | [optional] |
| **modify_index** | **Integer** |  | [optional] |
| **modify_time** | **Integer** |  | [optional] |
| **name** | **String** |  | [optional] |
| **namespace** | **String** |  | [optional] |
| **node_id** | **String** |  | [optional] |
| **node_name** | **String** |  | [optional] |
| **preempted_allocations** | **Array&lt;String&gt;** |  | [optional] |
| **preempted_by_allocation** | **String** |  | [optional] |
| **reschedule_tracker** | [**RescheduleTracker**](RescheduleTracker.md) |  | [optional] |
| **task_group** | **String** |  | [optional] |
| **task_states** | [**Hash&lt;String, TaskState&gt;**](TaskState.md) |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::AllocationListStub.new(
  allocated_resources: null,
  client_description: null,
  client_status: null,
  create_index: null,
  create_time: null,
  deployment_status: null,
  desired_description: null,
  desired_status: null,
  eval_id: null,
  followup_eval_id: null,
  id: null,
  job_id: null,
  job_type: null,
  job_version: null,
  modify_index: null,
  modify_time: null,
  name: null,
  namespace: null,
  node_id: null,
  node_name: null,
  preempted_allocations: null,
  preempted_by_allocation: null,
  reschedule_tracker: null,
  task_group: null,
  task_states: null
)
```

