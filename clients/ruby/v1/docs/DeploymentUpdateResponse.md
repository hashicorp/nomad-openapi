# NomadClient::DeploymentUpdateResponse

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **deployment_modify_index** | **Integer** |  | [optional] |
| **eval_create_index** | **Integer** |  | [optional] |
| **eval_id** | **String** |  | [optional] |
| **last_index** | **Integer** |  | [optional] |
| **request_time** | **Integer** |  | [optional] |
| **reverted_job_version** | **Integer** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::DeploymentUpdateResponse.new(
  deployment_modify_index: null,
  eval_create_index: null,
  eval_id: null,
  last_index: null,
  request_time: null,
  reverted_job_version: null
)
```

