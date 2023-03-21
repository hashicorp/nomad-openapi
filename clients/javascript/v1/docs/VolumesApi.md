# nomad-client.VolumesApi

All URIs are relative to *https://127.0.0.1:4646/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createVolume**](VolumesApi.md#createVolume) | **POST** /volume/csi/{volumeId}/{action} | 
[**deleteSnapshot**](VolumesApi.md#deleteSnapshot) | **DELETE** /volumes/snapshot | 
[**deleteVolumeRegistration**](VolumesApi.md#deleteVolumeRegistration) | **DELETE** /volume/csi/{volumeId} | 
[**detachOrDeleteVolume**](VolumesApi.md#detachOrDeleteVolume) | **DELETE** /volume/csi/{volumeId}/{action} | 
[**getExternalVolumes**](VolumesApi.md#getExternalVolumes) | **GET** /volumes/external | 
[**getSnapshots**](VolumesApi.md#getSnapshots) | **GET** /volumes/snapshot | 
[**getVolume**](VolumesApi.md#getVolume) | **GET** /volume/csi/{volumeId} | 
[**getVolumes**](VolumesApi.md#getVolumes) | **GET** /volumes | 
[**postSnapshot**](VolumesApi.md#postSnapshot) | **POST** /volumes/snapshot | 
[**postVolume**](VolumesApi.md#postVolume) | **POST** /volumes | 
[**postVolumeRegistration**](VolumesApi.md#postVolumeRegistration) | **POST** /volume/csi/{volumeId} | 



## createVolume

> createVolume(volumeId, action, cSIVolumeCreateRequest, opts)



### Example

```javascript
import nomad-client from 'nomad';
let defaultClient = nomad-client.ApiClient.instance;
// Configure API key authorization: X-Nomad-Token
let X-Nomad-Token = defaultClient.authentications['X-Nomad-Token'];
X-Nomad-Token.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//X-Nomad-Token.apiKeyPrefix = 'Token';

let apiInstance = new nomad-client.VolumesApi();
let volumeId = "volumeId_example"; // String | Volume unique identifier.
let action = "action_example"; // String | The action to perform on the Volume (create, detach, delete).
let cSIVolumeCreateRequest = new nomad-client.CSIVolumeCreateRequest(); // CSIVolumeCreateRequest | 
let opts = {
  'region': "region_example", // String | Filters results based on the specified region.
  'namespace': "namespace_example", // String | Filters results based on the specified namespace.
  'xNomadToken': "xNomadToken_example", // String | A Nomad ACL token.
  'idempotencyToken': "idempotencyToken_example" // String | Can be used to ensure operations are only run once.
};
apiInstance.createVolume(volumeId, action, cSIVolumeCreateRequest, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully.');
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **volumeId** | **String**| Volume unique identifier. | 
 **action** | **String**| The action to perform on the Volume (create, detach, delete). | 
 **cSIVolumeCreateRequest** | [**CSIVolumeCreateRequest**](CSIVolumeCreateRequest.md)|  | 
 **region** | **String**| Filters results based on the specified region. | [optional] 
 **namespace** | **String**| Filters results based on the specified namespace. | [optional] 
 **xNomadToken** | **String**| A Nomad ACL token. | [optional] 
 **idempotencyToken** | **String**| Can be used to ensure operations are only run once. | [optional] 

### Return type

null (empty response body)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined


## deleteSnapshot

> deleteSnapshot(opts)



### Example

```javascript
import nomad-client from 'nomad';
let defaultClient = nomad-client.ApiClient.instance;
// Configure API key authorization: X-Nomad-Token
let X-Nomad-Token = defaultClient.authentications['X-Nomad-Token'];
X-Nomad-Token.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//X-Nomad-Token.apiKeyPrefix = 'Token';

let apiInstance = new nomad-client.VolumesApi();
let opts = {
  'region': "region_example", // String | Filters results based on the specified region.
  'namespace': "namespace_example", // String | Filters results based on the specified namespace.
  'xNomadToken': "xNomadToken_example", // String | A Nomad ACL token.
  'idempotencyToken': "idempotencyToken_example", // String | Can be used to ensure operations are only run once.
  'pluginId': "pluginId_example", // String | Filters volume lists by plugin ID.
  'snapshotId': "snapshotId_example" // String | The ID of the snapshot to target.
};
apiInstance.deleteSnapshot(opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully.');
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **region** | **String**| Filters results based on the specified region. | [optional] 
 **namespace** | **String**| Filters results based on the specified namespace. | [optional] 
 **xNomadToken** | **String**| A Nomad ACL token. | [optional] 
 **idempotencyToken** | **String**| Can be used to ensure operations are only run once. | [optional] 
 **pluginId** | **String**| Filters volume lists by plugin ID. | [optional] 
 **snapshotId** | **String**| The ID of the snapshot to target. | [optional] 

### Return type

null (empty response body)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined


## deleteVolumeRegistration

> deleteVolumeRegistration(volumeId, opts)



### Example

```javascript
import nomad-client from 'nomad';
let defaultClient = nomad-client.ApiClient.instance;
// Configure API key authorization: X-Nomad-Token
let X-Nomad-Token = defaultClient.authentications['X-Nomad-Token'];
X-Nomad-Token.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//X-Nomad-Token.apiKeyPrefix = 'Token';

let apiInstance = new nomad-client.VolumesApi();
let volumeId = "volumeId_example"; // String | Volume unique identifier.
let opts = {
  'region': "region_example", // String | Filters results based on the specified region.
  'namespace': "namespace_example", // String | Filters results based on the specified namespace.
  'xNomadToken': "xNomadToken_example", // String | A Nomad ACL token.
  'idempotencyToken': "idempotencyToken_example", // String | Can be used to ensure operations are only run once.
  'force': "force_example" // String | Used to force the de-registration of a volume.
};
apiInstance.deleteVolumeRegistration(volumeId, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully.');
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **volumeId** | **String**| Volume unique identifier. | 
 **region** | **String**| Filters results based on the specified region. | [optional] 
 **namespace** | **String**| Filters results based on the specified namespace. | [optional] 
 **xNomadToken** | **String**| A Nomad ACL token. | [optional] 
 **idempotencyToken** | **String**| Can be used to ensure operations are only run once. | [optional] 
 **force** | **String**| Used to force the de-registration of a volume. | [optional] 

### Return type

null (empty response body)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined


## detachOrDeleteVolume

> detachOrDeleteVolume(volumeId, action, opts)



### Example

```javascript
import nomad-client from 'nomad';
let defaultClient = nomad-client.ApiClient.instance;
// Configure API key authorization: X-Nomad-Token
let X-Nomad-Token = defaultClient.authentications['X-Nomad-Token'];
X-Nomad-Token.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//X-Nomad-Token.apiKeyPrefix = 'Token';

let apiInstance = new nomad-client.VolumesApi();
let volumeId = "volumeId_example"; // String | Volume unique identifier.
let action = "action_example"; // String | The action to perform on the Volume (create, detach, delete).
let opts = {
  'region': "region_example", // String | Filters results based on the specified region.
  'namespace': "namespace_example", // String | Filters results based on the specified namespace.
  'xNomadToken': "xNomadToken_example", // String | A Nomad ACL token.
  'idempotencyToken': "idempotencyToken_example", // String | Can be used to ensure operations are only run once.
  'node': "node_example" // String | Specifies node to target volume operation for.
};
apiInstance.detachOrDeleteVolume(volumeId, action, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully.');
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **volumeId** | **String**| Volume unique identifier. | 
 **action** | **String**| The action to perform on the Volume (create, detach, delete). | 
 **region** | **String**| Filters results based on the specified region. | [optional] 
 **namespace** | **String**| Filters results based on the specified namespace. | [optional] 
 **xNomadToken** | **String**| A Nomad ACL token. | [optional] 
 **idempotencyToken** | **String**| Can be used to ensure operations are only run once. | [optional] 
 **node** | **String**| Specifies node to target volume operation for. | [optional] 

### Return type

null (empty response body)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined


## getExternalVolumes

> CSIVolumeListExternalResponse getExternalVolumes(opts)



### Example

```javascript
import nomad-client from 'nomad';
let defaultClient = nomad-client.ApiClient.instance;
// Configure API key authorization: X-Nomad-Token
let X-Nomad-Token = defaultClient.authentications['X-Nomad-Token'];
X-Nomad-Token.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//X-Nomad-Token.apiKeyPrefix = 'Token';

let apiInstance = new nomad-client.VolumesApi();
let opts = {
  'region': "region_example", // String | Filters results based on the specified region.
  'namespace': "namespace_example", // String | Filters results based on the specified namespace.
  'index': 56, // Number | If set, wait until query exceeds given index. Must be provided with WaitParam.
  'wait': "wait_example", // String | Provided with IndexParam to wait for change.
  'stale': "stale_example", // String | If present, results will include stale reads.
  'prefix': "prefix_example", // String | Constrains results to jobs that start with the defined prefix
  'xNomadToken': "xNomadToken_example", // String | A Nomad ACL token.
  'perPage': 56, // Number | Maximum number of results to return.
  'nextToken': "nextToken_example", // String | Indicates where to start paging for queries that support pagination.
  'pluginId': "pluginId_example" // String | Filters volume lists by plugin ID.
};
apiInstance.getExternalVolumes(opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **region** | **String**| Filters results based on the specified region. | [optional] 
 **namespace** | **String**| Filters results based on the specified namespace. | [optional] 
 **index** | **Number**| If set, wait until query exceeds given index. Must be provided with WaitParam. | [optional] 
 **wait** | **String**| Provided with IndexParam to wait for change. | [optional] 
 **stale** | **String**| If present, results will include stale reads. | [optional] 
 **prefix** | **String**| Constrains results to jobs that start with the defined prefix | [optional] 
 **xNomadToken** | **String**| A Nomad ACL token. | [optional] 
 **perPage** | **Number**| Maximum number of results to return. | [optional] 
 **nextToken** | **String**| Indicates where to start paging for queries that support pagination. | [optional] 
 **pluginId** | **String**| Filters volume lists by plugin ID. | [optional] 

### Return type

[**CSIVolumeListExternalResponse**](CSIVolumeListExternalResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getSnapshots

> CSISnapshotListResponse getSnapshots(opts)



### Example

```javascript
import nomad-client from 'nomad';
let defaultClient = nomad-client.ApiClient.instance;
// Configure API key authorization: X-Nomad-Token
let X-Nomad-Token = defaultClient.authentications['X-Nomad-Token'];
X-Nomad-Token.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//X-Nomad-Token.apiKeyPrefix = 'Token';

let apiInstance = new nomad-client.VolumesApi();
let opts = {
  'region': "region_example", // String | Filters results based on the specified region.
  'namespace': "namespace_example", // String | Filters results based on the specified namespace.
  'index': 56, // Number | If set, wait until query exceeds given index. Must be provided with WaitParam.
  'wait': "wait_example", // String | Provided with IndexParam to wait for change.
  'stale': "stale_example", // String | If present, results will include stale reads.
  'prefix': "prefix_example", // String | Constrains results to jobs that start with the defined prefix
  'xNomadToken': "xNomadToken_example", // String | A Nomad ACL token.
  'perPage': 56, // Number | Maximum number of results to return.
  'nextToken': "nextToken_example", // String | Indicates where to start paging for queries that support pagination.
  'pluginId': "pluginId_example" // String | Filters volume lists by plugin ID.
};
apiInstance.getSnapshots(opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **region** | **String**| Filters results based on the specified region. | [optional] 
 **namespace** | **String**| Filters results based on the specified namespace. | [optional] 
 **index** | **Number**| If set, wait until query exceeds given index. Must be provided with WaitParam. | [optional] 
 **wait** | **String**| Provided with IndexParam to wait for change. | [optional] 
 **stale** | **String**| If present, results will include stale reads. | [optional] 
 **prefix** | **String**| Constrains results to jobs that start with the defined prefix | [optional] 
 **xNomadToken** | **String**| A Nomad ACL token. | [optional] 
 **perPage** | **Number**| Maximum number of results to return. | [optional] 
 **nextToken** | **String**| Indicates where to start paging for queries that support pagination. | [optional] 
 **pluginId** | **String**| Filters volume lists by plugin ID. | [optional] 

### Return type

[**CSISnapshotListResponse**](CSISnapshotListResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getVolume

> CSIVolume getVolume(volumeId, opts)



### Example

```javascript
import nomad-client from 'nomad';
let defaultClient = nomad-client.ApiClient.instance;
// Configure API key authorization: X-Nomad-Token
let X-Nomad-Token = defaultClient.authentications['X-Nomad-Token'];
X-Nomad-Token.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//X-Nomad-Token.apiKeyPrefix = 'Token';

let apiInstance = new nomad-client.VolumesApi();
let volumeId = "volumeId_example"; // String | Volume unique identifier.
let opts = {
  'region': "region_example", // String | Filters results based on the specified region.
  'namespace': "namespace_example", // String | Filters results based on the specified namespace.
  'index': 56, // Number | If set, wait until query exceeds given index. Must be provided with WaitParam.
  'wait': "wait_example", // String | Provided with IndexParam to wait for change.
  'stale': "stale_example", // String | If present, results will include stale reads.
  'prefix': "prefix_example", // String | Constrains results to jobs that start with the defined prefix
  'xNomadToken': "xNomadToken_example", // String | A Nomad ACL token.
  'perPage': 56, // Number | Maximum number of results to return.
  'nextToken': "nextToken_example" // String | Indicates where to start paging for queries that support pagination.
};
apiInstance.getVolume(volumeId, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **volumeId** | **String**| Volume unique identifier. | 
 **region** | **String**| Filters results based on the specified region. | [optional] 
 **namespace** | **String**| Filters results based on the specified namespace. | [optional] 
 **index** | **Number**| If set, wait until query exceeds given index. Must be provided with WaitParam. | [optional] 
 **wait** | **String**| Provided with IndexParam to wait for change. | [optional] 
 **stale** | **String**| If present, results will include stale reads. | [optional] 
 **prefix** | **String**| Constrains results to jobs that start with the defined prefix | [optional] 
 **xNomadToken** | **String**| A Nomad ACL token. | [optional] 
 **perPage** | **Number**| Maximum number of results to return. | [optional] 
 **nextToken** | **String**| Indicates where to start paging for queries that support pagination. | [optional] 

### Return type

[**CSIVolume**](CSIVolume.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getVolumes

> [CSIVolumeListStub] getVolumes(opts)



### Example

```javascript
import nomad-client from 'nomad';
let defaultClient = nomad-client.ApiClient.instance;
// Configure API key authorization: X-Nomad-Token
let X-Nomad-Token = defaultClient.authentications['X-Nomad-Token'];
X-Nomad-Token.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//X-Nomad-Token.apiKeyPrefix = 'Token';

let apiInstance = new nomad-client.VolumesApi();
let opts = {
  'region': "region_example", // String | Filters results based on the specified region.
  'namespace': "namespace_example", // String | Filters results based on the specified namespace.
  'index': 56, // Number | If set, wait until query exceeds given index. Must be provided with WaitParam.
  'wait': "wait_example", // String | Provided with IndexParam to wait for change.
  'stale': "stale_example", // String | If present, results will include stale reads.
  'prefix': "prefix_example", // String | Constrains results to jobs that start with the defined prefix
  'xNomadToken': "xNomadToken_example", // String | A Nomad ACL token.
  'perPage': 56, // Number | Maximum number of results to return.
  'nextToken': "nextToken_example", // String | Indicates where to start paging for queries that support pagination.
  'nodeId': "nodeId_example", // String | Filters volume lists by node ID.
  'pluginId': "pluginId_example", // String | Filters volume lists by plugin ID.
  'type': "type_example" // String | Filters volume lists to a specific type.
};
apiInstance.getVolumes(opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **region** | **String**| Filters results based on the specified region. | [optional] 
 **namespace** | **String**| Filters results based on the specified namespace. | [optional] 
 **index** | **Number**| If set, wait until query exceeds given index. Must be provided with WaitParam. | [optional] 
 **wait** | **String**| Provided with IndexParam to wait for change. | [optional] 
 **stale** | **String**| If present, results will include stale reads. | [optional] 
 **prefix** | **String**| Constrains results to jobs that start with the defined prefix | [optional] 
 **xNomadToken** | **String**| A Nomad ACL token. | [optional] 
 **perPage** | **Number**| Maximum number of results to return. | [optional] 
 **nextToken** | **String**| Indicates where to start paging for queries that support pagination. | [optional] 
 **nodeId** | **String**| Filters volume lists by node ID. | [optional] 
 **pluginId** | **String**| Filters volume lists by plugin ID. | [optional] 
 **type** | **String**| Filters volume lists to a specific type. | [optional] 

### Return type

[**[CSIVolumeListStub]**](CSIVolumeListStub.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## postSnapshot

> CSISnapshotCreateResponse postSnapshot(cSISnapshotCreateRequest, opts)



### Example

```javascript
import nomad-client from 'nomad';
let defaultClient = nomad-client.ApiClient.instance;
// Configure API key authorization: X-Nomad-Token
let X-Nomad-Token = defaultClient.authentications['X-Nomad-Token'];
X-Nomad-Token.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//X-Nomad-Token.apiKeyPrefix = 'Token';

let apiInstance = new nomad-client.VolumesApi();
let cSISnapshotCreateRequest = new nomad-client.CSISnapshotCreateRequest(); // CSISnapshotCreateRequest | 
let opts = {
  'region': "region_example", // String | Filters results based on the specified region.
  'namespace': "namespace_example", // String | Filters results based on the specified namespace.
  'xNomadToken': "xNomadToken_example", // String | A Nomad ACL token.
  'idempotencyToken': "idempotencyToken_example" // String | Can be used to ensure operations are only run once.
};
apiInstance.postSnapshot(cSISnapshotCreateRequest, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cSISnapshotCreateRequest** | [**CSISnapshotCreateRequest**](CSISnapshotCreateRequest.md)|  | 
 **region** | **String**| Filters results based on the specified region. | [optional] 
 **namespace** | **String**| Filters results based on the specified namespace. | [optional] 
 **xNomadToken** | **String**| A Nomad ACL token. | [optional] 
 **idempotencyToken** | **String**| Can be used to ensure operations are only run once. | [optional] 

### Return type

[**CSISnapshotCreateResponse**](CSISnapshotCreateResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## postVolume

> postVolume(cSIVolumeRegisterRequest, opts)



### Example

```javascript
import nomad-client from 'nomad';
let defaultClient = nomad-client.ApiClient.instance;
// Configure API key authorization: X-Nomad-Token
let X-Nomad-Token = defaultClient.authentications['X-Nomad-Token'];
X-Nomad-Token.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//X-Nomad-Token.apiKeyPrefix = 'Token';

let apiInstance = new nomad-client.VolumesApi();
let cSIVolumeRegisterRequest = new nomad-client.CSIVolumeRegisterRequest(); // CSIVolumeRegisterRequest | 
let opts = {
  'region': "region_example", // String | Filters results based on the specified region.
  'namespace': "namespace_example", // String | Filters results based on the specified namespace.
  'xNomadToken': "xNomadToken_example", // String | A Nomad ACL token.
  'idempotencyToken': "idempotencyToken_example" // String | Can be used to ensure operations are only run once.
};
apiInstance.postVolume(cSIVolumeRegisterRequest, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully.');
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cSIVolumeRegisterRequest** | [**CSIVolumeRegisterRequest**](CSIVolumeRegisterRequest.md)|  | 
 **region** | **String**| Filters results based on the specified region. | [optional] 
 **namespace** | **String**| Filters results based on the specified namespace. | [optional] 
 **xNomadToken** | **String**| A Nomad ACL token. | [optional] 
 **idempotencyToken** | **String**| Can be used to ensure operations are only run once. | [optional] 

### Return type

null (empty response body)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined


## postVolumeRegistration

> postVolumeRegistration(volumeId, cSIVolumeRegisterRequest, opts)



### Example

```javascript
import nomad-client from 'nomad';
let defaultClient = nomad-client.ApiClient.instance;
// Configure API key authorization: X-Nomad-Token
let X-Nomad-Token = defaultClient.authentications['X-Nomad-Token'];
X-Nomad-Token.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//X-Nomad-Token.apiKeyPrefix = 'Token';

let apiInstance = new nomad-client.VolumesApi();
let volumeId = "volumeId_example"; // String | Volume unique identifier.
let cSIVolumeRegisterRequest = new nomad-client.CSIVolumeRegisterRequest(); // CSIVolumeRegisterRequest | 
let opts = {
  'region': "region_example", // String | Filters results based on the specified region.
  'namespace': "namespace_example", // String | Filters results based on the specified namespace.
  'xNomadToken': "xNomadToken_example", // String | A Nomad ACL token.
  'idempotencyToken': "idempotencyToken_example" // String | Can be used to ensure operations are only run once.
};
apiInstance.postVolumeRegistration(volumeId, cSIVolumeRegisterRequest, opts, (error, data, response) => {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully.');
  }
});
```

### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **volumeId** | **String**| Volume unique identifier. | 
 **cSIVolumeRegisterRequest** | [**CSIVolumeRegisterRequest**](CSIVolumeRegisterRequest.md)|  | 
 **region** | **String**| Filters results based on the specified region. | [optional] 
 **namespace** | **String**| Filters results based on the specified namespace. | [optional] 
 **xNomadToken** | **String**| A Nomad ACL token. | [optional] 
 **idempotencyToken** | **String**| Can be used to ensure operations are only run once. | [optional] 

### Return type

null (empty response body)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

