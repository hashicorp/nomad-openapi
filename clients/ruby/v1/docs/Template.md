# NomadClient::Template

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **change_mode** | **String** |  | [optional] |
| **change_signal** | **String** |  | [optional] |
| **dest_path** | **String** |  | [optional] |
| **embedded_tmpl** | **String** |  | [optional] |
| **envvars** | **Boolean** |  | [optional] |
| **left_delim** | **String** |  | [optional] |
| **perms** | **String** |  | [optional] |
| **right_delim** | **String** |  | [optional] |
| **source_path** | **String** |  | [optional] |
| **splay** | **Integer** |  | [optional] |
| **vault_grace** | **Integer** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::Template.new(
  change_mode: null,
  change_signal: null,
  dest_path: null,
  embedded_tmpl: null,
  envvars: null,
  left_delim: null,
  perms: null,
  right_delim: null,
  source_path: null,
  splay: null,
  vault_grace: null
)
```

