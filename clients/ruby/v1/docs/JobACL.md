# NomadClient::JobACL

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **group** | **String** |  | [optional] |
| **job_id** | **String** |  | [optional] |
| **namespace** | **String** |  | [optional] |
| **task** | **String** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::JobACL.new(
  group: null,
  job_id: null,
  namespace: null,
  task: null
)
```

