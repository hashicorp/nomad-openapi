# nomad-client.EvaluationsApi

All URIs are relative to *https://127.0.0.1:4646/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**getEvaluation**](EvaluationsApi.md#getEvaluation) | **GET** /evaluation/{evalID} | 
[**getEvaluationAllocations**](EvaluationsApi.md#getEvaluationAllocations) | **GET** /evaluation/{evalID}/allocations | 
[**getEvaluations**](EvaluationsApi.md#getEvaluations) | **GET** /evaluations | 


# **getEvaluation**
> Evaluation getEvaluation()


### Example


```typescript
import { nomad-client } from '';
import * as fs from 'fs';

const configuration = nomad-client.createConfiguration();
const apiInstance = new nomad-client.EvaluationsApi(configuration);

let body:nomad-client.EvaluationsApiGetEvaluationRequest = {
  // string | Evaluation ID.
  evalID: "evalID_example",
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

apiInstance.getEvaluation(body).then((data:any) => {
  console.log('API called successfully. Returned data: ' + data);
}).catch((error:any) => console.error(error));
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **evalID** | [**string**] | Evaluation ID. | defaults to undefined
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

**Evaluation**

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

# **getEvaluationAllocations**
> Array<AllocationListStub> getEvaluationAllocations()


### Example


```typescript
import { nomad-client } from '';
import * as fs from 'fs';

const configuration = nomad-client.createConfiguration();
const apiInstance = new nomad-client.EvaluationsApi(configuration);

let body:nomad-client.EvaluationsApiGetEvaluationAllocationsRequest = {
  // string | Evaluation ID.
  evalID: "evalID_example",
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

apiInstance.getEvaluationAllocations(body).then((data:any) => {
  console.log('API called successfully. Returned data: ' + data);
}).catch((error:any) => console.error(error));
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **evalID** | [**string**] | Evaluation ID. | defaults to undefined
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

# **getEvaluations**
> Array<Evaluation> getEvaluations()


### Example


```typescript
import { nomad-client } from '';
import * as fs from 'fs';

const configuration = nomad-client.createConfiguration();
const apiInstance = new nomad-client.EvaluationsApi(configuration);

let body:nomad-client.EvaluationsApiGetEvaluationsRequest = {
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

apiInstance.getEvaluations(body).then((data:any) => {
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


### Return type

**Array<Evaluation>**

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


