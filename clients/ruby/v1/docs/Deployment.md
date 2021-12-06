# NomadClient::Deployment

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **create_index** | **Integer** |  | [optional] |
| **id** | **String** |  | [optional] |
| **is_multiregion** | **Boolean** |  | [optional] |
| **job_create_index** | **Integer** |  | [optional] |
| **job_id** | **String** |  | [optional] |
| **job_modify_index** | **Integer** |  | [optional] |
| **job_spec_modify_index** | **Integer** |  | [optional] |
| **job_version** | **Integer** |  | [optional] |
| **modify_index** | **Integer** |  | [optional] |
| **namespace** | **String** |  | [optional] |
| **status** | **String** |  | [optional] |
| **status_description** | **String** |  | [optional] |
| **task_groups** | [**Hash&lt;String, DeploymentState&gt;**](DeploymentState.md) |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::Deployment.new(
  create_index: null,
  id: null,
  is_multiregion: null,
  job_create_index: null,
  job_id: null,
  job_modify_index: null,
  job_spec_modify_index: null,
  job_version: null,
  modify_index: null,
  namespace: null,
  status: null,
  status_description: null,
  task_groups: null
)
```

