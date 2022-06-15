# NomadClient::EnterpriseApi

All URIs are relative to *https://127.0.0.1:4646/v1*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**create_quota_spec**](EnterpriseApi.md#create_quota_spec) | **POST** /quota |  |
| [**delete_quota_spec**](EnterpriseApi.md#delete_quota_spec) | **DELETE** /quota/{specName} |  |
| [**get_quota_spec**](EnterpriseApi.md#get_quota_spec) | **GET** /quota/{specName} |  |
| [**get_quotas**](EnterpriseApi.md#get_quotas) | **GET** /quotas |  |
| [**post_quota_spec**](EnterpriseApi.md#post_quota_spec) | **POST** /quota/{specName} |  |


## create_quota_spec

> create_quota_spec(quota_spec, opts)



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

api_instance = NomadClient::EnterpriseApi.new
quota_spec = NomadClient::QuotaSpec.new # QuotaSpec | 
opts = {
  region: 'region_example', # String | Filters results based on the specified region.
  namespace: 'namespace_example', # String | Filters results based on the specified namespace.
  x_nomad_token: 'x_nomad_token_example', # String | A Nomad ACL token.
  idempotency_token: 'idempotency_token_example' # String | Can be used to ensure operations are only run once.
}

begin
  
  api_instance.create_quota_spec(quota_spec, opts)
rescue NomadClient::ApiError => e
  puts "Error when calling EnterpriseApi->create_quota_spec: #{e}"
end
```

#### Using the create_quota_spec_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> create_quota_spec_with_http_info(quota_spec, opts)

```ruby
begin
  
  data, status_code, headers = api_instance.create_quota_spec_with_http_info(quota_spec, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue NomadClient::ApiError => e
  puts "Error when calling EnterpriseApi->create_quota_spec_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **quota_spec** | [**QuotaSpec**](QuotaSpec.md) |  |  |
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


## delete_quota_spec

> delete_quota_spec(spec_name, opts)



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

api_instance = NomadClient::EnterpriseApi.new
spec_name = 'spec_name_example' # String | The quota spec identifier.
opts = {
  region: 'region_example', # String | Filters results based on the specified region.
  namespace: 'namespace_example', # String | Filters results based on the specified namespace.
  x_nomad_token: 'x_nomad_token_example', # String | A Nomad ACL token.
  idempotency_token: 'idempotency_token_example' # String | Can be used to ensure operations are only run once.
}

begin
  
  api_instance.delete_quota_spec(spec_name, opts)
rescue NomadClient::ApiError => e
  puts "Error when calling EnterpriseApi->delete_quota_spec: #{e}"
end
```

#### Using the delete_quota_spec_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> delete_quota_spec_with_http_info(spec_name, opts)

```ruby
begin
  
  data, status_code, headers = api_instance.delete_quota_spec_with_http_info(spec_name, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue NomadClient::ApiError => e
  puts "Error when calling EnterpriseApi->delete_quota_spec_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **spec_name** | **String** | The quota spec identifier. |  |
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


## get_quota_spec

> <QuotaSpec> get_quota_spec(spec_name, opts)



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

api_instance = NomadClient::EnterpriseApi.new
spec_name = 'spec_name_example' # String | The quota spec identifier.
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
  
  result = api_instance.get_quota_spec(spec_name, opts)
  p result
rescue NomadClient::ApiError => e
  puts "Error when calling EnterpriseApi->get_quota_spec: #{e}"
end
```

#### Using the get_quota_spec_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<QuotaSpec>, Integer, Hash)> get_quota_spec_with_http_info(spec_name, opts)

```ruby
begin
  
  data, status_code, headers = api_instance.get_quota_spec_with_http_info(spec_name, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <QuotaSpec>
rescue NomadClient::ApiError => e
  puts "Error when calling EnterpriseApi->get_quota_spec_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **spec_name** | **String** | The quota spec identifier. |  |
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

[**QuotaSpec**](QuotaSpec.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## get_quotas

> Array&lt;Object&gt; get_quotas(opts)



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

api_instance = NomadClient::EnterpriseApi.new
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
  
  result = api_instance.get_quotas(opts)
  p result
rescue NomadClient::ApiError => e
  puts "Error when calling EnterpriseApi->get_quotas: #{e}"
end
```

#### Using the get_quotas_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(Array&lt;Object&gt;, Integer, Hash)> get_quotas_with_http_info(opts)

```ruby
begin
  
  data, status_code, headers = api_instance.get_quotas_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => Array&lt;Object&gt;
rescue NomadClient::ApiError => e
  puts "Error when calling EnterpriseApi->get_quotas_with_http_info: #{e}"
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

**Array&lt;Object&gt;**

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## post_quota_spec

> post_quota_spec(spec_name, quota_spec, opts)



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

api_instance = NomadClient::EnterpriseApi.new
spec_name = 'spec_name_example' # String | The quota spec identifier.
quota_spec = NomadClient::QuotaSpec.new # QuotaSpec | 
opts = {
  region: 'region_example', # String | Filters results based on the specified region.
  namespace: 'namespace_example', # String | Filters results based on the specified namespace.
  x_nomad_token: 'x_nomad_token_example', # String | A Nomad ACL token.
  idempotency_token: 'idempotency_token_example' # String | Can be used to ensure operations are only run once.
}

begin
  
  api_instance.post_quota_spec(spec_name, quota_spec, opts)
rescue NomadClient::ApiError => e
  puts "Error when calling EnterpriseApi->post_quota_spec: #{e}"
end
```

#### Using the post_quota_spec_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> post_quota_spec_with_http_info(spec_name, quota_spec, opts)

```ruby
begin
  
  data, status_code, headers = api_instance.post_quota_spec_with_http_info(spec_name, quota_spec, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue NomadClient::ApiError => e
  puts "Error when calling EnterpriseApi->post_quota_spec_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **spec_name** | **String** | The quota spec identifier. |  |
| **quota_spec** | [**QuotaSpec**](QuotaSpec.md) |  |  |
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

