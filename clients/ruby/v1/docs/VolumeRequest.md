# NomadClient::VolumeRequest

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **access_mode** | **String** |  | [optional] |
| **attachment_mode** | **String** |  | [optional] |
| **mount_options** | [**CSIMountOptions**](CSIMountOptions.md) |  | [optional] |
| **name** | **String** |  | [optional] |
| **per_alloc** | **Boolean** |  | [optional] |
| **read_only** | **Boolean** |  | [optional] |
| **source** | **String** |  | [optional] |
| **type** | **String** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::VolumeRequest.new(
  access_mode: null,
  attachment_mode: null,
  mount_options: null,
  name: null,
  per_alloc: null,
  read_only: null,
  source: null,
  type: null
)
```

