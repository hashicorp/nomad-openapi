# NomadClient::ConsulSidecarService

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **disable_default_tcp_check** | **Boolean** |  | [optional] |
| **port** | **String** |  | [optional] |
| **proxy** | [**ConsulProxy**](ConsulProxy.md) |  | [optional] |
| **tags** | **Array&lt;String&gt;** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::ConsulSidecarService.new(
  disable_default_tcp_check: null,
  port: null,
  proxy: null,
  tags: null
)
```

