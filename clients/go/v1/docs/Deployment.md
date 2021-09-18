# Deployment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateIndex** | Pointer to **int32** |  | [optional] 
**ID** | Pointer to **string** |  | [optional] 
**IsMultiregion** | Pointer to **bool** |  | [optional] 
**JobCreateIndex** | Pointer to **int32** |  | [optional] 
**JobID** | Pointer to **string** |  | [optional] 
**JobModifyIndex** | Pointer to **int32** |  | [optional] 
**JobSpecModifyIndex** | Pointer to **int32** |  | [optional] 
**JobVersion** | Pointer to **int32** |  | [optional] 
**ModifyIndex** | Pointer to **int32** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusDescription** | Pointer to **string** |  | [optional] 
**TaskGroups** | Pointer to [**map[string]DeploymentState**](DeploymentState.md) |  | [optional] 

## Methods

### NewDeployment

`func NewDeployment() *Deployment`

NewDeployment instantiates a new Deployment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentWithDefaults

`func NewDeploymentWithDefaults() *Deployment`

NewDeploymentWithDefaults instantiates a new Deployment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateIndex

`func (o *Deployment) GetCreateIndex() int32`

GetCreateIndex returns the CreateIndex field if non-nil, zero value otherwise.

### GetCreateIndexOk

`func (o *Deployment) GetCreateIndexOk() (*int32, bool)`

GetCreateIndexOk returns a tuple with the CreateIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateIndex

`func (o *Deployment) SetCreateIndex(v int32)`

SetCreateIndex sets CreateIndex field to given value.

### HasCreateIndex

`func (o *Deployment) HasCreateIndex() bool`

HasCreateIndex returns a boolean if a field has been set.

### GetID

`func (o *Deployment) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *Deployment) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *Deployment) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *Deployment) HasID() bool`

HasID returns a boolean if a field has been set.

### GetIsMultiregion

`func (o *Deployment) GetIsMultiregion() bool`

GetIsMultiregion returns the IsMultiregion field if non-nil, zero value otherwise.

### GetIsMultiregionOk

`func (o *Deployment) GetIsMultiregionOk() (*bool, bool)`

GetIsMultiregionOk returns a tuple with the IsMultiregion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMultiregion

`func (o *Deployment) SetIsMultiregion(v bool)`

SetIsMultiregion sets IsMultiregion field to given value.

### HasIsMultiregion

`func (o *Deployment) HasIsMultiregion() bool`

HasIsMultiregion returns a boolean if a field has been set.

### GetJobCreateIndex

`func (o *Deployment) GetJobCreateIndex() int32`

GetJobCreateIndex returns the JobCreateIndex field if non-nil, zero value otherwise.

### GetJobCreateIndexOk

`func (o *Deployment) GetJobCreateIndexOk() (*int32, bool)`

GetJobCreateIndexOk returns a tuple with the JobCreateIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobCreateIndex

`func (o *Deployment) SetJobCreateIndex(v int32)`

SetJobCreateIndex sets JobCreateIndex field to given value.

### HasJobCreateIndex

`func (o *Deployment) HasJobCreateIndex() bool`

HasJobCreateIndex returns a boolean if a field has been set.

### GetJobID

`func (o *Deployment) GetJobID() string`

GetJobID returns the JobID field if non-nil, zero value otherwise.

### GetJobIDOk

`func (o *Deployment) GetJobIDOk() (*string, bool)`

GetJobIDOk returns a tuple with the JobID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobID

`func (o *Deployment) SetJobID(v string)`

SetJobID sets JobID field to given value.

### HasJobID

`func (o *Deployment) HasJobID() bool`

HasJobID returns a boolean if a field has been set.

### GetJobModifyIndex

`func (o *Deployment) GetJobModifyIndex() int32`

GetJobModifyIndex returns the JobModifyIndex field if non-nil, zero value otherwise.

