# CSIVolumeListStub

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessMode** | Pointer to **string** |  | [optional] 
**AttachmentMode** | Pointer to **string** |  | [optional] 
**ControllerRequired** | Pointer to **bool** |  | [optional] 
**ControllersExpected** | Pointer to **int32** |  | [optional] 
**ControllersHealthy** | Pointer to **int32** |  | [optional] 
**CreateIndex** | Pointer to **int32** |  | [optional] 
**ExternalID** | Pointer to **string** |  | [optional] 
**ID** | Pointer to **string** |  | [optional] 
**ModifyIndex** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**NodesExpected** | Pointer to **int32** |  | [optional] 
**NodesHealthy** | Pointer to **int32** |  | [optional] 
**PluginID** | Pointer to **string** |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**ResourceExhausted** | Pointer to **NullableTime** |  | [optional] 
**Schedulable** | Pointer to **bool** |  | [optional] 
**Topologies** | Pointer to [**[]CSITopology**](CSITopology.md) |  | [optional] 

## Methods

### NewCSIVolumeListStub

`func NewCSIVolumeListStub() *CSIVolumeListStub`

NewCSIVolumeListStub instantiates a new CSIVolumeListStub object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSIVolumeListStubWithDefaults

`func NewCSIVolumeListStubWithDefaults() *CSIVolumeListStub`

NewCSIVolumeListStubWithDefaults instantiates a new CSIVolumeListStub object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessMode

`func (o *CSIVolumeListStub) GetAccessMode() string`

GetAccessMode returns the AccessMode field if non-nil, zero value otherwise.

### GetAccessModeOk

`func (o *CSIVolumeListStub) GetAccessModeOk() (*string, bool)`

GetAccessModeOk returns a tuple with the AccessMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessMode

`func (o *CSIVolumeListStub) SetAccessMode(v string)`

SetAccessMode sets AccessMode field to given value.

### HasAccessMode

`func (o *CSIVolumeListStub) HasAccessMode() bool`

HasAccessMode returns a boolean if a field has been set.

### GetAttachmentMode

`func (o *CSIVolumeListStub) GetAttachmentMode() string`

GetAttachmentMode returns the AttachmentMode field if non-nil, zero value otherwise.

### GetAttachmentModeOk

`func (o *CSIVolumeListStub) GetAttachmentModeOk() (*string, bool)`

GetAttachmentModeOk returns a tuple with the AttachmentMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentMode

`func (o *CSIVolumeListStub) SetAttachmentMode(v string)`

SetAttachmentMode sets AttachmentMode field to given value.

### HasAttachmentMode

`func (o *CSIVolumeListStub) HasAttachmentMode() bool`

HasAttachmentMode returns a boolean if a field has been set.

### GetControllerRequired

`func (o *CSIVolumeListStub) GetControllerRequired() bool`

GetControllerRequired returns the ControllerRequired field if non-nil, zero value otherwise.

### GetControllerRequiredOk

`func (o *CSIVolumeListStub) GetControllerRequiredOk() (*bool, bool)`

GetControllerRequiredOk returns a tuple with the ControllerRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerRequired

`func (o *CSIVolumeListStub) SetControllerRequired(v bool)`

SetControllerRequired sets ControllerRequired field to given value.

### HasControllerRequired

`func (o *CSIVolumeListStub) HasControllerRequired() bool`

HasControllerRequired returns a boolean if a field has been set.

### GetControllersExpected

`func (o *CSIVolumeListStub) GetControllersExpected() int32`

GetControllersExpected returns the ControllersExpected field if non-nil, zero value otherwise.

### GetControllersExpectedOk

`func (o *CSIVolumeListStub) GetControllersExpectedOk() (*int32, bool)`

GetControllersExpectedOk returns a tuple with the ControllersExpected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllersExpected

`func (o *CSIVolumeListStub) SetControllersExpected(v int32)`

SetControllersExpected sets ControllersExpected field to given value.

