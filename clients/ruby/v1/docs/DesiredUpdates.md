# NomadClient::DesiredUpdates

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **canary** | **Integer** |  | [optional] |
| **destructive_update** | **Integer** |  | [optional] |
| **ignore** | **Integer** |  | [optional] |
| **in_place_update** | **Integer** |  | [optional] |
| **migrate** | **Integer** |  | [optional] |
| **place** | **Integer** |  | [optional] |
| **preemptions** | **Integer** |  | [optional] |
| **stop** | **Integer** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::DesiredUpdates.new(
  canary: null,
  destructive_update: null,
  ignore: null,
  in_place_update: null,
  migrate: null,
  place: null,
  preemptions: null,
  stop: null
)
```

