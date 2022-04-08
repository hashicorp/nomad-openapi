# NomadClient::CSITopologyRequest

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **preferred** | [**Array&lt;CSITopology&gt;**](CSITopology.md) |  | [optional] |
| **required** | [**Array&lt;CSITopology&gt;**](CSITopology.md) |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::CSITopologyRequest.new(
  preferred: null,
  required: null
)
```

