# NomadClient::DesiredTransition

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **migrate** | **Boolean** |  | [optional] |
| **reschedule** | **Boolean** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::DesiredTransition.new(
  migrate: null,
  reschedule: null
)
```

