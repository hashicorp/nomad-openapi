# nomad_client.MetricsApi

All URIs are relative to *https://127.0.0.1:4646/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get_metrics_summary**](MetricsApi.md#get_metrics_summary) | **GET** /metrics | 


# **get_metrics_summary**
> MetricsSummary get_metrics_summary()



### Example

* Api Key Authentication (X-Nomad-Token):

```python
import time
import nomad_client
from nomad_client.api import metrics_api
from nomad_client.model.metrics_summary import MetricsSummary
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
    api_instance = metrics_api.MetricsApi(api_client)
    format = "format_example" # str | The format the user requested for the metrics summary (e.g. prometheus) (optional)

    # example passing only required values which don't have defaults set
    # and optional values
    try:
        api_response = api_instance.get_metrics_summary(format=format)
        pprint(api_response)
    except nomad_client.ApiException as e:
        print("Exception when calling MetricsApi->get_metrics_summary: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **str**| The format the user requested for the metrics summary (e.g. prometheus) | [optional]

### Return type

[**MetricsSummary**](MetricsSummary.md)

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

