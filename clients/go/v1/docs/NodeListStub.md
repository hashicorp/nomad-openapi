# NodeListStub

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** |  | [optional] 
**Attributes** | Pointer to **map[string]string** |  | [optional] 
**CreateIndex** | Pointer to **int32** |  | [optional] 
**Datacenter** | Pointer to **string** |  | [optional] 
**Drain** | Pointer to **bool** |  | [optional] 
**Drivers** | Pointer to [**map[string]DriverInfo**](DriverInfo.md) |  | [optional] 
**ID** | Pointer to **string** |  | [optional] 
**LastDrain** | Pointer to [**NullableDrainMetadata**](DrainMetadata.md) |  | [optional] 
**ModifyIndex** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NodeClass** | Pointer to **string** |  | [optional] 
**NodeResources** | Pointer to [**NullableNodeResources**](NodeResources.md) |  | [optional] 
**ReservedResources** | Pointer to [**NullableNodeReservedResources**](NodeReservedResources.md) |  | [optional] 
**SchedulingEligibility** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusDescription** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewNodeListStub

`func NewNodeListStub() *NodeListStub`

NewNodeListStub instantiates a new NodeListStub object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeListStubWithDefaults

`func NewNodeListStubWithDefaults() *NodeListStub`

NewNodeListStubWithDefaults instantiates a new NodeListStub object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *NodeListStub) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *NodeListStub) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *NodeListStub) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *NodeListStub) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAttributes

`func (o *NodeListStub) GetAttributes() map[string]string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *NodeListStub) GetAttributesOk() (*map[string]string, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *NodeListStub) SetAttributes(v map[string]string)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *NodeListStub) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetCreateIndex

`func (o *NodeListStub) GetCreateIndex() int32`

GetCreateIndex returns the CreateIndex field if non-nil, zero value otherwise.

### GetCreateIndexOk

`func (o *NodeListStub) GetCreateIndexOk() (*int32, bool)`

GetCreateIndexOk returns a tuple with the CreateIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateIndex

`func (o *NodeListStub) SetCreateIndex(v int32)`

SetCreateIndex sets CreateIndex field to given value.

### HasCreateIndex

`func (o *NodeListStub) HasCreateIndex() bool`

HasCreateIndex returns a boolean if a field has been set.

### GetDatacenter

`func (o *NodeListStub) GetDatacenter() string`

GetDatacenter returns the Datacenter field if non-nil, zero value otherwise.

### GetDatacenterOk

`func (o *NodeListStub) GetDatacenterOk() (*string, bool)`

GetDatacenterOk returns a tuple with the Datacenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenter

`func (o *NodeListStub) SetDatacenter(v string)`

SetDatacenter sets Datacenter field to given value.

### HasDatacenter

`func (o *NodeListStub) HasDatacenter() bool`

HasDatacenter returns a boolean if a field has been set.

### GetDrain

`func (o *NodeListStub) GetDrain() bool`

GetDrain returns the Drain field if non-nil, zero value otherwise.

### GetDrainOk

`func (o *NodeListStub) GetDrainOk() (*bool, bool)`

GetDrainOk returns a tuple with the Drain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrain

`func (o *NodeListStub) SetDrain(v bool)`

SetDrain sets Drain field to given value.

### HasDrain

`func (o *NodeListStub) HasDrain() bool`

HasDrain returns a boolean if a field has been set.

### GetDrivers

`func (o *NodeListStub) GetDrivers() map[string]DriverInfo`

GetDrivers returns the Drivers field if non-nil, zero value otherwise.

### GetDriversOk

`func (o *NodeListStub) GetDriversOk() (*map[string]DriverInfo, bool)`

GetDriversOk returns a tuple with the Drivers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrivers

`func (o *NodeListStub) SetDrivers(v map[string]DriverInfo)`

SetDrivers sets Drivers field to given value.

### HasDrivers

`func (o *NodeListStub) HasDrivers() bool`

HasDrivers returns a boolean if a field has been set.

### GetID

`func (o *NodeListStub) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *NodeListStub) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *NodeListStub) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *NodeListStub) HasID() bool`

HasID returns a boolean if a field has been set.

### GetLastDrain

`func (o *NodeListStub) GetLastDrain() DrainMetadata`

GetLastDrain returns the LastDrain field if non-nil, zero value otherwise.

### GetLastDrainOk

`func (o *NodeListStub) GetLastDrainOk() (*DrainMetadata, bool)`

GetLastDrainOk returns a tuple with the LastDrain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDrain

`func (o *NodeListStub) SetLastDrain(v DrainMetadata)`

SetLastDrain sets LastDrain field to given value.

### HasLastDrain

`func (o *NodeListStub) HasLastDrain() bool`

HasLastDrain returns a boolean if a field has been set.

### SetLastDrainNil

`func (o *NodeListStub) SetLastDrainNil(b bool)`

 SetLastDrainNil sets the value for LastDrain to be an explicit nil

### UnsetLastDrain
`func (o *NodeListStub) UnsetLastDrain()`

UnsetLastDrain ensures that no value is present for LastDrain, not even an explicit nil
### GetModifyIndex

`func (o *NodeListStub) GetModifyIndex() int32`

GetModifyIndex returns the ModifyIndex field if non-nil, zero value otherwise.

### GetModifyIndexOk

`func (o *NodeListStub) GetModifyIndexOk() (*int32, bool)`

GetModifyIndexOk returns a tuple with the ModifyIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifyIndex

`func (o *NodeListStub) SetModifyIndex(v int32)`

SetModifyIndex sets ModifyIndex field to given value.

### HasModifyIndex

`func (o *NodeListStub) HasModifyIndex() bool`

HasModifyIndex returns a boolean if a field has been set.

### GetName

`func (o *NodeListStub) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NodeListStub) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NodeListStub) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NodeListStub) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNodeClass

