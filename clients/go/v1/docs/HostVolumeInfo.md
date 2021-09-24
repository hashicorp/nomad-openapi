# HostVolumeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** |  | [optional] 
**ReadOnly** | Pointer to **bool** |  | [optional] 

## Methods

### NewHostVolumeInfo

`func NewHostVolumeInfo() *HostVolumeInfo`

NewHostVolumeInfo instantiates a new HostVolumeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostVolumeInfoWithDefaults

`func NewHostVolumeInfoWithDefaults() *HostVolumeInfo`

NewHostVolumeInfoWithDefaults instantiates a new HostVolumeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *HostVolumeInfo) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *HostVolumeInfo) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *HostVolumeInfo) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *HostVolumeInfo) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetReadOnly

`func (o *HostVolumeInfo) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *HostVolumeInfo) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *HostVolumeInfo) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *HostVolumeInfo) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


