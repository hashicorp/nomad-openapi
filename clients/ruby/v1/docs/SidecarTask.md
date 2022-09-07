# NomadClient::SidecarTask

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **config** | **Hash&lt;String, Object&gt;** |  | [optional] |
| **driver** | **String** |  | [optional] |
| **env** | **Hash&lt;String, String&gt;** |  | [optional] |
| **kill_signal** | **String** |  | [optional] |
| **kill_timeout** | **Integer** |  | [optional] |
| **log_config** | [**LogConfig**](LogConfig.md) |  | [optional] |
| **meta** | **Hash&lt;String, String&gt;** |  | [optional] |
| **name** | **String** |  | [optional] |
| **resources** | [**Resources**](Resources.md) |  | [optional] |
| **shutdown_delay** | **Integer** |  | [optional] |
| **user** | **String** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::SidecarTask.new(
  config: null,
  driver: null,
  env: null,
  kill_signal: null,
  kill_timeout: null,
  log_config: null,
  meta: null,
  name: null,
  resources: null,
  shutdown_delay: null,
  user: null
)
```

