# NomadClient::Template

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **change_mode** | **String** |  | [optional] |
| **change_script** | [**ChangeScript**](ChangeScript.md) |  | [optional] |
| **change_signal** | **String** |  | [optional] |
| **dest_path** | **String** |  | [optional] |
| **embedded_tmpl** | **String** |  | [optional] |
| **envvars** | **Boolean** |  | [optional] |
| **err_missing_key** | **Boolean** |  | [optional] |
| **gid** | **Integer** |  | [optional] |
| **left_delim** | **String** |  | [optional] |
| **perms** | **String** |  | [optional] |
| **right_delim** | **String** |  | [optional] |
| **source_path** | **String** |  | [optional] |
| **splay** | **Integer** |  | [optional] |
| **uid** | **Integer** |  | [optional] |
| **vault_grace** | **Integer** |  | [optional] |
| **wait** | [**WaitConfig**](WaitConfig.md) |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::Template.new(
  change_mode: null,
  change_script: null,
  change_signal: null,
  dest_path: null,
  embedded_tmpl: null,
  envvars: null,
  err_missing_key: null,
  gid: null,
  left_delim: null,
  perms: null,
  right_delim: null,
  source_path: null,
  splay: null,
  uid: null,
  vault_grace: null,
  wait: null
)
```

