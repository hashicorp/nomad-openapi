# NomadClient::CSIInfo

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **alloc_id** | **String** |  | [optional] |
| **controller_info** | [**CSIControllerInfo**](CSIControllerInfo.md) |  | [optional] |
| **health_description** | **String** |  | [optional] |
| **healthy** | **Boolean** |  | [optional] |
| **node_info** | [**CSINodeInfo**](CSINodeInfo.md) |  | [optional] |
| **plugin_id** | **String** |  | [optional] |
| **requires_controller_plugin** | **Boolean** |  | [optional] |
| **requires_topologies** | **Boolean** |  | [optional] |
| **update_time** | **Time** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::CSIInfo.new(
  alloc_id: null,
  controller_info: null,
  health_description: null,
  healthy: null,
  node_info: null,
  plugin_id: null,
  requires_controller_plugin: null,
  requires_topologies: null,
  update_time: null
)
```

