# NomadClient::FuzzySearchRequest

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **allow_stale** | **Boolean** |  | [optional] |
| **auth_token** | **String** |  | [optional] |
| **context** | **String** |  | [optional] |
| **namespace** | **String** |  | [optional] |
| **next_token** | **String** |  | [optional] |
| **params** | **Hash&lt;String, String&gt;** |  | [optional] |
| **per_page** | **Integer** |  | [optional] |
| **prefix** | **String** |  | [optional] |
| **region** | **String** |  | [optional] |
| **text** | **String** |  | [optional] |
| **wait_index** | **Integer** |  | [optional] |
| **wait_time** | **Integer** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::FuzzySearchRequest.new(
  allow_stale: null,
  auth_token: null,
  context: null,
  namespace: null,
  next_token: null,
  params: null,
  per_page: null,
  prefix: null,
  region: null,
  text: null,
  wait_index: null,
  wait_time: null
)
```

