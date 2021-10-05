# CSIControllerInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SupportsAttachDetach** | Pointer to **bool** |  | [optional] 
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


