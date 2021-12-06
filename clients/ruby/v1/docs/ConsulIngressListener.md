# NomadClient::ConsulIngressListener

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **port** | **Integer** |  | [optional] |
| **protocol** | **String** |  | [optional] |
| **services** | [**Array&lt;ConsulIngressService&gt;**](ConsulIngressService.md) |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::ConsulIngressListener.new(
  port: null,
  protocol: null,
  services: null
)
```

