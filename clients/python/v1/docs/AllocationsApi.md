# nomad_client.AllocationsApi

All URIs are relative to *https://127.0.0.1:4646/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get_allocation**](AllocationsApi.md#get_allocation) | **GET** /allocation/{allocID} | 
[**get_allocation_services**](AllocationsApi.md#get_allocation_services) | **GET** /allocation/{allocID}/services | 
[**get_allocations**](AllocationsApi.md#get_allocations) | **GET** /allocations | 
[**post_allocation_stop**](AllocationsApi.md#post_allocation_stop) | **POST** /allocation/{allocID}/stop | 


# **get_allocation**
> Allocation get_allocation(alloc_id)



### Example

* Api Key Authentication (X-Nomad-Token):

```python
import time
import nomad_client
from nomad_client.api import allocations_api
from nomad_client.model.allocation import Allocation
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
    api_instance = allocations_api.AllocationsApi(api_client)
    alloc_id = "allocID_example" # str | Allocation ID.
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
        api_response = api_instance.get_allocation(alloc_id)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling AllocationsApi->get_allocation: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        api_response = api_instance.get_allocation(alloc_id, region=region, namespace=namespace, index=index, wait=wait, stale=stale, prefix=prefix, x_nomad_token=x_nomad_token, per_page=per_page, next_token=next_token)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling AllocationsApi->get_allocation: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **alloc_id** | **str**| Allocation ID. |
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

[**Allocation**](Allocation.md)

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

# **get_allocation_services**
> [ServiceRegistration] get_allocation_services(alloc_id)



### Example

* Api Key Authentication (X-Nomad-Token):

```python
import time
import nomad_client
from nomad_client.api import allocations_api
from nomad_client.model.service_registration import ServiceRegistration
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
    api_instance = allocations_api.AllocationsApi(api_client)
    alloc_id = "allocID_example" # str | Allocation ID.
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
        api_response = api_instance.get_allocation_services(alloc_id)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling AllocationsApi->get_allocation_services: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        api_response = api_instance.get_allocation_services(alloc_id, region=region, namespace=namespace, index=index, wait=wait, stale=stale, prefix=prefix, x_nomad_token=x_nomad_token, per_page=per_page, next_token=next_token)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling AllocationsApi->get_allocation_services: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **alloc_id** | **str**| Allocation ID. |
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

[**[ServiceRegistration]**](ServiceRegistration.md)

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

# **get_allocations**
> [AllocationListStub] get_allocations()



### Example

* Api Key Authentication (X-Nomad-Token):

```python
import time
import nomad_client
from nomad_client.api import allocations_api
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
    api_instance = allocations_api.AllocationsApi(api_client)
    region = "region_example" # str | Filters results based on the specified region. (optional)
    namespace = "namespace_example" # str | Filters results based on the specified namespace. (optional)
    index = 1 # int | If set, wait until query exceeds given index. Must be provided with WaitParam. (optional)
    wait = "wait_example" # str | Provided with IndexParam to wait for change. (optional)
    stale = "stale_example" # str | If present, results will include stale reads. (optional)
    prefix = "prefix_example" # str | Constrains results to jobs that start with the defined prefix (optional)
    x_nomad_token = "X-Nomad-Token_example" # str | A Nomad ACL token. (optional)
    per_page = 1 # int | Maximum number of results to return. (optional)
    next_token = "next_token_example" # str | Indicates where to start paging for queries that support pagination. (optional)
    resources = True # bool | Flag indicating whether to include resources in response. (optional)
    task_states = True # bool | Flag indicating whether to include task states in response. (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        api_response = api_instance.get_allocations(region=region, namespace=namespace, index=index, wait=wait, stale=stale, prefix=prefix, x_nomad_token=x_nomad_token, per_page=per_page, next_token=next_token, resources=resources, task_states=task_states)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling AllocationsApi->get_allocations: %s\n" % e)
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
 **resources** | **bool**| Flag indicating whether to include resources in response. | [optional]
 **task_states** | **bool**| Flag indicating whether to include task states in response. | [optional]

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

# **post_allocation_stop**
> AllocStopResponse post_allocation_stop(alloc_id)



### Example

* Api Key Authentication (X-Nomad-Token):

```python
import time
import nomad_client
from nomad_client.api import allocations_api
from nomad_client.model.alloc_stop_response import AllocStopResponse
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
    api_instance = allocations_api.AllocationsApi(api_client)
    alloc_id = "allocID_example" # str | Allocation ID.
    region = "region_example" # str | Filters results based on the specified region. (optional)
    namespace = "namespace_example" # str | Filters results based on the specified namespace. (optional)
    index = 1 # int | If set, wait until query exceeds given index. Must be provided with WaitParam. (optional)
    wait = "wait_example" # str | Provided with IndexParam to wait for change. (optional)
    stale = "stale_example" # str | If present, results will include stale reads. (optional)
    prefix = "prefix_example" # str | Constrains results to jobs that start with the defined prefix (optional)
    x_nomad_token = "X-Nomad-Token_example" # str | A Nomad ACL token. (optional)
    per_page = 1 # int | Maximum number of results to return. (optional)
    next_token = "next_token_example" # str | Indicates where to start paging for queries that support pagination. (optional)
    no_shutdown_delay = True # bool | Flag indicating whether to delay shutdown when requesting an allocation stop. (optional)

    # example passing only required values which don't have defaults set
    try:
        api_response = api_instance.post_allocation_stop(alloc_id)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling AllocationsApi->post_allocation_stop: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        api_response = api_instance.post_allocation_stop(alloc_id, region=region, namespace=namespace, index=index, wait=wait, stale=stale, prefix=prefix, x_nomad_token=x_nomad_token, per_page=per_page, next_token=next_token, no_shutdown_delay=no_shutdown_delay)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling AllocationsApi->post_allocation_stop: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **alloc_id** | **str**| Allocation ID. |
 **region** | **str**| Filters results based on the specified region. | [optional]
 **namespace** | **str**| Filters results based on the specified namespace. | [optional]
 **index** | **int**| If set, wait until query exceeds given index. Must be provided with WaitParam. | [optional]
 **wait** | **str**| Provided with IndexParam to wait for change. | [optional]
 **stale** | **str**| If present, results will include stale reads. | [optional]
 **prefix** | **str**| Constrains results to jobs that start with the defined prefix | [optional]
 **x_nomad_token** | **str**| A Nomad ACL token. | [optional]
 **per_page** | **int**| Maximum number of results to return. | [optional]
 **next_token** | **str**| Indicates where to start paging for queries that support pagination. | [optional]
 **no_shutdown_delay** | **bool**| Flag indicating whether to delay shutdown when requesting an allocation stop. | [optional]

### Return type

[**AllocStopResponse**](AllocStopResponse.md)

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

