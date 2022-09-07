# Allocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllocModifyIndex** | Pointer to **int32** |  | [optional] 
**AllocatedResources** | Pointer to [**NullableAllocatedResources**](AllocatedResources.md) |  | [optional] 
**ClientDescription** | Pointer to **string** |  | [optional] 
**ClientStatus** | Pointer to **string** |  | [optional] 
**CreateIndex** | Pointer to **int32** |  | [optional] 
**CreateTime** | Pointer to **int64** |  | [optional] 
**DeploymentID** | Pointer to **string** |  | [optional] 
**DeploymentStatus** | Pointer to [**NullableAllocDeploymentStatus**](AllocDeploymentStatus.md) |  | [optional] 
**DesiredDescription** | Pointer to **string** |  | [optional] 
**DesiredStatus** | Pointer to **string** |  | [optional] 
**DesiredTransition** | Pointer to [**DesiredTransition**](DesiredTransition.md) |  | [optional] 
**EvalID** | Pointer to **string** |  | [optional] 
**FollowupEvalID** | Pointer to **string** |  | [optional] 
**ID** | Pointer to **string** |  | [optional] 
**Job** | Pointer to [**NullableJob**](Job.md) |  | [optional] 
**JobID** | Pointer to **string** |  | [optional] 
**Metrics** | Pointer to [**NullableAllocationMetric**](AllocationMetric.md) |  | [optional] 
**ModifyIndex** | Pointer to **int32** |  | [optional] 
**ModifyTime** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**NextAllocation** | Pointer to **string** |  | [optional] 
**NodeID** | Pointer to **string** |  | [optional] 
**NodeName** | Pointer to **string** |  | [optional] 
**PreemptedAllocations** | Pointer to **[]string** |  | [optional] 
**PreemptedByAllocation** | Pointer to **string** |  | [optional] 
**PreviousAllocation** | Pointer to **string** |  | [optional] 
**RescheduleTracker** | Pointer to [**NullableRescheduleTracker**](RescheduleTracker.md) |  | [optional] 
**Resources** | Pointer to [**NullableResources**](Resources.md) |  | [optional] 
**Services** | Pointer to **map[string]string** |  | [optional] 
**TaskGroup** | Pointer to **string** |  | [optional] 
**TaskResources** | Pointer to [**map[string]Resources**](Resources.md) |  | [optional] 
**TaskStates** | Pointer to [**map[string]TaskState**](TaskState.md) |  | [optional] 

## Methods

### NewAllocation

`func NewAllocation() *Allocation`

NewAllocation instantiates a new Allocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllocationWithDefaults

`func NewAllocationWithDefaults() *Allocation`

NewAllocationWithDefaults instantiates a new Allocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllocModifyIndex

`func (o *Allocation) GetAllocModifyIndex() int32`

GetAllocModifyIndex returns the AllocModifyIndex field if non-nil, zero value otherwise.

### GetAllocModifyIndexOk

`func (o *Allocation) GetAllocModifyIndexOk() (*int32, bool)`

GetAllocModifyIndexOk returns a tuple with the AllocModifyIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocModifyIndex

`func (o *Allocation) SetAllocModifyIndex(v int32)`

SetAllocModifyIndex sets AllocModifyIndex field to given value.

### HasAllocModifyIndex

`func (o *Allocation) HasAllocModifyIndex() bool`

HasAllocModifyIndex returns a boolean if a field has been set.

### GetAllocatedResources

`func (o *Allocation) GetAllocatedResources() AllocatedResources`

GetAllocatedResources returns the AllocatedResources field if non-nil, zero value otherwise.

### GetAllocatedResourcesOk

`func (o *Allocation) GetAllocatedResourcesOk() (*AllocatedResources, bool)`

GetAllocatedResourcesOk returns a tuple with the AllocatedResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedResources

`func (o *Allocation) SetAllocatedResources(v AllocatedResources)`

SetAllocatedResources sets AllocatedResources field to given value.

### HasAllocatedResources

`func (o *Allocation) HasAllocatedResources() bool`

HasAllocatedResources returns a boolean if a field has been set.

### SetAllocatedResourcesNil

`func (o *Allocation) SetAllocatedResourcesNil(b bool)`

 SetAllocatedResourcesNil sets the value for AllocatedResources to be an explicit nil

### UnsetAllocatedResources
`func (o *Allocation) UnsetAllocatedResources()`

