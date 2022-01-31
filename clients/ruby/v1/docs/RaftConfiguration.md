# NomadClient::RaftConfiguration

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **index** | **Integer** |  | [optional] |
| **servers** | [**Array&lt;RaftServer&gt;**](RaftServer.md) |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::RaftConfiguration.new(
  index: null,
  servers: null
)
```

