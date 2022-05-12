# NomadClient::ServiceCheck

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **address_mode** | **String** |  | [optional] |
| **advertise** | **String** |  | [optional] |
| **args** | **Array&lt;String&gt;** |  | [optional] |
| **body** | **String** |  | [optional] |
| **check_restart** | [**CheckRestart**](CheckRestart.md) |  | [optional] |
| **command** | **String** |  | [optional] |
| **expose** | **Boolean** |  | [optional] |
| **failures_before_critical** | **Integer** |  | [optional] |
| **grpc_service** | **String** |  | [optional] |
| **grpc_use_tls** | **Boolean** |  | [optional] |
| **header** | **Hash&lt;String, Array&lt;String&gt;&gt;** |  | [optional] |
| **initial_status** | **String** |  | [optional] |
| **interval** | **Integer** |  | [optional] |
| **method** | **String** |  | [optional] |
| **name** | **String** |  | [optional] |
| **on_update** | **String** |  | [optional] |
| **path** | **String** |  | [optional] |
| **port_label** | **String** |  | [optional] |
| **protocol** | **String** |  | [optional] |
| **success_before_passing** | **Integer** |  | [optional] |
| **tls_skip_verify** | **Boolean** |  | [optional] |
| **task_name** | **String** |  | [optional] |
| **timeout** | **Integer** |  | [optional] |
| **type** | **String** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::ServiceCheck.new(
  address_mode: null,
  advertise: null,
  args: null,
  body: null,
  check_restart: null,
  command: null,
  expose: null,
  failures_before_critical: null,
  grpc_service: null,
  grpc_use_tls: null,
  header: null,
  initial_status: null,
  interval: null,
  method: null,
  name: null,
  on_update: null,
  path: null,
  port_label: null,
  protocol: null,
  success_before_passing: null,
  tls_skip_verify: null,
  task_name: null,
  timeout: null,
  type: null
)
```

