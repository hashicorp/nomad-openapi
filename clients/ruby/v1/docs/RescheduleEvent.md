# NomadClient::RescheduleEvent

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **prev_alloc_id** | **String** |  | [optional] |
| **prev_node_id** | **String** |  | [optional] |
| **reschedule_time** | **Integer** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::RescheduleEvent.new(
  prev_alloc_id: null,
  prev_node_id: null,
  reschedule_time: null
)
```

