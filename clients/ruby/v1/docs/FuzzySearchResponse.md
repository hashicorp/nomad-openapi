# NomadClient::FuzzySearchResponse

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **known_leader** | **Boolean** |  | [optional] |
| **last_contact** | **Integer** |  | [optional] |
| **last_index** | **Integer** |  | [optional] |
| **matches** | **Hash&lt;String, Array&lt;FuzzyMatch&gt;&gt;** |  | [optional] |
| **request_time** | **Integer** |  | [optional] |
| **truncations** | **Hash&lt;String, Boolean&gt;** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::FuzzySearchResponse.new(
  known_leader: null,
  last_contact: null,
  last_index: null,
  matches: null,
  request_time: null,
  truncations: null
)
```

