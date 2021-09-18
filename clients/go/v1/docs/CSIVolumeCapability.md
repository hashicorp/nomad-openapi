# CSIVolumeCapability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessMode** | Pointer to **string** |  | [optional] 
**AttachmentMode** | Pointer to **string** |  | [optional] 

## Methods

### NewCSIVolumeCapability

`func NewCSIVolumeCapability() *CSIVolumeCapability`

NewCSIVolumeCapability instantiates a new CSIVolumeCapability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSIVolumeCapabilityWithDefaults

`func NewCSIVolumeCapabilityWithDefaults() *CSIVolumeCapability`

NewCSIVolumeCapabilityWithDefaults instantiates a new CSIVolumeCapability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessMode

`func (o *CSIVolumeCapability) GetAccessMode() string`

GetAccessMode returns the AccessMode field if non-nil, zero value otherwise.

### GetAccessModeOk

`func (o *CSIVolumeCapability) GetAccessModeOk() (*string, bool)`

GetAccessModeOk returns a tuple with the AccessMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessMode

`func (o *CSIVolumeCapability) SetAccessMode(v string)`

SetAccessMode sets AccessMode field to given value.

### HasAccessMode

`func (o *CSIVolumeCapability) HasAccessMode() bool`

HasAccessMode returns a boolean if a field has been set.

### GetAttachmentMode

`func (o *CSIVolumeCapability) GetAttachmentMode() string`

GetAttachmentMode returns the AttachmentMode field if non-nil, zero value otherwise.

### GetAttachmentModeOk

`func (o *CSIVolumeCapability) GetAttachmentModeOk() (*string, bool)`

GetAttachmentModeOk returns a tuple with the AttachmentMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentMode

`func (o *CSIVolumeCapability) SetAttachmentMode(v string)`

SetAttachmentMode sets AttachmentMode field to given value.

### HasAttachmentMode

`func (o *CSIVolumeCapability) HasAttachmentMode() bool`

HasAttachmentMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


