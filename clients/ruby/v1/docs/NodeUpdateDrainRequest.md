# NomadClient::NodeUpdateDrainRequest

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **drain_spec** | [**DrainSpec**](DrainSpec.md) |  | [optional] |
| **mark_eligible** | **Boolean** |  | [optional] |
| **meta** | **Hash&lt;String, String&gt;** |  | [optional] |
| **node_id** | **String** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::NodeUpdateDrainRequest.new(
  drain_spec: null,
  mark_eligible: null,
  meta: null,
  node_id: null
)
```

