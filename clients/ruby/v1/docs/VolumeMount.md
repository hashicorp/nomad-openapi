# NomadClient::VolumeMount

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **destination** | **String** |  | [optional] |
| **propagation_mode** | **String** |  | [optional] |
| **read_only** | **Boolean** |  | [optional] |
| **volume** | **String** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::VolumeMount.new(
  destination: null,
  propagation_mode: null,
  read_only: null,
  volume: null
)
```

