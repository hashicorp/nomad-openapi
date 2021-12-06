# NomadClient::NodeEvent

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **create_index** | **Integer** |  | [optional] |
| **details** | **Hash&lt;String, String&gt;** |  | [optional] |
| **message** | **String** |  | [optional] |
| **subsystem** | **String** |  | [optional] |
| **timestamp** | **Time** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::NodeEvent.new(
  create_index: null,
  details: null,
  message: null,
  subsystem: null,
  timestamp: null
)
```

