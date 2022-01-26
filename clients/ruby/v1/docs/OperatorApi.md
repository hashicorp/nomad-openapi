# openapi_client.OperatorApi

All URIs are relative to *https://127.0.0.1:4646/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**delete_operator_raft_peer**](OperatorApi.md#delete_operator_raft_peer) | **DELETE** /operator/raft/peer | 
[**get_operator_autopilot_configuration**](OperatorApi.md#get_operator_autopilot_configuration) | **GET** /operator/autopilot/configuration | 
[**get_operator_autopilot_health**](OperatorApi.md#get_operator_autopilot_health) | **GET** /operator/autopilot/health | 
[**get_operator_raft_configuration**](OperatorApi.md#get_operator_raft_configuration) | **GET** /operator/raft/configuration | 
[**get_operator_scheduler_configuration**](OperatorApi.md#get_operator_scheduler_configuration) | **GET** /operator/scheduler/configuration | 
[**post_operator_scheduler_configuration**](OperatorApi.md#post_operator_scheduler_configuration) | **POST** /operator/scheduler/configuration | 
[**put_operator_autopilot_configuration**](OperatorApi.md#put_operator_autopilot_configuration) | **PUT** /operator/autopilot/configuration | 


# **delete_operator_raft_peer**
> delete_operator_raft_peer()



### Example

* Api Key Authentication (X-Nomad-Token):
```python
import time
import openapi_client
from openapi_client.api import operator_api
from pprint import pprint
# Defining the host is optional and defaults to https://127.0.0.1:4646/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
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
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = operator_api.OperatorApi(api_client)
    region = "region_example" # str | Filters results based on the specified region. (optional)
    namespace = "namespace_example" # str | Filters results based on the specified namespace. (optional)
    x_nomad_token = "X-Nomad-Token_example" # str | A Nomad ACL token. (optional)
    idempotency_token = "idempotency_token_example" # str | Can be used to ensure operations are only run once. (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        api_instance.delete_operator_raft_peer(region=region, namespace=namespace, x_nomad_token=x_nomad_token, idempotency_token=idempotency_token)
    except openapi_client.ApiException as e:
        print("Exception when calling OperatorApi->delete_operator_raft_peer: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **region** | **str**| Filters results based on the specified region. | [optional]
 **namespace** | **str**| Filters results based on the specified namespace. | [optional]
 **x_nomad_token** | **str**| A Nomad ACL token. | [optional]
 **idempotency_token** | **str**| Can be used to ensure operations are only run once. | [optional]

### Return type

void (empty response body)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** |  |  -  |
**400** | Bad request |  -  |
**403** | Forbidden |  -  |
**405** | Method not allowed |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_operator_autopilot_configuration**
> AutopilotConfiguration get_operator_autopilot_configuration()



### Example

* Api Key Authentication (X-Nomad-Token):
```python
import time
import openapi_client
from openapi_client.api import operator_api
from openapi_client.model.autopilot_configuration import AutopilotConfiguration
from pprint import pprint
# Defining the host is optional and defaults to https://127.0.0.1:4646/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
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
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = operator_api.OperatorApi(api_client)
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
        api_response = api_instance.get_operator_autopilot_configuration(region=region, namespace=namespace, index=index, wait=wait, stale=stale, prefix=prefix, x_nomad_token=x_nomad_token, per_page=per_page, next_token=next_token)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling OperatorApi->get_operator_autopilot_configuration: %s\n" % e)
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

[**AutopilotConfiguration**](AutopilotConfiguration.md)

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

# **get_operator_autopilot_health**
> OperatorHealthReply get_operator_autopilot_health()



### Example

* Api Key Authentication (X-Nomad-Token):
```python
import time
import openapi_client
from openapi_client.api import operator_api
from openapi_client.model.operator_health_reply import OperatorHealthReply
from pprint import pprint
# Defining the host is optional and defaults to https://127.0.0.1:4646/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
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
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = operator_api.OperatorApi(api_client)
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
        api_response = api_instance.get_operator_autopilot_health(region=region, namespace=namespace, index=index, wait=wait, stale=stale, prefix=prefix, x_nomad_token=x_nomad_token, per_page=per_page, next_token=next_token)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling OperatorApi->get_operator_autopilot_health: %s\n" % e)
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

[**OperatorHealthReply**](OperatorHealthReply.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

<<<<<<< HEAD
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

# **get_operator_raft_configuration**
> [RaftServer] get_operator_raft_configuration()



### Example

* Api Key Authentication (X-Nomad-Token):
```python
import time
import openapi_client
from openapi_client.api import operator_api
from openapi_client.model.raft_server import RaftServer
from pprint import pprint
# Defining the host is optional and defaults to https://127.0.0.1:4646/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
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
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = operator_api.OperatorApi(api_client)
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
        api_response = api_instance.get_operator_raft_configuration(region=region, namespace=namespace, index=index, wait=wait, stale=stale, prefix=prefix, x_nomad_token=x_nomad_token, per_page=per_page, next_token=next_token)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling OperatorApi->get_operator_raft_configuration: %s\n" % e)
```

=======
- **Content-Type**: Not defined
- **Accept**: application/json


## get_operator_raft_configuration

> <Array<RaftConfigurationResponse>> get_operator_raft_configuration(opts)



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

api_instance = NomadClient::OperatorApi.new
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
  
  result = api_instance.get_operator_raft_configuration(opts)
  p result
rescue NomadClient::ApiError => e
  puts "Error when calling OperatorApi->get_operator_raft_configuration: #{e}"
end
```

#### Using the get_operator_raft_configuration_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Array<RaftConfigurationResponse>>, Integer, Hash)> get_operator_raft_configuration_with_http_info(opts)

```ruby
begin
  
  data, status_code, headers = api_instance.get_operator_raft_configuration_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Array<RaftConfigurationResponse>>
rescue NomadClient::ApiError => e
  puts "Error when calling OperatorApi->get_operator_raft_configuration_with_http_info: #{e}"
end
```
>>>>>>> 730e030 (updated tests, TestPutAutopilotConfiguration and TestPostSchedulerConfiguration working)

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

<<<<<<< HEAD
[**[RaftServer]**](RaftServer.md)
=======
[**Array&lt;RaftConfigurationResponse&gt;**](RaftConfigurationResponse.md)
>>>>>>> 730e030 (updated tests, TestPutAutopilotConfiguration and TestPostSchedulerConfiguration working)

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

# **get_operator_scheduler_configuration**
> SchedulerConfigurationResponse get_operator_scheduler_configuration()



### Example

* Api Key Authentication (X-Nomad-Token):
```python
import time
import openapi_client
from openapi_client.api import operator_api
from openapi_client.model.scheduler_configuration_response import SchedulerConfigurationResponse
from pprint import pprint
# Defining the host is optional and defaults to https://127.0.0.1:4646/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
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
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = operator_api.OperatorApi(api_client)
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
        api_response = api_instance.get_operator_scheduler_configuration(region=region, namespace=namespace, index=index, wait=wait, stale=stale, prefix=prefix, x_nomad_token=x_nomad_token, per_page=per_page, next_token=next_token)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling OperatorApi->get_operator_scheduler_configuration: %s\n" % e)
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

[**SchedulerConfigurationResponse**](SchedulerConfigurationResponse.md)

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

# **post_operator_scheduler_configuration**
> SchedulerSetConfigurationResponse post_operator_scheduler_configuration(scheduler_configuration)



### Example

* Api Key Authentication (X-Nomad-Token):
```python
import time
import openapi_client
from openapi_client.api import operator_api
from openapi_client.model.scheduler_set_configuration_response import SchedulerSetConfigurationResponse
from openapi_client.model.scheduler_configuration import SchedulerConfiguration
from pprint import pprint
# Defining the host is optional and defaults to https://127.0.0.1:4646/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
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
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = operator_api.OperatorApi(api_client)
    scheduler_configuration = SchedulerConfiguration(
        create_index=0,
        memory_oversubscription_enabled=True,
        modify_index=0,
        preemption_config=PreemptionConfig(
            batch_scheduler_enabled=True,
            service_scheduler_enabled=True,
            sys_batch_scheduler_enabled=True,
            system_scheduler_enabled=True,
        ),
        scheduler_algorithm="scheduler_algorithm_example",
    ) # SchedulerConfiguration | 
    region = "region_example" # str | Filters results based on the specified region. (optional)
    namespace = "namespace_example" # str | Filters results based on the specified namespace. (optional)
    x_nomad_token = "X-Nomad-Token_example" # str | A Nomad ACL token. (optional)
    idempotency_token = "idempotency_token_example" # str | Can be used to ensure operations are only run once. (optional)

    # example passing only required values which don't have defaults set
    try:
        api_response = api_instance.post_operator_scheduler_configuration(scheduler_configuration)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling OperatorApi->post_operator_scheduler_configuration: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        api_response = api_instance.post_operator_scheduler_configuration(scheduler_configuration, region=region, namespace=namespace, x_nomad_token=x_nomad_token, idempotency_token=idempotency_token)
        pprint(api_response)
    except openapi_client.ApiException as e:
        print("Exception when calling OperatorApi->post_operator_scheduler_configuration: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scheduler_configuration** | [**SchedulerConfiguration**](SchedulerConfiguration.md)|  |
 **region** | **str**| Filters results based on the specified region. | [optional]
 **namespace** | **str**| Filters results based on the specified namespace. | [optional]
 **x_nomad_token** | **str**| A Nomad ACL token. | [optional]
 **idempotency_token** | **str**| Can be used to ensure operations are only run once. | [optional]

### Return type

[**SchedulerSetConfigurationResponse**](SchedulerSetConfigurationResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

<<<<<<< HEAD
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

# **put_operator_autopilot_configuration**
> put_operator_autopilot_configuration(autopilot_configuration)



### Example

* Api Key Authentication (X-Nomad-Token):
```python
import time
import openapi_client
from openapi_client.api import operator_api
from openapi_client.model.autopilot_configuration import AutopilotConfiguration
from pprint import pprint
# Defining the host is optional and defaults to https://127.0.0.1:4646/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
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
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = operator_api.OperatorApi(api_client)
    autopilot_configuration = AutopilotConfiguration(
        cleanup_dead_servers=True,
        create_index=0,
        disable_upgrade_migration=True,
        enable_custom_upgrades=True,
        enable_redundancy_zones=True,
        last_contact_threshold=1,
        max_trailing_logs=0,
        min_quorum=0,
        modify_index=0,
        server_stabilization_time=1,
    ) # AutopilotConfiguration | 
    region = "region_example" # str | Filters results based on the specified region. (optional)
    namespace = "namespace_example" # str | Filters results based on the specified namespace. (optional)
    x_nomad_token = "X-Nomad-Token_example" # str | A Nomad ACL token. (optional)
    idempotency_token = "idempotency_token_example" # str | Can be used to ensure operations are only run once. (optional)

    # example passing only required values which don't have defaults set
    try:
        api_instance.put_operator_autopilot_configuration(autopilot_configuration)
    except openapi_client.ApiException as e:
        print("Exception when calling OperatorApi->put_operator_autopilot_configuration: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        api_instance.put_operator_autopilot_configuration(autopilot_configuration, region=region, namespace=namespace, x_nomad_token=x_nomad_token, idempotency_token=idempotency_token)
    except openapi_client.ApiException as e:
        print("Exception when calling OperatorApi->put_operator_autopilot_configuration: %s\n" % e)
```

=======
- **Content-Type**: application/json
- **Accept**: application/json


## put_operator_autopilot_configuration

> Boolean put_operator_autopilot_configuration(autopilot_configuration, opts)



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

api_instance = NomadClient::OperatorApi.new
autopilot_configuration = NomadClient::AutopilotConfiguration.new # AutopilotConfiguration | 
opts = {
  region: 'region_example', # String | Filters results based on the specified region.
  namespace: 'namespace_example', # String | Filters results based on the specified namespace.
  x_nomad_token: 'x_nomad_token_example', # String | A Nomad ACL token.
  idempotency_token: 'idempotency_token_example' # String | Can be used to ensure operations are only run once.
}

begin
  
  result = api_instance.put_operator_autopilot_configuration(autopilot_configuration, opts)
  p result
rescue NomadClient::ApiError => e
  puts "Error when calling OperatorApi->put_operator_autopilot_configuration: #{e}"
end
```

#### Using the put_operator_autopilot_configuration_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(Boolean, Integer, Hash)> put_operator_autopilot_configuration_with_http_info(autopilot_configuration, opts)

```ruby
begin
  
  data, status_code, headers = api_instance.put_operator_autopilot_configuration_with_http_info(autopilot_configuration, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => Boolean
rescue NomadClient::ApiError => e
  puts "Error when calling OperatorApi->put_operator_autopilot_configuration_with_http_info: #{e}"
end
```
>>>>>>> 730e030 (updated tests, TestPutAutopilotConfiguration and TestPostSchedulerConfiguration working)

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **autopilot_configuration** | [**AutopilotConfiguration**](AutopilotConfiguration.md)|  |
 **region** | **str**| Filters results based on the specified region. | [optional]
 **namespace** | **str**| Filters results based on the specified namespace. | [optional]
 **x_nomad_token** | **str**| A Nomad ACL token. | [optional]
 **idempotency_token** | **str**| Can be used to ensure operations are only run once. | [optional]

### Return type

<<<<<<< HEAD
void (empty response body)
=======
**Boolean**
>>>>>>> 730e030 (updated tests, TestPutAutopilotConfiguration and TestPostSchedulerConfiguration working)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

<<<<<<< HEAD
 - **Content-Type**: application/json
 - **Accept**: Not defined


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** |  |  -  |
**400** | Bad request |  -  |
**403** | Forbidden |  -  |
**405** | Method not allowed |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
=======
- **Content-Type**: application/json
- **Accept**: application/json
>>>>>>> 730e030 (updated tests, TestPutAutopilotConfiguration and TestPostSchedulerConfiguration working)

