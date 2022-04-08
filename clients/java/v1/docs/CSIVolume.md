

# CSIVolume


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**accessMode** | **String** |  |  [optional]
**allocations** | [**List&lt;AllocationListStub&gt;**](AllocationListStub.md) |  |  [optional]
**attachmentMode** | **String** |  |  [optional]
**capacity** | **Long** |  |  [optional]
**cloneID** | **String** |  |  [optional]
**context** | **Map&lt;String, String&gt;** |  |  [optional]
**controllerRequired** | **Boolean** |  |  [optional]
**controllersExpected** | **Integer** |  |  [optional]
**controllersHealthy** | **Integer** |  |  [optional]
**createIndex** | **Integer** |  |  [optional]
**externalID** | **String** |  |  [optional]
**ID** | **String** |  |  [optional]
**modifyIndex** | **Integer** |  |  [optional]
**mountOptions** | [**CSIMountOptions**](CSIMountOptions.md) |  |  [optional]
**name** | **String** |  |  [optional]
**namespace** | **String** |  |  [optional]
**nodesExpected** | **Integer** |  |  [optional]
**nodesHealthy** | **Integer** |  |  [optional]
**parameters** | **Map&lt;String, String&gt;** |  |  [optional]
**pluginID** | **String** |  |  [optional]
**provider** | **String** |  |  [optional]
**providerVersion** | **String** |  |  [optional]
**readAllocs** | [**Map&lt;String, Allocation&gt;**](Allocation.md) |  |  [optional]
**requestedCapabilities** | [**List&lt;CSIVolumeCapability&gt;**](CSIVolumeCapability.md) |  |  [optional]
**requestedCapacityMax** | **Long** |  |  [optional]
**requestedCapacityMin** | **Long** |  |  [optional]
**requestedTopologies** | [**CSITopologyRequest**](CSITopologyRequest.md) |  |  [optional]
**resourceExhausted** | **OffsetDateTime** |  |  [optional]
**schedulable** | **Boolean** |  |  [optional]
**secrets** | **Map&lt;String, String&gt;** |  |  [optional]
**snapshotID** | **String** |  |  [optional]
**topologies** | [**List&lt;CSITopology&gt;**](CSITopology.md) |  |  [optional]
**writeAllocs** | [**Map&lt;String, Allocation&gt;**](Allocation.md) |  |  [optional]



