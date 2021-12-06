# NomadClient::JobDeregisterResponse

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **eval_create_index** | **Integer** |  | [optional] |
| **eval_id** | **String** |  | [optional] |
| **job_modify_index** | **Integer** |  | [optional] |
| **known_leader** | **Boolean** |  | [optional] |
| **last_contact** | **Integer** |  | [optional] |
| **last_index** | **Integer** |  | [optional] |
| **request_time** | **Integer** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::JobDeregisterResponse.new(
  eval_create_index: null,
  eval_id: null,
  job_modify_index: null,
  known_leader: null,
  last_contact: null,
  last_index: null,
  request_time: null
)
```

