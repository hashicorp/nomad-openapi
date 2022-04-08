# nomad_client.JobsApi

All URIs are relative to *https://127.0.0.1:4646/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**delete_job**](JobsApi.md#delete_job) | **DELETE** /job/{jobName} | 
[**get_job**](JobsApi.md#get_job) | **GET** /job/{jobName} | 
[**get_job_allocations**](JobsApi.md#get_job_allocations) | **GET** /job/{jobName}/allocations | 
[**get_job_deployment**](JobsApi.md#get_job_deployment) | **GET** /job/{jobName}/deployment | 
[**get_job_deployments**](JobsApi.md#get_job_deployments) | **GET** /job/{jobName}/deployments | 
[**get_job_evaluations**](JobsApi.md#get_job_evaluations) | **GET** /job/{jobName}/evaluations | 
[**get_job_scale_status**](JobsApi.md#get_job_scale_status) | **GET** /job/{jobName}/scale | 
[**get_job_summary**](JobsApi.md#get_job_summary) | **GET** /job/{jobName}/summary | 
[**get_job_versions**](JobsApi.md#get_job_versions) | **GET** /job/{jobName}/versions | 
[**get_jobs**](JobsApi.md#get_jobs) | **GET** /jobs | 
[**post_job**](JobsApi.md#post_job) | **POST** /job/{jobName} | 
[**post_job_dispatch**](JobsApi.md#post_job_dispatch) | **POST** /job/{jobName}/dispatch | 
[**post_job_evaluate**](JobsApi.md#post_job_evaluate) | **POST** /job/{jobName}/evaluate | 
[**post_job_parse**](JobsApi.md#post_job_parse) | **POST** /jobs/parse | 
[**post_job_periodic_force**](JobsApi.md#post_job_periodic_force) | **POST** /job/{jobName}/periodic/force | 
[**post_job_plan**](JobsApi.md#post_job_plan) | **POST** /job/{jobName}/plan | 
[**post_job_revert**](JobsApi.md#post_job_revert) | **POST** /job/{jobName}/revert | 
[**post_job_scaling_request**](JobsApi.md#post_job_scaling_request) | **POST** /job/{jobName}/scale | 
[**post_job_stability**](JobsApi.md#post_job_stability) | **POST** /job/{jobName}/stable | 
[**post_job_validate_request**](JobsApi.md#post_job_validate_request) | **POST** /validate/job | 
[**register_job**](JobsApi.md#register_job) | **POST** /jobs | 


# **delete_job**
> JobDeregisterResponse delete_job(job_name)



### Example

* Api Key Authentication (X-Nomad-Token):
```python
import time
import nomad_client
from nomad_client.api import jobs_api
from nomad_client.model.job_deregister_response import JobDeregisterResponse
from pprint import pprint
# Defining the host is optional and defaults to https://127.0.0.1:4646/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = nomad_client.Configuration(
    host = "https://127.0.0.1:4646/v1"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: X-Nomad-Token
configuration.api_key['X-Nomad-Token'] = 'YOUR_API_KEY'

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['X-Nomad-Token'] = 'Bearer'

# Enter a context with an instance of the API client
with nomad_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = jobs_api.JobsApi(api_client)
    job_name = "jobName_example" # str | The job identifier.
    region = "region_example" # str | Filters results based on the specified region. (optional)
    namespace = "namespace_example" # str | Filters results based on the specified namespace. (optional)
    x_nomad_token = "X-Nomad-Token_example" # str | A Nomad ACL token. (optional)
    idempotency_token = "idempotency_token_example" # str | Can be used to ensure operations are only run once. (optional)
    purge = True # bool | Boolean flag indicating whether to purge allocations of the job after deleting. (optional)
    _global = True # bool | Boolean flag indicating whether the operation should apply to all instances of the job globally. (optional)

    # example passing only required values which don't have defaults set
    try:
        api_response = api_instance.delete_job(job_name)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling JobsApi->delete_job: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        api_response = api_instance.delete_job(job_name, region=region, namespace=namespace, x_nomad_token=x_nomad_token, idempotency_token=idempotency_token, purge=purge, _global=_global)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling JobsApi->delete_job: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **job_name** | **str**| The job identifier. |
 **region** | **str**| Filters results based on the specified region. | [optional]
 **namespace** | **str**| Filters results based on the specified namespace. | [optional]
 **x_nomad_token** | **str**| A Nomad ACL token. | [optional]
 **idempotency_token** | **str**| Can be used to ensure operations are only run once. | [optional]
 **purge** | **bool**| Boolean flag indicating whether to purge allocations of the job after deleting. | [optional]
 **_global** | **bool**| Boolean flag indicating whether the operation should apply to all instances of the job globally. | [optional]

### Return type

[**JobDeregisterResponse**](JobDeregisterResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** |  |  * X-Nomad-Index - A unique identifier representing the current state of the requested resource. On a new Nomad cluster the value of this index starts at 1. <br>  |
**400** | Bad request |  -  |
**403** | Forbidden |  -  |
**405** | Method not allowed |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_job**
> Job get_job(job_name)



### Example

* Api Key Authentication (X-Nomad-Token):
```python
import time
import nomad_client
from nomad_client.api import jobs_api
from nomad_client.model.job import Job
from pprint import pprint
# Defining the host is optional and defaults to https://127.0.0.1:4646/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = nomad_client.Configuration(
    host = "https://127.0.0.1:4646/v1"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: X-Nomad-Token
configuration.api_key['X-Nomad-Token'] = 'YOUR_API_KEY'

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['X-Nomad-Token'] = 'Bearer'

# Enter a context with an instance of the API client
with nomad_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = jobs_api.JobsApi(api_client)
    job_name = "jobName_example" # str | The job identifier.
    region = "region_example" # str | Filters results based on the specified region. (optional)
    namespace = "namespace_example" # str | Filters results based on the specified namespace. (optional)
    index = 1 # int | If set, wait until query exceeds given index. Must be provided with WaitParam. (optional)
    wait = "wait_example" # str | Provided with IndexParam to wait for change. (optional)
    stale = "stale_example" # str | If present, results will include stale reads. (optional)
    prefix = "prefix_example" # str | Constrains results to jobs that start with the defined prefix (optional)
    x_nomad_token = "X-Nomad-Token_example" # str | A Nomad ACL token. (optional)
    per_page = 1 # int | Maximum number of results to return. (optional)
    next_token = "next_token_example" # str | Indicates where to start paging for queries that support pagination. (optional)

    # example passing only required values which don't have defaults set
    try:
        api_response = api_instance.get_job(job_name)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling JobsApi->get_job: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        api_response = api_instance.get_job(job_name, region=region, namespace=namespace, index=index, wait=wait, stale=stale, prefix=prefix, x_nomad_token=x_nomad_token, per_page=per_page, next_token=next_token)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling JobsApi->get_job: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **job_name** | **str**| The job identifier. |
 **region** | **str**| Filters results based on the specified region. | [optional]
 **namespace** | **str**| Filters results based on the specified namespace. | [optional]
 **index** | **int**| If set, wait until query exceeds given index. Must be provided with WaitParam. | [optional]
 **wait** | **str**| Provided with IndexParam to wait for change. | [optional]
 **stale** | **str**| If present, results will include stale reads. | [optional]
 **prefix** | **str**| Constrains results to jobs that start with the defined prefix | [optional]
 **x_nomad_token** | **str**| A Nomad ACL token. | [optional]
 **per_page** | **int**| Maximum number of results to return. | [optional]
 **next_token** | **str**| Indicates where to start paging for queries that support pagination. | [optional]

### Return type

[**Job**](Job.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** |  |  * X-Nomad-Index - A unique identifier representing the current state of the requested resource. On a new Nomad cluster the value of this index starts at 1. <br>  * X-Nomad-KnownLeader - Boolean indicating if there is a known cluster leader. <br>  * X-Nomad-LastContact - The time in milliseconds that a server was last contacted by the leader node. <br>  |
**400** | Bad request |  -  |
**403** | Forbidden |  -  |
**405** | Method not allowed |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_job_allocations**
> [AllocationListStub] get_job_allocations(job_name)



### Example

* Api Key Authentication (X-Nomad-Token):
```python
import time
import nomad_client
from nomad_client.api import jobs_api
from nomad_client.model.allocation_list_stub import AllocationListStub
from pprint import pprint
# Defining the host is optional and defaults to https://127.0.0.1:4646/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = nomad_client.Configuration(
    host = "https://127.0.0.1:4646/v1"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: X-Nomad-Token
configuration.api_key['X-Nomad-Token'] = 'YOUR_API_KEY'

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['X-Nomad-Token'] = 'Bearer'

# Enter a context with an instance of the API client
with nomad_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = jobs_api.JobsApi(api_client)
    job_name = "jobName_example" # str | The job identifier.
    region = "region_example" # str | Filters results based on the specified region. (optional)
    namespace = "namespace_example" # str | Filters results based on the specified namespace. (optional)
    index = 1 # int | If set, wait until query exceeds given index. Must be provided with WaitParam. (optional)
    wait = "wait_example" # str | Provided with IndexParam to wait for change. (optional)
    stale = "stale_example" # str | If present, results will include stale reads. (optional)
    prefix = "prefix_example" # str | Constrains results to jobs that start with the defined prefix (optional)
    x_nomad_token = "X-Nomad-Token_example" # str | A Nomad ACL token. (optional)
    per_page = 1 # int | Maximum number of results to return. (optional)
    next_token = "next_token_example" # str | Indicates where to start paging for queries that support pagination. (optional)
    all = True # bool | Specifies whether the list of allocations should include allocations from a previously registered job with the same ID. This is possible if the job is deregistered and reregistered. (optional)

    # example passing only required values which don't have defaults set
    try:
        api_response = api_instance.get_job_allocations(job_name)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling JobsApi->get_job_allocations: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        api_response = api_instance.get_job_allocations(job_name, region=region, namespace=namespace, index=index, wait=wait, stale=stale, prefix=prefix, x_nomad_token=x_nomad_token, per_page=per_page, next_token=next_token, all=all)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling JobsApi->get_job_allocations: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **job_name** | **str**| The job identifier. |
 **region** | **str**| Filters results based on the specified region. | [optional]
 **namespace** | **str**| Filters results based on the specified namespace. | [optional]
 **index** | **int**| If set, wait until query exceeds given index. Must be provided with WaitParam. | [optional]
 **wait** | **str**| Provided with IndexParam to wait for change. | [optional]
 **stale** | **str**| If present, results will include stale reads. | [optional]
 **prefix** | **str**| Constrains results to jobs that start with the defined prefix | [optional]
 **x_nomad_token** | **str**| A Nomad ACL token. | [optional]
 **per_page** | **int**| Maximum number of results to return. | [optional]
 **next_token** | **str**| Indicates where to start paging for queries that support pagination. | [optional]
 **all** | **bool**| Specifies whether the list of allocations should include allocations from a previously registered job with the same ID. This is possible if the job is deregistered and reregistered. | [optional]

### Return type

[**[AllocationListStub]**](AllocationListStub.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** |  |  * X-Nomad-Index - A unique identifier representing the current state of the requested resource. On a new Nomad cluster the value of this index starts at 1. <br>  * X-Nomad-KnownLeader - Boolean indicating if there is a known cluster leader. <br>  * X-Nomad-LastContact - The time in milliseconds that a server was last contacted by the leader node. <br>  |
**400** | Bad request |  -  |
**403** | Forbidden |  -  |
**405** | Method not allowed |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_job_deployment**
> Deployment get_job_deployment(job_name)



### Example

* Api Key Authentication (X-Nomad-Token):
```python
import time
import nomad_client
from nomad_client.api import jobs_api
from nomad_client.model.deployment import Deployment
from pprint import pprint
# Defining the host is optional and defaults to https://127.0.0.1:4646/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = nomad_client.Configuration(
    host = "https://127.0.0.1:4646/v1"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: X-Nomad-Token
configuration.api_key['X-Nomad-Token'] = 'YOUR_API_KEY'

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['X-Nomad-Token'] = 'Bearer'

# Enter a context with an instance of the API client
with nomad_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = jobs_api.JobsApi(api_client)
    job_name = "jobName_example" # str | The job identifier.
    region = "region_example" # str | Filters results based on the specified region. (optional)
    namespace = "namespace_example" # str | Filters results based on the specified namespace. (optional)
    index = 1 # int | If set, wait until query exceeds given index. Must be provided with WaitParam. (optional)
    wait = "wait_example" # str | Provided with IndexParam to wait for change. (optional)
    stale = "stale_example" # str | If present, results will include stale reads. (optional)
    prefix = "prefix_example" # str | Constrains results to jobs that start with the defined prefix (optional)
    x_nomad_token = "X-Nomad-Token_example" # str | A Nomad ACL token. (optional)
    per_page = 1 # int | Maximum number of results to return. (optional)
    next_token = "next_token_example" # str | Indicates where to start paging for queries that support pagination. (optional)

    # example passing only required values which don't have defaults set
    try:
        api_response = api_instance.get_job_deployment(job_name)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling JobsApi->get_job_deployment: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        api_response = api_instance.get_job_deployment(job_name, region=region, namespace=namespace, index=index, wait=wait, stale=stale, prefix=prefix, x_nomad_token=x_nomad_token, per_page=per_page, next_token=next_token)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling JobsApi->get_job_deployment: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **job_name** | **str**| The job identifier. |
 **region** | **str**| Filters results based on the specified region. | [optional]
 **namespace** | **str**| Filters results based on the specified namespace. | [optional]
 **index** | **int**| If set, wait until query exceeds given index. Must be provided with WaitParam. | [optional]
 **wait** | **str**| Provided with IndexParam to wait for change. | [optional]
 **stale** | **str**| If present, results will include stale reads. | [optional]
 **prefix** | **str**| Constrains results to jobs that start with the defined prefix | [optional]
 **x_nomad_token** | **str**| A Nomad ACL token. | [optional]
 **per_page** | **int**| Maximum number of results to return. | [optional]
 **next_token** | **str**| Indicates where to start paging for queries that support pagination. | [optional]

### Return type

[**Deployment**](Deployment.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** |  |  * X-Nomad-Index - A unique identifier representing the current state of the requested resource. On a new Nomad cluster the value of this index starts at 1. <br>  * X-Nomad-KnownLeader - Boolean indicating if there is a known cluster leader. <br>  * X-Nomad-LastContact - The time in milliseconds that a server was last contacted by the leader node. <br>  |
**400** | Bad request |  -  |
**403** | Forbidden |  -  |
**405** | Method not allowed |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_job_deployments**
> [Deployment] get_job_deployments(job_name)



### Example

* Api Key Authentication (X-Nomad-Token):
```python
import time
import nomad_client
from nomad_client.api import jobs_api
from nomad_client.model.deployment import Deployment
from pprint import pprint
# Defining the host is optional and defaults to https://127.0.0.1:4646/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = nomad_client.Configuration(
    host = "https://127.0.0.1:4646/v1"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: X-Nomad-Token
configuration.api_key['X-Nomad-Token'] = 'YOUR_API_KEY'

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['X-Nomad-Token'] = 'Bearer'

# Enter a context with an instance of the API client
with nomad_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = jobs_api.JobsApi(api_client)
    job_name = "jobName_example" # str | The job identifier.
    region = "region_example" # str | Filters results based on the specified region. (optional)
    namespace = "namespace_example" # str | Filters results based on the specified namespace. (optional)
    index = 1 # int | If set, wait until query exceeds given index. Must be provided with WaitParam. (optional)
    wait = "wait_example" # str | Provided with IndexParam to wait for change. (optional)
    stale = "stale_example" # str | If present, results will include stale reads. (optional)
    prefix = "prefix_example" # str | Constrains results to jobs that start with the defined prefix (optional)
    x_nomad_token = "X-Nomad-Token_example" # str | A Nomad ACL token. (optional)
    per_page = 1 # int | Maximum number of results to return. (optional)
    next_token = "next_token_example" # str | Indicates where to start paging for queries that support pagination. (optional)
    all = 1 # int | Flag indicating whether to constrain by job creation index or not. (optional)

    # example passing only required values which don't have defaults set
    try:
        api_response = api_instance.get_job_deployments(job_name)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling JobsApi->get_job_deployments: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        api_response = api_instance.get_job_deployments(job_name, region=region, namespace=namespace, index=index, wait=wait, stale=stale, prefix=prefix, x_nomad_token=x_nomad_token, per_page=per_page, next_token=next_token, all=all)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling JobsApi->get_job_deployments: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **job_name** | **str**| The job identifier. |
 **region** | **str**| Filters results based on the specified region. | [optional]
 **namespace** | **str**| Filters results based on the specified namespace. | [optional]
 **index** | **int**| If set, wait until query exceeds given index. Must be provided with WaitParam. | [optional]
 **wait** | **str**| Provided with IndexParam to wait for change. | [optional]
 **stale** | **str**| If present, results will include stale reads. | [optional]
 **prefix** | **str**| Constrains results to jobs that start with the defined prefix | [optional]
 **x_nomad_token** | **str**| A Nomad ACL token. | [optional]
 **per_page** | **int**| Maximum number of results to return. | [optional]
 **next_token** | **str**| Indicates where to start paging for queries that support pagination. | [optional]
 **all** | **int**| Flag indicating whether to constrain by job creation index or not. | [optional]

### Return type

[**[Deployment]**](Deployment.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** |  |  * X-Nomad-Index - A unique identifier representing the current state of the requested resource. On a new Nomad cluster the value of this index starts at 1. <br>  * X-Nomad-KnownLeader - Boolean indicating if there is a known cluster leader. <br>  * X-Nomad-LastContact - The time in milliseconds that a server was last contacted by the leader node. <br>  |
**400** | Bad request |  -  |
**403** | Forbidden |  -  |
**405** | Method not allowed |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_job_evaluations**
> [Evaluation] get_job_evaluations(job_name)



### Example

* Api Key Authentication (X-Nomad-Token):
```python
import time
import nomad_client
from nomad_client.api import jobs_api
from nomad_client.model.evaluation import Evaluation
from pprint import pprint
# Defining the host is optional and defaults to https://127.0.0.1:4646/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = nomad_client.Configuration(
    host = "https://127.0.0.1:4646/v1"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: X-Nomad-Token
configuration.api_key['X-Nomad-Token'] = 'YOUR_API_KEY'

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['X-Nomad-Token'] = 'Bearer'

# Enter a context with an instance of the API client
with nomad_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = jobs_api.JobsApi(api_client)
    job_name = "jobName_example" # str | The job identifier.
    region = "region_example" # str | Filters results based on the specified region. (optional)
    namespace = "namespace_example" # str | Filters results based on the specified namespace. (optional)
    index = 1 # int | If set, wait until query exceeds given index. Must be provided with WaitParam. (optional)
    wait = "wait_example" # str | Provided with IndexParam to wait for change. (optional)
    stale = "stale_example" # str | If present, results will include stale reads. (optional)
    prefix = "prefix_example" # str | Constrains results to jobs that start with the defined prefix (optional)
    x_nomad_token = "X-Nomad-Token_example" # str | A Nomad ACL token. (optional)
    per_page = 1 # int | Maximum number of results to return. (optional)
    next_token = "next_token_example" # str | Indicates where to start paging for queries that support pagination. (optional)

    # example passing only required values which don't have defaults set
    try:
        api_response = api_instance.get_job_evaluations(job_name)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling JobsApi->get_job_evaluations: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        api_response = api_instance.get_job_evaluations(job_name, region=region, namespace=namespace, index=index, wait=wait, stale=stale, prefix=prefix, x_nomad_token=x_nomad_token, per_page=per_page, next_token=next_token)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling JobsApi->get_job_evaluations: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **job_name** | **str**| The job identifier. |
 **region** | **str**| Filters results based on the specified region. | [optional]
 **namespace** | **str**| Filters results based on the specified namespace. | [optional]
 **index** | **int**| If set, wait until query exceeds given index. Must be provided with WaitParam. | [optional]
 **wait** | **str**| Provided with IndexParam to wait for change. | [optional]
 **stale** | **str**| If present, results will include stale reads. | [optional]
 **prefix** | **str**| Constrains results to jobs that start with the defined prefix | [optional]
 **x_nomad_token** | **str**| A Nomad ACL token. | [optional]
 **per_page** | **int**| Maximum number of results to return. | [optional]
 **next_token** | **str**| Indicates where to start paging for queries that support pagination. | [optional]

### Return type

[**[Evaluation]**](Evaluation.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** |  |  * X-Nomad-Index - A unique identifier representing the current state of the requested resource. On a new Nomad cluster the value of this index starts at 1. <br>  * X-Nomad-KnownLeader - Boolean indicating if there is a known cluster leader. <br>  * X-Nomad-LastContact - The time in milliseconds that a server was last contacted by the leader node. <br>  |
**400** | Bad request |  -  |
**403** | Forbidden |  -  |
**405** | Method not allowed |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_job_scale_status**
> JobScaleStatusResponse get_job_scale_status(job_name)



### Example

* Api Key Authentication (X-Nomad-Token):
```python
import time
import nomad_client
from nomad_client.api import jobs_api
from nomad_client.model.job_scale_status_response import JobScaleStatusResponse
from pprint import pprint
# Defining the host is optional and defaults to https://127.0.0.1:4646/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = nomad_client.Configuration(
    host = "https://127.0.0.1:4646/v1"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: X-Nomad-Token
configuration.api_key['X-Nomad-Token'] = 'YOUR_API_KEY'

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['X-Nomad-Token'] = 'Bearer'

# Enter a context with an instance of the API client
with nomad_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = jobs_api.JobsApi(api_client)
    job_name = "jobName_example" # str | The job identifier.
    region = "region_example" # str | Filters results based on the specified region. (optional)
    namespace = "namespace_example" # str | Filters results based on the specified namespace. (optional)
    index = 1 # int | If set, wait until query exceeds given index. Must be provided with WaitParam. (optional)
    wait = "wait_example" # str | Provided with IndexParam to wait for change. (optional)
    stale = "stale_example" # str | If present, results will include stale reads. (optional)
    prefix = "prefix_example" # str | Constrains results to jobs that start with the defined prefix (optional)
    x_nomad_token = "X-Nomad-Token_example" # str | A Nomad ACL token. (optional)
    per_page = 1 # int | Maximum number of results to return. (optional)
    next_token = "next_token_example" # str | Indicates where to start paging for queries that support pagination. (optional)

    # example passing only required values which don't have defaults set
    try:
        api_response = api_instance.get_job_scale_status(job_name)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling JobsApi->get_job_scale_status: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        api_response = api_instance.get_job_scale_status(job_name, region=region, namespace=namespace, index=index, wait=wait, stale=stale, prefix=prefix, x_nomad_token=x_nomad_token, per_page=per_page, next_token=next_token)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling JobsApi->get_job_scale_status: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **job_name** | **str**| The job identifier. |
 **region** | **str**| Filters results based on the specified region. | [optional]
 **namespace** | **str**| Filters results based on the specified namespace. | [optional]
 **index** | **int**| If set, wait until query exceeds given index. Must be provided with WaitParam. | [optional]
 **wait** | **str**| Provided with IndexParam to wait for change. | [optional]
 **stale** | **str**| If present, results will include stale reads. | [optional]
 **prefix** | **str**| Constrains results to jobs that start with the defined prefix | [optional]
 **x_nomad_token** | **str**| A Nomad ACL token. | [optional]
 **per_page** | **int**| Maximum number of results to return. | [optional]
 **next_token** | **str**| Indicates where to start paging for queries that support pagination. | [optional]

### Return type

[**JobScaleStatusResponse**](JobScaleStatusResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** |  |  * X-Nomad-Index - A unique identifier representing the current state of the requested resource. On a new Nomad cluster the value of this index starts at 1. <br>  * X-Nomad-KnownLeader - Boolean indicating if there is a known cluster leader. <br>  * X-Nomad-LastContact - The time in milliseconds that a server was last contacted by the leader node. <br>  |
**400** | Bad request |  -  |
**403** | Forbidden |  -  |
**405** | Method not allowed |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_job_summary**
> JobSummary get_job_summary(job_name)



### Example

* Api Key Authentication (X-Nomad-Token):
```python
import time
import nomad_client
from nomad_client.api import jobs_api
from nomad_client.model.job_summary import JobSummary
from pprint import pprint
# Defining the host is optional and defaults to https://127.0.0.1:4646/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = nomad_client.Configuration(
    host = "https://127.0.0.1:4646/v1"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: X-Nomad-Token
configuration.api_key['X-Nomad-Token'] = 'YOUR_API_KEY'

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['X-Nomad-Token'] = 'Bearer'

# Enter a context with an instance of the API client
with nomad_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = jobs_api.JobsApi(api_client)
    job_name = "jobName_example" # str | The job identifier.
    region = "region_example" # str | Filters results based on the specified region. (optional)
    namespace = "namespace_example" # str | Filters results based on the specified namespace. (optional)
    index = 1 # int | If set, wait until query exceeds given index. Must be provided with WaitParam. (optional)
    wait = "wait_example" # str | Provided with IndexParam to wait for change. (optional)
    stale = "stale_example" # str | If present, results will include stale reads. (optional)
    prefix = "prefix_example" # str | Constrains results to jobs that start with the defined prefix (optional)
    x_nomad_token = "X-Nomad-Token_example" # str | A Nomad ACL token. (optional)
    per_page = 1 # int | Maximum number of results to return. (optional)
    next_token = "next_token_example" # str | Indicates where to start paging for queries that support pagination. (optional)

    # example passing only required values which don't have defaults set
    try:
        api_response = api_instance.get_job_summary(job_name)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling JobsApi->get_job_summary: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        api_response = api_instance.get_job_summary(job_name, region=region, namespace=namespace, index=index, wait=wait, stale=stale, prefix=prefix, x_nomad_token=x_nomad_token, per_page=per_page, next_token=next_token)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling JobsApi->get_job_summary: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **job_name** | **str**| The job identifier. |
 **region** | **str**| Filters results based on the specified region. | [optional]
 **namespace** | **str**| Filters results based on the specified namespace. | [optional]
 **index** | **int**| If set, wait until query exceeds given index. Must be provided with WaitParam. | [optional]
 **wait** | **str**| Provided with IndexParam to wait for change. | [optional]
 **stale** | **str**| If present, results will include stale reads. | [optional]
 **prefix** | **str**| Constrains results to jobs that start with the defined prefix | [optional]
 **x_nomad_token** | **str**| A Nomad ACL token. | [optional]
 **per_page** | **int**| Maximum number of results to return. | [optional]
 **next_token** | **str**| Indicates where to start paging for queries that support pagination. | [optional]

### Return type

[**JobSummary**](JobSummary.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** |  |  * X-Nomad-Index - A unique identifier representing the current state of the requested resource. On a new Nomad cluster the value of this index starts at 1. <br>  * X-Nomad-KnownLeader - Boolean indicating if there is a known cluster leader. <br>  * X-Nomad-LastContact - The time in milliseconds that a server was last contacted by the leader node. <br>  |
**400** | Bad request |  -  |
**403** | Forbidden |  -  |
**405** | Method not allowed |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_job_versions**
> JobVersionsResponse get_job_versions(job_name)



### Example

* Api Key Authentication (X-Nomad-Token):
```python
import time
import nomad_client
from nomad_client.api import jobs_api
from nomad_client.model.job_versions_response import JobVersionsResponse
from pprint import pprint
# Defining the host is optional and defaults to https://127.0.0.1:4646/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = nomad_client.Configuration(
    host = "https://127.0.0.1:4646/v1"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: X-Nomad-Token
configuration.api_key['X-Nomad-Token'] = 'YOUR_API_KEY'

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['X-Nomad-Token'] = 'Bearer'

# Enter a context with an instance of the API client
with nomad_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = jobs_api.JobsApi(api_client)
    job_name = "jobName_example" # str | The job identifier.
    region = "region_example" # str | Filters results based on the specified region. (optional)
    namespace = "namespace_example" # str | Filters results based on the specified namespace. (optional)
    index = 1 # int | If set, wait until query exceeds given index. Must be provided with WaitParam. (optional)
    wait = "wait_example" # str | Provided with IndexParam to wait for change. (optional)
    stale = "stale_example" # str | If present, results will include stale reads. (optional)
    prefix = "prefix_example" # str | Constrains results to jobs that start with the defined prefix (optional)
    x_nomad_token = "X-Nomad-Token_example" # str | A Nomad ACL token. (optional)
    per_page = 1 # int | Maximum number of results to return. (optional)
    next_token = "next_token_example" # str | Indicates where to start paging for queries that support pagination. (optional)
    diffs = True # bool | Boolean flag indicating whether to compute job diffs. (optional)

    # example passing only required values which don't have defaults set
    try:
        api_response = api_instance.get_job_versions(job_name)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling JobsApi->get_job_versions: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        api_response = api_instance.get_job_versions(job_name, region=region, namespace=namespace, index=index, wait=wait, stale=stale, prefix=prefix, x_nomad_token=x_nomad_token, per_page=per_page, next_token=next_token, diffs=diffs)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling JobsApi->get_job_versions: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **job_name** | **str**| The job identifier. |
 **region** | **str**| Filters results based on the specified region. | [optional]
 **namespace** | **str**| Filters results based on the specified namespace. | [optional]
 **index** | **int**| If set, wait until query exceeds given index. Must be provided with WaitParam. | [optional]
 **wait** | **str**| Provided with IndexParam to wait for change. | [optional]
 **stale** | **str**| If present, results will include stale reads. | [optional]
 **prefix** | **str**| Constrains results to jobs that start with the defined prefix | [optional]
 **x_nomad_token** | **str**| A Nomad ACL token. | [optional]
 **per_page** | **int**| Maximum number of results to return. | [optional]
 **next_token** | **str**| Indicates where to start paging for queries that support pagination. | [optional]
 **diffs** | **bool**| Boolean flag indicating whether to compute job diffs. | [optional]

### Return type

[**JobVersionsResponse**](JobVersionsResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** |  |  * X-Nomad-Index - A unique identifier representing the current state of the requested resource. On a new Nomad cluster the value of this index starts at 1. <br>  * X-Nomad-KnownLeader - Boolean indicating if there is a known cluster leader. <br>  * X-Nomad-LastContact - The time in milliseconds that a server was last contacted by the leader node. <br>  |
**400** | Bad request |  -  |
**403** | Forbidden |  -  |
**405** | Method not allowed |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_jobs**
> [JobListStub] get_jobs()



### Example

* Api Key Authentication (X-Nomad-Token):
```python
import time
import nomad_client
from nomad_client.api import jobs_api
from nomad_client.model.job_list_stub import JobListStub
from pprint import pprint
# Defining the host is optional and defaults to https://127.0.0.1:4646/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = nomad_client.Configuration(
    host = "https://127.0.0.1:4646/v1"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: X-Nomad-Token
configuration.api_key['X-Nomad-Token'] = 'YOUR_API_KEY'

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['X-Nomad-Token'] = 'Bearer'

# Enter a context with an instance of the API client
with nomad_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = jobs_api.JobsApi(api_client)
    region = "region_example" # str | Filters results based on the specified region. (optional)
    namespace = "namespace_example" # str | Filters results based on the specified namespace. (optional)
    index = 1 # int | If set, wait until query exceeds given index. Must be provided with WaitParam. (optional)
    wait = "wait_example" # str | Provided with IndexParam to wait for change. (optional)
    stale = "stale_example" # str | If present, results will include stale reads. (optional)
    prefix = "prefix_example" # str | Constrains results to jobs that start with the defined prefix (optional)
    x_nomad_token = "X-Nomad-Token_example" # str | A Nomad ACL token. (optional)
    per_page = 1 # int | Maximum number of results to return. (optional)
    next_token = "next_token_example" # str | Indicates where to start paging for queries that support pagination. (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        api_response = api_instance.get_jobs(region=region, namespace=namespace, index=index, wait=wait, stale=stale, prefix=prefix, x_nomad_token=x_nomad_token, per_page=per_page, next_token=next_token)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling JobsApi->get_jobs: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **region** | **str**| Filters results based on the specified region. | [optional]
 **namespace** | **str**| Filters results based on the specified namespace. | [optional]
 **index** | **int**| If set, wait until query exceeds given index. Must be provided with WaitParam. | [optional]
 **wait** | **str**| Provided with IndexParam to wait for change. | [optional]
 **stale** | **str**| If present, results will include stale reads. | [optional]
 **prefix** | **str**| Constrains results to jobs that start with the defined prefix | [optional]
 **x_nomad_token** | **str**| A Nomad ACL token. | [optional]
 **per_page** | **int**| Maximum number of results to return. | [optional]
 **next_token** | **str**| Indicates where to start paging for queries that support pagination. | [optional]

### Return type

[**[JobListStub]**](JobListStub.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** |  |  * X-Nomad-Index - A unique identifier representing the current state of the requested resource. On a new Nomad cluster the value of this index starts at 1. <br>  * X-Nomad-KnownLeader - Boolean indicating if there is a known cluster leader. <br>  * X-Nomad-LastContact - The time in milliseconds that a server was last contacted by the leader node. <br>  |
**400** | Bad request |  -  |
**403** | Forbidden |  -  |
**405** | Method not allowed |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **post_job**
> JobRegisterResponse post_job(job_name, job_register_request)



### Example

* Api Key Authentication (X-Nomad-Token):
```python
import time
import nomad_client
from nomad_client.api import jobs_api
from nomad_client.model.job_register_request import JobRegisterRequest
from nomad_client.model.job_register_response import JobRegisterResponse
from pprint import pprint
# Defining the host is optional and defaults to https://127.0.0.1:4646/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = nomad_client.Configuration(
    host = "https://127.0.0.1:4646/v1"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: X-Nomad-Token
configuration.api_key['X-Nomad-Token'] = 'YOUR_API_KEY'

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['X-Nomad-Token'] = 'Bearer'

# Enter a context with an instance of the API client
with nomad_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = jobs_api.JobsApi(api_client)
    job_name = "jobName_example" # str | The job identifier.
    job_register_request = JobRegisterRequest(
        enforce_index=True,
        eval_priority=1,
        job=Job(
            affinities=[
                Affinity(
                    l_target="l_target_example",
                    operand="operand_example",
                    r_target="r_target_example",
                    weight=-128,
                ),
            ],
            all_at_once=True,
            constraints=[
                Constraint(
                    l_target="l_target_example",
                    operand="operand_example",
                    r_target="r_target_example",
                ),
            ],
            consul_namespace="consul_namespace_example",
            consul_token="consul_token_example",
            create_index=0,
            datacenters=[
                "datacenters_example",
            ],
            dispatch_idempotency_token="dispatch_idempotency_token_example",
            dispatched=True,
            id="id_example",
            job_modify_index=0,
            meta={
                "key": "key_example",
            },
            migrate=MigrateStrategy(
                health_check="health_check_example",
                healthy_deadline=1,
                max_parallel=1,
                min_healthy_time=1,
            ),
            modify_index=0,
            multiregion=Multiregion(
                regions=[
                    MultiregionRegion(
                        count=1,
                        datacenters=[
                            "datacenters_example",
                        ],
                        meta={
                            "key": "key_example",
                        },
                        name="name_example",
                    ),
                ],
                strategy=MultiregionStrategy(
                    max_parallel=1,
                    on_failure="on_failure_example",
                ),
            ),
            name="name_example",
            namespace="namespace_example",
            nomad_token_id="nomad_token_id_example",
            parameterized_job=ParameterizedJobConfig(
                meta_optional=[
                    "meta_optional_example",
                ],
                meta_required=[
                    "meta_required_example",
                ],
                payload="payload_example",
            ),
            parent_id="parent_id_example",
            payload='YQ==',
            periodic=PeriodicConfig(
                enabled=True,
                prohibit_overlap=True,
                spec="spec_example",
                spec_type="spec_type_example",
                time_zone="time_zone_example",
            ),
            priority=1,
            region="region_example",
            reschedule=ReschedulePolicy(
                attempts=1,
                delay=1,
                delay_function="delay_function_example",
                interval=1,
                max_delay=1,
                unlimited=True,
            ),
            spreads=[
                Spread(
                    attribute="attribute_example",
                    spread_target=[
                        SpreadTarget(
                            percent=0,
                            value="value_example",
                        ),
                    ],
                    weight=-128,
                ),
            ],
            stable=True,
            status="status_example",
            status_description="status_description_example",
            stop=True,
            submit_time=1,
            task_groups=[
                TaskGroup(
                    affinities=[
                        Affinity(
                            l_target="l_target_example",
                            operand="operand_example",
                            r_target="r_target_example",
                            weight=-128,
                        ),
                    ],
                    constraints=[
                        Constraint(
                            l_target="l_target_example",
                            operand="operand_example",
                            r_target="r_target_example",
                        ),
                    ],
                    consul=Consul(
                        namespace="namespace_example",
                    ),
                    count=1,
                    ephemeral_disk=EphemeralDisk(
                        migrate=True,
                        size_mb=1,
                        sticky=True,
                    ),
                    max_client_disconnect=1,
                    meta={
                        "key": "key_example",
                    },
                    migrate=MigrateStrategy(
                        health_check="health_check_example",
                        healthy_deadline=1,
                        max_parallel=1,
                        min_healthy_time=1,
                    ),
                    name="name_example",
                    networks=[
                        NetworkResource(
                            cidr="cidr_example",
                            dns=DNSConfig(
                                options=[
                                    "options_example",
                                ],
                                searches=[
                                    "searches_example",
                                ],
                                servers=[
                                    "servers_example",
                                ],
                            ),
                            device="device_example",
                            dynamic_ports=[
                                Port(
                                    host_network="host_network_example",
                                    label="label_example",
                                    to=1,
                                    value=1,
                                ),
                            ],
                            hostname="hostname_example",
                            ip="ip_example",
                            m_bits=1,
                            mode="mode_example",
                            reserved_ports=[
                                Port(
                                    host_network="host_network_example",
                                    label="label_example",
                                    to=1,
                                    value=1,
                                ),
                            ],
                        ),
                    ],
                    reschedule_policy=ReschedulePolicy(
                        attempts=1,
                        delay=1,
                        delay_function="delay_function_example",
                        interval=1,
                        max_delay=1,
                        unlimited=True,
                    ),
                    restart_policy=RestartPolicy(
                        attempts=1,
                        delay=1,
                        interval=1,
                        mode="mode_example",
                    ),
                    scaling=ScalingPolicy(
                        create_index=0,
                        enabled=True,
                        id="id_example",
                        max=1,
                        min=1,
                        modify_index=0,
                        namespace="namespace_example",
                        policy={
                            "key": None,
                        },
                        target={
                            "key": "key_example",
                        },
                        type="type_example",
                    ),
                    services=[
                        Service(
                            address_mode="address_mode_example",
                            canary_meta={
                                "key": "key_example",
                            },
                            canary_tags=[
                                "canary_tags_example",
                            ],
                            check_restart=CheckRestart(
                                grace=1,
                                ignore_warnings=True,
                                limit=1,
                            ),
                            checks=[
                                ServiceCheck(
                                    address_mode="address_mode_example",
                                    args=[
                                        "args_example",
                                    ],
                                    body="body_example",
                                    check_restart=CheckRestart(
                                        grace=1,
                                        ignore_warnings=True,
                                        limit=1,
                                    ),
                                    command="command_example",
                                    expose=True,
                                    failures_before_critical=1,
                                    grpc_service="grpc_service_example",
                                    grpc_use_tls=True,
                                    header={
                                        "key": [
                                            "key_example",
                                        ],
                                    },
                                    id="id_example",
                                    initial_status="initial_status_example",
                                    interval=1,
                                    method="method_example",
                                    name="name_example",
                                    on_update="on_update_example",
                                    path="path_example",
                                    port_label="port_label_example",
                                    protocol="protocol_example",
                                    success_before_passing=1,
                                    tls_skip_verify=True,
                                    task_name="task_name_example",
                                    timeout=1,
                                    type="type_example",
                                ),
                            ],
                            connect=ConsulConnect(
                                gateway=ConsulGateway(
                                    ingress=ConsulIngressConfigEntry(
                                        listeners=[
                                            ConsulIngressListener(
                                                port=1,
                                                protocol="protocol_example",
                                                services=[
                                                    ConsulIngressService(
                                                        hosts=[
                                                            "hosts_example",
                                                        ],
                                                        name="name_example",
                                                    ),
                                                ],
                                            ),
                                        ],
                                        tls=ConsulGatewayTLSConfig(
                                            enabled=True,
                                        ),
                                    ),
                                    mesh=None,
                                    proxy=ConsulGatewayProxy(
                                        config={
                                            "key": None,
                                        },
                                        connect_timeout=1,
                                        envoy_dns_discovery_type="envoy_dns_discovery_type_example",
                                        envoy_gateway_bind_addresses={
                                            "key": ConsulGatewayBindAddress(
                                                address="address_example",
                                                name="name_example",
                                                port=1,
                                            ),
                                        },
                                        envoy_gateway_bind_tagged_addresses=True,
                                        envoy_gateway_no_default_bind=True,
                                    ),
                                    terminating=ConsulTerminatingConfigEntry(
                                        services=[
                                            ConsulLinkedService(
                                                ca_file="ca_file_example",
                                                cert_file="cert_file_example",
                                                key_file="key_file_example",
                                                name="name_example",
                                                sni="sni_example",
                                            ),
                                        ],
                                    ),
                                ),
                                native=True,
                                sidecar_service=ConsulSidecarService(
                                    disable_default_tcp_check=True,
                                    port="port_example",
                                    proxy=ConsulProxy(
                                        config={
                                            "key": None,
                                        },
                                        expose_config=ConsulExposeConfig(
                                            path=[
                                                ConsulExposePath(
                                                    listener_port="listener_port_example",
                                                    local_path_port=1,
                                                    path="path_example",
                                                    protocol="protocol_example",
                                                ),
                                            ],
                                        ),
                                        local_service_address="local_service_address_example",
                                        local_service_port=1,
                                        upstreams=[
                                            ConsulUpstream(
                                                datacenter="datacenter_example",
                                                destination_name="destination_name_example",
                                                local_bind_address="local_bind_address_example",
                                                local_bind_port=1,
                                                mesh_gateway=ConsulMeshGateway(
                                                    mode="mode_example",
                                                ),
                                            ),
                                        ],
                                    ),
                                    tags=[
                                        "tags_example",
                                    ],
                                ),
                                sidecar_task=SidecarTask(
                                    config={
                                        "key": None,
                                    },
                                    driver="driver_example",
                                    env={
                                        "key": "key_example",
                                    },
                                    kill_signal="kill_signal_example",
                                    kill_timeout=1,
                                    log_config=LogConfig(
                                        max_file_size_mb=1,
                                        max_files=1,
                                    ),
                                    meta={
                                        "key": "key_example",
                                    },
                                    name="name_example",
                                    resources=Resources(
                                        cpu=1,
                                        cores=1,
                                        devices=[
                                            RequestedDevice(
                                                affinities=[
                                                    Affinity(
                                                        l_target="l_target_example",
                                                        operand="operand_example",
                                                        r_target="r_target_example",
                                                        weight=-128,
                                                    ),
                                                ],
                                                constraints=[
                                                    Constraint(
                                                        l_target="l_target_example",
                                                        operand="operand_example",
                                                        r_target="r_target_example",
                                                    ),
                                                ],
                                                count=0,
                                                name="name_example",
                                            ),
                                        ],
                                        disk_mb=1,
                                        iops=1,
                                        memory_mb=1,
                                        memory_max_mb=1,
                                        networks=[
                                            NetworkResource(
                                                cidr="cidr_example",
                                                dns=DNSConfig(
                                                    options=[
                                                        "options_example",
                                                    ],
                                                    searches=[
                                                        "searches_example",
                                                    ],
                                                    servers=[
                                                        "servers_example",
                                                    ],
                                                ),
                                                device="device_example",
                                                dynamic_ports=[
                                                    Port(
                                                        host_network="host_network_example",
                                                        label="label_example",
                                                        to=1,
                                                        value=1,
                                                    ),
                                                ],
                                                hostname="hostname_example",
                                                ip="ip_example",
                                                m_bits=1,
                                                mode="mode_example",
                                                reserved_ports=[
                                                    Port(
                                                        host_network="host_network_example",
                                                        label="label_example",
                                                        to=1,
                                                        value=1,
                                                    ),
                                                ],
                                            ),
                                        ],
                                    ),
                                    shutdown_delay=1,
                                    user="user_example",
                                ),
                            ),
                            enable_tag_override=True,
                            id="id_example",
                            meta={
                                "key": "key_example",
                            },
                            name="name_example",
                            on_update="on_update_example",
                            port_label="port_label_example",
                            provider="provider_example",
                            tags=[
                                "tags_example",
                            ],
                            task_name="task_name_example",
                        ),
                    ],
                    shutdown_delay=1,
                    spreads=[
                        Spread(
                            attribute="attribute_example",
                            spread_target=[
                                SpreadTarget(
                                    percent=0,
                                    value="value_example",
                                ),
                            ],
                            weight=-128,
                        ),
                    ],
                    stop_after_client_disconnect=1,
                    tasks=[
                        Task(
                            affinities=[
                                Affinity(
                                    l_target="l_target_example",
                                    operand="operand_example",
                                    r_target="r_target_example",
                                    weight=-128,
                                ),
                            ],
                            artifacts=[
                                TaskArtifact(
                                    getter_headers={
                                        "key": "key_example",
                                    },
                                    getter_mode="getter_mode_example",
                                    getter_options={
                                        "key": "key_example",
                                    },
                                    getter_source="getter_source_example",
                                    relative_dest="relative_dest_example",
                                ),
                            ],
                            csi_plugin_config=TaskCSIPluginConfig(
                                id="id_example",
                                mount_dir="mount_dir_example",
                                type="type_example",
                            ),
                            config={
                                "key": None,
                            },
                            constraints=[
                                Constraint(
                                    l_target="l_target_example",
                                    operand="operand_example",
                                    r_target="r_target_example",
                                ),
                            ],
                            dispatch_payload=DispatchPayloadConfig(
                                file="file_example",
                            ),
                            driver="driver_example",
                            env={
                                "key": "key_example",
                            },
                            kill_signal="kill_signal_example",
                            kill_timeout=1,
                            kind="kind_example",
                            leader=True,
                            lifecycle=TaskLifecycle(
                                hook="hook_example",
                                sidecar=True,
                            ),
                            log_config=LogConfig(
                                max_file_size_mb=1,
                                max_files=1,
                            ),
                            meta={
                                "key": "key_example",
                            },
                            name="name_example",
                            resources=Resources(
                                cpu=1,
                                cores=1,
                                devices=[
                                    RequestedDevice(
                                        affinities=[
                                            Affinity(
                                                l_target="l_target_example",
                                                operand="operand_example",
                                                r_target="r_target_example",
                                                weight=-128,
                                            ),
                                        ],
                                        constraints=[
                                            Constraint(
                                                l_target="l_target_example",
                                                operand="operand_example",
                                                r_target="r_target_example",
                                            ),
                                        ],
                                        count=0,
                                        name="name_example",
                                    ),
                                ],
                                disk_mb=1,
                                iops=1,
                                memory_mb=1,
                                memory_max_mb=1,
                                networks=[
                                    NetworkResource(
                                        cidr="cidr_example",
                                        dns=DNSConfig(
                                            options=[
                                                "options_example",
                                            ],
                                            searches=[
                                                "searches_example",
                                            ],
                                            servers=[
                                                "servers_example",
                                            ],
                                        ),
                                        device="device_example",
                                        dynamic_ports=[
                                            Port(
                                                host_network="host_network_example",
                                                label="label_example",
                                                to=1,
                                                value=1,
                                            ),
                                        ],
                                        hostname="hostname_example",
                                        ip="ip_example",
                                        m_bits=1,
                                        mode="mode_example",
                                        reserved_ports=[
                                            Port(
                                                host_network="host_network_example",
                                                label="label_example",
                                                to=1,
                                                value=1,
                                            ),
                                        ],
                                    ),
                                ],
                            ),
                            restart_policy=RestartPolicy(
                                attempts=1,
                                delay=1,
                                interval=1,
                                mode="mode_example",
                            ),
                            scaling_policies=[
                                ScalingPolicy(
                                    create_index=0,
                                    enabled=True,
                                    id="id_example",
                                    max=1,
                                    min=1,
                                    modify_index=0,
                                    namespace="namespace_example",
                                    policy={
                                        "key": None,
                                    },
                                    target={
                                        "key": "key_example",
                                    },
                                    type="type_example",
                                ),
                            ],
                            services=[
                                Service(
                                    address_mode="address_mode_example",
                                    canary_meta={
                                        "key": "key_example",
                                    },
                                    canary_tags=[
                                        "canary_tags_example",
                                    ],
                                    check_restart=CheckRestart(
                                        grace=1,
                                        ignore_warnings=True,
                                        limit=1,
                                    ),
                                    checks=[
                                        ServiceCheck(
                                            address_mode="address_mode_example",
                                            args=[
                                                "args_example",
                                            ],
                                            body="body_example",
                                            check_restart=CheckRestart(
                                                grace=1,
                                                ignore_warnings=True,
                                                limit=1,
                                            ),
                                            command="command_example",
                                            expose=True,
                                            failures_before_critical=1,
                                            grpc_service="grpc_service_example",
                                            grpc_use_tls=True,
                                            header={
                                                "key": [
                                                    "key_example",
                                                ],
                                            },
                                            id="id_example",
                                            initial_status="initial_status_example",
                                            interval=1,
                                            method="method_example",
                                            name="name_example",
                                            on_update="on_update_example",
                                            path="path_example",
                                            port_label="port_label_example",
                                            protocol="protocol_example",
                                            success_before_passing=1,
                                            tls_skip_verify=True,
                                            task_name="task_name_example",
                                            timeout=1,
                                            type="type_example",
                                        ),
                                    ],
                                    connect=ConsulConnect(
                                        gateway=ConsulGateway(
                                            ingress=ConsulIngressConfigEntry(
                                                listeners=[
                                                    ConsulIngressListener(
                                                        port=1,
                                                        protocol="protocol_example",
                                                        services=[
                                                            ConsulIngressService(
                                                                hosts=[
                                                                    "hosts_example",
                                                                ],
                                                                name="name_example",
                                                            ),
                                                        ],
                                                    ),
                                                ],
                                                tls=ConsulGatewayTLSConfig(
                                                    enabled=True,
                                                ),
                                            ),
                                            mesh=None,
                                            proxy=ConsulGatewayProxy(
                                                config={
                                                    "key": None,
                                                },
                                                connect_timeout=1,
                                                envoy_dns_discovery_type="envoy_dns_discovery_type_example",
                                                envoy_gateway_bind_addresses={
                                                    "key": ConsulGatewayBindAddress(
                                                        address="address_example",
                                                        name="name_example",
                                                        port=1,
                                                    ),
                                                },
                                                envoy_gateway_bind_tagged_addresses=True,
                                                envoy_gateway_no_default_bind=True,
                                            ),
                                            terminating=ConsulTerminatingConfigEntry(
                                                services=[
                                                    ConsulLinkedService(
                                                        ca_file="ca_file_example",
                                                        cert_file="cert_file_example",
                                                        key_file="key_file_example",
                                                        name="name_example",
                                                        sni="sni_example",
                                                    ),
                                                ],
                                            ),
                                        ),
                                        native=True,
                                        sidecar_service=ConsulSidecarService(
                                            disable_default_tcp_check=True,
                                            port="port_example",
                                            proxy=ConsulProxy(
                                                config={
                                                    "key": None,
                                                },
                                                expose_config=ConsulExposeConfig(
                                                    path=[
                                                        ConsulExposePath(
                                                            listener_port="listener_port_example",
                                                            local_path_port=1,
                                                            path="path_example",
                                                            protocol="protocol_example",
                                                        ),
                                                    ],
                                                ),
                                                local_service_address="local_service_address_example",
                                                local_service_port=1,
                                                upstreams=[
                                                    ConsulUpstream(
                                                        datacenter="datacenter_example",
                                                        destination_name="destination_name_example",
                                                        local_bind_address="local_bind_address_example",
                                                        local_bind_port=1,
                                                        mesh_gateway=ConsulMeshGateway(
                                                            mode="mode_example",
                                                        ),
                                                    ),
                                                ],
                                            ),
                                            tags=[
                                                "tags_example",
                                            ],
                                        ),
                                        sidecar_task=SidecarTask(
                                            config={
                                                "key": None,
                                            },
                                            driver="driver_example",
                                            env={
                                                "key": "key_example",
                                            },
                                            kill_signal="kill_signal_example",
                                            kill_timeout=1,
                                            log_config=LogConfig(
                                                max_file_size_mb=1,
                                                max_files=1,
                                            ),
                                            meta={
                                                "key": "key_example",
                                            },
                                            name="name_example",
                                            resources=Resources(
                                                cpu=1,
                                                cores=1,
                                                devices=[
                                                    RequestedDevice(
                                                        affinities=[
                                                            Affinity(
                                                                l_target="l_target_example",
                                                                operand="operand_example",
                                                                r_target="r_target_example",
                                                                weight=-128,
                                                            ),
                                                        ],
                                                        constraints=[
                                                            Constraint(
                                                                l_target="l_target_example",
                                                                operand="operand_example",
                                                                r_target="r_target_example",
                                                            ),
                                                        ],
                                                        count=0,
                                                        name="name_example",
                                                    ),
                                                ],
                                                disk_mb=1,
                                                iops=1,
                                                memory_mb=1,
                                                memory_max_mb=1,
                                                networks=[
                                                    NetworkResource(
                                                        cidr="cidr_example",
                                                        dns=DNSConfig(
                                                            options=[
                                                                "options_example",
                                                            ],
                                                            searches=[
                                                                "searches_example",
                                                            ],
                                                            servers=[
                                                                "servers_example",
                                                            ],
                                                        ),
                                                        device="device_example",
                                                        dynamic_ports=[
                                                            Port(
                                                                host_network="host_network_example",
                                                                label="label_example",
                                                                to=1,
                                                                value=1,
                                                            ),
                                                        ],
                                                        hostname="hostname_example",
                                                        ip="ip_example",
                                                        m_bits=1,
                                                        mode="mode_example",
                                                        reserved_ports=[
                                                            Port(
                                                                host_network="host_network_example",
                                                                label="label_example",
                                                                to=1,
                                                                value=1,
                                                            ),
                                                        ],
                                                    ),
                                                ],
                                            ),
                                            shutdown_delay=1,
                                            user="user_example",
                                        ),
                                    ),
                                    enable_tag_override=True,
                                    id="id_example",
                                    meta={
                                        "key": "key_example",
                                    },
                                    name="name_example",
                                    on_update="on_update_example",
                                    port_label="port_label_example",
                                    provider="provider_example",
                                    tags=[
                                        "tags_example",
                                    ],
                                    task_name="task_name_example",
                                ),
                            ],
                            shutdown_delay=1,
                            templates=[
                                Template(
                                    change_mode="change_mode_example",
                                    change_signal="change_signal_example",
                                    dest_path="dest_path_example",
                                    embedded_tmpl="embedded_tmpl_example",
                                    envvars=True,
                                    left_delim="left_delim_example",
                                    perms="perms_example",
                                    right_delim="right_delim_example",
                                    source_path="source_path_example",
                                    splay=1,
                                    vault_grace=1,
                                    wait=WaitConfig(
                                        max=1,
                                        min=1,
                                    ),
                                ),
                            ],
                            user="user_example",
                            vault=Vault(
                                change_mode="change_mode_example",
                                change_signal="change_signal_example",
                                entity_alias="entity_alias_example",
                                env=True,
                                namespace="namespace_example",
                                policies=[
                                    "policies_example",
                                ],
                            ),
                            volume_mounts=[
                                VolumeMount(
                                    destination="destination_example",
                                    propagation_mode="propagation_mode_example",
                                    read_only=True,
                                    volume="volume_example",
                                ),
                            ],
                        ),
                    ],
                    update=UpdateStrategy(
                        auto_promote=True,
                        auto_revert=True,
                        canary=1,
                        health_check="health_check_example",
                        healthy_deadline=1,
                        max_parallel=1,
                        min_healthy_time=1,
                        progress_deadline=1,
                        stagger=1,
                    ),
                    volumes={
                        "key": VolumeRequest(
                            access_mode="access_mode_example",
                            attachment_mode="attachment_mode_example",
                            mount_options=CSIMountOptions(
                                fs_type="fs_type_example",
                                mount_flags=[
                                    "mount_flags_example",
                                ],
                            ),
                            name="name_example",
                            per_alloc=True,
                            read_only=True,
                            source="source_example",
                            type="type_example",
                        ),
                    },
                ),
            ],
            type="type_example",
            update=UpdateStrategy(
                auto_promote=True,
                auto_revert=True,
                canary=1,
                health_check="health_check_example",
                healthy_deadline=1,
                max_parallel=1,
                min_healthy_time=1,
                progress_deadline=1,
                stagger=1,
            ),
            vault_namespace="vault_namespace_example",
            vault_token="vault_token_example",
            version=0,
        ),
        job_modify_index=0,
        namespace="namespace_example",
        policy_override=True,
        preserve_counts=True,
        region="region_example",
        secret_id="secret_id_example",
    ) # JobRegisterRequest | 
    region = "region_example" # str | Filters results based on the specified region. (optional)
    namespace = "namespace_example" # str | Filters results based on the specified namespace. (optional)
    x_nomad_token = "X-Nomad-Token_example" # str | A Nomad ACL token. (optional)
    idempotency_token = "idempotency_token_example" # str | Can be used to ensure operations are only run once. (optional)

    # example passing only required values which don't have defaults set
    try:
        api_response = api_instance.post_job(job_name, job_register_request)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling JobsApi->post_job: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        api_response = api_instance.post_job(job_name, job_register_request, region=region, namespace=namespace, x_nomad_token=x_nomad_token, idempotency_token=idempotency_token)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling JobsApi->post_job: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **job_name** | **str**| The job identifier. |
 **job_register_request** | [**JobRegisterRequest**](JobRegisterRequest.md)|  |
 **region** | **str**| Filters results based on the specified region. | [optional]
 **namespace** | **str**| Filters results based on the specified namespace. | [optional]
 **x_nomad_token** | **str**| A Nomad ACL token. | [optional]
 **idempotency_token** | **str**| Can be used to ensure operations are only run once. | [optional]

### Return type

[**JobRegisterResponse**](JobRegisterResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** |  |  * X-Nomad-Index - A unique identifier representing the current state of the requested resource. On a new Nomad cluster the value of this index starts at 1. <br>  * X-Nomad-KnownLeader - Boolean indicating if there is a known cluster leader. <br>  * X-Nomad-LastContact - The time in milliseconds that a server was last contacted by the leader node. <br>  |
**400** | Bad request |  -  |
**403** | Forbidden |  -  |
**405** | Method not allowed |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **post_job_dispatch**
> JobDispatchResponse post_job_dispatch(job_name, job_dispatch_request)



### Example

* Api Key Authentication (X-Nomad-Token):
```python
import time
import nomad_client
from nomad_client.api import jobs_api
from nomad_client.model.job_dispatch_request import JobDispatchRequest
from nomad_client.model.job_dispatch_response import JobDispatchResponse
from pprint import pprint
# Defining the host is optional and defaults to https://127.0.0.1:4646/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = nomad_client.Configuration(
    host = "https://127.0.0.1:4646/v1"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: X-Nomad-Token
configuration.api_key['X-Nomad-Token'] = 'YOUR_API_KEY'

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['X-Nomad-Token'] = 'Bearer'

# Enter a context with an instance of the API client
with nomad_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = jobs_api.JobsApi(api_client)
    job_name = "jobName_example" # str | The job identifier.
    job_dispatch_request = JobDispatchRequest(
        job_id="job_id_example",
        meta={
            "key": "key_example",
        },
        payload='YQ==',
    ) # JobDispatchRequest | 
    region = "region_example" # str | Filters results based on the specified region. (optional)
    namespace = "namespace_example" # str | Filters results based on the specified namespace. (optional)
    x_nomad_token = "X-Nomad-Token_example" # str | A Nomad ACL token. (optional)
    idempotency_token = "idempotency_token_example" # str | Can be used to ensure operations are only run once. (optional)

    # example passing only required values which don't have defaults set
    try:
        api_response = api_instance.post_job_dispatch(job_name, job_dispatch_request)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling JobsApi->post_job_dispatch: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        api_response = api_instance.post_job_dispatch(job_name, job_dispatch_request, region=region, namespace=namespace, x_nomad_token=x_nomad_token, idempotency_token=idempotency_token)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling JobsApi->post_job_dispatch: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **job_name** | **str**| The job identifier. |
 **job_dispatch_request** | [**JobDispatchRequest**](JobDispatchRequest.md)|  |
 **region** | **str**| Filters results based on the specified region. | [optional]
 **namespace** | **str**| Filters results based on the specified namespace. | [optional]
 **x_nomad_token** | **str**| A Nomad ACL token. | [optional]
 **idempotency_token** | **str**| Can be used to ensure operations are only run once. | [optional]

### Return type

[**JobDispatchResponse**](JobDispatchResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** |  |  * X-Nomad-Index - A unique identifier representing the current state of the requested resource. On a new Nomad cluster the value of this index starts at 1. <br>  |
**400** | Bad request |  -  |
**403** | Forbidden |  -  |
**405** | Method not allowed |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **post_job_evaluate**
> JobRegisterResponse post_job_evaluate(job_name, job_evaluate_request)



### Example

* Api Key Authentication (X-Nomad-Token):
```python
import time
import nomad_client
from nomad_client.api import jobs_api
from nomad_client.model.job_evaluate_request import JobEvaluateRequest
from nomad_client.model.job_register_response import JobRegisterResponse
from pprint import pprint
# Defining the host is optional and defaults to https://127.0.0.1:4646/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = nomad_client.Configuration(
    host = "https://127.0.0.1:4646/v1"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: X-Nomad-Token
configuration.api_key['X-Nomad-Token'] = 'YOUR_API_KEY'

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['X-Nomad-Token'] = 'Bearer'

# Enter a context with an instance of the API client
with nomad_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = jobs_api.JobsApi(api_client)
    job_name = "jobName_example" # str | The job identifier.
    job_evaluate_request = JobEvaluateRequest(
        eval_options=EvalOptions(
            force_reschedule=True,
        ),
        job_id="job_id_example",
        namespace="namespace_example",
        region="region_example",
        secret_id="secret_id_example",
    ) # JobEvaluateRequest | 
    region = "region_example" # str | Filters results based on the specified region. (optional)
    namespace = "namespace_example" # str | Filters results based on the specified namespace. (optional)
    x_nomad_token = "X-Nomad-Token_example" # str | A Nomad ACL token. (optional)
    idempotency_token = "idempotency_token_example" # str | Can be used to ensure operations are only run once. (optional)

    # example passing only required values which don't have defaults set
    try:
        api_response = api_instance.post_job_evaluate(job_name, job_evaluate_request)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling JobsApi->post_job_evaluate: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        api_response = api_instance.post_job_evaluate(job_name, job_evaluate_request, region=region, namespace=namespace, x_nomad_token=x_nomad_token, idempotency_token=idempotency_token)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling JobsApi->post_job_evaluate: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **job_name** | **str**| The job identifier. |
 **job_evaluate_request** | [**JobEvaluateRequest**](JobEvaluateRequest.md)|  |
 **region** | **str**| Filters results based on the specified region. | [optional]
 **namespace** | **str**| Filters results based on the specified namespace. | [optional]
 **x_nomad_token** | **str**| A Nomad ACL token. | [optional]
 **idempotency_token** | **str**| Can be used to ensure operations are only run once. | [optional]

### Return type

[**JobRegisterResponse**](JobRegisterResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** |  |  * X-Nomad-Index - A unique identifier representing the current state of the requested resource. On a new Nomad cluster the value of this index starts at 1. <br>  |
**400** | Bad request |  -  |
**403** | Forbidden |  -  |
**405** | Method not allowed |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **post_job_parse**
> Job post_job_parse(jobs_parse_request)



### Example

* Api Key Authentication (X-Nomad-Token):
```python
import time
import nomad_client
from nomad_client.api import jobs_api
from nomad_client.model.jobs_parse_request import JobsParseRequest
from nomad_client.model.job import Job
from pprint import pprint
# Defining the host is optional and defaults to https://127.0.0.1:4646/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = nomad_client.Configuration(
    host = "https://127.0.0.1:4646/v1"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: X-Nomad-Token
configuration.api_key['X-Nomad-Token'] = 'YOUR_API_KEY'

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['X-Nomad-Token'] = 'Bearer'

# Enter a context with an instance of the API client
with nomad_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = jobs_api.JobsApi(api_client)
    jobs_parse_request = JobsParseRequest(
        canonicalize=True,
        job_hcl="job_hcl_example",
        hclv1=True,
    ) # JobsParseRequest | 

    # example passing only required values which don't have defaults set
    try:
        api_response = api_instance.post_job_parse(jobs_parse_request)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling JobsApi->post_job_parse: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jobs_parse_request** | [**JobsParseRequest**](JobsParseRequest.md)|  |

### Return type

[**Job**](Job.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** |  |  -  |
**400** | Bad request |  -  |
**403** | Forbidden |  -  |
**405** | Method not allowed |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **post_job_periodic_force**
> PeriodicForceResponse post_job_periodic_force(job_name)



### Example

* Api Key Authentication (X-Nomad-Token):
```python
import time
import nomad_client
from nomad_client.api import jobs_api
from nomad_client.model.periodic_force_response import PeriodicForceResponse
from pprint import pprint
# Defining the host is optional and defaults to https://127.0.0.1:4646/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = nomad_client.Configuration(
    host = "https://127.0.0.1:4646/v1"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: X-Nomad-Token
configuration.api_key['X-Nomad-Token'] = 'YOUR_API_KEY'

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['X-Nomad-Token'] = 'Bearer'

# Enter a context with an instance of the API client
with nomad_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = jobs_api.JobsApi(api_client)
    job_name = "jobName_example" # str | The job identifier.
    region = "region_example" # str | Filters results based on the specified region. (optional)
    namespace = "namespace_example" # str | Filters results based on the specified namespace. (optional)
    x_nomad_token = "X-Nomad-Token_example" # str | A Nomad ACL token. (optional)
    idempotency_token = "idempotency_token_example" # str | Can be used to ensure operations are only run once. (optional)

    # example passing only required values which don't have defaults set
    try:
        api_response = api_instance.post_job_periodic_force(job_name)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling JobsApi->post_job_periodic_force: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        api_response = api_instance.post_job_periodic_force(job_name, region=region, namespace=namespace, x_nomad_token=x_nomad_token, idempotency_token=idempotency_token)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling JobsApi->post_job_periodic_force: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **job_name** | **str**| The job identifier. |
 **region** | **str**| Filters results based on the specified region. | [optional]
 **namespace** | **str**| Filters results based on the specified namespace. | [optional]
 **x_nomad_token** | **str**| A Nomad ACL token. | [optional]
 **idempotency_token** | **str**| Can be used to ensure operations are only run once. | [optional]

### Return type

[**PeriodicForceResponse**](PeriodicForceResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** |  |  * X-Nomad-Index - A unique identifier representing the current state of the requested resource. On a new Nomad cluster the value of this index starts at 1. <br>  |
**400** | Bad request |  -  |
**403** | Forbidden |  -  |
**405** | Method not allowed |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **post_job_plan**
> JobPlanResponse post_job_plan(job_name, job_plan_request)



### Example

* Api Key Authentication (X-Nomad-Token):
```python
import time
import nomad_client
from nomad_client.api import jobs_api
from nomad_client.model.job_plan_request import JobPlanRequest
from nomad_client.model.job_plan_response import JobPlanResponse
from pprint import pprint
# Defining the host is optional and defaults to https://127.0.0.1:4646/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = nomad_client.Configuration(
    host = "https://127.0.0.1:4646/v1"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: X-Nomad-Token
configuration.api_key['X-Nomad-Token'] = 'YOUR_API_KEY'

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['X-Nomad-Token'] = 'Bearer'

# Enter a context with an instance of the API client
with nomad_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = jobs_api.JobsApi(api_client)
    job_name = "jobName_example" # str | The job identifier.
    job_plan_request = JobPlanRequest(
        diff=True,
        job=Job(
            affinities=[
                Affinity(
                    l_target="l_target_example",
                    operand="operand_example",
                    r_target="r_target_example",
                    weight=-128,
                ),
            ],
            all_at_once=True,
            constraints=[
                Constraint(
                    l_target="l_target_example",
                    operand="operand_example",
                    r_target="r_target_example",
                ),
            ],
            consul_namespace="consul_namespace_example",
            consul_token="consul_token_example",
            create_index=0,
            datacenters=[
                "datacenters_example",
            ],
            dispatch_idempotency_token="dispatch_idempotency_token_example",
            dispatched=True,
            id="id_example",
            job_modify_index=0,
            meta={
                "key": "key_example",
            },
            migrate=MigrateStrategy(
                health_check="health_check_example",
                healthy_deadline=1,
                max_parallel=1,
                min_healthy_time=1,
            ),
            modify_index=0,
            multiregion=Multiregion(
                regions=[
                    MultiregionRegion(
                        count=1,
                        datacenters=[
                            "datacenters_example",
                        ],
                        meta={
                            "key": "key_example",
                        },
                        name="name_example",
                    ),
                ],
                strategy=MultiregionStrategy(
                    max_parallel=1,
                    on_failure="on_failure_example",
                ),
            ),
            name="name_example",
            namespace="namespace_example",
            nomad_token_id="nomad_token_id_example",
            parameterized_job=ParameterizedJobConfig(
                meta_optional=[
                    "meta_optional_example",
                ],
                meta_required=[
                    "meta_required_example",
                ],
                payload="payload_example",
            ),
            parent_id="parent_id_example",
            payload='YQ==',
            periodic=PeriodicConfig(
                enabled=True,
                prohibit_overlap=True,
                spec="spec_example",
                spec_type="spec_type_example",
                time_zone="time_zone_example",
            ),
            priority=1,
            region="region_example",
            reschedule=ReschedulePolicy(
                attempts=1,
                delay=1,
                delay_function="delay_function_example",
                interval=1,
                max_delay=1,
                unlimited=True,
            ),
            spreads=[
                Spread(
                    attribute="attribute_example",
                    spread_target=[
                        SpreadTarget(
                            percent=0,
                            value="value_example",
                        ),
                    ],
                    weight=-128,
                ),
            ],
            stable=True,
            status="status_example",
            status_description="status_description_example",
            stop=True,
            submit_time=1,
            task_groups=[
                TaskGroup(
                    affinities=[
                        Affinity(
                            l_target="l_target_example",
                            operand="operand_example",
                            r_target="r_target_example",
                            weight=-128,
                        ),
                    ],
                    constraints=[
                        Constraint(
                            l_target="l_target_example",
                            operand="operand_example",
                            r_target="r_target_example",
                        ),
                    ],
                    consul=Consul(
                        namespace="namespace_example",
                    ),
                    count=1,
                    ephemeral_disk=EphemeralDisk(
                        migrate=True,
                        size_mb=1,
                        sticky=True,
                    ),
                    max_client_disconnect=1,
                    meta={
                        "key": "key_example",
                    },
                    migrate=MigrateStrategy(
                        health_check="health_check_example",
                        healthy_deadline=1,
                        max_parallel=1,
                        min_healthy_time=1,
                    ),
                    name="name_example",
                    networks=[
                        NetworkResource(
                            cidr="cidr_example",
                            dns=DNSConfig(
                                options=[
                                    "options_example",
                                ],
                                searches=[
                                    "searches_example",
                                ],
                                servers=[
                                    "servers_example",
                                ],
                            ),
                            device="device_example",
                            dynamic_ports=[
                                Port(
                                    host_network="host_network_example",
                                    label="label_example",
                                    to=1,
                                    value=1,
                                ),
                            ],
                            hostname="hostname_example",
                            ip="ip_example",
                            m_bits=1,
                            mode="mode_example",
                            reserved_ports=[
                                Port(
                                    host_network="host_network_example",
                                    label="label_example",
                                    to=1,
                                    value=1,
                                ),
                            ],
                        ),
                    ],
                    reschedule_policy=ReschedulePolicy(
                        attempts=1,
                        delay=1,
                        delay_function="delay_function_example",
                        interval=1,
                        max_delay=1,
                        unlimited=True,
                    ),
                    restart_policy=RestartPolicy(
                        attempts=1,
                        delay=1,
                        interval=1,
                        mode="mode_example",
                    ),
                    scaling=ScalingPolicy(
                        create_index=0,
                        enabled=True,
                        id="id_example",
                        max=1,
                        min=1,
                        modify_index=0,
                        namespace="namespace_example",
                        policy={
                            "key": None,
                        },
                        target={
                            "key": "key_example",
                        },
                        type="type_example",
                    ),
                    services=[
                        Service(
                            address_mode="address_mode_example",
                            canary_meta={
                                "key": "key_example",
                            },
                            canary_tags=[
                                "canary_tags_example",
                            ],
                            check_restart=CheckRestart(
                                grace=1,
                                ignore_warnings=True,
                                limit=1,
                            ),
                            checks=[
                                ServiceCheck(
                                    address_mode="address_mode_example",
                                    args=[
                                        "args_example",
                                    ],
                                    body="body_example",
                                    check_restart=CheckRestart(
                                        grace=1,
                                        ignore_warnings=True,
                                        limit=1,
                                    ),
                                    command="command_example",
                                    expose=True,
                                    failures_before_critical=1,
                                    grpc_service="grpc_service_example",
                                    grpc_use_tls=True,
                                    header={
                                        "key": [
                                            "key_example",
                                        ],
                                    },
                                    id="id_example",
                                    initial_status="initial_status_example",
                                    interval=1,
                                    method="method_example",
                                    name="name_example",
                                    on_update="on_update_example",
                                    path="path_example",
                                    port_label="port_label_example",
                                    protocol="protocol_example",
                                    success_before_passing=1,
                                    tls_skip_verify=True,
                                    task_name="task_name_example",
                                    timeout=1,
                                    type="type_example",
                                ),
                            ],
                            connect=ConsulConnect(
                                gateway=ConsulGateway(
                                    ingress=ConsulIngressConfigEntry(
                                        listeners=[
                                            ConsulIngressListener(
                                                port=1,
                                                protocol="protocol_example",
                                                services=[
                                                    ConsulIngressService(
                                                        hosts=[
                                                            "hosts_example",
                                                        ],
                                                        name="name_example",
                                                    ),
                                                ],
                                            ),
                                        ],
                                        tls=ConsulGatewayTLSConfig(
                                            enabled=True,
                                        ),
                                    ),
                                    mesh=None,
                                    proxy=ConsulGatewayProxy(
                                        config={
                                            "key": None,
                                        },
                                        connect_timeout=1,
                                        envoy_dns_discovery_type="envoy_dns_discovery_type_example",
                                        envoy_gateway_bind_addresses={
                                            "key": ConsulGatewayBindAddress(
                                                address="address_example",
                                                name="name_example",
                                                port=1,
                                            ),
                                        },
                                        envoy_gateway_bind_tagged_addresses=True,
                                        envoy_gateway_no_default_bind=True,
                                    ),
                                    terminating=ConsulTerminatingConfigEntry(
                                        services=[
                                            ConsulLinkedService(
                                                ca_file="ca_file_example",
                                                cert_file="cert_file_example",
                                                key_file="key_file_example",
                                                name="name_example",
                                                sni="sni_example",
                                            ),
                                        ],
                                    ),
                                ),
                                native=True,
                                sidecar_service=ConsulSidecarService(
                                    disable_default_tcp_check=True,
                                    port="port_example",
                                    proxy=ConsulProxy(
                                        config={
                                            "key": None,
                                        },
                                        expose_config=ConsulExposeConfig(
                                            path=[
                                                ConsulExposePath(
                                                    listener_port="listener_port_example",
                                                    local_path_port=1,
                                                    path="path_example",
                                                    protocol="protocol_example",
                                                ),
                                            ],
                                        ),
                                        local_service_address="local_service_address_example",
                                        local_service_port=1,
                                        upstreams=[
                                            ConsulUpstream(
                                                datacenter="datacenter_example",
                                                destination_name="destination_name_example",
                                                local_bind_address="local_bind_address_example",
                                                local_bind_port=1,
                                                mesh_gateway=ConsulMeshGateway(
                                                    mode="mode_example",
                                                ),
                                            ),
                                        ],
                                    ),
                                    tags=[
                                        "tags_example",
                                    ],
                                ),
                                sidecar_task=SidecarTask(
                                    config={
                                        "key": None,
                                    },
                                    driver="driver_example",
                                    env={
                                        "key": "key_example",
                                    },
                                    kill_signal="kill_signal_example",
                                    kill_timeout=1,
                                    log_config=LogConfig(
                                        max_file_size_mb=1,
                                        max_files=1,
                                    ),
                                    meta={
                                        "key": "key_example",
                                    },
                                    name="name_example",
                                    resources=Resources(
                                        cpu=1,
                                        cores=1,
                                        devices=[
                                            RequestedDevice(
                                                affinities=[
                                                    Affinity(
                                                        l_target="l_target_example",
                                                        operand="operand_example",
                                                        r_target="r_target_example",
                                                        weight=-128,
                                                    ),
                                                ],
                                                constraints=[
                                                    Constraint(
                                                        l_target="l_target_example",
                                                        operand="operand_example",
                                                        r_target="r_target_example",
                                                    ),
                                                ],
                                                count=0,
                                                name="name_example",
                                            ),
                                        ],
                                        disk_mb=1,
                                        iops=1,
                                        memory_mb=1,
                                        memory_max_mb=1,
                                        networks=[
                                            NetworkResource(
                                                cidr="cidr_example",
                                                dns=DNSConfig(
                                                    options=[
                                                        "options_example",
                                                    ],
                                                    searches=[
                                                        "searches_example",
                                                    ],
                                                    servers=[
                                                        "servers_example",
                                                    ],
                                                ),
                                                device="device_example",
                                                dynamic_ports=[
                                                    Port(
                                                        host_network="host_network_example",
                                                        label="label_example",
                                                        to=1,
                                                        value=1,
                                                    ),
                                                ],
                                                hostname="hostname_example",
                                                ip="ip_example",
                                                m_bits=1,
                                                mode="mode_example",
                                                reserved_ports=[
                                                    Port(
                                                        host_network="host_network_example",
                                                        label="label_example",
                                                        to=1,
                                                        value=1,
                                                    ),
                                                ],
                                            ),
                                        ],
                                    ),
                                    shutdown_delay=1,
                                    user="user_example",
                                ),
                            ),
                            enable_tag_override=True,
                            id="id_example",
                            meta={
                                "key": "key_example",
                            },
                            name="name_example",
                            on_update="on_update_example",
                            port_label="port_label_example",
                            provider="provider_example",
                            tags=[
                                "tags_example",
                            ],
                            task_name="task_name_example",
                        ),
                    ],
                    shutdown_delay=1,
                    spreads=[
                        Spread(
                            attribute="attribute_example",
                            spread_target=[
                                SpreadTarget(
                                    percent=0,
                                    value="value_example",
                                ),
                            ],
                            weight=-128,
                        ),
                    ],
                    stop_after_client_disconnect=1,
                    tasks=[
                        Task(
                            affinities=[
                                Affinity(
                                    l_target="l_target_example",
                                    operand="operand_example",
                                    r_target="r_target_example",
                                    weight=-128,
                                ),
                            ],
                            artifacts=[
                                TaskArtifact(
                                    getter_headers={
                                        "key": "key_example",
                                    },
                                    getter_mode="getter_mode_example",
                                    getter_options={
                                        "key": "key_example",
                                    },
                                    getter_source="getter_source_example",
                                    relative_dest="relative_dest_example",
                                ),
                            ],
                            csi_plugin_config=TaskCSIPluginConfig(
                                id="id_example",
                                mount_dir="mount_dir_example",
                                type="type_example",
                            ),
                            config={
                                "key": None,
                            },
                            constraints=[
                                Constraint(
                                    l_target="l_target_example",
                                    operand="operand_example",
                                    r_target="r_target_example",
                                ),
                            ],
                            dispatch_payload=DispatchPayloadConfig(
                                file="file_example",
                            ),
                            driver="driver_example",
                            env={
                                "key": "key_example",
                            },
                            kill_signal="kill_signal_example",
                            kill_timeout=1,
                            kind="kind_example",
                            leader=True,
                            lifecycle=TaskLifecycle(
                                hook="hook_example",
                                sidecar=True,
                            ),
                            log_config=LogConfig(
                                max_file_size_mb=1,
                                max_files=1,
                            ),
                            meta={
                                "key": "key_example",
                            },
                            name="name_example",
                            resources=Resources(
                                cpu=1,
                                cores=1,
                                devices=[
                                    RequestedDevice(
                                        affinities=[
                                            Affinity(
                                                l_target="l_target_example",
                                                operand="operand_example",
                                                r_target="r_target_example",
                                                weight=-128,
                                            ),
                                        ],
                                        constraints=[
                                            Constraint(
                                                l_target="l_target_example",
                                                operand="operand_example",
                                                r_target="r_target_example",
                                            ),
                                        ],
                                        count=0,
                                        name="name_example",
                                    ),
                                ],
                                disk_mb=1,
                                iops=1,
                                memory_mb=1,
                                memory_max_mb=1,
                                networks=[
                                    NetworkResource(
                                        cidr="cidr_example",
                                        dns=DNSConfig(
                                            options=[
                                                "options_example",
                                            ],
                                            searches=[
                                                "searches_example",
                                            ],
                                            servers=[
                                                "servers_example",
                                            ],
                                        ),
                                        device="device_example",
                                        dynamic_ports=[
                                            Port(
                                                host_network="host_network_example",
                                                label="label_example",
                                                to=1,
                                                value=1,
                                            ),
                                        ],
                                        hostname="hostname_example",
                                        ip="ip_example",
                                        m_bits=1,
                                        mode="mode_example",
                                        reserved_ports=[
                                            Port(
                                                host_network="host_network_example",
                                                label="label_example",
                                                to=1,
                                                value=1,
                                            ),
                                        ],
                                    ),
                                ],
                            ),
                            restart_policy=RestartPolicy(
                                attempts=1,
                                delay=1,
                                interval=1,
                                mode="mode_example",
                            ),
                            scaling_policies=[
                                ScalingPolicy(
                                    create_index=0,
                                    enabled=True,
                                    id="id_example",
                                    max=1,
                                    min=1,
                                    modify_index=0,
                                    namespace="namespace_example",
                                    policy={
                                        "key": None,
                                    },
                                    target={
                                        "key": "key_example",
                                    },
                                    type="type_example",
                                ),
                            ],
                            services=[
                                Service(
                                    address_mode="address_mode_example",
                                    canary_meta={
                                        "key": "key_example",
                                    },
                                    canary_tags=[
                                        "canary_tags_example",
                                    ],
                                    check_restart=CheckRestart(
                                        grace=1,
                                        ignore_warnings=True,
                                        limit=1,
                                    ),
                                    checks=[
                                        ServiceCheck(
                                            address_mode="address_mode_example",
                                            args=[
                                                "args_example",
                                            ],
                                            body="body_example",
                                            check_restart=CheckRestart(
                                                grace=1,
                                                ignore_warnings=True,
                                                limit=1,
                                            ),
                                            command="command_example",
                                            expose=True,
                                            failures_before_critical=1,
                                            grpc_service="grpc_service_example",
                                            grpc_use_tls=True,
                                            header={
                                                "key": [
                                                    "key_example",
                                                ],
                                            },
                                            id="id_example",
                                            initial_status="initial_status_example",
                                            interval=1,
                                            method="method_example",
                                            name="name_example",
                                            on_update="on_update_example",
                                            path="path_example",
                                            port_label="port_label_example",
                                            protocol="protocol_example",
                                            success_before_passing=1,
                                            tls_skip_verify=True,
                                            task_name="task_name_example",
                                            timeout=1,
                                            type="type_example",
                                        ),
                                    ],
                                    connect=ConsulConnect(
                                        gateway=ConsulGateway(
                                            ingress=ConsulIngressConfigEntry(
                                                listeners=[
                                                    ConsulIngressListener(
                                                        port=1,
                                                        protocol="protocol_example",
                                                        services=[
                                                            ConsulIngressService(
                                                                hosts=[
                                                                    "hosts_example",
                                                                ],
                                                                name="name_example",
                                                            ),
                                                        ],
                                                    ),
                                                ],
                                                tls=ConsulGatewayTLSConfig(
                                                    enabled=True,
                                                ),
                                            ),
                                            mesh=None,
                                            proxy=ConsulGatewayProxy(
                                                config={
                                                    "key": None,
                                                },
                                                connect_timeout=1,
                                                envoy_dns_discovery_type="envoy_dns_discovery_type_example",
                                                envoy_gateway_bind_addresses={
                                                    "key": ConsulGatewayBindAddress(
                                                        address="address_example",
                                                        name="name_example",
                                                        port=1,
                                                    ),
                                                },
                                                envoy_gateway_bind_tagged_addresses=True,
                                                envoy_gateway_no_default_bind=True,
                                            ),
                                            terminating=ConsulTerminatingConfigEntry(
                                                services=[
                                                    ConsulLinkedService(
                                                        ca_file="ca_file_example",
                                                        cert_file="cert_file_example",
                                                        key_file="key_file_example",
                                                        name="name_example",
                                                        sni="sni_example",
                                                    ),
                                                ],
                                            ),
                                        ),
                                        native=True,
                                        sidecar_service=ConsulSidecarService(
                                            disable_default_tcp_check=True,
                                            port="port_example",
                                            proxy=ConsulProxy(
                                                config={
                                                    "key": None,
                                                },
                                                expose_config=ConsulExposeConfig(
                                                    path=[
                                                        ConsulExposePath(
                                                            listener_port="listener_port_example",
                                                            local_path_port=1,
                                                            path="path_example",
                                                            protocol="protocol_example",
                                                        ),
                                                    ],
                                                ),
                                                local_service_address="local_service_address_example",
                                                local_service_port=1,
                                                upstreams=[
                                                    ConsulUpstream(
                                                        datacenter="datacenter_example",
                                                        destination_name="destination_name_example",
                                                        local_bind_address="local_bind_address_example",
                                                        local_bind_port=1,
                                                        mesh_gateway=ConsulMeshGateway(
                                                            mode="mode_example",
                                                        ),
                                                    ),
                                                ],
                                            ),
                                            tags=[
                                                "tags_example",
                                            ],
                                        ),
                                        sidecar_task=SidecarTask(
                                            config={
                                                "key": None,
                                            },
                                            driver="driver_example",
                                            env={
                                                "key": "key_example",
                                            },
                                            kill_signal="kill_signal_example",
                                            kill_timeout=1,
                                            log_config=LogConfig(
                                                max_file_size_mb=1,
                                                max_files=1,
                                            ),
                                            meta={
                                                "key": "key_example",
                                            },
                                            name="name_example",
                                            resources=Resources(
                                                cpu=1,
                                                cores=1,
                                                devices=[
                                                    RequestedDevice(
                                                        affinities=[
                                                            Affinity(
                                                                l_target="l_target_example",
                                                                operand="operand_example",
                                                                r_target="r_target_example",
                                                                weight=-128,
                                                            ),
                                                        ],
                                                        constraints=[
                                                            Constraint(
                                                                l_target="l_target_example",
                                                                operand="operand_example",
                                                                r_target="r_target_example",
                                                            ),
                                                        ],
                                                        count=0,
                                                        name="name_example",
                                                    ),
                                                ],
                                                disk_mb=1,
                                                iops=1,
                                                memory_mb=1,
                                                memory_max_mb=1,
                                                networks=[
                                                    NetworkResource(
                                                        cidr="cidr_example",
                                                        dns=DNSConfig(
                                                            options=[
                                                                "options_example",
                                                            ],
                                                            searches=[
                                                                "searches_example",
                                                            ],
                                                            servers=[
                                                                "servers_example",
                                                            ],
                                                        ),
                                                        device="device_example",
                                                        dynamic_ports=[
                                                            Port(
                                                                host_network="host_network_example",
                                                                label="label_example",
                                                                to=1,
                                                                value=1,
                                                            ),
                                                        ],
                                                        hostname="hostname_example",
                                                        ip="ip_example",
                                                        m_bits=1,
                                                        mode="mode_example",
                                                        reserved_ports=[
                                                            Port(
                                                                host_network="host_network_example",
                                                                label="label_example",
                                                                to=1,
                                                                value=1,
                                                            ),
                                                        ],
                                                    ),
                                                ],
                                            ),
                                            shutdown_delay=1,
                                            user="user_example",
                                        ),
                                    ),
                                    enable_tag_override=True,
                                    id="id_example",
                                    meta={
                                        "key": "key_example",
                                    },
                                    name="name_example",
                                    on_update="on_update_example",
                                    port_label="port_label_example",
                                    provider="provider_example",
                                    tags=[
                                        "tags_example",
                                    ],
                                    task_name="task_name_example",
                                ),
                            ],
                            shutdown_delay=1,
                            templates=[
                                Template(
                                    change_mode="change_mode_example",
                                    change_signal="change_signal_example",
                                    dest_path="dest_path_example",
                                    embedded_tmpl="embedded_tmpl_example",
                                    envvars=True,
                                    left_delim="left_delim_example",
                                    perms="perms_example",
                                    right_delim="right_delim_example",
                                    source_path="source_path_example",
                                    splay=1,
                                    vault_grace=1,
                                    wait=WaitConfig(
                                        max=1,
                                        min=1,
                                    ),
                                ),
                            ],
                            user="user_example",
                            vault=Vault(
                                change_mode="change_mode_example",
                                change_signal="change_signal_example",
                                entity_alias="entity_alias_example",
                                env=True,
                                namespace="namespace_example",
                                policies=[
                                    "policies_example",
                                ],
                            ),
                            volume_mounts=[
                                VolumeMount(
                                    destination="destination_example",
                                    propagation_mode="propagation_mode_example",
                                    read_only=True,
                                    volume="volume_example",
                                ),
                            ],
                        ),
                    ],
                    update=UpdateStrategy(
                        auto_promote=True,
                        auto_revert=True,
                        canary=1,
                        health_check="health_check_example",
                        healthy_deadline=1,
                        max_parallel=1,
                        min_healthy_time=1,
                        progress_deadline=1,
                        stagger=1,
                    ),
                    volumes={
                        "key": VolumeRequest(
                            access_mode="access_mode_example",
                            attachment_mode="attachment_mode_example",
                            mount_options=CSIMountOptions(
                                fs_type="fs_type_example",
                                mount_flags=[
                                    "mount_flags_example",
                                ],
                            ),
                            name="name_example",
                            per_alloc=True,
                            read_only=True,
                            source="source_example",
                            type="type_example",
                        ),
                    },
                ),
            ],
            type="type_example",
            update=UpdateStrategy(
                auto_promote=True,
                auto_revert=True,
                canary=1,
                health_check="health_check_example",
                healthy_deadline=1,
                max_parallel=1,
                min_healthy_time=1,
                progress_deadline=1,
                stagger=1,
            ),
            vault_namespace="vault_namespace_example",
            vault_token="vault_token_example",
            version=0,
        ),
        namespace="namespace_example",
        policy_override=True,
        region="region_example",
        secret_id="secret_id_example",
    ) # JobPlanRequest | 
    region = "region_example" # str | Filters results based on the specified region. (optional)
    namespace = "namespace_example" # str | Filters results based on the specified namespace. (optional)
    x_nomad_token = "X-Nomad-Token_example" # str | A Nomad ACL token. (optional)
    idempotency_token = "idempotency_token_example" # str | Can be used to ensure operations are only run once. (optional)

    # example passing only required values which don't have defaults set
    try:
        api_response = api_instance.post_job_plan(job_name, job_plan_request)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling JobsApi->post_job_plan: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        api_response = api_instance.post_job_plan(job_name, job_plan_request, region=region, namespace=namespace, x_nomad_token=x_nomad_token, idempotency_token=idempotency_token)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling JobsApi->post_job_plan: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **job_name** | **str**| The job identifier. |
 **job_plan_request** | [**JobPlanRequest**](JobPlanRequest.md)|  |
 **region** | **str**| Filters results based on the specified region. | [optional]
 **namespace** | **str**| Filters results based on the specified namespace. | [optional]
 **x_nomad_token** | **str**| A Nomad ACL token. | [optional]
 **idempotency_token** | **str**| Can be used to ensure operations are only run once. | [optional]

### Return type

[**JobPlanResponse**](JobPlanResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** |  |  * X-Nomad-Index - A unique identifier representing the current state of the requested resource. On a new Nomad cluster the value of this index starts at 1. <br>  |
**400** | Bad request |  -  |
**403** | Forbidden |  -  |
**405** | Method not allowed |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **post_job_revert**
> JobRegisterResponse post_job_revert(job_name, job_revert_request)



### Example

* Api Key Authentication (X-Nomad-Token):
```python
import time
import nomad_client
from nomad_client.api import jobs_api
from nomad_client.model.job_register_response import JobRegisterResponse
from nomad_client.model.job_revert_request import JobRevertRequest
from pprint import pprint
# Defining the host is optional and defaults to https://127.0.0.1:4646/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = nomad_client.Configuration(
    host = "https://127.0.0.1:4646/v1"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: X-Nomad-Token
configuration.api_key['X-Nomad-Token'] = 'YOUR_API_KEY'

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['X-Nomad-Token'] = 'Bearer'

# Enter a context with an instance of the API client
with nomad_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = jobs_api.JobsApi(api_client)
    job_name = "jobName_example" # str | The job identifier.
    job_revert_request = JobRevertRequest(
        consul_token="consul_token_example",
        enforce_prior_version=0,
        job_id="job_id_example",
        job_version=0,
        namespace="namespace_example",
        region="region_example",
        secret_id="secret_id_example",
        vault_token="vault_token_example",
    ) # JobRevertRequest | 
    region = "region_example" # str | Filters results based on the specified region. (optional)
    namespace = "namespace_example" # str | Filters results based on the specified namespace. (optional)
    x_nomad_token = "X-Nomad-Token_example" # str | A Nomad ACL token. (optional)
    idempotency_token = "idempotency_token_example" # str | Can be used to ensure operations are only run once. (optional)

    # example passing only required values which don't have defaults set
    try:
        api_response = api_instance.post_job_revert(job_name, job_revert_request)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling JobsApi->post_job_revert: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        api_response = api_instance.post_job_revert(job_name, job_revert_request, region=region, namespace=namespace, x_nomad_token=x_nomad_token, idempotency_token=idempotency_token)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling JobsApi->post_job_revert: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **job_name** | **str**| The job identifier. |
 **job_revert_request** | [**JobRevertRequest**](JobRevertRequest.md)|  |
 **region** | **str**| Filters results based on the specified region. | [optional]
 **namespace** | **str**| Filters results based on the specified namespace. | [optional]
 **x_nomad_token** | **str**| A Nomad ACL token. | [optional]
 **idempotency_token** | **str**| Can be used to ensure operations are only run once. | [optional]

### Return type

[**JobRegisterResponse**](JobRegisterResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** |  |  * X-Nomad-Index - A unique identifier representing the current state of the requested resource. On a new Nomad cluster the value of this index starts at 1. <br>  |
**400** | Bad request |  -  |
**403** | Forbidden |  -  |
**405** | Method not allowed |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **post_job_scaling_request**
> JobRegisterResponse post_job_scaling_request(job_name, scaling_request)



### Example

* Api Key Authentication (X-Nomad-Token):
```python
import time
import nomad_client
from nomad_client.api import jobs_api
from nomad_client.model.job_register_response import JobRegisterResponse
from nomad_client.model.scaling_request import ScalingRequest
from pprint import pprint
# Defining the host is optional and defaults to https://127.0.0.1:4646/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = nomad_client.Configuration(
    host = "https://127.0.0.1:4646/v1"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: X-Nomad-Token
configuration.api_key['X-Nomad-Token'] = 'YOUR_API_KEY'

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['X-Nomad-Token'] = 'Bearer'

# Enter a context with an instance of the API client
with nomad_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = jobs_api.JobsApi(api_client)
    job_name = "jobName_example" # str | The job identifier.
    scaling_request = ScalingRequest(
        count=1,
        error=True,
        message="message_example",
        meta={
            "key": None,
        },
        namespace="namespace_example",
        policy_override=True,
        region="region_example",
        secret_id="secret_id_example",
        target={
            "key": "key_example",
        },
    ) # ScalingRequest | 
    region = "region_example" # str | Filters results based on the specified region. (optional)
    namespace = "namespace_example" # str | Filters results based on the specified namespace. (optional)
    x_nomad_token = "X-Nomad-Token_example" # str | A Nomad ACL token. (optional)
    idempotency_token = "idempotency_token_example" # str | Can be used to ensure operations are only run once. (optional)

    # example passing only required values which don't have defaults set
    try:
        api_response = api_instance.post_job_scaling_request(job_name, scaling_request)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling JobsApi->post_job_scaling_request: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        api_response = api_instance.post_job_scaling_request(job_name, scaling_request, region=region, namespace=namespace, x_nomad_token=x_nomad_token, idempotency_token=idempotency_token)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling JobsApi->post_job_scaling_request: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **job_name** | **str**| The job identifier. |
 **scaling_request** | [**ScalingRequest**](ScalingRequest.md)|  |
 **region** | **str**| Filters results based on the specified region. | [optional]
 **namespace** | **str**| Filters results based on the specified namespace. | [optional]
 **x_nomad_token** | **str**| A Nomad ACL token. | [optional]
 **idempotency_token** | **str**| Can be used to ensure operations are only run once. | [optional]

### Return type

[**JobRegisterResponse**](JobRegisterResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** |  |  * X-Nomad-Index - A unique identifier representing the current state of the requested resource. On a new Nomad cluster the value of this index starts at 1. <br>  |
**400** | Bad request |  -  |
**403** | Forbidden |  -  |
**405** | Method not allowed |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **post_job_stability**
> JobStabilityResponse post_job_stability(job_name, job_stability_request)



### Example

* Api Key Authentication (X-Nomad-Token):
```python
import time
import nomad_client
from nomad_client.api import jobs_api
from nomad_client.model.job_stability_response import JobStabilityResponse
from nomad_client.model.job_stability_request import JobStabilityRequest
from pprint import pprint
# Defining the host is optional and defaults to https://127.0.0.1:4646/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = nomad_client.Configuration(
    host = "https://127.0.0.1:4646/v1"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: X-Nomad-Token
configuration.api_key['X-Nomad-Token'] = 'YOUR_API_KEY'

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['X-Nomad-Token'] = 'Bearer'

# Enter a context with an instance of the API client
with nomad_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = jobs_api.JobsApi(api_client)
    job_name = "jobName_example" # str | The job identifier.
    job_stability_request = JobStabilityRequest(
        job_id="job_id_example",
        job_version=0,
        namespace="namespace_example",
        region="region_example",
        secret_id="secret_id_example",
        stable=True,
    ) # JobStabilityRequest | 
    region = "region_example" # str | Filters results based on the specified region. (optional)
    namespace = "namespace_example" # str | Filters results based on the specified namespace. (optional)
    x_nomad_token = "X-Nomad-Token_example" # str | A Nomad ACL token. (optional)
    idempotency_token = "idempotency_token_example" # str | Can be used to ensure operations are only run once. (optional)

    # example passing only required values which don't have defaults set
    try:
        api_response = api_instance.post_job_stability(job_name, job_stability_request)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling JobsApi->post_job_stability: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        api_response = api_instance.post_job_stability(job_name, job_stability_request, region=region, namespace=namespace, x_nomad_token=x_nomad_token, idempotency_token=idempotency_token)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling JobsApi->post_job_stability: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **job_name** | **str**| The job identifier. |
 **job_stability_request** | [**JobStabilityRequest**](JobStabilityRequest.md)|  |
 **region** | **str**| Filters results based on the specified region. | [optional]
 **namespace** | **str**| Filters results based on the specified namespace. | [optional]
 **x_nomad_token** | **str**| A Nomad ACL token. | [optional]
 **idempotency_token** | **str**| Can be used to ensure operations are only run once. | [optional]

### Return type

[**JobStabilityResponse**](JobStabilityResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** |  |  * X-Nomad-Index - A unique identifier representing the current state of the requested resource. On a new Nomad cluster the value of this index starts at 1. <br>  |
**400** | Bad request |  -  |
**403** | Forbidden |  -  |
**405** | Method not allowed |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **post_job_validate_request**
> JobValidateResponse post_job_validate_request(job_validate_request)



### Example

* Api Key Authentication (X-Nomad-Token):
```python
import time
import nomad_client
from nomad_client.api import jobs_api
from nomad_client.model.job_validate_request import JobValidateRequest
from nomad_client.model.job_validate_response import JobValidateResponse
from pprint import pprint
# Defining the host is optional and defaults to https://127.0.0.1:4646/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = nomad_client.Configuration(
    host = "https://127.0.0.1:4646/v1"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: X-Nomad-Token
configuration.api_key['X-Nomad-Token'] = 'YOUR_API_KEY'

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['X-Nomad-Token'] = 'Bearer'

# Enter a context with an instance of the API client
with nomad_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = jobs_api.JobsApi(api_client)
    job_validate_request = JobValidateRequest(
        job=Job(
            affinities=[
                Affinity(
                    l_target="l_target_example",
                    operand="operand_example",
                    r_target="r_target_example",
                    weight=-128,
                ),
            ],
            all_at_once=True,
            constraints=[
                Constraint(
                    l_target="l_target_example",
                    operand="operand_example",
                    r_target="r_target_example",
                ),
            ],
            consul_namespace="consul_namespace_example",
            consul_token="consul_token_example",
            create_index=0,
            datacenters=[
                "datacenters_example",
            ],
            dispatch_idempotency_token="dispatch_idempotency_token_example",
            dispatched=True,
            id="id_example",
            job_modify_index=0,
            meta={
                "key": "key_example",
            },
            migrate=MigrateStrategy(
                health_check="health_check_example",
                healthy_deadline=1,
                max_parallel=1,
                min_healthy_time=1,
            ),
            modify_index=0,
            multiregion=Multiregion(
                regions=[
                    MultiregionRegion(
                        count=1,
                        datacenters=[
                            "datacenters_example",
                        ],
                        meta={
                            "key": "key_example",
                        },
                        name="name_example",
                    ),
                ],
                strategy=MultiregionStrategy(
                    max_parallel=1,
                    on_failure="on_failure_example",
                ),
            ),
            name="name_example",
            namespace="namespace_example",
            nomad_token_id="nomad_token_id_example",
            parameterized_job=ParameterizedJobConfig(
                meta_optional=[
                    "meta_optional_example",
                ],
                meta_required=[
                    "meta_required_example",
                ],
                payload="payload_example",
            ),
            parent_id="parent_id_example",
            payload='YQ==',
            periodic=PeriodicConfig(
                enabled=True,
                prohibit_overlap=True,
                spec="spec_example",
                spec_type="spec_type_example",
                time_zone="time_zone_example",
            ),
            priority=1,
            region="region_example",
            reschedule=ReschedulePolicy(
                attempts=1,
                delay=1,
                delay_function="delay_function_example",
                interval=1,
                max_delay=1,
                unlimited=True,
            ),
            spreads=[
                Spread(
                    attribute="attribute_example",
                    spread_target=[
                        SpreadTarget(
                            percent=0,
                            value="value_example",
                        ),
                    ],
                    weight=-128,
                ),
            ],
            stable=True,
            status="status_example",
            status_description="status_description_example",
            stop=True,
            submit_time=1,
            task_groups=[
                TaskGroup(
                    affinities=[
                        Affinity(
                            l_target="l_target_example",
                            operand="operand_example",
                            r_target="r_target_example",
                            weight=-128,
                        ),
                    ],
                    constraints=[
                        Constraint(
                            l_target="l_target_example",
                            operand="operand_example",
                            r_target="r_target_example",
                        ),
                    ],
                    consul=Consul(
                        namespace="namespace_example",
                    ),
                    count=1,
                    ephemeral_disk=EphemeralDisk(
                        migrate=True,
                        size_mb=1,
                        sticky=True,
                    ),
                    max_client_disconnect=1,
                    meta={
                        "key": "key_example",
                    },
                    migrate=MigrateStrategy(
                        health_check="health_check_example",
                        healthy_deadline=1,
                        max_parallel=1,
                        min_healthy_time=1,
                    ),
                    name="name_example",
                    networks=[
                        NetworkResource(
                            cidr="cidr_example",
                            dns=DNSConfig(
                                options=[
                                    "options_example",
                                ],
                                searches=[
                                    "searches_example",
                                ],
                                servers=[
                                    "servers_example",
                                ],
                            ),
                            device="device_example",
                            dynamic_ports=[
                                Port(
                                    host_network="host_network_example",
                                    label="label_example",
                                    to=1,
                                    value=1,
                                ),
                            ],
                            hostname="hostname_example",
                            ip="ip_example",
                            m_bits=1,
                            mode="mode_example",
                            reserved_ports=[
                                Port(
                                    host_network="host_network_example",
                                    label="label_example",
                                    to=1,
                                    value=1,
                                ),
                            ],
                        ),
                    ],
                    reschedule_policy=ReschedulePolicy(
                        attempts=1,
                        delay=1,
                        delay_function="delay_function_example",
                        interval=1,
                        max_delay=1,
                        unlimited=True,
                    ),
                    restart_policy=RestartPolicy(
                        attempts=1,
                        delay=1,
                        interval=1,
                        mode="mode_example",
                    ),
                    scaling=ScalingPolicy(
                        create_index=0,
                        enabled=True,
                        id="id_example",
                        max=1,
                        min=1,
                        modify_index=0,
                        namespace="namespace_example",
                        policy={
                            "key": None,
                        },
                        target={
                            "key": "key_example",
                        },
                        type="type_example",
                    ),
                    services=[
                        Service(
                            address_mode="address_mode_example",
                            canary_meta={
                                "key": "key_example",
                            },
                            canary_tags=[
                                "canary_tags_example",
                            ],
                            check_restart=CheckRestart(
                                grace=1,
                                ignore_warnings=True,
                                limit=1,
                            ),
                            checks=[
                                ServiceCheck(
                                    address_mode="address_mode_example",
                                    args=[
                                        "args_example",
                                    ],
                                    body="body_example",
                                    check_restart=CheckRestart(
                                        grace=1,
                                        ignore_warnings=True,
                                        limit=1,
                                    ),
                                    command="command_example",
                                    expose=True,
                                    failures_before_critical=1,
                                    grpc_service="grpc_service_example",
                                    grpc_use_tls=True,
                                    header={
                                        "key": [
                                            "key_example",
                                        ],
                                    },
                                    id="id_example",
                                    initial_status="initial_status_example",
                                    interval=1,
                                    method="method_example",
                                    name="name_example",
                                    on_update="on_update_example",
                                    path="path_example",
                                    port_label="port_label_example",
                                    protocol="protocol_example",
                                    success_before_passing=1,
                                    tls_skip_verify=True,
                                    task_name="task_name_example",
                                    timeout=1,
                                    type="type_example",
                                ),
                            ],
                            connect=ConsulConnect(
                                gateway=ConsulGateway(
                                    ingress=ConsulIngressConfigEntry(
                                        listeners=[
                                            ConsulIngressListener(
                                                port=1,
                                                protocol="protocol_example",
                                                services=[
                                                    ConsulIngressService(
                                                        hosts=[
                                                            "hosts_example",
                                                        ],
                                                        name="name_example",
                                                    ),
                                                ],
                                            ),
                                        ],
                                        tls=ConsulGatewayTLSConfig(
                                            enabled=True,
                                        ),
                                    ),
                                    mesh=None,
                                    proxy=ConsulGatewayProxy(
                                        config={
                                            "key": None,
                                        },
                                        connect_timeout=1,
                                        envoy_dns_discovery_type="envoy_dns_discovery_type_example",
                                        envoy_gateway_bind_addresses={
                                            "key": ConsulGatewayBindAddress(
                                                address="address_example",
                                                name="name_example",
                                                port=1,
                                            ),
                                        },
                                        envoy_gateway_bind_tagged_addresses=True,
                                        envoy_gateway_no_default_bind=True,
                                    ),
                                    terminating=ConsulTerminatingConfigEntry(
                                        services=[
                                            ConsulLinkedService(
                                                ca_file="ca_file_example",
                                                cert_file="cert_file_example",
                                                key_file="key_file_example",
                                                name="name_example",
                                                sni="sni_example",
                                            ),
                                        ],
                                    ),
                                ),
                                native=True,
                                sidecar_service=ConsulSidecarService(
                                    disable_default_tcp_check=True,
                                    port="port_example",
                                    proxy=ConsulProxy(
                                        config={
                                            "key": None,
                                        },
                                        expose_config=ConsulExposeConfig(
                                            path=[
                                                ConsulExposePath(
                                                    listener_port="listener_port_example",
                                                    local_path_port=1,
                                                    path="path_example",
                                                    protocol="protocol_example",
                                                ),
                                            ],
                                        ),
                                        local_service_address="local_service_address_example",
                                        local_service_port=1,
                                        upstreams=[
                                            ConsulUpstream(
                                                datacenter="datacenter_example",
                                                destination_name="destination_name_example",
                                                local_bind_address="local_bind_address_example",
                                                local_bind_port=1,
                                                mesh_gateway=ConsulMeshGateway(
                                                    mode="mode_example",
                                                ),
                                            ),
                                        ],
                                    ),
                                    tags=[
                                        "tags_example",
                                    ],
                                ),
                                sidecar_task=SidecarTask(
                                    config={
                                        "key": None,
                                    },
                                    driver="driver_example",
                                    env={
                                        "key": "key_example",
                                    },
                                    kill_signal="kill_signal_example",
                                    kill_timeout=1,
                                    log_config=LogConfig(
                                        max_file_size_mb=1,
                                        max_files=1,
                                    ),
                                    meta={
                                        "key": "key_example",
                                    },
                                    name="name_example",
                                    resources=Resources(
                                        cpu=1,
                                        cores=1,
                                        devices=[
                                            RequestedDevice(
                                                affinities=[
                                                    Affinity(
                                                        l_target="l_target_example",
                                                        operand="operand_example",
                                                        r_target="r_target_example",
                                                        weight=-128,
                                                    ),
                                                ],
                                                constraints=[
                                                    Constraint(
                                                        l_target="l_target_example",
                                                        operand="operand_example",
                                                        r_target="r_target_example",
                                                    ),
                                                ],
                                                count=0,
                                                name="name_example",
                                            ),
                                        ],
                                        disk_mb=1,
                                        iops=1,
                                        memory_mb=1,
                                        memory_max_mb=1,
                                        networks=[
                                            NetworkResource(
                                                cidr="cidr_example",
                                                dns=DNSConfig(
                                                    options=[
                                                        "options_example",
                                                    ],
                                                    searches=[
                                                        "searches_example",
                                                    ],
                                                    servers=[
                                                        "servers_example",
                                                    ],
                                                ),
                                                device="device_example",
                                                dynamic_ports=[
                                                    Port(
                                                        host_network="host_network_example",
                                                        label="label_example",
                                                        to=1,
                                                        value=1,
                                                    ),
                                                ],
                                                hostname="hostname_example",
                                                ip="ip_example",
                                                m_bits=1,
                                                mode="mode_example",
                                                reserved_ports=[
                                                    Port(
                                                        host_network="host_network_example",
                                                        label="label_example",
                                                        to=1,
                                                        value=1,
                                                    ),
                                                ],
                                            ),
                                        ],
                                    ),
                                    shutdown_delay=1,
                                    user="user_example",
                                ),
                            ),
                            enable_tag_override=True,
                            id="id_example",
                            meta={
                                "key": "key_example",
                            },
                            name="name_example",
                            on_update="on_update_example",
                            port_label="port_label_example",
                            provider="provider_example",
                            tags=[
                                "tags_example",
                            ],
                            task_name="task_name_example",
                        ),
                    ],
                    shutdown_delay=1,
                    spreads=[
                        Spread(
                            attribute="attribute_example",
                            spread_target=[
                                SpreadTarget(
                                    percent=0,
                                    value="value_example",
                                ),
                            ],
                            weight=-128,
                        ),
                    ],
                    stop_after_client_disconnect=1,
                    tasks=[
                        Task(
                            affinities=[
                                Affinity(
                                    l_target="l_target_example",
                                    operand="operand_example",
                                    r_target="r_target_example",
                                    weight=-128,
                                ),
                            ],
                            artifacts=[
                                TaskArtifact(
                                    getter_headers={
                                        "key": "key_example",
                                    },
                                    getter_mode="getter_mode_example",
                                    getter_options={
                                        "key": "key_example",
                                    },
                                    getter_source="getter_source_example",
                                    relative_dest="relative_dest_example",
                                ),
                            ],
                            csi_plugin_config=TaskCSIPluginConfig(
                                id="id_example",
                                mount_dir="mount_dir_example",
                                type="type_example",
                            ),
                            config={
                                "key": None,
                            },
                            constraints=[
                                Constraint(
                                    l_target="l_target_example",
                                    operand="operand_example",
                                    r_target="r_target_example",
                                ),
                            ],
                            dispatch_payload=DispatchPayloadConfig(
                                file="file_example",
                            ),
                            driver="driver_example",
                            env={
                                "key": "key_example",
                            },
                            kill_signal="kill_signal_example",
                            kill_timeout=1,
                            kind="kind_example",
                            leader=True,
                            lifecycle=TaskLifecycle(
                                hook="hook_example",
                                sidecar=True,
                            ),
                            log_config=LogConfig(
                                max_file_size_mb=1,
                                max_files=1,
                            ),
                            meta={
                                "key": "key_example",
                            },
                            name="name_example",
                            resources=Resources(
                                cpu=1,
                                cores=1,
                                devices=[
                                    RequestedDevice(
                                        affinities=[
                                            Affinity(
                                                l_target="l_target_example",
                                                operand="operand_example",
                                                r_target="r_target_example",
                                                weight=-128,
                                            ),
                                        ],
                                        constraints=[
                                            Constraint(
                                                l_target="l_target_example",
                                                operand="operand_example",
                                                r_target="r_target_example",
                                            ),
                                        ],
                                        count=0,
                                        name="name_example",
                                    ),
                                ],
                                disk_mb=1,
                                iops=1,
                                memory_mb=1,
                                memory_max_mb=1,
                                networks=[
                                    NetworkResource(
                                        cidr="cidr_example",
                                        dns=DNSConfig(
                                            options=[
                                                "options_example",
                                            ],
                                            searches=[
                                                "searches_example",
                                            ],
                                            servers=[
                                                "servers_example",
                                            ],
                                        ),
                                        device="device_example",
                                        dynamic_ports=[
                                            Port(
                                                host_network="host_network_example",
                                                label="label_example",
                                                to=1,
                                                value=1,
                                            ),
                                        ],
                                        hostname="hostname_example",
                                        ip="ip_example",
                                        m_bits=1,
                                        mode="mode_example",
                                        reserved_ports=[
                                            Port(
                                                host_network="host_network_example",
                                                label="label_example",
                                                to=1,
                                                value=1,
                                            ),
                                        ],
                                    ),
                                ],
                            ),
                            restart_policy=RestartPolicy(
                                attempts=1,
                                delay=1,
                                interval=1,
                                mode="mode_example",
                            ),
                            scaling_policies=[
                                ScalingPolicy(
                                    create_index=0,
                                    enabled=True,
                                    id="id_example",
                                    max=1,
                                    min=1,
                                    modify_index=0,
                                    namespace="namespace_example",
                                    policy={
                                        "key": None,
                                    },
                                    target={
                                        "key": "key_example",
                                    },
                                    type="type_example",
                                ),
                            ],
                            services=[
                                Service(
                                    address_mode="address_mode_example",
                                    canary_meta={
                                        "key": "key_example",
                                    },
                                    canary_tags=[
                                        "canary_tags_example",
                                    ],
                                    check_restart=CheckRestart(
                                        grace=1,
                                        ignore_warnings=True,
                                        limit=1,
                                    ),
                                    checks=[
                                        ServiceCheck(
                                            address_mode="address_mode_example",
                                            args=[
                                                "args_example",
                                            ],
                                            body="body_example",
                                            check_restart=CheckRestart(
                                                grace=1,
                                                ignore_warnings=True,
                                                limit=1,
                                            ),
                                            command="command_example",
                                            expose=True,
                                            failures_before_critical=1,
                                            grpc_service="grpc_service_example",
                                            grpc_use_tls=True,
                                            header={
                                                "key": [
                                                    "key_example",
                                                ],
                                            },
                                            id="id_example",
                                            initial_status="initial_status_example",
                                            interval=1,
                                            method="method_example",
                                            name="name_example",
                                            on_update="on_update_example",
                                            path="path_example",
                                            port_label="port_label_example",
                                            protocol="protocol_example",
                                            success_before_passing=1,
                                            tls_skip_verify=True,
                                            task_name="task_name_example",
                                            timeout=1,
                                            type="type_example",
                                        ),
                                    ],
                                    connect=ConsulConnect(
                                        gateway=ConsulGateway(
                                            ingress=ConsulIngressConfigEntry(
                                                listeners=[
                                                    ConsulIngressListener(
                                                        port=1,
                                                        protocol="protocol_example",
                                                        services=[
                                                            ConsulIngressService(
                                                                hosts=[
                                                                    "hosts_example",
                                                                ],
                                                                name="name_example",
                                                            ),
                                                        ],
                                                    ),
                                                ],
                                                tls=ConsulGatewayTLSConfig(
                                                    enabled=True,
                                                ),
                                            ),
                                            mesh=None,
                                            proxy=ConsulGatewayProxy(
                                                config={
                                                    "key": None,
                                                },
                                                connect_timeout=1,
                                                envoy_dns_discovery_type="envoy_dns_discovery_type_example",
                                                envoy_gateway_bind_addresses={
                                                    "key": ConsulGatewayBindAddress(
                                                        address="address_example",
                                                        name="name_example",
                                                        port=1,
                                                    ),
                                                },
                                                envoy_gateway_bind_tagged_addresses=True,
                                                envoy_gateway_no_default_bind=True,
                                            ),
                                            terminating=ConsulTerminatingConfigEntry(
                                                services=[
                                                    ConsulLinkedService(
                                                        ca_file="ca_file_example",
                                                        cert_file="cert_file_example",
                                                        key_file="key_file_example",
                                                        name="name_example",
                                                        sni="sni_example",
                                                    ),
                                                ],
                                            ),
                                        ),
                                        native=True,
                                        sidecar_service=ConsulSidecarService(
                                            disable_default_tcp_check=True,
                                            port="port_example",
                                            proxy=ConsulProxy(
                                                config={
                                                    "key": None,
                                                },
                                                expose_config=ConsulExposeConfig(
                                                    path=[
                                                        ConsulExposePath(
                                                            listener_port="listener_port_example",
                                                            local_path_port=1,
                                                            path="path_example",
                                                            protocol="protocol_example",
                                                        ),
                                                    ],
                                                ),
                                                local_service_address="local_service_address_example",
                                                local_service_port=1,
                                                upstreams=[
                                                    ConsulUpstream(
                                                        datacenter="datacenter_example",
                                                        destination_name="destination_name_example",
                                                        local_bind_address="local_bind_address_example",
                                                        local_bind_port=1,
                                                        mesh_gateway=ConsulMeshGateway(
                                                            mode="mode_example",
                                                        ),
                                                    ),
                                                ],
                                            ),
                                            tags=[
                                                "tags_example",
                                            ],
                                        ),
                                        sidecar_task=SidecarTask(
                                            config={
                                                "key": None,
                                            },
                                            driver="driver_example",
                                            env={
                                                "key": "key_example",
                                            },
                                            kill_signal="kill_signal_example",
                                            kill_timeout=1,
                                            log_config=LogConfig(
                                                max_file_size_mb=1,
                                                max_files=1,
                                            ),
                                            meta={
                                                "key": "key_example",
                                            },
                                            name="name_example",
                                            resources=Resources(
                                                cpu=1,
                                                cores=1,
                                                devices=[
                                                    RequestedDevice(
                                                        affinities=[
                                                            Affinity(
                                                                l_target="l_target_example",
                                                                operand="operand_example",
                                                                r_target="r_target_example",
                                                                weight=-128,
                                                            ),
                                                        ],
                                                        constraints=[
                                                            Constraint(
                                                                l_target="l_target_example",
                                                                operand="operand_example",
                                                                r_target="r_target_example",
                                                            ),
                                                        ],
                                                        count=0,
                                                        name="name_example",
                                                    ),
                                                ],
                                                disk_mb=1,
                                                iops=1,
                                                memory_mb=1,
                                                memory_max_mb=1,
                                                networks=[
                                                    NetworkResource(
                                                        cidr="cidr_example",
                                                        dns=DNSConfig(
                                                            options=[
                                                                "options_example",
                                                            ],
                                                            searches=[
                                                                "searches_example",
                                                            ],
                                                            servers=[
                                                                "servers_example",
                                                            ],
                                                        ),
                                                        device="device_example",
                                                        dynamic_ports=[
                                                            Port(
                                                                host_network="host_network_example",
                                                                label="label_example",
                                                                to=1,
                                                                value=1,
                                                            ),
                                                        ],
                                                        hostname="hostname_example",
                                                        ip="ip_example",
                                                        m_bits=1,
                                                        mode="mode_example",
                                                        reserved_ports=[
                                                            Port(
                                                                host_network="host_network_example",
                                                                label="label_example",
                                                                to=1,
                                                                value=1,
                                                            ),
                                                        ],
                                                    ),
                                                ],
                                            ),
                                            shutdown_delay=1,
                                            user="user_example",
                                        ),
                                    ),
                                    enable_tag_override=True,
                                    id="id_example",
                                    meta={
                                        "key": "key_example",
                                    },
                                    name="name_example",
                                    on_update="on_update_example",
                                    port_label="port_label_example",
                                    provider="provider_example",
                                    tags=[
                                        "tags_example",
                                    ],
                                    task_name="task_name_example",
                                ),
                            ],
                            shutdown_delay=1,
                            templates=[
                                Template(
                                    change_mode="change_mode_example",
                                    change_signal="change_signal_example",
                                    dest_path="dest_path_example",
                                    embedded_tmpl="embedded_tmpl_example",
                                    envvars=True,
                                    left_delim="left_delim_example",
                                    perms="perms_example",
                                    right_delim="right_delim_example",
                                    source_path="source_path_example",
                                    splay=1,
                                    vault_grace=1,
                                    wait=WaitConfig(
                                        max=1,
                                        min=1,
                                    ),
                                ),
                            ],
                            user="user_example",
                            vault=Vault(
                                change_mode="change_mode_example",
                                change_signal="change_signal_example",
                                entity_alias="entity_alias_example",
                                env=True,
                                namespace="namespace_example",
                                policies=[
                                    "policies_example",
                                ],
                            ),
                            volume_mounts=[
                                VolumeMount(
                                    destination="destination_example",
                                    propagation_mode="propagation_mode_example",
                                    read_only=True,
                                    volume="volume_example",
                                ),
                            ],
                        ),
                    ],
                    update=UpdateStrategy(
                        auto_promote=True,
                        auto_revert=True,
                        canary=1,
                        health_check="health_check_example",
                        healthy_deadline=1,
                        max_parallel=1,
                        min_healthy_time=1,
                        progress_deadline=1,
                        stagger=1,
                    ),
                    volumes={
                        "key": VolumeRequest(
                            access_mode="access_mode_example",
                            attachment_mode="attachment_mode_example",
                            mount_options=CSIMountOptions(
                                fs_type="fs_type_example",
                                mount_flags=[
                                    "mount_flags_example",
                                ],
                            ),
                            name="name_example",
                            per_alloc=True,
                            read_only=True,
                            source="source_example",
                            type="type_example",
                        ),
                    },
                ),
            ],
            type="type_example",
            update=UpdateStrategy(
                auto_promote=True,
                auto_revert=True,
                canary=1,
                health_check="health_check_example",
                healthy_deadline=1,
                max_parallel=1,
                min_healthy_time=1,
                progress_deadline=1,
                stagger=1,
            ),
            vault_namespace="vault_namespace_example",
            vault_token="vault_token_example",
            version=0,
        ),
        namespace="namespace_example",
        region="region_example",
        secret_id="secret_id_example",
    ) # JobValidateRequest | 
    region = "region_example" # str | Filters results based on the specified region. (optional)
    namespace = "namespace_example" # str | Filters results based on the specified namespace. (optional)
    x_nomad_token = "X-Nomad-Token_example" # str | A Nomad ACL token. (optional)
    idempotency_token = "idempotency_token_example" # str | Can be used to ensure operations are only run once. (optional)

    # example passing only required values which don't have defaults set
    try:
        api_response = api_instance.post_job_validate_request(job_validate_request)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling JobsApi->post_job_validate_request: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        api_response = api_instance.post_job_validate_request(job_validate_request, region=region, namespace=namespace, x_nomad_token=x_nomad_token, idempotency_token=idempotency_token)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling JobsApi->post_job_validate_request: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **job_validate_request** | [**JobValidateRequest**](JobValidateRequest.md)|  |
 **region** | **str**| Filters results based on the specified region. | [optional]
 **namespace** | **str**| Filters results based on the specified namespace. | [optional]
 **x_nomad_token** | **str**| A Nomad ACL token. | [optional]
 **idempotency_token** | **str**| Can be used to ensure operations are only run once. | [optional]

### Return type

[**JobValidateResponse**](JobValidateResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** |  |  * X-Nomad-Index - A unique identifier representing the current state of the requested resource. On a new Nomad cluster the value of this index starts at 1. <br>  |
**400** | Bad request |  -  |
**403** | Forbidden |  -  |
**405** | Method not allowed |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **register_job**
> JobRegisterResponse register_job(job_register_request)



### Example

* Api Key Authentication (X-Nomad-Token):
```python
import time
import nomad_client
from nomad_client.api import jobs_api
from nomad_client.model.job_register_request import JobRegisterRequest
from nomad_client.model.job_register_response import JobRegisterResponse
from pprint import pprint
# Defining the host is optional and defaults to https://127.0.0.1:4646/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = nomad_client.Configuration(
    host = "https://127.0.0.1:4646/v1"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure API key authorization: X-Nomad-Token
configuration.api_key['X-Nomad-Token'] = 'YOUR_API_KEY'

# Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
# configuration.api_key_prefix['X-Nomad-Token'] = 'Bearer'

# Enter a context with an instance of the API client
with nomad_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = jobs_api.JobsApi(api_client)
    job_register_request = JobRegisterRequest(
        enforce_index=True,
        eval_priority=1,
        job=Job(
            affinities=[
                Affinity(
                    l_target="l_target_example",
                    operand="operand_example",
                    r_target="r_target_example",
                    weight=-128,
                ),
            ],
            all_at_once=True,
            constraints=[
                Constraint(
                    l_target="l_target_example",
                    operand="operand_example",
                    r_target="r_target_example",
                ),
            ],
            consul_namespace="consul_namespace_example",
            consul_token="consul_token_example",
            create_index=0,
            datacenters=[
                "datacenters_example",
            ],
            dispatch_idempotency_token="dispatch_idempotency_token_example",
            dispatched=True,
            id="id_example",
            job_modify_index=0,
            meta={
                "key": "key_example",
            },
            migrate=MigrateStrategy(
                health_check="health_check_example",
                healthy_deadline=1,
                max_parallel=1,
                min_healthy_time=1,
            ),
            modify_index=0,
            multiregion=Multiregion(
                regions=[
                    MultiregionRegion(
                        count=1,
                        datacenters=[
                            "datacenters_example",
                        ],
                        meta={
                            "key": "key_example",
                        },
                        name="name_example",
                    ),
                ],
                strategy=MultiregionStrategy(
                    max_parallel=1,
                    on_failure="on_failure_example",
                ),
            ),
            name="name_example",
            namespace="namespace_example",
            nomad_token_id="nomad_token_id_example",
            parameterized_job=ParameterizedJobConfig(
                meta_optional=[
                    "meta_optional_example",
                ],
                meta_required=[
                    "meta_required_example",
                ],
                payload="payload_example",
            ),
            parent_id="parent_id_example",
            payload='YQ==',
            periodic=PeriodicConfig(
                enabled=True,
                prohibit_overlap=True,
                spec="spec_example",
                spec_type="spec_type_example",
                time_zone="time_zone_example",
            ),
            priority=1,
            region="region_example",
            reschedule=ReschedulePolicy(
                attempts=1,
                delay=1,
                delay_function="delay_function_example",
                interval=1,
                max_delay=1,
                unlimited=True,
            ),
            spreads=[
                Spread(
                    attribute="attribute_example",
                    spread_target=[
                        SpreadTarget(
                            percent=0,
                            value="value_example",
                        ),
                    ],
                    weight=-128,
                ),
            ],
            stable=True,
            status="status_example",
            status_description="status_description_example",
            stop=True,
            submit_time=1,
            task_groups=[
                TaskGroup(
                    affinities=[
                        Affinity(
                            l_target="l_target_example",
                            operand="operand_example",
                            r_target="r_target_example",
                            weight=-128,
                        ),
                    ],
                    constraints=[
                        Constraint(
                            l_target="l_target_example",
                            operand="operand_example",
                            r_target="r_target_example",
                        ),
                    ],
                    consul=Consul(
                        namespace="namespace_example",
                    ),
                    count=1,
                    ephemeral_disk=EphemeralDisk(
                        migrate=True,
                        size_mb=1,
                        sticky=True,
                    ),
                    max_client_disconnect=1,
                    meta={
                        "key": "key_example",
                    },
                    migrate=MigrateStrategy(
                        health_check="health_check_example",
                        healthy_deadline=1,
                        max_parallel=1,
                        min_healthy_time=1,
                    ),
                    name="name_example",
                    networks=[
                        NetworkResource(
                            cidr="cidr_example",
                            dns=DNSConfig(
                                options=[
                                    "options_example",
                                ],
                                searches=[
                                    "searches_example",
                                ],
                                servers=[
                                    "servers_example",
                                ],
                            ),
                            device="device_example",
                            dynamic_ports=[
                                Port(
                                    host_network="host_network_example",
                                    label="label_example",
                                    to=1,
                                    value=1,
                                ),
                            ],
                            hostname="hostname_example",
                            ip="ip_example",
                            m_bits=1,
                            mode="mode_example",
                            reserved_ports=[
                                Port(
                                    host_network="host_network_example",
                                    label="label_example",
                                    to=1,
                                    value=1,
                                ),
                            ],
                        ),
                    ],
                    reschedule_policy=ReschedulePolicy(
                        attempts=1,
                        delay=1,
                        delay_function="delay_function_example",
                        interval=1,
                        max_delay=1,
                        unlimited=True,
                    ),
                    restart_policy=RestartPolicy(
                        attempts=1,
                        delay=1,
                        interval=1,
                        mode="mode_example",
                    ),
                    scaling=ScalingPolicy(
                        create_index=0,
                        enabled=True,
                        id="id_example",
                        max=1,
                        min=1,
                        modify_index=0,
                        namespace="namespace_example",
                        policy={
                            "key": None,
                        },
                        target={
                            "key": "key_example",
                        },
                        type="type_example",
                    ),
                    services=[
                        Service(
                            address_mode="address_mode_example",
                            canary_meta={
                                "key": "key_example",
                            },
                            canary_tags=[
                                "canary_tags_example",
                            ],
                            check_restart=CheckRestart(
                                grace=1,
                                ignore_warnings=True,
                                limit=1,
                            ),
                            checks=[
                                ServiceCheck(
                                    address_mode="address_mode_example",
                                    args=[
                                        "args_example",
                                    ],
                                    body="body_example",
                                    check_restart=CheckRestart(
                                        grace=1,
                                        ignore_warnings=True,
                                        limit=1,
                                    ),
                                    command="command_example",
                                    expose=True,
                                    failures_before_critical=1,
                                    grpc_service="grpc_service_example",
                                    grpc_use_tls=True,
                                    header={
                                        "key": [
                                            "key_example",
                                        ],
                                    },
                                    id="id_example",
                                    initial_status="initial_status_example",
                                    interval=1,
                                    method="method_example",
                                    name="name_example",
                                    on_update="on_update_example",
                                    path="path_example",
                                    port_label="port_label_example",
                                    protocol="protocol_example",
                                    success_before_passing=1,
                                    tls_skip_verify=True,
                                    task_name="task_name_example",
                                    timeout=1,
                                    type="type_example",
                                ),
                            ],
                            connect=ConsulConnect(
                                gateway=ConsulGateway(
                                    ingress=ConsulIngressConfigEntry(
                                        listeners=[
                                            ConsulIngressListener(
                                                port=1,
                                                protocol="protocol_example",
                                                services=[
                                                    ConsulIngressService(
                                                        hosts=[
                                                            "hosts_example",
                                                        ],
                                                        name="name_example",
                                                    ),
                                                ],
                                            ),
                                        ],
                                        tls=ConsulGatewayTLSConfig(
                                            enabled=True,
                                        ),
                                    ),
                                    mesh=None,
                                    proxy=ConsulGatewayProxy(
                                        config={
                                            "key": None,
                                        },
                                        connect_timeout=1,
                                        envoy_dns_discovery_type="envoy_dns_discovery_type_example",
                                        envoy_gateway_bind_addresses={
                                            "key": ConsulGatewayBindAddress(
                                                address="address_example",
                                                name="name_example",
                                                port=1,
                                            ),
                                        },
                                        envoy_gateway_bind_tagged_addresses=True,
                                        envoy_gateway_no_default_bind=True,
                                    ),
                                    terminating=ConsulTerminatingConfigEntry(
                                        services=[
                                            ConsulLinkedService(
                                                ca_file="ca_file_example",
                                                cert_file="cert_file_example",
                                                key_file="key_file_example",
                                                name="name_example",
                                                sni="sni_example",
                                            ),
                                        ],
                                    ),
                                ),
                                native=True,
                                sidecar_service=ConsulSidecarService(
                                    disable_default_tcp_check=True,
                                    port="port_example",
                                    proxy=ConsulProxy(
                                        config={
                                            "key": None,
                                        },
                                        expose_config=ConsulExposeConfig(
                                            path=[
                                                ConsulExposePath(
                                                    listener_port="listener_port_example",
                                                    local_path_port=1,
                                                    path="path_example",
                                                    protocol="protocol_example",
                                                ),
                                            ],
                                        ),
                                        local_service_address="local_service_address_example",
                                        local_service_port=1,
                                        upstreams=[
                                            ConsulUpstream(
                                                datacenter="datacenter_example",
                                                destination_name="destination_name_example",
                                                local_bind_address="local_bind_address_example",
                                                local_bind_port=1,
                                                mesh_gateway=ConsulMeshGateway(
                                                    mode="mode_example",
                                                ),
                                            ),
                                        ],
                                    ),
                                    tags=[
                                        "tags_example",
                                    ],
                                ),
                                sidecar_task=SidecarTask(
                                    config={
                                        "key": None,
                                    },
                                    driver="driver_example",
                                    env={
                                        "key": "key_example",
                                    },
                                    kill_signal="kill_signal_example",
                                    kill_timeout=1,
                                    log_config=LogConfig(
                                        max_file_size_mb=1,
                                        max_files=1,
                                    ),
                                    meta={
                                        "key": "key_example",
                                    },
                                    name="name_example",
                                    resources=Resources(
                                        cpu=1,
                                        cores=1,
                                        devices=[
                                            RequestedDevice(
                                                affinities=[
                                                    Affinity(
                                                        l_target="l_target_example",
                                                        operand="operand_example",
                                                        r_target="r_target_example",
                                                        weight=-128,
                                                    ),
                                                ],
                                                constraints=[
                                                    Constraint(
                                                        l_target="l_target_example",
                                                        operand="operand_example",
                                                        r_target="r_target_example",
                                                    ),
                                                ],
                                                count=0,
                                                name="name_example",
                                            ),
                                        ],
                                        disk_mb=1,
                                        iops=1,
                                        memory_mb=1,
                                        memory_max_mb=1,
                                        networks=[
                                            NetworkResource(
                                                cidr="cidr_example",
                                                dns=DNSConfig(
                                                    options=[
                                                        "options_example",
                                                    ],
                                                    searches=[
                                                        "searches_example",
                                                    ],
                                                    servers=[
                                                        "servers_example",
                                                    ],
                                                ),
                                                device="device_example",
                                                dynamic_ports=[
                                                    Port(
                                                        host_network="host_network_example",
                                                        label="label_example",
                                                        to=1,
                                                        value=1,
                                                    ),
                                                ],
                                                hostname="hostname_example",
                                                ip="ip_example",
                                                m_bits=1,
                                                mode="mode_example",
                                                reserved_ports=[
                                                    Port(
                                                        host_network="host_network_example",
                                                        label="label_example",
                                                        to=1,
                                                        value=1,
                                                    ),
                                                ],
                                            ),
                                        ],
                                    ),
                                    shutdown_delay=1,
                                    user="user_example",
                                ),
                            ),
                            enable_tag_override=True,
                            id="id_example",
                            meta={
                                "key": "key_example",
                            },
                            name="name_example",
                            on_update="on_update_example",
                            port_label="port_label_example",
                            provider="provider_example",
                            tags=[
                                "tags_example",
                            ],
                            task_name="task_name_example",
                        ),
                    ],
                    shutdown_delay=1,
                    spreads=[
                        Spread(
                            attribute="attribute_example",
                            spread_target=[
                                SpreadTarget(
                                    percent=0,
                                    value="value_example",
                                ),
                            ],
                            weight=-128,
                        ),
                    ],
                    stop_after_client_disconnect=1,
                    tasks=[
                        Task(
                            affinities=[
                                Affinity(
                                    l_target="l_target_example",
                                    operand="operand_example",
                                    r_target="r_target_example",
                                    weight=-128,
                                ),
                            ],
                            artifacts=[
                                TaskArtifact(
                                    getter_headers={
                                        "key": "key_example",
                                    },
                                    getter_mode="getter_mode_example",
                                    getter_options={
                                        "key": "key_example",
                                    },
                                    getter_source="getter_source_example",
                                    relative_dest="relative_dest_example",
                                ),
                            ],
                            csi_plugin_config=TaskCSIPluginConfig(
                                id="id_example",
                                mount_dir="mount_dir_example",
                                type="type_example",
                            ),
                            config={
                                "key": None,
                            },
                            constraints=[
                                Constraint(
                                    l_target="l_target_example",
                                    operand="operand_example",
                                    r_target="r_target_example",
                                ),
                            ],
                            dispatch_payload=DispatchPayloadConfig(
                                file="file_example",
                            ),
                            driver="driver_example",
                            env={
                                "key": "key_example",
                            },
                            kill_signal="kill_signal_example",
                            kill_timeout=1,
                            kind="kind_example",
                            leader=True,
                            lifecycle=TaskLifecycle(
                                hook="hook_example",
                                sidecar=True,
                            ),
                            log_config=LogConfig(
                                max_file_size_mb=1,
                                max_files=1,
                            ),
                            meta={
                                "key": "key_example",
                            },
                            name="name_example",
                            resources=Resources(
                                cpu=1,
                                cores=1,
                                devices=[
                                    RequestedDevice(
                                        affinities=[
                                            Affinity(
                                                l_target="l_target_example",
                                                operand="operand_example",
                                                r_target="r_target_example",
                                                weight=-128,
                                            ),
                                        ],
                                        constraints=[
                                            Constraint(
                                                l_target="l_target_example",
                                                operand="operand_example",
                                                r_target="r_target_example",
                                            ),
                                        ],
                                        count=0,
                                        name="name_example",
                                    ),
                                ],
                                disk_mb=1,
                                iops=1,
                                memory_mb=1,
                                memory_max_mb=1,
                                networks=[
                                    NetworkResource(
                                        cidr="cidr_example",
                                        dns=DNSConfig(
                                            options=[
                                                "options_example",
                                            ],
                                            searches=[
                                                "searches_example",
                                            ],
                                            servers=[
                                                "servers_example",
                                            ],
                                        ),
                                        device="device_example",
                                        dynamic_ports=[
                                            Port(
                                                host_network="host_network_example",
                                                label="label_example",
                                                to=1,
                                                value=1,
                                            ),
                                        ],
                                        hostname="hostname_example",
                                        ip="ip_example",
                                        m_bits=1,
                                        mode="mode_example",
                                        reserved_ports=[
                                            Port(
                                                host_network="host_network_example",
                                                label="label_example",
                                                to=1,
                                                value=1,
                                            ),
                                        ],
                                    ),
                                ],
                            ),
                            restart_policy=RestartPolicy(
                                attempts=1,
                                delay=1,
                                interval=1,
                                mode="mode_example",
                            ),
                            scaling_policies=[
                                ScalingPolicy(
                                    create_index=0,
                                    enabled=True,
                                    id="id_example",
                                    max=1,
                                    min=1,
                                    modify_index=0,
                                    namespace="namespace_example",
                                    policy={
                                        "key": None,
                                    },
                                    target={
                                        "key": "key_example",
                                    },
                                    type="type_example",
                                ),
                            ],
                            services=[
                                Service(
                                    address_mode="address_mode_example",
                                    canary_meta={
                                        "key": "key_example",
                                    },
                                    canary_tags=[
                                        "canary_tags_example",
                                    ],
                                    check_restart=CheckRestart(
                                        grace=1,
                                        ignore_warnings=True,
                                        limit=1,
                                    ),
                                    checks=[
                                        ServiceCheck(
                                            address_mode="address_mode_example",
                                            args=[
                                                "args_example",
                                            ],
                                            body="body_example",
                                            check_restart=CheckRestart(
                                                grace=1,
                                                ignore_warnings=True,
                                                limit=1,
                                            ),
                                            command="command_example",
                                            expose=True,
                                            failures_before_critical=1,
                                            grpc_service="grpc_service_example",
                                            grpc_use_tls=True,
                                            header={
                                                "key": [
                                                    "key_example",
                                                ],
                                            },
                                            id="id_example",
                                            initial_status="initial_status_example",
                                            interval=1,
                                            method="method_example",
                                            name="name_example",
                                            on_update="on_update_example",
                                            path="path_example",
                                            port_label="port_label_example",
                                            protocol="protocol_example",
                                            success_before_passing=1,
                                            tls_skip_verify=True,
                                            task_name="task_name_example",
                                            timeout=1,
                                            type="type_example",
                                        ),
                                    ],
                                    connect=ConsulConnect(
                                        gateway=ConsulGateway(
                                            ingress=ConsulIngressConfigEntry(
                                                listeners=[
                                                    ConsulIngressListener(
                                                        port=1,
                                                        protocol="protocol_example",
                                                        services=[
                                                            ConsulIngressService(
                                                                hosts=[
                                                                    "hosts_example",
                                                                ],
                                                                name="name_example",
                                                            ),
                                                        ],
                                                    ),
                                                ],
                                                tls=ConsulGatewayTLSConfig(
                                                    enabled=True,
                                                ),
                                            ),
                                            mesh=None,
                                            proxy=ConsulGatewayProxy(
                                                config={
                                                    "key": None,
                                                },
                                                connect_timeout=1,
                                                envoy_dns_discovery_type="envoy_dns_discovery_type_example",
                                                envoy_gateway_bind_addresses={
                                                    "key": ConsulGatewayBindAddress(
                                                        address="address_example",
                                                        name="name_example",
                                                        port=1,
                                                    ),
                                                },
                                                envoy_gateway_bind_tagged_addresses=True,
                                                envoy_gateway_no_default_bind=True,
                                            ),
                                            terminating=ConsulTerminatingConfigEntry(
                                                services=[
                                                    ConsulLinkedService(
                                                        ca_file="ca_file_example",
                                                        cert_file="cert_file_example",
                                                        key_file="key_file_example",
                                                        name="name_example",
                                                        sni="sni_example",
                                                    ),
                                                ],
                                            ),
                                        ),
                                        native=True,
                                        sidecar_service=ConsulSidecarService(
                                            disable_default_tcp_check=True,
                                            port="port_example",
                                            proxy=ConsulProxy(
                                                config={
                                                    "key": None,
                                                },
                                                expose_config=ConsulExposeConfig(
                                                    path=[
                                                        ConsulExposePath(
                                                            listener_port="listener_port_example",
                                                            local_path_port=1,
                                                            path="path_example",
                                                            protocol="protocol_example",
                                                        ),
                                                    ],
                                                ),
                                                local_service_address="local_service_address_example",
                                                local_service_port=1,
                                                upstreams=[
                                                    ConsulUpstream(
                                                        datacenter="datacenter_example",
                                                        destination_name="destination_name_example",
                                                        local_bind_address="local_bind_address_example",
                                                        local_bind_port=1,
                                                        mesh_gateway=ConsulMeshGateway(
                                                            mode="mode_example",
                                                        ),
                                                    ),
                                                ],
                                            ),
                                            tags=[
                                                "tags_example",
                                            ],
                                        ),
                                        sidecar_task=SidecarTask(
                                            config={
                                                "key": None,
                                            },
                                            driver="driver_example",
                                            env={
                                                "key": "key_example",
                                            },
                                            kill_signal="kill_signal_example",
                                            kill_timeout=1,
                                            log_config=LogConfig(
                                                max_file_size_mb=1,
                                                max_files=1,
                                            ),
                                            meta={
                                                "key": "key_example",
                                            },
                                            name="name_example",
                                            resources=Resources(
                                                cpu=1,
                                                cores=1,
                                                devices=[
                                                    RequestedDevice(
                                                        affinities=[
                                                            Affinity(
                                                                l_target="l_target_example",
                                                                operand="operand_example",
                                                                r_target="r_target_example",
                                                                weight=-128,
                                                            ),
                                                        ],
                                                        constraints=[
                                                            Constraint(
                                                                l_target="l_target_example",
                                                                operand="operand_example",
                                                                r_target="r_target_example",
                                                            ),
                                                        ],
                                                        count=0,
                                                        name="name_example",
                                                    ),
                                                ],
                                                disk_mb=1,
                                                iops=1,
                                                memory_mb=1,
                                                memory_max_mb=1,
                                                networks=[
                                                    NetworkResource(
                                                        cidr="cidr_example",
                                                        dns=DNSConfig(
                                                            options=[
                                                                "options_example",
                                                            ],
                                                            searches=[
                                                                "searches_example",
                                                            ],
                                                            servers=[
                                                                "servers_example",
                                                            ],
                                                        ),
                                                        device="device_example",
                                                        dynamic_ports=[
                                                            Port(
                                                                host_network="host_network_example",
                                                                label="label_example",
                                                                to=1,
                                                                value=1,
                                                            ),
                                                        ],
                                                        hostname="hostname_example",
                                                        ip="ip_example",
                                                        m_bits=1,
                                                        mode="mode_example",
                                                        reserved_ports=[
                                                            Port(
                                                                host_network="host_network_example",
                                                                label="label_example",
                                                                to=1,
                                                                value=1,
                                                            ),
                                                        ],
                                                    ),
                                                ],
                                            ),
                                            shutdown_delay=1,
                                            user="user_example",
                                        ),
                                    ),
                                    enable_tag_override=True,
                                    id="id_example",
                                    meta={
                                        "key": "key_example",
                                    },
                                    name="name_example",
                                    on_update="on_update_example",
                                    port_label="port_label_example",
                                    provider="provider_example",
                                    tags=[
                                        "tags_example",
                                    ],
                                    task_name="task_name_example",
                                ),
                            ],
                            shutdown_delay=1,
                            templates=[
                                Template(
                                    change_mode="change_mode_example",
                                    change_signal="change_signal_example",
                                    dest_path="dest_path_example",
                                    embedded_tmpl="embedded_tmpl_example",
                                    envvars=True,
                                    left_delim="left_delim_example",
                                    perms="perms_example",
                                    right_delim="right_delim_example",
                                    source_path="source_path_example",
                                    splay=1,
                                    vault_grace=1,
                                    wait=WaitConfig(
                                        max=1,
                                        min=1,
                                    ),
                                ),
                            ],
                            user="user_example",
                            vault=Vault(
                                change_mode="change_mode_example",
                                change_signal="change_signal_example",
                                entity_alias="entity_alias_example",
                                env=True,
                                namespace="namespace_example",
                                policies=[
                                    "policies_example",
                                ],
                            ),
                            volume_mounts=[
                                VolumeMount(
                                    destination="destination_example",
                                    propagation_mode="propagation_mode_example",
                                    read_only=True,
                                    volume="volume_example",
                                ),
                            ],
                        ),
                    ],
                    update=UpdateStrategy(
                        auto_promote=True,
                        auto_revert=True,
                        canary=1,
                        health_check="health_check_example",
                        healthy_deadline=1,
                        max_parallel=1,
                        min_healthy_time=1,
                        progress_deadline=1,
                        stagger=1,
                    ),
                    volumes={
                        "key": VolumeRequest(
                            access_mode="access_mode_example",
                            attachment_mode="attachment_mode_example",
                            mount_options=CSIMountOptions(
                                fs_type="fs_type_example",
                                mount_flags=[
                                    "mount_flags_example",
                                ],
                            ),
                            name="name_example",
                            per_alloc=True,
                            read_only=True,
                            source="source_example",
                            type="type_example",
                        ),
                    },
                ),
            ],
            type="type_example",
            update=UpdateStrategy(
                auto_promote=True,
                auto_revert=True,
                canary=1,
                health_check="health_check_example",
                healthy_deadline=1,
                max_parallel=1,
                min_healthy_time=1,
                progress_deadline=1,
                stagger=1,
            ),
            vault_namespace="vault_namespace_example",
            vault_token="vault_token_example",
            version=0,
        ),
        job_modify_index=0,
        namespace="namespace_example",
        policy_override=True,
        preserve_counts=True,
        region="region_example",
        secret_id="secret_id_example",
    ) # JobRegisterRequest | 
    region = "region_example" # str | Filters results based on the specified region. (optional)
    namespace = "namespace_example" # str | Filters results based on the specified namespace. (optional)
    x_nomad_token = "X-Nomad-Token_example" # str | A Nomad ACL token. (optional)
    idempotency_token = "idempotency_token_example" # str | Can be used to ensure operations are only run once. (optional)

    # example passing only required values which don't have defaults set
    try:
        api_response = api_instance.register_job(job_register_request)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling JobsApi->register_job: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        api_response = api_instance.register_job(job_register_request, region=region, namespace=namespace, x_nomad_token=x_nomad_token, idempotency_token=idempotency_token)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling JobsApi->register_job: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **job_register_request** | [**JobRegisterRequest**](JobRegisterRequest.md)|  |
 **region** | **str**| Filters results based on the specified region. | [optional]
 **namespace** | **str**| Filters results based on the specified namespace. | [optional]
 **x_nomad_token** | **str**| A Nomad ACL token. | [optional]
 **idempotency_token** | **str**| Can be used to ensure operations are only run once. | [optional]

### Return type

[**JobRegisterResponse**](JobRegisterResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** |  |  * X-Nomad-Index - A unique identifier representing the current state of the requested resource. On a new Nomad cluster the value of this index starts at 1. <br>  * X-Nomad-KnownLeader - Boolean indicating if there is a known cluster leader. <br>  * X-Nomad-LastContact - The time in milliseconds that a server was last contacted by the leader node. <br>  |
**400** | Bad request |  -  |
**403** | Forbidden |  -  |
**405** | Method not allowed |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

