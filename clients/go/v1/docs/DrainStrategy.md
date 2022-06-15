# DrainStrategy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deadline** | Pointer to **int64** |  | [optional] 
**ForceDeadline** | Pointer to **NullableTime** |  | [optional] 
**IgnoreSystemJobs** | Pointer to **bool** |  | [optional] 
**StartedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewDrainStrategy

`func NewDrainStrategy() *DrainStrategy`

NewDrainStrategy instantiates a new DrainStrategy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDrainStrategyWithDefaults

`func NewDrainStrategyWithDefaults() *DrainStrategy`

NewDrainStrategyWithDefaults instantiates a new DrainStrategy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeadline

`func (o *DrainStrategy) GetDeadline() int64`

GetDeadline returns the Deadline field if non-nil, zero value otherwise.

### GetDeadlineOk

`func (o *DrainStrategy) GetDeadlineOk() (*int64, bool)`

GetDeadlineOk returns a tuple with the Deadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadline

`func (o *DrainStrategy) SetDeadline(v int64)`

SetDeadline sets Deadline field to given value.

### HasDeadline

`func (o *DrainStrategy) HasDeadline() bool`

HasDeadline returns a boolean if a field has been set.

### GetForceDeadline

`func (o *DrainStrategy) GetForceDeadline() time.Time`

GetForceDeadline returns the ForceDeadline field if non-nil, zero value otherwise.

### GetForceDeadlineOk

`func (o *DrainStrategy) GetForceDeadlineOk() (*time.Time, bool)`

GetForceDeadlineOk returns a tuple with the ForceDeadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceDeadline

`func (o *DrainStrategy) SetForceDeadline(v time.Time)`

SetForceDeadline sets ForceDeadline field to given value.

### HasForceDeadline

`func (o *DrainStrategy) HasForceDeadline() bool`

HasForceDeadline returns a boolean if a field has been set.

### SetForceDeadlineNil

`func (o *DrainStrategy) SetForceDeadlineNil(b bool)`

 SetForceDeadlineNil sets the value for ForceDeadline to be an explicit nil

### UnsetForceDeadline
`func (o *DrainStrategy) UnsetForceDeadline()`

UnsetForceDeadline ensures that no value is present for ForceDeadline, not even an explicit nil
### GetIgnoreSystemJobs

`func (o *DrainStrategy) GetIgnoreSystemJobs() bool`

GetIgnoreSystemJobs returns the IgnoreSystemJobs field if non-nil, zero value otherwise.

### GetIgnoreSystemJobsOk

`func (o *DrainStrategy) GetIgnoreSystemJobsOk() (*bool, bool)`

GetIgnoreSystemJobsOk returns a tuple with the IgnoreSystemJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreSystemJobs

`func (o *DrainStrategy) SetIgnoreSystemJobs(v bool)`

SetIgnoreSystemJobs sets IgnoreSystemJobs field to given value.

### HasIgnoreSystemJobs

`func (o *DrainStrategy) HasIgnoreSystemJobs() bool`

HasIgnoreSystemJobs returns a boolean if a field has been set.

### GetStartedAt

`func (o *DrainStrategy) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *DrainStrategy) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *DrainStrategy) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *DrainStrategy) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *DrainStrategy) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *DrainStrategy) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


