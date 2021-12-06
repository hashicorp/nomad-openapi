# NomadClient::JobEvaluateRequest

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **eval_options** | [**EvalOptions**](EvalOptions.md) |  | [optional] |
| **job_id** | **String** |  | [optional] |
| **namespace** | **String** |  | [optional] |
| **region** | **String** |  | [optional] |
| **secret_id** | **String** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::JobEvaluateRequest.new(
  eval_options: null,
  job_id: null,
  namespace: null,
  region: null,
  secret_id: null
)
```

