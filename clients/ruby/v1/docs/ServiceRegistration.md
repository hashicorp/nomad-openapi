# NomadClient::ServiceRegistration

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **address** | **String** |  | [optional] |
| **alloc_id** | **String** |  | [optional] |
| **create_index** | **Integer** |  | [optional] |
| **datacenter** | **String** |  | [optional] |
| **id** | **String** |  | [optional] |
| **job_id** | **String** |  | [optional] |
| **modify_index** | **Integer** |  | [optional] |
| **namespace** | **String** |  | [optional] |
| **node_id** | **String** |  | [optional] |
| **port** | **Integer** |  | [optional] |
| **service_name** | **String** |  | [optional] |
| **tags** | **Array&lt;String&gt;** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::ServiceRegistration.new(
  address: null,
  alloc_id: null,
  create_index: null,
  datacenter: null,
  id: null,
  job_id: null,
  modify_index: null,
  namespace: null,
  node_id: null,
  port: null,
  service_name: null,
  tags: null
)
```

