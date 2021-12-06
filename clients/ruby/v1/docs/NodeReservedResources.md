# NomadClient::NodeReservedResources

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **cpu** | [**NodeReservedCpuResources**](NodeReservedCpuResources.md) |  | [optional] |
| **disk** | [**NodeReservedDiskResources**](NodeReservedDiskResources.md) |  | [optional] |
| **memory** | [**NodeReservedMemoryResources**](NodeReservedMemoryResources.md) |  | [optional] |
| **networks** | [**NodeReservedNetworkResources**](NodeReservedNetworkResources.md) |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::NodeReservedResources.new(
  cpu: null,
  disk: null,
  memory: null,
  networks: null
)
```

