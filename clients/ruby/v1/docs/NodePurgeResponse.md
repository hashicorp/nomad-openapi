# NomadClient::NodePurgeResponse

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **eval_create_index** | **Integer** |  | [optional] |
| **eval_ids** | **Array&lt;String&gt;** |  | [optional] |
| **node_modify_index** | **Integer** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::NodePurgeResponse.new(
  eval_create_index: null,
  eval_ids: null,
  node_modify_index: null
)
```

