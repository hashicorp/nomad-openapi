# NomadClient::SchedulerConfiguration

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **create_index** | **Integer** |  | [optional] |
| **memory_oversubscription_enabled** | **Boolean** |  | [optional] |
| **modify_index** | **Integer** |  | [optional] |
| **preemption_config** | [**PreemptionConfig**](PreemptionConfig.md) |  | [optional] |
| **scheduler_algorithm** | **String** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::SchedulerConfiguration.new(
  create_index: null,
  memory_oversubscription_enabled: null,
  modify_index: null,
  preemption_config: null,
  scheduler_algorithm: null
)
```

