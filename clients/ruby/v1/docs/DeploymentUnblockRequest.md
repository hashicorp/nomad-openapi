# NomadClient::DeploymentUnblockRequest

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **deployment_id** | **String** |  | [optional] |
| **namespace** | **String** |  | [optional] |
| **region** | **String** |  | [optional] |
| **secret_id** | **String** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::DeploymentUnblockRequest.new(
  deployment_id: null,
  namespace: null,
  region: null,
  secret_id: null
)
```

