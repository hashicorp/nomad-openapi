# NomadClient::ACLApi

All URIs are relative to *https://127.0.0.1:4646/v1*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**delete_acl_policy**](ACLApi.md#delete_acl_policy) | **DELETE** /acl/policy/{policyName} |  |
| [**delete_acl_token**](ACLApi.md#delete_acl_token) | **DELETE** /acl/token/{tokenAccessor} |  |
| [**get_acl_policies**](ACLApi.md#get_acl_policies) | **GET** /acl/policies |  |
| [**get_acl_policy**](ACLApi.md#get_acl_policy) | **GET** /acl/policy/{policyName} |  |
| [**get_acl_token**](ACLApi.md#get_acl_token) | **GET** /acl/token/{tokenAccessor} |  |
| [**get_acl_token_self**](ACLApi.md#get_acl_token_self) | **GET** /acl/token |  |
| [**get_acl_tokens**](ACLApi.md#get_acl_tokens) | **GET** /acl/tokens |  |
| [**post_acl_bootstrap**](ACLApi.md#post_acl_bootstrap) | **POST** /acl/bootstrap |  |
| [**post_acl_policy**](ACLApi.md#post_acl_policy) | **POST** /acl/policy/{policyName} |  |
| [**post_acl_token**](ACLApi.md#post_acl_token) | **POST** /acl/token/{tokenAccessor} |  |
| [**post_acl_token_onetime**](ACLApi.md#post_acl_token_onetime) | **POST** /acl/token/onetime |  |
| [**post_acl_token_onetime_exchange**](ACLApi.md#post_acl_token_onetime_exchange) | **POST** /acl/token/onetime/exchange |  |


## delete_acl_policy

> delete_acl_policy(policy_name, opts)



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

api_instance = NomadClient::ACLApi.new
policy_name = 'policy_name_example' # String | The ACL policy name.
opts = {
  region: 'region_example', # String | Filters results based on the specified region.
  namespace: 'namespace_example', # String | Filters results based on the specified namespace.
  x_nomad_token: 'x_nomad_token_example', # String | A Nomad ACL token.
  idempotency_token: 'idempotency_token_example' # String | Can be used to ensure operations are only run once.
}

begin
  
  api_instance.delete_acl_policy(policy_name, opts)
rescue NomadClient::ApiError => e
  puts "Error when calling ACLApi->delete_acl_policy: #{e}"
end
```

#### Using the delete_acl_policy_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> delete_acl_policy_with_http_info(policy_name, opts)

```ruby
begin
  
  data, status_code, headers = api_instance.delete_acl_policy_with_http_info(policy_name, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue NomadClient::ApiError => e
  puts "Error when calling ACLApi->delete_acl_policy_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **policy_name** | **String** | The ACL policy name. |  |
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


## delete_acl_token

> delete_acl_token(token_accessor, opts)



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

api_instance = NomadClient::ACLApi.new
token_accessor = 'token_accessor_example' # String | The token accessor ID.
opts = {
  region: 'region_example', # String | Filters results based on the specified region.
  namespace: 'namespace_example', # String | Filters results based on the specified namespace.
  x_nomad_token: 'x_nomad_token_example', # String | A Nomad ACL token.
  idempotency_token: 'idempotency_token_example' # String | Can be used to ensure operations are only run once.
}

begin
  
  api_instance.delete_acl_token(token_accessor, opts)
rescue NomadClient::ApiError => e
  puts "Error when calling ACLApi->delete_acl_token: #{e}"
end
```

#### Using the delete_acl_token_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> delete_acl_token_with_http_info(token_accessor, opts)

```ruby
begin
  
  data, status_code, headers = api_instance.delete_acl_token_with_http_info(token_accessor, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue NomadClient::ApiError => e
  puts "Error when calling ACLApi->delete_acl_token_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **token_accessor** | **String** | The token accessor ID. |  |
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


## get_acl_policies

> <Array<ACLPolicyListStub>> get_acl_policies(opts)



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

api_instance = NomadClient::ACLApi.new
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
  
  result = api_instance.get_acl_policies(opts)
  p result
rescue NomadClient::ApiError => e
  puts "Error when calling ACLApi->get_acl_policies: #{e}"
end
```

#### Using the get_acl_policies_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Array<ACLPolicyListStub>>, Integer, Hash)> get_acl_policies_with_http_info(opts)

```ruby
begin
  
  data, status_code, headers = api_instance.get_acl_policies_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Array<ACLPolicyListStub>>
rescue NomadClient::ApiError => e
  puts "Error when calling ACLApi->get_acl_policies_with_http_info: #{e}"
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

[**Array&lt;ACLPolicyListStub&gt;**](ACLPolicyListStub.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## get_acl_policy

> <ACLPolicy> get_acl_policy(policy_name, opts)



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

api_instance = NomadClient::ACLApi.new
policy_name = 'policy_name_example' # String | The ACL policy name.
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
  
  result = api_instance.get_acl_policy(policy_name, opts)
  p result
rescue NomadClient::ApiError => e
  puts "Error when calling ACLApi->get_acl_policy: #{e}"
end
```

#### Using the get_acl_policy_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<ACLPolicy>, Integer, Hash)> get_acl_policy_with_http_info(policy_name, opts)

```ruby
begin
  
  data, status_code, headers = api_instance.get_acl_policy_with_http_info(policy_name, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <ACLPolicy>
rescue NomadClient::ApiError => e
  puts "Error when calling ACLApi->get_acl_policy_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **policy_name** | **String** | The ACL policy name. |  |
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

[**ACLPolicy**](ACLPolicy.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## get_acl_token

> <ACLToken> get_acl_token(token_accessor, opts)



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

api_instance = NomadClient::ACLApi.new
token_accessor = 'token_accessor_example' # String | The token accessor ID.
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
  
  result = api_instance.get_acl_token(token_accessor, opts)
  p result
rescue NomadClient::ApiError => e
  puts "Error when calling ACLApi->get_acl_token: #{e}"
end
```

#### Using the get_acl_token_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<ACLToken>, Integer, Hash)> get_acl_token_with_http_info(token_accessor, opts)

```ruby
begin
  
  data, status_code, headers = api_instance.get_acl_token_with_http_info(token_accessor, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <ACLToken>
rescue NomadClient::ApiError => e
  puts "Error when calling ACLApi->get_acl_token_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **token_accessor** | **String** | The token accessor ID. |  |
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

[**ACLToken**](ACLToken.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## get_acl_token_self

> <ACLToken> get_acl_token_self(opts)



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

api_instance = NomadClient::ACLApi.new
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
  
  result = api_instance.get_acl_token_self(opts)
  p result
rescue NomadClient::ApiError => e
  puts "Error when calling ACLApi->get_acl_token_self: #{e}"
end
```

#### Using the get_acl_token_self_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<ACLToken>, Integer, Hash)> get_acl_token_self_with_http_info(opts)

```ruby
begin
  
  data, status_code, headers = api_instance.get_acl_token_self_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <ACLToken>
rescue NomadClient::ApiError => e
  puts "Error when calling ACLApi->get_acl_token_self_with_http_info: #{e}"
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

[**ACLToken**](ACLToken.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## get_acl_tokens

> <Array<ACLTokenListStub>> get_acl_tokens(opts)



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

api_instance = NomadClient::ACLApi.new
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
  
  result = api_instance.get_acl_tokens(opts)
  p result
rescue NomadClient::ApiError => e
  puts "Error when calling ACLApi->get_acl_tokens: #{e}"
end
```

#### Using the get_acl_tokens_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Array<ACLTokenListStub>>, Integer, Hash)> get_acl_tokens_with_http_info(opts)

```ruby
begin
  
  data, status_code, headers = api_instance.get_acl_tokens_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Array<ACLTokenListStub>>
rescue NomadClient::ApiError => e
  puts "Error when calling ACLApi->get_acl_tokens_with_http_info: #{e}"
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

[**Array&lt;ACLTokenListStub&gt;**](ACLTokenListStub.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## post_acl_bootstrap

> <Array<ACLToken>> post_acl_bootstrap(opts)



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

api_instance = NomadClient::ACLApi.new
opts = {
  region: 'region_example', # String | Filters results based on the specified region.
  namespace: 'namespace_example', # String | Filters results based on the specified namespace.
  x_nomad_token: 'x_nomad_token_example', # String | A Nomad ACL token.
  idempotency_token: 'idempotency_token_example' # String | Can be used to ensure operations are only run once.
}

begin
  
  result = api_instance.post_acl_bootstrap(opts)
  p result
rescue NomadClient::ApiError => e
  puts "Error when calling ACLApi->post_acl_bootstrap: #{e}"
end
```

#### Using the post_acl_bootstrap_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Array<ACLToken>>, Integer, Hash)> post_acl_bootstrap_with_http_info(opts)

```ruby
begin
  
  data, status_code, headers = api_instance.post_acl_bootstrap_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Array<ACLToken>>
rescue NomadClient::ApiError => e
  puts "Error when calling ACLApi->post_acl_bootstrap_with_http_info: #{e}"
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

[**Array&lt;ACLToken&gt;**](ACLToken.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## post_acl_policy

> post_acl_policy(policy_name, acl_policy, opts)



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

api_instance = NomadClient::ACLApi.new
policy_name = 'policy_name_example' # String | The ACL policy name.
acl_policy = NomadClient::ACLPolicy.new # ACLPolicy | 
opts = {
  region: 'region_example', # String | Filters results based on the specified region.
  namespace: 'namespace_example', # String | Filters results based on the specified namespace.
  x_nomad_token: 'x_nomad_token_example', # String | A Nomad ACL token.
  idempotency_token: 'idempotency_token_example' # String | Can be used to ensure operations are only run once.
}

begin
  
  api_instance.post_acl_policy(policy_name, acl_policy, opts)
rescue NomadClient::ApiError => e
  puts "Error when calling ACLApi->post_acl_policy: #{e}"
end
```

#### Using the post_acl_policy_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> post_acl_policy_with_http_info(policy_name, acl_policy, opts)

```ruby
begin
  
  data, status_code, headers = api_instance.post_acl_policy_with_http_info(policy_name, acl_policy, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue NomadClient::ApiError => e
  puts "Error when calling ACLApi->post_acl_policy_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **policy_name** | **String** | The ACL policy name. |  |
| **acl_policy** | [**ACLPolicy**](ACLPolicy.md) |  |  |
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


## post_acl_token

> <ACLToken> post_acl_token(token_accessor, acl_token, opts)



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

api_instance = NomadClient::ACLApi.new
token_accessor = 'token_accessor_example' # String | The token accessor ID.
acl_token = NomadClient::ACLToken.new # ACLToken | 
opts = {
  region: 'region_example', # String | Filters results based on the specified region.
  namespace: 'namespace_example', # String | Filters results based on the specified namespace.
  x_nomad_token: 'x_nomad_token_example', # String | A Nomad ACL token.
  idempotency_token: 'idempotency_token_example' # String | Can be used to ensure operations are only run once.
}

begin
  
  result = api_instance.post_acl_token(token_accessor, acl_token, opts)
  p result
rescue NomadClient::ApiError => e
  puts "Error when calling ACLApi->post_acl_token: #{e}"
end
```

#### Using the post_acl_token_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<ACLToken>, Integer, Hash)> post_acl_token_with_http_info(token_accessor, acl_token, opts)

```ruby
begin
  
  data, status_code, headers = api_instance.post_acl_token_with_http_info(token_accessor, acl_token, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <ACLToken>
rescue NomadClient::ApiError => e
  puts "Error when calling ACLApi->post_acl_token_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **token_accessor** | **String** | The token accessor ID. |  |
| **acl_token** | [**ACLToken**](ACLToken.md) |  |  |
| **region** | **String** | Filters results based on the specified region. | [optional] |
| **namespace** | **String** | Filters results based on the specified namespace. | [optional] |
| **x_nomad_token** | **String** | A Nomad ACL token. | [optional] |
| **idempotency_token** | **String** | Can be used to ensure operations are only run once. | [optional] |

### Return type

[**ACLToken**](ACLToken.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## post_acl_token_onetime

> <OneTimeToken> post_acl_token_onetime(opts)



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

api_instance = NomadClient::ACLApi.new
opts = {
  region: 'region_example', # String | Filters results based on the specified region.
  namespace: 'namespace_example', # String | Filters results based on the specified namespace.
  x_nomad_token: 'x_nomad_token_example', # String | A Nomad ACL token.
  idempotency_token: 'idempotency_token_example' # String | Can be used to ensure operations are only run once.
}

begin
  
  result = api_instance.post_acl_token_onetime(opts)
  p result
rescue NomadClient::ApiError => e
  puts "Error when calling ACLApi->post_acl_token_onetime: #{e}"
end
```

#### Using the post_acl_token_onetime_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<OneTimeToken>, Integer, Hash)> post_acl_token_onetime_with_http_info(opts)

```ruby
begin
  
  data, status_code, headers = api_instance.post_acl_token_onetime_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <OneTimeToken>
rescue NomadClient::ApiError => e
  puts "Error when calling ACLApi->post_acl_token_onetime_with_http_info: #{e}"
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

[**OneTimeToken**](OneTimeToken.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## post_acl_token_onetime_exchange

> <ACLToken> post_acl_token_onetime_exchange(one_time_token_exchange_request, opts)



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

api_instance = NomadClient::ACLApi.new
one_time_token_exchange_request = NomadClient::OneTimeTokenExchangeRequest.new # OneTimeTokenExchangeRequest | 
opts = {
  region: 'region_example', # String | Filters results based on the specified region.
  namespace: 'namespace_example', # String | Filters results based on the specified namespace.
  x_nomad_token: 'x_nomad_token_example', # String | A Nomad ACL token.
  idempotency_token: 'idempotency_token_example' # String | Can be used to ensure operations are only run once.
}

begin
  
  result = api_instance.post_acl_token_onetime_exchange(one_time_token_exchange_request, opts)
  p result
rescue NomadClient::ApiError => e
  puts "Error when calling ACLApi->post_acl_token_onetime_exchange: #{e}"
end
```

#### Using the post_acl_token_onetime_exchange_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<ACLToken>, Integer, Hash)> post_acl_token_onetime_exchange_with_http_info(one_time_token_exchange_request, opts)

```ruby
begin
  
  data, status_code, headers = api_instance.post_acl_token_onetime_exchange_with_http_info(one_time_token_exchange_request, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <ACLToken>
rescue NomadClient::ApiError => e
  puts "Error when calling ACLApi->post_acl_token_onetime_exchange_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **one_time_token_exchange_request** | [**OneTimeTokenExchangeRequest**](OneTimeTokenExchangeRequest.md) |  |  |
| **region** | **String** | Filters results based on the specified region. | [optional] |
| **namespace** | **String** | Filters results based on the specified namespace. | [optional] |
| **x_nomad_token** | **String** | A Nomad ACL token. | [optional] |
| **idempotency_token** | **String** | Can be used to ensure operations are only run once. | [optional] |

### Return type

[**ACLToken**](ACLToken.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

