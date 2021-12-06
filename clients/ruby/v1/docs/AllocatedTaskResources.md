# NomadClient::AllocatedTaskResources

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **cpu** | [**AllocatedCpuResources**](AllocatedCpuResources.md) |  | [optional] |
| **devices** | [**Array&lt;AllocatedDeviceResource&gt;**](AllocatedDeviceResource.md) |  | [optional] |
| **memory** | [**AllocatedMemoryResources**](AllocatedMemoryResources.md) |  | [optional] |
| **networks** | [**Array&lt;NetworkResource&gt;**](NetworkResource.md) |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::AllocatedTaskResources.new(
  cpu: null,
  devices: null,
  memory: null,
  networks: null
)
```

