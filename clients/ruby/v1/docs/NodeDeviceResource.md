# NomadClient::NodeDeviceResource

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **attributes** | [**Hash&lt;String, Attribute&gt;**](Attribute.md) |  | [optional] |
| **instances** | [**Array&lt;NodeDevice&gt;**](NodeDevice.md) |  | [optional] |
| **name** | **String** |  | [optional] |
| **type** | **String** |  | [optional] |
| **vendor** | **String** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::NodeDeviceResource.new(
  attributes: null,
  instances: null,
  name: null,
  type: null,
  vendor: null
)
```

