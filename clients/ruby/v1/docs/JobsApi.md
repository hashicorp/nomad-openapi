# NomadClient::JobsApi

All URIs are relative to *http://127.0.0.1:4646/v1*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**delete_job**](JobsApi.md#delete_job) | **DELETE** /job/{jobName} |  |
| [**get_job**](JobsApi.md#get_job) | **GET** /job/{jobName} |  |
| [**get_job_allocations**](JobsApi.md#get_job_allocations) | **GET** /job/{jobName}/allocations |  |
| [**get_job_deployment**](JobsApi.md#get_job_deployment) | **GET** /job/{jobName}/deployment |  |
| [**get_job_deployments**](JobsApi.md#get_job_deployments) | **GET** /job/{jobName}/deployments |  |
| [**get_job_evaluations**](JobsApi.md#get_job_evaluations) | **GET** /job/{jobName}/evaluations |  |
| [**get_job_scale_status**](JobsApi.md#get_job_scale_status) | **GET** /job/{jobName}/scale |  |
| [**get_job_summary**](JobsApi.md#get_job_summary) | **GET** /job/{jobName}/summary |  |
| [**get_job_versions**](JobsApi.md#get_job_versions) | **GET** /job/{jobName}/versions |  |
| [**get_jobs**](JobsApi.md#get_jobs) | **GET** /jobs |  |
| [**post_job**](JobsApi.md#post_job) | **POST** /job/{jobName} |  |
| [**post_job_dispatch**](JobsApi.md#post_job_dispatch) | **POST** /job/{jobName}/dispatch |  |
| [**post_job_evaluate**](JobsApi.md#post_job_evaluate) | **POST** /job/{jobName}/evaluate |  |
| [**post_job_parse**](JobsApi.md#post_job_parse) | **POST** /jobs/parse |  |
| [**post_job_periodic_force**](JobsApi.md#post_job_periodic_force) | **POST** /job/{jobName}/periodic/force |  |
| [**post_job_plan**](JobsApi.md#post_job_plan) | **POST** /job/{jobName}/plan |  |
| [**post_job_revert**](JobsApi.md#post_job_revert) | **POST** /job/{jobName}/revert |  |
| [**post_job_scaling_request**](JobsApi.md#post_job_scaling_request) | **POST** /job/{jobName}/scale |  |
| [**post_job_stability**](JobsApi.md#post_job_stability) | **POST** /job/{jobName}/stable |  |
| [**post_job_validate_request**](JobsApi.md#post_job_validate_request) | **POST** /validate/job |  |
| [**register_job**](JobsApi.md#register_job) | **POST** /jobs |  |


## delete_job

> <JobDeregisterResponse> delete_job(job_name, opts)



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

api_instance = NomadClient::JobsApi.new
job_name = 'job_name_example' # String | The job identifier.
opts = {
  region: 'region_example', # String | Filters results based on the specified region.
  namespace: 'namespace_example', # String | Filters results based on the specified namespace.
  x_nomad_token: 'x_nomad_token_example', # String | A Nomad ACL token.
  idempotency_token: 'idempotency_token_example', # String | Can be used to ensure operations are only run once.
  purge: true, # Boolean | Boolean flag indicating whether to purge allocations of the job after deleting.
  global: true # Boolean | Boolean flag indicating whether the operation should apply to all instances of the job globally.
}

begin
  
  result = api_instance.delete_job(job_name, opts)
  p result
rescue NomadClient::ApiError => e
  puts "Error when calling JobsApi->delete_job: #{e}"
end
```

#### Using the delete_job_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<JobDeregisterResponse>, Integer, Hash)> delete_job_with_http_info(job_name, opts)

```ruby
begin
  
  data, status_code, headers = api_instance.delete_job_with_http_info(job_name, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <JobDeregisterResponse>
rescue NomadClient::ApiError => e
  puts "Error when calling JobsApi->delete_job_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **job_name** | **String** | The job identifier. |  |
| **region** | **String** | Filters results based on the specified region. | [optional] |
| **namespace** | **String** | Filters results based on the specified namespace. | [optional] |
| **x_nomad_token** | **String** | A Nomad ACL token. | [optional] |
| **idempotency_token** | **String** | Can be used to ensure operations are only run once. | [optional] |
| **purge** | **Boolean** | Boolean flag indicating whether to purge allocations of the job after deleting. | [optional] |
| **global** | **Boolean** | Boolean flag indicating whether the operation should apply to all instances of the job globally. | [optional] |

### Return type

[**JobDeregisterResponse**](JobDeregisterResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## get_job

> <Job> get_job(job_name, opts)



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

api_instance = NomadClient::JobsApi.new
job_name = 'job_name_example' # String | The job identifier.
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
  
  result = api_instance.get_job(job_name, opts)
  p result
rescue NomadClient::ApiError => e
  puts "Error when calling JobsApi->get_job: #{e}"
end
```

#### Using the get_job_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Job>, Integer, Hash)> get_job_with_http_info(job_name, opts)

```ruby
begin
  
  data, status_code, headers = api_instance.get_job_with_http_info(job_name, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Job>
rescue NomadClient::ApiError => e
  puts "Error when calling JobsApi->get_job_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **job_name** | **String** | The job identifier. |  |
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

[**Job**](Job.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## get_job_allocations

> <Array<AllocationListStub>> get_job_allocations(job_name, opts)



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

api_instance = NomadClient::JobsApi.new
job_name = 'job_name_example' # String | The job identifier.
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
  all: true # Boolean | Specifies whether the list of allocations should include allocations from a previously registered job with the same ID. This is possible if the job is deregistered and reregistered.
}

begin
  
  result = api_instance.get_job_allocations(job_name, opts)
  p result
rescue NomadClient::ApiError => e
  puts "Error when calling JobsApi->get_job_allocations: #{e}"
end
```

#### Using the get_job_allocations_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Array<AllocationListStub>>, Integer, Hash)> get_job_allocations_with_http_info(job_name, opts)

```ruby
begin
  
  data, status_code, headers = api_instance.get_job_allocations_with_http_info(job_name, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Array<AllocationListStub>>
rescue NomadClient::ApiError => e
  puts "Error when calling JobsApi->get_job_allocations_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **job_name** | **String** | The job identifier. |  |
| **region** | **String** | Filters results based on the specified region. | [optional] |
| **namespace** | **String** | Filters results based on the specified namespace. | [optional] |
| **index** | **Integer** | If set, wait until query exceeds given index. Must be provided with WaitParam. | [optional] |
| **wait** | **String** | Provided with IndexParam to wait for change. | [optional] |
| **stale** | **String** | If present, results will include stale reads. | [optional] |
| **prefix** | **String** | Constrains results to jobs that start with the defined prefix | [optional] |
| **x_nomad_token** | **String** | A Nomad ACL token. | [optional] |
| **per_page** | **Integer** | Maximum number of results to return. | [optional] |
| **next_token** | **String** | Indicates where to start paging for queries that support pagination. | [optional] |
| **all** | **Boolean** | Specifies whether the list of allocations should include allocations from a previously registered job with the same ID. This is possible if the job is deregistered and reregistered. | [optional] |

### Return type

[**Array&lt;AllocationListStub&gt;**](AllocationListStub.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## get_job_deployment

> <Deployment> get_job_deployment(job_name, opts)



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

api_instance = NomadClient::JobsApi.new
job_name = 'job_name_example' # String | The job identifier.
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
  
  result = api_instance.get_job_deployment(job_name, opts)
  p result
rescue NomadClient::ApiError => e
  puts "Error when calling JobsApi->get_job_deployment: #{e}"
end
```

#### Using the get_job_deployment_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Deployment>, Integer, Hash)> get_job_deployment_with_http_info(job_name, opts)

```ruby
begin
  
  data, status_code, headers = api_instance.get_job_deployment_with_http_info(job_name, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Deployment>
rescue NomadClient::ApiError => e
  puts "Error when calling JobsApi->get_job_deployment_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **job_name** | **String** | The job identifier. |  |
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


## get_job_deployments

> <Array<Deployment>> get_job_deployments(job_name, opts)



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

api_instance = NomadClient::JobsApi.new
job_name = 'job_name_example' # String | The job identifier.
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
  all: 56 # Integer | Flag indicating whether to constrain by job creation index or not.
}

begin
  
  result = api_instance.get_job_deployments(job_name, opts)
  p result
rescue NomadClient::ApiError => e
  puts "Error when calling JobsApi->get_job_deployments: #{e}"
end
```

#### Using the get_job_deployments_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Array<Deployment>>, Integer, Hash)> get_job_deployments_with_http_info(job_name, opts)

```ruby
begin
  
  data, status_code, headers = api_instance.get_job_deployments_with_http_info(job_name, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Array<Deployment>>
rescue NomadClient::ApiError => e
  puts "Error when calling JobsApi->get_job_deployments_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **job_name** | **String** | The job identifier. |  |
| **region** | **String** | Filters results based on the specified region. | [optional] |
| **namespace** | **String** | Filters results based on the specified namespace. | [optional] |
| **index** | **Integer** | If set, wait until query exceeds given index. Must be provided with WaitParam. | [optional] |
| **wait** | **String** | Provided with IndexParam to wait for change. | [optional] |
| **stale** | **String** | If present, results will include stale reads. | [optional] |
| **prefix** | **String** | Constrains results to jobs that start with the defined prefix | [optional] |
| **x_nomad_token** | **String** | A Nomad ACL token. | [optional] |
| **per_page** | **Integer** | Maximum number of results to return. | [optional] |
| **next_token** | **String** | Indicates where to start paging for queries that support pagination. | [optional] |
| **all** | **Integer** | Flag indicating whether to constrain by job creation index or not. | [optional] |

### Return type

[**Array&lt;Deployment&gt;**](Deployment.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## get_job_evaluations

> <Array<Evaluation>> get_job_evaluations(job_name, opts)



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

api_instance = NomadClient::JobsApi.new
job_name = 'job_name_example' # String | The job identifier.
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
  
  result = api_instance.get_job_evaluations(job_name, opts)
  p result
rescue NomadClient::ApiError => e
  puts "Error when calling JobsApi->get_job_evaluations: #{e}"
end
```

#### Using the get_job_evaluations_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Array<Evaluation>>, Integer, Hash)> get_job_evaluations_with_http_info(job_name, opts)

```ruby
begin
  
  data, status_code, headers = api_instance.get_job_evaluations_with_http_info(job_name, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Array<Evaluation>>
rescue NomadClient::ApiError => e
  puts "Error when calling JobsApi->get_job_evaluations_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **job_name** | **String** | The job identifier. |  |
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


## get_job_scale_status

> <JobScaleStatusResponse> get_job_scale_status(job_name, opts)



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

api_instance = NomadClient::JobsApi.new
job_name = 'job_name_example' # String | The job identifier.
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
  
  result = api_instance.get_job_scale_status(job_name, opts)
  p result
rescue NomadClient::ApiError => e
  puts "Error when calling JobsApi->get_job_scale_status: #{e}"
end
```

#### Using the get_job_scale_status_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<JobScaleStatusResponse>, Integer, Hash)> get_job_scale_status_with_http_info(job_name, opts)

```ruby
begin
  
  data, status_code, headers = api_instance.get_job_scale_status_with_http_info(job_name, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <JobScaleStatusResponse>
rescue NomadClient::ApiError => e
  puts "Error when calling JobsApi->get_job_scale_status_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **job_name** | **String** | The job identifier. |  |
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

[**JobScaleStatusResponse**](JobScaleStatusResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## get_job_summary

> <JobSummary> get_job_summary(job_name, opts)



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

api_instance = NomadClient::JobsApi.new
job_name = 'job_name_example' # String | The job identifier.
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
  
  result = api_instance.get_job_summary(job_name, opts)
  p result
rescue NomadClient::ApiError => e
  puts "Error when calling JobsApi->get_job_summary: #{e}"
end
```

#### Using the get_job_summary_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<JobSummary>, Integer, Hash)> get_job_summary_with_http_info(job_name, opts)

```ruby
begin
  
  data, status_code, headers = api_instance.get_job_summary_with_http_info(job_name, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <JobSummary>
rescue NomadClient::ApiError => e
  puts "Error when calling JobsApi->get_job_summary_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **job_name** | **String** | The job identifier. |  |
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

[**JobSummary**](JobSummary.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## get_job_versions

> <JobVersionsResponse> get_job_versions(job_name, opts)



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

api_instance = NomadClient::JobsApi.new
job_name = 'job_name_example' # String | The job identifier.
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
  diffs: true # Boolean | Boolean flag indicating whether to compute job diffs.
}

begin
  
  result = api_instance.get_job_versions(job_name, opts)
  p result
rescue NomadClient::ApiError => e
  puts "Error when calling JobsApi->get_job_versions: #{e}"
end
```

#### Using the get_job_versions_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<JobVersionsResponse>, Integer, Hash)> get_job_versions_with_http_info(job_name, opts)

```ruby
begin
  
  data, status_code, headers = api_instance.get_job_versions_with_http_info(job_name, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <JobVersionsResponse>
rescue NomadClient::ApiError => e
  puts "Error when calling JobsApi->get_job_versions_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **job_name** | **String** | The job identifier. |  |
| **region** | **String** | Filters results based on the specified region. | [optional] |
| **namespace** | **String** | Filters results based on the specified namespace. | [optional] |
| **index** | **Integer** | If set, wait until query exceeds given index. Must be provided with WaitParam. | [optional] |
| **wait** | **String** | Provided with IndexParam to wait for change. | [optional] |
| **stale** | **String** | If present, results will include stale reads. | [optional] |
| **prefix** | **String** | Constrains results to jobs that start with the defined prefix | [optional] |
| **x_nomad_token** | **String** | A Nomad ACL token. | [optional] |
| **per_page** | **Integer** | Maximum number of results to return. | [optional] |
| **next_token** | **String** | Indicates where to start paging for queries that support pagination. | [optional] |
| **diffs** | **Boolean** | Boolean flag indicating whether to compute job diffs. | [optional] |

### Return type

[**JobVersionsResponse**](JobVersionsResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## get_jobs

> <Array<JobListStub>> get_jobs(opts)



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

api_instance = NomadClient::JobsApi.new
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
  
  result = api_instance.get_jobs(opts)
  p result
rescue NomadClient::ApiError => e
  puts "Error when calling JobsApi->get_jobs: #{e}"
end
```

#### Using the get_jobs_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Array<JobListStub>>, Integer, Hash)> get_jobs_with_http_info(opts)

```ruby
begin
  
  data, status_code, headers = api_instance.get_jobs_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Array<JobListStub>>
rescue NomadClient::ApiError => e
  puts "Error when calling JobsApi->get_jobs_with_http_info: #{e}"
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

[**Array&lt;JobListStub&gt;**](JobListStub.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## post_job

> <JobRegisterResponse> post_job(job_name, job_register_request, opts)



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

api_instance = NomadClient::JobsApi.new
job_name = 'job_name_example' # String | The job identifier.
job_register_request = NomadClient::JobRegisterRequest.new # JobRegisterRequest | 
opts = {
  region: 'region_example', # String | Filters results based on the specified region.
  namespace: 'namespace_example', # String | Filters results based on the specified namespace.
  x_nomad_token: 'x_nomad_token_example', # String | A Nomad ACL token.
  idempotency_token: 'idempotency_token_example' # String | Can be used to ensure operations are only run once.
}

begin
  
  result = api_instance.post_job(job_name, job_register_request, opts)
  p result
rescue NomadClient::ApiError => e
  puts "Error when calling JobsApi->post_job: #{e}"
end
```

#### Using the post_job_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<JobRegisterResponse>, Integer, Hash)> post_job_with_http_info(job_name, job_register_request, opts)

```ruby
begin
  
  data, status_code, headers = api_instance.post_job_with_http_info(job_name, job_register_request, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <JobRegisterResponse>
rescue NomadClient::ApiError => e
  puts "Error when calling JobsApi->post_job_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **job_name** | **String** | The job identifier. |  |
| **job_register_request** | [**JobRegisterRequest**](JobRegisterRequest.md) |  |  |
| **region** | **String** | Filters results based on the specified region. | [optional] |
| **namespace** | **String** | Filters results based on the specified namespace. | [optional] |
| **x_nomad_token** | **String** | A Nomad ACL token. | [optional] |
| **idempotency_token** | **String** | Can be used to ensure operations are only run once. | [optional] |

### Return type

[**JobRegisterResponse**](JobRegisterResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## post_job_dispatch

> <JobDispatchResponse> post_job_dispatch(job_name, job_dispatch_request, opts)



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

api_instance = NomadClient::JobsApi.new
job_name = 'job_name_example' # String | The job identifier.
job_dispatch_request = NomadClient::JobDispatchRequest.new # JobDispatchRequest | 
opts = {
  region: 'region_example', # String | Filters results based on the specified region.
  namespace: 'namespace_example', # String | Filters results based on the specified namespace.
  x_nomad_token: 'x_nomad_token_example', # String | A Nomad ACL token.
  idempotency_token: 'idempotency_token_example' # String | Can be used to ensure operations are only run once.
}

begin
  
  result = api_instance.post_job_dispatch(job_name, job_dispatch_request, opts)
  p result
rescue NomadClient::ApiError => e
  puts "Error when calling JobsApi->post_job_dispatch: #{e}"
end
```

#### Using the post_job_dispatch_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<JobDispatchResponse>, Integer, Hash)> post_job_dispatch_with_http_info(job_name, job_dispatch_request, opts)

```ruby
begin
  
  data, status_code, headers = api_instance.post_job_dispatch_with_http_info(job_name, job_dispatch_request, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <JobDispatchResponse>
rescue NomadClient::ApiError => e
  puts "Error when calling JobsApi->post_job_dispatch_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **job_name** | **String** | The job identifier. |  |
| **job_dispatch_request** | [**JobDispatchRequest**](JobDispatchRequest.md) |  |  |
| **region** | **String** | Filters results based on the specified region. | [optional] |
| **namespace** | **String** | Filters results based on the specified namespace. | [optional] |
| **x_nomad_token** | **String** | A Nomad ACL token. | [optional] |
| **idempotency_token** | **String** | Can be used to ensure operations are only run once. | [optional] |

### Return type

[**JobDispatchResponse**](JobDispatchResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## post_job_evaluate

> <JobRegisterResponse> post_job_evaluate(job_name, job_evaluate_request, opts)



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

api_instance = NomadClient::JobsApi.new
job_name = 'job_name_example' # String | The job identifier.
job_evaluate_request = NomadClient::JobEvaluateRequest.new # JobEvaluateRequest | 
opts = {
  region: 'region_example', # String | Filters results based on the specified region.
  namespace: 'namespace_example', # String | Filters results based on the specified namespace.
  x_nomad_token: 'x_nomad_token_example', # String | A Nomad ACL token.
  idempotency_token: 'idempotency_token_example' # String | Can be used to ensure operations are only run once.
}

begin
  
  result = api_instance.post_job_evaluate(job_name, job_evaluate_request, opts)
  p result
rescue NomadClient::ApiError => e
  puts "Error when calling JobsApi->post_job_evaluate: #{e}"
end
```

#### Using the post_job_evaluate_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<JobRegisterResponse>, Integer, Hash)> post_job_evaluate_with_http_info(job_name, job_evaluate_request, opts)

```ruby
begin
  
  data, status_code, headers = api_instance.post_job_evaluate_with_http_info(job_name, job_evaluate_request, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <JobRegisterResponse>
rescue NomadClient::ApiError => e
  puts "Error when calling JobsApi->post_job_evaluate_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **job_name** | **String** | The job identifier. |  |
| **job_evaluate_request** | [**JobEvaluateRequest**](JobEvaluateRequest.md) |  |  |
| **region** | **String** | Filters results based on the specified region. | [optional] |
| **namespace** | **String** | Filters results based on the specified namespace. | [optional] |
| **x_nomad_token** | **String** | A Nomad ACL token. | [optional] |
| **idempotency_token** | **String** | Can be used to ensure operations are only run once. | [optional] |

### Return type

[**JobRegisterResponse**](JobRegisterResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## post_job_parse

> <Job> post_job_parse(jobs_parse_request)



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

api_instance = NomadClient::JobsApi.new
jobs_parse_request = NomadClient::JobsParseRequest.new # JobsParseRequest | 

begin
  
  result = api_instance.post_job_parse(jobs_parse_request)
  p result
rescue NomadClient::ApiError => e
  puts "Error when calling JobsApi->post_job_parse: #{e}"
end
```

#### Using the post_job_parse_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Job>, Integer, Hash)> post_job_parse_with_http_info(jobs_parse_request)

```ruby
begin
  
  data, status_code, headers = api_instance.post_job_parse_with_http_info(jobs_parse_request)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Job>
rescue NomadClient::ApiError => e
  puts "Error when calling JobsApi->post_job_parse_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **jobs_parse_request** | [**JobsParseRequest**](JobsParseRequest.md) |  |  |

### Return type

[**Job**](Job.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## post_job_periodic_force

> <PeriodicForceResponse> post_job_periodic_force(job_name, opts)



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

api_instance = NomadClient::JobsApi.new
job_name = 'job_name_example' # String | The job identifier.
opts = {
  region: 'region_example', # String | Filters results based on the specified region.
  namespace: 'namespace_example', # String | Filters results based on the specified namespace.
  x_nomad_token: 'x_nomad_token_example', # String | A Nomad ACL token.
  idempotency_token: 'idempotency_token_example' # String | Can be used to ensure operations are only run once.
}

begin
  
  result = api_instance.post_job_periodic_force(job_name, opts)
  p result
rescue NomadClient::ApiError => e
  puts "Error when calling JobsApi->post_job_periodic_force: #{e}"
end
```

#### Using the post_job_periodic_force_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<PeriodicForceResponse>, Integer, Hash)> post_job_periodic_force_with_http_info(job_name, opts)

```ruby
begin
  
  data, status_code, headers = api_instance.post_job_periodic_force_with_http_info(job_name, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <PeriodicForceResponse>
rescue NomadClient::ApiError => e
  puts "Error when calling JobsApi->post_job_periodic_force_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **job_name** | **String** | The job identifier. |  |
| **region** | **String** | Filters results based on the specified region. | [optional] |
| **namespace** | **String** | Filters results based on the specified namespace. | [optional] |
| **x_nomad_token** | **String** | A Nomad ACL token. | [optional] |
| **idempotency_token** | **String** | Can be used to ensure operations are only run once. | [optional] |

### Return type

[**PeriodicForceResponse**](PeriodicForceResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## post_job_plan

> <JobPlanResponse> post_job_plan(job_name, job_plan_request, opts)



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

api_instance = NomadClient::JobsApi.new
job_name = 'job_name_example' # String | The job identifier.
job_plan_request = NomadClient::JobPlanRequest.new # JobPlanRequest | 
opts = {
  region: 'region_example', # String | Filters results based on the specified region.
  namespace: 'namespace_example', # String | Filters results based on the specified namespace.
  x_nomad_token: 'x_nomad_token_example', # String | A Nomad ACL token.
  idempotency_token: 'idempotency_token_example' # String | Can be used to ensure operations are only run once.
}

begin
  
  result = api_instance.post_job_plan(job_name, job_plan_request, opts)
  p result
rescue NomadClient::ApiError => e
  puts "Error when calling JobsApi->post_job_plan: #{e}"
end
```

#### Using the post_job_plan_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<JobPlanResponse>, Integer, Hash)> post_job_plan_with_http_info(job_name, job_plan_request, opts)

```ruby
begin
  
  data, status_code, headers = api_instance.post_job_plan_with_http_info(job_name, job_plan_request, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <JobPlanResponse>
rescue NomadClient::ApiError => e
  puts "Error when calling JobsApi->post_job_plan_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **job_name** | **String** | The job identifier. |  |
| **job_plan_request** | [**JobPlanRequest**](JobPlanRequest.md) |  |  |
| **region** | **String** | Filters results based on the specified region. | [optional] |
| **namespace** | **String** | Filters results based on the specified namespace. | [optional] |
| **x_nomad_token** | **String** | A Nomad ACL token. | [optional] |
| **idempotency_token** | **String** | Can be used to ensure operations are only run once. | [optional] |

### Return type

[**JobPlanResponse**](JobPlanResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## post_job_revert

> <JobRegisterResponse> post_job_revert(job_name, job_revert_request, opts)



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

api_instance = NomadClient::JobsApi.new
job_name = 'job_name_example' # String | The job identifier.
job_revert_request = NomadClient::JobRevertRequest.new # JobRevertRequest | 
opts = {
  region: 'region_example', # String | Filters results based on the specified region.
  namespace: 'namespace_example', # String | Filters results based on the specified namespace.
  x_nomad_token: 'x_nomad_token_example', # String | A Nomad ACL token.
  idempotency_token: 'idempotency_token_example' # String | Can be used to ensure operations are only run once.
}

begin
  
  result = api_instance.post_job_revert(job_name, job_revert_request, opts)
  p result
rescue NomadClient::ApiError => e
  puts "Error when calling JobsApi->post_job_revert: #{e}"
end
```

#### Using the post_job_revert_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<JobRegisterResponse>, Integer, Hash)> post_job_revert_with_http_info(job_name, job_revert_request, opts)

```ruby
begin
  
  data, status_code, headers = api_instance.post_job_revert_with_http_info(job_name, job_revert_request, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <JobRegisterResponse>
rescue NomadClient::ApiError => e
  puts "Error when calling JobsApi->post_job_revert_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **job_name** | **String** | The job identifier. |  |
| **job_revert_request** | [**JobRevertRequest**](JobRevertRequest.md) |  |  |
| **region** | **String** | Filters results based on the specified region. | [optional] |
| **namespace** | **String** | Filters results based on the specified namespace. | [optional] |
| **x_nomad_token** | **String** | A Nomad ACL token. | [optional] |
| **idempotency_token** | **String** | Can be used to ensure operations are only run once. | [optional] |

### Return type

[**JobRegisterResponse**](JobRegisterResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## post_job_scaling_request

> <JobRegisterResponse> post_job_scaling_request(job_name, scaling_request, opts)



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

api_instance = NomadClient::JobsApi.new
job_name = 'job_name_example' # String | The job identifier.
scaling_request = NomadClient::ScalingRequest.new # ScalingRequest | 
opts = {
  region: 'region_example', # String | Filters results based on the specified region.
  namespace: 'namespace_example', # String | Filters results based on the specified namespace.
  x_nomad_token: 'x_nomad_token_example', # String | A Nomad ACL token.
  idempotency_token: 'idempotency_token_example' # String | Can be used to ensure operations are only run once.
}

begin
  
  result = api_instance.post_job_scaling_request(job_name, scaling_request, opts)
  p result
rescue NomadClient::ApiError => e
  puts "Error when calling JobsApi->post_job_scaling_request: #{e}"
end
```

#### Using the post_job_scaling_request_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<JobRegisterResponse>, Integer, Hash)> post_job_scaling_request_with_http_info(job_name, scaling_request, opts)

```ruby
begin
  
  data, status_code, headers = api_instance.post_job_scaling_request_with_http_info(job_name, scaling_request, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <JobRegisterResponse>
rescue NomadClient::ApiError => e
  puts "Error when calling JobsApi->post_job_scaling_request_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **job_name** | **String** | The job identifier. |  |
| **scaling_request** | [**ScalingRequest**](ScalingRequest.md) |  |  |
| **region** | **String** | Filters results based on the specified region. | [optional] |
| **namespace** | **String** | Filters results based on the specified namespace. | [optional] |
| **x_nomad_token** | **String** | A Nomad ACL token. | [optional] |
| **idempotency_token** | **String** | Can be used to ensure operations are only run once. | [optional] |

### Return type

[**JobRegisterResponse**](JobRegisterResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## post_job_stability

> <JobStabilityResponse> post_job_stability(job_name, job_stability_request, opts)



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

api_instance = NomadClient::JobsApi.new
job_name = 'job_name_example' # String | The job identifier.
job_stability_request = NomadClient::JobStabilityRequest.new # JobStabilityRequest | 
opts = {
  region: 'region_example', # String | Filters results based on the specified region.
  namespace: 'namespace_example', # String | Filters results based on the specified namespace.
  x_nomad_token: 'x_nomad_token_example', # String | A Nomad ACL token.
  idempotency_token: 'idempotency_token_example' # String | Can be used to ensure operations are only run once.
}

begin
  
  result = api_instance.post_job_stability(job_name, job_stability_request, opts)
  p result
rescue NomadClient::ApiError => e
  puts "Error when calling JobsApi->post_job_stability: #{e}"
end
```

#### Using the post_job_stability_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<JobStabilityResponse>, Integer, Hash)> post_job_stability_with_http_info(job_name, job_stability_request, opts)

```ruby
begin
  
  data, status_code, headers = api_instance.post_job_stability_with_http_info(job_name, job_stability_request, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <JobStabilityResponse>
rescue NomadClient::ApiError => e
  puts "Error when calling JobsApi->post_job_stability_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **job_name** | **String** | The job identifier. |  |
| **job_stability_request** | [**JobStabilityRequest**](JobStabilityRequest.md) |  |  |
| **region** | **String** | Filters results based on the specified region. | [optional] |
| **namespace** | **String** | Filters results based on the specified namespace. | [optional] |
| **x_nomad_token** | **String** | A Nomad ACL token. | [optional] |
| **idempotency_token** | **String** | Can be used to ensure operations are only run once. | [optional] |

### Return type

[**JobStabilityResponse**](JobStabilityResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## post_job_validate_request

> <JobValidateResponse> post_job_validate_request(job_validate_request, opts)



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

api_instance = NomadClient::JobsApi.new
job_validate_request = NomadClient::JobValidateRequest.new # JobValidateRequest | 
opts = {
  region: 'region_example', # String | Filters results based on the specified region.
  namespace: 'namespace_example', # String | Filters results based on the specified namespace.
  x_nomad_token: 'x_nomad_token_example', # String | A Nomad ACL token.
  idempotency_token: 'idempotency_token_example' # String | Can be used to ensure operations are only run once.
}

begin
  
  result = api_instance.post_job_validate_request(job_validate_request, opts)
  p result
rescue NomadClient::ApiError => e
  puts "Error when calling JobsApi->post_job_validate_request: #{e}"
end
```

#### Using the post_job_validate_request_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<JobValidateResponse>, Integer, Hash)> post_job_validate_request_with_http_info(job_validate_request, opts)

```ruby
begin
  
  data, status_code, headers = api_instance.post_job_validate_request_with_http_info(job_validate_request, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <JobValidateResponse>
rescue NomadClient::ApiError => e
  puts "Error when calling JobsApi->post_job_validate_request_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **job_validate_request** | [**JobValidateRequest**](JobValidateRequest.md) |  |  |
| **region** | **String** | Filters results based on the specified region. | [optional] |
| **namespace** | **String** | Filters results based on the specified namespace. | [optional] |
| **x_nomad_token** | **String** | A Nomad ACL token. | [optional] |
| **idempotency_token** | **String** | Can be used to ensure operations are only run once. | [optional] |

### Return type

[**JobValidateResponse**](JobValidateResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## register_job

> <JobRegisterResponse> register_job(job_register_request, opts)



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

api_instance = NomadClient::JobsApi.new
job_register_request = NomadClient::JobRegisterRequest.new # JobRegisterRequest | 
opts = {
  region: 'region_example', # String | Filters results based on the specified region.
  namespace: 'namespace_example', # String | Filters results based on the specified namespace.
  x_nomad_token: 'x_nomad_token_example', # String | A Nomad ACL token.
  idempotency_token: 'idempotency_token_example' # String | Can be used to ensure operations are only run once.
}

begin
  
  result = api_instance.register_job(job_register_request, opts)
  p result
rescue NomadClient::ApiError => e
  puts "Error when calling JobsApi->register_job: #{e}"
end
```

#### Using the register_job_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<JobRegisterResponse>, Integer, Hash)> register_job_with_http_info(job_register_request, opts)

```ruby
begin
  
  data, status_code, headers = api_instance.register_job_with_http_info(job_register_request, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <JobRegisterResponse>
rescue NomadClient::ApiError => e
  puts "Error when calling JobsApi->register_job_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **job_register_request** | [**JobRegisterRequest**](JobRegisterRequest.md) |  |  |
| **region** | **String** | Filters results based on the specified region. | [optional] |
| **namespace** | **String** | Filters results based on the specified namespace. | [optional] |
| **x_nomad_token** | **String** | A Nomad ACL token. | [optional] |
| **idempotency_token** | **String** | Can be used to ensure operations are only run once. | [optional] |

### Return type

[**JobRegisterResponse**](JobRegisterResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

