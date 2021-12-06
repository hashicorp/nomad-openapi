# NomadClient::AllocatedDeviceResource

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **device_ids** | **Array&lt;String&gt;** |  | [optional] |
| **name** | **String** |  | [optional] |
| **type** | **String** |  | [optional] |
| **vendor** | **String** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::AllocatedDeviceResource.new(
  device_ids: null,
  name: null,
  type: null,
  vendor: null
)
```

