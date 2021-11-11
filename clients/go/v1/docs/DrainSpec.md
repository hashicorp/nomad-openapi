# DrainSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deadline** | Pointer to **int64** |  | [optional] 
**IgnoreSystemJobs** | Pointer to **bool** |  | [optional] 

## Methods

### NewDrainSpec

`func NewDrainSpec() *DrainSpec`

NewDrainSpec instantiates a new DrainSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDrainSpecWithDefaults

`func NewDrainSpecWithDefaults() *DrainSpec`

NewDrainSpecWithDefaults instantiates a new DrainSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeadline

`func (o *DrainSpec) GetDeadline() int64`

GetDeadline returns the Deadline field if non-nil, zero value otherwise.

### GetDeadlineOk

`func (o *DrainSpec) GetDeadlineOk() (*int64, bool)`

GetDeadlineOk returns a tuple with the Deadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadline

`func (o *DrainSpec) SetDeadline(v int64)`

SetDeadline sets Deadline field to given value.

### HasDeadline

`func (o *DrainSpec) HasDeadline() bool`

HasDeadline returns a boolean if a field has been set.

### GetIgnoreSystemJobs

`func (o *DrainSpec) GetIgnoreSystemJobs() bool`

GetIgnoreSystemJobs returns the IgnoreSystemJobs field if non-nil, zero value otherwise.

### GetIgnoreSystemJobsOk

`func (o *DrainSpec) GetIgnoreSystemJobsOk() (*bool, bool)`

GetIgnoreSystemJobsOk returns a tuple with the IgnoreSystemJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreSystemJobs

`func (o *DrainSpec) SetIgnoreSystemJobs(v bool)`

SetIgnoreSystemJobs sets IgnoreSystemJobs field to given value.

### HasIgnoreSystemJobs

`func (o *DrainSpec) HasIgnoreSystemJobs() bool`

HasIgnoreSystemJobs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


