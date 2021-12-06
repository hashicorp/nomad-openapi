# NomadClient::JobPlanResponse

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **annotations** | [**PlanAnnotations**](PlanAnnotations.md) |  | [optional] |
| **created_evals** | [**Array&lt;Evaluation&gt;**](Evaluation.md) |  | [optional] |
| **diff** | [**JobDiff**](JobDiff.md) |  | [optional] |
| **failed_tg_allocs** | [**Hash&lt;String, AllocationMetric&gt;**](AllocationMetric.md) |  | [optional] |
| **job_modify_index** | **Integer** |  | [optional] |
| **next_periodic_launch** | **Time** |  | [optional] |
| **warnings** | **String** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::JobPlanResponse.new(
  annotations: null,
  created_evals: null,
  diff: null,
  failed_tg_allocs: null,
  job_modify_index: null,
  next_periodic_launch: null,
  warnings: null
)
```

