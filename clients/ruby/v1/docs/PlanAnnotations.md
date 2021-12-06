# NomadClient::PlanAnnotations

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **desired_tg_updates** | [**Hash&lt;String, DesiredUpdates&gt;**](DesiredUpdates.md) |  | [optional] |
| **preempted_allocs** | [**Array&lt;AllocationListStub&gt;**](AllocationListStub.md) |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::PlanAnnotations.new(
  desired_tg_updates: null,
  preempted_allocs: null
)
```

