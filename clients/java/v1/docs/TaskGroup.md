

# TaskGroup


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**affinities** | [**List&lt;Affinity&gt;**](Affinity.md) |  |  [optional] |
|**constraints** | [**List&lt;Constraint&gt;**](Constraint.md) |  |  [optional] |
|**consul** | [**Consul**](Consul.md) |  |  [optional] |
|**count** | **Integer** |  |  [optional] |
|**ephemeralDisk** | [**EphemeralDisk**](EphemeralDisk.md) |  |  [optional] |
|**maxClientDisconnect** | **Long** |  |  [optional] |
|**meta** | **Map&lt;String, String&gt;** |  |  [optional] |
|**migrate** | [**MigrateStrategy**](MigrateStrategy.md) |  |  [optional] |
|**name** | **String** |  |  [optional] |
|**networks** | [**List&lt;NetworkResource&gt;**](NetworkResource.md) |  |  [optional] |
|**reschedulePolicy** | [**ReschedulePolicy**](ReschedulePolicy.md) |  |  [optional] |
|**restartPolicy** | [**RestartPolicy**](RestartPolicy.md) |  |  [optional] |
|**scaling** | [**ScalingPolicy**](ScalingPolicy.md) |  |  [optional] |
|**services** | [**List&lt;Service&gt;**](Service.md) |  |  [optional] |
|**shutdownDelay** | **Long** |  |  [optional] |
|**spreads** | [**List&lt;Spread&gt;**](Spread.md) |  |  [optional] |
|**stopAfterClientDisconnect** | **Long** |  |  [optional] |
|**tasks** | [**List&lt;Task&gt;**](Task.md) |  |  [optional] |
|**update** | [**UpdateStrategy**](UpdateStrategy.md) |  |  [optional] |
|**volumes** | [**Map&lt;String, VolumeRequest&gt;**](VolumeRequest.md) |  |  [optional] |



