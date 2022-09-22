# NomadClient::Variable

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **create_index** | **Integer** |  | [optional] |
| **create_time** | **Integer** |  | [optional] |
| **items** | **Hash&lt;String, String&gt;** |  | [optional] |
| **modify_index** | **Integer** |  | [optional] |
| **modify_time** | **Integer** |  | [optional] |
| **namespace** | **String** |  | [optional] |
| **path** | **String** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::Variable.new(
  create_index: null,
  create_time: null,
  items: null,
  modify_index: null,
  modify_time: null,
  namespace: null,
  path: null
)
```

