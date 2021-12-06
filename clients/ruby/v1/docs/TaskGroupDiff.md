# NomadClient::TaskGroupDiff

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **fields** | [**Array&lt;FieldDiff&gt;**](FieldDiff.md) |  | [optional] |
| **name** | **String** |  | [optional] |
| **objects** | [**Array&lt;ObjectDiff&gt;**](ObjectDiff.md) |  | [optional] |
| **tasks** | [**Array&lt;TaskDiff&gt;**](TaskDiff.md) |  | [optional] |
| **type** | **String** |  | [optional] |
| **updates** | **Hash&lt;String, Integer&gt;** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::TaskGroupDiff.new(
  fields: null,
  name: null,
  objects: null,
  tasks: null,
  type: null,
  updates: null
)
```

