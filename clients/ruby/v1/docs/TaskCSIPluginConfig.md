# NomadClient::TaskCSIPluginConfig

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **health_timeout** | **Integer** |  | [optional] |
| **id** | **String** |  | [optional] |
| **mount_dir** | **String** |  | [optional] |
| **type** | **String** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::TaskCSIPluginConfig.new(
  health_timeout: null,
  id: null,
  mount_dir: null,
  type: null
)
```

