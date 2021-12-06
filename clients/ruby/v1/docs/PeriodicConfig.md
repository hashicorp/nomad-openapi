# NomadClient::PeriodicConfig

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **enabled** | **Boolean** |  | [optional] |
| **prohibit_overlap** | **Boolean** |  | [optional] |
| **spec** | **String** |  | [optional] |
| **spec_type** | **String** |  | [optional] |
| **time_zone** | **String** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::PeriodicConfig.new(
  enabled: null,
  prohibit_overlap: null,
  spec: null,
  spec_type: null,
  time_zone: null
)
```

