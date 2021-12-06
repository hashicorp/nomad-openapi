# NomadClient::QuotaSpec

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **create_index** | **Integer** |  | [optional] |
| **description** | **String** |  | [optional] |
| **limits** | [**Array&lt;QuotaLimit&gt;**](QuotaLimit.md) |  | [optional] |
| **modify_index** | **Integer** |  | [optional] |
| **name** | **String** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::QuotaSpec.new(
  create_index: null,
  description: null,
  limits: null,
  modify_index: null,
  name: null
)
```

