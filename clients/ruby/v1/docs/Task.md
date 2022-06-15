# NomadClient::Task

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **affinities** | [**Array&lt;Affinity&gt;**](Affinity.md) |  | [optional] |
| **artifacts** | [**Array&lt;TaskArtifact&gt;**](TaskArtifact.md) |  | [optional] |
| **csi_plugin_config** | [**TaskCSIPluginConfig**](TaskCSIPluginConfig.md) |  | [optional] |
| **config** | **Hash&lt;String, Object&gt;** |  | [optional] |
| **constraints** | [**Array&lt;Constraint&gt;**](Constraint.md) |  | [optional] |
| **dispatch_payload** | [**DispatchPayloadConfig**](DispatchPayloadConfig.md) |  | [optional] |
| **driver** | **String** |  | [optional] |
| **env** | **Hash&lt;String, String&gt;** |  | [optional] |
| **kill_signal** | **String** |  | [optional] |
| **kill_timeout** | **Integer** |  | [optional] |
| **kind** | **String** |  | [optional] |
| **leader** | **Boolean** |  | [optional] |
| **lifecycle** | [**TaskLifecycle**](TaskLifecycle.md) |  | [optional] |
| **log_config** | [**LogConfig**](LogConfig.md) |  | [optional] |
| **meta** | **Hash&lt;String, String&gt;** |  | [optional] |
| **name** | **String** |  | [optional] |
| **resources** | [**Resources**](Resources.md) |  | [optional] |
| **restart_policy** | [**RestartPolicy**](RestartPolicy.md) |  | [optional] |
| **scaling_policies** | [**Array&lt;ScalingPolicy&gt;**](ScalingPolicy.md) |  | [optional] |
| **services** | [**Array&lt;Service&gt;**](Service.md) |  | [optional] |
| **shutdown_delay** | **Integer** |  | [optional] |
| **templates** | [**Array&lt;Template&gt;**](Template.md) |  | [optional] |
| **user** | **String** |  | [optional] |
| **vault** | [**Vault**](Vault.md) |  | [optional] |
| **volume_mounts** | [**Array&lt;VolumeMount&gt;**](VolumeMount.md) |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::Task.new(
  affinities: null,
  artifacts: null,
  csi_plugin_config: null,
  config: null,
  constraints: null,
  dispatch_payload: null,
  driver: null,
  env: null,
  kill_signal: null,
  kill_timeout: null,
  kind: null,
  leader: null,
  lifecycle: null,
  log_config: null,
  meta: null,
  name: null,
  resources: null,
  restart_policy: null,
  scaling_policies: null,
  services: null,
  shutdown_delay: null,
  templates: null,
  user: null,
  vault: null,
  volume_mounts: null
)
```

