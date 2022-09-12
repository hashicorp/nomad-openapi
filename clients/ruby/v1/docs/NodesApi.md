# NomadClient::NodesApi

All URIs are relative to *http://127.0.0.1:4646/v1*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**get_node**](NodesApi.md#get_node) | **GET** /node/{nodeId} |  |
| [**get_node_allocations**](NodesApi.md#get_node_allocations) | **GET** /node/{nodeId}/allocations |  |
| [**get_nodes**](NodesApi.md#get_nodes) | **GET** /nodes |  |
| [**update_node_drain**](NodesApi.md#update_node_drain) | **POST** /node/{nodeId}/drain |  |
| [**update_node_eligibility**](NodesApi.md#update_node_eligibility) | **POST** /node/{nodeId}/eligibility |  |
| [**update_node_purge**](NodesApi.md#update_node_purge) | **POST** /node/{nodeId}/purge |  |


## get_node

> <Node> get_node(node_id, opts)



### Examples

```ruby
require 'time'
require 'nomad_client'
# setup authorization
NomadClient.configure do |config|
  # Configure API key authorization: X-Nomad-Token
  config.api_key['X-Nomad-Token'] = 'YOUR API KEY'
  # Uncomment the following line to set a prefix for the API key, e.g. 'Bearer' (defaults to nil)
  # config.api_key_prefix['X-Nomad-Token'] = 'Bearer'
end

api_instance = NomadClient::NodesApi.new
node_id = 'node_id_example' # String | The ID of the node.
opts = {
  region: 'region_example', # String | Filters results based on the specified region.
  namespace: 'namespace_example', # String | Filters results based on the specified namespace.
  index: 56, # Integer | If set, wait until query exceeds given index. Must be provided with WaitParam.
  wait: 'wait_example', # String | Provided with IndexParam to wait for change.
  stale: 'stale_example', # String | If present, results will include stale reads.
  prefix: 'prefix_example', # String | Constrains results to jobs that start with the defined prefix
  x_nomad_token: 'x_nomad_token_example', # String | A Nomad ACL token.
  per_page: 56, # Integer | Maximum number of results to return.
  next_token: 'next_token_example' # String | Indicates where to start paging for queries that support pagination.
}

begin
  
  result = api_instance.get_node(node_id, opts)
  p result
rescue NomadClient::ApiError => e
  puts "Error when calling NodesApi->get_node: #{e}"
end
```

#### Using the get_node_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Node>, Integer, Hash)> get_node_with_http_info(node_id, opts)

```ruby
begin
  
  data, status_code, headers = api_instance.get_node_with_http_info(node_id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Node>
rescue NomadClient::ApiError => e
  puts "Error when calling NodesApi->get_node_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **node_id** | **String** | The ID of the node. |  |
| **region** | **String** | Filters results based on the specified region. | [optional] |
| **namespace** | **String** | Filters results based on the specified namespace. | [optional] |
| **index** | **Integer** | If set, wait until query exceeds given index. Must be provided with WaitParam. | [optional] |
| **wait** | **String** | Provided with IndexParam to wait for change. | [optional] |
| **stale** | **String** | If present, results will include stale reads. | [optional] |
| **prefix** | **String** | Constrains results to jobs that start with the defined prefix | [optional] |
| **x_nomad_token** | **String** | A Nomad ACL token. | [optional] |
| **per_page** | **Integer** | Maximum number of results to return. | [optional] |
| **next_token** | **String** | Indicates where to start paging for queries that support pagination. | [optional] |

### Return type

[**Node**](Node.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## get_node_allocations

> <Array<AllocationListStub>> get_node_allocations(node_id, opts)



### Examples

```ruby
require 'time'
require 'nomad_client'
# setup authorization
NomadClient.configure do |config|
  # Configure API key authorization: X-Nomad-Token
  config.api_key['X-Nomad-Token'] = 'YOUR API KEY'
  # Uncomment the following line to set a prefix for the API key, e.g. 'Bearer' (defaults to nil)
  # config.api_key_prefix['X-Nomad-Token'] = 'Bearer'
end

api_instance = NomadClient::NodesApi.new
node_id = 'node_id_example' # String | The ID of the node.
opts = {
  region: 'region_example', # String | Filters results based on the specified region.
  namespace: 'namespace_example', # String | Filters results based on the specified namespace.
  index: 56, # Integer | If set, wait until query exceeds given index. Must be provided with WaitParam.
  wait: 'wait_example', # String | Provided with IndexParam to wait for change.
  stale: 'stale_example', # String | If present, results will include stale reads.
  prefix: 'prefix_example', # String | Constrains results to jobs that start with the defined prefix
  x_nomad_token: 'x_nomad_token_example', # String | A Nomad ACL token.
  per_page: 56, # Integer | Maximum number of results to return.
  next_token: 'next_token_example' # String | Indicates where to start paging for queries that support pagination.
}

begin
  
  result = api_instance.get_node_allocations(node_id, opts)
  p result
rescue NomadClient::ApiError => e
  puts "Error when calling NodesApi->get_node_allocations: #{e}"
end
```

#### Using the get_node_allocations_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Array<AllocationListStub>>, Integer, Hash)> get_node_allocations_with_http_info(node_id, opts)

```ruby
begin
  
  data, status_code, headers = api_instance.get_node_allocations_with_http_info(node_id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Array<AllocationListStub>>
rescue NomadClient::ApiError => e
  puts "Error when calling NodesApi->get_node_allocations_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **node_id** | **String** | The ID of the node. |  |
| **region** | **String** | Filters results based on the specified region. | [optional] |
| **namespace** | **String** | Filters results based on the specified namespace. | [optional] |
| **index** | **Integer** | If set, wait until query exceeds given index. Must be provided with WaitParam. | [optional] |
| **wait** | **String** | Provided with IndexParam to wait for change. | [optional] |
| **stale** | **String** | If present, results will include stale reads. | [optional] |
| **prefix** | **String** | Constrains results to jobs that start with the defined prefix | [optional] |
| **x_nomad_token** | **String** | A Nomad ACL token. | [optional] |
| **per_page** | **Integer** | Maximum number of results to return. | [optional] |
| **next_token** | **String** | Indicates where to start paging for queries that support pagination. | [optional] |

### Return type

[**Array&lt;AllocationListStub&gt;**](AllocationListStub.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## get_nodes

> <Array<NodeListStub>> get_nodes(opts)



### Examples

```ruby
require 'time'
require 'nomad_client'
# setup authorization
NomadClient.configure do |config|
  # Configure API key authorization: X-Nomad-Token
  config.api_key['X-Nomad-Token'] = 'YOUR API KEY'
  # Uncomment the following line to set a prefix for the API key, e.g. 'Bearer' (defaults to nil)
  # config.api_key_prefix['X-Nomad-Token'] = 'Bearer'
end

api_instance = NomadClient::NodesApi.new
opts = {
  region: 'region_example', # String | Filters results based on the specified region.
  namespace: 'namespace_example', # String | Filters results based on the specified namespace.
  index: 56, # Integer | If set, wait until query exceeds given index. Must be provided with WaitParam.
  wait: 'wait_example', # String | Provided with IndexParam to wait for change.
  stale: 'stale_example', # String | If present, results will include stale reads.
  prefix: 'prefix_example', # String | Constrains results to jobs that start with the defined prefix
  x_nomad_token: 'x_nomad_token_example', # String | A Nomad ACL token.
  per_page: 56, # Integer | Maximum number of results to return.
  next_token: 'next_token_example', # String | Indicates where to start paging for queries that support pagination.
  resources: true # Boolean | Whether or not to include the NodeResources and ReservedResources fields in the response.
}

begin
  
  result = api_instance.get_nodes(opts)
  p result
rescue NomadClient::ApiError => e
  puts "Error when calling NodesApi->get_nodes: #{e}"
end
```

#### Using the get_nodes_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Array<NodeListStub>>, Integer, Hash)> get_nodes_with_http_info(opts)

```ruby
begin
  
  data, status_code, headers = api_instance.get_nodes_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Array<NodeListStub>>
rescue NomadClient::ApiError => e
  puts "Error when calling NodesApi->get_nodes_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **region** | **String** | Filters results based on the specified region. | [optional] |
| **namespace** | **String** | Filters results based on the specified namespace. | [optional] |
| **index** | **Integer** | If set, wait until query exceeds given index. Must be provided with WaitParam. | [optional] |
| **wait** | **String** | Provided with IndexParam to wait for change. | [optional] |
| **stale** | **String** | If present, results will include stale reads. | [optional] |
| **prefix** | **String** | Constrains results to jobs that start with the defined prefix | [optional] |
| **x_nomad_token** | **String** | A Nomad ACL token. | [optional] |
| **per_page** | **Integer** | Maximum number of results to return. | [optional] |
| **next_token** | **String** | Indicates where to start paging for queries that support pagination. | [optional] |
| **resources** | **Boolean** | Whether or not to include the NodeResources and ReservedResources fields in the response. | [optional] |

### Return type

[**Array&lt;NodeListStub&gt;**](NodeListStub.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## update_node_drain

> <NodeDrainUpdateResponse> update_node_drain(node_id, node_update_drain_request, opts)



### Examples

```ruby
require 'time'
require 'nomad_client'
# setup authorization
NomadClient.configure do |config|
  # Configure API key authorization: X-Nomad-Token
  config.api_key['X-Nomad-Token'] = 'YOUR API KEY'
  # Uncomment the following line to set a prefix for the API key, e.g. 'Bearer' (defaults to nil)
  # config.api_key_prefix['X-Nomad-Token'] = 'Bearer'
end

api_instance = NomadClient::NodesApi.new
node_id = 'node_id_example' # String | The ID of the node.
node_update_drain_request = NomadClient::NodeUpdateDrainRequest.new # NodeUpdateDrainRequest | 
opts = {
  region: 'region_example', # String | Filters results based on the specified region.
  namespace: 'namespace_example', # String | Filters results based on the specified namespace.
  index: 56, # Integer | If set, wait until query exceeds given index. Must be provided with WaitParam.
  wait: 'wait_example', # String | Provided with IndexParam to wait for change.
  stale: 'stale_example', # String | If present, results will include stale reads.
  prefix: 'prefix_example', # String | Constrains results to jobs that start with the defined prefix
  x_nomad_token: 'x_nomad_token_example', # String | A Nomad ACL token.
  per_page: 56, # Integer | Maximum number of results to return.
  next_token: 'next_token_example' # String | Indicates where to start paging for queries that support pagination.
}

begin
  
  result = api_instance.update_node_drain(node_id, node_update_drain_request, opts)
  p result
rescue NomadClient::ApiError => e
  puts "Error when calling NodesApi->update_node_drain: #{e}"
end
```

#### Using the update_node_drain_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<NodeDrainUpdateResponse>, Integer, Hash)> update_node_drain_with_http_info(node_id, node_update_drain_request, opts)

```ruby
begin
  
  data, status_code, headers = api_instance.update_node_drain_with_http_info(node_id, node_update_drain_request, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <NodeDrainUpdateResponse>
rescue NomadClient::ApiError => e
  puts "Error when calling NodesApi->update_node_drain_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **node_id** | **String** | The ID of the node. |  |
| **node_update_drain_request** | [**NodeUpdateDrainRequest**](NodeUpdateDrainRequest.md) |  |  |
| **region** | **String** | Filters results based on the specified region. | [optional] |
| **namespace** | **String** | Filters results based on the specified namespace. | [optional] |
| **index** | **Integer** | If set, wait until query exceeds given index. Must be provided with WaitParam. | [optional] |
| **wait** | **String** | Provided with IndexParam to wait for change. | [optional] |
| **stale** | **String** | If present, results will include stale reads. | [optional] |
| **prefix** | **String** | Constrains results to jobs that start with the defined prefix | [optional] |
| **x_nomad_token** | **String** | A Nomad ACL token. | [optional] |
| **per_page** | **Integer** | Maximum number of results to return. | [optional] |
| **next_token** | **String** | Indicates where to start paging for queries that support pagination. | [optional] |

### Return type

[**NodeDrainUpdateResponse**](NodeDrainUpdateResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## update_node_eligibility

> <NodeEligibilityUpdateResponse> update_node_eligibility(node_id, node_update_eligibility_request, opts)



### Examples

```ruby
require 'time'
require 'nomad_client'
# setup authorization
NomadClient.configure do |config|
  # Configure API key authorization: X-Nomad-Token
  config.api_key['X-Nomad-Token'] = 'YOUR API KEY'
  # Uncomment the following line to set a prefix for the API key, e.g. 'Bearer' (defaults to nil)
  # config.api_key_prefix['X-Nomad-Token'] = 'Bearer'
end

api_instance = NomadClient::NodesApi.new
node_id = 'node_id_example' # String | The ID of the node.
node_update_eligibility_request = NomadClient::NodeUpdateEligibilityRequest.new # NodeUpdateEligibilityRequest | 
opts = {
  region: 'region_example', # String | Filters results based on the specified region.
  namespace: 'namespace_example', # String | Filters results based on the specified namespace.
  index: 56, # Integer | If set, wait until query exceeds given index. Must be provided with WaitParam.
  wait: 'wait_example', # String | Provided with IndexParam to wait for change.
  stale: 'stale_example', # String | If present, results will include stale reads.
  prefix: 'prefix_example', # String | Constrains results to jobs that start with the defined prefix
  x_nomad_token: 'x_nomad_token_example', # String | A Nomad ACL token.
  per_page: 56, # Integer | Maximum number of results to return.
  next_token: 'next_token_example' # String | Indicates where to start paging for queries that support pagination.
}

begin
  
  result = api_instance.update_node_eligibility(node_id, node_update_eligibility_request, opts)
  p result
rescue NomadClient::ApiError => e
  puts "Error when calling NodesApi->update_node_eligibility: #{e}"
end
```

#### Using the update_node_eligibility_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<NodeEligibilityUpdateResponse>, Integer, Hash)> update_node_eligibility_with_http_info(node_id, node_update_eligibility_request, opts)

```ruby
begin
  
  data, status_code, headers = api_instance.update_node_eligibility_with_http_info(node_id, node_update_eligibility_request, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <NodeEligibilityUpdateResponse>
rescue NomadClient::ApiError => e
  puts "Error when calling NodesApi->update_node_eligibility_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **node_id** | **String** | The ID of the node. |  |
| **node_update_eligibility_request** | [**NodeUpdateEligibilityRequest**](NodeUpdateEligibilityRequest.md) |  |  |
| **region** | **String** | Filters results based on the specified region. | [optional] |
| **namespace** | **String** | Filters results based on the specified namespace. | [optional] |
| **index** | **Integer** | If set, wait until query exceeds given index. Must be provided with WaitParam. | [optional] |
| **wait** | **String** | Provided with IndexParam to wait for change. | [optional] |
| **stale** | **String** | If present, results will include stale reads. | [optional] |
| **prefix** | **String** | Constrains results to jobs that start with the defined prefix | [optional] |
| **x_nomad_token** | **String** | A Nomad ACL token. | [optional] |
| **per_page** | **Integer** | Maximum number of results to return. | [optional] |
| **next_token** | **String** | Indicates where to start paging for queries that support pagination. | [optional] |

### Return type

[**NodeEligibilityUpdateResponse**](NodeEligibilityUpdateResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## update_node_purge

> <NodePurgeResponse> update_node_purge(node_id, opts)



### Examples

```ruby
require 'time'
require 'nomad_client'
# setup authorization
NomadClient.configure do |config|
  # Configure API key authorization: X-Nomad-Token
  config.api_key['X-Nomad-Token'] = 'YOUR API KEY'
  # Uncomment the following line to set a prefix for the API key, e.g. 'Bearer' (defaults to nil)
  # config.api_key_prefix['X-Nomad-Token'] = 'Bearer'
end

api_instance = NomadClient::NodesApi.new
node_id = 'node_id_example' # String | The ID of the node.
opts = {
  region: 'region_example', # String | Filters results based on the specified region.
  namespace: 'namespace_example', # String | Filters results based on the specified namespace.
  index: 56, # Integer | If set, wait until query exceeds given index. Must be provided with WaitParam.
  wait: 'wait_example', # String | Provided with IndexParam to wait for change.
  stale: 'stale_example', # String | If present, results will include stale reads.
  prefix: 'prefix_example', # String | Constrains results to jobs that start with the defined prefix
  x_nomad_token: 'x_nomad_token_example', # String | A Nomad ACL token.
  per_page: 56, # Integer | Maximum number of results to return.
  next_token: 'next_token_example' # String | Indicates where to start paging for queries that support pagination.
}

begin
  
  result = api_instance.update_node_purge(node_id, opts)
  p result
rescue NomadClient::ApiError => e
  puts "Error when calling NodesApi->update_node_purge: #{e}"
end
```

#### Using the update_node_purge_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<NodePurgeResponse>, Integer, Hash)> update_node_purge_with_http_info(node_id, opts)

```ruby
begin
  
  data, status_code, headers = api_instance.update_node_purge_with_http_info(node_id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <NodePurgeResponse>
rescue NomadClient::ApiError => e
  puts "Error when calling NodesApi->update_node_purge_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **node_id** | **String** | The ID of the node. |  |
| **region** | **String** | Filters results based on the specified region. | [optional] |
| **namespace** | **String** | Filters results based on the specified namespace. | [optional] |
| **index** | **Integer** | If set, wait until query exceeds given index. Must be provided with WaitParam. | [optional] |
| **wait** | **String** | Provided with IndexParam to wait for change. | [optional] |
| **stale** | **String** | If present, results will include stale reads. | [optional] |
| **prefix** | **String** | Constrains results to jobs that start with the defined prefix | [optional] |
| **x_nomad_token** | **String** | A Nomad ACL token. | [optional] |
| **per_page** | **Integer** | Maximum number of results to return. | [optional] |
| **next_token** | **String** | Indicates where to start paging for queries that support pagination. | [optional] |

### Return type

[**NodePurgeResponse**](NodePurgeResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

