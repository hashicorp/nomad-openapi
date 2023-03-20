# NomadClient::JobListStub

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **create_index** | **Integer** |  | [optional] |
| **datacenters** | **Array&lt;String&gt;** |  | [optional] |
| **id** | **String** |  | [optional] |
| **job_modify_index** | **Integer** |  | [optional] |
| **job_summary** | [**JobSummary**](JobSummary.md) |  | [optional] |
| **meta** | **Hash&lt;String, String&gt;** |  | [optional] |
| **modify_index** | **Integer** |  | [optional] |
| **name** | **String** |  | [optional] |
| **namespace** | **String** |  | [optional] |
| **parameterized_job** | **Boolean** |  | [optional] |
| **parent_id** | **String** |  | [optional] |
| **periodic** | **Boolean** |  | [optional] |
| **priority** | **Integer** |  | [optional] |
| **status** | **String** |  | [optional] |
| **status_description** | **String** |  | [optional] |
| **stop** | **Boolean** |  | [optional] |
| **submit_time** | **Integer** |  | [optional] |
| **type** | **String** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::JobListStub.new(
  create_index: null,
  datacenters: null,
  id: null,
  job_modify_index: null,
  job_summary: null,
  meta: null,
  modify_index: null,
  name: null,
  namespace: null,
  parameterized_job: null,
  parent_id: null,
  periodic: null,
  priority: null,
  status: null,
  status_description: null,
  stop: null,
  submit_time: null,
  type: null
)
```

