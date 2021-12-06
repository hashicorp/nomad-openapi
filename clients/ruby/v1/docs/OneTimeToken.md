# NomadClient::OneTimeToken

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **accessor_id** | **String** |  | [optional] |
| **create_index** | **Integer** |  | [optional] |
| **expires_at** | **Time** |  | [optional] |
| **modify_index** | **Integer** |  | [optional] |
| **one_time_secret_id** | **String** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::OneTimeToken.new(
  accessor_id: null,
  create_index: null,
  expires_at: null,
  modify_index: null,
  one_time_secret_id: null
)
```

