# nomad_client.VariablesApi

All URIs are relative to *http://127.0.0.1:4646/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**delete_variable**](VariablesApi.md#delete_variable) | **DELETE** /var/{path} | 
[**get_variable_query**](VariablesApi.md#get_variable_query) | **GET** /var/{path} | 
[**get_variables_list_request**](VariablesApi.md#get_variables_list_request) | **GET** /vars | 
[**post_variable**](VariablesApi.md#post_variable) | **POST** /var/{path} | 
[**put_variable**](VariablesApi.md#put_variable) | **PUT** /var/{path} | 


# **delete_variable**
> delete_variable(path, variable)



### Example

* Api Key Authentication (X-Nomad-Token):
```python
import time
import nomad_client
from nomad_client.api import variables_api
from nomad_client.model.variable import Variable
from pprint import pprint
# Defining the host is optional and defaults to http://127.0.0.1:4646/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = nomad_client.Configuration(
    host = "http://127.0.0.1:4646/v1"
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
    api_instance = variables_api.VariablesApi(api_client)
    path = "path_example" # str | A path to a Nomad Variable
    variable = Variable(
        create_index=0,
        create_time=1,
        items=VariableItems(
            key="key_example",
        ),
        modify_index=0,
        modify_time=1,
        namespace="namespace_example",
        path="path_example",
    ) # Variable | 
    region = "region_example" # str | Filters results based on the specified region. (optional)
    namespace = "namespace_example" # str | Filters results based on the specified namespace. (optional)
    x_nomad_token = "X-Nomad-Token_example" # str | A Nomad ACL token. (optional)
    idempotency_token = "idempotency_token_example" # str | Can be used to ensure operations are only run once. (optional)
    cas = 1 # int | A compare-and-set parameter for Nomad Variables (optional)

    # example passing only required values which don't have defaults set
    try:
        api_instance.delete_variable(path, variable)
    except nomad_client.ApiException as e:
        print("Exception when calling VariablesApi->delete_variable: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        api_instance.delete_variable(path, variable, region=region, namespace=namespace, x_nomad_token=x_nomad_token, idempotency_token=idempotency_token, cas=cas)
    except nomad_client.ApiException as e:
        print("Exception when calling VariablesApi->delete_variable: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **str**| A path to a Nomad Variable |
 **variable** | [**Variable**](Variable.md)|  |
 **region** | **str**| Filters results based on the specified region. | [optional]
 **namespace** | **str**| Filters results based on the specified namespace. | [optional]
 **x_nomad_token** | **str**| A Nomad ACL token. | [optional]
 **idempotency_token** | **str**| Can be used to ensure operations are only run once. | [optional]
 **cas** | **int**| A compare-and-set parameter for Nomad Variables | [optional]

### Return type

void (empty response body)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** |  |  * X-Nomad-Index - A unique identifier representing the current state of the requested resource. On a new Nomad cluster the value of this index starts at 1. <br>  |
**400** | Bad request |  -  |
**403** | Forbidden |  -  |
**405** | Method not allowed |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_variable_query**
> Variable get_variable_query(path)



### Example

* Api Key Authentication (X-Nomad-Token):
```python
import time
import nomad_client
from nomad_client.api import variables_api
from nomad_client.model.variable import Variable
from pprint import pprint
# Defining the host is optional and defaults to http://127.0.0.1:4646/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = nomad_client.Configuration(
    host = "http://127.0.0.1:4646/v1"
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
    api_instance = variables_api.VariablesApi(api_client)
    path = "path_example" # str | A path to a Nomad Variable
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
        api_response = api_instance.get_variable_query(path)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling VariablesApi->get_variable_query: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        api_response = api_instance.get_variable_query(path, region=region, namespace=namespace, index=index, wait=wait, stale=stale, prefix=prefix, x_nomad_token=x_nomad_token, per_page=per_page, next_token=next_token)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling VariablesApi->get_variable_query: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **str**| A path to a Nomad Variable |
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

[**Variable**](Variable.md)

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

# **get_variables_list_request**
> [VariableMetadata] get_variables_list_request()



### Example

* Api Key Authentication (X-Nomad-Token):
```python
import time
import nomad_client
from nomad_client.api import variables_api
from nomad_client.model.variable_metadata import VariableMetadata
from pprint import pprint
# Defining the host is optional and defaults to http://127.0.0.1:4646/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = nomad_client.Configuration(
    host = "http://127.0.0.1:4646/v1"
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
    api_instance = variables_api.VariablesApi(api_client)
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
        api_response = api_instance.get_variables_list_request(region=region, namespace=namespace, index=index, wait=wait, stale=stale, prefix=prefix, x_nomad_token=x_nomad_token, per_page=per_page, next_token=next_token)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling VariablesApi->get_variables_list_request: %s\n" % e)
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

[**[VariableMetadata]**](VariableMetadata.md)

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

# **post_variable**
> Variable post_variable(path, variable)



### Example

* Api Key Authentication (X-Nomad-Token):
```python
import time
import nomad_client
from nomad_client.api import variables_api
from nomad_client.model.variable import Variable
from pprint import pprint
# Defining the host is optional and defaults to http://127.0.0.1:4646/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = nomad_client.Configuration(
    host = "http://127.0.0.1:4646/v1"
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
    api_instance = variables_api.VariablesApi(api_client)
    path = "path_example" # str | A path to a Nomad Variable
    variable = Variable(
        create_index=0,
        create_time=1,
        items=VariableItems(
            key="key_example",
        ),
        modify_index=0,
        modify_time=1,
        namespace="namespace_example",
        path="path_example",
    ) # Variable | 
    region = "region_example" # str | Filters results based on the specified region. (optional)
    namespace = "namespace_example" # str | Filters results based on the specified namespace. (optional)
    x_nomad_token = "X-Nomad-Token_example" # str | A Nomad ACL token. (optional)
    idempotency_token = "idempotency_token_example" # str | Can be used to ensure operations are only run once. (optional)
    cas = 1 # int | A compare-and-set parameter for Nomad Variables (optional)

    # example passing only required values which don't have defaults set
    try:
        api_response = api_instance.post_variable(path, variable)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling VariablesApi->post_variable: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        api_response = api_instance.post_variable(path, variable, region=region, namespace=namespace, x_nomad_token=x_nomad_token, idempotency_token=idempotency_token, cas=cas)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling VariablesApi->post_variable: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **str**| A path to a Nomad Variable |
 **variable** | [**Variable**](Variable.md)|  |
 **region** | **str**| Filters results based on the specified region. | [optional]
 **namespace** | **str**| Filters results based on the specified namespace. | [optional]
 **x_nomad_token** | **str**| A Nomad ACL token. | [optional]
 **idempotency_token** | **str**| Can be used to ensure operations are only run once. | [optional]
 **cas** | **int**| A compare-and-set parameter for Nomad Variables | [optional]

### Return type

[**Variable**](Variable.md)

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

# **put_variable**
> Variable put_variable(path, variable)



### Example

* Api Key Authentication (X-Nomad-Token):
```python
import time
import nomad_client
from nomad_client.api import variables_api
from nomad_client.model.variable import Variable
from pprint import pprint
# Defining the host is optional and defaults to http://127.0.0.1:4646/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = nomad_client.Configuration(
    host = "http://127.0.0.1:4646/v1"
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
    api_instance = variables_api.VariablesApi(api_client)
    path = "path_example" # str | A path to a Nomad Variable
    variable = Variable(
        create_index=0,
        create_time=1,
        items=VariableItems(
            key="key_example",
        ),
        modify_index=0,
        modify_time=1,
        namespace="namespace_example",
        path="path_example",
    ) # Variable | 
    region = "region_example" # str | Filters results based on the specified region. (optional)
    namespace = "namespace_example" # str | Filters results based on the specified namespace. (optional)
    x_nomad_token = "X-Nomad-Token_example" # str | A Nomad ACL token. (optional)
    idempotency_token = "idempotency_token_example" # str | Can be used to ensure operations are only run once. (optional)
    cas = 1 # int | A compare-and-set parameter for Nomad Variables (optional)

    # example passing only required values which don't have defaults set
    try:
        api_response = api_instance.put_variable(path, variable)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling VariablesApi->put_variable: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        api_response = api_instance.put_variable(path, variable, region=region, namespace=namespace, x_nomad_token=x_nomad_token, idempotency_token=idempotency_token, cas=cas)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling VariablesApi->put_variable: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **str**| A path to a Nomad Variable |
 **variable** | [**Variable**](Variable.md)|  |
 **region** | **str**| Filters results based on the specified region. | [optional]
 **namespace** | **str**| Filters results based on the specified namespace. | [optional]
 **x_nomad_token** | **str**| A Nomad ACL token. | [optional]
 **idempotency_token** | **str**| Can be used to ensure operations are only run once. | [optional]
 **cas** | **int**| A compare-and-set parameter for Nomad Variables | [optional]

### Return type

[**Variable**](Variable.md)

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

