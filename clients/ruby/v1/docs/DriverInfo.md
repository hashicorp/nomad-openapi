# NomadClient::DriverInfo

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **attributes** | **Hash&lt;String, String&gt;** |  | [optional] |
| **detected** | **Boolean** |  | [optional] |
| **health_description** | **String** |  | [optional] |
| **healthy** | **Boolean** |  | [optional] |
| **update_time** | **Time** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::DriverInfo.new(
  attributes: null,
  detected: null,
  health_description: null,
  healthy: null,
  update_time: null
)
```

