# NomadClient::ChangeScript

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **args** | **Array&lt;String&gt;** |  | [optional] |
| **command** | **String** |  | [optional] |
| **fail_on_error** | **Boolean** |  | [optional] |
| **timeout** | **Integer** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::ChangeScript.new(
  args: null,
  command: null,
  fail_on_error: null,
  timeout: null
)
```

