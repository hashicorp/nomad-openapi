# NomadClient::DeploymentPromoteRequest

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **all** | **Boolean** |  | [optional] |
| **deployment_id** | **String** |  | [optional] |
| **groups** | **Array&lt;String&gt;** |  | [optional] |
| **namespace** | **String** |  | [optional] |
| **region** | **String** |  | [optional] |
| **secret_id** | **String** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::DeploymentPromoteRequest.new(
  all: null,
  deployment_id: null,
  groups: null,
  namespace: null,
  region: null,
  secret_id: null
)
```