UnsetAllocatedResources ensures that no value is present for AllocatedResources, not even an explicit nil
### GetClientDescription

`func (o *Allocation) GetClientDescription() string`

GetClientDescription returns the ClientDescription field if non-nil, zero value otherwise.

### GetClientDescriptionOk

`func (o *Allocation) GetClientDescriptionOk() (*string, bool)`

GetClientDescriptionOk returns a tuple with the ClientDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientDescription

`func (o *Allocation) SetClientDescription(v string)`

SetClientDescription sets ClientDescription field to given value.

### HasClientDescription

`func (o *Allocation) HasClientDescription() bool`

HasClientDescription returns a boolean if a field has been set.

### GetClientStatus

`func (o *Allocation) GetClientStatus() string`

GetClientStatus returns the ClientStatus field if non-nil, zero value otherwise.

### GetClientStatusOk

`func (o *Allocation) GetClientStatusOk() (*string, bool)`

GetClientStatusOk returns a tuple with the ClientStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientStatus

`func (o *Allocation) SetClientStatus(v string)`

SetClientStatus sets ClientStatus field to given value.

### HasClientStatus

`func (o *Allocation) HasClientStatus() bool`

HasClientStatus returns a boolean if a field has been set.

### GetCreateIndex

`func (o *Allocation) GetCreateIndex() int32`

GetCreateIndex returns the CreateIndex field if non-nil, zero value otherwise.

### GetCreateIndexOk

`func (o *Allocation) GetCreateIndexOk() (*int32, bool)`

GetCreateIndexOk returns a tuple with the CreateIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateIndex

`func (o *Allocation) SetCreateIndex(v int32)`

SetCreateIndex sets CreateIndex field to given value.

### HasCreateIndex

`func (o *Allocation) HasCreateIndex() bool`

HasCreateIndex returns a boolean if a field has been set.

### GetCreateTime

