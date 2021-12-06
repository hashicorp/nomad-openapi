# NomadClient::TaskLifecycle

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **hook** | **String** |  | [optional] |
| **sidecar** | **Boolean** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::TaskLifecycle.new(
  hook: null,
  sidecar: null
)
```

