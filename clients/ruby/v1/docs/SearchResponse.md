# NomadClient::SearchResponse

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **known_leader** | **Boolean** |  | [optional] |
| **last_contact** | **Integer** |  | [optional] |
| **last_index** | **Integer** |  | [optional] |
| **matches** | **Hash&lt;String, Array&lt;String&gt;&gt;** |  | [optional] |
| **next_token** | **String** |  | [optional] |
| **request_time** | **Integer** |  | [optional] |
| **truncations** | **Hash&lt;String, Boolean&gt;** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::SearchResponse.new(
  known_leader: null,
  last_contact: null,
  last_index: null,
  matches: null,
  next_token: null,
  request_time: null,
  truncations: null
)
```

