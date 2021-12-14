# NomadClient::RaftServer

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **address** | **String** |  | [optional] |
| **id** | **String** |  | [optional] |
| **leader** | **Boolean** |  | [optional] |
| **node** | **String** |  | [optional] |
| **raft_protocol** | **String** |  | [optional] |
| **voter** | **Boolean** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::RaftServer.new(
  address: null,
  id: null,
  leader: null,
  node: null,
  raft_protocol: null,
  voter: null
)
```

