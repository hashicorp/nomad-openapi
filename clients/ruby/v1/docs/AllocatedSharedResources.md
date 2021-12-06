# NomadClient::AllocatedSharedResources

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **disk_mb** | **Integer** |  | [optional] |
| **networks** | [**Array&lt;NetworkResource&gt;**](NetworkResource.md) |  | [optional] |
| **ports** | [**Array&lt;PortMapping&gt;**](PortMapping.md) |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::AllocatedSharedResources.new(
  disk_mb: null,
  networks: null,
  ports: null
)
```

