# NomadClient::CSIVolumeListStub

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **access_mode** | **String** |  | [optional] |
| **attachment_mode** | **String** |  | [optional] |
| **controller_required** | **Boolean** |  | [optional] |
| **controllers_expected** | **Integer** |  | [optional] |
| **controllers_healthy** | **Integer** |  | [optional] |
| **create_index** | **Integer** |  | [optional] |
| **current_readers** | **Integer** |  | [optional] |
| **current_writers** | **Integer** |  | [optional] |
| **external_id** | **String** |  | [optional] |
| **id** | **String** |  | [optional] |
| **modify_index** | **Integer** |  | [optional] |
| **name** | **String** |  | [optional] |
| **namespace** | **String** |  | [optional] |
| **nodes_expected** | **Integer** |  | [optional] |
| **nodes_healthy** | **Integer** |  | [optional] |
| **plugin_id** | **String** |  | [optional] |
| **provider** | **String** |  | [optional] |
| **resource_exhausted** | **Time** |  | [optional] |
| **schedulable** | **Boolean** |  | [optional] |
| **topologies** | [**Array&lt;CSITopology&gt;**](CSITopology.md) |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::CSIVolumeListStub.new(
  access_mode: null,
  attachment_mode: null,
  controller_required: null,
  controllers_expected: null,
  controllers_healthy: null,
  create_index: null,
  current_readers: null,
  current_writers: null,
  external_id: null,
  id: null,
  modify_index: null,
  name: null,
  namespace: null,
  nodes_expected: null,
  nodes_healthy: null,
  plugin_id: null,
  provider: null,
  resource_exhausted: null,
  schedulable: null,
  topologies: null
)
```

