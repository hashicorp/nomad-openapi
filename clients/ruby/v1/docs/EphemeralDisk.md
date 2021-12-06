# NomadClient::EphemeralDisk

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **migrate** | **Boolean** |  | [optional] |
| **size_mb** | **Integer** |  | [optional] |
| **sticky** | **Boolean** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::EphemeralDisk.new(
  migrate: null,
  size_mb: null,
  sticky: null
)
```

