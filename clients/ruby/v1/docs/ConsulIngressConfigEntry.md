# NomadClient::ConsulIngressConfigEntry

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **listeners** | [**Array&lt;ConsulIngressListener&gt;**](ConsulIngressListener.md) |  | [optional] |
| **tls** | [**ConsulGatewayTLSConfig**](ConsulGatewayTLSConfig.md) |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::ConsulIngressConfigEntry.new(
  listeners: null,
  tls: null
)
```

