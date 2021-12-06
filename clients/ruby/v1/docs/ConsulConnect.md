# NomadClient::ConsulConnect

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **gateway** | [**ConsulGateway**](ConsulGateway.md) |  | [optional] |
| **native** | **Boolean** |  | [optional] |
| **sidecar_service** | [**ConsulSidecarService**](ConsulSidecarService.md) |  | [optional] |
| **sidecar_task** | [**SidecarTask**](SidecarTask.md) |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::ConsulConnect.new(
  gateway: null,
  native: null,
  sidecar_service: null,
  sidecar_task: null
)
```

