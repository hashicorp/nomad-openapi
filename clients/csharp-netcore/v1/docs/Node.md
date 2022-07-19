# Nomad.Client.Model.Node

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | **Dictionary&lt;string, string&gt;** |  | [optional] 
**CSIControllerPlugins** | [**Dictionary&lt;string, CSIInfo&gt;**](CSIInfo.md) |  | [optional] 
**CSINodePlugins** | [**Dictionary&lt;string, CSIInfo&gt;**](CSIInfo.md) |  | [optional] 
**CgroupParent** | **string** |  | [optional] 
**CreateIndex** | **int** |  | [optional] 
**Datacenter** | **string** |  | [optional] 
**Drain** | **bool** |  | [optional] 
**DrainStrategy** | [**DrainStrategy**](DrainStrategy.md) |  | [optional] 
**Drivers** | [**Dictionary&lt;string, DriverInfo&gt;**](DriverInfo.md) |  | [optional] 
**Events** | [**List&lt;NodeEvent&gt;**](NodeEvent.md) |  | [optional] 
**HTTPAddr** | **string** |  | [optional] 
**HostNetworks** | [**Dictionary&lt;string, HostNetworkInfo&gt;**](HostNetworkInfo.md) |  | [optional] 
**HostVolumes** | [**Dictionary&lt;string, HostVolumeInfo&gt;**](HostVolumeInfo.md) |  | [optional] 
**ID** | **string** |  | [optional] 
**LastDrain** | [**DrainMetadata**](DrainMetadata.md) |  | [optional] 
**Links** | **Dictionary&lt;string, string&gt;** |  | [optional] 
**Meta** | **Dictionary&lt;string, string&gt;** |  | [optional] 
**ModifyIndex** | **int** |  | [optional] 
**Name** | **string** |  | [optional] 
**NodeClass** | **string** |  | [optional] 
**NodeResources** | [**NodeResources**](NodeResources.md) |  | [optional] 
**Reserved** | [**Resources**](Resources.md) |  | [optional] 
**ReservedResources** | [**NodeReservedResources**](NodeReservedResources.md) |  | [optional] 
**Resources** | [**Resources**](Resources.md) |  | [optional] 
**SchedulingEligibility** | **string** |  | [optional] 
**Status** | **string** |  | [optional] 
**StatusDescription** | **string** |  | [optional] 
**StatusUpdatedAt** | **long** |  | [optional] 
**TLSEnabled** | **bool** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

