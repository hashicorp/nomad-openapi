# NomadClient::Namespace

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **capabilities** | [**NamespaceCapabilities**](NamespaceCapabilities.md) |  | [optional] |
| **create_index** | **Integer** |  | [optional] |
| **description** | **String** |  | [optional] |
| **meta** | **Hash&lt;String, String&gt;** |  | [optional] |
| **modify_index** | **Integer** |  | [optional] |
| **name** | **String** |  | [optional] |
| **quota** | **String** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::Namespace.new(
  capabilities: null,
  create_index: null,
  description: null,
  meta: null,
  modify_index: null,
  name: null,
  quota: null
)
```

