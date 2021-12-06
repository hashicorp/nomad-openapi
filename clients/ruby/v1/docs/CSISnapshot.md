# NomadClient::CSISnapshot

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **create_time** | **Integer** |  | [optional] |
| **external_source_volume_id** | **String** |  | [optional] |
| **id** | **String** |  | [optional] |
| **is_ready** | **Boolean** |  | [optional] |
| **name** | **String** |  | [optional] |
| **parameters** | **Hash&lt;String, String&gt;** |  | [optional] |
| **plugin_id** | **String** |  | [optional] |
| **secrets** | **Hash&lt;String, String&gt;** |  | [optional] |
| **size_bytes** | **Integer** |  | [optional] |
| **source_volume_id** | **String** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::CSISnapshot.new(
  create_time: null,
  external_source_volume_id: null,
  id: null,
  is_ready: null,
  name: null,
  parameters: null,
  plugin_id: null,
  secrets: null,
  size_bytes: null,
  source_volume_id: null
)
```

