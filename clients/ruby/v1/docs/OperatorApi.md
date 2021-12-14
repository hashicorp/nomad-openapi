# NomadClient::OperatorApi

All URIs are relative to *https://127.0.0.1:4646/v1*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**delete_operator_raft_peer**](OperatorApi.md#delete_operator_raft_peer) | **DELETE** /operator/raft/peer |  |
| [**get_operator_autopilot_configuration**](OperatorApi.md#get_operator_autopilot_configuration) | **GET** /operator/autopilot/configuration |  |
| [**get_operator_autopilot_health**](OperatorApi.md#get_operator_autopilot_health) | **GET** /operator/autopilot/health |  |
| [**get_operator_raft_configuration**](OperatorApi.md#get_operator_raft_configuration) | **GET** /operator/raft/configuration |  |
| [**get_operator_scheduler_configuration**](OperatorApi.md#get_operator_scheduler_configuration) | **GET** /operator/scheduler/configuration |  |
| [**post_operator_scheduler_configuration**](OperatorApi.md#post_operator_scheduler_configuration) | **POST** /operator/scheduler/configuration |  |
| [**put_operator_autopilot_configuration**](OperatorApi.md#put_operator_autopilot_configuration) | **PUT** /operator/autopilot/configuration |  |


## delete_operator_raft_peer

> delete_operator_raft_peer(opts)



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

api_instance = NomadClient::OperatorApi.new
opts = {
  region: 'region_example', # String | Filters results based on the specified region.
  namespace: 'namespace_example', # String | Filters results based on the specified namespace.
  x_nomad_token: 'x_nomad_token_example', # String | A Nomad ACL token.
  idempotency_token: 'idempotency_token_example' # String | Can be used to ensure operations are only run once.
}

begin
  
  api_instance.delete_operator_raft_peer(opts)
rescue NomadClient::ApiError => e
  puts "Error when calling OperatorApi->delete_operator_raft_peer: #{e}"
end
```

#### Using the delete_operator_raft_peer_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> delete_operator_raft_peer_with_http_info(opts)

```ruby
begin
  
  data, status_code, headers = api_instance.delete_operator_raft_peer_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue NomadClient::ApiError => e
  puts "Error when calling OperatorApi->delete_operator_raft_peer_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **region** | **String** | Filters results based on the specified region. | [optional] |
| **namespace** | **String** | Filters results based on the specified namespace. | [optional] |
| **x_nomad_token** | **String** | A Nomad ACL token. | [optional] |
| **idempotency_token** | **String** | Can be used to ensure operations are only run once. | [optional] |

### Return type

nil (empty response body)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined


## get_operator_autopilot_configuration

> <AutopilotConfiguration> get_operator_autopilot_configuration(opts)



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

api_instance = NomadClient::OperatorApi.new
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
  
  result = api_instance.get_operator_autopilot_configuration(opts)
  p result
rescue NomadClient::ApiError => e
  puts "Error when calling OperatorApi->get_operator_autopilot_configuration: #{e}"
end
```

#### Using the get_operator_autopilot_configuration_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<AutopilotConfiguration>, Integer, Hash)> get_operator_autopilot_configuration_with_http_info(opts)

```ruby
begin
  
  data, status_code, headers = api_instance.get_operator_autopilot_configuration_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <AutopilotConfiguration>
rescue NomadClient::ApiError => e
  puts "Error when calling OperatorApi->get_operator_autopilot_configuration_with_http_info: #{e}"
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

[**AutopilotConfiguration**](AutopilotConfiguration.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## get_operator_autopilot_health

> <OperatorHealthReply> get_operator_autopilot_health(opts)



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

api_instance = NomadClient::OperatorApi.new
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
  
  result = api_instance.get_operator_autopilot_health(opts)
  p result
rescue NomadClient::ApiError => e
  puts "Error when calling OperatorApi->get_operator_autopilot_health: #{e}"
end
```

#### Using the get_operator_autopilot_health_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OperatorHealthReply>, Integer, Hash)> get_operator_autopilot_health_with_http_info(opts)

```ruby
begin
  
  data, status_code, headers = api_instance.get_operator_autopilot_health_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OperatorHealthReply>
rescue NomadClient::ApiError => e
  puts "Error when calling OperatorApi->get_operator_autopilot_health_with_http_info: #{e}"
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

[**OperatorHealthReply**](OperatorHealthReply.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## get_operator_raft_configuration

> <Array<RaftServer>> get_operator_raft_configuration(opts)



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

api_instance = NomadClient::OperatorApi.new
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
  
  result = api_instance.get_operator_raft_configuration(opts)
  p result
rescue NomadClient::ApiError => e
  puts "Error when calling OperatorApi->get_operator_raft_configuration: #{e}"
end
```

#### Using the get_operator_raft_configuration_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Array<RaftServer>>, Integer, Hash)> get_operator_raft_configuration_with_http_info(opts)

```ruby
begin
  
  data, status_code, headers = api_instance.get_operator_raft_configuration_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Array<RaftServer>>
rescue NomadClient::ApiError => e
  puts "Error when calling OperatorApi->get_operator_raft_configuration_with_http_info: #{e}"
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

[**Array&lt;RaftServer&gt;**](RaftServer.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## get_operator_scheduler_configuration

> <SchedulerConfigurationResponse> get_operator_scheduler_configuration(opts)



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

api_instance = NomadClient::OperatorApi.new
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
  
  result = api_instance.get_operator_scheduler_configuration(opts)
  p result
rescue NomadClient::ApiError => e
  puts "Error when calling OperatorApi->get_operator_scheduler_configuration: #{e}"
end
```

#### Using the get_operator_scheduler_configuration_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<SchedulerConfigurationResponse>, Integer, Hash)> get_operator_scheduler_configuration_with_http_info(opts)

```ruby
begin
  
  data, status_code, headers = api_instance.get_operator_scheduler_configuration_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <SchedulerConfigurationResponse>
rescue NomadClient::ApiError => e
  puts "Error when calling OperatorApi->get_operator_scheduler_configuration_with_http_info: #{e}"
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

[**SchedulerConfigurationResponse**](SchedulerConfigurationResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## post_operator_scheduler_configuration

> <SchedulerSetConfigurationResponse> post_operator_scheduler_configuration(scheduler_configuration, opts)



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

api_instance = NomadClient::OperatorApi.new
scheduler_configuration = NomadClient::SchedulerConfiguration.new # SchedulerConfiguration | 
opts = {
  region: 'region_example', # String | Filters results based on the specified region.
  namespace: 'namespace_example', # String | Filters results based on the specified namespace.
  x_nomad_token: 'x_nomad_token_example', # String | A Nomad ACL token.
  idempotency_token: 'idempotency_token_example' # String | Can be used to ensure operations are only run once.
}

begin
  
  result = api_instance.post_operator_scheduler_configuration(scheduler_configuration, opts)
  p result
rescue NomadClient::ApiError => e
  puts "Error when calling OperatorApi->post_operator_scheduler_configuration: #{e}"
end
```

#### Using the post_operator_scheduler_configuration_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<SchedulerSetConfigurationResponse>, Integer, Hash)> post_operator_scheduler_configuration_with_http_info(scheduler_configuration, opts)

```ruby
begin
  
  data, status_code, headers = api_instance.post_operator_scheduler_configuration_with_http_info(scheduler_configuration, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <SchedulerSetConfigurationResponse>
rescue NomadClient::ApiError => e
  puts "Error when calling OperatorApi->post_operator_scheduler_configuration_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **scheduler_configuration** | [**SchedulerConfiguration**](SchedulerConfiguration.md) |  |  |
| **region** | **String** | Filters results based on the specified region. | [optional] |
| **namespace** | **String** | Filters results based on the specified namespace. | [optional] |
| **x_nomad_token** | **String** | A Nomad ACL token. | [optional] |
| **idempotency_token** | **String** | Can be used to ensure operations are only run once. | [optional] |

### Return type

[**SchedulerSetConfigurationResponse**](SchedulerSetConfigurationResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## put_operator_autopilot_configuration

> put_operator_autopilot_configuration(autopilot_configuration, opts)



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

api_instance = NomadClient::OperatorApi.new
autopilot_configuration = NomadClient::AutopilotConfiguration.new # AutopilotConfiguration | 
opts = {
  region: 'region_example', # String | Filters results based on the specified region.
  namespace: 'namespace_example', # String | Filters results based on the specified namespace.
  x_nomad_token: 'x_nomad_token_example', # String | A Nomad ACL token.
  idempotency_token: 'idempotency_token_example' # String | Can be used to ensure operations are only run once.
}

begin
  
  api_instance.put_operator_autopilot_configuration(autopilot_configuration, opts)
rescue NomadClient::ApiError => e
  puts "Error when calling OperatorApi->put_operator_autopilot_configuration: #{e}"
end
```

#### Using the put_operator_autopilot_configuration_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> put_operator_autopilot_configuration_with_http_info(autopilot_configuration, opts)

```ruby
begin
  
  data, status_code, headers = api_instance.put_operator_autopilot_configuration_with_http_info(autopilot_configuration, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue NomadClient::ApiError => e
  puts "Error when calling OperatorApi->put_operator_autopilot_configuration_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **autopilot_configuration** | [**AutopilotConfiguration**](AutopilotConfiguration.md) |  |  |
| **region** | **String** | Filters results based on the specified region. | [optional] |
| **namespace** | **String** | Filters results based on the specified namespace. | [optional] |
| **x_nomad_token** | **String** | A Nomad ACL token. | [optional] |
| **idempotency_token** | **String** | Can be used to ensure operations are only run once. | [optional] |

### Return type

nil (empty response body)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

