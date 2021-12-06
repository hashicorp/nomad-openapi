# NomadClient::TaskState

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **events** | [**Array&lt;TaskEvent&gt;**](TaskEvent.md) |  | [optional] |
| **failed** | **Boolean** |  | [optional] |
| **finished_at** | **Time** |  | [optional] |
| **last_restart** | **Time** |  | [optional] |
| **restarts** | **Integer** |  | [optional] |
| **started_at** | **Time** |  | [optional] |
| **state** | **String** |  | [optional] |
| **task_handle** | [**TaskHandle**](TaskHandle.md) |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::TaskState.new(
  events: null,
  failed: null,
  finished_at: null,
  last_restart: null,
  restarts: null,
  started_at: null,
  state: null,
  task_handle: null
)
```

