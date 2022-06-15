# SchedulerConfigurationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KnownLeader** | Pointer to **bool** |  | [optional] 
**LastContact** | Pointer to **int64** |  | [optional] 
**LastIndex** | Pointer to **int32** |  | [optional] 
**NextToken** | Pointer to **string** |  | [optional] 
**RequestTime** | Pointer to **int64** |  | [optional] 
**SchedulerConfig** | Pointer to [**NullableSchedulerConfiguration**](SchedulerConfiguration.md) |  | [optional] 

## Methods

### NewSchedulerConfigurationResponse

`func NewSchedulerConfigurationResponse() *SchedulerConfigurationResponse`

NewSchedulerConfigurationResponse instantiates a new SchedulerConfigurationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulerConfigurationResponseWithDefaults

`func NewSchedulerConfigurationResponseWithDefaults() *SchedulerConfigurationResponse`

NewSchedulerConfigurationResponseWithDefaults instantiates a new SchedulerConfigurationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKnownLeader

`func (o *SchedulerConfigurationResponse) GetKnownLeader() bool`

GetKnownLeader returns the KnownLeader field if non-nil, zero value otherwise.

### GetKnownLeaderOk

`func (o *SchedulerConfigurationResponse) GetKnownLeaderOk() (*bool, bool)`

GetKnownLeaderOk returns a tuple with the KnownLeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnownLeader

`func (o *SchedulerConfigurationResponse) SetKnownLeader(v bool)`

SetKnownLeader sets KnownLeader field to given value.

### HasKnownLeader

`func (o *SchedulerConfigurationResponse) HasKnownLeader() bool`

HasKnownLeader returns a boolean if a field has been set.

### GetLastContact

`func (o *SchedulerConfigurationResponse) GetLastContact() int64`

GetLastContact returns the LastContact field if non-nil, zero value otherwise.

### GetLastContactOk

`func (o *SchedulerConfigurationResponse) GetLastContactOk() (*int64, bool)`

GetLastContactOk returns a tuple with the LastContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastContact

`func (o *SchedulerConfigurationResponse) SetLastContact(v int64)`

SetLastContact sets LastContact field to given value.

### HasLastContact

`func (o *SchedulerConfigurationResponse) HasLastContact() bool`

HasLastContact returns a boolean if a field has been set.

### GetLastIndex

`func (o *SchedulerConfigurationResponse) GetLastIndex() int32`

GetLastIndex returns the LastIndex field if non-nil, zero value otherwise.

### GetLastIndexOk

`func (o *SchedulerConfigurationResponse) GetLastIndexOk() (*int32, bool)`

GetLastIndexOk returns a tuple with the LastIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastIndex

`func (o *SchedulerConfigurationResponse) SetLastIndex(v int32)`

SetLastIndex sets LastIndex field to given value.

### HasLastIndex

`func (o *SchedulerConfigurationResponse) HasLastIndex() bool`

HasLastIndex returns a boolean if a field has been set.

### GetNextToken

`func (o *SchedulerConfigurationResponse) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *SchedulerConfigurationResponse) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *SchedulerConfigurationResponse) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *SchedulerConfigurationResponse) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.

### GetRequestTime

`func (o *SchedulerConfigurationResponse) GetRequestTime() int64`

GetRequestTime returns the RequestTime field if non-nil, zero value otherwise.

### GetRequestTimeOk

`func (o *SchedulerConfigurationResponse) GetRequestTimeOk() (*int64, bool)`

GetRequestTimeOk returns a tuple with the RequestTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTime

`func (o *SchedulerConfigurationResponse) SetRequestTime(v int64)`

SetRequestTime sets RequestTime field to given value.

### HasRequestTime

`func (o *SchedulerConfigurationResponse) HasRequestTime() bool`

HasRequestTime returns a boolean if a field has been set.

### GetSchedulerConfig

`func (o *SchedulerConfigurationResponse) GetSchedulerConfig() SchedulerConfiguration`

GetSchedulerConfig returns the SchedulerConfig field if non-nil, zero value otherwise.

### GetSchedulerConfigOk

`func (o *SchedulerConfigurationResponse) GetSchedulerConfigOk() (*SchedulerConfiguration, bool)`

GetSchedulerConfigOk returns a tuple with the SchedulerConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedulerConfig

`func (o *SchedulerConfigurationResponse) SetSchedulerConfig(v SchedulerConfiguration)`

SetSchedulerConfig sets SchedulerConfig field to given value.

### HasSchedulerConfig

`func (o *SchedulerConfigurationResponse) HasSchedulerConfig() bool`

HasSchedulerConfig returns a boolean if a field has been set.

### SetSchedulerConfigNil

`func (o *SchedulerConfigurationResponse) SetSchedulerConfigNil(b bool)`

 SetSchedulerConfigNil sets the value for SchedulerConfig to be an explicit nil

### UnsetSchedulerConfig
`func (o *SchedulerConfigurationResponse) UnsetSchedulerConfig()`

UnsetSchedulerConfig ensures that no value is present for SchedulerConfig, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


