# NomadClient::DeploymentPauseRequest

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **deployment_id** | **String** |  | [optional] |
| **namespace** | **String** |  | [optional] |
| **pause** | **Boolean** |  | [optional] |
| **region** | **String** |  | [optional] |
| **secret_id** | **String** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::DeploymentPauseRequest.new(
  deployment_id: null,
  namespace: null,
  pause: null,
  region: null,
  secret_id: null
)
```

