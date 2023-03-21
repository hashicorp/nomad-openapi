# NomadClient::DeploymentsApi

All URIs are relative to *https://127.0.0.1:4646/v1*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**get_deployment**](DeploymentsApi.md#get_deployment) | **GET** /deployment/{deploymentID} |  |
| [**get_deployment_allocations**](DeploymentsApi.md#get_deployment_allocations) | **GET** /deployment/allocations/{deploymentID} |  |
| [**get_deployments**](DeploymentsApi.md#get_deployments) | **GET** /deployments |  |
| [**post_deployment_allocation_health**](DeploymentsApi.md#post_deployment_allocation_health) | **POST** /deployment/allocation-health/{deploymentID} |  |
| [**post_deployment_fail**](DeploymentsApi.md#post_deployment_fail) | **POST** /deployment/fail/{deploymentID} |  |
| [**post_deployment_pause**](DeploymentsApi.md#post_deployment_pause) | **POST** /deployment/pause/{deploymentID} |  |
| [**post_deployment_promote**](DeploymentsApi.md#post_deployment_promote) | **POST** /deployment/promote/{deploymentID} |  |
| [**post_deployment_unblock**](DeploymentsApi.md#post_deployment_unblock) | **POST** /deployment/unblock/{deploymentID} |  |


## get_deployment

> <Deployment> get_deployment(deployment_id, opts)



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

api_instance = NomadClient::DeploymentsApi.new
deployment_id = 'deployment_id_example' # String | Deployment ID.
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
  
  result = api_instance.get_deployment(deployment_id, opts)
  p result
rescue NomadClient::ApiError => e
  puts "Error when calling DeploymentsApi->get_deployment: #{e}"
end
```

#### Using the get_deployment_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Deployment>, Integer, Hash)> get_deployment_with_http_info(deployment_id, opts)

```ruby
begin
  
  data, status_code, headers = api_instance.get_deployment_with_http_info(deployment_id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Deployment>
rescue NomadClient::ApiError => e
  puts "Error when calling DeploymentsApi->get_deployment_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **deployment_id** | **String** | Deployment ID. |  |
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

[**Deployment**](Deployment.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## get_deployment_allocations

> <Array<AllocationListStub>> get_deployment_allocations(deployment_id, opts)



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

api_instance = NomadClient::DeploymentsApi.new
deployment_id = 'deployment_id_example' # String | Deployment ID.
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
  
  result = api_instance.get_deployment_allocations(deployment_id, opts)
  p result
rescue NomadClient::ApiError => e
  puts "Error when calling DeploymentsApi->get_deployment_allocations: #{e}"
end
```

#### Using the get_deployment_allocations_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Array<AllocationListStub>>, Integer, Hash)> get_deployment_allocations_with_http_info(deployment_id, opts)

```ruby
begin
  
  data, status_code, headers = api_instance.get_deployment_allocations_with_http_info(deployment_id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Array<AllocationListStub>>
rescue NomadClient::ApiError => e
  puts "Error when calling DeploymentsApi->get_deployment_allocations_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **deployment_id** | **String** | Deployment ID. |  |
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


## get_deployments

> <Array<Deployment>> get_deployments(opts)



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

api_instance = NomadClient::DeploymentsApi.new
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
  
  result = api_instance.get_deployments(opts)
  p result
rescue NomadClient::ApiError => e
  puts "Error when calling DeploymentsApi->get_deployments: #{e}"
end
```

#### Using the get_deployments_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Array<Deployment>>, Integer, Hash)> get_deployments_with_http_info(opts)

```ruby
begin
  
  data, status_code, headers = api_instance.get_deployments_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Array<Deployment>>
rescue NomadClient::ApiError => e
  puts "Error when calling DeploymentsApi->get_deployments_with_http_info: #{e}"
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

### Return type

[**Array&lt;Deployment&gt;**](Deployment.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## post_deployment_allocation_health

> <DeploymentUpdateResponse> post_deployment_allocation_health(deployment_id, deployment_alloc_health_request, opts)



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

api_instance = NomadClient::DeploymentsApi.new
deployment_id = 'deployment_id_example' # String | Deployment ID.
deployment_alloc_health_request = NomadClient::DeploymentAllocHealthRequest.new # DeploymentAllocHealthRequest | 
opts = {
  region: 'region_example', # String | Filters results based on the specified region.
  namespace: 'namespace_example', # String | Filters results based on the specified namespace.
  x_nomad_token: 'x_nomad_token_example', # String | A Nomad ACL token.
  idempotency_token: 'idempotency_token_example' # String | Can be used to ensure operations are only run once.
}

begin
  
  result = api_instance.post_deployment_allocation_health(deployment_id, deployment_alloc_health_request, opts)
  p result
rescue NomadClient::ApiError => e
  puts "Error when calling DeploymentsApi->post_deployment_allocation_health: #{e}"
end
```

#### Using the post_deployment_allocation_health_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<DeploymentUpdateResponse>, Integer, Hash)> post_deployment_allocation_health_with_http_info(deployment_id, deployment_alloc_health_request, opts)

```ruby
begin
  
  data, status_code, headers = api_instance.post_deployment_allocation_health_with_http_info(deployment_id, deployment_alloc_health_request, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <DeploymentUpdateResponse>
rescue NomadClient::ApiError => e
  puts "Error when calling DeploymentsApi->post_deployment_allocation_health_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **deployment_id** | **String** | Deployment ID. |  |
| **deployment_alloc_health_request** | [**DeploymentAllocHealthRequest**](DeploymentAllocHealthRequest.md) |  |  |
| **region** | **String** | Filters results based on the specified region. | [optional] |
| **namespace** | **String** | Filters results based on the specified namespace. | [optional] |
| **x_nomad_token** | **String** | A Nomad ACL token. | [optional] |
| **idempotency_token** | **String** | Can be used to ensure operations are only run once. | [optional] |

### Return type

[**DeploymentUpdateResponse**](DeploymentUpdateResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## post_deployment_fail

> <DeploymentUpdateResponse> post_deployment_fail(deployment_id, opts)



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

api_instance = NomadClient::DeploymentsApi.new
deployment_id = 'deployment_id_example' # String | Deployment ID.
opts = {
  region: 'region_example', # String | Filters results based on the specified region.
  namespace: 'namespace_example', # String | Filters results based on the specified namespace.
  x_nomad_token: 'x_nomad_token_example', # String | A Nomad ACL token.
  idempotency_token: 'idempotency_token_example' # String | Can be used to ensure operations are only run once.
}

begin
  
  result = api_instance.post_deployment_fail(deployment_id, opts)
  p result
rescue NomadClient::ApiError => e
  puts "Error when calling DeploymentsApi->post_deployment_fail: #{e}"
end
```

#### Using the post_deployment_fail_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<DeploymentUpdateResponse>, Integer, Hash)> post_deployment_fail_with_http_info(deployment_id, opts)

```ruby
begin
  
  data, status_code, headers = api_instance.post_deployment_fail_with_http_info(deployment_id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <DeploymentUpdateResponse>
rescue NomadClient::ApiError => e
  puts "Error when calling DeploymentsApi->post_deployment_fail_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **deployment_id** | **String** | Deployment ID. |  |
| **region** | **String** | Filters results based on the specified region. | [optional] |
| **namespace** | **String** | Filters results based on the specified namespace. | [optional] |
| **x_nomad_token** | **String** | A Nomad ACL token. | [optional] |
| **idempotency_token** | **String** | Can be used to ensure operations are only run once. | [optional] |

### Return type

[**DeploymentUpdateResponse**](DeploymentUpdateResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## post_deployment_pause

> <DeploymentUpdateResponse> post_deployment_pause(deployment_id, deployment_pause_request, opts)



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

api_instance = NomadClient::DeploymentsApi.new
deployment_id = 'deployment_id_example' # String | Deployment ID.
deployment_pause_request = NomadClient::DeploymentPauseRequest.new # DeploymentPauseRequest | 
opts = {
  region: 'region_example', # String | Filters results based on the specified region.
  namespace: 'namespace_example', # String | Filters results based on the specified namespace.
  x_nomad_token: 'x_nomad_token_example', # String | A Nomad ACL token.
  idempotency_token: 'idempotency_token_example' # String | Can be used to ensure operations are only run once.
}

begin
  
  result = api_instance.post_deployment_pause(deployment_id, deployment_pause_request, opts)
  p result
rescue NomadClient::ApiError => e
  puts "Error when calling DeploymentsApi->post_deployment_pause: #{e}"
end
```

#### Using the post_deployment_pause_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<DeploymentUpdateResponse>, Integer, Hash)> post_deployment_pause_with_http_info(deployment_id, deployment_pause_request, opts)

```ruby
begin
  
  data, status_code, headers = api_instance.post_deployment_pause_with_http_info(deployment_id, deployment_pause_request, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <DeploymentUpdateResponse>
rescue NomadClient::ApiError => e
  puts "Error when calling DeploymentsApi->post_deployment_pause_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **deployment_id** | **String** | Deployment ID. |  |
| **deployment_pause_request** | [**DeploymentPauseRequest**](DeploymentPauseRequest.md) |  |  |
| **region** | **String** | Filters results based on the specified region. | [optional] |
| **namespace** | **String** | Filters results based on the specified namespace. | [optional] |
| **x_nomad_token** | **String** | A Nomad ACL token. | [optional] |
| **idempotency_token** | **String** | Can be used to ensure operations are only run once. | [optional] |

### Return type

[**DeploymentUpdateResponse**](DeploymentUpdateResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## post_deployment_promote

> <DeploymentUpdateResponse> post_deployment_promote(deployment_id, deployment_promote_request, opts)



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

api_instance = NomadClient::DeploymentsApi.new
deployment_id = 'deployment_id_example' # String | Deployment ID.
deployment_promote_request = NomadClient::DeploymentPromoteRequest.new # DeploymentPromoteRequest | 
opts = {
  region: 'region_example', # String | Filters results based on the specified region.
  namespace: 'namespace_example', # String | Filters results based on the specified namespace.
  x_nomad_token: 'x_nomad_token_example', # String | A Nomad ACL token.
  idempotency_token: 'idempotency_token_example' # String | Can be used to ensure operations are only run once.
}

begin
  
  result = api_instance.post_deployment_promote(deployment_id, deployment_promote_request, opts)
  p result
rescue NomadClient::ApiError => e
  puts "Error when calling DeploymentsApi->post_deployment_promote: #{e}"
end
```

#### Using the post_deployment_promote_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<DeploymentUpdateResponse>, Integer, Hash)> post_deployment_promote_with_http_info(deployment_id, deployment_promote_request, opts)

```ruby
begin
  
  data, status_code, headers = api_instance.post_deployment_promote_with_http_info(deployment_id, deployment_promote_request, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <DeploymentUpdateResponse>
rescue NomadClient::ApiError => e
  puts "Error when calling DeploymentsApi->post_deployment_promote_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **deployment_id** | **String** | Deployment ID. |  |
| **deployment_promote_request** | [**DeploymentPromoteRequest**](DeploymentPromoteRequest.md) |  |  |
| **region** | **String** | Filters results based on the specified region. | [optional] |
| **namespace** | **String** | Filters results based on the specified namespace. | [optional] |
| **x_nomad_token** | **String** | A Nomad ACL token. | [optional] |
| **idempotency_token** | **String** | Can be used to ensure operations are only run once. | [optional] |

### Return type

[**DeploymentUpdateResponse**](DeploymentUpdateResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## post_deployment_unblock

> <DeploymentUpdateResponse> post_deployment_unblock(deployment_id, deployment_unblock_request, opts)



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

api_instance = NomadClient::DeploymentsApi.new
deployment_id = 'deployment_id_example' # String | Deployment ID.
deployment_unblock_request = NomadClient::DeploymentUnblockRequest.new # DeploymentUnblockRequest | 
opts = {
  region: 'region_example', # String | Filters results based on the specified region.
  namespace: 'namespace_example', # String | Filters results based on the specified namespace.
  x_nomad_token: 'x_nomad_token_example', # String | A Nomad ACL token.
  idempotency_token: 'idempotency_token_example' # String | Can be used to ensure operations are only run once.
}

begin
  
  result = api_instance.post_deployment_unblock(deployment_id, deployment_unblock_request, opts)
  p result
rescue NomadClient::ApiError => e
  puts "Error when calling DeploymentsApi->post_deployment_unblock: #{e}"
end
```

#### Using the post_deployment_unblock_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<DeploymentUpdateResponse>, Integer, Hash)> post_deployment_unblock_with_http_info(deployment_id, deployment_unblock_request, opts)

```ruby
begin
  
  data, status_code, headers = api_instance.post_deployment_unblock_with_http_info(deployment_id, deployment_unblock_request, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <DeploymentUpdateResponse>
rescue NomadClient::ApiError => e
  puts "Error when calling DeploymentsApi->post_deployment_unblock_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **deployment_id** | **String** | Deployment ID. |  |
| **deployment_unblock_request** | [**DeploymentUnblockRequest**](DeploymentUnblockRequest.md) |  |  |
| **region** | **String** | Filters results based on the specified region. | [optional] |
| **namespace** | **String** | Filters results based on the specified namespace. | [optional] |
| **x_nomad_token** | **String** | A Nomad ACL token. | [optional] |
| **idempotency_token** | **String** | Can be used to ensure operations are only run once. | [optional] |

### Return type

[**DeploymentUpdateResponse**](DeploymentUpdateResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

