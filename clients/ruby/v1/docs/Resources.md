# NomadClient::Resources

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **cpu** | **Integer** |  | [optional] |
| **cores** | **Integer** |  | [optional] |
| **devices** | [**Array&lt;RequestedDevice&gt;**](RequestedDevice.md) |  | [optional] |
| **disk_mb** | **Integer** |  | [optional] |
| **iops** | **Integer** |  | [optional] |
| **memory_mb** | **Integer** |  | [optional] |
| **memory_max_mb** | **Integer** |  | [optional] |
| **networks** | [**Array&lt;NetworkResource&gt;**](NetworkResource.md) |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::Resources.new(
  cpu: null,
  cores: null,
  devices: null,
  disk_mb: null,
  iops: null,
  memory_mb: null,
  memory_max_mb: null,
  networks: null
)
```

