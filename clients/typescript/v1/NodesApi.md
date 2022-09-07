# nomad-client.NodesApi

All URIs are relative to *https://127.0.0.1:4646/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**getNode**](NodesApi.md#getNode) | **GET** /node/{nodeId} | 
[**getNodeAllocations**](NodesApi.md#getNodeAllocations) | **GET** /node/{nodeId}/allocations | 
[**getNodes**](NodesApi.md#getNodes) | **GET** /nodes | 
[**updateNodeDrain**](NodesApi.md#updateNodeDrain) | **POST** /node/{nodeId}/drain | 
[**updateNodeEligibility**](NodesApi.md#updateNodeEligibility) | **POST** /node/{nodeId}/eligibility | 
[**updateNodePurge**](NodesApi.md#updateNodePurge) | **POST** /node/{nodeId}/purge | 


# **getNode**
> Node getNode()


### Example


```typescript
import { nomad-client } from '';
import * as fs from 'fs';

const configuration = nomad-client.createConfiguration();
const apiInstance = new nomad-client.NodesApi(configuration);

let body:nomad-client.NodesApiGetNodeRequest = {
  // string | The ID of the node.
  nodeId: "nodeId_example",
  // string | Filters results based on the specified region. (optional)
  region: "region_example",
  // string | Filters results based on the specified namespace. (optional)
  namespace: "namespace_example",
  // number | If set, wait until query exceeds given index. Must be provided with WaitParam. (optional)
  index: 1,
  // string | Provided with IndexParam to wait for change. (optional)
  wait: "wait_example",
  // string | If present, results will include stale reads. (optional)
  stale: "stale_example",
  // string | Constrains results to jobs that start with the defined prefix (optional)
  prefix: "prefix_example",
  // string | A Nomad ACL token. (optional)
  xNomadToken: "X-Nomad-Token_example",
  // number | Maximum number of results to return. (optional)
  perPage: 1,
  // string | Indicates where to start paging for queries that support pagination. (optional)
  nextToken: "next_token_example",
};

apiInstance.getNode(body).then((data:any) => {
  console.log('API called successfully. Returned data: ' + data);
}).catch((error:any) => console.error(error));
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nodeId** | [**string**] | The ID of the node. | defaults to undefined
 **region** | [**string**] | Filters results based on the specified region. | (optional) defaults to undefined
 **namespace** | [**string**] | Filters results based on the specified namespace. | (optional) defaults to undefined
 **index** | [**number**] | If set, wait until query exceeds given index. Must be provided with WaitParam. | (optional) defaults to undefined
 **wait** | [**string**] | Provided with IndexParam to wait for change. | (optional) defaults to undefined
 **stale** | [**string**] | If present, results will include stale reads. | (optional) defaults to undefined
 **prefix** | [**string**] | Constrains results to jobs that start with the defined prefix | (optional) defaults to undefined
 **xNomadToken** | [**string**] | A Nomad ACL token. | (optional) defaults to undefined
 **perPage** | [**number**] | Maximum number of results to return. | (optional) defaults to undefined
 **nextToken** | [**string**] | Indicates where to start paging for queries that support pagination. | (optional) defaults to undefined


### Return type

**Node**

### Authorization

[X-Nomad-Token](README.md#X-Nomad-Token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** |  |  * X-Nomad-Index - A unique identifier representing the current state of the requested resource. On a new Nomad cluster the value of this index starts at 1. <br>  * X-Nomad-KnownLeader - Boolean indicating if there is a known cluster leader. <br>  * X-Nomad-LastContact - The time in milliseconds that a server was last contacted by the leader node. <br>  |
**400** | Bad request |  -  |
**403** | Forbidden |  -  |
**405** | Method not allowed |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](README.md#documentation-for-api-endpoints) [[Back to Model list]](README.md#documentation-for-models) [[Back to README]](README.md)

# **getNodeAllocations**
> Array<AllocationListStub> getNodeAllocations()


### Example


```typescript
import { nomad-client } from '';
import * as fs from 'fs';

const configuration = nomad-client.createConfiguration();
const apiInstance = new nomad-client.NodesApi(configuration);

let body:nomad-client.NodesApiGetNodeAllocationsRequest = {
  // string | The ID of the node.
  nodeId: "nodeId_example",
  // string | Filters results based on the specified region. (optional)
  region: "region_example",
  // string | Filters results based on the specified namespace. (optional)
  namespace: "namespace_example",
  // number | If set, wait until query exceeds given index. Must be provided with WaitParam. (optional)
  index: 1,
  // string | Provided with IndexParam to wait for change. (optional)
  wait: "wait_example",
  // string | If present, results will include stale reads. (optional)
  stale: "stale_example",
  // string | Constrains results to jobs that start with the defined prefix (optional)
  prefix: "prefix_example",
  // string | A Nomad ACL token. (optional)
  xNomadToken: "X-Nomad-Token_example",
  // number | Maximum number of results to return. (optional)
  perPage: 1,
  // string | Indicates where to start paging for queries that support pagination. (optional)
  nextToken: "next_token_example",
};

apiInstance.getNodeAllocations(body).then((data:any) => {
  console.log('API called successfully. Returned data: ' + data);
}).catch((error:any) => console.error(error));
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nodeId** | [**string**] | The ID of the node. | defaults to undefined
 **region** | [**string**] | Filters results based on the specified region. | (optional) defaults to undefined
 **namespace** | [**string**] | Filters results based on the specified namespace. | (optional) defaults to undefined
 **index** | [**number**] | If set, wait until query exceeds given index. Must be provided with WaitParam. | (optional) defaults to undefined
 **wait** | [**string**] | Provided with IndexParam to wait for change. | (optional) defaults to undefined
 **stale** | [**string**] | If present, results will include stale reads. | (optional) defaults to undefined
 **prefix** | [**string**] | Constrains results to jobs that start with the defined prefix | (optional) defaults to undefined
 **xNomadToken** | [**string**] | A Nomad ACL token. | (optional) defaults to undefined
 **perPage** | [**number**] | Maximum number of results to return. | (optional) defaults to undefined
 **nextToken** | [**string**] | Indicates where to start paging for queries that support pagination. | (optional) defaults to undefined


### Return type

**Array<AllocationListStub>**

### Authorization

[X-Nomad-Token](README.md#X-Nomad-Token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** |  |  * X-Nomad-Index - A unique identifier representing the current state of the requested resource. On a new Nomad cluster the value of this index starts at 1. <br>  * X-Nomad-KnownLeader - Boolean indicating if there is a known cluster leader. <br>  * X-Nomad-LastContact - The time in milliseconds that a server was last contacted by the leader node. <br>  |
**400** | Bad request |  -  |
**403** | Forbidden |  -  |
**405** | Method not allowed |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](README.md#documentation-for-api-endpoints) [[Back to Model list]](README.md#documentation-for-models) [[Back to README]](README.md)

# **getNodes**
> Array<NodeListStub> getNodes()


### Example


```typescript
import { nomad-client } from '';
import * as fs from 'fs';

const configuration = nomad-client.createConfiguration();
const apiInstance = new nomad-client.NodesApi(configuration);

let body:nomad-client.NodesApiGetNodesRequest = {
  // string | Filters results based on the specified region. (optional)
  region: "region_example",
  // string | Filters results based on the specified namespace. (optional)
  namespace: "namespace_example",
  // number | If set, wait until query exceeds given index. Must be provided with WaitParam. (optional)
  index: 1,
  // string | Provided with IndexParam to wait for change. (optional)
  wait: "wait_example",
  // string | If present, results will include stale reads. (optional)
  stale: "stale_example",
  // string | Constrains results to jobs that start with the defined prefix (optional)
  prefix: "prefix_example",
  // string | A Nomad ACL token. (optional)
  xNomadToken: "X-Nomad-Token_example",
  // number | Maximum number of results to return. (optional)
  perPage: 1,
  // string | Indicates where to start paging for queries that support pagination. (optional)
  nextToken: "next_token_example",
  // boolean | Whether or not to include the NodeResources and ReservedResources fields in the response. (optional)
  resources: true,
};

apiInstance.getNodes(body).then((data:any) => {
  console.log('API called successfully. Returned data: ' + data);
}).catch((error:any) => console.error(error));
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **region** | [**string**] | Filters results based on the specified region. | (optional) defaults to undefined
 **namespace** | [**string**] | Filters results based on the specified namespace. | (optional) defaults to undefined
 **index** | [**number**] | If set, wait until query exceeds given index. Must be provided with WaitParam. | (optional) defaults to undefined
 **wait** | [**string**] | Provided with IndexParam to wait for change. | (optional) defaults to undefined
 **stale** | [**string**] | If present, results will include stale reads. | (optional) defaults to undefined
 **prefix** | [**string**] | Constrains results to jobs that start with the defined prefix | (optional) defaults to undefined
 **xNomadToken** | [**string**] | A Nomad ACL token. | (optional) defaults to undefined
 **perPage** | [**number**] | Maximum number of results to return. | (optional) defaults to undefined
 **nextToken** | [**string**] | Indicates where to start paging for queries that support pagination. | (optional) defaults to undefined
 **resources** | [**boolean**] | Whether or not to include the NodeResources and ReservedResources fields in the response. | (optional) defaults to undefined


### Return type

**Array<NodeListStub>**

### Authorization

[X-Nomad-Token](README.md#X-Nomad-Token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** |  |  * X-Nomad-Index - A unique identifier representing the current state of the requested resource. On a new Nomad cluster the value of this index starts at 1. <br>  * X-Nomad-KnownLeader - Boolean indicating if there is a known cluster leader. <br>  * X-Nomad-LastContact - The time in milliseconds that a server was last contacted by the leader node. <br>  |
**400** | Bad request |  -  |
**403** | Forbidden |  -  |
**405** | Method not allowed |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](README.md#documentation-for-api-endpoints) [[Back to Model list]](README.md#documentation-for-models) [[Back to README]](README.md)

# **updateNodeDrain**
> NodeDrainUpdateResponse updateNodeDrain(nodeUpdateDrainRequest)


### Example


```typescript
import { nomad-client } from '';
import * as fs from 'fs';

const configuration = nomad-client.createConfiguration();
const apiInstance = new nomad-client.NodesApi(configuration);

let body:nomad-client.NodesApiUpdateNodeDrainRequest = {
  // string | The ID of the node.
  nodeId: "nodeId_example",
  // NodeUpdateDrainRequest
  nodeUpdateDrainRequest: {
    drainSpec: {
      deadline: 1,
      ignoreSystemJobs: true,
    },
    markEligible: true,
    meta: {
      "key": "key_example",
    },
    nodeID: "nodeID_example",
  },
  // string | Filters results based on the specified region. (optional)
  region: "region_example",
  // string | Filters results based on the specified namespace. (optional)
  namespace: "namespace_example",
  // number | If set, wait until query exceeds given index. Must be provided with WaitParam. (optional)
  index: 1,
  // string | Provided with IndexParam to wait for change. (optional)
  wait: "wait_example",
  // string | If present, results will include stale reads. (optional)
  stale: "stale_example",
  // string | Constrains results to jobs that start with the defined prefix (optional)
  prefix: "prefix_example",
  // string | A Nomad ACL token. (optional)
  xNomadToken: "X-Nomad-Token_example",
  // number | Maximum number of results to return. (optional)
  perPage: 1,
  // string | Indicates where to start paging for queries that support pagination. (optional)
  nextToken: "next_token_example",
};

apiInstance.updateNodeDrain(body).then((data:any) => {
  console.log('API called successfully. Returned data: ' + data);
}).catch((error:any) => console.error(error));
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nodeUpdateDrainRequest** | **NodeUpdateDrainRequest**|  |
 **nodeId** | [**string**] | The ID of the node. | defaults to undefined
 **region** | [**string**] | Filters results based on the specified region. | (optional) defaults to undefined
 **namespace** | [**string**] | Filters results based on the specified namespace. | (optional) defaults to undefined
 **index** | [**number**] | If set, wait until query exceeds given index. Must be provided with WaitParam. | (optional) defaults to undefined
 **wait** | [**string**] | Provided with IndexParam to wait for change. | (optional) defaults to undefined
 **stale** | [**string**] | If present, results will include stale reads. | (optional) defaults to undefined
 **prefix** | [**string**] | Constrains results to jobs that start with the defined prefix | (optional) defaults to undefined
 **xNomadToken** | [**string**] | A Nomad ACL token. | (optional) defaults to undefined
 **perPage** | [**number**] | Maximum number of results to return. | (optional) defaults to undefined
 **nextToken** | [**string**] | Indicates where to start paging for queries that support pagination. | (optional) defaults to undefined


### Return type

**NodeDrainUpdateResponse**

### Authorization

[X-Nomad-Token](README.md#X-Nomad-Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** |  |  * X-Nomad-Index - A unique identifier representing the current state of the requested resource. On a new Nomad cluster the value of this index starts at 1. <br>  |
**400** | Bad request |  -  |
**403** | Forbidden |  -  |
**405** | Method not allowed |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](README.md#documentation-for-api-endpoints) [[Back to Model list]](README.md#documentation-for-models) [[Back to README]](README.md)

# **updateNodeEligibility**
> NodeEligibilityUpdateResponse updateNodeEligibility(nodeUpdateEligibilityRequest)


### Example


```typescript
import { nomad-client } from '';
import * as fs from 'fs';

const configuration = nomad-client.createConfiguration();
const apiInstance = new nomad-client.NodesApi(configuration);

let body:nomad-client.NodesApiUpdateNodeEligibilityRequest = {
  // string | The ID of the node.
  nodeId: "nodeId_example",
  // NodeUpdateEligibilityRequest
  nodeUpdateEligibilityRequest: {
    eligibility: "eligibility_example",
    nodeID: "nodeID_example",
  },
  // string | Filters results based on the specified region. (optional)
  region: "region_example",
  // string | Filters results based on the specified namespace. (optional)
  namespace: "namespace_example",
  // number | If set, wait until query exceeds given index. Must be provided with WaitParam. (optional)
  index: 1,
  // string | Provided with IndexParam to wait for change. (optional)
  wait: "wait_example",
  // string | If present, results will include stale reads. (optional)
  stale: "stale_example",
  // string | Constrains results to jobs that start with the defined prefix (optional)
  prefix: "prefix_example",
  // string | A Nomad ACL token. (optional)
  xNomadToken: "X-Nomad-Token_example",
  // number | Maximum number of results to return. (optional)
  perPage: 1,
  // string | Indicates where to start paging for queries that support pagination. (optional)
  nextToken: "next_token_example",
};

apiInstance.updateNodeEligibility(body).then((data:any) => {
  console.log('API called successfully. Returned data: ' + data);
}).catch((error:any) => console.error(error));
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nodeUpdateEligibilityRequest** | **NodeUpdateEligibilityRequest**|  |
 **nodeId** | [**string**] | The ID of the node. | defaults to undefined
 **region** | [**string**] | Filters results based on the specified region. | (optional) defaults to undefined
 **namespace** | [**string**] | Filters results based on the specified namespace. | (optional) defaults to undefined
 **index** | [**number**] | If set, wait until query exceeds given index. Must be provided with WaitParam. | (optional) defaults to undefined
 **wait** | [**string**] | Provided with IndexParam to wait for change. | (optional) defaults to undefined
 **stale** | [**string**] | If present, results will include stale reads. | (optional) defaults to undefined
 **prefix** | [**string**] | Constrains results to jobs that start with the defined prefix | (optional) defaults to undefined
 **xNomadToken** | [**string**] | A Nomad ACL token. | (optional) defaults to undefined
 **perPage** | [**number**] | Maximum number of results to return. | (optional) defaults to undefined
 **nextToken** | [**string**] | Indicates where to start paging for queries that support pagination. | (optional) defaults to undefined


### Return type

**NodeEligibilityUpdateResponse**

### Authorization

[X-Nomad-Token](README.md#X-Nomad-Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** |  |  * X-Nomad-Index - A unique identifier representing the current state of the requested resource. On a new Nomad cluster the value of this index starts at 1. <br>  |
**400** | Bad request |  -  |
**403** | Forbidden |  -  |
**405** | Method not allowed |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](README.md#documentation-for-api-endpoints) [[Back to Model list]](README.md#documentation-for-models) [[Back to README]](README.md)

# **updateNodePurge**
> NodePurgeResponse updateNodePurge()


### Example


```typescript
import { nomad-client } from '';
import * as fs from 'fs';

const configuration = nomad-client.createConfiguration();
const apiInstance = new nomad-client.NodesApi(configuration);

let body:nomad-client.NodesApiUpdateNodePurgeRequest = {
  // string | The ID of the node.
  nodeId: "nodeId_example",
  // string | Filters results based on the specified region. (optional)
  region: "region_example",
  // string | Filters results based on the specified namespace. (optional)
  namespace: "namespace_example",
  // number | If set, wait until query exceeds given index. Must be provided with WaitParam. (optional)
  index: 1,
  // string | Provided with IndexParam to wait for change. (optional)
  wait: "wait_example",
  // string | If present, results will include stale reads. (optional)
  stale: "stale_example",
  // string | Constrains results to jobs that start with the defined prefix (optional)
  prefix: "prefix_example",
  // string | A Nomad ACL token. (optional)
  xNomadToken: "X-Nomad-Token_example",
  // number | Maximum number of results to return. (optional)
  perPage: 1,
  // string | Indicates where to start paging for queries that support pagination. (optional)
  nextToken: "next_token_example",
};

apiInstance.updateNodePurge(body).then((data:any) => {
  console.log('API called successfully. Returned data: ' + data);
}).catch((error:any) => console.error(error));
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nodeId** | [**string**] | The ID of the node. | defaults to undefined
 **region** | [**string**] | Filters results based on the specified region. | (optional) defaults to undefined
 **namespace** | [**string**] | Filters results based on the specified namespace. | (optional) defaults to undefined
 **index** | [**number**] | If set, wait until query exceeds given index. Must be provided with WaitParam. | (optional) defaults to undefined
 **wait** | [**string**] | Provided with IndexParam to wait for change. | (optional) defaults to undefined
 **stale** | [**string**] | If present, results will include stale reads. | (optional) defaults to undefined
 **prefix** | [**string**] | Constrains results to jobs that start with the defined prefix | (optional) defaults to undefined
 **xNomadToken** | [**string**] | A Nomad ACL token. | (optional) defaults to undefined
 **perPage** | [**number**] | Maximum number of results to return. | (optional) defaults to undefined
 **nextToken** | [**string**] | Indicates where to start paging for queries that support pagination. | (optional) defaults to undefined


### Return type

**NodePurgeResponse**

### Authorization

[X-Nomad-Token](README.md#X-Nomad-Token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** |  |  * X-Nomad-Index - A unique identifier representing the current state of the requested resource. On a new Nomad cluster the value of this index starts at 1. <br>  |
**400** | Bad request |  -  |
**403** | Forbidden |  -  |
**405** | Method not allowed |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](README.md#documentation-for-api-endpoints) [[Back to Model list]](README.md#documentation-for-models) [[Back to README]](README.md)


