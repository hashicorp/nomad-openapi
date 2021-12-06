# NomadClient::CSIPluginListStub

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **controller_required** | **Boolean** |  | [optional] |
| **controllers_expected** | **Integer** |  | [optional] |
| **controllers_healthy** | **Integer** |  | [optional] |
| **create_index** | **Integer** |  | [optional] |
| **id** | **String** |  | [optional] |
| **modify_index** | **Integer** |  | [optional] |
| **nodes_expected** | **Integer** |  | [optional] |
| **nodes_healthy** | **Integer** |  | [optional] |
| **provider** | **String** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::CSIPluginListStub.new(
  controller_required: null,
  controllers_expected: null,
  controllers_healthy: null,
  create_index: null,
  id: null,
  modify_index: null,
  nodes_expected: null,
  nodes_healthy: null,
  provider: null
)
```

