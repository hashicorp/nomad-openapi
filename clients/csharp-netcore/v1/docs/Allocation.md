# Nomad.Client.Model.Allocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllocModifyIndex** | **int** |  | [optional] 
**AllocatedResources** | [**AllocatedResources**](AllocatedResources.md) |  | [optional] 
**ClientDescription** | **string** |  | [optional] 
**ClientStatus** | **string** |  | [optional] 
**CreateIndex** | **int** |  | [optional] 
**CreateTime** | **long** |  | [optional] 
**DeploymentID** | **string** |  | [optional] 
**DeploymentStatus** | [**AllocDeploymentStatus**](AllocDeploymentStatus.md) |  | [optional] 
**DesiredDescription** | **string** |  | [optional] 
**DesiredStatus** | **string** |  | [optional] 
**DesiredTransition** | [**DesiredTransition**](DesiredTransition.md) |  | [optional] 
**EvalID** | **string** |  | [optional] 
**FollowupEvalID** | **string** |  | [optional] 
**ID** | **string** |  | [optional] 
**Job** | [**Job**](Job.md) |  | [optional] 
**JobID** | **string** |  | [optional] 
**Metrics** | [**AllocationMetric**](AllocationMetric.md) |  | [optional] 
**ModifyIndex** | **int** |  | [optional] 
**ModifyTime** | **long** |  | [optional] 
**Name** | **string** |  | [optional] 
**Namespace** | **string** |  | [optional] 
**NextAllocation** | **string** |  | [optional] 
**NodeID** | **string** |  | [optional] 
**NodeName** | **string** |  | [optional] 
**PreemptedAllocations** | **List&lt;string&gt;** |  | [optional] 
**PreemptedByAllocation** | **string** |  | [optional] 
**PreviousAllocation** | **string** |  | [optional] 
**RescheduleTracker** | [**RescheduleTracker**](RescheduleTracker.md) |  | [optional] 
**Resources** | [**Resources**](Resources.md) |  | [optional] 
**Services** | **Dictionary&lt;string, string&gt;** |  | [optional] 
**TaskGroup** | **string** |  | [optional] 
**TaskResources** | [**Dictionary&lt;string, Resources&gt;**](Resources.md) |  | [optional] 
**TaskStates** | [**Dictionary&lt;string, TaskState&gt;**](TaskState.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

