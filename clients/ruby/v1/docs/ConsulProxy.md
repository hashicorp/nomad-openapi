# NomadClient::ConsulProxy

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **config** | [**Hash&lt;String, AnyType&gt;**](AnyType.md) |  | [optional] |
| **expose** | [**ConsulExposeConfig**](ConsulExposeConfig.md) |  | [optional] |
| **expose_config** | [**ConsulExposeConfig**](ConsulExposeConfig.md) |  | [optional] |
| **local_service_address** | **String** |  | [optional] |
| **local_service_port** | **Integer** |  | [optional] |
| **upstreams** | [**Array&lt;ConsulUpstream&gt;**](ConsulUpstream.md) |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::ConsulProxy.new(
  config: null,
  expose: null,
  expose_config: null,
  local_service_address: null,
  local_service_port: null,
  upstreams: null
)
```

