# CSISnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateTime** | Pointer to **int64** |  | [optional] 
**ExternalSourceVolumeID** | Pointer to **string** |  | [optional] 
**ID** | Pointer to **string** |  | [optional] 
**IsReady** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Parameters** | Pointer to **map[string]string** |  | [optional] 
**PluginID** | Pointer to **string** |  | [optional] 
**Secrets** | Pointer to **map[string]string** |  | [optional] 
**SizeBytes** | Pointer to **int64** |  | [optional] 
**SourceVolumeID** | Pointer to **string** |  | [optional] 

## Methods

### NewCSISnapshot

`func NewCSISnapshot() *CSISnapshot`

NewCSISnapshot instantiates a new CSISnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSISnapshotWithDefaults

`func NewCSISnapshotWithDefaults() *CSISnapshot`

NewCSISnapshotWithDefaults instantiates a new CSISnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateTime

`func (o *CSISnapshot) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *CSISnapshot) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *CSISnapshot) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *CSISnapshot) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetExternalSourceVolumeID

`func (o *CSISnapshot) GetExternalSourceVolumeID() string`

GetExternalSourceVolumeID returns the ExternalSourceVolumeID field if non-nil, zero value otherwise.

### GetExternalSourceVolumeIDOk

`func (o *CSISnapshot) GetExternalSourceVolumeIDOk() (*string, bool)`

GetExternalSourceVolumeIDOk returns a tuple with the ExternalSourceVolumeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSourceVolumeID

`func (o *CSISnapshot) SetExternalSourceVolumeID(v string)`

SetExternalSourceVolumeID sets ExternalSourceVolumeID field to given value.

### HasExternalSourceVolumeID

`func (o *CSISnapshot) HasExternalSourceVolumeID() bool`

HasExternalSourceVolumeID returns a boolean if a field has been set.

### GetID

`func (o *CSISnapshot) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *CSISnapshot) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *CSISnapshot) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *CSISnapshot) HasID() bool`

HasID returns a boolean if a field has been set.

### GetIsReady

`func (o *CSISnapshot) GetIsReady() bool`

GetIsReady returns the IsReady field if non-nil, zero value otherwise.

### GetIsReadyOk

`func (o *CSISnapshot) GetIsReadyOk() (*bool, bool)`

GetIsReadyOk returns a tuple with the IsReady field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReady

`func (o *CSISnapshot) SetIsReady(v bool)`

SetIsReady sets IsReady field to given value.

### HasIsReady

`func (o *CSISnapshot) HasIsReady() bool`

HasIsReady returns a boolean if a field has been set.

### GetName

`func (o *CSISnapshot) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CSISnapshot) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CSISnapshot) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CSISnapshot) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParameters

`func (o *CSISnapshot) GetParameters() map[string]string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *CSISnapshot) GetParametersOk() (*map[string]string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *CSISnapshot) SetParameters(v map[string]string)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *CSISnapshot) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetPluginID

`func (o *CSISnapshot) GetPluginID() string`

GetPluginID returns the PluginID field if non-nil, zero value otherwise.

### GetPluginIDOk

`func (o *CSISnapshot) GetPluginIDOk() (*string, bool)`

GetPluginIDOk returns a tuple with the PluginID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginID

`func (o *CSISnapshot) SetPluginID(v string)`

SetPluginID sets PluginID field to given value.

### HasPluginID

`func (o *CSISnapshot) HasPluginID() bool`

HasPluginID returns a boolean if a field has been set.

### GetSecrets

`func (o *CSISnapshot) GetSecrets() map[string]string`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *CSISnapshot) GetSecretsOk() (*map[string]string, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *CSISnapshot) SetSecrets(v map[string]string)`

SetSecrets sets Secrets field to given value.

### HasSecrets

`func (o *CSISnapshot) HasSecrets() bool`

HasSecrets returns a boolean if a field has been set.

### GetSizeBytes

`func (o *CSISnapshot) GetSizeBytes() int64`

GetSizeBytes returns the SizeBytes field if non-nil, zero value otherwise.

### GetSizeBytesOk

`func (o *CSISnapshot) GetSizeBytesOk() (*int64, bool)`

GetSizeBytesOk returns a tuple with the SizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBytes

`func (o *CSISnapshot) SetSizeBytes(v int64)`

SetSizeBytes sets SizeBytes field to given value.

### HasSizeBytes

`func (o *CSISnapshot) HasSizeBytes() bool`

HasSizeBytes returns a boolean if a field has been set.

### GetSourceVolumeID

`func (o *CSISnapshot) GetSourceVolumeID() string`

GetSourceVolumeID returns the SourceVolumeID field if non-nil, zero value otherwise.

### GetSourceVolumeIDOk

`func (o *CSISnapshot) GetSourceVolumeIDOk() (*string, bool)`

GetSourceVolumeIDOk returns a tuple with the SourceVolumeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceVolumeID

`func (o *CSISnapshot) SetSourceVolumeID(v string)`

SetSourceVolumeID sets SourceVolumeID field to given value.

### HasSourceVolumeID

`func (o *CSISnapshot) HasSourceVolumeID() bool`

HasSourceVolumeID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


