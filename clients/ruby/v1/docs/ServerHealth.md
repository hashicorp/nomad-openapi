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

