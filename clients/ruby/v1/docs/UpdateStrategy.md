# NomadClient::UpdateStrategy

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **auto_promote** | **Boolean** |  | [optional] |
| **auto_revert** | **Boolean** |  | [optional] |
| **canary** | **Integer** |  | [optional] |
| **health_check** | **String** |  | [optional] |
| **healthy_deadline** | **Integer** |  | [optional] |
| **max_parallel** | **Integer** |  | [optional] |
| **min_healthy_time** | **Integer** |  | [optional] |
| **progress_deadline** | **Integer** |  | [optional] |
| **stagger** | **Integer** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::UpdateStrategy.new(
  auto_promote: null,
  auto_revert: null,
  canary: null,
  health_check: null,
  healthy_deadline: null,
  max_parallel: null,
  min_healthy_time: null,
  progress_deadline: null,
  stagger: null
)
```

