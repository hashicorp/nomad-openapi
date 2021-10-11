# AllocStopResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EvalID** | Pointer to **string** |  | [optional] 
**LastIndex** | Pointer to **int32** |  | [optional] 
**RequestTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewAllocStopResponse

`func NewAllocStopResponse() *AllocStopResponse`

NewAllocStopResponse instantiates a new AllocStopResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllocStopResponseWithDefaults

`func NewAllocStopResponseWithDefaults() *AllocStopResponse`

NewAllocStopResponseWithDefaults instantiates a new AllocStopResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvalID

`func (o *AllocStopResponse) GetEvalID() string`

GetEvalID returns the EvalID field if non-nil, zero value otherwise.

### GetEvalIDOk

`func (o *AllocStopResponse) GetEvalIDOk() (*string, bool)`

GetEvalIDOk returns a tuple with the EvalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvalID

`func (o *AllocStopResponse) SetEvalID(v string)`

SetEvalID sets EvalID field to given value.

### HasEvalID

`func (o *AllocStopResponse) HasEvalID() bool`

HasEvalID returns a boolean if a field has been set.

### GetLastIndex

`func (o *AllocStopResponse) GetLastIndex() int32`

GetLastIndex returns the LastIndex field if non-nil, zero value otherwise.

### GetLastIndexOk

`func (o *AllocStopResponse) GetLastIndexOk() (*int32, bool)`

GetLastIndexOk returns a tuple with the LastIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastIndex

`func (o *AllocStopResponse) SetLastIndex(v int32)`

SetLastIndex sets LastIndex field to given value.

### HasLastIndex

`func (o *AllocStopResponse) HasLastIndex() bool`

HasLastIndex returns a boolean if a field has been set.

### GetRequestTime

`func (o *AllocStopResponse) GetRequestTime() int64`

GetRequestTime returns the RequestTime field if non-nil, zero value otherwise.

### GetRequestTimeOk

`func (o *AllocStopResponse) GetRequestTimeOk() (*int64, bool)`

GetRequestTimeOk returns a tuple with the RequestTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTime

`func (o *AllocStopResponse) SetRequestTime(v int64)`

SetRequestTime sets RequestTime field to given value.

### HasRequestTime

`func (o *AllocStopResponse) HasRequestTime() bool`

HasRequestTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


