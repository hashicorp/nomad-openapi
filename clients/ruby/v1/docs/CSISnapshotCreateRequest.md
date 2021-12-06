# NomadClient::CSISnapshotCreateRequest

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **namespace** | **String** |  | [optional] |
| **region** | **String** |  | [optional] |
| **secret_id** | **String** |  | [optional] |
| **snapshots** | [**Array&lt;CSISnapshot&gt;**](CSISnapshot.md) |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::CSISnapshotCreateRequest.new(
  namespace: null,
  region: null,
  secret_id: null,
  snapshots: null
)
```

