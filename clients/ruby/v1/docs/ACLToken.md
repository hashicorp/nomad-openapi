# NomadClient::ACLToken

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **accessor_id** | **String** |  | [optional] |
| **create_index** | **Integer** |  | [optional] |
| **create_time** | **Time** |  | [optional] |
| **expiration_ttl** | **Integer** |  | [optional] |
| **expiration_time** | **Time** |  | [optional] |
| **global** | **Boolean** |  | [optional] |
| **modify_index** | **Integer** |  | [optional] |
| **name** | **String** |  | [optional] |
| **policies** | **Array&lt;String&gt;** |  | [optional] |
| **roles** | [**Array&lt;ACLTokenRoleLink&gt;**](ACLTokenRoleLink.md) |  | [optional] |
| **secret_id** | **String** |  | [optional] |
| **type** | **String** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::ACLToken.new(
  accessor_id: null,
  create_index: null,
  create_time: null,
  expiration_ttl: null,
  expiration_time: null,
  global: null,
  modify_index: null,
  name: null,
  policies: null,
  roles: null,
  secret_id: null,
  type: null
)
```

