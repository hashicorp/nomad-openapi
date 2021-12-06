# NomadClient::MultiregionStrategy

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **max_parallel** | **Integer** |  | [optional] |
| **on_failure** | **String** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::MultiregionStrategy.new(
  max_parallel: null,
  on_failure: null
)
```

