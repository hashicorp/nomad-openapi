# \JobsApi

All URIs are relative to *https://127.0.0.1:4646/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteJob**](JobsApi.md#DeleteJob) | **Delete** /job/{jobName} | 
[**GetJob**](JobsApi.md#GetJob) | **Get** /job/{jobName} | 
[**GetJobAllocations**](JobsApi.md#GetJobAllocations) | **Get** /job/{jobName}/allocations | 
[**GetJobDeployment**](JobsApi.md#GetJobDeployment) | **Get** /job/{jobName}/deployment | 
[**GetJobDeployments**](JobsApi.md#GetJobDeployments) | **Get** /job/{jobName}/deployments | 
[**GetJobEvaluations**](JobsApi.md#GetJobEvaluations) | **Get** /job/{jobName}/evaluations | 
[**GetJobScaleStatus**](JobsApi.md#GetJobScaleStatus) | **Get** /job/{jobName}/scale | 
[**GetJobSummary**](JobsApi.md#GetJobSummary) | **Get** /job/{jobName}/summary | 
[**GetJobVersions**](JobsApi.md#GetJobVersions) | **Get** /job/{jobName}/versions | 
[**GetJobs**](JobsApi.md#GetJobs) | **Get** /jobs | 
[**PostJob**](JobsApi.md#PostJob) | **Post** /job/{jobName} | 
[**PostJobDispatch**](JobsApi.md#PostJobDispatch) | **Post** /job/{jobName}/dispatch | 
[**PostJobEvaluate**](JobsApi.md#PostJobEvaluate) | **Post** /job/{jobName}/evaluate | 
[**PostJobParse**](JobsApi.md#PostJobParse) | **Post** /jobs/parse | 
[**PostJobPeriodicForce**](JobsApi.md#PostJobPeriodicForce) | **Post** /job/{jobName}/periodic/force | 
[**PostJobPlan**](JobsApi.md#PostJobPlan) | **Post** /job/{jobName}/plan | 
[**PostJobRevert**](JobsApi.md#PostJobRevert) | **Post** /job/{jobName}/revert | 
[**PostJobScalingRequest**](JobsApi.md#PostJobScalingRequest) | **Post** /job/{jobName}/scale | 
[**PostJobStability**](JobsApi.md#PostJobStability) | **Post** /job/{jobName}/stable | 
[**PostJobValidateRequest**](JobsApi.md#PostJobValidateRequest) | **Post** /validate/job | 
[**RegisterJob**](JobsApi.md#RegisterJob) | **Post** /jobs | 



## DeleteJob

> JobDeregisterResponse DeleteJob(ctx, jobName).Region(region).Namespace(namespace).XNomadToken(xNomadToken).IdempotencyToken(idempotencyToken).Purge(purge).Global(global).Execute()



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
    jobName := "jobName_example" // string | The job identifier.
    region := "region_example" // string | Filters results based on the specified region. (optional)
    namespace := "namespace_example" // string | Filters results based on the specified namespace. (optional)
    xNomadToken := "xNomadToken_example" // string | A Nomad ACL token. (optional)
    idempotencyToken := "idempotencyToken_example" // string | Can be used to ensure operations are only run once. (optional)
    purge := true // bool | Boolean flag indicating whether to purge allocations of the job after deleting. (optional)
    global := true // bool | Boolean flag indicating whether the operation should apply to all instances of the job globally. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JobsApi.DeleteJob(context.Background(), jobName).Region(region).Namespace(namespace).XNomadToken(xNomadToken).IdempotencyToken(idempotencyToken).Purge(purge).Global(global).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobsApi.DeleteJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteJob`: JobDeregisterResponse
    fmt.Fprintf(os.Stdout, "Response from `JobsApi.DeleteJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobName** | **string** | The job identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **region** | **string** | Filters results based on the specified region. | 
 **namespace** | **string** | Filters results based on the specified namespace. | 
 **xNomadToken** | **string** | A Nomad ACL token. | 
 **idempotencyToken** | **string** | Can be used to ensure operations are only run once. | 
 **purge** | **bool** | Boolean flag indicating whether to purge allocations of the job after deleting. | 
 **global** | **bool** | Boolean flag indicating whether the operation should apply to all instances of the job globally. | 

### Return type

[**JobDeregisterResponse**](JobDeregisterResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJob

> Job GetJob(ctx, jobName).Region(region).Namespace(namespace).Index(index).Wait(wait).Stale(stale).Prefix(prefix).XNomadToken(xNomadToken).PerPage(perPage).NextToken(nextToken).Execute()



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
    jobName := "jobName_example" // string | The job identifier.
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
    resp, r, err := apiClient.JobsApi.GetJob(context.Background(), jobName).Region(region).Namespace(namespace).Index(index).Wait(wait).Stale(stale).Prefix(prefix).XNomadToken(xNomadToken).PerPage(perPage).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobsApi.GetJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJob`: Job
    fmt.Fprintf(os.Stdout, "Response from `JobsApi.GetJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobName** | **string** | The job identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJobRequest struct via the builder pattern


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

[**Job**](Job.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJobAllocations

> []AllocationListStub GetJobAllocations(ctx, jobName).Region(region).Namespace(namespace).Index(index).Wait(wait).Stale(stale).Prefix(prefix).XNomadToken(xNomadToken).PerPage(perPage).NextToken(nextToken).All(all).Execute()



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
    jobName := "jobName_example" // string | The job identifier.
    region := "region_example" // string | Filters results based on the specified region. (optional)
    namespace := "namespace_example" // string | Filters results based on the specified namespace. (optional)
    index := int32(56) // int32 | If set, wait until query exceeds given index. Must be provided with WaitParam. (optional)
    wait := "wait_example" // string | Provided with IndexParam to wait for change. (optional)
    stale := "stale_example" // string | If present, results will include stale reads. (optional)
    prefix := "prefix_example" // string | Constrains results to jobs that start with the defined prefix (optional)
    xNomadToken := "xNomadToken_example" // string | A Nomad ACL token. (optional)
    perPage := int32(56) // int32 | Maximum number of results to return. (optional)
    nextToken := "nextToken_example" // string | Indicates where to start paging for queries that support pagination. (optional)
    all := true // bool | Specifies whether the list of allocations should include allocations from a previously registered job with the same ID. This is possible if the job is deregistered and reregistered. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JobsApi.GetJobAllocations(context.Background(), jobName).Region(region).Namespace(namespace).Index(index).Wait(wait).Stale(stale).Prefix(prefix).XNomadToken(xNomadToken).PerPage(perPage).NextToken(nextToken).All(all).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobsApi.GetJobAllocations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJobAllocations`: []AllocationListStub
    fmt.Fprintf(os.Stdout, "Response from `JobsApi.GetJobAllocations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobName** | **string** | The job identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJobAllocationsRequest struct via the builder pattern


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
 **all** | **bool** | Specifies whether the list of allocations should include allocations from a previously registered job with the same ID. This is possible if the job is deregistered and reregistered. | 

### Return type

[**[]AllocationListStub**](AllocationListStub.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJobDeployment

> Deployment GetJobDeployment(ctx, jobName).Region(region).Namespace(namespace).Index(index).Wait(wait).Stale(stale).Prefix(prefix).XNomadToken(xNomadToken).PerPage(perPage).NextToken(nextToken).Execute()



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
    jobName := "jobName_example" // string | The job identifier.
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
    resp, r, err := apiClient.JobsApi.GetJobDeployment(context.Background(), jobName).Region(region).Namespace(namespace).Index(index).Wait(wait).Stale(stale).Prefix(prefix).XNomadToken(xNomadToken).PerPage(perPage).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobsApi.GetJobDeployment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJobDeployment`: Deployment
    fmt.Fprintf(os.Stdout, "Response from `JobsApi.GetJobDeployment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobName** | **string** | The job identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJobDeploymentRequest struct via the builder pattern


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

[**Deployment**](Deployment.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJobDeployments

> []Deployment GetJobDeployments(ctx, jobName).Region(region).Namespace(namespace).Index(index).Wait(wait).Stale(stale).Prefix(prefix).XNomadToken(xNomadToken).PerPage(perPage).NextToken(nextToken).All(all).Execute()



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
    jobName := "jobName_example" // string | The job identifier.
    region := "region_example" // string | Filters results based on the specified region. (optional)
    namespace := "namespace_example" // string | Filters results based on the specified namespace. (optional)
    index := int32(56) // int32 | If set, wait until query exceeds given index. Must be provided with WaitParam. (optional)
    wait := "wait_example" // string | Provided with IndexParam to wait for change. (optional)
    stale := "stale_example" // string | If present, results will include stale reads. (optional)
    prefix := "prefix_example" // string | Constrains results to jobs that start with the defined prefix (optional)
    xNomadToken := "xNomadToken_example" // string | A Nomad ACL token. (optional)
    perPage := int32(56) // int32 | Maximum number of results to return. (optional)
    nextToken := "nextToken_example" // string | Indicates where to start paging for queries that support pagination. (optional)
    all := int32(56) // int32 | Flag indicating whether to constrain by job creation index or not. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JobsApi.GetJobDeployments(context.Background(), jobName).Region(region).Namespace(namespace).Index(index).Wait(wait).Stale(stale).Prefix(prefix).XNomadToken(xNomadToken).PerPage(perPage).NextToken(nextToken).All(all).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobsApi.GetJobDeployments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJobDeployments`: []Deployment
    fmt.Fprintf(os.Stdout, "Response from `JobsApi.GetJobDeployments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobName** | **string** | The job identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJobDeploymentsRequest struct via the builder pattern


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
 **all** | **int32** | Flag indicating whether to constrain by job creation index or not. | 

### Return type

[**[]Deployment**](Deployment.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJobEvaluations

> []Evaluation GetJobEvaluations(ctx, jobName).Region(region).Namespace(namespace).Index(index).Wait(wait).Stale(stale).Prefix(prefix).XNomadToken(xNomadToken).PerPage(perPage).NextToken(nextToken).Execute()



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
    jobName := "jobName_example" // string | The job identifier.
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
    resp, r, err := apiClient.JobsApi.GetJobEvaluations(context.Background(), jobName).Region(region).Namespace(namespace).Index(index).Wait(wait).Stale(stale).Prefix(prefix).XNomadToken(xNomadToken).PerPage(perPage).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobsApi.GetJobEvaluations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJobEvaluations`: []Evaluation
    fmt.Fprintf(os.Stdout, "Response from `JobsApi.GetJobEvaluations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobName** | **string** | The job identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJobEvaluationsRequest struct via the builder pattern


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

[**[]Evaluation**](Evaluation.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJobScaleStatus

> JobScaleStatusResponse GetJobScaleStatus(ctx, jobName).Region(region).Namespace(namespace).Index(index).Wait(wait).Stale(stale).Prefix(prefix).XNomadToken(xNomadToken).PerPage(perPage).NextToken(nextToken).Execute()



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
    jobName := "jobName_example" // string | The job identifier.
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
    resp, r, err := apiClient.JobsApi.GetJobScaleStatus(context.Background(), jobName).Region(region).Namespace(namespace).Index(index).Wait(wait).Stale(stale).Prefix(prefix).XNomadToken(xNomadToken).PerPage(perPage).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobsApi.GetJobScaleStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJobScaleStatus`: JobScaleStatusResponse
    fmt.Fprintf(os.Stdout, "Response from `JobsApi.GetJobScaleStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobName** | **string** | The job identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJobScaleStatusRequest struct via the builder pattern


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

[**JobScaleStatusResponse**](JobScaleStatusResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJobSummary

> JobSummary GetJobSummary(ctx, jobName).Region(region).Namespace(namespace).Index(index).Wait(wait).Stale(stale).Prefix(prefix).XNomadToken(xNomadToken).PerPage(perPage).NextToken(nextToken).Execute()



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
    jobName := "jobName_example" // string | The job identifier.
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
    resp, r, err := apiClient.JobsApi.GetJobSummary(context.Background(), jobName).Region(region).Namespace(namespace).Index(index).Wait(wait).Stale(stale).Prefix(prefix).XNomadToken(xNomadToken).PerPage(perPage).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobsApi.GetJobSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJobSummary`: JobSummary
    fmt.Fprintf(os.Stdout, "Response from `JobsApi.GetJobSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobName** | **string** | The job identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJobSummaryRequest struct via the builder pattern


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

[**JobSummary**](JobSummary.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJobVersions

> JobVersionsResponse GetJobVersions(ctx, jobName).Region(region).Namespace(namespace).Index(index).Wait(wait).Stale(stale).Prefix(prefix).XNomadToken(xNomadToken).PerPage(perPage).NextToken(nextToken).Diffs(diffs).Execute()



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
    jobName := "jobName_example" // string | The job identifier.
    region := "region_example" // string | Filters results based on the specified region. (optional)
    namespace := "namespace_example" // string | Filters results based on the specified namespace. (optional)
    index := int32(56) // int32 | If set, wait until query exceeds given index. Must be provided with WaitParam. (optional)
    wait := "wait_example" // string | Provided with IndexParam to wait for change. (optional)
    stale := "stale_example" // string | If present, results will include stale reads. (optional)
    prefix := "prefix_example" // string | Constrains results to jobs that start with the defined prefix (optional)
    xNomadToken := "xNomadToken_example" // string | A Nomad ACL token. (optional)
    perPage := int32(56) // int32 | Maximum number of results to return. (optional)
    nextToken := "nextToken_example" // string | Indicates where to start paging for queries that support pagination. (optional)
    diffs := true // bool | Boolean flag indicating whether to compute job diffs. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JobsApi.GetJobVersions(context.Background(), jobName).Region(region).Namespace(namespace).Index(index).Wait(wait).Stale(stale).Prefix(prefix).XNomadToken(xNomadToken).PerPage(perPage).NextToken(nextToken).Diffs(diffs).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobsApi.GetJobVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJobVersions`: JobVersionsResponse
    fmt.Fprintf(os.Stdout, "Response from `JobsApi.GetJobVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobName** | **string** | The job identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJobVersionsRequest struct via the builder pattern


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
 **diffs** | **bool** | Boolean flag indicating whether to compute job diffs. | 

### Return type

[**JobVersionsResponse**](JobVersionsResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJobs

> []JobListStub GetJobs(ctx).Region(region).Namespace(namespace).Index(index).Wait(wait).Stale(stale).Prefix(prefix).XNomadToken(xNomadToken).PerPage(perPage).NextToken(nextToken).Execute()



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
    resp, r, err := apiClient.JobsApi.GetJobs(context.Background()).Region(region).Namespace(namespace).Index(index).Wait(wait).Stale(stale).Prefix(prefix).XNomadToken(xNomadToken).PerPage(perPage).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobsApi.GetJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJobs`: []JobListStub
    fmt.Fprintf(os.Stdout, "Response from `JobsApi.GetJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetJobsRequest struct via the builder pattern


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

[**[]JobListStub**](JobListStub.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostJob

> JobRegisterResponse PostJob(ctx, jobName).JobRegisterRequest(jobRegisterRequest).Region(region).Namespace(namespace).XNomadToken(xNomadToken).IdempotencyToken(idempotencyToken).Execute()



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
    jobName := "jobName_example" // string | The job identifier.
    jobRegisterRequest := *openapiclient.NewJobRegisterRequest() // JobRegisterRequest | 
    region := "region_example" // string | Filters results based on the specified region. (optional)
    namespace := "namespace_example" // string | Filters results based on the specified namespace. (optional)
    xNomadToken := "xNomadToken_example" // string | A Nomad ACL token. (optional)
    idempotencyToken := "idempotencyToken_example" // string | Can be used to ensure operations are only run once. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JobsApi.PostJob(context.Background(), jobName).JobRegisterRequest(jobRegisterRequest).Region(region).Namespace(namespace).XNomadToken(xNomadToken).IdempotencyToken(idempotencyToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobsApi.PostJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostJob`: JobRegisterResponse
    fmt.Fprintf(os.Stdout, "Response from `JobsApi.PostJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobName** | **string** | The job identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jobRegisterRequest** | [**JobRegisterRequest**](JobRegisterRequest.md) |  | 
 **region** | **string** | Filters results based on the specified region. | 
 **namespace** | **string** | Filters results based on the specified namespace. | 
 **xNomadToken** | **string** | A Nomad ACL token. | 
 **idempotencyToken** | **string** | Can be used to ensure operations are only run once. | 

### Return type

[**JobRegisterResponse**](JobRegisterResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostJobDispatch

> JobDispatchResponse PostJobDispatch(ctx, jobName).JobDispatchRequest(jobDispatchRequest).Region(region).Namespace(namespace).XNomadToken(xNomadToken).IdempotencyToken(idempotencyToken).Execute()



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
    jobName := "jobName_example" // string | The job identifier.
    jobDispatchRequest := *openapiclient.NewJobDispatchRequest() // JobDispatchRequest | 
    region := "region_example" // string | Filters results based on the specified region. (optional)
    namespace := "namespace_example" // string | Filters results based on the specified namespace. (optional)
    xNomadToken := "xNomadToken_example" // string | A Nomad ACL token. (optional)
    idempotencyToken := "idempotencyToken_example" // string | Can be used to ensure operations are only run once. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JobsApi.PostJobDispatch(context.Background(), jobName).JobDispatchRequest(jobDispatchRequest).Region(region).Namespace(namespace).XNomadToken(xNomadToken).IdempotencyToken(idempotencyToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobsApi.PostJobDispatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostJobDispatch`: JobDispatchResponse
    fmt.Fprintf(os.Stdout, "Response from `JobsApi.PostJobDispatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobName** | **string** | The job identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostJobDispatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jobDispatchRequest** | [**JobDispatchRequest**](JobDispatchRequest.md) |  | 
 **region** | **string** | Filters results based on the specified region. | 
 **namespace** | **string** | Filters results based on the specified namespace. | 
 **xNomadToken** | **string** | A Nomad ACL token. | 
 **idempotencyToken** | **string** | Can be used to ensure operations are only run once. | 

### Return type

[**JobDispatchResponse**](JobDispatchResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostJobEvaluate

> JobRegisterResponse PostJobEvaluate(ctx, jobName).JobEvaluateRequest(jobEvaluateRequest).Region(region).Namespace(namespace).XNomadToken(xNomadToken).IdempotencyToken(idempotencyToken).Execute()



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
    jobName := "jobName_example" // string | The job identifier.
    jobEvaluateRequest := *openapiclient.NewJobEvaluateRequest() // JobEvaluateRequest | 
    region := "region_example" // string | Filters results based on the specified region. (optional)
    namespace := "namespace_example" // string | Filters results based on the specified namespace. (optional)
    xNomadToken := "xNomadToken_example" // string | A Nomad ACL token. (optional)
    idempotencyToken := "idempotencyToken_example" // string | Can be used to ensure operations are only run once. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JobsApi.PostJobEvaluate(context.Background(), jobName).JobEvaluateRequest(jobEvaluateRequest).Region(region).Namespace(namespace).XNomadToken(xNomadToken).IdempotencyToken(idempotencyToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobsApi.PostJobEvaluate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostJobEvaluate`: JobRegisterResponse
    fmt.Fprintf(os.Stdout, "Response from `JobsApi.PostJobEvaluate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobName** | **string** | The job identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostJobEvaluateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jobEvaluateRequest** | [**JobEvaluateRequest**](JobEvaluateRequest.md) |  | 
 **region** | **string** | Filters results based on the specified region. | 
 **namespace** | **string** | Filters results based on the specified namespace. | 
 **xNomadToken** | **string** | A Nomad ACL token. | 
 **idempotencyToken** | **string** | Can be used to ensure operations are only run once. | 

### Return type

[**JobRegisterResponse**](JobRegisterResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostJobParse

> Job PostJobParse(ctx).JobsParseRequest(jobsParseRequest).Execute()



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
    jobsParseRequest := *openapiclient.NewJobsParseRequest() // JobsParseRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JobsApi.PostJobParse(context.Background()).JobsParseRequest(jobsParseRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobsApi.PostJobParse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostJobParse`: Job
    fmt.Fprintf(os.Stdout, "Response from `JobsApi.PostJobParse`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostJobParseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jobsParseRequest** | [**JobsParseRequest**](JobsParseRequest.md) |  | 

### Return type

[**Job**](Job.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostJobPeriodicForce

> PeriodicForceResponse PostJobPeriodicForce(ctx, jobName).Region(region).Namespace(namespace).XNomadToken(xNomadToken).IdempotencyToken(idempotencyToken).Execute()



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
    jobName := "jobName_example" // string | The job identifier.
    region := "region_example" // string | Filters results based on the specified region. (optional)
    namespace := "namespace_example" // string | Filters results based on the specified namespace. (optional)
    xNomadToken := "xNomadToken_example" // string | A Nomad ACL token. (optional)
    idempotencyToken := "idempotencyToken_example" // string | Can be used to ensure operations are only run once. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JobsApi.PostJobPeriodicForce(context.Background(), jobName).Region(region).Namespace(namespace).XNomadToken(xNomadToken).IdempotencyToken(idempotencyToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobsApi.PostJobPeriodicForce``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostJobPeriodicForce`: PeriodicForceResponse
    fmt.Fprintf(os.Stdout, "Response from `JobsApi.PostJobPeriodicForce`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobName** | **string** | The job identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostJobPeriodicForceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **region** | **string** | Filters results based on the specified region. | 
 **namespace** | **string** | Filters results based on the specified namespace. | 
 **xNomadToken** | **string** | A Nomad ACL token. | 
 **idempotencyToken** | **string** | Can be used to ensure operations are only run once. | 

### Return type

[**PeriodicForceResponse**](PeriodicForceResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostJobPlan

> JobPlanResponse PostJobPlan(ctx, jobName).JobPlanRequest(jobPlanRequest).Region(region).Namespace(namespace).XNomadToken(xNomadToken).IdempotencyToken(idempotencyToken).Execute()



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
    jobName := "jobName_example" // string | The job identifier.
    jobPlanRequest := *openapiclient.NewJobPlanRequest() // JobPlanRequest | 
    region := "region_example" // string | Filters results based on the specified region. (optional)
    namespace := "namespace_example" // string | Filters results based on the specified namespace. (optional)
    xNomadToken := "xNomadToken_example" // string | A Nomad ACL token. (optional)
    idempotencyToken := "idempotencyToken_example" // string | Can be used to ensure operations are only run once. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JobsApi.PostJobPlan(context.Background(), jobName).JobPlanRequest(jobPlanRequest).Region(region).Namespace(namespace).XNomadToken(xNomadToken).IdempotencyToken(idempotencyToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobsApi.PostJobPlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostJobPlan`: JobPlanResponse
    fmt.Fprintf(os.Stdout, "Response from `JobsApi.PostJobPlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobName** | **string** | The job identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostJobPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jobPlanRequest** | [**JobPlanRequest**](JobPlanRequest.md) |  | 
 **region** | **string** | Filters results based on the specified region. | 
 **namespace** | **string** | Filters results based on the specified namespace. | 
 **xNomadToken** | **string** | A Nomad ACL token. | 
 **idempotencyToken** | **string** | Can be used to ensure operations are only run once. | 

### Return type

[**JobPlanResponse**](JobPlanResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostJobRevert

> JobRegisterResponse PostJobRevert(ctx, jobName).JobRevertRequest(jobRevertRequest).Region(region).Namespace(namespace).XNomadToken(xNomadToken).IdempotencyToken(idempotencyToken).Execute()



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
    jobName := "jobName_example" // string | The job identifier.
    jobRevertRequest := *openapiclient.NewJobRevertRequest() // JobRevertRequest | 
    region := "region_example" // string | Filters results based on the specified region. (optional)
    namespace := "namespace_example" // string | Filters results based on the specified namespace. (optional)
    xNomadToken := "xNomadToken_example" // string | A Nomad ACL token. (optional)
    idempotencyToken := "idempotencyToken_example" // string | Can be used to ensure operations are only run once. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JobsApi.PostJobRevert(context.Background(), jobName).JobRevertRequest(jobRevertRequest).Region(region).Namespace(namespace).XNomadToken(xNomadToken).IdempotencyToken(idempotencyToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobsApi.PostJobRevert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostJobRevert`: JobRegisterResponse
    fmt.Fprintf(os.Stdout, "Response from `JobsApi.PostJobRevert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobName** | **string** | The job identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostJobRevertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jobRevertRequest** | [**JobRevertRequest**](JobRevertRequest.md) |  | 
 **region** | **string** | Filters results based on the specified region. | 
 **namespace** | **string** | Filters results based on the specified namespace. | 
 **xNomadToken** | **string** | A Nomad ACL token. | 
 **idempotencyToken** | **string** | Can be used to ensure operations are only run once. | 

### Return type

[**JobRegisterResponse**](JobRegisterResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostJobScalingRequest

> JobRegisterResponse PostJobScalingRequest(ctx, jobName).ScalingRequest(scalingRequest).Region(region).Namespace(namespace).XNomadToken(xNomadToken).IdempotencyToken(idempotencyToken).Execute()



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
    jobName := "jobName_example" // string | The job identifier.
    scalingRequest := *openapiclient.NewScalingRequest() // ScalingRequest | 
    region := "region_example" // string | Filters results based on the specified region. (optional)
    namespace := "namespace_example" // string | Filters results based on the specified namespace. (optional)
    xNomadToken := "xNomadToken_example" // string | A Nomad ACL token. (optional)
    idempotencyToken := "idempotencyToken_example" // string | Can be used to ensure operations are only run once. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JobsApi.PostJobScalingRequest(context.Background(), jobName).ScalingRequest(scalingRequest).Region(region).Namespace(namespace).XNomadToken(xNomadToken).IdempotencyToken(idempotencyToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobsApi.PostJobScalingRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostJobScalingRequest`: JobRegisterResponse
    fmt.Fprintf(os.Stdout, "Response from `JobsApi.PostJobScalingRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobName** | **string** | The job identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostJobScalingRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scalingRequest** | [**ScalingRequest**](ScalingRequest.md) |  | 
 **region** | **string** | Filters results based on the specified region. | 
 **namespace** | **string** | Filters results based on the specified namespace. | 
 **xNomadToken** | **string** | A Nomad ACL token. | 
 **idempotencyToken** | **string** | Can be used to ensure operations are only run once. | 

### Return type

[**JobRegisterResponse**](JobRegisterResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostJobStability

> JobStabilityResponse PostJobStability(ctx, jobName).JobStabilityRequest(jobStabilityRequest).Region(region).Namespace(namespace).XNomadToken(xNomadToken).IdempotencyToken(idempotencyToken).Execute()



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
    jobName := "jobName_example" // string | The job identifier.
    jobStabilityRequest := *openapiclient.NewJobStabilityRequest() // JobStabilityRequest | 
    region := "region_example" // string | Filters results based on the specified region. (optional)
    namespace := "namespace_example" // string | Filters results based on the specified namespace. (optional)
    xNomadToken := "xNomadToken_example" // string | A Nomad ACL token. (optional)
    idempotencyToken := "idempotencyToken_example" // string | Can be used to ensure operations are only run once. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JobsApi.PostJobStability(context.Background(), jobName).JobStabilityRequest(jobStabilityRequest).Region(region).Namespace(namespace).XNomadToken(xNomadToken).IdempotencyToken(idempotencyToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobsApi.PostJobStability``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostJobStability`: JobStabilityResponse
    fmt.Fprintf(os.Stdout, "Response from `JobsApi.PostJobStability`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobName** | **string** | The job identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostJobStabilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **jobStabilityRequest** | [**JobStabilityRequest**](JobStabilityRequest.md) |  | 
 **region** | **string** | Filters results based on the specified region. | 
 **namespace** | **string** | Filters results based on the specified namespace. | 
 **xNomadToken** | **string** | A Nomad ACL token. | 
 **idempotencyToken** | **string** | Can be used to ensure operations are only run once. | 

### Return type

[**JobStabilityResponse**](JobStabilityResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostJobValidateRequest

> JobValidateResponse PostJobValidateRequest(ctx).JobValidateRequest(jobValidateRequest).Region(region).Namespace(namespace).XNomadToken(xNomadToken).IdempotencyToken(idempotencyToken).Execute()



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
    jobValidateRequest := *openapiclient.NewJobValidateRequest() // JobValidateRequest | 
    region := "region_example" // string | Filters results based on the specified region. (optional)
    namespace := "namespace_example" // string | Filters results based on the specified namespace. (optional)
    xNomadToken := "xNomadToken_example" // string | A Nomad ACL token. (optional)
    idempotencyToken := "idempotencyToken_example" // string | Can be used to ensure operations are only run once. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JobsApi.PostJobValidateRequest(context.Background()).JobValidateRequest(jobValidateRequest).Region(region).Namespace(namespace).XNomadToken(xNomadToken).IdempotencyToken(idempotencyToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobsApi.PostJobValidateRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostJobValidateRequest`: JobValidateResponse
    fmt.Fprintf(os.Stdout, "Response from `JobsApi.PostJobValidateRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostJobValidateRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jobValidateRequest** | [**JobValidateRequest**](JobValidateRequest.md) |  | 
 **region** | **string** | Filters results based on the specified region. | 
 **namespace** | **string** | Filters results based on the specified namespace. | 
 **xNomadToken** | **string** | A Nomad ACL token. | 
 **idempotencyToken** | **string** | Can be used to ensure operations are only run once. | 

### Return type

[**JobValidateResponse**](JobValidateResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterJob

> JobRegisterResponse RegisterJob(ctx).JobRegisterRequest(jobRegisterRequest).Region(region).Namespace(namespace).XNomadToken(xNomadToken).IdempotencyToken(idempotencyToken).Execute()



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
    jobRegisterRequest := *openapiclient.NewJobRegisterRequest() // JobRegisterRequest | 
    region := "region_example" // string | Filters results based on the specified region. (optional)
    namespace := "namespace_example" // string | Filters results based on the specified namespace. (optional)
    xNomadToken := "xNomadToken_example" // string | A Nomad ACL token. (optional)
    idempotencyToken := "idempotencyToken_example" // string | Can be used to ensure operations are only run once. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JobsApi.RegisterJob(context.Background()).JobRegisterRequest(jobRegisterRequest).Region(region).Namespace(namespace).XNomadToken(xNomadToken).IdempotencyToken(idempotencyToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobsApi.RegisterJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegisterJob`: JobRegisterResponse
    fmt.Fprintf(os.Stdout, "Response from `JobsApi.RegisterJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jobRegisterRequest** | [**JobRegisterRequest**](JobRegisterRequest.md) |  | 
 **region** | **string** | Filters results based on the specified region. | 
 **namespace** | **string** | Filters results based on the specified namespace. | 
 **xNomadToken** | **string** | A Nomad ACL token. | 
 **idempotencyToken** | **string** | Can be used to ensure operations are only run once. | 

### Return type

[**JobRegisterResponse**](JobRegisterResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

