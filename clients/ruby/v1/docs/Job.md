# Job


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**affinities** | [**[Affinity]**](Affinity.md) |  | [optional] 
**all_at_once** | **bool** |  | [optional] 
**constraints** | [**[Constraint]**](Constraint.md) |  | [optional] 
**consul_namespace** | **str** |  | [optional] 
**consul_token** | **str** |  | [optional] 
**create_index** | **int** |  | [optional] 
**datacenters** | **[str]** |  | [optional] 
**dispatched** | **bool** |  | [optional] 
**id** | **str** |  | [optional] 
**job_modify_index** | **int** |  | [optional] 
**meta** | **{str: (str,)}** |  | [optional] 
**migrate** | [**MigrateStrategy**](MigrateStrategy.md) |  | [optional] 
**modify_index** | **int** |  | [optional] 
**multiregion** | [**Multiregion**](Multiregion.md) |  | [optional] 
**name** | **str** |  | [optional] 
**namespace** | **str** |  | [optional] 
**nomad_token_id** | **str** |  | [optional] 
**parameterized_job** | [**ParameterizedJobConfig**](ParameterizedJobConfig.md) |  | [optional] 
**parent_id** | **str** |  | [optional] 
**payload** | **str** |  | [optional] 
**periodic** | [**PeriodicConfig**](PeriodicConfig.md) |  | [optional] 
**priority** | **int** |  | [optional] 
**region** | **str** |  | [optional] 
**reschedule** | [**ReschedulePolicy**](ReschedulePolicy.md) |  | [optional] 
**spreads** | [**[Spread]**](Spread.md) |  | [optional] 
**stable** | **bool** |  | [optional] 
**status** | **str** |  | [optional] 
**status_description** | **str** |  | [optional] 
**stop** | **bool** |  | [optional] 
**submit_time** | **int** |  | [optional] 
**task_groups** | [**[TaskGroup]**](TaskGroup.md) |  | [optional] 
**type** | **str** |  | [optional] 
**update** | [**UpdateStrategy**](UpdateStrategy.md) |  | [optional] 
**vault_namespace** | **str** |  | [optional] 
**vault_token** | **str** |  | [optional] 
**version** | **int** |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


