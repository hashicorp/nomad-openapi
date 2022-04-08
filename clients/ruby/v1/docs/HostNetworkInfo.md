# NomadClient::HostNetworkInfo

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **cidr** | **String** |  | [optional] |
| **interface** | **String** |  | [optional] |
| **name** | **String** |  | [optional] |
| **reserved_ports** | **String** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::HostNetworkInfo.new(
  cidr: null,
  interface: null,
  name: null,
  reserved_ports: null
)
```

