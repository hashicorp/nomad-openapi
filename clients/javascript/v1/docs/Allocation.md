# nomad-client.Allocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**allocModifyIndex** | **Number** |  | [optional] 
**allocatedResources** | [**AllocatedResources**](AllocatedResources.md) |  | [optional] 
**clientDescription** | **String** |  | [optional] 
**clientStatus** | **String** |  | [optional] 
**createIndex** | **Number** |  | [optional] 
**createTime** | **Number** |  | [optional] 
**deploymentID** | **String** |  | [optional] 
**deploymentStatus** | [**AllocDeploymentStatus**](AllocDeploymentStatus.md) |  | [optional] 
**desiredDescription** | **String** |  | [optional] 
**desiredStatus** | **String** |  | [optional] 
**desiredTransition** | [**DesiredTransition**](DesiredTransition.md) |  | [optional] 
**evalID** | **String** |  | [optional] 
**followupEvalID** | **String** |  | [optional] 
**ID** | **String** |  | [optional] 
**job** | [**Job**](Job.md) |  | [optional] 
**jobID** | **String** |  | [optional] 
**metrics** | [**AllocationMetric**](AllocationMetric.md) |  | [optional] 
**modifyIndex** | **Number** |  | [optional] 
**modifyTime** | **Number** |  | [optional] 
**name** | **String** |  | [optional] 
**namespace** | **String** |  | [optional] 
**nextAllocation** | **String** |  | [optional] 
**nodeID** | **String** |  | [optional] 
**nodeName** | **String** |  | [optional] 
**preemptedAllocations** | **[String]** |  | [optional] 
**preemptedByAllocation** | **String** |  | [optional] 
**previousAllocation** | **String** |  | [optional] 
**rescheduleTracker** | [**RescheduleTracker**](RescheduleTracker.md) |  | [optional] 
**resources** | [**Resources**](Resources.md) |  | [optional] 
**services** | **{String: String}** |  | [optional] 
**taskGroup** | **String** |  | [optional] 
**taskResources** | [**{String: Resources}**](Resources.md) |  | [optional] 
**taskStates** | [**{String: TaskState}**](TaskState.md) |  | [optional] 


