# NomadClient::Allocation

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **alloc_modify_index** | **Integer** |  | [optional] |
| **allocated_resources** | [**AllocatedResources**](AllocatedResources.md) |  | [optional] |
| **client_description** | **String** |  | [optional] |
| **client_status** | **String** |  | [optional] |
| **create_index** | **Integer** |  | [optional] |
| **create_time** | **Integer** |  | [optional] |
| **deployment_id** | **String** |  | [optional] |
| **deployment_status** | [**AllocDeploymentStatus**](AllocDeploymentStatus.md) |  | [optional] |
| **desired_description** | **String** |  | [optional] |
| **desired_status** | **String** |  | [optional] |
| **desired_transition** | [**DesiredTransition**](DesiredTransition.md) |  | [optional] |
| **eval_id** | **String** |  | [optional] |
| **followup_eval_id** | **String** |  | [optional] |
| **id** | **String** |  | [optional] |
| **job** | [**Job**](Job.md) |  | [optional] |
| **job_id** | **String** |  | [optional] |
| **metrics** | [**AllocationMetric**](AllocationMetric.md) |  | [optional] |
| **modify_index** | **Integer** |  | [optional] |
| **modify_time** | **Integer** |  | [optional] |
| **name** | **String** |  | [optional] |
| **namespace** | **String** |  | [optional] |
| **next_allocation** | **String** |  | [optional] |
| **node_id** | **String** |  | [optional] |
| **node_name** | **String** |  | [optional] |
| **preempted_allocations** | **Array&lt;String&gt;** |  | [optional] |
| **preempted_by_allocation** | **String** |  | [optional] |
| **previous_allocation** | **String** |  | [optional] |
| **reschedule_tracker** | [**RescheduleTracker**](RescheduleTracker.md) |  | [optional] |
| **resources** | [**Resources**](Resources.md) |  | [optional] |
| **services** | **Hash&lt;String, String&gt;** |  | [optional] |
| **task_group** | **String** |  | [optional] |
| **task_resources** | [**Hash&lt;String, Resources&gt;**](Resources.md) |  | [optional] |
| **task_states** | [**Hash&lt;String, TaskState&gt;**](TaskState.md) |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::Allocation.new(
  alloc_modify_index: null,
  allocated_resources: null,
  client_description: null,
  client_status: null,
  create_index: null,
  create_time: null,
  deployment_id: null,
  deployment_status: null,
  desired_description: null,
  desired_status: null,
  desired_transition: null,
  eval_id: null,
  followup_eval_id: null,
  id: null,
  job: null,
  job_id: null,
  metrics: null,
  modify_index: null,
  modify_time: null,
  name: null,
  namespace: null,
  next_allocation: null,
  node_id: null,
  node_name: null,
  preempted_allocations: null,
  preempted_by_allocation: null,
  previous_allocation: null,
  reschedule_tracker: null,
  resources: null,
  services: null,
  task_group: null,
  task_resources: null,
  task_states: null
)
```

