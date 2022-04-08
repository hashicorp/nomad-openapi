# NomadClient::CSIControllerInfo

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **supports_attach_detach** | **Boolean** |  | [optional] |
| **supports_clone** | **Boolean** |  | [optional] |
| **supports_condition** | **Boolean** |  | [optional] |
| **supports_create_delete** | **Boolean** |  | [optional] |
| **supports_create_delete_snapshot** | **Boolean** |  | [optional] |
| **supports_expand** | **Boolean** |  | [optional] |
| **supports_get** | **Boolean** |  | [optional] |
| **supports_get_capacity** | **Boolean** |  | [optional] |
| **supports_list_snapshots** | **Boolean** |  | [optional] |
| **supports_list_volumes** | **Boolean** |  | [optional] |
| **supports_list_volumes_attached_nodes** | **Boolean** |  | [optional] |
| **supports_read_only_attach** | **Boolean** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::CSIControllerInfo.new(
  supports_attach_detach: null,
  supports_clone: null,
  supports_condition: null,
  supports_create_delete: null,
  supports_create_delete_snapshot: null,
  supports_expand: null,
  supports_get: null,
  supports_get_capacity: null,
  supports_list_snapshots: null,
  supports_list_volumes: null,
  supports_list_volumes_attached_nodes: null,
  supports_read_only_attach: null
)
```

