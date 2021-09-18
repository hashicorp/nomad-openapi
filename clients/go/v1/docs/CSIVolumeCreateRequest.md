# CSIVolumeCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Namespace** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**SecretID** | Pointer to **string** |  | [optional] 
**Volumes** | Pointer to [**[]CSIVolume**](CSIVolume.md) |  | [optional] 

## Methods

### NewCSIVolumeCreateRequest

`func NewCSIVolumeCreateRequest() *CSIVolumeCreateRequest`

NewCSIVolumeCreateRequest instantiates a new CSIVolumeCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSIVolumeCreateRequestWithDefaults

`func NewCSIVolumeCreateRequestWithDefaults() *CSIVolumeCreateRequest`

NewCSIVolumeCreateRequestWithDefaults instantiates a new CSIVolumeCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamespace

`func (o *CSIVolumeCreateRequest) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *CSIVolumeCreateRequest) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *CSIVolumeCreateRequest) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *CSIVolumeCreateRequest) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetRegion

`func (o *CSIVolumeCreateRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CSIVolumeCreateRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CSIVolumeCreateRequest) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CSIVolumeCreateRequest) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSecretID

`func (o *CSIVolumeCreateRequest) GetSecretID() string`

GetSecretID returns the SecretID field if non-nil, zero value otherwise.

### GetSecretIDOk

`func (o *CSIVolumeCreateRequest) GetSecretIDOk() (*string, bool)`

GetSecretIDOk returns a tuple with the SecretID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretID

`func (o *CSIVolumeCreateRequest) SetSecretID(v string)`

SetSecretID sets SecretID field to given value.

### HasSecretID

`func (o *CSIVolumeCreateRequest) HasSecretID() bool`

HasSecretID returns a boolean if a field has been set.

### GetVolumes

`func (o *CSIVolumeCreateRequest) GetVolumes() []CSIVolume`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *CSIVolumeCreateRequest) GetVolumesOk() (*[]CSIVolume, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *CSIVolumeCreateRequest) SetVolumes(v []CSIVolume)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *CSIVolumeCreateRequest) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


