# NomadClient::JobValidateResponse

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **driver_config_validated** | **Boolean** |  | [optional] |
| **error** | **String** |  | [optional] |
| **validation_errors** | **Array&lt;String&gt;** |  | [optional] |
| **warnings** | **String** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::JobValidateResponse.new(
  driver_config_validated: null,
  error: null,
  validation_errors: null,
  warnings: null
)
```

