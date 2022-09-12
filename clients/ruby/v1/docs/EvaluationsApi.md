# NomadClient::EvaluationsApi

All URIs are relative to *http://127.0.0.1:4646/v1*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**get_evaluation**](EvaluationsApi.md#get_evaluation) | **GET** /evaluation/{evalID} |  |
| [**get_evaluation_allocations**](EvaluationsApi.md#get_evaluation_allocations) | **GET** /evaluation/{evalID}/allocations |  |
| [**get_evaluations**](EvaluationsApi.md#get_evaluations) | **GET** /evaluations |  |


## get_evaluation

> <Evaluation> get_evaluation(eval_id, opts)



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

api_instance = NomadClient::EvaluationsApi.new
eval_id = 'eval_id_example' # String | Evaluation ID.
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
  
  result = api_instance.get_evaluation(eval_id, opts)
  p result
rescue NomadClient::ApiError => e
  puts "Error when calling EvaluationsApi->get_evaluation: #{e}"
end
```

#### Using the get_evaluation_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Evaluation>, Integer, Hash)> get_evaluation_with_http_info(eval_id, opts)

```ruby
begin
  
  data, status_code, headers = api_instance.get_evaluation_with_http_info(eval_id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Evaluation>
rescue NomadClient::ApiError => e
  puts "Error when calling EvaluationsApi->get_evaluation_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **eval_id** | **String** | Evaluation ID. |  |
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

[**Evaluation**](Evaluation.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## get_evaluation_allocations

> <Array<AllocationListStub>> get_evaluation_allocations(eval_id, opts)



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

api_instance = NomadClient::EvaluationsApi.new
eval_id = 'eval_id_example' # String | Evaluation ID.
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
  
  result = api_instance.get_evaluation_allocations(eval_id, opts)
  p result
rescue NomadClient::ApiError => e
  puts "Error when calling EvaluationsApi->get_evaluation_allocations: #{e}"
end
```

#### Using the get_evaluation_allocations_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Array<AllocationListStub>>, Integer, Hash)> get_evaluation_allocations_with_http_info(eval_id, opts)

```ruby
begin
  
  data, status_code, headers = api_instance.get_evaluation_allocations_with_http_info(eval_id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Array<AllocationListStub>>
rescue NomadClient::ApiError => e
  puts "Error when calling EvaluationsApi->get_evaluation_allocations_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **eval_id** | **String** | Evaluation ID. |  |
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


## get_evaluations

> <Array<Evaluation>> get_evaluations(opts)



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

api_instance = NomadClient::EvaluationsApi.new
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
  
  result = api_instance.get_evaluations(opts)
  p result
rescue NomadClient::ApiError => e
  puts "Error when calling EvaluationsApi->get_evaluations: #{e}"
end
```

#### Using the get_evaluations_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Array<Evaluation>>, Integer, Hash)> get_evaluations_with_http_info(opts)

```ruby
begin
  
  data, status_code, headers = api_instance.get_evaluations_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Array<Evaluation>>
rescue NomadClient::ApiError => e
  puts "Error when calling EvaluationsApi->get_evaluations_with_http_info: #{e}"
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

[**Array&lt;Evaluation&gt;**](Evaluation.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

