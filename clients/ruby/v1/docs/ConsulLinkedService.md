# NomadClient::ConsulLinkedService

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **ca_file** | **String** |  | [optional] |
| **cert_file** | **String** |  | [optional] |
| **key_file** | **String** |  | [optional] |
| **name** | **String** |  | [optional] |
| **sni** | **String** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::ConsulLinkedService.new(
  ca_file: null,
  cert_file: null,
  key_file: null,
  name: null,
  sni: null
)
```

