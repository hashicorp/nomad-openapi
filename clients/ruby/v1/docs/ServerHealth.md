<<<<<<< HEAD
# NomadClient::ServerHealth

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **address** | **String** |  | [optional] |
| **healthy** | **Boolean** |  | [optional] |
| **id** | **String** |  | [optional] |
| **last_contact** | **Integer** |  | [optional] |
| **last_index** | **Integer** |  | [optional] |
| **last_term** | **Integer** |  | [optional] |
| **leader** | **Boolean** |  | [optional] |
| **name** | **String** |  | [optional] |
| **serf_status** | **String** |  | [optional] |
| **stable_since** | **Time** |  | [optional] |
| **version** | **String** |  | [optional] |
| **voter** | **Boolean** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::ServerHealth.new(
  address: null,
  healthy: null,
  id: null,
  last_contact: null,
  last_index: null,
  last_term: null,
  leader: null,
  name: null,
  serf_status: null,
  stable_since: null,
  version: null,
  voter: null
)
```
=======
# ServerHealth


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**address** | **str** |  | [optional] 
**healthy** | **bool** |  | [optional] 
**id** | **str** |  | [optional] 
**last_contact** | **int** |  | [optional] 
**last_index** | **int** |  | [optional] 
**last_term** | **int** |  | [optional] 
**leader** | **bool** |  | [optional] 
**name** | **str** |  | [optional] 
**serf_status** | **str** |  | [optional] 
**stable_since** | **datetime** |  | [optional] 
**version** | **str** |  | [optional] 
**voter** | **bool** |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

>>>>>>> ec8f030ff3631498657522324f2a4ddaece26cc6

