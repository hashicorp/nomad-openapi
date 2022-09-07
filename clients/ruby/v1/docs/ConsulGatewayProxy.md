# NomadClient::ConsulGatewayProxy

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **config** | **Hash&lt;String, Object&gt;** |  | [optional] |
| **connect_timeout** | **Integer** |  | [optional] |
| **envoy_dns_discovery_type** | **String** |  | [optional] |
| **envoy_gateway_bind_addresses** | [**Hash&lt;String, ConsulGatewayBindAddress&gt;**](ConsulGatewayBindAddress.md) |  | [optional] |
| **envoy_gateway_bind_tagged_addresses** | **Boolean** |  | [optional] |
| **envoy_gateway_no_default_bind** | **Boolean** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::ConsulGatewayProxy.new(
  config: null,
  connect_timeout: null,
  envoy_dns_discovery_type: null,
  envoy_gateway_bind_addresses: null,
  envoy_gateway_bind_tagged_addresses: null,
  envoy_gateway_no_default_bind: null
)
```

