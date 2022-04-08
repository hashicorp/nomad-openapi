# NomadClient::TaskGroupSummary

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **complete** | **Integer** |  | [optional] |
| **failed** | **Integer** |  | [optional] |
| **lost** | **Integer** |  | [optional] |
| **queued** | **Integer** |  | [optional] |
| **running** | **Integer** |  | [optional] |
| **starting** | **Integer** |  | [optional] |
| **unknown** | **Integer** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::TaskGroupSummary.new(
  complete: null,
  failed: null,
  lost: null,
  queued: null,
  running: null,
  starting: null,
  unknown: null
)
```

