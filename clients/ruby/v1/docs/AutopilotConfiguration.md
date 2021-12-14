# NomadClient::AutopilotConfiguration

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **cleanup_dead_servers** | **Boolean** |  | [optional] |
| **create_index** | **Integer** |  | [optional] |
| **disable_upgrade_migration** | **Boolean** |  | [optional] |
| **enable_custom_upgrades** | **Boolean** |  | [optional] |
| **enable_redundancy_zones** | **Boolean** |  | [optional] |
| **last_contact_threshold** | **Integer** |  | [optional] |
| **max_trailing_logs** | **Integer** |  | [optional] |
| **min_quorum** | **Integer** |  | [optional] |
| **modify_index** | **Integer** |  | [optional] |
| **server_stabilization_time** | **Integer** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::AutopilotConfiguration.new(
  cleanup_dead_servers: null,
  create_index: null,
  disable_upgrade_migration: null,
  enable_custom_upgrades: null,
  enable_redundancy_zones: null,
  last_contact_threshold: null,
  max_trailing_logs: null,
  min_quorum: null,
  modify_index: null,
  server_stabilization_time: null
)
```

