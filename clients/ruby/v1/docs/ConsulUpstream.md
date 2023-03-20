# NomadClient::ConsulUpstream

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **config** | [**Hash&lt;String, AnyType&gt;**](AnyType.md) |  | [optional] |
| **datacenter** | **String** |  | [optional] |
| **destination_name** | **String** |  | [optional] |
| **destination_namespace** | **String** |  | [optional] |
| **local_bind_address** | **String** |  | [optional] |
| **local_bind_port** | **Integer** |  | [optional] |
| **mesh_gateway** | [**ConsulMeshGateway**](ConsulMeshGateway.md) |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::ConsulUpstream.new(
  config: null,
  datacenter: null,
  destination_name: null,
  destination_namespace: null,
  local_bind_address: null,
  local_bind_port: null,
  mesh_gateway: null
)
```

