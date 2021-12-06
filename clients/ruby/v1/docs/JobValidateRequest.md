# NomadClient::JobValidateRequest

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **job** | [**Job**](Job.md) |  | [optional] |
| **namespace** | **String** |  | [optional] |
| **region** | **String** |  | [optional] |
| **secret_id** | **String** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::JobValidateRequest.new(
  job: null,
  namespace: null,
  region: null,
  secret_id: null
)
```

