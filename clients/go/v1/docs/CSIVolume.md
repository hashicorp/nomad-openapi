# CSIVolume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessMode** | Pointer to **string** |  | [optional] 
**Allocations** | Pointer to [**[]AllocationListStub**](AllocationListStub.md) |  | [optional] 
**AttachmentMode** | Pointer to **string** |  | [optional] 
**Capacity** | Pointer to **int64** |  | [optional] 
**CloneID** | Pointer to **string** |  | [optional] 
**Context** | Pointer to **map[string]string** |  | [optional] 
**ControllerRequired** | Pointer to **bool** |  | [optional] 
**ControllersExpected** | Pointer to **int32** |  | [optional] 
**ControllersHealthy** | Pointer to **int32** |  | [optional] 
**CreateIndex** | Pointer to **int32** |  | [optional] 
**ExternalID** | Pointer to **string** |  | [optional] 
**ID** | Pointer to **string** |  | [optional] 
**ModifyIndex** | Pointer to **int32** |  | [optional] 
**MountOptions** | Pointer to [**NullableCSIMountOptions**](CSIMountOptions.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**NodesExpected** | Pointer to **int32** |  | [optional] 
**NodesHealthy** | Pointer to **int32** |  | [optional] 
**Parameters** | Pointer to **map[string]string** |  | [optional] 
**PluginID** | Pointer to **string** |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**ProviderVersion** | Pointer to **string** |  | [optional] 
**ReadAllocs** | Pointer to [**map[string]Allocation**](Allocation.md) |  | [optional] 
**RequestedCapabilities** | Pointer to [**[]CSIVolumeCapability**](CSIVolumeCapability.md) |  | [optional] 
**RequestedCapacityMax** | Pointer to **int64** |  | [optional] 
**RequestedCapacityMin** | Pointer to **int64** |  | [optional] 
**RequestedTopologies** | Pointer to [**NullableCSITopologyRequest**](CSITopologyRequest.md) |  | [optional] 
**ResourceExhausted** | Pointer to **NullableTime** |  | [optional] 
**Schedulable** | Pointer to **bool** |  | [optional] 
**Secrets** | Pointer to **map[string]string** |  | [optional] 
**SnapshotID** | Pointer to **string** |  | [optional] 
**Topologies** | Pointer to [**[]CSITopology**](CSITopology.md) |  | [optional] 
**WriteAllocs** | Pointer to [**map[string]Allocation**](Allocation.md) |  | [optional] 

## Methods

### NewCSIVolume

`func NewCSIVolume() *CSIVolume`

NewCSIVolume instantiates a new CSIVolume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSIVolumeWithDefaults

`func NewCSIVolumeWithDefaults() *CSIVolume`

NewCSIVolumeWithDefaults instantiates a new CSIVolume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessMode

`func (o *CSIVolume) GetAccessMode() string`

GetAccessMode returns the AccessMode field if non-nil, zero value otherwise.

### GetAccessModeOk

`func (o *CSIVolume) GetAccessModeOk() (*string, bool)`

GetAccessModeOk returns a tuple with the AccessMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessMode

`func (o *CSIVolume) SetAccessMode(v string)`

SetAccessMode sets AccessMode field to given value.

### HasAccessMode

`func (o *CSIVolume) HasAccessMode() bool`

HasAccessMode returns a boolean if a field has been set.

### GetAllocations

`func (o *CSIVolume) GetAllocations() []AllocationListStub`

GetAllocations returns the Allocations field if non-nil, zero value otherwise.

### GetAllocationsOk

`func (o *CSIVolume) GetAllocationsOk() (*[]AllocationListStub, bool)`

GetAllocationsOk returns a tuple with the Allocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocations

`func (o *CSIVolume) SetAllocations(v []AllocationListStub)`

SetAllocations sets Allocations field to given value.

### HasAllocations

`func (o *CSIVolume) HasAllocations() bool`

HasAllocations returns a boolean if a field has been set.

### GetAttachmentMode

`func (o *CSIVolume) GetAttachmentMode() string`

GetAttachmentMode returns the AttachmentMode field if non-nil, zero value otherwise.

### GetAttachmentModeOk

`func (o *CSIVolume) GetAttachmentModeOk() (*string, bool)`

GetAttachmentModeOk returns a tuple with the AttachmentMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentMode

`func (o *CSIVolume) SetAttachmentMode(v string)`

SetAttachmentMode sets AttachmentMode field to given value.

### HasAttachmentMode

`func (o *CSIVolume) HasAttachmentMode() bool`

HasAttachmentMode returns a boolean if a field has been set.

### GetCapacity

`func (o *CSIVolume) GetCapacity() int64`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *CSIVolume) GetCapacityOk() (*int64, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *CSIVolume) SetCapacity(v int64)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *CSIVolume) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetCloneID

