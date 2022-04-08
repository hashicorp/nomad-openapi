# NomadClient::NamespaceCapabilities

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **disabled_task_drivers** | **Array&lt;String&gt;** |  | [optional] |
| **enabled_task_drivers** | **Array&lt;String&gt;** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::NamespaceCapabilities.new(
  disabled_task_drivers: null,
  enabled_task_drivers: null
)
```

