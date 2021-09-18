# nomad-client.TaskGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**affinities** | [**[Affinity]**](Affinity.md) |  | [optional] 
**constraints** | [**[Constraint]**](Constraint.md) |  | [optional] 
**consul** | [**Consul**](Consul.md) |  | [optional] 
**count** | **Number** |  | [optional] 
**ephemeralDisk** | [**EphemeralDisk**](EphemeralDisk.md) |  | [optional] 
**meta** | **{String: String}** |  | [optional] 
**migrate** | [**MigrateStrategy**](MigrateStrategy.md) |  | [optional] 
**name** | **String** |  | [optional] 
**networks** | [**[NetworkResource]**](NetworkResource.md) |  | [optional] 
**reschedulePolicy** | [**ReschedulePolicy**](ReschedulePolicy.md) |  | [optional] 
**restartPolicy** | [**RestartPolicy**](RestartPolicy.md) |  | [optional] 
**scaling** | [**ScalingPolicy**](ScalingPolicy.md) |  | [optional] 
**services** | [**[Service]**](Service.md) |  | [optional] 
**shutdownDelay** | **Number** |  | [optional] 
**spreads** | [**[Spread]**](Spread.md) |  | [optional] 
**stopAfterClientDisconnect** | **Number** |  | [optional] 
**tasks** | [**[Task]**](Task.md) |  | [optional] 
**update** | [**UpdateStrategy**](UpdateStrategy.md) |  | [optional] 
**volumes** | [**{String: VolumeRequest}**](VolumeRequest.md) |  | [optional] 


