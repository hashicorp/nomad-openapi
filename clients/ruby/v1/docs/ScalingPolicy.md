# NomadClient::ScalingPolicy

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **create_index** | **Integer** |  | [optional] |
| **enabled** | **Boolean** |  | [optional] |
| **id** | **String** |  | [optional] |
| **max** | **Integer** |  | [optional] |
| **min** | **Integer** |  | [optional] |
| **modify_index** | **Integer** |  | [optional] |
| **namespace** | **String** |  | [optional] |
| **policy** | **Hash&lt;String, Object&gt;** |  | [optional] |
| **target** | **Hash&lt;String, String&gt;** |  | [optional] |
| **type** | **String** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::ScalingPolicy.new(
  create_index: null,
  enabled: null,
  id: null,
  max: null,
  min: null,
  modify_index: null,
  namespace: null,
  policy: null,
  target: null,
  type: null
)
```

