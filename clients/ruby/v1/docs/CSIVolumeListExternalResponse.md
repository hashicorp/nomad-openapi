# NomadClient::CSIVolumeListExternalResponse

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **next_token** | **String** |  | [optional] |
| **volumes** | [**Array&lt;CSIVolumeExternalStub&gt;**](CSIVolumeExternalStub.md) |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::CSIVolumeListExternalResponse.new(
  next_token: null,
  volumes: null
)
```

