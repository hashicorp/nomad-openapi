# NomadClient::JobDiff

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **fields** | [**Array&lt;FieldDiff&gt;**](FieldDiff.md) |  | [optional] |
| **id** | **String** |  | [optional] |
| **objects** | [**Array&lt;ObjectDiff&gt;**](ObjectDiff.md) |  | [optional] |
| **task_groups** | [**Array&lt;TaskGroupDiff&gt;**](TaskGroupDiff.md) |  | [optional] |
| **type** | **String** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::JobDiff.new(
  fields: null,
  id: null,
  objects: null,
  task_groups: null,
  type: null
)
```

