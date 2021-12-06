# NomadClient::NodeListStub

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **address** | **String** |  | [optional] |
| **create_index** | **Integer** |  | [optional] |
| **datacenter** | **String** |  | [optional] |
| **drain** | **Boolean** |  | [optional] |
| **drivers** | [**Hash&lt;String, DriverInfo&gt;**](DriverInfo.md) |  | [optional] |
| **id** | **String** |  | [optional] |
| **last_drain** | [**DrainMetadata**](DrainMetadata.md) |  | [optional] |
| **modify_index** | **Integer** |  | [optional] |
| **name** | **String** |  | [optional] |
| **node_class** | **String** |  | [optional] |
| **node_resources** | [**NodeResources**](NodeResources.md) |  | [optional] |
| **reserved_resources** | [**NodeReservedResources**](NodeReservedResources.md) |  | [optional] |
| **scheduling_eligibility** | **String** |  | [optional] |
| **status** | **String** |  | [optional] |
| **status_description** | **String** |  | [optional] |
| **version** | **String** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::NodeListStub.new(
  address: null,
  create_index: null,
  datacenter: null,
  drain: null,
  drivers: null,
  id: null,
  last_drain: null,
  modify_index: null,
  name: null,
  node_class: null,
  node_resources: null,
  reserved_resources: null,
  scheduling_eligibility: null,
  status: null,
  status_description: null,
  version: null
)
```

