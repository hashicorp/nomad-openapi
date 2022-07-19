# Nomad.Client.Model.TaskGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Affinities** | [**List&lt;Affinity&gt;**](Affinity.md) |  | [optional] 
**Constraints** | [**List&lt;Constraint&gt;**](Constraint.md) |  | [optional] 
**Consul** | [**Consul**](Consul.md) |  | [optional] 
**Count** | **int** |  | [optional] 
**EphemeralDisk** | [**EphemeralDisk**](EphemeralDisk.md) |  | [optional] 
**MaxClientDisconnect** | **long** |  | [optional] 
**Meta** | **Dictionary&lt;string, string&gt;** |  | [optional] 
**Migrate** | [**MigrateStrategy**](MigrateStrategy.md) |  | [optional] 
**Name** | **string** |  | [optional] 
**Networks** | [**List&lt;NetworkResource&gt;**](NetworkResource.md) |  | [optional] 
**ReschedulePolicy** | [**ReschedulePolicy**](ReschedulePolicy.md) |  | [optional] 
**RestartPolicy** | [**RestartPolicy**](RestartPolicy.md) |  | [optional] 
**Scaling** | [**ScalingPolicy**](ScalingPolicy.md) |  | [optional] 
**Services** | [**List&lt;Service&gt;**](Service.md) |  | [optional] 
**ShutdownDelay** | **long** |  | [optional] 
**Spreads** | [**List&lt;Spread&gt;**](Spread.md) |  | [optional] 
**StopAfterClientDisconnect** | **long** |  | [optional] 
**Tasks** | [**List&lt;Task&gt;**](Task.md) |  | [optional] 
**Update** | [**UpdateStrategy**](UpdateStrategy.md) |  | [optional] 
**Volumes** | [**Dictionary&lt;string, VolumeRequest&gt;**](VolumeRequest.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

