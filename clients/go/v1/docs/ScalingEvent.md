# ScalingEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int64** |  | [optional] 
**CreateIndex** | Pointer to **int32** |  | [optional] 
**Error** | Pointer to **bool** |  | [optional] 
**EvalID** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to **map[string]interface{}** |  | [optional] 
**PreviousCount** | Pointer to **int64** |  | [optional] 
**Time** | Pointer to **int32** |  | [optional] 

## Methods

### NewScalingEvent

`func NewScalingEvent() *ScalingEvent`

NewScalingEvent instantiates a new ScalingEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScalingEventWithDefaults

`func NewScalingEventWithDefaults() *ScalingEvent`

NewScalingEventWithDefaults instantiates a new ScalingEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *ScalingEvent) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ScalingEvent) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ScalingEvent) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *ScalingEvent) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetCreateIndex

`func (o *ScalingEvent) GetCreateIndex() int32`

GetCreateIndex returns the CreateIndex field if non-nil, zero value otherwise.

### GetCreateIndexOk

`func (o *ScalingEvent) GetCreateIndexOk() (*int32, bool)`

GetCreateIndexOk returns a tuple with the CreateIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateIndex

`func (o *ScalingEvent) SetCreateIndex(v int32)`

SetCreateIndex sets CreateIndex field to given value.

### HasCreateIndex

`func (o *ScalingEvent) HasCreateIndex() bool`

HasCreateIndex returns a boolean if a field has been set.

### GetError

`func (o *ScalingEvent) GetError() bool`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ScalingEvent) GetErrorOk() (*bool, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ScalingEvent) SetError(v bool)`

SetError sets Error field to given value.

### HasError

`func (o *ScalingEvent) HasError() bool`

HasError returns a boolean if a field has been set.

### GetEvalID

`func (o *ScalingEvent) GetEvalID() string`

GetEvalID returns the EvalID field if non-nil, zero value otherwise.

### GetEvalIDOk

`func (o *ScalingEvent) GetEvalIDOk() (*string, bool)`

GetEvalIDOk returns a tuple with the EvalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvalID

`func (o *ScalingEvent) SetEvalID(v string)`

SetEvalID sets EvalID field to given value.

### HasEvalID

`func (o *ScalingEvent) HasEvalID() bool`

HasEvalID returns a boolean if a field has been set.

### GetMessage

`func (o *ScalingEvent) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ScalingEvent) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ScalingEvent) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ScalingEvent) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetMeta

`func (o *ScalingEvent) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ScalingEvent) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ScalingEvent) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ScalingEvent) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetPreviousCount

`func (o *ScalingEvent) GetPreviousCount() int64`

GetPreviousCount returns the PreviousCount field if non-nil, zero value otherwise.

### GetPreviousCountOk

`func (o *ScalingEvent) GetPreviousCountOk() (*int64, bool)`

GetPreviousCountOk returns a tuple with the PreviousCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousCount

`func (o *ScalingEvent) SetPreviousCount(v int64)`

SetPreviousCount sets PreviousCount field to given value.

### HasPreviousCount

`func (o *ScalingEvent) HasPreviousCount() bool`

HasPreviousCount returns a boolean if a field has been set.

### GetTime

`func (o *ScalingEvent) GetTime() int32`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ScalingEvent) GetTimeOk() (*int32, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ScalingEvent) SetTime(v int32)`

SetTime sets Time field to given value.

### HasTime

`func (o *ScalingEvent) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


