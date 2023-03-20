# Task


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**affinities** | [**[Affinity]**](Affinity.md) |  | [optional] 
**artifacts** | [**[TaskArtifact]**](TaskArtifact.md) |  | [optional] 
**csi_plugin_config** | [**TaskCSIPluginConfig**](TaskCSIPluginConfig.md) |  | [optional] 
**config** | **{str: (bool, date, datetime, dict, float, int, list, str, none_type)}** |  | [optional] 
**constraints** | [**[Constraint]**](Constraint.md) |  | [optional] 
**dispatch_payload** | [**DispatchPayloadConfig**](DispatchPayloadConfig.md) |  | [optional] 
**driver** | **str** |  | [optional] 
**env** | **{str: (str,)}** |  | [optional] 
**identity** | [**WorkloadIdentity**](WorkloadIdentity.md) |  | [optional] 
**kill_signal** | **str** |  | [optional] 
**kill_timeout** | **int** |  | [optional] 
**kind** | **str** |  | [optional] 
**leader** | **bool** |  | [optional] 
**lifecycle** | [**TaskLifecycle**](TaskLifecycle.md) |  | [optional] 
**log_config** | [**LogConfig**](LogConfig.md) |  | [optional] 
**meta** | **{str: (str,)}** |  | [optional] 
**name** | **str** |  | [optional] 
**resources** | [**Resources**](Resources.md) |  | [optional] 
**restart_policy** | [**RestartPolicy**](RestartPolicy.md) |  | [optional] 
**scaling_policies** | [**[ScalingPolicy]**](ScalingPolicy.md) |  | [optional] 
**services** | [**[Service]**](Service.md) |  | [optional] 
**shutdown_delay** | **int** |  | [optional] 
**templates** | [**[Template]**](Template.md) |  | [optional] 
**user** | **str** |  | [optional] 
**vault** | [**Vault**](Vault.md) |  | [optional] 
**volume_mounts** | [**[VolumeMount]**](VolumeMount.md) |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


