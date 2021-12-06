# NomadClient::SampledValue

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **count** | **Integer** |  | [optional] |
| **labels** | **Hash&lt;String, String&gt;** |  | [optional] |
| **max** | **Float** |  | [optional] |
| **mean** | **Float** |  | [optional] |
| **min** | **Float** |  | [optional] |
| **name** | **String** |  | [optional] |
| **rate** | **Float** |  | [optional] |
| **stddev** | **Float** |  | [optional] |
| **sum** | **Float** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::SampledValue.new(
  count: null,
  labels: null,
  max: null,
  mean: null,
  min: null,
  name: null,
  rate: null,
  stddev: null,
  sum: null
)
```

