# NomadClient::ObjectDiff

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **fields** | [**Array&lt;FieldDiff&gt;**](FieldDiff.md) |  | [optional] |
| **name** | **String** |  | [optional] |
| **objects** | [**Array&lt;ObjectDiff&gt;**](ObjectDiff.md) |  | [optional] |
| **type** | **String** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::ObjectDiff.new(
  fields: null,
  name: null,
  objects: null,
  type: null
)
```

