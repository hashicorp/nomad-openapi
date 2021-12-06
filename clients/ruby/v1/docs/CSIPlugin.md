# NomadClient::CSIPlugin

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **allocations** | [**Array&lt;AllocationListStub&gt;**](AllocationListStub.md) |  | [optional] |
| **controller_required** | **Boolean** |  | [optional] |
| **controllers** | [**Hash&lt;String, CSIInfo&gt;**](CSIInfo.md) |  | [optional] |
| **controllers_expected** | **Integer** |  | [optional] |
| **controllers_healthy** | **Integer** |  | [optional] |
| **create_index** | **Integer** |  | [optional] |
| **id** | **String** |  | [optional] |
| **modify_index** | **Integer** |  | [optional] |
| **nodes** | [**Hash&lt;String, CSIInfo&gt;**](CSIInfo.md) |  | [optional] |
| **nodes_expected** | **Integer** |  | [optional] |
| **nodes_healthy** | **Integer** |  | [optional] |
| **provider** | **String** |  | [optional] |
| **version** | **String** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::CSIPlugin.new(
  allocations: null,
  controller_required: null,
  controllers: null,
  controllers_expected: null,
  controllers_healthy: null,
  create_index: null,
  id: null,
  modify_index: null,
  nodes: null,
  nodes_expected: null,
  nodes_healthy: null,
  provider: null,
  version: null
)
```

