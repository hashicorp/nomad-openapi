# \VolumesApi

All URIs are relative to *https://127.0.0.1:4646/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVolume**](VolumesApi.md#CreateVolume) | **Post** /volume/csi/{volumeId}/{action} | 
[**DeleteSnapshot**](VolumesApi.md#DeleteSnapshot) | **Delete** /volumes/snapshot | 
[**DeleteVolumeRegistration**](VolumesApi.md#DeleteVolumeRegistration) | **Delete** /volume/csi/{volumeId} | 
[**DetachOrDeleteVolume**](VolumesApi.md#DetachOrDeleteVolume) | **Delete** /volume/csi/{volumeId}/{action} | 
[**GetExternalVolumes**](VolumesApi.md#GetExternalVolumes) | **Get** /volumes/external | 
[**GetSnapshots**](VolumesApi.md#GetSnapshots) | **Get** /volumes/snapshot | 
[**GetVolume**](VolumesApi.md#GetVolume) | **Get** /volume/csi/{volumeId} | 
[**GetVolumes**](VolumesApi.md#GetVolumes) | **Get** /volumes | 
[**PostSnapshot**](VolumesApi.md#PostSnapshot) | **Post** /volumes/snapshot | 
[**PostVolume**](VolumesApi.md#PostVolume) | **Post** /volumes | 
[**PostVolumeRegistration**](VolumesApi.md#PostVolumeRegistration) | **Post** /volume/csi/{volumeId} | 



## CreateVolume

> CreateVolume(ctx, volumeId, action).CSIVolumeCreateRequest(cSIVolumeCreateRequest).Region(region).Namespace(namespace).XNomadToken(xNomadToken).IdempotencyToken(idempotencyToken).Execute()



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
    volumeId := "volumeId_example" // string | Volume unique identifier.
    action := "action_example" // string | The action to perform on the Volume (create, detach, delete).
    cSIVolumeCreateRequest := *openapiclient.NewCSIVolumeCreateRequest() // CSIVolumeCreateRequest | 
    region := "region_example" // string | Filters results based on the specified region. (optional)
    namespace := "namespace_example" // string | Filters results based on the specified namespace. (optional)
    xNomadToken := "xNomadToken_example" // string | A Nomad ACL token. (optional)
    idempotencyToken := "idempotencyToken_example" // string | Can be used to ensure operations are only run once. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VolumesApi.CreateVolume(context.Background(), volumeId, action).CSIVolumeCreateRequest(cSIVolumeCreateRequest).Region(region).Namespace(namespace).XNomadToken(xNomadToken).IdempotencyToken(idempotencyToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumesApi.CreateVolume``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **string** | Volume unique identifier. | 
**action** | **string** | The action to perform on the Volume (create, detach, delete). | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cSIVolumeCreateRequest** | [**CSIVolumeCreateRequest**](CSIVolumeCreateRequest.md) |  | 
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


## DeleteSnapshot

> DeleteSnapshot(ctx).Region(region).Namespace(namespace).XNomadToken(xNomadToken).IdempotencyToken(idempotencyToken).PluginId(pluginId).SnapshotId(snapshotId).Execute()



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
    pluginId := "pluginId_example" // string | Filters volume lists by plugin ID. (optional)
    snapshotId := "snapshotId_example" // string | The ID of the snapshot to target. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VolumesApi.DeleteSnapshot(context.Background()).Region(region).Namespace(namespace).XNomadToken(xNomadToken).IdempotencyToken(idempotencyToken).PluginId(pluginId).SnapshotId(snapshotId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumesApi.DeleteSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **region** | **string** | Filters results based on the specified region. | 
 **namespace** | **string** | Filters results based on the specified namespace. | 
 **xNomadToken** | **string** | A Nomad ACL token. | 
 **idempotencyToken** | **string** | Can be used to ensure operations are only run once. | 
 **pluginId** | **string** | Filters volume lists by plugin ID. | 
 **snapshotId** | **string** | The ID of the snapshot to target. | 

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


## DeleteVolumeRegistration

> DeleteVolumeRegistration(ctx, volumeId).Region(region).Namespace(namespace).XNomadToken(xNomadToken).IdempotencyToken(idempotencyToken).Force(force).Execute()



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
    volumeId := "volumeId_example" // string | Volume unique identifier.
    region := "region_example" // string | Filters results based on the specified region. (optional)
    namespace := "namespace_example" // string | Filters results based on the specified namespace. (optional)
    xNomadToken := "xNomadToken_example" // string | A Nomad ACL token. (optional)
    idempotencyToken := "idempotencyToken_example" // string | Can be used to ensure operations are only run once. (optional)
    force := "force_example" // string | Used to force the de-registration of a volume. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VolumesApi.DeleteVolumeRegistration(context.Background(), volumeId).Region(region).Namespace(namespace).XNomadToken(xNomadToken).IdempotencyToken(idempotencyToken).Force(force).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumesApi.DeleteVolumeRegistration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **string** | Volume unique identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVolumeRegistrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **region** | **string** | Filters results based on the specified region. | 
 **namespace** | **string** | Filters results based on the specified namespace. | 
 **xNomadToken** | **string** | A Nomad ACL token. | 
 **idempotencyToken** | **string** | Can be used to ensure operations are only run once. | 
 **force** | **string** | Used to force the de-registration of a volume. | 

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


## DetachOrDeleteVolume

> DetachOrDeleteVolume(ctx, volumeId, action).Region(region).Namespace(namespace).XNomadToken(xNomadToken).IdempotencyToken(idempotencyToken).Node(node).Execute()



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
    volumeId := "volumeId_example" // string | Volume unique identifier.
    action := "action_example" // string | The action to perform on the Volume (create, detach, delete).
    region := "region_example" // string | Filters results based on the specified region. (optional)
    namespace := "namespace_example" // string | Filters results based on the specified namespace. (optional)
    xNomadToken := "xNomadToken_example" // string | A Nomad ACL token. (optional)
    idempotencyToken := "idempotencyToken_example" // string | Can be used to ensure operations are only run once. (optional)
    node := "node_example" // string | Specifies node to target volume operation for. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VolumesApi.DetachOrDeleteVolume(context.Background(), volumeId, action).Region(region).Namespace(namespace).XNomadToken(xNomadToken).IdempotencyToken(idempotencyToken).Node(node).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumesApi.DetachOrDeleteVolume``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **string** | Volume unique identifier. | 
**action** | **string** | The action to perform on the Volume (create, detach, delete). | 

### Other Parameters

Other parameters are passed through a pointer to a apiDetachOrDeleteVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **region** | **string** | Filters results based on the specified region. | 
 **namespace** | **string** | Filters results based on the specified namespace. | 
 **xNomadToken** | **string** | A Nomad ACL token. | 
 **idempotencyToken** | **string** | Can be used to ensure operations are only run once. | 
 **node** | **string** | Specifies node to target volume operation for. | 

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


## GetExternalVolumes

> CSIVolumeListExternalResponse GetExternalVolumes(ctx).Region(region).Namespace(namespace).Index(index).Wait(wait).Stale(stale).Prefix(prefix).XNomadToken(xNomadToken).PerPage(perPage).NextToken(nextToken).PluginId(pluginId).Execute()



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
    pluginId := "pluginId_example" // string | Filters volume lists by plugin ID. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VolumesApi.GetExternalVolumes(context.Background()).Region(region).Namespace(namespace).Index(index).Wait(wait).Stale(stale).Prefix(prefix).XNomadToken(xNomadToken).PerPage(perPage).NextToken(nextToken).PluginId(pluginId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumesApi.GetExternalVolumes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExternalVolumes`: CSIVolumeListExternalResponse
    fmt.Fprintf(os.Stdout, "Response from `VolumesApi.GetExternalVolumes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExternalVolumesRequest struct via the builder pattern


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
 **pluginId** | **string** | Filters volume lists by plugin ID. | 

### Return type

[**CSIVolumeListExternalResponse**](CSIVolumeListExternalResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSnapshots

> CSISnapshotListResponse GetSnapshots(ctx).Region(region).Namespace(namespace).Index(index).Wait(wait).Stale(stale).Prefix(prefix).XNomadToken(xNomadToken).PerPage(perPage).NextToken(nextToken).PluginId(pluginId).Execute()



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
    pluginId := "pluginId_example" // string | Filters volume lists by plugin ID. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VolumesApi.GetSnapshots(context.Background()).Region(region).Namespace(namespace).Index(index).Wait(wait).Stale(stale).Prefix(prefix).XNomadToken(xNomadToken).PerPage(perPage).NextToken(nextToken).PluginId(pluginId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumesApi.GetSnapshots``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSnapshots`: CSISnapshotListResponse
    fmt.Fprintf(os.Stdout, "Response from `VolumesApi.GetSnapshots`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSnapshotsRequest struct via the builder pattern


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
 **pluginId** | **string** | Filters volume lists by plugin ID. | 

### Return type

[**CSISnapshotListResponse**](CSISnapshotListResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVolume

> CSIVolume GetVolume(ctx, volumeId).Region(region).Namespace(namespace).Index(index).Wait(wait).Stale(stale).Prefix(prefix).XNomadToken(xNomadToken).PerPage(perPage).NextToken(nextToken).Execute()



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
    volumeId := "volumeId_example" // string | Volume unique identifier.
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
    resp, r, err := apiClient.VolumesApi.GetVolume(context.Background(), volumeId).Region(region).Namespace(namespace).Index(index).Wait(wait).Stale(stale).Prefix(prefix).XNomadToken(xNomadToken).PerPage(perPage).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumesApi.GetVolume``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVolume`: CSIVolume
    fmt.Fprintf(os.Stdout, "Response from `VolumesApi.GetVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **string** | Volume unique identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVolumeRequest struct via the builder pattern


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

[**CSIVolume**](CSIVolume.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVolumes

> []CSIVolumeListStub GetVolumes(ctx).Region(region).Namespace(namespace).Index(index).Wait(wait).Stale(stale).Prefix(prefix).XNomadToken(xNomadToken).PerPage(perPage).NextToken(nextToken).NodeId(nodeId).PluginId(pluginId).Type_(type_).Execute()



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
    nodeId := "nodeId_example" // string | Filters volume lists by node ID. (optional)
    pluginId := "pluginId_example" // string | Filters volume lists by plugin ID. (optional)
    type_ := "type__example" // string | Filters volume lists to a specific type. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VolumesApi.GetVolumes(context.Background()).Region(region).Namespace(namespace).Index(index).Wait(wait).Stale(stale).Prefix(prefix).XNomadToken(xNomadToken).PerPage(perPage).NextToken(nextToken).NodeId(nodeId).PluginId(pluginId).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumesApi.GetVolumes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVolumes`: []CSIVolumeListStub
    fmt.Fprintf(os.Stdout, "Response from `VolumesApi.GetVolumes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVolumesRequest struct via the builder pattern


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
 **nodeId** | **string** | Filters volume lists by node ID. | 
 **pluginId** | **string** | Filters volume lists by plugin ID. | 
 **type_** | **string** | Filters volume lists to a specific type. | 

### Return type

[**[]CSIVolumeListStub**](CSIVolumeListStub.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSnapshot

> CSISnapshotCreateResponse PostSnapshot(ctx).CSISnapshotCreateRequest(cSISnapshotCreateRequest).Region(region).Namespace(namespace).XNomadToken(xNomadToken).IdempotencyToken(idempotencyToken).Execute()



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
    cSISnapshotCreateRequest := *openapiclient.NewCSISnapshotCreateRequest() // CSISnapshotCreateRequest | 
    region := "region_example" // string | Filters results based on the specified region. (optional)
    namespace := "namespace_example" // string | Filters results based on the specified namespace. (optional)
    xNomadToken := "xNomadToken_example" // string | A Nomad ACL token. (optional)
    idempotencyToken := "idempotencyToken_example" // string | Can be used to ensure operations are only run once. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VolumesApi.PostSnapshot(context.Background()).CSISnapshotCreateRequest(cSISnapshotCreateRequest).Region(region).Namespace(namespace).XNomadToken(xNomadToken).IdempotencyToken(idempotencyToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumesApi.PostSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostSnapshot`: CSISnapshotCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `VolumesApi.PostSnapshot`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cSISnapshotCreateRequest** | [**CSISnapshotCreateRequest**](CSISnapshotCreateRequest.md) |  | 
 **region** | **string** | Filters results based on the specified region. | 
 **namespace** | **string** | Filters results based on the specified namespace. | 
 **xNomadToken** | **string** | A Nomad ACL token. | 
 **idempotencyToken** | **string** | Can be used to ensure operations are only run once. | 

### Return type

[**CSISnapshotCreateResponse**](CSISnapshotCreateResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostVolume

> PostVolume(ctx).CSIVolumeRegisterRequest(cSIVolumeRegisterRequest).Region(region).Namespace(namespace).XNomadToken(xNomadToken).IdempotencyToken(idempotencyToken).Execute()



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
    cSIVolumeRegisterRequest := *openapiclient.NewCSIVolumeRegisterRequest() // CSIVolumeRegisterRequest | 
    region := "region_example" // string | Filters results based on the specified region. (optional)
    namespace := "namespace_example" // string | Filters results based on the specified namespace. (optional)
    xNomadToken := "xNomadToken_example" // string | A Nomad ACL token. (optional)
    idempotencyToken := "idempotencyToken_example" // string | Can be used to ensure operations are only run once. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VolumesApi.PostVolume(context.Background()).CSIVolumeRegisterRequest(cSIVolumeRegisterRequest).Region(region).Namespace(namespace).XNomadToken(xNomadToken).IdempotencyToken(idempotencyToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumesApi.PostVolume``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cSIVolumeRegisterRequest** | [**CSIVolumeRegisterRequest**](CSIVolumeRegisterRequest.md) |  | 
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


## PostVolumeRegistration

> PostVolumeRegistration(ctx, volumeId).CSIVolumeRegisterRequest(cSIVolumeRegisterRequest).Region(region).Namespace(namespace).XNomadToken(xNomadToken).IdempotencyToken(idempotencyToken).Execute()



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
    volumeId := "volumeId_example" // string | Volume unique identifier.
    cSIVolumeRegisterRequest := *openapiclient.NewCSIVolumeRegisterRequest() // CSIVolumeRegisterRequest | 
    region := "region_example" // string | Filters results based on the specified region. (optional)
    namespace := "namespace_example" // string | Filters results based on the specified namespace. (optional)
    xNomadToken := "xNomadToken_example" // string | A Nomad ACL token. (optional)
    idempotencyToken := "idempotencyToken_example" // string | Can be used to ensure operations are only run once. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VolumesApi.PostVolumeRegistration(context.Background(), volumeId).CSIVolumeRegisterRequest(cSIVolumeRegisterRequest).Region(region).Namespace(namespace).XNomadToken(xNomadToken).IdempotencyToken(idempotencyToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumesApi.PostVolumeRegistration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **string** | Volume unique identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostVolumeRegistrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cSIVolumeRegisterRequest** | [**CSIVolumeRegisterRequest**](CSIVolumeRegisterRequest.md) |  | 
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

