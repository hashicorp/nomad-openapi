# JobScaleStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobCreateIndex** | Pointer to **int32** |  | [optional] 
**JobID** | Pointer to **string** |  | [optional] 
**JobModifyIndex** | Pointer to **int32** |  | [optional] 
**JobStopped** | Pointer to **bool** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**TaskGroups** | Pointer to [**map[string]TaskGroupScaleStatus**](TaskGroupScaleStatus.md) |  | [optional] 

## Methods

### NewJobScaleStatusResponse

`func NewJobScaleStatusResponse() *JobScaleStatusResponse`

NewJobScaleStatusResponse instantiates a new JobScaleStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobScaleStatusResponseWithDefaults

`func NewJobScaleStatusResponseWithDefaults() *JobScaleStatusResponse`

NewJobScaleStatusResponseWithDefaults instantiates a new JobScaleStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobCreateIndex

`func (o *JobScaleStatusResponse) GetJobCreateIndex() int32`

GetJobCreateIndex returns the JobCreateIndex field if non-nil, zero value otherwise.

### GetJobCreateIndexOk

`func (o *JobScaleStatusResponse) GetJobCreateIndexOk() (*int32, bool)`

GetJobCreateIndexOk returns a tuple with the JobCreateIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobCreateIndex

`func (o *JobScaleStatusResponse) SetJobCreateIndex(v int32)`

SetJobCreateIndex sets JobCreateIndex field to given value.

### HasJobCreateIndex

`func (o *JobScaleStatusResponse) HasJobCreateIndex() bool`

HasJobCreateIndex returns a boolean if a field has been set.

### GetJobID

`func (o *JobScaleStatusResponse) GetJobID() string`

GetJobID returns the JobID field if non-nil, zero value otherwise.

### GetJobIDOk

`func (o *JobScaleStatusResponse) GetJobIDOk() (*string, bool)`

GetJobIDOk returns a tuple with the JobID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobID

`func (o *JobScaleStatusResponse) SetJobID(v string)`

SetJobID sets JobID field to given value.

### HasJobID

`func (o *JobScaleStatusResponse) HasJobID() bool`

HasJobID returns a boolean if a field has been set.

### GetJobModifyIndex

`func (o *JobScaleStatusResponse) GetJobModifyIndex() int32`

GetJobModifyIndex returns the JobModifyIndex field if non-nil, zero value otherwise.

### GetJobModifyIndexOk

`func (o *JobScaleStatusResponse) GetJobModifyIndexOk() (*int32, bool)`

GetJobModifyIndexOk returns a tuple with the JobModifyIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobModifyIndex

`func (o *JobScaleStatusResponse) SetJobModifyIndex(v int32)`

SetJobModifyIndex sets JobModifyIndex field to given value.

### HasJobModifyIndex

`func (o *JobScaleStatusResponse) HasJobModifyIndex() bool`

HasJobModifyIndex returns a boolean if a field has been set.

### GetJobStopped

`func (o *JobScaleStatusResponse) GetJobStopped() bool`

GetJobStopped returns the JobStopped field if non-nil, zero value otherwise.

### GetJobStoppedOk

`func (o *JobScaleStatusResponse) GetJobStoppedOk() (*bool, bool)`

GetJobStoppedOk returns a tuple with the JobStopped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobStopped

`func (o *JobScaleStatusResponse) SetJobStopped(v bool)`

SetJobStopped sets JobStopped field to given value.

### HasJobStopped

`func (o *JobScaleStatusResponse) HasJobStopped() bool`

HasJobStopped returns a boolean if a field has been set.

### GetNamespace

`func (o *JobScaleStatusResponse) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *JobScaleStatusResponse) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *JobScaleStatusResponse) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *JobScaleStatusResponse) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetTaskGroups

`func (o *JobScaleStatusResponse) GetTaskGroups() map[string]TaskGroupScaleStatus`

GetTaskGroups returns the TaskGroups field if non-nil, zero value otherwise.

### GetTaskGroupsOk

`func (o *JobScaleStatusResponse) GetTaskGroupsOk() (*map[string]TaskGroupScaleStatus, bool)`

GetTaskGroupsOk returns a tuple with the TaskGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskGroups

`func (o *JobScaleStatusResponse) SetTaskGroups(v map[string]TaskGroupScaleStatus)`

SetTaskGroups sets TaskGroups field to given value.

### HasTaskGroups

`func (o *JobScaleStatusResponse) HasTaskGroups() bool`

HasTaskGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


