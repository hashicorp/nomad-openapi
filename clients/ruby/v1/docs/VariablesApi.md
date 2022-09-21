# NomadClient::VariablesApi

All URIs are relative to *http://127.0.0.1:4646/v1*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**delete_variable**](VariablesApi.md#delete_variable) | **DELETE** /var/{path} |  |
| [**get_variable_query**](VariablesApi.md#get_variable_query) | **GET** /var/{path} |  |
| [**get_variables_list_request**](VariablesApi.md#get_variables_list_request) | **GET** /vars |  |
| [**post_variable**](VariablesApi.md#post_variable) | **POST** /var/{path} |  |
| [**put_variable**](VariablesApi.md#put_variable) | **PUT** /var/{path} |  |


## delete_variable

> delete_variable(path, variable, opts)



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

api_instance = NomadClient::VariablesApi.new
path = 'path_example' # String | A path to a Nomad Variable
variable = NomadClient::Variable.new # Variable | 
opts = {
  region: 'region_example', # String | Filters results based on the specified region.
  namespace: 'namespace_example', # String | Filters results based on the specified namespace.
  x_nomad_token: 'x_nomad_token_example', # String | A Nomad ACL token.
  idempotency_token: 'idempotency_token_example', # String | Can be used to ensure operations are only run once.
  cas: 56 # Integer | A compare-and-set parameter for Nomad Variables
}

begin
  
  api_instance.delete_variable(path, variable, opts)
rescue NomadClient::ApiError => e
  puts "Error when calling VariablesApi->delete_variable: #{e}"
end
```

#### Using the delete_variable_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> delete_variable_with_http_info(path, variable, opts)

```ruby
begin
  
  data, status_code, headers = api_instance.delete_variable_with_http_info(path, variable, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue NomadClient::ApiError => e
  puts "Error when calling VariablesApi->delete_variable_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **path** | **String** | A path to a Nomad Variable |  |
| **variable** | [**Variable**](Variable.md) |  |  |
| **region** | **String** | Filters results based on the specified region. | [optional] |
| **namespace** | **String** | Filters results based on the specified namespace. | [optional] |
| **x_nomad_token** | **String** | A Nomad ACL token. | [optional] |
| **idempotency_token** | **String** | Can be used to ensure operations are only run once. | [optional] |
| **cas** | **Integer** | A compare-and-set parameter for Nomad Variables | [optional] |

### Return type

nil (empty response body)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined


## get_variable_query

> <Variable> get_variable_query(path, opts)



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

api_instance = NomadClient::VariablesApi.new
path = 'path_example' # String | A path to a Nomad Variable
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
  
  result = api_instance.get_variable_query(path, opts)
  p result
rescue NomadClient::ApiError => e
  puts "Error when calling VariablesApi->get_variable_query: #{e}"
end
```

#### Using the get_variable_query_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Variable>, Integer, Hash)> get_variable_query_with_http_info(path, opts)

```ruby
begin
  
  data, status_code, headers = api_instance.get_variable_query_with_http_info(path, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Variable>
rescue NomadClient::ApiError => e
  puts "Error when calling VariablesApi->get_variable_query_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **path** | **String** | A path to a Nomad Variable |  |
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

[**Variable**](Variable.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## get_variables_list_request

> <Array<VariableMetadata>> get_variables_list_request(opts)



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

api_instance = NomadClient::VariablesApi.new
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
  
  result = api_instance.get_variables_list_request(opts)
  p result
rescue NomadClient::ApiError => e
  puts "Error when calling VariablesApi->get_variables_list_request: #{e}"
end
```

#### Using the get_variables_list_request_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Array<VariableMetadata>>, Integer, Hash)> get_variables_list_request_with_http_info(opts)

```ruby
begin
  
  data, status_code, headers = api_instance.get_variables_list_request_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Array<VariableMetadata>>
rescue NomadClient::ApiError => e
  puts "Error when calling VariablesApi->get_variables_list_request_with_http_info: #{e}"
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

[**Array&lt;VariableMetadata&gt;**](VariableMetadata.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## post_variable

> <Variable> post_variable(path, variable, opts)



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

api_instance = NomadClient::VariablesApi.new
path = 'path_example' # String | A path to a Nomad Variable
variable = NomadClient::Variable.new # Variable | 
opts = {
  region: 'region_example', # String | Filters results based on the specified region.
  namespace: 'namespace_example', # String | Filters results based on the specified namespace.
  x_nomad_token: 'x_nomad_token_example', # String | A Nomad ACL token.
  idempotency_token: 'idempotency_token_example', # String | Can be used to ensure operations are only run once.
  cas: 56 # Integer | A compare-and-set parameter for Nomad Variables
}

begin
  
  result = api_instance.post_variable(path, variable, opts)
  p result
rescue NomadClient::ApiError => e
  puts "Error when calling VariablesApi->post_variable: #{e}"
end
```

#### Using the post_variable_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Variable>, Integer, Hash)> post_variable_with_http_info(path, variable, opts)

```ruby
begin
  
  data, status_code, headers = api_instance.post_variable_with_http_info(path, variable, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Variable>
rescue NomadClient::ApiError => e
  puts "Error when calling VariablesApi->post_variable_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **path** | **String** | A path to a Nomad Variable |  |
| **variable** | [**Variable**](Variable.md) |  |  |
| **region** | **String** | Filters results based on the specified region. | [optional] |
| **namespace** | **String** | Filters results based on the specified namespace. | [optional] |
| **x_nomad_token** | **String** | A Nomad ACL token. | [optional] |
| **idempotency_token** | **String** | Can be used to ensure operations are only run once. | [optional] |
| **cas** | **Integer** | A compare-and-set parameter for Nomad Variables | [optional] |

### Return type

[**Variable**](Variable.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## put_variable

> <Variable> put_variable(path, variable, opts)



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

api_instance = NomadClient::VariablesApi.new
path = 'path_example' # String | A path to a Nomad Variable
variable = NomadClient::Variable.new # Variable | 
opts = {
  region: 'region_example', # String | Filters results based on the specified region.
  namespace: 'namespace_example', # String | Filters results based on the specified namespace.
  x_nomad_token: 'x_nomad_token_example', # String | A Nomad ACL token.
  idempotency_token: 'idempotency_token_example', # String | Can be used to ensure operations are only run once.
  cas: 56 # Integer | A compare-and-set parameter for Nomad Variables
}

begin
  
  result = api_instance.put_variable(path, variable, opts)
  p result
rescue NomadClient::ApiError => e
  puts "Error when calling VariablesApi->put_variable: #{e}"
end
```

#### Using the put_variable_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Variable>, Integer, Hash)> put_variable_with_http_info(path, variable, opts)

```ruby
begin
  
  data, status_code, headers = api_instance.put_variable_with_http_info(path, variable, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Variable>
rescue NomadClient::ApiError => e
  puts "Error when calling VariablesApi->put_variable_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **path** | **String** | A path to a Nomad Variable |  |
| **variable** | [**Variable**](Variable.md) |  |  |
| **region** | **String** | Filters results based on the specified region. | [optional] |
| **namespace** | **String** | Filters results based on the specified namespace. | [optional] |
| **x_nomad_token** | **String** | A Nomad ACL token. | [optional] |
| **idempotency_token** | **String** | Can be used to ensure operations are only run once. | [optional] |
| **cas** | **Integer** | A compare-and-set parameter for Nomad Variables | [optional] |

### Return type

[**Variable**](Variable.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

