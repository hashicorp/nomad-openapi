# NomadClient::JobRevertRequest

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **consul_token** | **String** |  | [optional] |
| **enforce_prior_version** | **Integer** |  | [optional] |
| **job_id** | **String** |  | [optional] |
| **job_version** | **Integer** |  | [optional] |
| **namespace** | **String** |  | [optional] |
| **region** | **String** |  | [optional] |
| **secret_id** | **String** |  | [optional] |
| **vault_token** | **String** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::JobRevertRequest.new(
  consul_token: null,
  enforce_prior_version: null,
  job_id: null,
  job_version: null,
  namespace: null,
  region: null,
  secret_id: null,
  vault_token: null
)
```

