# EvaluationStub

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockedEval** | Pointer to **string** |  | [optional] 
**CreateIndex** | Pointer to **int32** |  | [optional] 
**CreateTime** | Pointer to **int64** |  | [optional] 
**DeploymentID** | Pointer to **string** |  | [optional] 
**ID** | Pointer to **string** |  | [optional] 
**JobID** | Pointer to **string** |  | [optional] 
**ModifyIndex** | Pointer to **int32** |  | [optional] 
**ModifyTime** | Pointer to **int64** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**NextEval** | Pointer to **string** |  | [optional] 
**NodeID** | Pointer to **string** |  | [optional] 
**PreviousEval** | Pointer to **string** |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusDescription** | Pointer to **string** |  | [optional] 
**TriggeredBy** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**WaitUntil** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewEvaluationStub

`func NewEvaluationStub() *EvaluationStub`

NewEvaluationStub instantiates a new EvaluationStub object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvaluationStubWithDefaults

`func NewEvaluationStubWithDefaults() *EvaluationStub`

NewEvaluationStubWithDefaults instantiates a new EvaluationStub object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockedEval

`func (o *EvaluationStub) GetBlockedEval() string`

GetBlockedEval returns the BlockedEval field if non-nil, zero value otherwise.

### GetBlockedEvalOk

`func (o *EvaluationStub) GetBlockedEvalOk() (*string, bool)`

GetBlockedEvalOk returns a tuple with the BlockedEval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedEval

`func (o *EvaluationStub) SetBlockedEval(v string)`

SetBlockedEval sets BlockedEval field to given value.

### HasBlockedEval

`func (o *EvaluationStub) HasBlockedEval() bool`

HasBlockedEval returns a boolean if a field has been set.

### GetCreateIndex

`func (o *EvaluationStub) GetCreateIndex() int32`

GetCreateIndex returns the CreateIndex field if non-nil, zero value otherwise.

### GetCreateIndexOk

`func (o *EvaluationStub) GetCreateIndexOk() (*int32, bool)`

GetCreateIndexOk returns a tuple with the CreateIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateIndex

`func (o *EvaluationStub) SetCreateIndex(v int32)`

SetCreateIndex sets CreateIndex field to given value.

### HasCreateIndex

`func (o *EvaluationStub) HasCreateIndex() bool`

HasCreateIndex returns a boolean if a field has been set.

### GetCreateTime

`func (o *EvaluationStub) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *EvaluationStub) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *EvaluationStub) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *EvaluationStub) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDeploymentID

`func (o *EvaluationStub) GetDeploymentID() string`

GetDeploymentID returns the DeploymentID field if non-nil, zero value otherwise.

### GetDeploymentIDOk

`func (o *EvaluationStub) GetDeploymentIDOk() (*string, bool)`

GetDeploymentIDOk returns a tuple with the DeploymentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentID

`func (o *EvaluationStub) SetDeploymentID(v string)`

SetDeploymentID sets DeploymentID field to given value.

### HasDeploymentID

`func (o *EvaluationStub) HasDeploymentID() bool`

HasDeploymentID returns a boolean if a field has been set.

### GetID

`func (o *EvaluationStub) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *EvaluationStub) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *EvaluationStub) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *EvaluationStub) HasID() bool`

HasID returns a boolean if a field has been set.

### GetJobID

`func (o *EvaluationStub) GetJobID() string`

GetJobID returns the JobID field if non-nil, zero value otherwise.

### GetJobIDOk

`func (o *EvaluationStub) GetJobIDOk() (*string, bool)`

GetJobIDOk returns a tuple with the JobID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobID

`func (o *EvaluationStub) SetJobID(v string)`

SetJobID sets JobID field to given value.

### HasJobID

`func (o *EvaluationStub) HasJobID() bool`

HasJobID returns a boolean if a field has been set.

### GetModifyIndex

`func (o *EvaluationStub) GetModifyIndex() int32`

GetModifyIndex returns the ModifyIndex field if non-nil, zero value otherwise.

### GetModifyIndexOk

`func (o *EvaluationStub) GetModifyIndexOk() (*int32, bool)`

GetModifyIndexOk returns a tuple with the ModifyIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifyIndex

`func (o *EvaluationStub) SetModifyIndex(v int32)`

SetModifyIndex sets ModifyIndex field to given value.

### HasModifyIndex

`func (o *EvaluationStub) HasModifyIndex() bool`

HasModifyIndex returns a boolean if a field has been set.

### GetModifyTime

`func (o *EvaluationStub) GetModifyTime() int64`

GetModifyTime returns the ModifyTime field if non-nil, zero value otherwise.

### GetModifyTimeOk

`func (o *EvaluationStub) GetModifyTimeOk() (*int64, bool)`

GetModifyTimeOk returns a tuple with the ModifyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifyTime

`func (o *EvaluationStub) SetModifyTime(v int64)`

SetModifyTime sets ModifyTime field to given value.

### HasModifyTime

`func (o *EvaluationStub) HasModifyTime() bool`

HasModifyTime returns a boolean if a field has been set.

### GetNamespace

`func (o *EvaluationStub) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *EvaluationStub) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *EvaluationStub) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *EvaluationStub) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetNextEval

