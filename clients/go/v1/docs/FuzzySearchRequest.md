# FuzzySearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowStale** | Pointer to **bool** |  | [optional] 
**AuthToken** | Pointer to **string** |  | [optional] 
**Context** | Pointer to **string** |  | [optional] 
**Filter** | Pointer to **string** |  | [optional] 
**Headers** | Pointer to **map[string]string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**NextToken** | Pointer to **string** |  | [optional] 
**Params** | Pointer to **map[string]string** |  | [optional] 
**PerPage** | Pointer to **int32** |  | [optional] 
**Prefix** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Reverse** | Pointer to **bool** |  | [optional] 
**Text** | Pointer to **string** |  | [optional] 
**WaitIndex** | Pointer to **int32** |  | [optional] 
**WaitTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewFuzzySearchRequest

`func NewFuzzySearchRequest() *FuzzySearchRequest`

NewFuzzySearchRequest instantiates a new FuzzySearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFuzzySearchRequestWithDefaults

`func NewFuzzySearchRequestWithDefaults() *FuzzySearchRequest`

NewFuzzySearchRequestWithDefaults instantiates a new FuzzySearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowStale

`func (o *FuzzySearchRequest) GetAllowStale() bool`

GetAllowStale returns the AllowStale field if non-nil, zero value otherwise.

### GetAllowStaleOk

`func (o *FuzzySearchRequest) GetAllowStaleOk() (*bool, bool)`

GetAllowStaleOk returns a tuple with the AllowStale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowStale

`func (o *FuzzySearchRequest) SetAllowStale(v bool)`

SetAllowStale sets AllowStale field to given value.

### HasAllowStale

`func (o *FuzzySearchRequest) HasAllowStale() bool`

HasAllowStale returns a boolean if a field has been set.

### GetAuthToken

`func (o *FuzzySearchRequest) GetAuthToken() string`

GetAuthToken returns the AuthToken field if non-nil, zero value otherwise.

### GetAuthTokenOk

`func (o *FuzzySearchRequest) GetAuthTokenOk() (*string, bool)`

GetAuthTokenOk returns a tuple with the AuthToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthToken

`func (o *FuzzySearchRequest) SetAuthToken(v string)`

SetAuthToken sets AuthToken field to given value.

### HasAuthToken

`func (o *FuzzySearchRequest) HasAuthToken() bool`

HasAuthToken returns a boolean if a field has been set.

### GetContext

`func (o *FuzzySearchRequest) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *FuzzySearchRequest) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *FuzzySearchRequest) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *FuzzySearchRequest) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetFilter

`func (o *FuzzySearchRequest) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *FuzzySearchRequest) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *FuzzySearchRequest) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *FuzzySearchRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetHeaders

`func (o *FuzzySearchRequest) GetHeaders() map[string]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *FuzzySearchRequest) GetHeadersOk() (*map[string]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *FuzzySearchRequest) SetHeaders(v map[string]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *FuzzySearchRequest) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetNamespace

`func (o *FuzzySearchRequest) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *FuzzySearchRequest) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *FuzzySearchRequest) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *FuzzySearchRequest) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetNextToken

`func (o *FuzzySearchRequest) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *FuzzySearchRequest) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *FuzzySearchRequest) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *FuzzySearchRequest) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.

### GetParams

`func (o *FuzzySearchRequest) GetParams() map[string]string`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *FuzzySearchRequest) GetParamsOk() (*map[string]string, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *FuzzySearchRequest) SetParams(v map[string]string)`

SetParams sets Params field to given value.

### HasParams

`func (o *FuzzySearchRequest) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetPerPage

`func (o *FuzzySearchRequest) GetPerPage() int32`

GetPerPage returns the PerPage field if non-nil, zero value otherwise.

### GetPerPageOk

`func (o *FuzzySearchRequest) GetPerPageOk() (*int32, bool)`

GetPerPageOk returns a tuple with the PerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPage

`func (o *FuzzySearchRequest) SetPerPage(v int32)`

SetPerPage sets PerPage field to given value.

### HasPerPage

`func (o *FuzzySearchRequest) HasPerPage() bool`

HasPerPage returns a boolean if a field has been set.

### GetPrefix

`func (o *FuzzySearchRequest) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *FuzzySearchRequest) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *FuzzySearchRequest) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *FuzzySearchRequest) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetRegion

`func (o *FuzzySearchRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *FuzzySearchRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *FuzzySearchRequest) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *FuzzySearchRequest) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetReverse

`func (o *FuzzySearchRequest) GetReverse() bool`

GetReverse returns the Reverse field if non-nil, zero value otherwise.

### GetReverseOk

`func (o *FuzzySearchRequest) GetReverseOk() (*bool, bool)`

GetReverseOk returns a tuple with the Reverse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReverse

`func (o *FuzzySearchRequest) SetReverse(v bool)`

SetReverse sets Reverse field to given value.

### HasReverse

`func (o *FuzzySearchRequest) HasReverse() bool`

HasReverse returns a boolean if a field has been set.

### GetText

`func (o *FuzzySearchRequest) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *FuzzySearchRequest) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *FuzzySearchRequest) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *FuzzySearchRequest) HasText() bool`

HasText returns a boolean if a field has been set.

### GetWaitIndex

`func (o *FuzzySearchRequest) GetWaitIndex() int32`

GetWaitIndex returns the WaitIndex field if non-nil, zero value otherwise.

### GetWaitIndexOk

`func (o *FuzzySearchRequest) GetWaitIndexOk() (*int32, bool)`

GetWaitIndexOk returns a tuple with the WaitIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitIndex

`func (o *FuzzySearchRequest) SetWaitIndex(v int32)`

SetWaitIndex sets WaitIndex field to given value.

### HasWaitIndex

`func (o *FuzzySearchRequest) HasWaitIndex() bool`

HasWaitIndex returns a boolean if a field has been set.

### GetWaitTime

`func (o *FuzzySearchRequest) GetWaitTime() int64`

GetWaitTime returns the WaitTime field if non-nil, zero value otherwise.

### GetWaitTimeOk

`func (o *FuzzySearchRequest) GetWaitTimeOk() (*int64, bool)`

GetWaitTimeOk returns a tuple with the WaitTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitTime

`func (o *FuzzySearchRequest) SetWaitTime(v int64)`

SetWaitTime sets WaitTime field to given value.

### HasWaitTime

`func (o *FuzzySearchRequest) HasWaitTime() bool`

HasWaitTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


