# nomad-client.CSIVolume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**accessMode** | **String** |  | [optional] 
**allocations** | [**[AllocationListStub]**](AllocationListStub.md) |  | [optional] 
**attachmentMode** | **String** |  | [optional] 
**capacity** | **Number** |  | [optional] 
**cloneID** | **String** |  | [optional] 
**context** | **{String: String}** |  | [optional] 
**controllerRequired** | **Boolean** |  | [optional] 
**controllersExpected** | **Number** |  | [optional] 
**controllersHealthy** | **Number** |  | [optional] 
**createIndex** | **Number** |  | [optional] 
**externalID** | **String** |  | [optional] 
**ID** | **String** |  | [optional] 
**modifyIndex** | **Number** |  | [optional] 
**mountOptions** | [**CSIMountOptions**](CSIMountOptions.md) |  | [optional] 
**name** | **String** |  | [optional] 
**namespace** | **String** |  | [optional] 
**nodesExpected** | **Number** |  | [optional] 
**nodesHealthy** | **Number** |  | [optional] 
**parameters** | **{String: String}** |  | [optional] 
**pluginID** | **String** |  | [optional] 
**provider** | **String** |  | [optional] 
**providerVersion** | **String** |  | [optional] 
**readAllocs** | [**{String: Allocation}**](Allocation.md) |  | [optional] 
**requestedCapabilities** | [**[CSIVolumeCapability]**](CSIVolumeCapability.md) |  | [optional] 
**requestedCapacityMax** | **Number** |  | [optional] 
**requestedCapacityMin** | **Number** |  | [optional] 
**resourceExhausted** | **Date** |  | [optional] 
**schedulable** | **Boolean** |  | [optional] 
**secrets** | **{String: String}** |  | [optional] 
**snapshotID** | **String** |  | [optional] 
**topologies** | [**[CSITopology]**](CSITopology.md) |  | [optional] 
**writeAllocs** | [**{String: Allocation}**](Allocation.md) |  | [optional] 


