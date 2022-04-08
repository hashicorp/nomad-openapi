# NomadClient::CSINodeInfo

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **accessible_topology** | [**CSITopology**](CSITopology.md) |  | [optional] |
| **id** | **String** |  | [optional] |
| **max_volumes** | **Integer** |  | [optional] |
| **requires_node_stage_volume** | **Boolean** |  | [optional] |
| **supports_condition** | **Boolean** |  | [optional] |
| **supports_expand** | **Boolean** |  | [optional] |
| **supports_stats** | **Boolean** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::CSINodeInfo.new(
  accessible_topology: null,
  id: null,
  max_volumes: null,
  requires_node_stage_volume: null,
  supports_condition: null,
  supports_expand: null,
  supports_stats: null
)
```

