# NomadClient::JobStabilityRequest

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **job_id** | **String** |  | [optional] |
| **job_version** | **Integer** |  | [optional] |
| **namespace** | **String** |  | [optional] |
| **region** | **String** |  | [optional] |
| **secret_id** | **String** |  | [optional] |
| **stable** | **Boolean** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::JobStabilityRequest.new(
  job_id: null,
  job_version: null,
  namespace: null,
  region: null,
  secret_id: null,
  stable: null
)
```

