# NomadClient::PreemptionConfig

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **batch_scheduler_enabled** | **Boolean** |  | [optional] |
| **service_scheduler_enabled** | **Boolean** |  | [optional] |
| **sys_batch_scheduler_enabled** | **Boolean** |  | [optional] |
| **system_scheduler_enabled** | **Boolean** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::PreemptionConfig.new(
  batch_scheduler_enabled: null,
  service_scheduler_enabled: null,
  sys_batch_scheduler_enabled: null,
  system_scheduler_enabled: null
)
```

