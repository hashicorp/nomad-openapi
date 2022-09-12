# NomadClient::AllocationsApi

All URIs are relative to *http://127.0.0.1:4646/v1*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**get_allocation**](AllocationsApi.md#get_allocation) | **GET** /allocation/{allocID} |  |
| [**get_allocation_services**](AllocationsApi.md#get_allocation_services) | **GET** /allocation/{allocID}/services |  |
| [**get_allocations**](AllocationsApi.md#get_allocations) | **GET** /allocations |  |
| [**post_allocation_stop**](AllocationsApi.md#post_allocation_stop) | **POST** /allocation/{allocID}/stop |  |


## get_allocation

> <Allocation> get_allocation(alloc_id, opts)



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

api_instance = NomadClient::AllocationsApi.new
alloc_id = 'alloc_id_example' # String | Allocation ID.
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
  
  result = api_instance.get_allocation(alloc_id, opts)
  p result
rescue NomadClient::ApiError => e
  puts "Error when calling AllocationsApi->get_allocation: #{e}"
end
```

#### Using the get_allocation_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Allocation>, Integer, Hash)> get_allocation_with_http_info(alloc_id, opts)

```ruby
begin
  
  data, status_code, headers = api_instance.get_allocation_with_http_info(alloc_id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Allocation>
rescue NomadClient::ApiError => e
  puts "Error when calling AllocationsApi->get_allocation_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **alloc_id** | **String** | Allocation ID. |  |
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

[**Allocation**](Allocation.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## get_allocation_services

> <Array<ServiceRegistration>> get_allocation_services(alloc_id, opts)



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

api_instance = NomadClient::AllocationsApi.new
alloc_id = 'alloc_id_example' # String | Allocation ID.
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
  
  result = api_instance.get_allocation_services(alloc_id, opts)
  p result
rescue NomadClient::ApiError => e
  puts "Error when calling AllocationsApi->get_allocation_services: #{e}"
end
```

#### Using the get_allocation_services_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Array<ServiceRegistration>>, Integer, Hash)> get_allocation_services_with_http_info(alloc_id, opts)

```ruby
begin
  
  data, status_code, headers = api_instance.get_allocation_services_with_http_info(alloc_id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Array<ServiceRegistration>>
rescue NomadClient::ApiError => e
  puts "Error when calling AllocationsApi->get_allocation_services_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **alloc_id** | **String** | Allocation ID. |  |
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

[**Array&lt;ServiceRegistration&gt;**](ServiceRegistration.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## get_allocations

> <Array<AllocationListStub>> get_allocations(opts)



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

api_instance = NomadClient::AllocationsApi.new
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
  resources: true, # Boolean | Flag indicating whether to include resources in response.
  task_states: true # Boolean | Flag indicating whether to include task states in response.
}

begin
  
  result = api_instance.get_allocations(opts)
  p result
rescue NomadClient::ApiError => e
  puts "Error when calling AllocationsApi->get_allocations: #{e}"
end
```

#### Using the get_allocations_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Array<AllocationListStub>>, Integer, Hash)> get_allocations_with_http_info(opts)

```ruby
begin
  
  data, status_code, headers = api_instance.get_allocations_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Array<AllocationListStub>>
rescue NomadClient::ApiError => e
  puts "Error when calling AllocationsApi->get_allocations_with_http_info: #{e}"
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
| **resources** | **Boolean** | Flag indicating whether to include resources in response. | [optional] |
| **task_states** | **Boolean** | Flag indicating whether to include task states in response. | [optional] |

### Return type

[**Array&lt;AllocationListStub&gt;**](AllocationListStub.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## post_allocation_stop

> <AllocStopResponse> post_allocation_stop(alloc_id, opts)



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

api_instance = NomadClient::AllocationsApi.new
alloc_id = 'alloc_id_example' # String | Allocation ID.
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
  no_shutdown_delay: true # Boolean | Flag indicating whether to delay shutdown when requesting an allocation stop.
}

begin
  
  result = api_instance.post_allocation_stop(alloc_id, opts)
  p result
rescue NomadClient::ApiError => e
  puts "Error when calling AllocationsApi->post_allocation_stop: #{e}"
end
```

#### Using the post_allocation_stop_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<AllocStopResponse>, Integer, Hash)> post_allocation_stop_with_http_info(alloc_id, opts)

```ruby
begin
  
  data, status_code, headers = api_instance.post_allocation_stop_with_http_info(alloc_id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <AllocStopResponse>
rescue NomadClient::ApiError => e
  puts "Error when calling AllocationsApi->post_allocation_stop_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **alloc_id** | **String** | Allocation ID. |  |
| **region** | **String** | Filters results based on the specified region. | [optional] |
| **namespace** | **String** | Filters results based on the specified namespace. | [optional] |
| **index** | **Integer** | If set, wait until query exceeds given index. Must be provided with WaitParam. | [optional] |
| **wait** | **String** | Provided with IndexParam to wait for change. | [optional] |
| **stale** | **String** | If present, results will include stale reads. | [optional] |
| **prefix** | **String** | Constrains results to jobs that start with the defined prefix | [optional] |
| **x_nomad_token** | **String** | A Nomad ACL token. | [optional] |
| **per_page** | **Integer** | Maximum number of results to return. | [optional] |
| **next_token** | **String** | Indicates where to start paging for queries that support pagination. | [optional] |
| **no_shutdown_delay** | **Boolean** | Flag indicating whether to delay shutdown when requesting an allocation stop. | [optional] |

### Return type

[**AllocStopResponse**](AllocStopResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

