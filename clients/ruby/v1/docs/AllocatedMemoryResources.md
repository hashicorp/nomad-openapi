# NomadClient::AllocatedMemoryResources

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **memory_mb** | **Integer** |  | [optional] |
| **memory_max_mb** | **Integer** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::AllocatedMemoryResources.new(
  memory_mb: null,
  memory_max_mb: null
)
```

