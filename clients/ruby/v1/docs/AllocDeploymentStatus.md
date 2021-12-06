# NomadClient::AllocDeploymentStatus

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **canary** | **Boolean** |  | [optional] |
| **healthy** | **Boolean** |  | [optional] |
| **modify_index** | **Integer** |  | [optional] |
| **timestamp** | **Time** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::AllocDeploymentStatus.new(
  canary: null,
  healthy: null,
  modify_index: null,
  timestamp: null
)
```

