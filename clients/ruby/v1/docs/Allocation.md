# Allocation


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**alloc_modify_index** | **int** |  | [optional] 
**allocated_resources** | [**AllocatedResources**](AllocatedResources.md) |  | [optional] 
**client_description** | **str** |  | [optional] 
**client_status** | **str** |  | [optional] 
**create_index** | **int** |  | [optional] 
**create_time** | **int** |  | [optional] 
**deployment_id** | **str** |  | [optional] 
**deployment_status** | [**AllocDeploymentStatus**](AllocDeploymentStatus.md) |  | [optional] 
**desired_description** | **str** |  | [optional] 
**desired_status** | **str** |  | [optional] 
**desired_transition** | [**DesiredTransition**](DesiredTransition.md) |  | [optional] 
**eval_id** | **str** |  | [optional] 
**followup_eval_id** | **str** |  | [optional] 
**id** | **str** |  | [optional] 
**job** | [**Job**](Job.md) |  | [optional] 
**job_id** | **str** |  | [optional] 
**metrics** | [**AllocationMetric**](AllocationMetric.md) |  | [optional] 
**modify_index** | **int** |  | [optional] 
**modify_time** | **int** |  | [optional] 
**name** | **str** |  | [optional] 
**namespace** | **str** |  | [optional] 
**next_allocation** | **str** |  | [optional] 
**node_id** | **str** |  | [optional] 
**node_name** | **str** |  | [optional] 
**preempted_allocations** | **[str]** |  | [optional] 
**preempted_by_allocation** | **str** |  | [optional] 
**previous_allocation** | **str** |  | [optional] 
**reschedule_tracker** | [**RescheduleTracker**](RescheduleTracker.md) |  | [optional] 
**resources** | [**Resources**](Resources.md) |  | [optional] 
**services** | **{str: (str,)}** |  | [optional] 
**task_group** | **str** |  | [optional] 
**task_resources** | [**{str: (Resources,)}**](Resources.md) |  | [optional] 
**task_states** | [**{str: (TaskState,)}**](TaskState.md) |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


