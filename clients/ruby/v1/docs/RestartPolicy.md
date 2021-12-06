# NomadClient::RestartPolicy

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **attempts** | **Integer** |  | [optional] |
| **delay** | **Integer** |  | [optional] |
| **interval** | **Integer** |  | [optional] |
| **mode** | **String** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::RestartPolicy.new(
  attempts: null,
  delay: null,
  interval: null,
  mode: null
)
```

