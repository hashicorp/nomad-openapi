# nomad_client.EnterpriseApi

All URIs are relative to *https://127.0.0.1:4646/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_quota_spec**](EnterpriseApi.md#create_quota_spec) | **POST** /quota | 
[**delete_quota_spec**](EnterpriseApi.md#delete_quota_spec) | **DELETE** /quota/{specName} | 
[**get_quota_spec**](EnterpriseApi.md#get_quota_spec) | **GET** /quota/{specName} | 
[**get_quotas**](EnterpriseApi.md#get_quotas) | **GET** /quotas | 
[**post_quota_spec**](EnterpriseApi.md#post_quota_spec) | **POST** /quota/{specName} | 


# **create_quota_spec**
> create_quota_spec(quota_spec)



### Example

* Api Key Authentication (X-Nomad-Token):

```python
import time
import nomad_client
from nomad_client.api import enterprise_api
from nomad_client.model.quota_spec import QuotaSpec
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
    api_instance = enterprise_api.EnterpriseApi(api_client)
    quota_spec = QuotaSpec(
        create_index=0,
        description="description_example",
        limits=[
            QuotaLimit(
                hash='YQ==',
                region="region_example",
                region_limit=Resources(
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
            ),
        ],
        modify_index=0,
        name="name_example",
    ) # QuotaSpec | 
    region = "region_example" # str | Filters results based on the specified region. (optional)
    namespace = "namespace_example" # str | Filters results based on the specified namespace. (optional)
    x_nomad_token = "X-Nomad-Token_example" # str | A Nomad ACL token. (optional)
    idempotency_token = "idempotency_token_example" # str | Can be used to ensure operations are only run once. (optional)

    # example passing only required values which don't have defaults set
    try:
        api_instance.create_quota_spec(quota_spec)
    except nomad_client.ApiException as e:
        print("Exception when calling EnterpriseApi->create_quota_spec: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        api_instance.create_quota_spec(quota_spec, region=region, namespace=namespace, x_nomad_token=x_nomad_token, idempotency_token=idempotency_token)
    except nomad_client.ApiException as e:
        print("Exception when calling EnterpriseApi->create_quota_spec: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **quota_spec** | [**QuotaSpec**](QuotaSpec.md)|  |
 **region** | **str**| Filters results based on the specified region. | [optional]
 **namespace** | **str**| Filters results based on the specified namespace. | [optional]
 **x_nomad_token** | **str**| A Nomad ACL token. | [optional]
 **idempotency_token** | **str**| Can be used to ensure operations are only run once. | [optional]

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

# **delete_quota_spec**
> delete_quota_spec(spec_name)



### Example

* Api Key Authentication (X-Nomad-Token):

```python
import time
import nomad_client
from nomad_client.api import enterprise_api
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
    api_instance = enterprise_api.EnterpriseApi(api_client)
    spec_name = "specName_example" # str | The quota spec identifier.
    region = "region_example" # str | Filters results based on the specified region. (optional)
    namespace = "namespace_example" # str | Filters results based on the specified namespace. (optional)
    x_nomad_token = "X-Nomad-Token_example" # str | A Nomad ACL token. (optional)
    idempotency_token = "idempotency_token_example" # str | Can be used to ensure operations are only run once. (optional)

    # example passing only required values which don't have defaults set
    try:
        api_instance.delete_quota_spec(spec_name)
    except nomad_client.ApiException as e:
        print("Exception when calling EnterpriseApi->delete_quota_spec: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        api_instance.delete_quota_spec(spec_name, region=region, namespace=namespace, x_nomad_token=x_nomad_token, idempotency_token=idempotency_token)
    except nomad_client.ApiException as e:
        print("Exception when calling EnterpriseApi->delete_quota_spec: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **spec_name** | **str**| The quota spec identifier. |
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
**200** |  |  * X-Nomad-Index - A unique identifier representing the current state of the requested resource. On a new Nomad cluster the value of this index starts at 1. <br>  |
**400** | Bad request |  -  |
**403** | Forbidden |  -  |
**405** | Method not allowed |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_quota_spec**
> QuotaSpec get_quota_spec(spec_name)



### Example

* Api Key Authentication (X-Nomad-Token):

```python
import time
import nomad_client
from nomad_client.api import enterprise_api
from nomad_client.model.quota_spec import QuotaSpec
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
    api_instance = enterprise_api.EnterpriseApi(api_client)
    spec_name = "specName_example" # str | The quota spec identifier.
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
        api_response = api_instance.get_quota_spec(spec_name)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling EnterpriseApi->get_quota_spec: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        api_response = api_instance.get_quota_spec(spec_name, region=region, namespace=namespace, index=index, wait=wait, stale=stale, prefix=prefix, x_nomad_token=x_nomad_token, per_page=per_page, next_token=next_token)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling EnterpriseApi->get_quota_spec: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **spec_name** | **str**| The quota spec identifier. |
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

[**QuotaSpec**](QuotaSpec.md)

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

# **get_quotas**
> [bool, date, datetime, dict, float, int, list, str, none_type] get_quotas()



### Example

* Api Key Authentication (X-Nomad-Token):

```python
import time
import nomad_client
from nomad_client.api import enterprise_api
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
    api_instance = enterprise_api.EnterpriseApi(api_client)
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
        api_response = api_instance.get_quotas(region=region, namespace=namespace, index=index, wait=wait, stale=stale, prefix=prefix, x_nomad_token=x_nomad_token, per_page=per_page, next_token=next_token)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling EnterpriseApi->get_quotas: %s\n" % e)
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

**[bool, date, datetime, dict, float, int, list, str, none_type]**

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

# **post_quota_spec**
> post_quota_spec(spec_name, quota_spec)



### Example

* Api Key Authentication (X-Nomad-Token):

```python
import time
import nomad_client
from nomad_client.api import enterprise_api
from nomad_client.model.quota_spec import QuotaSpec
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
    api_instance = enterprise_api.EnterpriseApi(api_client)
    spec_name = "specName_example" # str | The quota spec identifier.
    quota_spec = QuotaSpec(
        create_index=0,
        description="description_example",
        limits=[
            QuotaLimit(
                hash='YQ==',
                region="region_example",
                region_limit=Resources(
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
            ),
        ],
        modify_index=0,
        name="name_example",
    ) # QuotaSpec | 
    region = "region_example" # str | Filters results based on the specified region. (optional)
    namespace = "namespace_example" # str | Filters results based on the specified namespace. (optional)
    x_nomad_token = "X-Nomad-Token_example" # str | A Nomad ACL token. (optional)
    idempotency_token = "idempotency_token_example" # str | Can be used to ensure operations are only run once. (optional)

    # example passing only required values which don't have defaults set
    try:
        api_instance.post_quota_spec(spec_name, quota_spec)
    except nomad_client.ApiException as e:
        print("Exception when calling EnterpriseApi->post_quota_spec: %s\n" % e)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        api_instance.post_quota_spec(spec_name, quota_spec, region=region, namespace=namespace, x_nomad_token=x_nomad_token, idempotency_token=idempotency_token)
    except nomad_client.ApiException as e:
        print("Exception when calling EnterpriseApi->post_quota_spec: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **spec_name** | **str**| The quota spec identifier. |
 **quota_spec** | [**QuotaSpec**](QuotaSpec.md)|  |
 **region** | **str**| Filters results based on the specified region. | [optional]
 **namespace** | **str**| Filters results based on the specified namespace. | [optional]
 **x_nomad_token** | **str**| A Nomad ACL token. | [optional]
 **idempotency_token** | **str**| Can be used to ensure operations are only run once. | [optional]

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

