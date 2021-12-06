# NomadClient::QuotaLimit

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **hash** | **String** |  | [optional] |
| **region** | **String** |  | [optional] |
| **region_limit** | [**Resources**](Resources.md) |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::QuotaLimit.new(
  hash: null,
  region: null,
  region_limit: null
)
```

