# Nomad.Client.Model.CSIVolume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessMode** | **string** |  | [optional] 
**Allocations** | [**List&lt;AllocationListStub&gt;**](AllocationListStub.md) |  | [optional] 
**AttachmentMode** | **string** |  | [optional] 
**Capacity** | **long** |  | [optional] 
**CloneID** | **string** |  | [optional] 
**Context** | **Dictionary&lt;string, string&gt;** |  | [optional] 
**ControllerRequired** | **bool** |  | [optional] 
**ControllersExpected** | **int** |  | [optional] 
**ControllersHealthy** | **int** |  | [optional] 
**CreateIndex** | **int** |  | [optional] 
**ExternalID** | **string** |  | [optional] 
**ID** | **string** |  | [optional] 
**ModifyIndex** | **int** |  | [optional] 
**MountOptions** | [**CSIMountOptions**](CSIMountOptions.md) |  | [optional] 
**Name** | **string** |  | [optional] 
**Namespace** | **string** |  | [optional] 
**NodesExpected** | **int** |  | [optional] 
**NodesHealthy** | **int** |  | [optional] 
**Parameters** | **Dictionary&lt;string, string&gt;** |  | [optional] 
**PluginID** | **string** |  | [optional] 
**Provider** | **string** |  | [optional] 
**ProviderVersion** | **string** |  | [optional] 
**ReadAllocs** | [**Dictionary&lt;string, Allocation&gt;**](Allocation.md) |  | [optional] 
**RequestedCapabilities** | [**List&lt;CSIVolumeCapability&gt;**](CSIVolumeCapability.md) |  | [optional] 
**RequestedCapacityMax** | **long** |  | [optional] 
**RequestedCapacityMin** | **long** |  | [optional] 
**RequestedTopologies** | [**CSITopologyRequest**](CSITopologyRequest.md) |  | [optional] 
**ResourceExhausted** | **DateTime** |  | [optional] 
**Schedulable** | **bool** |  | [optional] 
**Secrets** | **Dictionary&lt;string, string&gt;** |  | [optional] 
**SnapshotID** | **string** |  | [optional] 
**Topologies** | [**List&lt;CSITopology&gt;**](CSITopology.md) |  | [optional] 
**WriteAllocs** | [**Dictionary&lt;string, Allocation&gt;**](Allocation.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

