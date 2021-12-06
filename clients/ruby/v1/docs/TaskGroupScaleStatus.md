# NomadClient::TaskGroupScaleStatus

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **desired** | **Integer** |  | [optional] |
| **events** | [**Array&lt;ScalingEvent&gt;**](ScalingEvent.md) |  | [optional] |
| **healthy** | **Integer** |  | [optional] |
| **placed** | **Integer** |  | [optional] |
| **running** | **Integer** |  | [optional] |
| **unhealthy** | **Integer** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::TaskGroupScaleStatus.new(
  desired: null,
  events: null,
  healthy: null,
  placed: null,
  running: null,
  unhealthy: null
)
```

