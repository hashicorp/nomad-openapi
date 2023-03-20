# NomadClient::JobDispatchRequest

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id_prefix_template** | **String** |  | [optional] |
| **job_id** | **String** |  | [optional] |
| **meta** | **Hash&lt;String, String&gt;** |  | [optional] |
| **payload** | **String** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::JobDispatchRequest.new(
  id_prefix_template: null,
  job_id: null,
  meta: null,
  payload: null
)
```

