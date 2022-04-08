# NomadClient::Evaluation

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **annotate_plan** | **Boolean** |  | [optional] |
| **blocked_eval** | **String** |  | [optional] |
| **class_eligibility** | **Hash&lt;String, Boolean&gt;** |  | [optional] |
| **create_index** | **Integer** |  | [optional] |
| **create_time** | **Integer** |  | [optional] |
| **deployment_id** | **String** |  | [optional] |
| **escaped_computed_class** | **Boolean** |  | [optional] |
| **failed_tg_allocs** | [**Hash&lt;String, AllocationMetric&gt;**](AllocationMetric.md) |  | [optional] |
| **id** | **String** |  | [optional] |
| **job_id** | **String** |  | [optional] |
| **job_modify_index** | **Integer** |  | [optional] |
| **modify_index** | **Integer** |  | [optional] |
| **modify_time** | **Integer** |  | [optional] |
| **namespace** | **String** |  | [optional] |
| **next_eval** | **String** |  | [optional] |
| **node_id** | **String** |  | [optional] |
| **node_modify_index** | **Integer** |  | [optional] |
| **previous_eval** | **String** |  | [optional] |
| **priority** | **Integer** |  | [optional] |
| **queued_allocations** | **Hash&lt;String, Integer&gt;** |  | [optional] |
| **quota_limit_reached** | **String** |  | [optional] |
| **related_evals** | [**Array&lt;EvaluationStub&gt;**](EvaluationStub.md) |  | [optional] |
| **snapshot_index** | **Integer** |  | [optional] |
| **status** | **String** |  | [optional] |
| **status_description** | **String** |  | [optional] |
| **triggered_by** | **String** |  | [optional] |
| **type** | **String** |  | [optional] |
| **wait** | **Integer** |  | [optional] |
| **wait_until** | **Time** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::Evaluation.new(
  annotate_plan: null,
  blocked_eval: null,
  class_eligibility: null,
  create_index: null,
  create_time: null,
  deployment_id: null,
  escaped_computed_class: null,
  failed_tg_allocs: null,
  id: null,
  job_id: null,
  job_modify_index: null,
  modify_index: null,
  modify_time: null,
  namespace: null,
  next_eval: null,
  node_id: null,
  node_modify_index: null,
  previous_eval: null,
  priority: null,
  queued_allocations: null,
  quota_limit_reached: null,
  related_evals: null,
  snapshot_index: null,
  status: null,
  status_description: null,
  triggered_by: null,
  type: null,
  wait: null,
  wait_until: null
)
```

