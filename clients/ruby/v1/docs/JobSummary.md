# NomadClient::JobSummary

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **children** | [**JobChildrenSummary**](JobChildrenSummary.md) |  | [optional] |
| **create_index** | **Integer** |  | [optional] |
| **job_id** | **String** |  | [optional] |
| **modify_index** | **Integer** |  | [optional] |
| **namespace** | **String** |  | [optional] |
| **summary** | [**Hash&lt;String, TaskGroupSummary&gt;**](TaskGroupSummary.md) |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::JobSummary.new(
  children: null,
  create_index: null,
  job_id: null,
  modify_index: null,
  namespace: null,
  summary: null
)
```

