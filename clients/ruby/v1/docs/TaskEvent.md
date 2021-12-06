# NomadClient::TaskEvent

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **details** | **Hash&lt;String, String&gt;** |  | [optional] |
| **disk_limit** | **Integer** |  | [optional] |
| **disk_size** | **Integer** |  | [optional] |
| **display_message** | **String** |  | [optional] |
| **download_error** | **String** |  | [optional] |
| **driver_error** | **String** |  | [optional] |
| **driver_message** | **String** |  | [optional] |
| **exit_code** | **Integer** |  | [optional] |
| **failed_sibling** | **String** |  | [optional] |
| **fails_task** | **Boolean** |  | [optional] |
| **generic_source** | **String** |  | [optional] |
| **kill_error** | **String** |  | [optional] |
| **kill_reason** | **String** |  | [optional] |
| **kill_timeout** | **Integer** |  | [optional] |
| **message** | **String** |  | [optional] |
| **restart_reason** | **String** |  | [optional] |
| **setup_error** | **String** |  | [optional] |
| **signal** | **Integer** |  | [optional] |
| **start_delay** | **Integer** |  | [optional] |
| **task_signal** | **String** |  | [optional] |
| **task_signal_reason** | **String** |  | [optional] |
| **time** | **Integer** |  | [optional] |
| **type** | **String** |  | [optional] |
| **validation_error** | **String** |  | [optional] |
| **vault_error** | **String** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::TaskEvent.new(
  details: null,
  disk_limit: null,
  disk_size: null,
  display_message: null,
  download_error: null,
  driver_error: null,
  driver_message: null,
  exit_code: null,
  failed_sibling: null,
  fails_task: null,
  generic_source: null,
  kill_error: null,
  kill_reason: null,
  kill_timeout: null,
  message: null,
  restart_reason: null,
  setup_error: null,
  signal: null,
  start_delay: null,
  task_signal: null,
  task_signal_reason: null,
  time: null,
  type: null,
  validation_error: null,
  vault_error: null
)
```

