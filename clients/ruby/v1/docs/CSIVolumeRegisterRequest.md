# NomadClient::CSIVolumeRegisterRequest

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **namespace** | **String** |  | [optional] |
| **region** | **String** |  | [optional] |
| **secret_id** | **String** |  | [optional] |
| **volumes** | [**Array&lt;CSIVolume&gt;**](CSIVolume.md) |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::CSIVolumeRegisterRequest.new(
  namespace: null,
  region: null,
  secret_id: null,
  volumes: null
)
```

