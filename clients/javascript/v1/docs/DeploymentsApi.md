# nomad-client.DeploymentsApi

All URIs are relative to *https://127.0.0.1:4646/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**getDeployment**](DeploymentsApi.md#getDeployment) | **GET** /deployment/{deploymentID} | 
[**getDeploymentAllocations**](DeploymentsApi.md#getDeploymentAllocations) | **GET** /deployment/allocations/{deploymentID} | 
[**getDeployments**](DeploymentsApi.md#getDeployments) | **GET** /deployments | 



## getDeployment

> Deployment getDeployment(deploymentID, opts)



### Example

```javascript
import nomad-client from 'nomad';
let defaultClient = nomad-client.ApiClient.instance;
// Configure API key authorization: X-Nomad-Token
let X-Nomad-Token = defaultClient.authentications['X-Nomad-Token'];
X-Nomad-Token.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//X-Nomad-Token.apiKeyPrefix = 'Token';

let apiInstance = new nomad-client.DeploymentsApi();
let deploymentID = "deploymentID_example"; // String | Deployment ID.
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
apiInstance.getDeployment(deploymentID, opts, (error, data, response) => {
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
 **deploymentID** | **String**| Deployment ID. | 
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

[**Deployment**](Deployment.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getDeploymentAllocations

> [AllocationListStub] getDeploymentAllocations(deploymentID, opts)



### Example

```javascript
import nomad-client from 'nomad';
let defaultClient = nomad-client.ApiClient.instance;
// Configure API key authorization: X-Nomad-Token
let X-Nomad-Token = defaultClient.authentications['X-Nomad-Token'];
X-Nomad-Token.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//X-Nomad-Token.apiKeyPrefix = 'Token';

let apiInstance = new nomad-client.DeploymentsApi();
let deploymentID = "deploymentID_example"; // String | Deployment ID.
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
apiInstance.getDeploymentAllocations(deploymentID, opts, (error, data, response) => {
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
 **deploymentID** | **String**| Deployment ID. | 
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

[**[AllocationListStub]**](AllocationListStub.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getDeployments

> [Deployment] getDeployments(opts)



### Example

```javascript
import nomad-client from 'nomad';
let defaultClient = nomad-client.ApiClient.instance;
// Configure API key authorization: X-Nomad-Token
let X-Nomad-Token = defaultClient.authentications['X-Nomad-Token'];
X-Nomad-Token.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//X-Nomad-Token.apiKeyPrefix = 'Token';

let apiInstance = new nomad-client.DeploymentsApi();
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
apiInstance.getDeployments(opts, (error, data, response) => {
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

### Return type

[**[Deployment]**](Deployment.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

