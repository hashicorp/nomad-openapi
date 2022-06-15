# NomadClient::ScalingRequest

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **count** | **Integer** |  | [optional] |
| **error** | **Boolean** |  | [optional] |
| **message** | **String** |  | [optional] |
| **meta** | **Hash&lt;String, Object&gt;** |  | [optional] |
| **namespace** | **String** |  | [optional] |
| **policy_override** | **Boolean** |  | [optional] |
| **region** | **String** |  | [optional] |
| **secret_id** | **String** |  | [optional] |
| **target** | **Hash&lt;String, String&gt;** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::ScalingRequest.new(
  count: null,
  error: null,
  message: null,
  meta: null,
  namespace: null,
  policy_override: null,
  region: null,
  secret_id: null,
  target: null
)
```

