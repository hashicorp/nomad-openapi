# Node


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**attributes** | **{str: (str,)}** |  | [optional] 
**csi_controller_plugins** | [**{str: (CSIInfo,)}**](CSIInfo.md) |  | [optional] 
**csi_node_plugins** | [**{str: (CSIInfo,)}**](CSIInfo.md) |  | [optional] 
**cgroup_parent** | **str** |  | [optional] 
**create_index** | **int** |  | [optional] 
**datacenter** | **str** |  | [optional] 
**drain** | **bool** |  | [optional] 
**drain_strategy** | [**DrainStrategy**](DrainStrategy.md) |  | [optional] 
**drivers** | [**{str: (DriverInfo,)}**](DriverInfo.md) |  | [optional] 
**events** | [**[NodeEvent]**](NodeEvent.md) |  | [optional] 
**http_addr** | **str** |  | [optional] 
**host_networks** | [**{str: (HostNetworkInfo,)}**](HostNetworkInfo.md) |  | [optional] 
**host_volumes** | [**{str: (HostVolumeInfo,)}**](HostVolumeInfo.md) |  | [optional] 
**id** | **str** |  | [optional] 
**last_drain** | [**DrainMetadata**](DrainMetadata.md) |  | [optional] 
**links** | **{str: (str,)}** |  | [optional] 
**meta** | **{str: (str,)}** |  | [optional] 
**modify_index** | **int** |  | [optional] 
**name** | **str** |  | [optional] 
**node_class** | **str** |  | [optional] 
**node_resources** | [**NodeResources**](NodeResources.md) |  | [optional] 
**reserved** | [**Resources**](Resources.md) |  | [optional] 
**reserved_resources** | [**NodeReservedResources**](NodeReservedResources.md) |  | [optional] 
**resources** | [**Resources**](Resources.md) |  | [optional] 
**scheduling_eligibility** | **str** |  | [optional] 
**status** | **str** |  | [optional] 
**status_description** | **str** |  | [optional] 
**status_updated_at** | **int** |  | [optional] 
**tls_enabled** | **bool** |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


