# CSIVolume


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**access_mode** | **str** |  | [optional] 
**allocations** | [**[AllocationListStub]**](AllocationListStub.md) |  | [optional] 
**attachment_mode** | **str** |  | [optional] 
**capacity** | **int** |  | [optional] 
**clone_id** | **str** |  | [optional] 
**context** | **{str: (str,)}** |  | [optional] 
**controller_required** | **bool** |  | [optional] 
**controllers_expected** | **int** |  | [optional] 
**controllers_healthy** | **int** |  | [optional] 
**create_index** | **int** |  | [optional] 
**external_id** | **str** |  | [optional] 
**id** | **str** |  | [optional] 
**modify_index** | **int** |  | [optional] 
**mount_options** | [**CSIMountOptions**](CSIMountOptions.md) |  | [optional] 
**name** | **str** |  | [optional] 
**namespace** | **str** |  | [optional] 
**nodes_expected** | **int** |  | [optional] 
**nodes_healthy** | **int** |  | [optional] 
**parameters** | **{str: (str,)}** |  | [optional] 
**plugin_id** | **str** |  | [optional] 
**provider** | **str** |  | [optional] 
**provider_version** | **str** |  | [optional] 
**read_allocs** | [**{str: (Allocation,)}**](Allocation.md) |  | [optional] 
**requested_capabilities** | [**[CSIVolumeCapability]**](CSIVolumeCapability.md) |  | [optional] 
**requested_capacity_max** | **int** |  | [optional] 
**requested_capacity_min** | **int** |  | [optional] 
**requested_topologies** | [**CSITopologyRequest**](CSITopologyRequest.md) |  | [optional] 
**resource_exhausted** | **datetime, none_type** |  | [optional] 
**schedulable** | **bool** |  | [optional] 
**secrets** | [**CSISecrets**](CSISecrets.md) |  | [optional] 
**snapshot_id** | **str** |  | [optional] 
**topologies** | [**[CSITopology]**](CSITopology.md) |  | [optional] 
**write_allocs** | [**{str: (Allocation,)}**](Allocation.md) |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