`func (o *EvaluationStub) GetNextEval() string`

GetNextEval returns the NextEval field if non-nil, zero value otherwise.

### GetNextEvalOk

`func (o *EvaluationStub) GetNextEvalOk() (*string, bool)`

GetNextEvalOk returns a tuple with the NextEval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextEval

`func (o *EvaluationStub) SetNextEval(v string)`

SetNextEval sets NextEval field to given value.

### HasNextEval

`func (o *EvaluationStub) HasNextEval() bool`

HasNextEval returns a boolean if a field has been set.

### GetNodeID

`func (o *EvaluationStub) GetNodeID() string`

GetNodeID returns the NodeID field if non-nil, zero value otherwise.

### GetNodeIDOk

`func (o *EvaluationStub) GetNodeIDOk() (*string, bool)`

GetNodeIDOk returns a tuple with the NodeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeID

`func (o *EvaluationStub) SetNodeID(v string)`

SetNodeID sets NodeID field to given value.

### HasNodeID

`func (o *EvaluationStub) HasNodeID() bool`

HasNodeID returns a boolean if a field has been set.

### GetPreviousEval

`func (o *EvaluationStub) GetPreviousEval() string`

GetPreviousEval returns the PreviousEval field if non-nil, zero value otherwise.

### GetPreviousEvalOk

`func (o *EvaluationStub) GetPreviousEvalOk() (*string, bool)`

GetPreviousEvalOk returns a tuple with the PreviousEval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousEval

`func (o *EvaluationStub) SetPreviousEval(v string)`

SetPreviousEval sets PreviousEval field to given value.

### HasPreviousEval

`func (o *EvaluationStub) HasPreviousEval() bool`

HasPreviousEval returns a boolean if a field has been set.

### GetPriority

`func (o *EvaluationStub) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *EvaluationStub) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *EvaluationStub) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *EvaluationStub) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetStatus

`func (o *EvaluationStub) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EvaluationStub) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EvaluationStub) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EvaluationStub) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusDescription

`func (o *EvaluationStub) GetStatusDescription() string`

GetStatusDescription returns the StatusDescription field if non-nil, zero value otherwise.

### GetStatusDescriptionOk

`func (o *EvaluationStub) GetStatusDescriptionOk() (*string, bool)`

GetStatusDescriptionOk returns a tuple with the StatusDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDescription

`func (o *EvaluationStub) SetStatusDescription(v string)`

SetStatusDescription sets StatusDescription field to given value.

### HasStatusDescription

`func (o *EvaluationStub) HasStatusDescription() bool`

HasStatusDescription returns a boolean if a field has been set.

### GetTriggeredBy

`func (o *EvaluationStub) GetTriggeredBy() string`

GetTriggeredBy returns the TriggeredBy field if non-nil, zero value otherwise.

### GetTriggeredByOk

`func (o *EvaluationStub) GetTriggeredByOk() (*string, bool)`

GetTriggeredByOk returns a tuple with the TriggeredBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeredBy

`func (o *EvaluationStub) SetTriggeredBy(v string)`

SetTriggeredBy sets TriggeredBy field to given value.

### HasTriggeredBy

`func (o *EvaluationStub) HasTriggeredBy() bool`

HasTriggeredBy returns a boolean if a field has been set.

### GetType

`func (o *EvaluationStub) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EvaluationStub) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EvaluationStub) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EvaluationStub) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWaitUntil

`func (o *EvaluationStub) GetWaitUntil() time.Time`

GetWaitUntil returns the WaitUntil field if non-nil, zero value otherwise.

### GetWaitUntilOk

`func (o *EvaluationStub) GetWaitUntilOk() (*time.Time, bool)`

GetWaitUntilOk returns a tuple with the WaitUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitUntil

`func (o *EvaluationStub) SetWaitUntil(v time.Time)`

SetWaitUntil sets WaitUntil field to given value.

### HasWaitUntil

`func (o *EvaluationStub) HasWaitUntil() bool`

HasWaitUntil returns a boolean if a field has been set.

### SetWaitUntilNil

`func (o *EvaluationStub) SetWaitUntilNil(b bool)`

 SetWaitUntilNil sets the value for WaitUntil to be an explicit nil

### UnsetWaitUntil
`func (o *EvaluationStub) UnsetWaitUntil()`

UnsetWaitUntil ensures that no value is present for WaitUntil, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


