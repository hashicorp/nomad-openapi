# NomadClient::NodeResources

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **cpu** | [**NodeCpuResources**](NodeCpuResources.md) |  | [optional] |
| **devices** | [**Array&lt;NodeDeviceResource&gt;**](NodeDeviceResource.md) |  | [optional] |
| **disk** | [**NodeDiskResources**](NodeDiskResources.md) |  | [optional] |
| **max_dynamic_port** | **Integer** |  | [optional] |
| **memory** | [**NodeMemoryResources**](NodeMemoryResources.md) |  | [optional] |
| **min_dynamic_port** | **Integer** |  | [optional] |
| **networks** | [**Array&lt;NetworkResource&gt;**](NetworkResource.md) |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::NodeResources.new(
  cpu: null,
  devices: null,
  disk: null,
  max_dynamic_port: null,
  memory: null,
  min_dynamic_port: null,
  networks: null
)
```

