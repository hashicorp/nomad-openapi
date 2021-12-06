# NomadClient::JobPlanRequest

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **diff** | **Boolean** |  | [optional] |
| **job** | [**Job**](Job.md) |  | [optional] |
| **namespace** | **String** |  | [optional] |
| **policy_override** | **Boolean** |  | [optional] |
| **region** | **String** |  | [optional] |
| **secret_id** | **String** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::JobPlanRequest.new(
  diff: null,
  job: null,
  namespace: null,
  policy_override: null,
  region: null,
  secret_id: null
)
```

