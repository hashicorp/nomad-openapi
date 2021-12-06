# NomadClient::NodeDevice

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **health_description** | **String** |  | [optional] |
| **healthy** | **Boolean** |  | [optional] |
| **id** | **String** |  | [optional] |
| **locality** | [**NodeDeviceLocality**](NodeDeviceLocality.md) |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::NodeDevice.new(
  health_description: null,
  healthy: null,
  id: null,
  locality: null
)
```

