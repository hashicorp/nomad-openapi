# NomadClient::JobRegisterRequest

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **enforce_index** | **Boolean** |  | [optional] |
| **eval_priority** | **Integer** |  | [optional] |
| **job** | [**Job**](Job.md) |  | [optional] |
| **job_modify_index** | **Integer** |  | [optional] |
| **namespace** | **String** |  | [optional] |
| **policy_override** | **Boolean** |  | [optional] |
| **preserve_counts** | **Boolean** |  | [optional] |
| **region** | **String** |  | [optional] |
| **secret_id** | **String** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::JobRegisterRequest.new(
  enforce_index: null,
  eval_priority: null,
  job: null,
  job_modify_index: null,
  namespace: null,
  policy_override: null,
  preserve_counts: null,
  region: null,
  secret_id: null
)
```

