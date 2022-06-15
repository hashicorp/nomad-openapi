# \OperatorApi

All URIs are relative to *https://127.0.0.1:4646/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteOperatorRaftPeer**](OperatorApi.md#DeleteOperatorRaftPeer) | **Delete** /operator/raft/peer | 
[**GetOperatorAutopilotConfiguration**](OperatorApi.md#GetOperatorAutopilotConfiguration) | **Get** /operator/autopilot/configuration | 
[**GetOperatorAutopilotHealth**](OperatorApi.md#GetOperatorAutopilotHealth) | **Get** /operator/autopilot/health | 
[**GetOperatorRaftConfiguration**](OperatorApi.md#GetOperatorRaftConfiguration) | **Get** /operator/raft/configuration | 
[**GetOperatorSchedulerConfiguration**](OperatorApi.md#GetOperatorSchedulerConfiguration) | **Get** /operator/scheduler/configuration | 
[**PostOperatorSchedulerConfiguration**](OperatorApi.md#PostOperatorSchedulerConfiguration) | **Post** /operator/scheduler/configuration | 
[**PutOperatorAutopilotConfiguration**](OperatorApi.md#PutOperatorAutopilotConfiguration) | **Put** /operator/autopilot/configuration | 



## DeleteOperatorRaftPeer

> DeleteOperatorRaftPeer(ctx).Region(region).Namespace(namespace).XNomadToken(xNomadToken).IdempotencyToken(idempotencyToken).Execute()



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
    resp, r, err := apiClient.OperatorApi.DeleteOperatorRaftPeer(context.Background()).Region(region).Namespace(namespace).XNomadToken(xNomadToken).IdempotencyToken(idempotencyToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatorApi.DeleteOperatorRaftPeer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOperatorRaftPeerRequest struct via the builder pattern


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


## GetOperatorAutopilotConfiguration

> AutopilotConfiguration GetOperatorAutopilotConfiguration(ctx).Region(region).Namespace(namespace).Index(index).Wait(wait).Stale(stale).Prefix(prefix).XNomadToken(xNomadToken).PerPage(perPage).NextToken(nextToken).Execute()



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
    resp, r, err := apiClient.OperatorApi.GetOperatorAutopilotConfiguration(context.Background()).Region(region).Namespace(namespace).Index(index).Wait(wait).Stale(stale).Prefix(prefix).XNomadToken(xNomadToken).PerPage(perPage).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatorApi.GetOperatorAutopilotConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOperatorAutopilotConfiguration`: AutopilotConfiguration
    fmt.Fprintf(os.Stdout, "Response from `OperatorApi.GetOperatorAutopilotConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOperatorAutopilotConfigurationRequest struct via the builder pattern


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

[**AutopilotConfiguration**](AutopilotConfiguration.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOperatorAutopilotHealth

> OperatorHealthReply GetOperatorAutopilotHealth(ctx).Region(region).Namespace(namespace).Index(index).Wait(wait).Stale(stale).Prefix(prefix).XNomadToken(xNomadToken).PerPage(perPage).NextToken(nextToken).Execute()



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
    resp, r, err := apiClient.OperatorApi.GetOperatorAutopilotHealth(context.Background()).Region(region).Namespace(namespace).Index(index).Wait(wait).Stale(stale).Prefix(prefix).XNomadToken(xNomadToken).PerPage(perPage).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatorApi.GetOperatorAutopilotHealth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOperatorAutopilotHealth`: OperatorHealthReply
    fmt.Fprintf(os.Stdout, "Response from `OperatorApi.GetOperatorAutopilotHealth`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOperatorAutopilotHealthRequest struct via the builder pattern


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

[**OperatorHealthReply**](OperatorHealthReply.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOperatorRaftConfiguration

> RaftConfiguration GetOperatorRaftConfiguration(ctx).Region(region).Namespace(namespace).Index(index).Wait(wait).Stale(stale).Prefix(prefix).XNomadToken(xNomadToken).PerPage(perPage).NextToken(nextToken).Execute()



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
    resp, r, err := apiClient.OperatorApi.GetOperatorRaftConfiguration(context.Background()).Region(region).Namespace(namespace).Index(index).Wait(wait).Stale(stale).Prefix(prefix).XNomadToken(xNomadToken).PerPage(perPage).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatorApi.GetOperatorRaftConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOperatorRaftConfiguration`: RaftConfiguration
    fmt.Fprintf(os.Stdout, "Response from `OperatorApi.GetOperatorRaftConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOperatorRaftConfigurationRequest struct via the builder pattern


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

[**RaftConfiguration**](RaftConfiguration.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOperatorSchedulerConfiguration

> SchedulerConfigurationResponse GetOperatorSchedulerConfiguration(ctx).Region(region).Namespace(namespace).Index(index).Wait(wait).Stale(stale).Prefix(prefix).XNomadToken(xNomadToken).PerPage(perPage).NextToken(nextToken).Execute()



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
    resp, r, err := apiClient.OperatorApi.GetOperatorSchedulerConfiguration(context.Background()).Region(region).Namespace(namespace).Index(index).Wait(wait).Stale(stale).Prefix(prefix).XNomadToken(xNomadToken).PerPage(perPage).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatorApi.GetOperatorSchedulerConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOperatorSchedulerConfiguration`: SchedulerConfigurationResponse
    fmt.Fprintf(os.Stdout, "Response from `OperatorApi.GetOperatorSchedulerConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOperatorSchedulerConfigurationRequest struct via the builder pattern


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

[**SchedulerConfigurationResponse**](SchedulerConfigurationResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOperatorSchedulerConfiguration

> SchedulerSetConfigurationResponse PostOperatorSchedulerConfiguration(ctx).SchedulerConfiguration(schedulerConfiguration).Region(region).Namespace(namespace).XNomadToken(xNomadToken).IdempotencyToken(idempotencyToken).Execute()



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
    schedulerConfiguration := *openapiclient.NewSchedulerConfiguration() // SchedulerConfiguration | 
    region := "region_example" // string | Filters results based on the specified region. (optional)
    namespace := "namespace_example" // string | Filters results based on the specified namespace. (optional)
    xNomadToken := "xNomadToken_example" // string | A Nomad ACL token. (optional)
    idempotencyToken := "idempotencyToken_example" // string | Can be used to ensure operations are only run once. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OperatorApi.PostOperatorSchedulerConfiguration(context.Background()).SchedulerConfiguration(schedulerConfiguration).Region(region).Namespace(namespace).XNomadToken(xNomadToken).IdempotencyToken(idempotencyToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatorApi.PostOperatorSchedulerConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostOperatorSchedulerConfiguration`: SchedulerSetConfigurationResponse
    fmt.Fprintf(os.Stdout, "Response from `OperatorApi.PostOperatorSchedulerConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostOperatorSchedulerConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **schedulerConfiguration** | [**SchedulerConfiguration**](SchedulerConfiguration.md) |  | 
 **region** | **string** | Filters results based on the specified region. | 
 **namespace** | **string** | Filters results based on the specified namespace. | 
 **xNomadToken** | **string** | A Nomad ACL token. | 
 **idempotencyToken** | **string** | Can be used to ensure operations are only run once. | 

### Return type

[**SchedulerSetConfigurationResponse**](SchedulerSetConfigurationResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutOperatorAutopilotConfiguration

> bool PutOperatorAutopilotConfiguration(ctx).AutopilotConfiguration(autopilotConfiguration).Region(region).Namespace(namespace).XNomadToken(xNomadToken).IdempotencyToken(idempotencyToken).Execute()



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
    autopilotConfiguration := *openapiclient.NewAutopilotConfiguration() // AutopilotConfiguration | 
    region := "region_example" // string | Filters results based on the specified region. (optional)
    namespace := "namespace_example" // string | Filters results based on the specified namespace. (optional)
    xNomadToken := "xNomadToken_example" // string | A Nomad ACL token. (optional)
    idempotencyToken := "idempotencyToken_example" // string | Can be used to ensure operations are only run once. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OperatorApi.PutOperatorAutopilotConfiguration(context.Background()).AutopilotConfiguration(autopilotConfiguration).Region(region).Namespace(namespace).XNomadToken(xNomadToken).IdempotencyToken(idempotencyToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatorApi.PutOperatorAutopilotConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutOperatorAutopilotConfiguration`: bool
    fmt.Fprintf(os.Stdout, "Response from `OperatorApi.PutOperatorAutopilotConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutOperatorAutopilotConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **autopilotConfiguration** | [**AutopilotConfiguration**](AutopilotConfiguration.md) |  | 
 **region** | **string** | Filters results based on the specified region. | 
 **namespace** | **string** | Filters results based on the specified namespace. | 
 **xNomadToken** | **string** | A Nomad ACL token. | 
 **idempotencyToken** | **string** | Can be used to ensure operations are only run once. | 

### Return type

**bool**

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

