<<<<<<< HEAD
# NomadClient::SchedulerConfiguration

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **create_index** | **Integer** |  | [optional] |
| **memory_oversubscription_enabled** | **Boolean** |  | [optional] |
| **modify_index** | **Integer** |  | [optional] |
| **preemption_config** | [**PreemptionConfig**](PreemptionConfig.md) |  | [optional] |
| **scheduler_algorithm** | **String** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::SchedulerConfiguration.new(
  create_index: null,
  memory_oversubscription_enabled: null,
  modify_index: null,
  preemption_config: null,
  scheduler_algorithm: null
)
```
=======
# SchedulerConfiguration


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**create_index** | **int** |  | [optional] 
**memory_oversubscription_enabled** | **bool** |  | [optional] 
**modify_index** | **int** |  | [optional] 
**preemption_config** | [**PreemptionConfig**](PreemptionConfig.md) |  | [optional] 
**scheduler_algorithm** | **str** |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

>>>>>>> ec8f030ff3631498657522324f2a4ddaece26cc6

