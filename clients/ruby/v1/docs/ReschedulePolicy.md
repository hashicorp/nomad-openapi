# NomadClient::ReschedulePolicy

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **attempts** | **Integer** |  | [optional] |
| **delay** | **Integer** |  | [optional] |
| **delay_function** | **String** |  | [optional] |
| **interval** | **Integer** |  | [optional] |
| **max_delay** | **Integer** |  | [optional] |
| **unlimited** | **Boolean** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::ReschedulePolicy.new(
  attempts: null,
  delay: null,
  delay_function: null,
  interval: null,
  max_delay: null,
  unlimited: null
)
```

