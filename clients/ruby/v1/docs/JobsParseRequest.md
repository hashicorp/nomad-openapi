# NomadClient::JobsParseRequest

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **canonicalize** | **Boolean** |  | [optional] |
| **job_hcl** | **String** |  | [optional] |
| **hclv1** | **Boolean** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::JobsParseRequest.new(
  canonicalize: null,
  job_hcl: null,
  hclv1: null
)
```

