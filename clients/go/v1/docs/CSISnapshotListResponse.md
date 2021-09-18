# CSISnapshotListResponse

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

### NewCSISnapshotListResponse

`func NewCSISnapshotListResponse() *CSISnapshotListResponse`

NewCSISnapshotListResponse instantiates a new CSISnapshotListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSISnapshotListResponseWithDefaults

`func NewCSISnapshotListResponseWithDefaults() *CSISnapshotListResponse`

NewCSISnapshotListResponseWithDefaults instantiates a new CSISnapshotListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKnownLeader

`func (o *CSISnapshotListResponse) GetKnownLeader() bool`

GetKnownLeader returns the KnownLeader field if non-nil, zero value otherwise.

### GetKnownLeaderOk

`func (o *CSISnapshotListResponse) GetKnownLeaderOk() (*bool, bool)`

GetKnownLeaderOk returns a tuple with the KnownLeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnownLeader

`func (o *CSISnapshotListResponse) SetKnownLeader(v bool)`

SetKnownLeader sets KnownLeader field to given value.

### HasKnownLeader

`func (o *CSISnapshotListResponse) HasKnownLeader() bool`

HasKnownLeader returns a boolean if a field has been set.

### GetLastContact

`func (o *CSISnapshotListResponse) GetLastContact() int64`

GetLastContact returns the LastContact field if non-nil, zero value otherwise.

### GetLastContactOk

`func (o *CSISnapshotListResponse) GetLastContactOk() (*int64, bool)`

GetLastContactOk returns a tuple with the LastContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastContact

`func (o *CSISnapshotListResponse) SetLastContact(v int64)`

SetLastContact sets LastContact field to given value.

### HasLastContact

`func (o *CSISnapshotListResponse) HasLastContact() bool`

HasLastContact returns a boolean if a field has been set.

### GetLastIndex

`func (o *CSISnapshotListResponse) GetLastIndex() int32`

GetLastIndex returns the LastIndex field if non-nil, zero value otherwise.

### GetLastIndexOk

`func (o *CSISnapshotListResponse) GetLastIndexOk() (*int32, bool)`

GetLastIndexOk returns a tuple with the LastIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastIndex

`func (o *CSISnapshotListResponse) SetLastIndex(v int32)`

SetLastIndex sets LastIndex field to given value.

### HasLastIndex

`func (o *CSISnapshotListResponse) HasLastIndex() bool`

HasLastIndex returns a boolean if a field has been set.

### GetNextToken

`func (o *CSISnapshotListResponse) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *CSISnapshotListResponse) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *CSISnapshotListResponse) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *CSISnapshotListResponse) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.

### GetRequestTime

`func (o *CSISnapshotListResponse) GetRequestTime() int64`

GetRequestTime returns the RequestTime field if non-nil, zero value otherwise.

### GetRequestTimeOk

`func (o *CSISnapshotListResponse) GetRequestTimeOk() (*int64, bool)`

GetRequestTimeOk returns a tuple with the RequestTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTime

`func (o *CSISnapshotListResponse) SetRequestTime(v int64)`

SetRequestTime sets RequestTime field to given value.

### HasRequestTime

`func (o *CSISnapshotListResponse) HasRequestTime() bool`

HasRequestTime returns a boolean if a field has been set.

### GetSnapshots

`func (o *CSISnapshotListResponse) GetSnapshots() []CSISnapshot`

GetSnapshots returns the Snapshots field if non-nil, zero value otherwise.

### GetSnapshotsOk

`func (o *CSISnapshotListResponse) GetSnapshotsOk() (*[]CSISnapshot, bool)`

GetSnapshotsOk returns a tuple with the Snapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshots

`func (o *CSISnapshotListResponse) SetSnapshots(v []CSISnapshot)`

SetSnapshots sets Snapshots field to given value.

### HasSnapshots

`func (o *CSISnapshotListResponse) HasSnapshots() bool`

HasSnapshots returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


