# NomadClient::TaskGroup

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **affinities** | [**Array&lt;Affinity&gt;**](Affinity.md) |  | [optional] |
| **constraints** | [**Array&lt;Constraint&gt;**](Constraint.md) |  | [optional] |
| **consul** | [**Consul**](Consul.md) |  | [optional] |
| **count** | **Integer** |  | [optional] |
| **ephemeral_disk** | [**EphemeralDisk**](EphemeralDisk.md) |  | [optional] |
| **meta** | **Hash&lt;String, String&gt;** |  | [optional] |
| **migrate** | [**MigrateStrategy**](MigrateStrategy.md) |  | [optional] |
| **name** | **String** |  | [optional] |
| **networks** | [**Array&lt;NetworkResource&gt;**](NetworkResource.md) |  | [optional] |
| **reschedule_policy** | [**ReschedulePolicy**](ReschedulePolicy.md) |  | [optional] |
| **restart_policy** | [**RestartPolicy**](RestartPolicy.md) |  | [optional] |
| **scaling** | [**ScalingPolicy**](ScalingPolicy.md) |  | [optional] |
| **services** | [**Array&lt;Service&gt;**](Service.md) |  | [optional] |
| **shutdown_delay** | **Integer** |  | [optional] |
| **spreads** | [**Array&lt;Spread&gt;**](Spread.md) |  | [optional] |
| **stop_after_client_disconnect** | **Integer** |  | [optional] |
| **tasks** | [**Array&lt;Task&gt;**](Task.md) |  | [optional] |
| **update** | [**UpdateStrategy**](UpdateStrategy.md) |  | [optional] |
| **volumes** | [**Hash&lt;String, VolumeRequest&gt;**](VolumeRequest.md) |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::TaskGroup.new(
  affinities: null,
  constraints: null,
  consul: null,
  count: null,
  ephemeral_disk: null,
  meta: null,
  migrate: null,
  name: null,
  networks: null,
  reschedule_policy: null,
  restart_policy: null,
  scaling: null,
  services: null,
  shutdown_delay: null,
  spreads: null,
  stop_after_client_disconnect: null,
  tasks: null,
  update: null,
  volumes: null
)
```

