# TaskGroup


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**affinities** | [**[Affinity]**](Affinity.md) |  | [optional] 
**constraints** | [**[Constraint]**](Constraint.md) |  | [optional] 
**consul** | [**Consul**](Consul.md) |  | [optional] 
**count** | **int** |  | [optional] 
**ephemeral_disk** | [**EphemeralDisk**](EphemeralDisk.md) |  | [optional] 
**meta** | **{str: (str,)}** |  | [optional] 
**migrate** | [**MigrateStrategy**](MigrateStrategy.md) |  | [optional] 
**name** | **str** |  | [optional] 
**networks** | [**[NetworkResource]**](NetworkResource.md) |  | [optional] 
**reschedule_policy** | [**ReschedulePolicy**](ReschedulePolicy.md) |  | [optional] 
**restart_policy** | [**RestartPolicy**](RestartPolicy.md) |  | [optional] 
**scaling** | [**ScalingPolicy**](ScalingPolicy.md) |  | [optional] 
**services** | [**[Service]**](Service.md) |  | [optional] 
**shutdown_delay** | **int** |  | [optional] 
**spreads** | [**[Spread]**](Spread.md) |  | [optional] 
**stop_after_client_disconnect** | **int** |  | [optional] 
**tasks** | [**[Task]**](Task.md) |  | [optional] 
**update** | [**UpdateStrategy**](UpdateStrategy.md) |  | [optional] 
**volumes** | [**{str: (VolumeRequest,)}**](VolumeRequest.md) |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


