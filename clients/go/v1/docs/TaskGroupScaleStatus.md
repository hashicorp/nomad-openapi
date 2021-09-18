# TaskGroupScaleStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Desired** | Pointer to **int32** |  | [optional] 
**Events** | Pointer to [**[]ScalingEvent**](ScalingEvent.md) |  | [optional] 
**Healthy** | Pointer to **int32** |  | [optional] 
**Placed** | Pointer to **int32** |  | [optional] 
**Running** | Pointer to **int32** |  | [optional] 
**Unhealthy** | Pointer to **int32** |  | [optional] 

## Methods

### NewTaskGroupScaleStatus

`func NewTaskGroupScaleStatus() *TaskGroupScaleStatus`

NewTaskGroupScaleStatus instantiates a new TaskGroupScaleStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskGroupScaleStatusWithDefaults

`func NewTaskGroupScaleStatusWithDefaults() *TaskGroupScaleStatus`

NewTaskGroupScaleStatusWithDefaults instantiates a new TaskGroupScaleStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDesired

`func (o *TaskGroupScaleStatus) GetDesired() int32`

GetDesired returns the Desired field if non-nil, zero value otherwise.

### GetDesiredOk

`func (o *TaskGroupScaleStatus) GetDesiredOk() (*int32, bool)`

GetDesiredOk returns a tuple with the Desired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesired

`func (o *TaskGroupScaleStatus) SetDesired(v int32)`

SetDesired sets Desired field to given value.

### HasDesired

`func (o *TaskGroupScaleStatus) HasDesired() bool`

HasDesired returns a boolean if a field has been set.

### GetEvents

`func (o *TaskGroupScaleStatus) GetEvents() []ScalingEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *TaskGroupScaleStatus) GetEventsOk() (*[]ScalingEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *TaskGroupScaleStatus) SetEvents(v []ScalingEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *TaskGroupScaleStatus) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetHealthy

`func (o *TaskGroupScaleStatus) GetHealthy() int32`

GetHealthy returns the Healthy field if non-nil, zero value otherwise.

### GetHealthyOk

`func (o *TaskGroupScaleStatus) GetHealthyOk() (*int32, bool)`

GetHealthyOk returns a tuple with the Healthy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthy

`func (o *TaskGroupScaleStatus) SetHealthy(v int32)`

SetHealthy sets Healthy field to given value.

### HasHealthy

`func (o *TaskGroupScaleStatus) HasHealthy() bool`

HasHealthy returns a boolean if a field has been set.

### GetPlaced

`func (o *TaskGroupScaleStatus) GetPlaced() int32`

GetPlaced returns the Placed field if non-nil, zero value otherwise.

### GetPlacedOk

`func (o *TaskGroupScaleStatus) GetPlacedOk() (*int32, bool)`

GetPlacedOk returns a tuple with the Placed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaced

`func (o *TaskGroupScaleStatus) SetPlaced(v int32)`

SetPlaced sets Placed field to given value.

### HasPlaced

`func (o *TaskGroupScaleStatus) HasPlaced() bool`

HasPlaced returns a boolean if a field has been set.

### GetRunning

`func (o *TaskGroupScaleStatus) GetRunning() int32`

GetRunning returns the Running field if non-nil, zero value otherwise.

### GetRunningOk

`func (o *TaskGroupScaleStatus) GetRunningOk() (*int32, bool)`

GetRunningOk returns a tuple with the Running field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunning

`func (o *TaskGroupScaleStatus) SetRunning(v int32)`

SetRunning sets Running field to given value.

### HasRunning

`func (o *TaskGroupScaleStatus) HasRunning() bool`

HasRunning returns a boolean if a field has been set.

### GetUnhealthy

`func (o *TaskGroupScaleStatus) GetUnhealthy() int32`

GetUnhealthy returns the Unhealthy field if non-nil, zero value otherwise.

### GetUnhealthyOk

`func (o *TaskGroupScaleStatus) GetUnhealthyOk() (*int32, bool)`

GetUnhealthyOk returns a tuple with the Unhealthy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnhealthy

`func (o *TaskGroupScaleStatus) SetUnhealthy(v int32)`

SetUnhealthy sets Unhealthy field to given value.

### HasUnhealthy

`func (o *TaskGroupScaleStatus) HasUnhealthy() bool`

HasUnhealthy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


