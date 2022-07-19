# NomadClient::SchedulerConfiguration

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **create_index** | **Integer** |  | [optional] |
| **memory_oversubscription_enabled** | **Boolean** |  | [optional] |
| **modify_index** | **Integer** |  | [optional] |
| **pause_eval_broker** | **Boolean** |  | [optional] |
| **preemption_config** | [**PreemptionConfig**](PreemptionConfig.md) |  | [optional] |
| **reject_job_registration** | **Boolean** |  | [optional] |
| **scheduler_algorithm** | **String** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::SchedulerConfiguration.new(
  create_index: null,
  memory_oversubscription_enabled: null,
  modify_index: null,
  pause_eval_broker: null,
  preemption_config: null,
  reject_job_registration: null,
  scheduler_algorithm: null
)
```

