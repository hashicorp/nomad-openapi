# NomadClient::MigrateStrategy

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **health_check** | **String** |  | [optional] |
| **healthy_deadline** | **Integer** |  | [optional] |
| **max_parallel** | **Integer** |  | [optional] |
| **min_healthy_time** | **Integer** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::MigrateStrategy.new(
  health_check: null,
  healthy_deadline: null,
  max_parallel: null,
  min_healthy_time: null
)
```

