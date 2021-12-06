# NomadClient::ConsulExposePath

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **listener_port** | **String** |  | [optional] |
| **local_path_port** | **Integer** |  | [optional] |
| **path** | **String** |  | [optional] |
| **protocol** | **String** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::ConsulExposePath.new(
  listener_port: null,
  local_path_port: null,
  path: null,
  protocol: null
)
```

