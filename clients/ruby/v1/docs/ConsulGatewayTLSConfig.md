# NomadClient::ConsulGatewayTLSConfig

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **cipher_suites** | **Array&lt;String&gt;** |  | [optional] |
| **enabled** | **Boolean** |  | [optional] |
| **tls_max_version** | **String** |  | [optional] |
| **tls_min_version** | **String** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::ConsulGatewayTLSConfig.new(
  cipher_suites: null,
  enabled: null,
  tls_max_version: null,
  tls_min_version: null
)
```