### HasControllersExpected

`func (o *CSIVolumeListStub) HasControllersExpected() bool`

HasControllersExpected returns a boolean if a field has been set.

### GetControllersHealthy

`func (o *CSIVolumeListStub) GetControllersHealthy() int32`

GetControllersHealthy returns the ControllersHealthy field if non-nil, zero value otherwise.

### GetControllersHealthyOk

`func (o *CSIVolumeListStub) GetControllersHealthyOk() (*int32, bool)`

GetControllersHealthyOk returns a tuple with the ControllersHealthy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllersHealthy

`func (o *CSIVolumeListStub) SetControllersHealthy(v int32)`

SetControllersHealthy sets ControllersHealthy field to given value.

### HasControllersHealthy

`func (o *CSIVolumeListStub) HasControllersHealthy() bool`

HasControllersHealthy returns a boolean if a field has been set.

### GetCreateIndex

`func (o *CSIVolumeListStub) GetCreateIndex() int32`

GetCreateIndex returns the CreateIndex field if non-nil, zero value otherwise.

### GetCreateIndexOk

`func (o *CSIVolumeListStub) GetCreateIndexOk() (*int32, bool)`

GetCreateIndexOk returns a tuple with the CreateIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateIndex

`func (o *CSIVolumeListStub) SetCreateIndex(v int32)`

SetCreateIndex sets CreateIndex field to given value.

### HasCreateIndex

`func (o *CSIVolumeListStub) HasCreateIndex() bool`

HasCreateIndex returns a boolean if a field has been set.

### GetExternalID

`func (o *CSIVolumeListStub) GetExternalID() string`

GetExternalID returns the ExternalID field if non-nil, zero value otherwise.

### GetExternalIDOk

`func (o *CSIVolumeListStub) GetExternalIDOk() (*string, bool)`

GetExternalIDOk returns a tuple with the ExternalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalID

`func (o *CSIVolumeListStub) SetExternalID(v string)`

SetExternalID sets ExternalID field to given value.

### HasExternalID

`func (o *CSIVolumeListStub) HasExternalID() bool`

HasExternalID returns a boolean if a field has been set.

### GetID

`func (o *CSIVolumeListStub) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *CSIVolumeListStub) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *CSIVolumeListStub) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *CSIVolumeListStub) HasID() bool`

HasID returns a boolean if a field has been set.

### GetModifyIndex

`func (o *CSIVolumeListStub) GetModifyIndex() int32`

GetModifyIndex returns the ModifyIndex field if non-nil, zero value otherwise.

### GetModifyIndexOk

`func (o *CSIVolumeListStub) GetModifyIndexOk() (*int32, bool)`

GetModifyIndexOk returns a tuple with the ModifyIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifyIndex

`func (o *CSIVolumeListStub) SetModifyIndex(v int32)`

SetModifyIndex sets ModifyIndex field to given value.

### HasModifyIndex

`func (o *CSIVolumeListStub) HasModifyIndex() bool`

HasModifyIndex returns a boolean if a field has been set.

### GetName

