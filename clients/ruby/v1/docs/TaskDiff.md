# NomadClient::TaskDiff

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **annotations** | **Array&lt;String&gt;** |  | [optional] |
| **fields** | [**Array&lt;FieldDiff&gt;**](FieldDiff.md) |  | [optional] |
| **name** | **String** |  | [optional] |
| **objects** | [**Array&lt;ObjectDiff&gt;**](ObjectDiff.md) |  | [optional] |
| **type** | **String** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::TaskDiff.new(
  annotations: null,
  fields: null,
  name: null,
  objects: null,
  type: null
)
```

