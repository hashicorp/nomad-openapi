# NomadClient::Service

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **address** | **String** |  | [optional] |
| **address_mode** | **String** |  | [optional] |
| **canary_meta** | **Hash&lt;String, String&gt;** |  | [optional] |
| **canary_tags** | **Array&lt;String&gt;** |  | [optional] |
| **check_restart** | [**CheckRestart**](CheckRestart.md) |  | [optional] |
| **checks** | [**Array&lt;ServiceCheck&gt;**](ServiceCheck.md) |  | [optional] |
| **connect** | [**ConsulConnect**](ConsulConnect.md) |  | [optional] |
| **enable_tag_override** | **Boolean** |  | [optional] |
| **meta** | **Hash&lt;String, String&gt;** |  | [optional] |
| **name** | **String** |  | [optional] |
| **on_update** | **String** |  | [optional] |
| **port_label** | **String** |  | [optional] |
| **provider** | **String** |  | [optional] |
| **tags** | **Array&lt;String&gt;** |  | [optional] |
| **task_name** | **String** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::Service.new(
  address: null,
  address_mode: null,
  canary_meta: null,
  canary_tags: null,
  check_restart: null,
  checks: null,
  connect: null,
  enable_tag_override: null,
  meta: null,
  name: null,
  on_update: null,
  port_label: null,
  provider: null,
  tags: null,
  task_name: null
)
```

