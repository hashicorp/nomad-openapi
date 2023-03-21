# NomadClient::ScalingApi

All URIs are relative to *https://127.0.0.1:4646/v1*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**get_scaling_policies**](ScalingApi.md#get_scaling_policies) | **GET** /scaling/policies |  |
| [**get_scaling_policy**](ScalingApi.md#get_scaling_policy) | **GET** /scaling/policy/{policyID} |  |


## get_scaling_policies

> <Array<ScalingPolicyListStub>> get_scaling_policies(opts)



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

api_instance = NomadClient::ScalingApi.new
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
  
  result = api_instance.get_scaling_policies(opts)
  p result
rescue NomadClient::ApiError => e
  puts "Error when calling ScalingApi->get_scaling_policies: #{e}"
end
```

#### Using the get_scaling_policies_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Array<ScalingPolicyListStub>>, Integer, Hash)> get_scaling_policies_with_http_info(opts)

```ruby
begin
  
  data, status_code, headers = api_instance.get_scaling_policies_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Array<ScalingPolicyListStub>>
rescue NomadClient::ApiError => e
  puts "Error when calling ScalingApi->get_scaling_policies_with_http_info: #{e}"
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

[**Array&lt;ScalingPolicyListStub&gt;**](ScalingPolicyListStub.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## get_scaling_policy

> <ScalingPolicy> get_scaling_policy(policy_id, opts)



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

api_instance = NomadClient::ScalingApi.new
policy_id = 'policy_id_example' # String | The scaling policy identifier.
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
  
  result = api_instance.get_scaling_policy(policy_id, opts)
  p result
rescue NomadClient::ApiError => e
  puts "Error when calling ScalingApi->get_scaling_policy: #{e}"
end
```

#### Using the get_scaling_policy_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<ScalingPolicy>, Integer, Hash)> get_scaling_policy_with_http_info(policy_id, opts)

```ruby
begin
  
  data, status_code, headers = api_instance.get_scaling_policy_with_http_info(policy_id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <ScalingPolicy>
rescue NomadClient::ApiError => e
  puts "Error when calling ScalingApi->get_scaling_policy_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **policy_id** | **String** | The scaling policy identifier. |  |
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

[**ScalingPolicy**](ScalingPolicy.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

