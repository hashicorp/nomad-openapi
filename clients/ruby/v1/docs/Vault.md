# NomadClient::Vault

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **change_mode** | **String** |  | [optional] |
| **change_signal** | **String** |  | [optional] |
| **env** | **Boolean** |  | [optional] |
| **namespace** | **String** |  | [optional] |
| **policies** | **Array&lt;String&gt;** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::Vault.new(
  change_mode: null,
  change_signal: null,
  env: null,
  namespace: null,
  policies: null
)
```

