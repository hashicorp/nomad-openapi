# NomadClient::CSIVolumeExternalStub

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **capacity_bytes** | **Integer** |  | [optional] |
| **clone_id** | **String** |  | [optional] |
| **external_id** | **String** |  | [optional] |
| **is_abnormal** | **Boolean** |  | [optional] |
| **published_external_node_ids** | **Array&lt;String&gt;** |  | [optional] |
| **snapshot_id** | **String** |  | [optional] |
| **status** | **String** |  | [optional] |
| **volume_context** | **Hash&lt;String, String&gt;** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::CSIVolumeExternalStub.new(
  capacity_bytes: null,
  clone_id: null,
  external_id: null,
  is_abnormal: null,
  published_external_node_ids: null,
  snapshot_id: null,
  status: null,
  volume_context: null
)
```

