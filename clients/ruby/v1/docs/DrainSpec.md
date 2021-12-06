# NomadClient::DrainSpec

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **deadline** | **Integer** |  | [optional] |
| **ignore_system_jobs** | **Boolean** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::DrainSpec.new(
  deadline: null,
  ignore_system_jobs: null
)
```

