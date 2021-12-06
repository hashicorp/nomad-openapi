# NomadClient::CSIVolumeCapability

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **access_mode** | **String** |  | [optional] |
| **attachment_mode** | **String** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::CSIVolumeCapability.new(
  access_mode: null,
  attachment_mode: null
)
```