`func (o *CSIVolume) GetCloneID() string`

GetCloneID returns the CloneID field if non-nil, zero value otherwise.

### GetCloneIDOk

`func (o *CSIVolume) GetCloneIDOk() (*string, bool)`

GetCloneIDOk returns a tuple with the CloneID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloneID

`func (o *CSIVolume) SetCloneID(v string)`

SetCloneID sets CloneID field to given value.

### HasCloneID

`func (o *CSIVolume) HasCloneID() bool`

HasCloneID returns a boolean if a field has been set.

### GetContext

`func (o *CSIVolume) GetContext() map[string]string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *CSIVolume) GetContextOk() (*map[string]string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *CSIVolume) SetContext(v map[string]string)`

SetContext sets Context field to given value.

### HasContext

`func (o *CSIVolume) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetControllerRequired

`func (o *CSIVolume) GetControllerRequired() bool`

GetControllerRequired returns the ControllerRequired field if non-nil, zero value otherwise.

### GetControllerRequiredOk

`func (o *CSIVolume) GetControllerRequiredOk() (*bool, bool)`

GetControllerRequiredOk returns a tuple with the ControllerRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerRequired

`func (o *CSIVolume) SetControllerRequired(v bool)`

SetControllerRequired sets ControllerRequired field to given value.

### HasControllerRequired

`func (o *CSIVolume) HasControllerRequired() bool`

HasControllerRequired returns a boolean if a field has been set.

### GetControllersExpected

`func (o *CSIVolume) GetControllersExpected() int32`

GetControllersExpected returns the ControllersExpected field if non-nil, zero value otherwise.

### GetControllersExpectedOk

`func (o *CSIVolume) GetControllersExpectedOk() (*int32, bool)`

GetControllersExpectedOk returns a tuple with the ControllersExpected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllersExpected

`func (o *CSIVolume) SetControllersExpected(v int32)`

SetControllersExpected sets ControllersExpected field to given value.

### HasControllersExpected

`func (o *CSIVolume) HasControllersExpected() bool`

HasControllersExpected returns a boolean if a field has been set.

### GetControllersHealthy

`func (o *CSIVolume) GetControllersHealthy() int32`

GetControllersHealthy returns the ControllersHealthy field if non-nil, zero value otherwise.

### GetControllersHealthyOk

`func (o *CSIVolume) GetControllersHealthyOk() (*int32, bool)`

GetControllersHealthyOk returns a tuple with the ControllersHealthy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllersHealthy

`func (o *CSIVolume) SetControllersHealthy(v int32)`

SetControllersHealthy sets ControllersHealthy field to given value.

### HasControllersHealthy

`func (o *CSIVolume) HasControllersHealthy() bool`

HasControllersHealthy returns a boolean if a field has been set.

### GetCreateIndex

`func (o *CSIVolume) GetCreateIndex() int32`

GetCreateIndex returns the CreateIndex field if non-nil, zero value otherwise.

### GetCreateIndexOk

`func (o *CSIVolume) GetCreateIndexOk() (*int32, bool)`

GetCreateIndexOk returns a tuple with the CreateIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateIndex

`func (o *CSIVolume) SetCreateIndex(v int32)`

SetCreateIndex sets CreateIndex field to given value.

### HasCreateIndex

`func (o *CSIVolume) HasCreateIndex() bool`

HasCreateIndex returns a boolean if a field has been set.

### GetExternalID

`func (o *CSIVolume) GetExternalID() string`

GetExternalID returns the ExternalID field if non-nil, zero value otherwise.

### GetExternalIDOk

`func (o *CSIVolume) GetExternalIDOk() (*string, bool)`

GetExternalIDOk returns a tuple with the ExternalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalID

`func (o *CSIVolume) SetExternalID(v string)`

SetExternalID sets ExternalID field to given value.

### HasExternalID

`func (o *CSIVolume) HasExternalID() bool`

HasExternalID returns a boolean if a field has been set.

### GetID

`func (o *CSIVolume) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *CSIVolume) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *CSIVolume) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *CSIVolume) HasID() bool`

HasID returns a boolean if a field has been set.

### GetModifyIndex

`func (o *CSIVolume) GetModifyIndex() int32`

GetModifyIndex returns the ModifyIndex field if non-nil, zero value otherwise.

### GetModifyIndexOk

`func (o *CSIVolume) GetModifyIndexOk() (*int32, bool)`

GetModifyIndexOk returns a tuple with the ModifyIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifyIndex

`func (o *CSIVolume) SetModifyIndex(v int32)`

SetModifyIndex sets ModifyIndex field to given value.

### HasModifyIndex

`func (o *CSIVolume) HasModifyIndex() bool`

HasModifyIndex returns a boolean if a field has been set.

### GetMountOptions

`func (o *CSIVolume) GetMountOptions() CSIMountOptions`

GetMountOptions returns the MountOptions field if non-nil, zero value otherwise.

### GetMountOptionsOk

`func (o *CSIVolume) GetMountOptionsOk() (*CSIMountOptions, bool)`

GetMountOptionsOk returns a tuple with the MountOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountOptions

`func (o *CSIVolume) SetMountOptions(v CSIMountOptions)`

SetMountOptions sets MountOptions field to given value.

### HasMountOptions

`func (o *CSIVolume) HasMountOptions() bool`

HasMountOptions returns a boolean if a field has been set.

### SetMountOptionsNil

`func (o *CSIVolume) SetMountOptionsNil(b bool)`

 SetMountOptionsNil sets the value for MountOptions to be an explicit nil

### UnsetMountOptions
`func (o *CSIVolume) UnsetMountOptions()`

UnsetMountOptions ensures that no value is present for MountOptions, not even an explicit nil
### GetName

`func (o *CSIVolume) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CSIVolume) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CSIVolume) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CSIVolume) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *CSIVolume) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *CSIVolume) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *CSIVolume) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *CSIVolume) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetNodesExpected

