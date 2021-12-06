# NomadClient::DrainStrategy

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **deadline** | **Integer** |  | [optional] |
| **force_deadline** | **Time** |  | [optional] |
| **ignore_system_jobs** | **Boolean** |  | [optional] |
| **started_at** | **Time** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::DrainStrategy.new(
  deadline: null,
  force_deadline: null,
  ignore_system_jobs: null,
  started_at: null
)
```

