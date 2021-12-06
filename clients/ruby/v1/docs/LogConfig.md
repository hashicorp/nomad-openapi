# NomadClient::LogConfig

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **max_file_size_mb** | **Integer** |  | [optional] |
| **max_files** | **Integer** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::LogConfig.new(
  max_file_size_mb: null,
  max_files: null
)
```

