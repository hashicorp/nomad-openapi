# NomadClient::NodeScoreMeta

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **node_id** | **String** |  | [optional] |
| **norm_score** | **Float** |  | [optional] |
| **scores** | **Hash&lt;String, Float&gt;** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::NodeScoreMeta.new(
  node_id: null,
  norm_score: null,
  scores: null
)
```

