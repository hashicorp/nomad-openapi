# Node

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to **map[string]string** |  | [optional] 
**CSIControllerPlugins** | Pointer to [**map[string]CSIInfo**](CSIInfo.md) |  | [optional] 
**CSINodePlugins** | Pointer to [**map[string]CSIInfo**](CSIInfo.md) |  | [optional] 
**CgroupParent** | Pointer to **string** |  | [optional] 
**CreateIndex** | Pointer to **int32** |  | [optional] 
**Datacenter** | Pointer to **string** |  | [optional] 
**Drain** | Pointer to **bool** |  | [optional] 
**DrainStrategy** | Pointer to [**NullableDrainStrategy**](DrainStrategy.md) |  | [optional] 
**Drivers** | Pointer to [**map[string]DriverInfo**](DriverInfo.md) |  | [optional] 
**Events** | Pointer to [**[]NodeEvent**](NodeEvent.md) |  | [optional] 
**HTTPAddr** | Pointer to **string** |  | [optional] 
**HostNetworks** | Pointer to [**map[string]HostNetworkInfo**](HostNetworkInfo.md) |  | [optional] 
**HostVolumes** | Pointer to [**map[string]HostVolumeInfo**](HostVolumeInfo.md) |  | [optional] 
**ID** | Pointer to **string** |  | [optional] 
**LastDrain** | Pointer to [**NullableDrainMetadata**](DrainMetadata.md) |  | [optional] 
**Links** | Pointer to **map[string]string** |  | [optional] 
**Meta** | Pointer to **map[string]string** |  | [optional] 
**ModifyIndex** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NodeClass** | Pointer to **string** |  | [optional] 
**NodeResources** | Pointer to [**NullableNodeResources**](NodeResources.md) |  | [optional] 
**Reserved** | Pointer to [**NullableResources**](Resources.md) |  | [optional] 
**ReservedResources** | Pointer to [**NullableNodeReservedResources**](NodeReservedResources.md) |  | [optional] 
**Resources** | Pointer to [**NullableResources**](Resources.md) |  | [optional] 
**SchedulingEligibility** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusDescription** | Pointer to **string** |  | [optional] 
**StatusUpdatedAt** | Pointer to **int64** |  | [optional] 
**TLSEnabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewNode

`func NewNode() *Node`

NewNode instantiates a new Node object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeWithDefaults

`func NewNodeWithDefaults() *Node`

NewNodeWithDefaults instantiates a new Node object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *Node) GetAttributes() map[string]string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Node) GetAttributesOk() (*map[string]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Node) SetAttributes(v map[string]string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *Node) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetCSIControllerPlugins

`func (o *Node) GetCSIControllerPlugins() map[string]CSIInfo`

GetCSIControllerPlugins returns the CSIControllerPlugins field if non-nil, zero value otherwise.

### GetCSIControllerPluginsOk

`func (o *Node) GetCSIControllerPluginsOk() (*map[string]CSIInfo, bool)`

GetCSIControllerPluginsOk returns a tuple with the CSIControllerPlugins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCSIControllerPlugins

`func (o *Node) SetCSIControllerPlugins(v map[string]CSIInfo)`

SetCSIControllerPlugins sets CSIControllerPlugins field to given value.

### HasCSIControllerPlugins

`func (o *Node) HasCSIControllerPlugins() bool`

HasCSIControllerPlugins returns a boolean if a field has been set.

### GetCSINodePlugins

`func (o *Node) GetCSINodePlugins() map[string]CSIInfo`

GetCSINodePlugins returns the CSINodePlugins field if non-nil, zero value otherwise.

### GetCSINodePluginsOk

`func (o *Node) GetCSINodePluginsOk() (*map[string]CSIInfo, bool)`

GetCSINodePluginsOk returns a tuple with the CSINodePlugins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCSINodePlugins

`func (o *Node) SetCSINodePlugins(v map[string]CSIInfo)`

SetCSINodePlugins sets CSINodePlugins field to given value.

### HasCSINodePlugins

`func (o *Node) HasCSINodePlugins() bool`

HasCSINodePlugins returns a boolean if a field has been set.

### GetCgroupParent

`func (o *Node) GetCgroupParent() string`

GetCgroupParent returns the CgroupParent field if non-nil, zero value otherwise.

### GetCgroupParentOk

`func (o *Node) GetCgroupParentOk() (*string, bool)`

GetCgroupParentOk returns a tuple with the CgroupParent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCgroupParent

`func (o *Node) SetCgroupParent(v string)`

SetCgroupParent sets CgroupParent field to given value.

### HasCgroupParent

