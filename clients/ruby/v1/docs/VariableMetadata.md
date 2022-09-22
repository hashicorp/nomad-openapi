# NomadClient::VariableMetadata

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **create_index** | **Integer** |  | [optional] |
| **create_time** | **Integer** |  | [optional] |
| **modify_index** | **Integer** |  | [optional] |
| **modify_time** | **Integer** |  | [optional] |
| **namespace** | **String** |  | [optional] |
| **path** | **String** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::VariableMetadata.new(
  create_index: null,
  create_time: null,
  modify_index: null,
  modify_time: null,
  namespace: null,
  path: null
)
```