`func (o *CSIVolume) GetNodesExpected() int32`

GetNodesExpected returns the NodesExpected field if non-nil, zero value otherwise.

### GetNodesExpectedOk

`func (o *CSIVolume) GetNodesExpectedOk() (*int32, bool)`

GetNodesExpectedOk returns a tuple with the NodesExpected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodesExpected

`func (o *CSIVolume) SetNodesExpected(v int32)`

SetNodesExpected sets NodesExpected field to given value.

### HasNodesExpected

`func (o *CSIVolume) HasNodesExpected() bool`

HasNodesExpected returns a boolean if a field has been set.

### GetNodesHealthy

`func (o *CSIVolume) GetNodesHealthy() int32`

GetNodesHealthy returns the NodesHealthy field if non-nil, zero value otherwise.

### GetNodesHealthyOk

`func (o *CSIVolume) GetNodesHealthyOk() (*int32, bool)`

GetNodesHealthyOk returns a tuple with the NodesHealthy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodesHealthy

`func (o *CSIVolume) SetNodesHealthy(v int32)`

SetNodesHealthy sets NodesHealthy field to given value.

### HasNodesHealthy

`func (o *CSIVolume) HasNodesHealthy() bool`

HasNodesHealthy returns a boolean if a field has been set.

### GetParameters

`func (o *CSIVolume) GetParameters() map[string]string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *CSIVolume) GetParametersOk() (*map[string]string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *CSIVolume) SetParameters(v map[string]string)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *CSIVolume) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetPluginID

`func (o *CSIVolume) GetPluginID() string`

GetPluginID returns the PluginID field if non-nil, zero value otherwise.

### GetPluginIDOk

`func (o *CSIVolume) GetPluginIDOk() (*string, bool)`

GetPluginIDOk returns a tuple with the PluginID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginID

`func (o *CSIVolume) SetPluginID(v string)`

SetPluginID sets PluginID field to given value.

### HasPluginID

`func (o *CSIVolume) HasPluginID() bool`

HasPluginID returns a boolean if a field has been set.

### GetProvider

`func (o *CSIVolume) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CSIVolume) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CSIVolume) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *CSIVolume) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetProviderVersion

`func (o *CSIVolume) GetProviderVersion() string`

