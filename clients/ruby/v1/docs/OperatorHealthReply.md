# NomadClient::OperatorHealthReply

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **failure_tolerance** | **Integer** |  | [optional] |
| **healthy** | **Boolean** |  | [optional] |
| **servers** | [**Array&lt;ServerHealth&gt;**](ServerHealth.md) |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::OperatorHealthReply.new(
  failure_tolerance: null,
  healthy: null,
  servers: null
)
```

