# NomadClient::JobScaleStatusResponse

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **job_create_index** | **Integer** |  | [optional] |
| **job_id** | **String** |  | [optional] |
| **job_modify_index** | **Integer** |  | [optional] |
| **job_stopped** | **Boolean** |  | [optional] |
| **namespace** | **String** |  | [optional] |
| **task_groups** | [**Hash&lt;String, TaskGroupScaleStatus&gt;**](TaskGroupScaleStatus.md) |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::JobScaleStatusResponse.new(
  job_create_index: null,
  job_id: null,
  job_modify_index: null,
  job_stopped: null,
  namespace: null,
  task_groups: null
)
```

