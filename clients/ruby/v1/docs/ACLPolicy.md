# NomadClient::ACLPolicy

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **create_index** | **Integer** |  | [optional] |
| **description** | **String** |  | [optional] |
| **modify_index** | **Integer** |  | [optional] |
| **name** | **String** |  | [optional] |
| **rules** | **String** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::ACLPolicy.new(
  create_index: null,
  description: null,
  modify_index: null,
  name: null,
  rules: null
)
```

