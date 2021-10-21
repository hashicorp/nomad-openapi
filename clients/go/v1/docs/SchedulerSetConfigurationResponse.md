# SchedulerSetConfigurationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastIndex** | Pointer to **int32** |  | [optional] 
**RequestTime** | Pointer to **int64** |  | [optional] 
**Updated** | Pointer to **bool** |  | [optional] 

## Methods

### NewSchedulerSetConfigurationResponse

`func NewSchedulerSetConfigurationResponse() *SchedulerSetConfigurationResponse`

NewSchedulerSetConfigurationResponse instantiates a new SchedulerSetConfigurationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulerSetConfigurationResponseWithDefaults

`func NewSchedulerSetConfigurationResponseWithDefaults() *SchedulerSetConfigurationResponse`

NewSchedulerSetConfigurationResponseWithDefaults instantiates a new SchedulerSetConfigurationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastIndex

`func (o *SchedulerSetConfigurationResponse) GetLastIndex() int32`

GetLastIndex returns the LastIndex field if non-nil, zero value otherwise.

### GetLastIndexOk

`func (o *SchedulerSetConfigurationResponse) GetLastIndexOk() (*int32, bool)`

GetLastIndexOk returns a tuple with the LastIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastIndex

`func (o *SchedulerSetConfigurationResponse) SetLastIndex(v int32)`

SetLastIndex sets LastIndex field to given value.

### HasLastIndex

`func (o *SchedulerSetConfigurationResponse) HasLastIndex() bool`

HasLastIndex returns a boolean if a field has been set.

### GetRequestTime

`func (o *SchedulerSetConfigurationResponse) GetRequestTime() int64`

GetRequestTime returns the RequestTime field if non-nil, zero value otherwise.

### GetRequestTimeOk

`func (o *SchedulerSetConfigurationResponse) GetRequestTimeOk() (*int64, bool)`

GetRequestTimeOk returns a tuple with the RequestTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTime

`func (o *SchedulerSetConfigurationResponse) SetRequestTime(v int64)`

SetRequestTime sets RequestTime field to given value.

### HasRequestTime

`func (o *SchedulerSetConfigurationResponse) HasRequestTime() bool`

HasRequestTime returns a boolean if a field has been set.

### GetUpdated

`func (o *SchedulerSetConfigurationResponse) GetUpdated() bool`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *SchedulerSetConfigurationResponse) GetUpdatedOk() (*bool, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *SchedulerSetConfigurationResponse) SetUpdated(v bool)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *SchedulerSetConfigurationResponse) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