### GetJobModifyIndexOk

`func (o *Deployment) GetJobModifyIndexOk() (*int32, bool)`

GetJobModifyIndexOk returns a tuple with the JobModifyIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobModifyIndex

`func (o *Deployment) SetJobModifyIndex(v int32)`

SetJobModifyIndex sets JobModifyIndex field to given value.

### HasJobModifyIndex

`func (o *Deployment) HasJobModifyIndex() bool`

HasJobModifyIndex returns a boolean if a field has been set.

### GetJobSpecModifyIndex

`func (o *Deployment) GetJobSpecModifyIndex() int32`

GetJobSpecModifyIndex returns the JobSpecModifyIndex field if non-nil, zero value otherwise.

### GetJobSpecModifyIndexOk

`func (o *Deployment) GetJobSpecModifyIndexOk() (*int32, bool)`

GetJobSpecModifyIndexOk returns a tuple with the JobSpecModifyIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobSpecModifyIndex

`func (o *Deployment) SetJobSpecModifyIndex(v int32)`

SetJobSpecModifyIndex sets JobSpecModifyIndex field to given value.

### HasJobSpecModifyIndex

`func (o *Deployment) HasJobSpecModifyIndex() bool`

HasJobSpecModifyIndex returns a boolean if a field has been set.

### GetJobVersion

`func (o *Deployment) GetJobVersion() int32`

GetJobVersion returns the JobVersion field if non-nil, zero value otherwise.

### GetJobVersionOk

`func (o *Deployment) GetJobVersionOk() (*int32, bool)`

GetJobVersionOk returns a tuple with the JobVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobVersion

`func (o *Deployment) SetJobVersion(v int32)`

SetJobVersion sets JobVersion field to given value.

### HasJobVersion

`func (o *Deployment) HasJobVersion() bool`

HasJobVersion returns a boolean if a field has been set.

### GetModifyIndex

`func (o *Deployment) GetModifyIndex() int32`

GetModifyIndex returns the ModifyIndex field if non-nil, zero value otherwise.

### GetModifyIndexOk

`func (o *Deployment) GetModifyIndexOk() (*int32, bool)`

GetModifyIndexOk returns a tuple with the ModifyIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifyIndex

`func (o *Deployment) SetModifyIndex(v int32)`

SetModifyIndex sets ModifyIndex field to given value.

### HasModifyIndex

`func (o *Deployment) HasModifyIndex() bool`

HasModifyIndex returns a boolean if a field has been set.

### GetNamespace

`func (o *Deployment) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *Deployment) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *Deployment) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *Deployment) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetStatus

`func (o *Deployment) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Deployment) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Deployment) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Deployment) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusDescription

`func (o *Deployment) GetStatusDescription() string`

GetStatusDescription returns the StatusDescription field if non-nil, zero value otherwise.

### GetStatusDescriptionOk

`func (o *Deployment) GetStatusDescriptionOk() (*string, bool)`

GetStatusDescriptionOk returns a tuple with the StatusDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDescription

`func (o *Deployment) SetStatusDescription(v string)`

SetStatusDescription sets StatusDescription field to given value.

### HasStatusDescription

`func (o *Deployment) HasStatusDescription() bool`

HasStatusDescription returns a boolean if a field has been set.

### GetTaskGroups

`func (o *Deployment) GetTaskGroups() map[string]DeploymentState`

GetTaskGroups returns the TaskGroups field if non-nil, zero value otherwise.

### GetTaskGroupsOk

`func (o *Deployment) GetTaskGroupsOk() (*map[string]DeploymentState, bool)`

GetTaskGroupsOk returns a tuple with the TaskGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskGroups

`func (o *Deployment) SetTaskGroups(v map[string]DeploymentState)`

SetTaskGroups sets TaskGroups field to given value.

### HasTaskGroups

`func (o *Deployment) HasTaskGroups() bool`

HasTaskGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


