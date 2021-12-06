# NomadClient::DeploymentState

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **auto_revert** | **Boolean** |  | [optional] |
| **desired_canaries** | **Integer** |  | [optional] |
| **desired_total** | **Integer** |  | [optional] |
| **healthy_allocs** | **Integer** |  | [optional] |
| **placed_allocs** | **Integer** |  | [optional] |
| **placed_canaries** | **Array&lt;String&gt;** |  | [optional] |
| **progress_deadline** | **Integer** |  | [optional] |
| **promoted** | **Boolean** |  | [optional] |
| **require_progress_by** | **Time** |  | [optional] |
| **unhealthy_allocs** | **Integer** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::DeploymentState.new(
  auto_revert: null,
  desired_canaries: null,
  desired_total: null,
  healthy_allocs: null,
  placed_allocs: null,
  placed_canaries: null,
  progress_deadline: null,
  promoted: null,
  require_progress_by: null,
  unhealthy_allocs: null
)
```

