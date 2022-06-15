

# Node


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**attributes** | **Map&lt;String, String&gt;** |  |  [optional] |
|**csIControllerPlugins** | [**Map&lt;String, CSIInfo&gt;**](CSIInfo.md) |  |  [optional] |
|**csINodePlugins** | [**Map&lt;String, CSIInfo&gt;**](CSIInfo.md) |  |  [optional] |
|**cgroupParent** | **String** |  |  [optional] |
|**createIndex** | **Integer** |  |  [optional] |
|**datacenter** | **String** |  |  [optional] |
|**drain** | **Boolean** |  |  [optional] |
|**drainStrategy** | [**DrainStrategy**](DrainStrategy.md) |  |  [optional] |
|**drivers** | [**Map&lt;String, DriverInfo&gt;**](DriverInfo.md) |  |  [optional] |
|**events** | [**List&lt;NodeEvent&gt;**](NodeEvent.md) |  |  [optional] |
|**htTPAddr** | **String** |  |  [optional] |
|**hostNetworks** | [**Map&lt;String, HostNetworkInfo&gt;**](HostNetworkInfo.md) |  |  [optional] |
|**hostVolumes** | [**Map&lt;String, HostVolumeInfo&gt;**](HostVolumeInfo.md) |  |  [optional] |
|**ID** | **String** |  |  [optional] |
|**lastDrain** | [**DrainMetadata**](DrainMetadata.md) |  |  [optional] |
|**links** | **Map&lt;String, String&gt;** |  |  [optional] |
|**meta** | **Map&lt;String, String&gt;** |  |  [optional] |
|**modifyIndex** | **Integer** |  |  [optional] |
|**name** | **String** |  |  [optional] |
|**nodeClass** | **String** |  |  [optional] |
|**nodeResources** | [**NodeResources**](NodeResources.md) |  |  [optional] |
|**reserved** | [**Resources**](Resources.md) |  |  [optional] |
|**reservedResources** | [**NodeReservedResources**](NodeReservedResources.md) |  |  [optional] |
|**resources** | [**Resources**](Resources.md) |  |  [optional] |
|**schedulingEligibility** | **String** |  |  [optional] |
|**status** | **String** |  |  [optional] |
|**statusDescription** | **String** |  |  [optional] |
|**statusUpdatedAt** | **Long** |  |  [optional] |
|**tlSEnabled** | **Boolean** |  |  [optional] |



