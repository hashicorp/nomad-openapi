# \MetricsApi

All URIs are relative to *https://127.0.0.1:4646/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMetricsSummary**](MetricsApi.md#GetMetricsSummary) | **Get** /metrics | 



## GetMetricsSummary

> MetricsSummary GetMetricsSummary(ctx).Format(format).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    format := "format_example" // string | The format the user requested for the metrics summary (e.g. prometheus) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetricsApi.GetMetricsSummary(context.Background()).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsApi.GetMetricsSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMetricsSummary`: MetricsSummary
    fmt.Fprintf(os.Stdout, "Response from `MetricsApi.GetMetricsSummary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMetricsSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string** | The format the user requested for the metrics summary (e.g. prometheus) | 

### Return type

[**MetricsSummary**](MetricsSummary.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