`func (o *Node) HasCgroupParent() bool`

HasCgroupParent returns a boolean if a field has been set.

### GetCreateIndex

`func (o *Node) GetCreateIndex() int32`

GetCreateIndex returns the CreateIndex field if non-nil, zero value otherwise.

### GetCreateIndexOk

`func (o *Node) GetCreateIndexOk() (*int32, bool)`

GetCreateIndexOk returns a tuple with the CreateIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateIndex

`func (o *Node) SetCreateIndex(v int32)`

SetCreateIndex sets CreateIndex field to given value.

### HasCreateIndex

`func (o *Node) HasCreateIndex() bool`

HasCreateIndex returns a boolean if a field has been set.

### GetDatacenter

`func (o *Node) GetDatacenter() string`

GetDatacenter returns the Datacenter field if non-nil, zero value otherwise.

### GetDatacenterOk

`func (o *Node) GetDatacenterOk() (*string, bool)`

GetDatacenterOk returns a tuple with the Datacenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenter

`func (o *Node) SetDatacenter(v string)`

SetDatacenter sets Datacenter field to given value.

### HasDatacenter

`func (o *Node) HasDatacenter() bool`

HasDatacenter returns a boolean if a field has been set.

### GetDrain

`func (o *Node) GetDrain() bool`

GetDrain returns the Drain field if non-nil, zero value otherwise.

### GetDrainOk

`func (o *Node) GetDrainOk() (*bool, bool)`

GetDrainOk returns a tuple with the Drain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrain

`func (o *Node) SetDrain(v bool)`

SetDrain sets Drain field to given value.

### HasDrain

`func (o *Node) HasDrain() bool`

HasDrain returns a boolean if a field has been set.

### GetDrainStrategy

`func (o *Node) GetDrainStrategy() DrainStrategy`

GetDrainStrategy returns the DrainStrategy field if non-nil, zero value otherwise.

### GetDrainStrategyOk

`func (o *Node) GetDrainStrategyOk() (*DrainStrategy, bool)`

GetDrainStrategyOk returns a tuple with the DrainStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrainStrategy

`func (o *Node) SetDrainStrategy(v DrainStrategy)`

SetDrainStrategy sets DrainStrategy field to given value.

### HasDrainStrategy

`func (o *Node) HasDrainStrategy() bool`

HasDrainStrategy returns a boolean if a field has been set.

### SetDrainStrategyNil

`func (o *Node) SetDrainStrategyNil(b bool)`

 SetDrainStrategyNil sets the value for DrainStrategy to be an explicit nil

### UnsetDrainStrategy
`func (o *Node) UnsetDrainStrategy()`

UnsetDrainStrategy ensures that no value is present for DrainStrategy, not even an explicit nil
### GetDrivers

`func (o *Node) GetDrivers() map[string]DriverInfo`

GetDrivers returns the Drivers field if non-nil, zero value otherwise.

### GetDriversOk

`func (o *Node) GetDriversOk() (*map[string]DriverInfo, bool)`

GetDriversOk returns a tuple with the Drivers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrivers

`func (o *Node) SetDrivers(v map[string]DriverInfo)`

SetDrivers sets Drivers field to given value.

### HasDrivers

`func (o *Node) HasDrivers() bool`

HasDrivers returns a boolean if a field has been set.

### GetEvents

