# NomadClient::JobDispatchResponse

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **dispatched_job_id** | **String** |  | [optional] |
| **eval_create_index** | **Integer** |  | [optional] |
| **eval_id** | **String** |  | [optional] |
| **job_create_index** | **Integer** |  | [optional] |
| **last_index** | **Integer** |  | [optional] |
| **request_time** | **Integer** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::JobDispatchResponse.new(
  dispatched_job_id: null,
  eval_create_index: null,
  eval_id: null,
  job_create_index: null,
  last_index: null,
  request_time: null
)
```

