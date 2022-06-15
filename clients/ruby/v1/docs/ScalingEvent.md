# NomadClient::ScalingEvent

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **count** | **Integer** |  | [optional] |
| **create_index** | **Integer** |  | [optional] |
| **error** | **Boolean** |  | [optional] |
| **eval_id** | **String** |  | [optional] |
| **message** | **String** |  | [optional] |
| **meta** | **Hash&lt;String, Object&gt;** |  | [optional] |
| **previous_count** | **Integer** |  | [optional] |
| **time** | **Integer** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::ScalingEvent.new(
  count: null,
  create_index: null,
  error: null,
  eval_id: null,
  message: null,
  meta: null,
  previous_count: null,
  time: null
)
```

