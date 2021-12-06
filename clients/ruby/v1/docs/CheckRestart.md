# NomadClient::CheckRestart

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **grace** | **Integer** |  | [optional] |
| **ignore_warnings** | **Boolean** |  | [optional] |
| **limit** | **Integer** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::CheckRestart.new(
  grace: null,
  ignore_warnings: null,
  limit: null
)
```