`func (o *Allocation) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *Allocation) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *Allocation) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *Allocation) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDeploymentID

`func (o *Allocation) GetDeploymentID() string`

GetDeploymentID returns the DeploymentID field if non-nil, zero value otherwise.

### GetDeploymentIDOk

`func (o *Allocation) GetDeploymentIDOk() (*string, bool)`

GetDeploymentIDOk returns a tuple with the DeploymentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentID

`func (o *Allocation) SetDeploymentID(v string)`

SetDeploymentID sets DeploymentID field to given value.

### HasDeploymentID

`func (o *Allocation) HasDeploymentID() bool`

HasDeploymentID returns a boolean if a field has been set.

### GetDeploymentStatus

`func (o *Allocation) GetDeploymentStatus() AllocDeploymentStatus`

GetDeploymentStatus returns the DeploymentStatus field if non-nil, zero value otherwise.

### GetDeploymentStatusOk

`func (o *Allocation) GetDeploymentStatusOk() (*AllocDeploymentStatus, bool)`

GetDeploymentStatusOk returns a tuple with the DeploymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentStatus

`func (o *Allocation) SetDeploymentStatus(v AllocDeploymentStatus)`

SetDeploymentStatus sets DeploymentStatus field to given value.

### HasDeploymentStatus

`func (o *Allocation) HasDeploymentStatus() bool`

HasDeploymentStatus returns a boolean if a field has been set.

### SetDeploymentStatusNil

`func (o *Allocation) SetDeploymentStatusNil(b bool)`

 SetDeploymentStatusNil sets the value for DeploymentStatus to be an explicit nil

### UnsetDeploymentStatus
`func (o *Allocation) UnsetDeploymentStatus()`

UnsetDeploymentStatus ensures that no value is present for DeploymentStatus, not even an explicit nil
### GetDesiredDescription

`func (o *Allocation) GetDesiredDescription() string`

GetDesiredDescription returns the DesiredDescription field if non-nil, zero value otherwise.

### GetDesiredDescriptionOk

`func (o *Allocation) GetDesiredDescriptionOk() (*string, bool)`

GetDesiredDescriptionOk returns a tuple with the DesiredDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredDescription

`func (o *Allocation) SetDesiredDescription(v string)`

SetDesiredDescription sets DesiredDescription field to given value.

### HasDesiredDescription

`func (o *Allocation) HasDesiredDescription() bool`

HasDesiredDescription returns a boolean if a field has been set.

### GetDesiredStatus

`func (o *Allocation) GetDesiredStatus() string`

GetDesiredStatus returns the DesiredStatus field if non-nil, zero value otherwise.

### GetDesiredStatusOk

`func (o *Allocation) GetDesiredStatusOk() (*string, bool)`

GetDesiredStatusOk returns a tuple with the DesiredStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredStatus

`func (o *Allocation) SetDesiredStatus(v string)`

SetDesiredStatus sets DesiredStatus field to given value.

### HasDesiredStatus

`func (o *Allocation) HasDesiredStatus() bool`

HasDesiredStatus returns a boolean if a field has been set.

### GetDesiredTransition

`func (o *Allocation) GetDesiredTransition() DesiredTransition`

GetDesiredTransition returns the DesiredTransition field if non-nil, zero value otherwise.

### GetDesiredTransitionOk

`func (o *Allocation) GetDesiredTransitionOk() (*DesiredTransition, bool)`

GetDesiredTransitionOk returns a tuple with the DesiredTransition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredTransition

`func (o *Allocation) SetDesiredTransition(v DesiredTransition)`

SetDesiredTransition sets DesiredTransition field to given value.

### HasDesiredTransition

`func (o *Allocation) HasDesiredTransition() bool`

HasDesiredTransition returns a boolean if a field has been set.

### GetEvalID

`func (o *Allocation) GetEvalID() string`

GetEvalID returns the EvalID field if non-nil, zero value otherwise.

### GetEvalIDOk

`func (o *Allocation) GetEvalIDOk() (*string, bool)`

GetEvalIDOk returns a tuple with the EvalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvalID

`func (o *Allocation) SetEvalID(v string)`

SetEvalID sets EvalID field to given value.

### HasEvalID

`func (o *Allocation) HasEvalID() bool`

HasEvalID returns a boolean if a field has been set.

### GetFollowupEvalID

`func (o *Allocation) GetFollowupEvalID() string`

GetFollowupEvalID returns the FollowupEvalID field if non-nil, zero value otherwise.

### GetFollowupEvalIDOk

`func (o *Allocation) GetFollowupEvalIDOk() (*string, bool)`

GetFollowupEvalIDOk returns a tuple with the FollowupEvalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowupEvalID

`func (o *Allocation) SetFollowupEvalID(v string)`

SetFollowupEvalID sets FollowupEvalID field to given value.

### HasFollowupEvalID

`func (o *Allocation) HasFollowupEvalID() bool`

HasFollowupEvalID returns a boolean if a field has been set.

### GetID

`func (o *Allocation) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *Allocation) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *Allocation) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *Allocation) HasID() bool`

HasID returns a boolean if a field has been set.

### GetJob

`func (o *Allocation) GetJob() Job`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *Allocation) GetJobOk() (*Job, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *Allocation) SetJob(v Job)`

SetJob sets Job field to given value.

### HasJob

`func (o *Allocation) HasJob() bool`

HasJob returns a boolean if a field has been set.

### SetJobNil

`func (o *Allocation) SetJobNil(b bool)`

 SetJobNil sets the value for Job to be an explicit nil

### UnsetJob
`func (o *Allocation) UnsetJob()`

UnsetJob ensures that no value is present for Job, not even an explicit nil
### GetJobID

`func (o *Allocation) GetJobID() string`

GetJobID returns the JobID field if non-nil, zero value otherwise.

### GetJobIDOk

`func (o *Allocation) GetJobIDOk() (*string, bool)`

GetJobIDOk returns a tuple with the JobID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobID

`func (o *Allocation) SetJobID(v string)`

SetJobID sets JobID field to given value.

### HasJobID

`func (o *Allocation) HasJobID() bool`

HasJobID returns a boolean if a field has been set.

### GetMetrics

`func (o *Allocation) GetMetrics() AllocationMetric`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *Allocation) GetMetricsOk() (*AllocationMetric, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *Allocation) SetMetrics(v AllocationMetric)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *Allocation) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### SetMetricsNil

`func (o *Allocation) SetMetricsNil(b bool)`

 SetMetricsNil sets the value for Metrics to be an explicit nil

### UnsetMetrics
`func (o *Allocation) UnsetMetrics()`

UnsetMetrics ensures that no value is present for Metrics, not even an explicit nil
### GetModifyIndex

`func (o *Allocation) GetModifyIndex() int32`

GetModifyIndex returns the ModifyIndex field if non-nil, zero value otherwise.

