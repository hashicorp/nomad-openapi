# Nomad.Client.Api.MetricsApi

All URIs are relative to *https://127.0.0.1:4646/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMetricsSummary**](MetricsApi.md#getmetricssummary) | **GET** /metrics | 


<a name="getmetricssummary"></a>
# **GetMetricsSummary**
> MetricsSummary GetMetricsSummary (string format = null)



### Example
```csharp
using System.Collections.Generic;
using System.Diagnostics;
using Nomad.Client.Api;
using Nomad.Client.Client;
using Nomad.Client.Model;

namespace Example
{
    public class GetMetricsSummaryExample
    {
        public static void Main()
        {
            Configuration config = new Configuration();
            config.BasePath = "https://127.0.0.1:4646/v1";
            // Configure API key authorization: X-Nomad-Token
            config.AddApiKey("X-Nomad-Token", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // config.AddApiKeyPrefix("X-Nomad-Token", "Bearer");

            var apiInstance = new MetricsApi(config);
            var format = format_example;  // string | The format the user requested for the metrics summary (e.g. prometheus) (optional) 

            try
            {
                MetricsSummary result = apiInstance.GetMetricsSummary(format);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling MetricsApi.GetMetricsSummary: " + e.Message );
                Debug.Print("Status Code: "+ e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string**| The format the user requested for the metrics summary (e.g. prometheus) | [optional] 

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
| **200** |  |  -  |
| **400** | Bad request |  -  |
| **403** | Forbidden |  -  |
| **405** | Method not allowed |  -  |
| **500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

