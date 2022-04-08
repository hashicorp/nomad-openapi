# CSISnapshotCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KnownLeader** | Pointer to **bool** |  | [optional] 
**LastContact** | Pointer to **int64** |  | [optional] 
**LastIndex** | Pointer to **int32** |  | [optional] 
**NextToken** | Pointer to **string** |  | [optional] 
**RequestTime** | Pointer to **int64** |  | [optional] 
**Snapshots** | Pointer to [**[]CSISnapshot**](CSISnapshot.md) |  | [optional] 

## Methods

### NewCSISnapshotCreateResponse

`func NewCSISnapshotCreateResponse() *CSISnapshotCreateResponse`

NewCSISnapshotCreateResponse instantiates a new CSISnapshotCreateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSISnapshotCreateResponseWithDefaults

`func NewCSISnapshotCreateResponseWithDefaults() *CSISnapshotCreateResponse`

NewCSISnapshotCreateResponseWithDefaults instantiates a new CSISnapshotCreateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKnownLeader

`func (o *CSISnapshotCreateResponse) GetKnownLeader() bool`

GetKnownLeader returns the KnownLeader field if non-nil, zero value otherwise.

### GetKnownLeaderOk

`func (o *CSISnapshotCreateResponse) GetKnownLeaderOk() (*bool, bool)`

GetKnownLeaderOk returns a tuple with the KnownLeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnownLeader

`func (o *CSISnapshotCreateResponse) SetKnownLeader(v bool)`

SetKnownLeader sets KnownLeader field to given value.

### HasKnownLeader

`func (o *CSISnapshotCreateResponse) HasKnownLeader() bool`

HasKnownLeader returns a boolean if a field has been set.

### GetLastContact

`func (o *CSISnapshotCreateResponse) GetLastContact() int64`

GetLastContact returns the LastContact field if non-nil, zero value otherwise.

### GetLastContactOk

`func (o *CSISnapshotCreateResponse) GetLastContactOk() (*int64, bool)`

GetLastContactOk returns a tuple with the LastContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastContact

`func (o *CSISnapshotCreateResponse) SetLastContact(v int64)`

SetLastContact sets LastContact field to given value.

### HasLastContact

`func (o *CSISnapshotCreateResponse) HasLastContact() bool`

HasLastContact returns a boolean if a field has been set.

### GetLastIndex

`func (o *CSISnapshotCreateResponse) GetLastIndex() int32`

GetLastIndex returns the LastIndex field if non-nil, zero value otherwise.

### GetLastIndexOk

`func (o *CSISnapshotCreateResponse) GetLastIndexOk() (*int32, bool)`

GetLastIndexOk returns a tuple with the LastIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastIndex

`func (o *CSISnapshotCreateResponse) SetLastIndex(v int32)`

SetLastIndex sets LastIndex field to given value.

### HasLastIndex

`func (o *CSISnapshotCreateResponse) HasLastIndex() bool`

HasLastIndex returns a boolean if a field has been set.

### GetNextToken

`func (o *CSISnapshotCreateResponse) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *CSISnapshotCreateResponse) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *CSISnapshotCreateResponse) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *CSISnapshotCreateResponse) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.

### GetRequestTime

`func (o *CSISnapshotCreateResponse) GetRequestTime() int64`

GetRequestTime returns the RequestTime field if non-nil, zero value otherwise.

### GetRequestTimeOk

`func (o *CSISnapshotCreateResponse) GetRequestTimeOk() (*int64, bool)`

GetRequestTimeOk returns a tuple with the RequestTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTime

`func (o *CSISnapshotCreateResponse) SetRequestTime(v int64)`

SetRequestTime sets RequestTime field to given value.

### HasRequestTime

`func (o *CSISnapshotCreateResponse) HasRequestTime() bool`

HasRequestTime returns a boolean if a field has been set.

### GetSnapshots

`func (o *CSISnapshotCreateResponse) GetSnapshots() []CSISnapshot`

GetSnapshots returns the Snapshots field if non-nil, zero value otherwise.

### GetSnapshotsOk

`func (o *CSISnapshotCreateResponse) GetSnapshotsOk() (*[]CSISnapshot, bool)`

GetSnapshotsOk returns a tuple with the Snapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshots

`func (o *CSISnapshotCreateResponse) SetSnapshots(v []CSISnapshot)`

SetSnapshots sets Snapshots field to given value.

### HasSnapshots

`func (o *CSISnapshotCreateResponse) HasSnapshots() bool`

HasSnapshots returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


