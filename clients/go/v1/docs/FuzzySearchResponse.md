# FuzzySearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KnownLeader** | Pointer to **bool** |  | [optional] 
**LastContact** | Pointer to **int64** |  | [optional] 
**LastIndex** | Pointer to **int32** |  | [optional] 
**Matches** | Pointer to [**map[string][]FuzzyMatch**](array.md) |  | [optional] 
**RequestTime** | Pointer to **int64** |  | [optional] 
**Truncations** | Pointer to **map[string]bool** |  | [optional] 

## Methods

### NewFuzzySearchResponse

`func NewFuzzySearchResponse() *FuzzySearchResponse`

NewFuzzySearchResponse instantiates a new FuzzySearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFuzzySearchResponseWithDefaults

`func NewFuzzySearchResponseWithDefaults() *FuzzySearchResponse`

NewFuzzySearchResponseWithDefaults instantiates a new FuzzySearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKnownLeader

`func (o *FuzzySearchResponse) GetKnownLeader() bool`

GetKnownLeader returns the KnownLeader field if non-nil, zero value otherwise.

### GetKnownLeaderOk

`func (o *FuzzySearchResponse) GetKnownLeaderOk() (*bool, bool)`

GetKnownLeaderOk returns a tuple with the KnownLeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnownLeader

`func (o *FuzzySearchResponse) SetKnownLeader(v bool)`

SetKnownLeader sets KnownLeader field to given value.

### HasKnownLeader

`func (o *FuzzySearchResponse) HasKnownLeader() bool`

HasKnownLeader returns a boolean if a field has been set.

### GetLastContact

`func (o *FuzzySearchResponse) GetLastContact() int64`

GetLastContact returns the LastContact field if non-nil, zero value otherwise.

### GetLastContactOk

`func (o *FuzzySearchResponse) GetLastContactOk() (*int64, bool)`

GetLastContactOk returns a tuple with the LastContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastContact

`func (o *FuzzySearchResponse) SetLastContact(v int64)`

SetLastContact sets LastContact field to given value.

### HasLastContact

`func (o *FuzzySearchResponse) HasLastContact() bool`

HasLastContact returns a boolean if a field has been set.

### GetLastIndex

`func (o *FuzzySearchResponse) GetLastIndex() int32`

GetLastIndex returns the LastIndex field if non-nil, zero value otherwise.

### GetLastIndexOk

`func (o *FuzzySearchResponse) GetLastIndexOk() (*int32, bool)`

GetLastIndexOk returns a tuple with the LastIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastIndex

`func (o *FuzzySearchResponse) SetLastIndex(v int32)`

SetLastIndex sets LastIndex field to given value.

### HasLastIndex

`func (o *FuzzySearchResponse) HasLastIndex() bool`

HasLastIndex returns a boolean if a field has been set.

### GetMatches

`func (o *FuzzySearchResponse) GetMatches() map[string][]FuzzyMatch`

GetMatches returns the Matches field if non-nil, zero value otherwise.

### GetMatchesOk

`func (o *FuzzySearchResponse) GetMatchesOk() (*map[string][]FuzzyMatch, bool)`

GetMatchesOk returns a tuple with the Matches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatches

`func (o *FuzzySearchResponse) SetMatches(v map[string][]FuzzyMatch)`

SetMatches sets Matches field to given value.

### HasMatches

`func (o *FuzzySearchResponse) HasMatches() bool`

HasMatches returns a boolean if a field has been set.

### GetRequestTime

`func (o *FuzzySearchResponse) GetRequestTime() int64`

GetRequestTime returns the RequestTime field if non-nil, zero value otherwise.

### GetRequestTimeOk

`func (o *FuzzySearchResponse) GetRequestTimeOk() (*int64, bool)`

GetRequestTimeOk returns a tuple with the RequestTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTime

`func (o *FuzzySearchResponse) SetRequestTime(v int64)`

SetRequestTime sets RequestTime field to given value.

### HasRequestTime

`func (o *FuzzySearchResponse) HasRequestTime() bool`

HasRequestTime returns a boolean if a field has been set.

### GetTruncations

`func (o *FuzzySearchResponse) GetTruncations() map[string]bool`

GetTruncations returns the Truncations field if non-nil, zero value otherwise.

### GetTruncationsOk

`func (o *FuzzySearchResponse) GetTruncationsOk() (*map[string]bool, bool)`

GetTruncationsOk returns a tuple with the Truncations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncations

`func (o *FuzzySearchResponse) SetTruncations(v map[string]bool)`

SetTruncations sets Truncations field to given value.

### HasTruncations

`func (o *FuzzySearchResponse) HasTruncations() bool`

HasTruncations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


