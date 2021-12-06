# NomadClient::DrainMetadata

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **accessor_id** | **String** |  | [optional] |
| **meta** | **Hash&lt;String, String&gt;** |  | [optional] |
| **started_at** | **Time** |  | [optional] |
| **status** | **String** |  | [optional] |
| **updated_at** | **Time** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::DrainMetadata.new(
  accessor_id: null,
  meta: null,
  started_at: null,
  status: null,
  updated_at: null
)
```

