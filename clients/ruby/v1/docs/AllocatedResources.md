# NomadClient::AllocatedResources

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **shared** | [**AllocatedSharedResources**](AllocatedSharedResources.md) |  | [optional] |
| **tasks** | [**Hash&lt;String, AllocatedTaskResources&gt;**](AllocatedTaskResources.md) |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::AllocatedResources.new(
  shared: null,
  tasks: null
)
```

