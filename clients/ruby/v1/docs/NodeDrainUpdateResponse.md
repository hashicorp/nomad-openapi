# NomadClient::NodeDrainUpdateResponse

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **eval_create_index** | **Integer** |  | [optional] |
| **eval_ids** | **Array&lt;String&gt;** |  | [optional] |
| **last_index** | **Integer** |  | [optional] |
| **node_modify_index** | **Integer** |  | [optional] |
| **request_time** | **Integer** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::NodeDrainUpdateResponse.new(
  eval_create_index: null,
  eval_ids: null,
  last_index: null,
  node_modify_index: null,
  request_time: null
)
```

