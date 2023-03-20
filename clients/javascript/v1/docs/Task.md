# nomad-client.Task

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**affinities** | [**[Affinity]**](Affinity.md) |  | [optional] 
**artifacts** | [**[TaskArtifact]**](TaskArtifact.md) |  | [optional] 
**cSIPluginConfig** | [**TaskCSIPluginConfig**](TaskCSIPluginConfig.md) |  | [optional] 
**config** | **{String: Object}** |  | [optional] 
**constraints** | [**[Constraint]**](Constraint.md) |  | [optional] 
**dispatchPayload** | [**DispatchPayloadConfig**](DispatchPayloadConfig.md) |  | [optional] 
**driver** | **String** |  | [optional] 
**env** | **{String: String}** |  | [optional] 
**identity** | [**WorkloadIdentity**](WorkloadIdentity.md) |  | [optional] 
**killSignal** | **String** |  | [optional] 
**killTimeout** | **Number** |  | [optional] 
**kind** | **String** |  | [optional] 
**leader** | **Boolean** |  | [optional] 
**lifecycle** | [**TaskLifecycle**](TaskLifecycle.md) |  | [optional] 
**logConfig** | [**LogConfig**](LogConfig.md) |  | [optional] 
**meta** | **{String: String}** |  | [optional] 
**name** | **String** |  | [optional] 
**resources** | [**Resources**](Resources.md) |  | [optional] 
**restartPolicy** | [**RestartPolicy**](RestartPolicy.md) |  | [optional] 
**scalingPolicies** | [**[ScalingPolicy]**](ScalingPolicy.md) |  | [optional] 
**services** | [**[Service]**](Service.md) |  | [optional] 
**shutdownDelay** | **Number** |  | [optional] 
**templates** | [**[Template]**](Template.md) |  | [optional] 
**user** | **String** |  | [optional] 
**vault** | [**Vault**](Vault.md) |  | [optional] 
**volumeMounts** | [**[VolumeMount]**](VolumeMount.md) |  | [optional] 


