# NomadClient::Node

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **attributes** | **Hash&lt;String, String&gt;** |  | [optional] |
| **csi_controller_plugins** | [**Hash&lt;String, CSIInfo&gt;**](CSIInfo.md) |  | [optional] |
| **csi_node_plugins** | [**Hash&lt;String, CSIInfo&gt;**](CSIInfo.md) |  | [optional] |
| **create_index** | **Integer** |  | [optional] |
| **datacenter** | **String** |  | [optional] |
| **drain** | **Boolean** |  | [optional] |
| **drain_strategy** | [**DrainStrategy**](DrainStrategy.md) |  | [optional] |
| **drivers** | [**Hash&lt;String, DriverInfo&gt;**](DriverInfo.md) |  | [optional] |
| **events** | [**Array&lt;NodeEvent&gt;**](NodeEvent.md) |  | [optional] |
| **http_addr** | **String** |  | [optional] |
| **host_volumes** | [**Hash&lt;String, HostVolumeInfo&gt;**](HostVolumeInfo.md) |  | [optional] |
| **id** | **String** |  | [optional] |
| **last_drain** | [**DrainMetadata**](DrainMetadata.md) |  | [optional] |
| **links** | **Hash&lt;String, String&gt;** |  | [optional] |
| **meta** | **Hash&lt;String, String&gt;** |  | [optional] |
| **modify_index** | **Integer** |  | [optional] |
| **name** | **String** |  | [optional] |
| **node_class** | **String** |  | [optional] |
| **node_resources** | [**NodeResources**](NodeResources.md) |  | [optional] |
| **reserved** | [**Resources**](Resources.md) |  | [optional] |
| **reserved_resources** | [**NodeReservedResources**](NodeReservedResources.md) |  | [optional] |
| **resources** | [**Resources**](Resources.md) |  | [optional] |
| **scheduling_eligibility** | **String** |  | [optional] |
| **status** | **String** |  | [optional] |
| **status_description** | **String** |  | [optional] |
| **status_updated_at** | **Integer** |  | [optional] |
| **tls_enabled** | **Boolean** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::Node.new(
  attributes: null,
  csi_controller_plugins: null,
  csi_node_plugins: null,
  create_index: null,
  datacenter: null,
  drain: null,
  drain_strategy: null,
  drivers: null,
  events: null,
  http_addr: null,
  host_volumes: null,
  id: null,
  last_drain: null,
  links: null,
  meta: null,
  modify_index: null,
  name: null,
  node_class: null,
  node_resources: null,
  reserved: null,
  reserved_resources: null,
  resources: null,
  scheduling_eligibility: null,
  status: null,
  status_description: null,
  status_updated_at: null,
  tls_enabled: null
)
```

