# CSIControllerInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SupportsAttachDetach** | Pointer to **bool** |  | [optional] 
**SupportsClone** | Pointer to **bool** |  | [optional] 
**SupportsCondition** | Pointer to **bool** |  | [optional] 
**SupportsCreateDelete** | Pointer to **bool** |  | [optional] 
**SupportsCreateDeleteSnapshot** | Pointer to **bool** |  | [optional] 
**SupportsExpand** | Pointer to **bool** |  | [optional] 
**SupportsGet** | Pointer to **bool** |  | [optional] 
**SupportsGetCapacity** | Pointer to **bool** |  | [optional] 
**SupportsListSnapshots** | Pointer to **bool** |  | [optional] 
**SupportsListVolumes** | Pointer to **bool** |  | [optional] 
**SupportsListVolumesAttachedNodes** | Pointer to **bool** |  | [optional] 
**SupportsReadOnlyAttach** | Pointer to **bool** |  | [optional] 

## Methods

### NewCSIControllerInfo

`func NewCSIControllerInfo() *CSIControllerInfo`

NewCSIControllerInfo instantiates a new CSIControllerInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSIControllerInfoWithDefaults

`func NewCSIControllerInfoWithDefaults() *CSIControllerInfo`

NewCSIControllerInfoWithDefaults instantiates a new CSIControllerInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupportsAttachDetach

`func (o *CSIControllerInfo) GetSupportsAttachDetach() bool`

GetSupportsAttachDetach returns the SupportsAttachDetach field if non-nil, zero value otherwise.

### GetSupportsAttachDetachOk

`func (o *CSIControllerInfo) GetSupportsAttachDetachOk() (*bool, bool)`

GetSupportsAttachDetachOk returns a tuple with the SupportsAttachDetach field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsAttachDetach

`func (o *CSIControllerInfo) SetSupportsAttachDetach(v bool)`

SetSupportsAttachDetach sets SupportsAttachDetach field to given value.

### HasSupportsAttachDetach

`func (o *CSIControllerInfo) HasSupportsAttachDetach() bool`

HasSupportsAttachDetach returns a boolean if a field has been set.

### GetSupportsClone

`func (o *CSIControllerInfo) GetSupportsClone() bool`

GetSupportsClone returns the SupportsClone field if non-nil, zero value otherwise.

### GetSupportsCloneOk

`func (o *CSIControllerInfo) GetSupportsCloneOk() (*bool, bool)`

GetSupportsCloneOk returns a tuple with the SupportsClone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsClone

`func (o *CSIControllerInfo) SetSupportsClone(v bool)`

SetSupportsClone sets SupportsClone field to given value.

### HasSupportsClone

`func (o *CSIControllerInfo) HasSupportsClone() bool`

HasSupportsClone returns a boolean if a field has been set.

### GetSupportsCondition

`func (o *CSIControllerInfo) GetSupportsCondition() bool`

GetSupportsCondition returns the SupportsCondition field if non-nil, zero value otherwise.

### GetSupportsConditionOk

`func (o *CSIControllerInfo) GetSupportsConditionOk() (*bool, bool)`

GetSupportsConditionOk returns a tuple with the SupportsCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsCondition

`func (o *CSIControllerInfo) SetSupportsCondition(v bool)`

SetSupportsCondition sets SupportsCondition field to given value.

### HasSupportsCondition

`func (o *CSIControllerInfo) HasSupportsCondition() bool`

HasSupportsCondition returns a boolean if a field has been set.

### GetSupportsCreateDelete

`func (o *CSIControllerInfo) GetSupportsCreateDelete() bool`

GetSupportsCreateDelete returns the SupportsCreateDelete field if non-nil, zero value otherwise.

### GetSupportsCreateDeleteOk

`func (o *CSIControllerInfo) GetSupportsCreateDeleteOk() (*bool, bool)`

GetSupportsCreateDeleteOk returns a tuple with the SupportsCreateDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsCreateDelete

`func (o *CSIControllerInfo) SetSupportsCreateDelete(v bool)`

SetSupportsCreateDelete sets SupportsCreateDelete field to given value.

### HasSupportsCreateDelete

`func (o *CSIControllerInfo) HasSupportsCreateDelete() bool`

HasSupportsCreateDelete returns a boolean if a field has been set.

### GetSupportsCreateDeleteSnapshot

`func (o *CSIControllerInfo) GetSupportsCreateDeleteSnapshot() bool`

GetSupportsCreateDeleteSnapshot returns the SupportsCreateDeleteSnapshot field if non-nil, zero value otherwise.

### GetSupportsCreateDeleteSnapshotOk

`func (o *CSIControllerInfo) GetSupportsCreateDeleteSnapshotOk() (*bool, bool)`

GetSupportsCreateDeleteSnapshotOk returns a tuple with the SupportsCreateDeleteSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsCreateDeleteSnapshot

`func (o *CSIControllerInfo) SetSupportsCreateDeleteSnapshot(v bool)`

SetSupportsCreateDeleteSnapshot sets SupportsCreateDeleteSnapshot field to given value.

### HasSupportsCreateDeleteSnapshot

`func (o *CSIControllerInfo) HasSupportsCreateDeleteSnapshot() bool`

HasSupportsCreateDeleteSnapshot returns a boolean if a field has been set.

### GetSupportsExpand

`func (o *CSIControllerInfo) GetSupportsExpand() bool`

GetSupportsExpand returns the SupportsExpand field if non-nil, zero value otherwise.

