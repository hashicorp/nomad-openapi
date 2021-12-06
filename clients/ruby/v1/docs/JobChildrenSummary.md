# NomadClient::JobChildrenSummary

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **dead** | **Integer** |  | [optional] |
| **pending** | **Integer** |  | [optional] |
| **running** | **Integer** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::JobChildrenSummary.new(
  dead: null,
  pending: null,
  running: null
)
```

