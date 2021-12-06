# NomadClient::NodeResources

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **cpu** | [**NodeCpuResources**](NodeCpuResources.md) |  | [optional] |
| **devices** | [**Array&lt;NodeDeviceResource&gt;**](NodeDeviceResource.md) |  | [optional] |
| **disk** | [**NodeDiskResources**](NodeDiskResources.md) |  | [optional] |
| **memory** | [**NodeMemoryResources**](NodeMemoryResources.md) |  | [optional] |
| **networks** | [**Array&lt;NetworkResource&gt;**](NetworkResource.md) |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::NodeResources.new(
  cpu: null,
  devices: null,
  disk: null,
  memory: null,
  networks: null
)
```

