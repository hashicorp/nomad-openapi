# nomad-client.Node

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**attributes** | **{String: String}** |  | [optional] 
**cSIControllerPlugins** | [**{String: CSIInfo}**](CSIInfo.md) |  | [optional] 
**cSINodePlugins** | [**{String: CSIInfo}**](CSIInfo.md) |  | [optional] 
**createIndex** | **Number** |  | [optional] 
**datacenter** | **String** |  | [optional] 
**drain** | **Boolean** |  | [optional] 
**drainStrategy** | [**DrainStrategy**](DrainStrategy.md) |  | [optional] 
**drivers** | [**{String: DriverInfo}**](DriverInfo.md) |  | [optional] 
**events** | [**[NodeEvent]**](NodeEvent.md) |  | [optional] 
**hTTPAddr** | **String** |  | [optional] 
**hostVolumes** | [**{String: HostVolumeInfo}**](HostVolumeInfo.md) |  | [optional] 
**ID** | **String** |  | [optional] 
**lastDrain** | [**DrainMetadata**](DrainMetadata.md) |  | [optional] 
**links** | **{String: String}** |  | [optional] 
**meta** | **{String: String}** |  | [optional] 
**modifyIndex** | **Number** |  | [optional] 
**name** | **String** |  | [optional] 
**nodeClass** | **String** |  | [optional] 
**nodeResources** | [**NodeResources**](NodeResources.md) |  | [optional] 
**reserved** | [**Resources**](Resources.md) |  | [optional] 
**reservedResources** | [**NodeReservedResources**](NodeReservedResources.md) |  | [optional] 
**resources** | [**Resources**](Resources.md) |  | [optional] 
**schedulingEligibility** | **String** |  | [optional] 
**status** | **String** |  | [optional] 
**statusDescription** | **String** |  | [optional] 
**statusUpdatedAt** | **Number** |  | [optional] 
**tLSEnabled** | **Boolean** |  | [optional] 