`func (o *CSIVolumeListStub) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CSIVolumeListStub) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CSIVolumeListStub) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CSIVolumeListStub) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *CSIVolumeListStub) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *CSIVolumeListStub) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *CSIVolumeListStub) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *CSIVolumeListStub) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetNodesExpected

`func (o *CSIVolumeListStub) GetNodesExpected() int32`

GetNodesExpected returns the NodesExpected field if non-nil, zero value otherwise.

### GetNodesExpectedOk

`func (o *CSIVolumeListStub) GetNodesExpectedOk() (*int32, bool)`

GetNodesExpectedOk returns a tuple with the NodesExpected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodesExpected

`func (o *CSIVolumeListStub) SetNodesExpected(v int32)`

SetNodesExpected sets NodesExpected field to given value.

### HasNodesExpected

`func (o *CSIVolumeListStub) HasNodesExpected() bool`

HasNodesExpected returns a boolean if a field has been set.

### GetNodesHealthy

`func (o *CSIVolumeListStub) GetNodesHealthy() int32`

GetNodesHealthy returns the NodesHealthy field if non-nil, zero value otherwise.

### GetNodesHealthyOk

`func (o *CSIVolumeListStub) GetNodesHealthyOk() (*int32, bool)`

GetNodesHealthyOk returns a tuple with the NodesHealthy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodesHealthy

`func (o *CSIVolumeListStub) SetNodesHealthy(v int32)`

SetNodesHealthy sets NodesHealthy field to given value.

### HasNodesHealthy

`func (o *CSIVolumeListStub) HasNodesHealthy() bool`

HasNodesHealthy returns a boolean if a field has been set.

### GetPluginID

`func (o *CSIVolumeListStub) GetPluginID() string`

GetPluginID returns the PluginID field if non-nil, zero value otherwise.

### GetPluginIDOk

`func (o *CSIVolumeListStub) GetPluginIDOk() (*string, bool)`

GetPluginIDOk returns a tuple with the PluginID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginID

`func (o *CSIVolumeListStub) SetPluginID(v string)`

SetPluginID sets PluginID field to given value.

### HasPluginID

`func (o *CSIVolumeListStub) HasPluginID() bool`

HasPluginID returns a boolean if a field has been set.

### GetProvider

`func (o *CSIVolumeListStub) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CSIVolumeListStub) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CSIVolumeListStub) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *CSIVolumeListStub) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetResourceExhausted

`func (o *CSIVolumeListStub) GetResourceExhausted() time.Time`

GetResourceExhausted returns the ResourceExhausted field if non-nil, zero value otherwise.

### GetResourceExhaustedOk

`func (o *CSIVolumeListStub) GetResourceExhaustedOk() (*time.Time, bool)`

GetResourceExhaustedOk returns a tuple with the ResourceExhausted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceExhausted

`func (o *CSIVolumeListStub) SetResourceExhausted(v time.Time)`

SetResourceExhausted sets ResourceExhausted field to given value.

### HasResourceExhausted

`func (o *CSIVolumeListStub) HasResourceExhausted() bool`

HasResourceExhausted returns a boolean if a field has been set.

### SetResourceExhaustedNil

`func (o *CSIVolumeListStub) SetResourceExhaustedNil(b bool)`

 SetResourceExhaustedNil sets the value for ResourceExhausted to be an explicit nil

### UnsetResourceExhausted
`func (o *CSIVolumeListStub) UnsetResourceExhausted()`

UnsetResourceExhausted ensures that no value is present for ResourceExhausted, not even an explicit nil
### GetSchedulable

`func (o *CSIVolumeListStub) GetSchedulable() bool`

GetSchedulable returns the Schedulable field if non-nil, zero value otherwise.

### GetSchedulableOk

`func (o *CSIVolumeListStub) GetSchedulableOk() (*bool, bool)`

GetSchedulableOk returns a tuple with the Schedulable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedulable

`func (o *CSIVolumeListStub) SetSchedulable(v bool)`

SetSchedulable sets Schedulable field to given value.

### HasSchedulable

`func (o *CSIVolumeListStub) HasSchedulable() bool`

HasSchedulable returns a boolean if a field has been set.

### GetTopologies

`func (o *CSIVolumeListStub) GetTopologies() []CSITopology`

GetTopologies returns the Topologies field if non-nil, zero value otherwise.

### GetTopologiesOk

`func (o *CSIVolumeListStub) GetTopologiesOk() (*[]CSITopology, bool)`

GetTopologiesOk returns a tuple with the Topologies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopologies

`func (o *CSIVolumeListStub) SetTopologies(v []CSITopology)`

SetTopologies sets Topologies field to given value.

### HasTopologies

`func (o *CSIVolumeListStub) HasTopologies() bool`

HasTopologies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