`func (o *Node) GetEvents() []NodeEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *Node) GetEventsOk() (*[]NodeEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *Node) SetEvents(v []NodeEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *Node) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetHTTPAddr

`func (o *Node) GetHTTPAddr() string`

GetHTTPAddr returns the HTTPAddr field if non-nil, zero value otherwise.

### GetHTTPAddrOk

`func (o *Node) GetHTTPAddrOk() (*string, bool)`

GetHTTPAddrOk returns a tuple with the HTTPAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHTTPAddr

`func (o *Node) SetHTTPAddr(v string)`

SetHTTPAddr sets HTTPAddr field to given value.

### HasHTTPAddr

`func (o *Node) HasHTTPAddr() bool`

HasHTTPAddr returns a boolean if a field has been set.

### GetHostNetworks

`func (o *Node) GetHostNetworks() map[string]HostNetworkInfo`

GetHostNetworks returns the HostNetworks field if non-nil, zero value otherwise.

### GetHostNetworksOk

`func (o *Node) GetHostNetworksOk() (*map[string]HostNetworkInfo, bool)`

GetHostNetworksOk returns a tuple with the HostNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostNetworks

`func (o *Node) SetHostNetworks(v map[string]HostNetworkInfo)`

SetHostNetworks sets HostNetworks field to given value.

### HasHostNetworks

`func (o *Node) HasHostNetworks() bool`

HasHostNetworks returns a boolean if a field has been set.

### GetHostVolumes

`func (o *Node) GetHostVolumes() map[string]HostVolumeInfo`

GetHostVolumes returns the HostVolumes field if non-nil, zero value otherwise.

### GetHostVolumesOk

`func (o *Node) GetHostVolumesOk() (*map[string]HostVolumeInfo, bool)`

GetHostVolumesOk returns a tuple with the HostVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostVolumes

`func (o *Node) SetHostVolumes(v map[string]HostVolumeInfo)`

SetHostVolumes sets HostVolumes field to given value.

### HasHostVolumes

`func (o *Node) HasHostVolumes() bool`

HasHostVolumes returns a boolean if a field has been set.

### GetID

`func (o *Node) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *Node) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *Node) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *Node) HasID() bool`

HasID returns a boolean if a field has been set.

### GetLastDrain

`func (o *Node) GetLastDrain() DrainMetadata`

GetLastDrain returns the LastDrain field if non-nil, zero value otherwise.

### GetLastDrainOk

`func (o *Node) GetLastDrainOk() (*DrainMetadata, bool)`

GetLastDrainOk returns a tuple with the LastDrain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDrain

`func (o *Node) SetLastDrain(v DrainMetadata)`

SetLastDrain sets LastDrain field to given value.

### HasLastDrain

`func (o *Node) HasLastDrain() bool`

HasLastDrain returns a boolean if a field has been set.

### SetLastDrainNil

`func (o *Node) SetLastDrainNil(b bool)`

 SetLastDrainNil sets the value for LastDrain to be an explicit nil

### UnsetLastDrain
`func (o *Node) UnsetLastDrain()`

UnsetLastDrain ensures that no value is present for LastDrain, not even an explicit nil
### GetLinks

`func (o *Node) GetLinks() map[string]string`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Node) GetLinksOk() (*map[string]string, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Node) SetLinks(v map[string]string)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Node) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *Node) GetMeta() map[string]string`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *Node) GetMetaOk() (*map[string]string, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *Node) SetMeta(v map[string]string)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *Node) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetModifyIndex

`func (o *Node) GetModifyIndex() int32`

GetModifyIndex returns the ModifyIndex field if non-nil, zero value otherwise.

### GetModifyIndexOk

`func (o *Node) GetModifyIndexOk() (*int32, bool)`

GetModifyIndexOk returns a tuple with the ModifyIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifyIndex

`func (o *Node) SetModifyIndex(v int32)`

SetModifyIndex sets ModifyIndex field to given value.

### HasModifyIndex

`func (o *Node) HasModifyIndex() bool`

HasModifyIndex returns a boolean if a field has been set.

### GetName

