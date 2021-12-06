# NomadClient::MetricsSummary

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **counters** | [**Array&lt;SampledValue&gt;**](SampledValue.md) |  | [optional] |
| **gauges** | [**Array&lt;GaugeValue&gt;**](GaugeValue.md) |  | [optional] |
| **points** | [**Array&lt;PointValue&gt;**](PointValue.md) |  | [optional] |
| **samples** | [**Array&lt;SampledValue&gt;**](SampledValue.md) |  | [optional] |
| **timestamp** | **String** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::MetricsSummary.new(
  counters: null,
  gauges: null,
  points: null,
  samples: null,
  timestamp: null
)
```