`func (o *NodeListStub) GetNodeClass() string`

GetNodeClass returns the NodeClass field if non-nil, zero value otherwise.

### GetNodeClassOk

`func (o *NodeListStub) GetNodeClassOk() (*string, bool)`

GetNodeClassOk returns a tuple with the NodeClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeClass

`func (o *NodeListStub) SetNodeClass(v string)`

SetNodeClass sets NodeClass field to given value.

### HasNodeClass

`func (o *NodeListStub) HasNodeClass() bool`

HasNodeClass returns a boolean if a field has been set.

### GetNodeResources

`func (o *NodeListStub) GetNodeResources() NodeResources`

GetNodeResources returns the NodeResources field if non-nil, zero value otherwise.

### GetNodeResourcesOk

`func (o *NodeListStub) GetNodeResourcesOk() (*NodeResources, bool)`

GetNodeResourcesOk returns a tuple with the NodeResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeResources

`func (o *NodeListStub) SetNodeResources(v NodeResources)`

SetNodeResources sets NodeResources field to given value.

### HasNodeResources

`func (o *NodeListStub) HasNodeResources() bool`

HasNodeResources returns a boolean if a field has been set.

### SetNodeResourcesNil

`func (o *NodeListStub) SetNodeResourcesNil(b bool)`

 SetNodeResourcesNil sets the value for NodeResources to be an explicit nil

### UnsetNodeResources
`func (o *NodeListStub) UnsetNodeResources()`

UnsetNodeResources ensures that no value is present for NodeResources, not even an explicit nil
### GetReservedResources

`func (o *NodeListStub) GetReservedResources() NodeReservedResources`

GetReservedResources returns the ReservedResources field if non-nil, zero value otherwise.

### GetReservedResourcesOk

`func (o *NodeListStub) GetReservedResourcesOk() (*NodeReservedResources, bool)`

GetReservedResourcesOk returns a tuple with the ReservedResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedResources

`func (o *NodeListStub) SetReservedResources(v NodeReservedResources)`

SetReservedResources sets ReservedResources field to given value.

### HasReservedResources

`func (o *NodeListStub) HasReservedResources() bool`

HasReservedResources returns a boolean if a field has been set.

### SetReservedResourcesNil

`func (o *NodeListStub) SetReservedResourcesNil(b bool)`

 SetReservedResourcesNil sets the value for ReservedResources to be an explicit nil

### UnsetReservedResources
`func (o *NodeListStub) UnsetReservedResources()`

UnsetReservedResources ensures that no value is present for ReservedResources, not even an explicit nil
### GetSchedulingEligibility

`func (o *NodeListStub) GetSchedulingEligibility() string`

GetSchedulingEligibility returns the SchedulingEligibility field if non-nil, zero value otherwise.

### GetSchedulingEligibilityOk

`func (o *NodeListStub) GetSchedulingEligibilityOk() (*string, bool)`

GetSchedulingEligibilityOk returns a tuple with the SchedulingEligibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedulingEligibility

`func (o *NodeListStub) SetSchedulingEligibility(v string)`

SetSchedulingEligibility sets SchedulingEligibility field to given value.

### HasSchedulingEligibility

`func (o *NodeListStub) HasSchedulingEligibility() bool`

HasSchedulingEligibility returns a boolean if a field has been set.

### GetStatus

`func (o *NodeListStub) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NodeListStub) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NodeListStub) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NodeListStub) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusDescription

`func (o *NodeListStub) GetStatusDescription() string`

GetStatusDescription returns the StatusDescription field if non-nil, zero value otherwise.

### GetStatusDescriptionOk

`func (o *NodeListStub) GetStatusDescriptionOk() (*string, bool)`

GetStatusDescriptionOk returns a tuple with the StatusDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDescription

`func (o *NodeListStub) SetStatusDescription(v string)`

SetStatusDescription sets StatusDescription field to given value.

### HasStatusDescription

`func (o *NodeListStub) HasStatusDescription() bool`

HasStatusDescription returns a boolean if a field has been set.

### GetVersion

`func (o *NodeListStub) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *NodeListStub) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *NodeListStub) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *NodeListStub) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


