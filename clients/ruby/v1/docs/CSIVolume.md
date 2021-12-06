# NomadClient::CSIVolume

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **access_mode** | **String** |  | [optional] |
| **allocations** | [**Array&lt;AllocationListStub&gt;**](AllocationListStub.md) |  | [optional] |
| **attachment_mode** | **String** |  | [optional] |
| **capacity** | **Integer** |  | [optional] |
| **clone_id** | **String** |  | [optional] |
| **context** | **Hash&lt;String, String&gt;** |  | [optional] |
| **controller_required** | **Boolean** |  | [optional] |
| **controllers_expected** | **Integer** |  | [optional] |
| **controllers_healthy** | **Integer** |  | [optional] |
| **create_index** | **Integer** |  | [optional] |
| **external_id** | **String** |  | [optional] |
| **id** | **String** |  | [optional] |
| **modify_index** | **Integer** |  | [optional] |
| **mount_options** | [**CSIMountOptions**](CSIMountOptions.md) |  | [optional] |
| **name** | **String** |  | [optional] |
| **namespace** | **String** |  | [optional] |
| **nodes_expected** | **Integer** |  | [optional] |
| **nodes_healthy** | **Integer** |  | [optional] |
| **parameters** | **Hash&lt;String, String&gt;** |  | [optional] |
| **plugin_id** | **String** |  | [optional] |
| **provider** | **String** |  | [optional] |
| **provider_version** | **String** |  | [optional] |
| **read_allocs** | [**Hash&lt;String, Allocation&gt;**](Allocation.md) |  | [optional] |
| **requested_capabilities** | [**Array&lt;CSIVolumeCapability&gt;**](CSIVolumeCapability.md) |  | [optional] |
| **requested_capacity_max** | **Integer** |  | [optional] |
| **requested_capacity_min** | **Integer** |  | [optional] |
| **resource_exhausted** | **Time** |  | [optional] |
| **schedulable** | **Boolean** |  | [optional] |
| **secrets** | **Hash&lt;String, String&gt;** |  | [optional] |
| **snapshot_id** | **String** |  | [optional] |
| **topologies** | [**Array&lt;CSITopology&gt;**](CSITopology.md) |  | [optional] |
| **write_allocs** | [**Hash&lt;String, Allocation&gt;**](Allocation.md) |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::CSIVolume.new(
  access_mode: null,
  allocations: null,
  attachment_mode: null,
  capacity: null,
  clone_id: null,
  context: null,
  controller_required: null,
  controllers_expected: null,
  controllers_healthy: null,
  create_index: null,
  external_id: null,
  id: null,
  modify_index: null,
  mount_options: null,
  name: null,
  namespace: null,
  nodes_expected: null,
  nodes_healthy: null,
  parameters: null,
  plugin_id: null,
  provider: null,
  provider_version: null,
  read_allocs: null,
  requested_capabilities: null,
  requested_capacity_max: null,
  requested_capacity_min: null,
  resource_exhausted: null,
  schedulable: null,
  secrets: null,
  snapshot_id: null,
  topologies: null,
  write_allocs: null
)
```

