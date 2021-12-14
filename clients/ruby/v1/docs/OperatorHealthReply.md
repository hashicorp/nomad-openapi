<<<<<<< HEAD
# NomadClient::OperatorHealthReply

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **failure_tolerance** | **Integer** |  | [optional] |
| **healthy** | **Boolean** |  | [optional] |
| **servers** | [**Array&lt;ServerHealth&gt;**](ServerHealth.md) |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::OperatorHealthReply.new(
  failure_tolerance: null,
  healthy: null,
  servers: null
)
```
=======
# OperatorHealthReply


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**failure_tolerance** | **int** |  | [optional] 
**healthy** | **bool** |  | [optional] 
**servers** | [**[ServerHealth]**](ServerHealth.md) |  | [optional] 
**any string name** | **bool, date, datetime, dict, float, int, list, str, none_type** | any string name can be used but the value must be the correct type | [optional]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

>>>>>>> ec8f030ff3631498657522324f2a4ddaece26cc6

