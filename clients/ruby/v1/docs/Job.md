# NomadClient::Job

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **affinities** | [**Array&lt;Affinity&gt;**](Affinity.md) |  | [optional] |
| **all_at_once** | **Boolean** |  | [optional] |
| **constraints** | [**Array&lt;Constraint&gt;**](Constraint.md) |  | [optional] |
| **consul_namespace** | **String** |  | [optional] |
| **consul_token** | **String** |  | [optional] |
| **create_index** | **Integer** |  | [optional] |
| **datacenters** | **Array&lt;String&gt;** |  | [optional] |
| **dispatched** | **Boolean** |  | [optional] |
| **id** | **String** |  | [optional] |
| **job_modify_index** | **Integer** |  | [optional] |
| **meta** | **Hash&lt;String, String&gt;** |  | [optional] |
| **migrate** | [**MigrateStrategy**](MigrateStrategy.md) |  | [optional] |
| **modify_index** | **Integer** |  | [optional] |
| **multiregion** | [**Multiregion**](Multiregion.md) |  | [optional] |
| **name** | **String** |  | [optional] |
| **namespace** | **String** |  | [optional] |
| **nomad_token_id** | **String** |  | [optional] |
| **parameterized_job** | [**ParameterizedJobConfig**](ParameterizedJobConfig.md) |  | [optional] |
| **parent_id** | **String** |  | [optional] |
| **payload** | **String** |  | [optional] |
| **periodic** | [**PeriodicConfig**](PeriodicConfig.md) |  | [optional] |
| **priority** | **Integer** |  | [optional] |
| **region** | **String** |  | [optional] |
| **reschedule** | [**ReschedulePolicy**](ReschedulePolicy.md) |  | [optional] |
| **spreads** | [**Array&lt;Spread&gt;**](Spread.md) |  | [optional] |
| **stable** | **Boolean** |  | [optional] |
| **status** | **String** |  | [optional] |
| **status_description** | **String** |  | [optional] |
| **stop** | **Boolean** |  | [optional] |
| **submit_time** | **Integer** |  | [optional] |
| **task_groups** | [**Array&lt;TaskGroup&gt;**](TaskGroup.md) |  | [optional] |
| **type** | **String** |  | [optional] |
| **update** | [**UpdateStrategy**](UpdateStrategy.md) |  | [optional] |
| **vault_namespace** | **String** |  | [optional] |
| **vault_token** | **String** |  | [optional] |
| **version** | **Integer** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::Job.new(
  affinities: null,
  all_at_once: null,
  constraints: null,
  consul_namespace: null,
  consul_token: null,
  create_index: null,
  datacenters: null,
  dispatched: null,
  id: null,
  job_modify_index: null,
  meta: null,
  migrate: null,
  modify_index: null,
  multiregion: null,
  name: null,
  namespace: null,
  nomad_token_id: null,
  parameterized_job: null,
  parent_id: null,
  payload: null,
  periodic: null,
  priority: null,
  region: null,
  reschedule: null,
  spreads: null,
  stable: null,
  status: null,
  status_description: null,
  stop: null,
  submit_time: null,
  task_groups: null,
  type: null,
  update: null,
  vault_namespace: null,
  vault_token: null,
  version: null
)
```