GetProviderVersion returns the ProviderVersion field if non-nil, zero value otherwise.

### GetProviderVersionOk

`func (o *CSIVolume) GetProviderVersionOk() (*string, bool)`

GetProviderVersionOk returns a tuple with the ProviderVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderVersion

`func (o *CSIVolume) SetProviderVersion(v string)`

SetProviderVersion sets ProviderVersion field to given value.

### HasProviderVersion

`func (o *CSIVolume) HasProviderVersion() bool`

HasProviderVersion returns a boolean if a field has been set.

### GetReadAllocs

`func (o *CSIVolume) GetReadAllocs() map[string]Allocation`

GetReadAllocs returns the ReadAllocs field if non-nil, zero value otherwise.

### GetReadAllocsOk

`func (o *CSIVolume) GetReadAllocsOk() (*map[string]Allocation, bool)`

GetReadAllocsOk returns a tuple with the ReadAllocs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadAllocs

`func (o *CSIVolume) SetReadAllocs(v map[string]Allocation)`

SetReadAllocs sets ReadAllocs field to given value.

### HasReadAllocs

`func (o *CSIVolume) HasReadAllocs() bool`

HasReadAllocs returns a boolean if a field has been set.

### GetRequestedCapabilities

`func (o *CSIVolume) GetRequestedCapabilities() []CSIVolumeCapability`

GetRequestedCapabilities returns the RequestedCapabilities field if non-nil, zero value otherwise.

### GetRequestedCapabilitiesOk

`func (o *CSIVolume) GetRequestedCapabilitiesOk() (*[]CSIVolumeCapability, bool)`

GetRequestedCapabilitiesOk returns a tuple with the RequestedCapabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedCapabilities

`func (o *CSIVolume) SetRequestedCapabilities(v []CSIVolumeCapability)`

SetRequestedCapabilities sets RequestedCapabilities field to given value.

### HasRequestedCapabilities

`func (o *CSIVolume) HasRequestedCapabilities() bool`

HasRequestedCapabilities returns a boolean if a field has been set.

### GetRequestedCapacityMax

`func (o *CSIVolume) GetRequestedCapacityMax() int64`

GetRequestedCapacityMax returns the RequestedCapacityMax field if non-nil, zero value otherwise.

### GetRequestedCapacityMaxOk

`func (o *CSIVolume) GetRequestedCapacityMaxOk() (*int64, bool)`

GetRequestedCapacityMaxOk returns a tuple with the RequestedCapacityMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedCapacityMax

`func (o *CSIVolume) SetRequestedCapacityMax(v int64)`

SetRequestedCapacityMax sets RequestedCapacityMax field to given value.

### HasRequestedCapacityMax

`func (o *CSIVolume) HasRequestedCapacityMax() bool`

HasRequestedCapacityMax returns a boolean if a field has been set.

### GetRequestedCapacityMin

`func (o *CSIVolume) GetRequestedCapacityMin() int64`

GetRequestedCapacityMin returns the RequestedCapacityMin field if non-nil, zero value otherwise.

### GetRequestedCapacityMinOk

`func (o *CSIVolume) GetRequestedCapacityMinOk() (*int64, bool)`

GetRequestedCapacityMinOk returns a tuple with the RequestedCapacityMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedCapacityMin

`func (o *CSIVolume) SetRequestedCapacityMin(v int64)`

SetRequestedCapacityMin sets RequestedCapacityMin field to given value.

### HasRequestedCapacityMin

`func (o *CSIVolume) HasRequestedCapacityMin() bool`

HasRequestedCapacityMin returns a boolean if a field has been set.

### GetRequestedTopologies

`func (o *CSIVolume) GetRequestedTopologies() CSITopologyRequest`

GetRequestedTopologies returns the RequestedTopologies field if non-nil, zero value otherwise.

### GetRequestedTopologiesOk

`func (o *CSIVolume) GetRequestedTopologiesOk() (*CSITopologyRequest, bool)`

GetRequestedTopologiesOk returns a tuple with the RequestedTopologies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedTopologies

`func (o *CSIVolume) SetRequestedTopologies(v CSITopologyRequest)`

SetRequestedTopologies sets RequestedTopologies field to given value.

### HasRequestedTopologies

`func (o *CSIVolume) HasRequestedTopologies() bool`

HasRequestedTopologies returns a boolean if a field has been set.

### SetRequestedTopologiesNil

