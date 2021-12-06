# NomadClient::CSIControllerInfo

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **supports_attach_detach** | **Boolean** |  | [optional] |
| **supports_list_volumes** | **Boolean** |  | [optional] |
| **supports_list_volumes_attached_nodes** | **Boolean** |  | [optional] |
| **supports_read_only_attach** | **Boolean** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::CSIControllerInfo.new(
  supports_attach_detach: null,
  supports_list_volumes: null,
  supports_list_volumes_attached_nodes: null,
  supports_read_only_attach: null
)
```

