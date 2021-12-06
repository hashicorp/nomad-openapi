# NomadClient::NodeCpuResources

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **cpu_shares** | **Integer** |  | [optional] |
| **reservable_cpu_cores** | **Array&lt;Integer&gt;** |  | [optional] |
| **total_cpu_cores** | **Integer** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::NodeCpuResources.new(
  cpu_shares: null,
  reservable_cpu_cores: null,
  total_cpu_cores: null
)
```

