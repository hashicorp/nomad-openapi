# NomadClient::JobDispatchRequest

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **job_id** | **String** |  | [optional] |
| **meta** | **Hash&lt;String, String&gt;** |  | [optional] |
| **payload** | **String** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::JobDispatchRequest.new(
  job_id: null,
  meta: null,
  payload: null
)
```