`func (o *CSIVolume) SetRequestedTopologiesNil(b bool)`

 SetRequestedTopologiesNil sets the value for RequestedTopologies to be an explicit nil

### UnsetRequestedTopologies
`func (o *CSIVolume) UnsetRequestedTopologies()`

UnsetRequestedTopologies ensures that no value is present for RequestedTopologies, not even an explicit nil
### GetResourceExhausted

`func (o *CSIVolume) GetResourceExhausted() time.Time`

GetResourceExhausted returns the ResourceExhausted field if non-nil, zero value otherwise.

### GetResourceExhaustedOk

`func (o *CSIVolume) GetResourceExhaustedOk() (*time.Time, bool)`

GetResourceExhaustedOk returns a tuple with the ResourceExhausted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceExhausted

`func (o *CSIVolume) SetResourceExhausted(v time.Time)`

SetResourceExhausted sets ResourceExhausted field to given value.

### HasResourceExhausted

`func (o *CSIVolume) HasResourceExhausted() bool`

HasResourceExhausted returns a boolean if a field has been set.

### SetResourceExhaustedNil

`func (o *CSIVolume) SetResourceExhaustedNil(b bool)`

 SetResourceExhaustedNil sets the value for ResourceExhausted to be an explicit nil

### UnsetResourceExhausted
`func (o *CSIVolume) UnsetResourceExhausted()`

UnsetResourceExhausted ensures that no value is present for ResourceExhausted, not even an explicit nil
### GetSchedulable

`func (o *CSIVolume) GetSchedulable() bool`

GetSchedulable returns the Schedulable field if non-nil, zero value otherwise.

### GetSchedulableOk

`func (o *CSIVolume) GetSchedulableOk() (*bool, bool)`

GetSchedulableOk returns a tuple with the Schedulable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedulable

`func (o *CSIVolume) SetSchedulable(v bool)`

SetSchedulable sets Schedulable field to given value.

### HasSchedulable

`func (o *CSIVolume) HasSchedulable() bool`

HasSchedulable returns a boolean if a field has been set.

### GetSecrets

`func (o *CSIVolume) GetSecrets() map[string]string`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *CSIVolume) GetSecretsOk() (*map[string]string, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *CSIVolume) SetSecrets(v map[string]string)`

SetSecrets sets Secrets field to given value.

### HasSecrets

`func (o *CSIVolume) HasSecrets() bool`

HasSecrets returns a boolean if a field has been set.

### GetSnapshotID

`func (o *CSIVolume) GetSnapshotID() string`

GetSnapshotID returns the SnapshotID field if non-nil, zero value otherwise.

### GetSnapshotIDOk

`func (o *CSIVolume) GetSnapshotIDOk() (*string, bool)`

GetSnapshotIDOk returns a tuple with the SnapshotID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotID

`func (o *CSIVolume) SetSnapshotID(v string)`

SetSnapshotID sets SnapshotID field to given value.

### HasSnapshotID

`func (o *CSIVolume) HasSnapshotID() bool`

HasSnapshotID returns a boolean if a field has been set.

### GetTopologies

`func (o *CSIVolume) GetTopologies() []CSITopology`

GetTopologies returns the Topologies field if non-nil, zero value otherwise.

### GetTopologiesOk

`func (o *CSIVolume) GetTopologiesOk() (*[]CSITopology, bool)`

GetTopologiesOk returns a tuple with the Topologies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopologies

`func (o *CSIVolume) SetTopologies(v []CSITopology)`

SetTopologies sets Topologies field to given value.

### HasTopologies

`func (o *CSIVolume) HasTopologies() bool`

HasTopologies returns a boolean if a field has been set.

### GetWriteAllocs

`func (o *CSIVolume) GetWriteAllocs() map[string]Allocation`

GetWriteAllocs returns the WriteAllocs field if non-nil, zero value otherwise.

### GetWriteAllocsOk

`func (o *CSIVolume) GetWriteAllocsOk() (*map[string]Allocation, bool)`

GetWriteAllocsOk returns a tuple with the WriteAllocs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWriteAllocs

`func (o *CSIVolume) SetWriteAllocs(v map[string]Allocation)`

SetWriteAllocs sets WriteAllocs field to given value.

### HasWriteAllocs

`func (o *CSIVolume) HasWriteAllocs() bool`

HasWriteAllocs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


