# nomad_client.DeploymentsApi

All URIs are relative to *https://127.0.0.1:4646/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get_deployment**](DeploymentsApi.md#get_deployment) | **GET** /deployment/{deploymentID} | 
[**get_deployment_allocations**](DeploymentsApi.md#get_deployment_allocations) | **GET** /deployment/allocations/{deploymentID} | 
[**get_deployments**](DeploymentsApi.md#get_deployments) | **GET** /deployments | 
[**post_deployment_allocation_health**](DeploymentsApi.md#post_deployment_allocation_health) | **POST** /deployment/allocation-health/{deploymentID} | 
[**post_deployment_fail**](DeploymentsApi.md#post_deployment_fail) | **POST** /deployment/fail/{deploymentID} | 
[**post_deployment_pause**](DeploymentsApi.md#post_deployment_pause) | **POST** /deployment/pause/{deploymentID} | 
[**post_deployment_promote**](DeploymentsApi.md#post_deployment_promote) | **POST** /deployment/promote/{deploymentID} | 
[**post_deployment_unblock**](DeploymentsApi.md#post_deployment_unblock) | **POST** /deployment/unblock/{deploymentID} | 


# **get_deployment**
> Deployment get_deployment(deployment_id)



### Example

* Api Key Authentication (X-Nomad-Token):

```python
import time
import nomad_client
from nomad_client.api import deployments_api
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
    api_instance = deployments_api.DeploymentsApi(api_client)
    deployment_id = "deploymentID_example" # str | Deployment ID.
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
        api_response = api_instance.get_deployment(deployment_id)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling DeploymentsApi->get_deployment: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        api_response = api_instance.get_deployment(deployment_id, region=region, namespace=namespace, index=index, wait=wait, stale=stale, prefix=prefix, x_nomad_token=x_nomad_token, per_page=per_page, next_token=next_token)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling DeploymentsApi->get_deployment: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deployment_id** | **str**| Deployment ID. |
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

# **get_deployment_allocations**
> [AllocationListStub] get_deployment_allocations(deployment_id)



### Example

* Api Key Authentication (X-Nomad-Token):

```python
import time
import nomad_client
from nomad_client.api import deployments_api
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
    api_instance = deployments_api.DeploymentsApi(api_client)
    deployment_id = "deploymentID_example" # str | Deployment ID.
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
        api_response = api_instance.get_deployment_allocations(deployment_id)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling DeploymentsApi->get_deployment_allocations: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        api_response = api_instance.get_deployment_allocations(deployment_id, region=region, namespace=namespace, index=index, wait=wait, stale=stale, prefix=prefix, x_nomad_token=x_nomad_token, per_page=per_page, next_token=next_token)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling DeploymentsApi->get_deployment_allocations: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deployment_id** | **str**| Deployment ID. |
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

# **get_deployments**
> [Deployment] get_deployments()



### Example

* Api Key Authentication (X-Nomad-Token):

```python
import time
import nomad_client
from nomad_client.api import deployments_api
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
    api_instance = deployments_api.DeploymentsApi(api_client)
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
        api_response = api_instance.get_deployments(region=region, namespace=namespace, index=index, wait=wait, stale=stale, prefix=prefix, x_nomad_token=x_nomad_token, per_page=per_page, next_token=next_token)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling DeploymentsApi->get_deployments: %s\n" % e)
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

# **post_deployment_allocation_health**
> DeploymentUpdateResponse post_deployment_allocation_health(deployment_id, deployment_alloc_health_request)



### Example

* Api Key Authentication (X-Nomad-Token):

```python
import time
import nomad_client
from nomad_client.api import deployments_api
from nomad_client.model.deployment_update_response import DeploymentUpdateResponse
from nomad_client.model.deployment_alloc_health_request import DeploymentAllocHealthRequest
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
    api_instance = deployments_api.DeploymentsApi(api_client)
    deployment_id = "deploymentID_example" # str | Deployment ID.
    deployment_alloc_health_request = DeploymentAllocHealthRequest(
        deployment_id="deployment_id_example",
        healthy_allocation_ids=[
            "healthy_allocation_ids_example",
        ],
        namespace="namespace_example",
        region="region_example",
        secret_id="secret_id_example",
        unhealthy_allocation_ids=[
            "unhealthy_allocation_ids_example",
        ],
    ) # DeploymentAllocHealthRequest | 
    region = "region_example" # str | Filters results based on the specified region. (optional)
    namespace = "namespace_example" # str | Filters results based on the specified namespace. (optional)
    x_nomad_token = "X-Nomad-Token_example" # str | A Nomad ACL token. (optional)
    idempotency_token = "idempotency_token_example" # str | Can be used to ensure operations are only run once. (optional)

    # example passing only required values which don't have defaults set
    try:
        api_response = api_instance.post_deployment_allocation_health(deployment_id, deployment_alloc_health_request)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling DeploymentsApi->post_deployment_allocation_health: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        api_response = api_instance.post_deployment_allocation_health(deployment_id, deployment_alloc_health_request, region=region, namespace=namespace, x_nomad_token=x_nomad_token, idempotency_token=idempotency_token)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling DeploymentsApi->post_deployment_allocation_health: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deployment_id** | **str**| Deployment ID. |
 **deployment_alloc_health_request** | [**DeploymentAllocHealthRequest**](DeploymentAllocHealthRequest.md)|  |
 **region** | **str**| Filters results based on the specified region. | [optional]
 **namespace** | **str**| Filters results based on the specified namespace. | [optional]
 **x_nomad_token** | **str**| A Nomad ACL token. | [optional]
 **idempotency_token** | **str**| Can be used to ensure operations are only run once. | [optional]

### Return type

[**DeploymentUpdateResponse**](DeploymentUpdateResponse.md)

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

# **post_deployment_fail**
> DeploymentUpdateResponse post_deployment_fail(deployment_id)



### Example

* Api Key Authentication (X-Nomad-Token):

```python
import time
import nomad_client
from nomad_client.api import deployments_api
from nomad_client.model.deployment_update_response import DeploymentUpdateResponse
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
    api_instance = deployments_api.DeploymentsApi(api_client)
    deployment_id = "deploymentID_example" # str | Deployment ID.
    region = "region_example" # str | Filters results based on the specified region. (optional)
    namespace = "namespace_example" # str | Filters results based on the specified namespace. (optional)
    x_nomad_token = "X-Nomad-Token_example" # str | A Nomad ACL token. (optional)
    idempotency_token = "idempotency_token_example" # str | Can be used to ensure operations are only run once. (optional)

    # example passing only required values which don't have defaults set
    try:
        api_response = api_instance.post_deployment_fail(deployment_id)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling DeploymentsApi->post_deployment_fail: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        api_response = api_instance.post_deployment_fail(deployment_id, region=region, namespace=namespace, x_nomad_token=x_nomad_token, idempotency_token=idempotency_token)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling DeploymentsApi->post_deployment_fail: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deployment_id** | **str**| Deployment ID. |
 **region** | **str**| Filters results based on the specified region. | [optional]
 **namespace** | **str**| Filters results based on the specified namespace. | [optional]
 **x_nomad_token** | **str**| A Nomad ACL token. | [optional]
 **idempotency_token** | **str**| Can be used to ensure operations are only run once. | [optional]

### Return type

[**DeploymentUpdateResponse**](DeploymentUpdateResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

 - **Content-Type**: Not defined
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

# **post_deployment_pause**
> DeploymentUpdateResponse post_deployment_pause(deployment_id, deployment_pause_request)



### Example

* Api Key Authentication (X-Nomad-Token):

```python
import time
import nomad_client
from nomad_client.api import deployments_api
from nomad_client.model.deployment_update_response import DeploymentUpdateResponse
from nomad_client.model.deployment_pause_request import DeploymentPauseRequest
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
    api_instance = deployments_api.DeploymentsApi(api_client)
    deployment_id = "deploymentID_example" # str | Deployment ID.
    deployment_pause_request = DeploymentPauseRequest(
        deployment_id="deployment_id_example",
        namespace="namespace_example",
        pause=True,
        region="region_example",
        secret_id="secret_id_example",
    ) # DeploymentPauseRequest | 
    region = "region_example" # str | Filters results based on the specified region. (optional)
    namespace = "namespace_example" # str | Filters results based on the specified namespace. (optional)
    x_nomad_token = "X-Nomad-Token_example" # str | A Nomad ACL token. (optional)
    idempotency_token = "idempotency_token_example" # str | Can be used to ensure operations are only run once. (optional)

    # example passing only required values which don't have defaults set
    try:
        api_response = api_instance.post_deployment_pause(deployment_id, deployment_pause_request)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling DeploymentsApi->post_deployment_pause: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        api_response = api_instance.post_deployment_pause(deployment_id, deployment_pause_request, region=region, namespace=namespace, x_nomad_token=x_nomad_token, idempotency_token=idempotency_token)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling DeploymentsApi->post_deployment_pause: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deployment_id** | **str**| Deployment ID. |
 **deployment_pause_request** | [**DeploymentPauseRequest**](DeploymentPauseRequest.md)|  |
 **region** | **str**| Filters results based on the specified region. | [optional]
 **namespace** | **str**| Filters results based on the specified namespace. | [optional]
 **x_nomad_token** | **str**| A Nomad ACL token. | [optional]
 **idempotency_token** | **str**| Can be used to ensure operations are only run once. | [optional]

### Return type

[**DeploymentUpdateResponse**](DeploymentUpdateResponse.md)

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

# **post_deployment_promote**
> DeploymentUpdateResponse post_deployment_promote(deployment_id, deployment_promote_request)



### Example

* Api Key Authentication (X-Nomad-Token):

```python
import time
import nomad_client
from nomad_client.api import deployments_api
from nomad_client.model.deployment_update_response import DeploymentUpdateResponse
from nomad_client.model.deployment_promote_request import DeploymentPromoteRequest
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
    api_instance = deployments_api.DeploymentsApi(api_client)
    deployment_id = "deploymentID_example" # str | Deployment ID.
    deployment_promote_request = DeploymentPromoteRequest(
        all=True,
        deployment_id="deployment_id_example",
        groups=[
            "groups_example",
        ],
        namespace="namespace_example",
        region="region_example",
        secret_id="secret_id_example",
    ) # DeploymentPromoteRequest | 
    region = "region_example" # str | Filters results based on the specified region. (optional)
    namespace = "namespace_example" # str | Filters results based on the specified namespace. (optional)
    x_nomad_token = "X-Nomad-Token_example" # str | A Nomad ACL token. (optional)
    idempotency_token = "idempotency_token_example" # str | Can be used to ensure operations are only run once. (optional)

    # example passing only required values which don't have defaults set
    try:
        api_response = api_instance.post_deployment_promote(deployment_id, deployment_promote_request)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling DeploymentsApi->post_deployment_promote: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        api_response = api_instance.post_deployment_promote(deployment_id, deployment_promote_request, region=region, namespace=namespace, x_nomad_token=x_nomad_token, idempotency_token=idempotency_token)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling DeploymentsApi->post_deployment_promote: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deployment_id** | **str**| Deployment ID. |
 **deployment_promote_request** | [**DeploymentPromoteRequest**](DeploymentPromoteRequest.md)|  |
 **region** | **str**| Filters results based on the specified region. | [optional]
 **namespace** | **str**| Filters results based on the specified namespace. | [optional]
 **x_nomad_token** | **str**| A Nomad ACL token. | [optional]
 **idempotency_token** | **str**| Can be used to ensure operations are only run once. | [optional]

### Return type

[**DeploymentUpdateResponse**](DeploymentUpdateResponse.md)

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

# **post_deployment_unblock**
> DeploymentUpdateResponse post_deployment_unblock(deployment_id, deployment_unblock_request)



### Example

* Api Key Authentication (X-Nomad-Token):

```python
import time
import nomad_client
from nomad_client.api import deployments_api
from nomad_client.model.deployment_update_response import DeploymentUpdateResponse
from nomad_client.model.deployment_unblock_request import DeploymentUnblockRequest
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
    api_instance = deployments_api.DeploymentsApi(api_client)
    deployment_id = "deploymentID_example" # str | Deployment ID.
    deployment_unblock_request = DeploymentUnblockRequest(
        deployment_id="deployment_id_example",
        namespace="namespace_example",
        region="region_example",
        secret_id="secret_id_example",
    ) # DeploymentUnblockRequest | 
    region = "region_example" # str | Filters results based on the specified region. (optional)
    namespace = "namespace_example" # str | Filters results based on the specified namespace. (optional)
    x_nomad_token = "X-Nomad-Token_example" # str | A Nomad ACL token. (optional)
    idempotency_token = "idempotency_token_example" # str | Can be used to ensure operations are only run once. (optional)

    # example passing only required values which don't have defaults set
    try:
        api_response = api_instance.post_deployment_unblock(deployment_id, deployment_unblock_request)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling DeploymentsApi->post_deployment_unblock: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        api_response = api_instance.post_deployment_unblock(deployment_id, deployment_unblock_request, region=region, namespace=namespace, x_nomad_token=x_nomad_token, idempotency_token=idempotency_token)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling DeploymentsApi->post_deployment_unblock: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deployment_id** | **str**| Deployment ID. |
 **deployment_unblock_request** | [**DeploymentUnblockRequest**](DeploymentUnblockRequest.md)|  |
 **region** | **str**| Filters results based on the specified region. | [optional]
 **namespace** | **str**| Filters results based on the specified namespace. | [optional]
 **x_nomad_token** | **str**| A Nomad ACL token. | [optional]
 **idempotency_token** | **str**| Can be used to ensure operations are only run once. | [optional]

### Return type

[**DeploymentUpdateResponse**](DeploymentUpdateResponse.md)

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

