# NomadClient::Multiregion

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **regions** | [**Array&lt;MultiregionRegion&gt;**](MultiregionRegion.md) |  | [optional] |
| **strategy** | [**MultiregionStrategy**](MultiregionStrategy.md) |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::Multiregion.new(
  regions: null,
  strategy: null
)
```