### GetSupportsExpandOk

`func (o *CSIControllerInfo) GetSupportsExpandOk() (*bool, bool)`

GetSupportsExpandOk returns a tuple with the SupportsExpand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsExpand

`func (o *CSIControllerInfo) SetSupportsExpand(v bool)`

SetSupportsExpand sets SupportsExpand field to given value.

### HasSupportsExpand

`func (o *CSIControllerInfo) HasSupportsExpand() bool`

HasSupportsExpand returns a boolean if a field has been set.

### GetSupportsGet

`func (o *CSIControllerInfo) GetSupportsGet() bool`

GetSupportsGet returns the SupportsGet field if non-nil, zero value otherwise.

### GetSupportsGetOk

`func (o *CSIControllerInfo) GetSupportsGetOk() (*bool, bool)`

GetSupportsGetOk returns a tuple with the SupportsGet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsGet

`func (o *CSIControllerInfo) SetSupportsGet(v bool)`

SetSupportsGet sets SupportsGet field to given value.

### HasSupportsGet

`func (o *CSIControllerInfo) HasSupportsGet() bool`

HasSupportsGet returns a boolean if a field has been set.

### GetSupportsGetCapacity

`func (o *CSIControllerInfo) GetSupportsGetCapacity() bool`

GetSupportsGetCapacity returns the SupportsGetCapacity field if non-nil, zero value otherwise.

### GetSupportsGetCapacityOk

`func (o *CSIControllerInfo) GetSupportsGetCapacityOk() (*bool, bool)`

GetSupportsGetCapacityOk returns a tuple with the SupportsGetCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsGetCapacity

`func (o *CSIControllerInfo) SetSupportsGetCapacity(v bool)`

SetSupportsGetCapacity sets SupportsGetCapacity field to given value.

### HasSupportsGetCapacity

`func (o *CSIControllerInfo) HasSupportsGetCapacity() bool`

HasSupportsGetCapacity returns a boolean if a field has been set.

### GetSupportsListSnapshots

`func (o *CSIControllerInfo) GetSupportsListSnapshots() bool`

GetSupportsListSnapshots returns the SupportsListSnapshots field if non-nil, zero value otherwise.

### GetSupportsListSnapshotsOk

`func (o *CSIControllerInfo) GetSupportsListSnapshotsOk() (*bool, bool)`

GetSupportsListSnapshotsOk returns a tuple with the SupportsListSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsListSnapshots

`func (o *CSIControllerInfo) SetSupportsListSnapshots(v bool)`

SetSupportsListSnapshots sets SupportsListSnapshots field to given value.

### HasSupportsListSnapshots

`func (o *CSIControllerInfo) HasSupportsListSnapshots() bool`

HasSupportsListSnapshots returns a boolean if a field has been set.

### GetSupportsListVolumes

`func (o *CSIControllerInfo) GetSupportsListVolumes() bool`

GetSupportsListVolumes returns the SupportsListVolumes field if non-nil, zero value otherwise.

### GetSupportsListVolumesOk

`func (o *CSIControllerInfo) GetSupportsListVolumesOk() (*bool, bool)`

GetSupportsListVolumesOk returns a tuple with the SupportsListVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsListVolumes

`func (o *CSIControllerInfo) SetSupportsListVolumes(v bool)`

SetSupportsListVolumes sets SupportsListVolumes field to given value.

### HasSupportsListVolumes

`func (o *CSIControllerInfo) HasSupportsListVolumes() bool`

HasSupportsListVolumes returns a boolean if a field has been set.

### GetSupportsListVolumesAttachedNodes

`func (o *CSIControllerInfo) GetSupportsListVolumesAttachedNodes() bool`

GetSupportsListVolumesAttachedNodes returns the SupportsListVolumesAttachedNodes field if non-nil, zero value otherwise.

### GetSupportsListVolumesAttachedNodesOk

`func (o *CSIControllerInfo) GetSupportsListVolumesAttachedNodesOk() (*bool, bool)`

GetSupportsListVolumesAttachedNodesOk returns a tuple with the SupportsListVolumesAttachedNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsListVolumesAttachedNodes

`func (o *CSIControllerInfo) SetSupportsListVolumesAttachedNodes(v bool)`

SetSupportsListVolumesAttachedNodes sets SupportsListVolumesAttachedNodes field to given value.

### HasSupportsListVolumesAttachedNodes

`func (o *CSIControllerInfo) HasSupportsListVolumesAttachedNodes() bool`

HasSupportsListVolumesAttachedNodes returns a boolean if a field has been set.

### GetSupportsReadOnlyAttach

`func (o *CSIControllerInfo) GetSupportsReadOnlyAttach() bool`

GetSupportsReadOnlyAttach returns the SupportsReadOnlyAttach field if non-nil, zero value otherwise.

### GetSupportsReadOnlyAttachOk

`func (o *CSIControllerInfo) GetSupportsReadOnlyAttachOk() (*bool, bool)`

GetSupportsReadOnlyAttachOk returns a tuple with the SupportsReadOnlyAttach field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsReadOnlyAttach

`func (o *CSIControllerInfo) SetSupportsReadOnlyAttach(v bool)`

SetSupportsReadOnlyAttach sets SupportsReadOnlyAttach field to given value.

### HasSupportsReadOnlyAttach

`func (o *CSIControllerInfo) HasSupportsReadOnlyAttach() bool`

HasSupportsReadOnlyAttach returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


