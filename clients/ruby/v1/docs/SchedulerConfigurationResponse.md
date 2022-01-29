# NomadClient::SchedulerConfigurationResponse

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **known_leader** | **Boolean** |  | [optional] |
| **last_contact** | **Integer** |  | [optional] |
| **last_index** | **Integer** |  | [optional] |
| **request_time** | **Integer** |  | [optional] |
| **scheduler_config** | [**SchedulerConfiguration**](SchedulerConfiguration.md) |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::SchedulerConfigurationResponse.new(
  known_leader: null,
  last_contact: null,
  last_index: null,
  request_time: null,
  scheduler_config: null
)
```

