# nomad_client.RegionsApi

All URIs are relative to *https://127.0.0.1:4646/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get_regions**](RegionsApi.md#get_regions) | **GET** /regions | 


# **get_regions**
> [str] get_regions()



### Example

* Api Key Authentication (X-Nomad-Token):

```python
import time
import nomad_client
from nomad_client.api import regions_api
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
    api_instance = regions_api.RegionsApi(api_client)

    # example, this endpoint has no required or optional parameters
    try:
        api_response = api_instance.get_regions()
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling RegionsApi->get_regions: %s\n" % e)
```


### Parameters
This endpoint does not need any parameter.

### Return type

**[str]**

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

