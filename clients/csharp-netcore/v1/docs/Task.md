# Nomad.Client.Model.Task

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Affinities** | [**List&lt;Affinity&gt;**](Affinity.md) |  | [optional] 
**Artifacts** | [**List&lt;TaskArtifact&gt;**](TaskArtifact.md) |  | [optional] 
**CSIPluginConfig** | [**TaskCSIPluginConfig**](TaskCSIPluginConfig.md) |  | [optional] 
**Config** | **Dictionary&lt;string, Object&gt;** |  | [optional] 
**Constraints** | [**List&lt;Constraint&gt;**](Constraint.md) |  | [optional] 
**DispatchPayload** | [**DispatchPayloadConfig**](DispatchPayloadConfig.md) |  | [optional] 
**Driver** | **string** |  | [optional] 
**Env** | **Dictionary&lt;string, string&gt;** |  | [optional] 
**KillSignal** | **string** |  | [optional] 
**KillTimeout** | **long** |  | [optional] 
**Kind** | **string** |  | [optional] 
**Leader** | **bool** |  | [optional] 
**Lifecycle** | [**TaskLifecycle**](TaskLifecycle.md) |  | [optional] 
**LogConfig** | [**LogConfig**](LogConfig.md) |  | [optional] 
**Meta** | **Dictionary&lt;string, string&gt;** |  | [optional] 
**Name** | **string** |  | [optional] 
**Resources** | [**Resources**](Resources.md) |  | [optional] 
**RestartPolicy** | [**RestartPolicy**](RestartPolicy.md) |  | [optional] 
**ScalingPolicies** | [**List&lt;ScalingPolicy&gt;**](ScalingPolicy.md) |  | [optional] 
**Services** | [**List&lt;Service&gt;**](Service.md) |  | [optional] 
**ShutdownDelay** | **long** |  | [optional] 
**Templates** | [**List&lt;Template&gt;**](Template.md) |  | [optional] 
**User** | **string** |  | [optional] 
**Vault** | [**Vault**](Vault.md) |  | [optional] 
**VolumeMounts** | [**List&lt;VolumeMount&gt;**](VolumeMount.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

