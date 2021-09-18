# CSISnapshotCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Namespace** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**SecretID** | Pointer to **string** |  | [optional] 
**Snapshots** | Pointer to [**[]CSISnapshot**](CSISnapshot.md) |  | [optional] 

## Methods

### NewCSISnapshotCreateRequest

`func NewCSISnapshotCreateRequest() *CSISnapshotCreateRequest`

NewCSISnapshotCreateRequest instantiates a new CSISnapshotCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSISnapshotCreateRequestWithDefaults

`func NewCSISnapshotCreateRequestWithDefaults() *CSISnapshotCreateRequest`

NewCSISnapshotCreateRequestWithDefaults instantiates a new CSISnapshotCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamespace

`func (o *CSISnapshotCreateRequest) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *CSISnapshotCreateRequest) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *CSISnapshotCreateRequest) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *CSISnapshotCreateRequest) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetRegion

`func (o *CSISnapshotCreateRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CSISnapshotCreateRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CSISnapshotCreateRequest) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CSISnapshotCreateRequest) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSecretID

`func (o *CSISnapshotCreateRequest) GetSecretID() string`

GetSecretID returns the SecretID field if non-nil, zero value otherwise.

### GetSecretIDOk

`func (o *CSISnapshotCreateRequest) GetSecretIDOk() (*string, bool)`

GetSecretIDOk returns a tuple with the SecretID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretID

`func (o *CSISnapshotCreateRequest) SetSecretID(v string)`

SetSecretID sets SecretID field to given value.

### HasSecretID

`func (o *CSISnapshotCreateRequest) HasSecretID() bool`

HasSecretID returns a boolean if a field has been set.

### GetSnapshots

`func (o *CSISnapshotCreateRequest) GetSnapshots() []CSISnapshot`

GetSnapshots returns the Snapshots field if non-nil, zero value otherwise.

### GetSnapshotsOk

`func (o *CSISnapshotCreateRequest) GetSnapshotsOk() (*[]CSISnapshot, bool)`

GetSnapshotsOk returns a tuple with the Snapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshots

`func (o *CSISnapshotCreateRequest) SetSnapshots(v []CSISnapshot)`

SetSnapshots sets Snapshots field to given value.

### HasSnapshots

`func (o *CSISnapshotCreateRequest) HasSnapshots() bool`

HasSnapshots returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


