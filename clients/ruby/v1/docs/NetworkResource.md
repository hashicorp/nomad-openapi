# NomadClient::NetworkResource

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **cidr** | **String** |  | [optional] |
| **dns** | [**DNSConfig**](DNSConfig.md) |  | [optional] |
| **device** | **String** |  | [optional] |
| **dynamic_ports** | [**Array&lt;Port&gt;**](Port.md) |  | [optional] |
| **ip** | **String** |  | [optional] |
| **m_bits** | **Integer** |  | [optional] |
| **mode** | **String** |  | [optional] |
| **reserved_ports** | [**Array&lt;Port&gt;**](Port.md) |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::NetworkResource.new(
  cidr: null,
  dns: null,
  device: null,
  dynamic_ports: null,
  ip: null,
  m_bits: null,
  mode: null,
  reserved_ports: null
)
```

