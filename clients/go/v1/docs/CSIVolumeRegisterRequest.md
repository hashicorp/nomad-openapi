# CSIVolumeRegisterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Namespace** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**SecretID** | Pointer to **string** |  | [optional] 
**Volumes** | Pointer to [**[]CSIVolume**](CSIVolume.md) |  | [optional] 

## Methods

### NewCSIVolumeRegisterRequest

`func NewCSIVolumeRegisterRequest() *CSIVolumeRegisterRequest`

NewCSIVolumeRegisterRequest instantiates a new CSIVolumeRegisterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSIVolumeRegisterRequestWithDefaults

`func NewCSIVolumeRegisterRequestWithDefaults() *CSIVolumeRegisterRequest`

NewCSIVolumeRegisterRequestWithDefaults instantiates a new CSIVolumeRegisterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamespace

`func (o *CSIVolumeRegisterRequest) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *CSIVolumeRegisterRequest) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *CSIVolumeRegisterRequest) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *CSIVolumeRegisterRequest) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetRegion

`func (o *CSIVolumeRegisterRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CSIVolumeRegisterRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CSIVolumeRegisterRequest) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CSIVolumeRegisterRequest) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSecretID

`func (o *CSIVolumeRegisterRequest) GetSecretID() string`

GetSecretID returns the SecretID field if non-nil, zero value otherwise.

### GetSecretIDOk

`func (o *CSIVolumeRegisterRequest) GetSecretIDOk() (*string, bool)`

GetSecretIDOk returns a tuple with the SecretID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretID

`func (o *CSIVolumeRegisterRequest) SetSecretID(v string)`

SetSecretID sets SecretID field to given value.

### HasSecretID

`func (o *CSIVolumeRegisterRequest) HasSecretID() bool`

HasSecretID returns a boolean if a field has been set.

### GetVolumes

`func (o *CSIVolumeRegisterRequest) GetVolumes() []CSIVolume`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *CSIVolumeRegisterRequest) GetVolumesOk() (*[]CSIVolume, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *CSIVolumeRegisterRequest) SetVolumes(v []CSIVolume)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *CSIVolumeRegisterRequest) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