### GetModifyIndexOk

`func (o *Allocation) GetModifyIndexOk() (*int32, bool)`

GetModifyIndexOk returns a tuple with the ModifyIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifyIndex

`func (o *Allocation) SetModifyIndex(v int32)`

SetModifyIndex sets ModifyIndex field to given value.

### HasModifyIndex

`func (o *Allocation) HasModifyIndex() bool`

HasModifyIndex returns a boolean if a field has been set.

### GetModifyTime

`func (o *Allocation) GetModifyTime() int64`

GetModifyTime returns the ModifyTime field if non-nil, zero value otherwise.

### GetModifyTimeOk

`func (o *Allocation) GetModifyTimeOk() (*int64, bool)`

GetModifyTimeOk returns a tuple with the ModifyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifyTime

`func (o *Allocation) SetModifyTime(v int64)`

SetModifyTime sets ModifyTime field to given value.

### HasModifyTime

`func (o *Allocation) HasModifyTime() bool`

HasModifyTime returns a boolean if a field has been set.

### GetName

`func (o *Allocation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Allocation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Allocation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Allocation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *Allocation) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *Allocation) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *Allocation) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *Allocation) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetNextAllocation

`func (o *Allocation) GetNextAllocation() string`

GetNextAllocation returns the NextAllocation field if non-nil, zero value otherwise.

### GetNextAllocationOk

`func (o *Allocation) GetNextAllocationOk() (*string, bool)`

GetNextAllocationOk returns a tuple with the NextAllocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextAllocation

`func (o *Allocation) SetNextAllocation(v string)`

SetNextAllocation sets NextAllocation field to given value.

### HasNextAllocation

`func (o *Allocation) HasNextAllocation() bool`

HasNextAllocation returns a boolean if a field has been set.

### GetNodeID

`func (o *Allocation) GetNodeID() string`

GetNodeID returns the NodeID field if non-nil, zero value otherwise.

### GetNodeIDOk

`func (o *Allocation) GetNodeIDOk() (*string, bool)`

GetNodeIDOk returns a tuple with the NodeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeID

`func (o *Allocation) SetNodeID(v string)`

SetNodeID sets NodeID field to given value.

### HasNodeID

`func (o *Allocation) HasNodeID() bool`

HasNodeID returns a boolean if a field has been set.

### GetNodeName

`func (o *Allocation) GetNodeName() string`

GetNodeName returns the NodeName field if non-nil, zero value otherwise.

### GetNodeNameOk

`func (o *Allocation) GetNodeNameOk() (*string, bool)`

GetNodeNameOk returns a tuple with the NodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeName

`func (o *Allocation) SetNodeName(v string)`

SetNodeName sets NodeName field to given value.

### HasNodeName

`func (o *Allocation) HasNodeName() bool`

HasNodeName returns a boolean if a field has been set.

### GetPreemptedAllocations

`func (o *Allocation) GetPreemptedAllocations() []string`

GetPreemptedAllocations returns the PreemptedAllocations field if non-nil, zero value otherwise.

### GetPreemptedAllocationsOk

`func (o *Allocation) GetPreemptedAllocationsOk() (*[]string, bool)`

GetPreemptedAllocationsOk returns a tuple with the PreemptedAllocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreemptedAllocations

`func (o *Allocation) SetPreemptedAllocations(v []string)`

SetPreemptedAllocations sets PreemptedAllocations field to given value.

### HasPreemptedAllocations

`func (o *Allocation) HasPreemptedAllocations() bool`

HasPreemptedAllocations returns a boolean if a field has been set.

### GetPreemptedByAllocation

`func (o *Allocation) GetPreemptedByAllocation() string`

GetPreemptedByAllocation returns the PreemptedByAllocation field if non-nil, zero value otherwise.

### GetPreemptedByAllocationOk

`func (o *Allocation) GetPreemptedByAllocationOk() (*string, bool)`

GetPreemptedByAllocationOk returns a tuple with the PreemptedByAllocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreemptedByAllocation

`func (o *Allocation) SetPreemptedByAllocation(v string)`

SetPreemptedByAllocation sets PreemptedByAllocation field to given value.

### HasPreemptedByAllocation

`func (o *Allocation) HasPreemptedByAllocation() bool`

HasPreemptedByAllocation returns a boolean if a field has been set.

### GetPreviousAllocation

`func (o *Allocation) GetPreviousAllocation() string`

GetPreviousAllocation returns the PreviousAllocation field if non-nil, zero value otherwise.

### GetPreviousAllocationOk

`func (o *Allocation) GetPreviousAllocationOk() (*string, bool)`

GetPreviousAllocationOk returns a tuple with the PreviousAllocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousAllocation

`func (o *Allocation) SetPreviousAllocation(v string)`

SetPreviousAllocation sets PreviousAllocation field to given value.

### HasPreviousAllocation

`func (o *Allocation) HasPreviousAllocation() bool`

HasPreviousAllocation returns a boolean if a field has been set.

### GetRescheduleTracker

`func (o *Allocation) GetRescheduleTracker() RescheduleTracker`

GetRescheduleTracker returns the RescheduleTracker field if non-nil, zero value otherwise.

### GetRescheduleTrackerOk

`func (o *Allocation) GetRescheduleTrackerOk() (*RescheduleTracker, bool)`

GetRescheduleTrackerOk returns a tuple with the RescheduleTracker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRescheduleTracker

`func (o *Allocation) SetRescheduleTracker(v RescheduleTracker)`

SetRescheduleTracker sets RescheduleTracker field to given value.

### HasRescheduleTracker

`func (o *Allocation) HasRescheduleTracker() bool`

HasRescheduleTracker returns a boolean if a field has been set.

### SetRescheduleTrackerNil

`func (o *Allocation) SetRescheduleTrackerNil(b bool)`

 SetRescheduleTrackerNil sets the value for RescheduleTracker to be an explicit nil

### UnsetRescheduleTracker
`func (o *Allocation) UnsetRescheduleTracker()`

UnsetRescheduleTracker ensures that no value is present for RescheduleTracker, not even an explicit nil
### GetResources

`func (o *Allocation) GetResources() Resources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *Allocation) GetResourcesOk() (*Resources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *Allocation) SetResources(v Resources)`

SetResources sets Resources field to given value.

### HasResources

`func (o *Allocation) HasResources() bool`

HasResources returns a boolean if a field has been set.

### SetResourcesNil

`func (o *Allocation) SetResourcesNil(b bool)`

 SetResourcesNil sets the value for Resources to be an explicit nil

### UnsetResources
`func (o *Allocation) UnsetResources()`

UnsetResources ensures that no value is present for Resources, not even an explicit nil
### GetServices

`func (o *Allocation) GetServices() map[string]string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *Allocation) GetServicesOk() (*map[string]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *Allocation) SetServices(v map[string]string)`

SetServices sets Services field to given value.

### HasServices

`func (o *Allocation) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetTaskGroup

