# NomadClient::GaugeValue

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **labels** | **Hash&lt;String, String&gt;** |  | [optional] |
| **name** | **String** |  | [optional] |
| **value** | **Float** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::GaugeValue.new(
  labels: null,
  name: null,
  value: null
)
```

