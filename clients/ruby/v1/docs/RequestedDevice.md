# NomadClient::RequestedDevice

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **affinities** | [**Array&lt;Affinity&gt;**](Affinity.md) |  | [optional] |
| **constraints** | [**Array&lt;Constraint&gt;**](Constraint.md) |  | [optional] |
| **count** | **Integer** |  | [optional] |
| **name** | **String** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::RequestedDevice.new(
  affinities: null,
  constraints: null,
  count: null,
  name: null
)
```

