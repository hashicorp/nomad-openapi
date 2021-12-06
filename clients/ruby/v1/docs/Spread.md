# NomadClient::Spread

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **attribute** | **String** |  | [optional] |
| **spread_target** | [**Array&lt;SpreadTarget&gt;**](SpreadTarget.md) |  | [optional] |
| **weight** | **Integer** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::Spread.new(
  attribute: null,
  spread_target: null,
  weight: null
)
```

