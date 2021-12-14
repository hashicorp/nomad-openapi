<<<<<<< HEAD
# NomadClient::SchedulerConfigurationResponse

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **known_leader** | **Boolean** |  | [optional] |
| **last_contact** | **Integer** |  | [optional] |
| **last_index** | **Integer** |  | [optional] |
| **request_time** | **Integer** |  | [optional] |
| **scheduler_config** | [**SchedulerConfiguration**](SchedulerConfiguration.md) |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::SchedulerConfigurationResponse.new(
  known_leader: null,
  last_contact: null,
  last_index: null,
  request_time: null,
  scheduler_config: null
)
```
=======
# SchedulerConfigurationResponse


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**known_leader** | **bool** |  | [optional] 
**last_contact** | **int** |  | [optional] 
**last_index** | **int** |  | [optional] 
**request_time** | **int** |  | [optional] 
**scheduler_config** | [**SchedulerConfiguration**](SchedulerConfiguration.md) |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

>>>>>>> ec8f030ff3631498657522324f2a4ddaece26cc6

