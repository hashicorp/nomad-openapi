# nomad-client.AllocationsApi

All URIs are relative to *https://127.0.0.1:4646/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**getAllocation**](AllocationsApi.md#getAllocation) | **GET** /allocation/{allocationID} | 
[**getAllocations**](AllocationsApi.md#getAllocations) | **GET** /allocations | 
[**stopAllocation**](AllocationsApi.md#stopAllocation) | **POST** /allocation/{allocationID}/stop | 



## getAllocation

> Allocation getAllocation(allocationID, opts)



### Example

```javascript
import nomad-client from 'nomad';
let defaultClient = nomad-client.ApiClient.instance;
// Configure API key authorization: X-Nomad-Token
let X-Nomad-Token = defaultClient.authentications['X-Nomad-Token'];
X-Nomad-Token.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//X-Nomad-Token.apiKeyPrefix = 'Token';

let apiInstance = new nomad-client.AllocationsApi();
let allocationID = "allocationID_example"; // String | Specifies the UUID of the allocation. This must be the full UUID, not the short 8-character one. This is specified as part of the path.
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
apiInstance.getAllocation(allocationID, opts, (error, data, response) => {
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
 **allocationID** | **String**| Specifies the UUID of the allocation. This must be the full UUID, not the short 8-character one. This is specified as part of the path. | 
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

[**Allocation**](Allocation.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getAllocations

> [AllocationListStub] getAllocations(opts)



### Example

```javascript
import nomad-client from 'nomad';
let defaultClient = nomad-client.ApiClient.instance;
// Configure API key authorization: X-Nomad-Token
let X-Nomad-Token = defaultClient.authentications['X-Nomad-Token'];
X-Nomad-Token.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//X-Nomad-Token.apiKeyPrefix = 'Token';

let apiInstance = new nomad-client.AllocationsApi();
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
  'resources': true, // Boolean | Flag indicating whether to include resources in response.
  'taskStates': true // Boolean | Flag indicating whether to include task states in response.
};
apiInstance.getAllocations(opts, (error, data, response) => {
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
 **resources** | **Boolean**| Flag indicating whether to include resources in response. | [optional] 
 **taskStates** | **Boolean**| Flag indicating whether to include task states in response. | [optional] 

### Return type

[**[AllocationListStub]**](AllocationListStub.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## stopAllocation

> AllocStopResponse stopAllocation(allocationID, opts)



### Example

```javascript
import nomad-client from 'nomad';
let defaultClient = nomad-client.ApiClient.instance;
// Configure API key authorization: X-Nomad-Token
let X-Nomad-Token = defaultClient.authentications['X-Nomad-Token'];
X-Nomad-Token.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//X-Nomad-Token.apiKeyPrefix = 'Token';

let apiInstance = new nomad-client.AllocationsApi();
let allocationID = "allocationID_example"; // String | Specifies the UUID of the allocation. This must be the full UUID, not the short 8-character one. This is specified as part of the path.
let opts = {
  'region': "region_example", // String | Filters results based on the specified region.
  'namespace': "namespace_example", // String | Filters results based on the specified namespace.
  'xNomadToken': "xNomadToken_example", // String | A Nomad ACL token.
  'idempotencyToken': "idempotencyToken_example" // String | Can be used to ensure operations are only run once.
};
apiInstance.stopAllocation(allocationID, opts, (error, data, response) => {
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
 **allocationID** | **String**| Specifies the UUID of the allocation. This must be the full UUID, not the short 8-character one. This is specified as part of the path. | 
 **region** | **String**| Filters results based on the specified region. | [optional] 
 **namespace** | **String**| Filters results based on the specified namespace. | [optional] 
 **xNomadToken** | **String**| A Nomad ACL token. | [optional] 
 **idempotencyToken** | **String**| Can be used to ensure operations are only run once. | [optional] 

### Return type

[**AllocStopResponse**](AllocStopResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

