# \SystemApi

All URIs are relative to *https://127.0.0.1:4646/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PutSystemGC**](SystemApi.md#PutSystemGC) | **Put** /system/gc | 
[**PutSystemReconcileSummaries**](SystemApi.md#PutSystemReconcileSummaries) | **Put** /system/reconcile/summaries | 



## PutSystemGC

> PutSystemGC(ctx).Region(region).Namespace(namespace).XNomadToken(xNomadToken).IdempotencyToken(idempotencyToken).Execute()



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
    region := "region_example" // string | Filters results based on the specified region. (optional)
    namespace := "namespace_example" // string | Filters results based on the specified namespace. (optional)
    xNomadToken := "xNomadToken_example" // string | A Nomad ACL token. (optional)
    idempotencyToken := "idempotencyToken_example" // string | Can be used to ensure operations are only run once. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.PutSystemGC(context.Background()).Region(region).Namespace(namespace).XNomadToken(xNomadToken).IdempotencyToken(idempotencyToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.PutSystemGC``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutSystemGCRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **region** | **string** | Filters results based on the specified region. | 
 **namespace** | **string** | Filters results based on the specified namespace. | 
 **xNomadToken** | **string** | A Nomad ACL token. | 
 **idempotencyToken** | **string** | Can be used to ensure operations are only run once. | 

### Return type

 (empty response body)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSystemReconcileSummaries

> PutSystemReconcileSummaries(ctx).Region(region).Namespace(namespace).XNomadToken(xNomadToken).IdempotencyToken(idempotencyToken).Execute()



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
    region := "region_example" // string | Filters results based on the specified region. (optional)
    namespace := "namespace_example" // string | Filters results based on the specified namespace. (optional)
    xNomadToken := "xNomadToken_example" // string | A Nomad ACL token. (optional)
    idempotencyToken := "idempotencyToken_example" // string | Can be used to ensure operations are only run once. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemApi.PutSystemReconcileSummaries(context.Background()).Region(region).Namespace(namespace).XNomadToken(xNomadToken).IdempotencyToken(idempotencyToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemApi.PutSystemReconcileSummaries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutSystemReconcileSummariesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **region** | **string** | Filters results based on the specified region. | 
 **namespace** | **string** | Filters results based on the specified namespace. | 
 **xNomadToken** | **string** | A Nomad ACL token. | 
 **idempotencyToken** | **string** | Can be used to ensure operations are only run once. | 

### Return type

 (empty response body)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