`func (o *Allocation) GetTaskGroup() string`

GetTaskGroup returns the TaskGroup field if non-nil, zero value otherwise.

### GetTaskGroupOk

`func (o *Allocation) GetTaskGroupOk() (*string, bool)`

GetTaskGroupOk returns a tuple with the TaskGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskGroup

`func (o *Allocation) SetTaskGroup(v string)`

SetTaskGroup sets TaskGroup field to given value.

### HasTaskGroup

`func (o *Allocation) HasTaskGroup() bool`

HasTaskGroup returns a boolean if a field has been set.

### GetTaskResources

`func (o *Allocation) GetTaskResources() map[string]Resources`

GetTaskResources returns the TaskResources field if non-nil, zero value otherwise.

### GetTaskResourcesOk

`func (o *Allocation) GetTaskResourcesOk() (*map[string]Resources, bool)`

GetTaskResourcesOk returns a tuple with the TaskResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskResources

`func (o *Allocation) SetTaskResources(v map[string]Resources)`

SetTaskResources sets TaskResources field to given value.

### HasTaskResources

`func (o *Allocation) HasTaskResources() bool`

HasTaskResources returns a boolean if a field has been set.

### GetTaskStates

`func (o *Allocation) GetTaskStates() map[string]TaskState`

GetTaskStates returns the TaskStates field if non-nil, zero value otherwise.

### GetTaskStatesOk

`func (o *Allocation) GetTaskStatesOk() (*map[string]TaskState, bool)`

GetTaskStatesOk returns a tuple with the TaskStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskStates

`func (o *Allocation) SetTaskStates(v map[string]TaskState)`

SetTaskStates sets TaskStates field to given value.

### HasTaskStates

`func (o *Allocation) HasTaskStates() bool`

HasTaskStates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


