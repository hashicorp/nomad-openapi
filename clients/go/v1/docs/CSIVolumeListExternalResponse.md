# CSIVolumeListExternalResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextToken** | Pointer to **string** |  | [optional] 
**Volumes** | Pointer to [**[]CSIVolumeExternalStub**](CSIVolumeExternalStub.md) |  | [optional] 

## Methods

### NewCSIVolumeListExternalResponse

`func NewCSIVolumeListExternalResponse() *CSIVolumeListExternalResponse`

NewCSIVolumeListExternalResponse instantiates a new CSIVolumeListExternalResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSIVolumeListExternalResponseWithDefaults

`func NewCSIVolumeListExternalResponseWithDefaults() *CSIVolumeListExternalResponse`

NewCSIVolumeListExternalResponseWithDefaults instantiates a new CSIVolumeListExternalResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextToken

`func (o *CSIVolumeListExternalResponse) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *CSIVolumeListExternalResponse) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *CSIVolumeListExternalResponse) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *CSIVolumeListExternalResponse) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.

### GetVolumes

`func (o *CSIVolumeListExternalResponse) GetVolumes() []CSIVolumeExternalStub`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *CSIVolumeListExternalResponse) GetVolumesOk() (*[]CSIVolumeExternalStub, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *CSIVolumeListExternalResponse) SetVolumes(v []CSIVolumeExternalStub)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *CSIVolumeListExternalResponse) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