`func (o *Node) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Node) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Node) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Node) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNodeClass

`func (o *Node) GetNodeClass() string`

GetNodeClass returns the NodeClass field if non-nil, zero value otherwise.

### GetNodeClassOk

`func (o *Node) GetNodeClassOk() (*string, bool)`

GetNodeClassOk returns a tuple with the NodeClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeClass

`func (o *Node) SetNodeClass(v string)`

SetNodeClass sets NodeClass field to given value.

### HasNodeClass

`func (o *Node) HasNodeClass() bool`

HasNodeClass returns a boolean if a field has been set.

### GetNodeResources

`func (o *Node) GetNodeResources() NodeResources`

GetNodeResources returns the NodeResources field if non-nil, zero value otherwise.

### GetNodeResourcesOk

`func (o *Node) GetNodeResourcesOk() (*NodeResources, bool)`

GetNodeResourcesOk returns a tuple with the NodeResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeResources

`func (o *Node) SetNodeResources(v NodeResources)`

SetNodeResources sets NodeResources field to given value.

### HasNodeResources

`func (o *Node) HasNodeResources() bool`

HasNodeResources returns a boolean if a field has been set.

### SetNodeResourcesNil

`func (o *Node) SetNodeResourcesNil(b bool)`

 SetNodeResourcesNil sets the value for NodeResources to be an explicit nil

### UnsetNodeResources
`func (o *Node) UnsetNodeResources()`

UnsetNodeResources ensures that no value is present for NodeResources, not even an explicit nil
### GetReserved

`func (o *Node) GetReserved() Resources`

GetReserved returns the Reserved field if non-nil, zero value otherwise.

### GetReservedOk

`func (o *Node) GetReservedOk() (*Resources, bool)`

GetReservedOk returns a tuple with the Reserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserved

`func (o *Node) SetReserved(v Resources)`

SetReserved sets Reserved field to given value.

### HasReserved

`func (o *Node) HasReserved() bool`

HasReserved returns a boolean if a field has been set.

### SetReservedNil

`func (o *Node) SetReservedNil(b bool)`

 SetReservedNil sets the value for Reserved to be an explicit nil

### UnsetReserved
`func (o *Node) UnsetReserved()`

UnsetReserved ensures that no value is present for Reserved, not even an explicit nil
### GetReservedResources

`func (o *Node) GetReservedResources() NodeReservedResources`

GetReservedResources returns the ReservedResources field if non-nil, zero value otherwise.

### GetReservedResourcesOk

`func (o *Node) GetReservedResourcesOk() (*NodeReservedResources, bool)`

GetReservedResourcesOk returns a tuple with the ReservedResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedResources

`func (o *Node) SetReservedResources(v NodeReservedResources)`

SetReservedResources sets ReservedResources field to given value.

### HasReservedResources

`func (o *Node) HasReservedResources() bool`

HasReservedResources returns a boolean if a field has been set.

### SetReservedResourcesNil

`func (o *Node) SetReservedResourcesNil(b bool)`

 SetReservedResourcesNil sets the value for ReservedResources to be an explicit nil

### UnsetReservedResources
`func (o *Node) UnsetReservedResources()`

UnsetReservedResources ensures that no value is present for ReservedResources, not even an explicit nil
### GetResources

`func (o *Node) GetResources() Resources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *Node) GetResourcesOk() (*Resources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *Node) SetResources(v Resources)`

SetResources sets Resources field to given value.

### HasResources

`func (o *Node) HasResources() bool`

HasResources returns a boolean if a field has been set.

### SetResourcesNil

`func (o *Node) SetResourcesNil(b bool)`

 SetResourcesNil sets the value for Resources to be an explicit nil

### UnsetResources
`func (o *Node) UnsetResources()`

UnsetResources ensures that no value is present for Resources, not even an explicit nil
### GetSchedulingEligibility

`func (o *Node) GetSchedulingEligibility() string`

GetSchedulingEligibility returns the SchedulingEligibility field if non-nil, zero value otherwise.

### GetSchedulingEligibilityOk

`func (o *Node) GetSchedulingEligibilityOk() (*string, bool)`

GetSchedulingEligibilityOk returns a tuple with the SchedulingEligibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedulingEligibility

`func (o *Node) SetSchedulingEligibility(v string)`

SetSchedulingEligibility sets SchedulingEligibility field to given value.

### HasSchedulingEligibility

`func (o *Node) HasSchedulingEligibility() bool`

HasSchedulingEligibility returns a boolean if a field has been set.

### GetStatus

`func (o *Node) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Node) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Node) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Node) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusDescription

`func (o *Node) GetStatusDescription() string`

GetStatusDescription returns the StatusDescription field if non-nil, zero value otherwise.

### GetStatusDescriptionOk

`func (o *Node) GetStatusDescriptionOk() (*string, bool)`

GetStatusDescriptionOk returns a tuple with the StatusDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDescription

`func (o *Node) SetStatusDescription(v string)`

SetStatusDescription sets StatusDescription field to given value.

### HasStatusDescription

`func (o *Node) HasStatusDescription() bool`

HasStatusDescription returns a boolean if a field has been set.

### GetStatusUpdatedAt

`func (o *Node) GetStatusUpdatedAt() int64`

GetStatusUpdatedAt returns the StatusUpdatedAt field if non-nil, zero value otherwise.

### GetStatusUpdatedAtOk

`func (o *Node) GetStatusUpdatedAtOk() (*int64, bool)`

GetStatusUpdatedAtOk returns a tuple with the StatusUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusUpdatedAt

`func (o *Node) SetStatusUpdatedAt(v int64)`

SetStatusUpdatedAt sets StatusUpdatedAt field to given value.

### HasStatusUpdatedAt

`func (o *Node) HasStatusUpdatedAt() bool`

HasStatusUpdatedAt returns a boolean if a field has been set.

### GetTLSEnabled

`func (o *Node) GetTLSEnabled() bool`

GetTLSEnabled returns the TLSEnabled field if non-nil, zero value otherwise.

### GetTLSEnabledOk

`func (o *Node) GetTLSEnabledOk() (*bool, bool)`

GetTLSEnabledOk returns a tuple with the TLSEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTLSEnabled

`func (o *Node) SetTLSEnabled(v bool)`

SetTLSEnabled sets TLSEnabled field to given value.

### HasTLSEnabled

`func (o *Node) HasTLSEnabled() bool`

HasTLSEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


