# nomad-client.MetricsApi

All URIs are relative to *https://127.0.0.1:4646/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**getMetricsSummary**](MetricsApi.md#getMetricsSummary) | **GET** /metrics | 


# **getMetricsSummary**
> MetricsSummary getMetricsSummary()


### Example


```typescript
import { nomad-client } from '';
import * as fs from 'fs';

const configuration = nomad-client.createConfiguration();
const apiInstance = new nomad-client.MetricsApi(configuration);

let body:nomad-client.MetricsApiGetMetricsSummaryRequest = {
  // string | The format the user requested for the metrics summary (e.g. prometheus) (optional)
  format: "format_example",
};

apiInstance.getMetricsSummary(body).then((data:any) => {
  console.log('API called successfully. Returned data: ' + data);
}).catch((error:any) => console.error(error));
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | [**string**] | The format the user requested for the metrics summary (e.g. prometheus) | (optional) defaults to undefined


### Return type

**MetricsSummary**

### Authorization

[X-Nomad-Token](README.md#X-Nomad-Token)

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

[[Back to top]](#) [[Back to API list]](README.md#documentation-for-api-endpoints) [[Back to Model list]](README.md#documentation-for-models) [[Back to README]](README.md)


