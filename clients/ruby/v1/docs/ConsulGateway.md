# NomadClient::ConsulGateway

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **ingress** | [**ConsulIngressConfigEntry**](ConsulIngressConfigEntry.md) |  | [optional] |
| **mesh** | **Object** |  | [optional] |
| **proxy** | [**ConsulGatewayProxy**](ConsulGatewayProxy.md) |  | [optional] |
| **terminating** | [**ConsulTerminatingConfigEntry**](ConsulTerminatingConfigEntry.md) |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::ConsulGateway.new(
  ingress: null,
  mesh: null,
  proxy: null,
  terminating: null
)
```

