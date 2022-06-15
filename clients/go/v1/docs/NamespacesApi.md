# \NamespacesApi

All URIs are relative to *https://127.0.0.1:4646/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNamespace**](NamespacesApi.md#CreateNamespace) | **Post** /namespace | 
[**DeleteNamespace**](NamespacesApi.md#DeleteNamespace) | **Delete** /namespace/{namespaceName} | 
[**GetNamespace**](NamespacesApi.md#GetNamespace) | **Get** /namespace/{namespaceName} | 
[**GetNamespaces**](NamespacesApi.md#GetNamespaces) | **Get** /namespaces | 
[**PostNamespace**](NamespacesApi.md#PostNamespace) | **Post** /namespace/{namespaceName} | 



## CreateNamespace

> CreateNamespace(ctx).Region(region).Namespace(namespace).XNomadToken(xNomadToken).IdempotencyToken(idempotencyToken).Execute()



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
    resp, r, err := apiClient.NamespacesApi.CreateNamespace(context.Background()).Region(region).Namespace(namespace).XNomadToken(xNomadToken).IdempotencyToken(idempotencyToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.CreateNamespace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNamespaceRequest struct via the builder pattern


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


## DeleteNamespace

> DeleteNamespace(ctx, namespaceName).Region(region).Namespace(namespace).XNomadToken(xNomadToken).IdempotencyToken(idempotencyToken).Execute()



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
    namespaceName := "namespaceName_example" // string | The namespace identifier.
    region := "region_example" // string | Filters results based on the specified region. (optional)
    namespace := "namespace_example" // string | Filters results based on the specified namespace. (optional)
    xNomadToken := "xNomadToken_example" // string | A Nomad ACL token. (optional)
    idempotencyToken := "idempotencyToken_example" // string | Can be used to ensure operations are only run once. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NamespacesApi.DeleteNamespace(context.Background(), namespaceName).Region(region).Namespace(namespace).XNomadToken(xNomadToken).IdempotencyToken(idempotencyToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.DeleteNamespace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespaceName** | **string** | The namespace identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNamespaceRequest struct via the builder pattern


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


## GetNamespace

> Namespace GetNamespace(ctx, namespaceName).Region(region).Namespace(namespace).Index(index).Wait(wait).Stale(stale).Prefix(prefix).XNomadToken(xNomadToken).PerPage(perPage).NextToken(nextToken).Execute()



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
    namespaceName := "namespaceName_example" // string | The namespace identifier.
    region := "region_example" // string | Filters results based on the specified region. (optional)
    namespace := "namespace_example" // string | Filters results based on the specified namespace. (optional)
    index := int32(56) // int32 | If set, wait until query exceeds given index. Must be provided with WaitParam. (optional)
    wait := "wait_example" // string | Provided with IndexParam to wait for change. (optional)
    stale := "stale_example" // string | If present, results will include stale reads. (optional)
    prefix := "prefix_example" // string | Constrains results to jobs that start with the defined prefix (optional)
    xNomadToken := "xNomadToken_example" // string | A Nomad ACL token. (optional)
    perPage := int32(56) // int32 | Maximum number of results to return. (optional)
    nextToken := "nextToken_example" // string | Indicates where to start paging for queries that support pagination. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NamespacesApi.GetNamespace(context.Background(), namespaceName).Region(region).Namespace(namespace).Index(index).Wait(wait).Stale(stale).Prefix(prefix).XNomadToken(xNomadToken).PerPage(perPage).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.GetNamespace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNamespace`: Namespace
    fmt.Fprintf(os.Stdout, "Response from `NamespacesApi.GetNamespace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespaceName** | **string** | The namespace identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNamespaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **region** | **string** | Filters results based on the specified region. | 
 **namespace** | **string** | Filters results based on the specified namespace. | 
 **index** | **int32** | If set, wait until query exceeds given index. Must be provided with WaitParam. | 
 **wait** | **string** | Provided with IndexParam to wait for change. | 
 **stale** | **string** | If present, results will include stale reads. | 
 **prefix** | **string** | Constrains results to jobs that start with the defined prefix | 
 **xNomadToken** | **string** | A Nomad ACL token. | 
 **perPage** | **int32** | Maximum number of results to return. | 
 **nextToken** | **string** | Indicates where to start paging for queries that support pagination. | 

### Return type

[**Namespace**](Namespace.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNamespaces

> []Namespace GetNamespaces(ctx).Region(region).Namespace(namespace).Index(index).Wait(wait).Stale(stale).Prefix(prefix).XNomadToken(xNomadToken).PerPage(perPage).NextToken(nextToken).Execute()



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
    index := int32(56) // int32 | If set, wait until query exceeds given index. Must be provided with WaitParam. (optional)
    wait := "wait_example" // string | Provided with IndexParam to wait for change. (optional)
    stale := "stale_example" // string | If present, results will include stale reads. (optional)
    prefix := "prefix_example" // string | Constrains results to jobs that start with the defined prefix (optional)
    xNomadToken := "xNomadToken_example" // string | A Nomad ACL token. (optional)
    perPage := int32(56) // int32 | Maximum number of results to return. (optional)
    nextToken := "nextToken_example" // string | Indicates where to start paging for queries that support pagination. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NamespacesApi.GetNamespaces(context.Background()).Region(region).Namespace(namespace).Index(index).Wait(wait).Stale(stale).Prefix(prefix).XNomadToken(xNomadToken).PerPage(perPage).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.GetNamespaces``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNamespaces`: []Namespace
    fmt.Fprintf(os.Stdout, "Response from `NamespacesApi.GetNamespaces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNamespacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **region** | **string** | Filters results based on the specified region. | 
 **namespace** | **string** | Filters results based on the specified namespace. | 
 **index** | **int32** | If set, wait until query exceeds given index. Must be provided with WaitParam. | 
 **wait** | **string** | Provided with IndexParam to wait for change. | 
 **stale** | **string** | If present, results will include stale reads. | 
 **prefix** | **string** | Constrains results to jobs that start with the defined prefix | 
 **xNomadToken** | **string** | A Nomad ACL token. | 
 **perPage** | **int32** | Maximum number of results to return. | 
 **nextToken** | **string** | Indicates where to start paging for queries that support pagination. | 

### Return type

[**[]Namespace**](Namespace.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostNamespace

> PostNamespace(ctx, namespaceName).Namespace2(namespace2).Region(region).Namespace(namespace).XNomadToken(xNomadToken).IdempotencyToken(idempotencyToken).Execute()



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
    namespaceName := "namespaceName_example" // string | The namespace identifier.
    namespace2 := *openapiclient.NewNamespace() // Namespace | 
    region := "region_example" // string | Filters results based on the specified region. (optional)
    namespace := "namespace_example" // string | Filters results based on the specified namespace. (optional)
    xNomadToken := "xNomadToken_example" // string | A Nomad ACL token. (optional)
    idempotencyToken := "idempotencyToken_example" // string | Can be used to ensure operations are only run once. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NamespacesApi.PostNamespace(context.Background(), namespaceName).Namespace2(namespace2).Region(region).Namespace(namespace).XNomadToken(xNomadToken).IdempotencyToken(idempotencyToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NamespacesApi.PostNamespace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespaceName** | **string** | The namespace identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostNamespaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **namespace2** | [**Namespace**](Namespace.md) |  | 
 **region** | **string** | Filters results based on the specified region. | 
 **namespace** | **string** | Filters results based on the specified namespace. | 
 **xNomadToken** | **string** | A Nomad ACL token. | 
 **idempotencyToken** | **string** | Can be used to ensure operations are only run once. | 

### Return type

 (empty response body)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

