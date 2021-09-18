# SearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowStale** | Pointer to **bool** |  | [optional] 
**AuthToken** | Pointer to **string** |  | [optional] 
**Context** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**NextToken** | Pointer to **string** |  | [optional] 
**Params** | Pointer to **map[string]string** |  | [optional] 
**PerPage** | Pointer to **int32** |  | [optional] 
**Prefix** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**WaitIndex** | Pointer to **int32** |  | [optional] 
**WaitTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewSearchRequest

`func NewSearchRequest() *SearchRequest`

NewSearchRequest instantiates a new SearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchRequestWithDefaults

`func NewSearchRequestWithDefaults() *SearchRequest`

NewSearchRequestWithDefaults instantiates a new SearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowStale

`func (o *SearchRequest) GetAllowStale() bool`

GetAllowStale returns the AllowStale field if non-nil, zero value otherwise.

### GetAllowStaleOk

`func (o *SearchRequest) GetAllowStaleOk() (*bool, bool)`

GetAllowStaleOk returns a tuple with the AllowStale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowStale

`func (o *SearchRequest) SetAllowStale(v bool)`

SetAllowStale sets AllowStale field to given value.

### HasAllowStale

`func (o *SearchRequest) HasAllowStale() bool`

HasAllowStale returns a boolean if a field has been set.

### GetAuthToken

`func (o *SearchRequest) GetAuthToken() string`

GetAuthToken returns the AuthToken field if non-nil, zero value otherwise.

### GetAuthTokenOk

`func (o *SearchRequest) GetAuthTokenOk() (*string, bool)`

GetAuthTokenOk returns a tuple with the AuthToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthToken

`func (o *SearchRequest) SetAuthToken(v string)`

SetAuthToken sets AuthToken field to given value.

### HasAuthToken

`func (o *SearchRequest) HasAuthToken() bool`

HasAuthToken returns a boolean if a field has been set.

### GetContext

`func (o *SearchRequest) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *SearchRequest) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *SearchRequest) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *SearchRequest) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetNamespace

`func (o *SearchRequest) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *SearchRequest) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *SearchRequest) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *SearchRequest) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetNextToken

`func (o *SearchRequest) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *SearchRequest) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *SearchRequest) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *SearchRequest) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.

### GetParams

`func (o *SearchRequest) GetParams() map[string]string`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *SearchRequest) GetParamsOk() (*map[string]string, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *SearchRequest) SetParams(v map[string]string)`

SetParams sets Params field to given value.

### HasParams

`func (o *SearchRequest) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetPerPage

`func (o *SearchRequest) GetPerPage() int32`

GetPerPage returns the PerPage field if non-nil, zero value otherwise.

### GetPerPageOk

`func (o *SearchRequest) GetPerPageOk() (*int32, bool)`

GetPerPageOk returns a tuple with the PerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPage

`func (o *SearchRequest) SetPerPage(v int32)`

SetPerPage sets PerPage field to given value.

### HasPerPage

`func (o *SearchRequest) HasPerPage() bool`

HasPerPage returns a boolean if a field has been set.

### GetPrefix

`func (o *SearchRequest) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *SearchRequest) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *SearchRequest) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *SearchRequest) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetRegion

`func (o *SearchRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *SearchRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *SearchRequest) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *SearchRequest) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetWaitIndex

`func (o *SearchRequest) GetWaitIndex() int32`

GetWaitIndex returns the WaitIndex field if non-nil, zero value otherwise.

### GetWaitIndexOk

`func (o *SearchRequest) GetWaitIndexOk() (*int32, bool)`

GetWaitIndexOk returns a tuple with the WaitIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitIndex

`func (o *SearchRequest) SetWaitIndex(v int32)`

SetWaitIndex sets WaitIndex field to given value.

### HasWaitIndex

`func (o *SearchRequest) HasWaitIndex() bool`

HasWaitIndex returns a boolean if a field has been set.

### GetWaitTime

`func (o *SearchRequest) GetWaitTime() int64`

GetWaitTime returns the WaitTime field if non-nil, zero value otherwise.

### GetWaitTimeOk

`func (o *SearchRequest) GetWaitTimeOk() (*int64, bool)`

GetWaitTimeOk returns a tuple with the WaitTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitTime

`func (o *SearchRequest) SetWaitTime(v int64)`

SetWaitTime sets WaitTime field to given value.

### HasWaitTime

`func (o *SearchRequest) HasWaitTime() bool`

HasWaitTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


