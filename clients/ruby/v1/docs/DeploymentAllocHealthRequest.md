# NomadClient::DeploymentAllocHealthRequest

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **deployment_id** | **String** |  | [optional] |
| **healthy_allocation_ids** | **Array&lt;String&gt;** |  | [optional] |
| **namespace** | **String** |  | [optional] |
| **region** | **String** |  | [optional] |
| **secret_id** | **String** |  | [optional] |
| **unhealthy_allocation_ids** | **Array&lt;String&gt;** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::DeploymentAllocHealthRequest.new(
  deployment_id: null,
  healthy_allocation_ids: null,
  namespace: null,
  region: null,
  secret_id: null,
  unhealthy_allocation_ids: null
)
```

