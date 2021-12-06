# NomadClient::AllocationMetric

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **allocation_time** | **Integer** |  | [optional] |
| **class_exhausted** | **Hash&lt;String, Integer&gt;** |  | [optional] |
| **class_filtered** | **Hash&lt;String, Integer&gt;** |  | [optional] |
| **coalesced_failures** | **Integer** |  | [optional] |
| **constraint_filtered** | **Hash&lt;String, Integer&gt;** |  | [optional] |
| **dimension_exhausted** | **Hash&lt;String, Integer&gt;** |  | [optional] |
| **nodes_available** | **Hash&lt;String, Integer&gt;** |  | [optional] |
| **nodes_evaluated** | **Integer** |  | [optional] |
| **nodes_exhausted** | **Integer** |  | [optional] |
| **nodes_filtered** | **Integer** |  | [optional] |
| **quota_exhausted** | **Array&lt;String&gt;** |  | [optional] |
| **resources_exhausted** | [**Hash&lt;String, Resources&gt;**](Resources.md) |  | [optional] |
| **score_meta_data** | [**Array&lt;NodeScoreMeta&gt;**](NodeScoreMeta.md) |  | [optional] |
| **scores** | **Hash&lt;String, Float&gt;** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::AllocationMetric.new(
  allocation_time: null,
  class_exhausted: null,
  class_filtered: null,
  coalesced_failures: null,
  constraint_filtered: null,
  dimension_exhausted: null,
  nodes_available: null,
  nodes_evaluated: null,
  nodes_exhausted: null,
  nodes_filtered: null,
  quota_exhausted: null,
  resources_exhausted: null,
  score_meta_data: null,
  scores: null
)
```

