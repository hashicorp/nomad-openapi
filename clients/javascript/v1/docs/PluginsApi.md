# nomad-client.PluginsApi

All URIs are relative to *https://127.0.0.1:4646/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**getPluginCSI**](PluginsApi.md#getPluginCSI) | **GET** /plugin/csi/{pluginID} | 
[**getPlugins**](PluginsApi.md#getPlugins) | **GET** /plugins | 



## getPluginCSI

<<<<<<< HEAD
<<<<<<< HEAD
> [CSIPlugin] getPluginCSI(pluginID, opts)
=======
> [CSIPlugin] getPluginCSI(pluginId, opts)
>>>>>>> d99c3e2 (populated generator/plugins.go, added v1/plugins.go)
=======
> [CSIPlugin] getPluginCSI(pluginID, opts)
>>>>>>> 13a3eee (added tests for plugins)



### Example

```javascript
import nomad-client from 'nomad';
let defaultClient = nomad-client.ApiClient.instance;
// Configure API key authorization: X-Nomad-Token
let X-Nomad-Token = defaultClient.authentications['X-Nomad-Token'];
X-Nomad-Token.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//X-Nomad-Token.apiKeyPrefix = 'Token';

let apiInstance = new nomad-client.PluginsApi();
<<<<<<< HEAD
<<<<<<< HEAD
let pluginID = "pluginID_example"; // String | The CSI plugin identifier.
=======
let pluginId = "pluginId_example"; // String | The CSI plugin identifier.
>>>>>>> d99c3e2 (populated generator/plugins.go, added v1/plugins.go)
=======
let pluginID = "pluginID_example"; // String | The CSI plugin identifier.
>>>>>>> 13a3eee (added tests for plugins)
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
<<<<<<< HEAD
<<<<<<< HEAD
apiInstance.getPluginCSI(pluginID, opts, (error, data, response) => {
=======
apiInstance.getPluginCSI(pluginId, opts, (error, data, response) => {
>>>>>>> d99c3e2 (populated generator/plugins.go, added v1/plugins.go)
=======
apiInstance.getPluginCSI(pluginID, opts, (error, data, response) => {
>>>>>>> 13a3eee (added tests for plugins)
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
<<<<<<< HEAD
<<<<<<< HEAD
 **pluginID** | **String**| The CSI plugin identifier. | 
=======
 **pluginId** | **String**| The CSI plugin identifier. | 
>>>>>>> d99c3e2 (populated generator/plugins.go, added v1/plugins.go)
=======
 **pluginID** | **String**| The CSI plugin identifier. | 
>>>>>>> 13a3eee (added tests for plugins)
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

[**[CSIPlugin]**](CSIPlugin.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## getPlugins

> [CSIPluginListStub] getPlugins(opts)



### Example

```javascript
import nomad-client from 'nomad';
let defaultClient = nomad-client.ApiClient.instance;
// Configure API key authorization: X-Nomad-Token
let X-Nomad-Token = defaultClient.authentications['X-Nomad-Token'];
X-Nomad-Token.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//X-Nomad-Token.apiKeyPrefix = 'Token';

let apiInstance = new nomad-client.PluginsApi();
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
apiInstance.getPlugins(opts, (error, data, response) => {
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

[**[CSIPluginListStub]**](CSIPluginListStub.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

