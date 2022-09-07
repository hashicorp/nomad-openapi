# nomad-client.ACLApi

All URIs are relative to *https://127.0.0.1:4646/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**deleteACLPolicy**](ACLApi.md#deleteACLPolicy) | **DELETE** /acl/policy/{policyName} | 
[**deleteACLToken**](ACLApi.md#deleteACLToken) | **DELETE** /acl/token/{tokenAccessor} | 
[**getACLPolicies**](ACLApi.md#getACLPolicies) | **GET** /acl/policies | 
[**getACLPolicy**](ACLApi.md#getACLPolicy) | **GET** /acl/policy/{policyName} | 
[**getACLToken**](ACLApi.md#getACLToken) | **GET** /acl/token/{tokenAccessor} | 
[**getACLTokenSelf**](ACLApi.md#getACLTokenSelf) | **GET** /acl/token | 
[**getACLTokens**](ACLApi.md#getACLTokens) | **GET** /acl/tokens | 
[**postACLBootstrap**](ACLApi.md#postACLBootstrap) | **POST** /acl/bootstrap | 
[**postACLPolicy**](ACLApi.md#postACLPolicy) | **POST** /acl/policy/{policyName} | 
[**postACLToken**](ACLApi.md#postACLToken) | **POST** /acl/token/{tokenAccessor} | 
[**postACLTokenOnetime**](ACLApi.md#postACLTokenOnetime) | **POST** /acl/token/onetime | 
[**postACLTokenOnetimeExchange**](ACLApi.md#postACLTokenOnetimeExchange) | **POST** /acl/token/onetime/exchange | 


# **deleteACLPolicy**
> void deleteACLPolicy()


### Example


```typescript
import { nomad-client } from '';
import * as fs from 'fs';

const configuration = nomad-client.createConfiguration();
const apiInstance = new nomad-client.ACLApi(configuration);

let body:nomad-client.ACLApiDeleteACLPolicyRequest = {
  // string | The ACL policy name.
  policyName: "policyName_example",
  // string | Filters results based on the specified region. (optional)
  region: "region_example",
  // string | Filters results based on the specified namespace. (optional)
  namespace: "namespace_example",
  // string | A Nomad ACL token. (optional)
  xNomadToken: "X-Nomad-Token_example",
  // string | Can be used to ensure operations are only run once. (optional)
  idempotencyToken: "idempotency_token_example",
};

apiInstance.deleteACLPolicy(body).then((data:any) => {
  console.log('API called successfully. Returned data: ' + data);
}).catch((error:any) => console.error(error));
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **policyName** | [**string**] | The ACL policy name. | defaults to undefined
 **region** | [**string**] | Filters results based on the specified region. | (optional) defaults to undefined
 **namespace** | [**string**] | Filters results based on the specified namespace. | (optional) defaults to undefined
 **xNomadToken** | [**string**] | A Nomad ACL token. | (optional) defaults to undefined
 **idempotencyToken** | [**string**] | Can be used to ensure operations are only run once. | (optional) defaults to undefined


### Return type

**void**

### Authorization

[X-Nomad-Token](README.md#X-Nomad-Token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** |  |  * X-Nomad-Index - A unique identifier representing the current state of the requested resource. On a new Nomad cluster the value of this index starts at 1. <br>  |
**400** | Bad request |  -  |
**403** | Forbidden |  -  |
**405** | Method not allowed |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](README.md#documentation-for-api-endpoints) [[Back to Model list]](README.md#documentation-for-models) [[Back to README]](README.md)

# **deleteACLToken**
> void deleteACLToken()


### Example


```typescript
import { nomad-client } from '';
import * as fs from 'fs';

const configuration = nomad-client.createConfiguration();
const apiInstance = new nomad-client.ACLApi(configuration);

let body:nomad-client.ACLApiDeleteACLTokenRequest = {
  // string | The token accessor ID.
  tokenAccessor: "tokenAccessor_example",
  // string | Filters results based on the specified region. (optional)
  region: "region_example",
  // string | Filters results based on the specified namespace. (optional)
  namespace: "namespace_example",
  // string | A Nomad ACL token. (optional)
  xNomadToken: "X-Nomad-Token_example",
  // string | Can be used to ensure operations are only run once. (optional)
  idempotencyToken: "idempotency_token_example",
};

apiInstance.deleteACLToken(body).then((data:any) => {
  console.log('API called successfully. Returned data: ' + data);
}).catch((error:any) => console.error(error));
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenAccessor** | [**string**] | The token accessor ID. | defaults to undefined
 **region** | [**string**] | Filters results based on the specified region. | (optional) defaults to undefined
 **namespace** | [**string**] | Filters results based on the specified namespace. | (optional) defaults to undefined
 **xNomadToken** | [**string**] | A Nomad ACL token. | (optional) defaults to undefined
 **idempotencyToken** | [**string**] | Can be used to ensure operations are only run once. | (optional) defaults to undefined


### Return type

**void**

### Authorization

[X-Nomad-Token](README.md#X-Nomad-Token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** |  |  * X-Nomad-Index - A unique identifier representing the current state of the requested resource. On a new Nomad cluster the value of this index starts at 1. <br>  |
**400** | Bad request |  -  |
**403** | Forbidden |  -  |
**405** | Method not allowed |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](README.md#documentation-for-api-endpoints) [[Back to Model list]](README.md#documentation-for-models) [[Back to README]](README.md)

# **getACLPolicies**
> Array<ACLPolicyListStub> getACLPolicies()


### Example


```typescript
import { nomad-client } from '';
import * as fs from 'fs';

const configuration = nomad-client.createConfiguration();
const apiInstance = new nomad-client.ACLApi(configuration);

let body:nomad-client.ACLApiGetACLPoliciesRequest = {
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

apiInstance.getACLPolicies(body).then((data:any) => {
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

**Array<ACLPolicyListStub>**

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

# **getACLPolicy**
> ACLPolicy getACLPolicy()


### Example


```typescript
import { nomad-client } from '';
import * as fs from 'fs';

const configuration = nomad-client.createConfiguration();
const apiInstance = new nomad-client.ACLApi(configuration);

let body:nomad-client.ACLApiGetACLPolicyRequest = {
  // string | The ACL policy name.
  policyName: "policyName_example",
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

apiInstance.getACLPolicy(body).then((data:any) => {
  console.log('API called successfully. Returned data: ' + data);
}).catch((error:any) => console.error(error));
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **policyName** | [**string**] | The ACL policy name. | defaults to undefined
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

**ACLPolicy**

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

# **getACLToken**
> ACLToken getACLToken()


### Example


```typescript
import { nomad-client } from '';
import * as fs from 'fs';

const configuration = nomad-client.createConfiguration();
const apiInstance = new nomad-client.ACLApi(configuration);

let body:nomad-client.ACLApiGetACLTokenRequest = {
  // string | The token accessor ID.
  tokenAccessor: "tokenAccessor_example",
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

apiInstance.getACLToken(body).then((data:any) => {
  console.log('API called successfully. Returned data: ' + data);
}).catch((error:any) => console.error(error));
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenAccessor** | [**string**] | The token accessor ID. | defaults to undefined
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

**ACLToken**

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

# **getACLTokenSelf**
> ACLToken getACLTokenSelf()


### Example


```typescript
import { nomad-client } from '';
import * as fs from 'fs';

const configuration = nomad-client.createConfiguration();
const apiInstance = new nomad-client.ACLApi(configuration);

let body:nomad-client.ACLApiGetACLTokenSelfRequest = {
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

apiInstance.getACLTokenSelf(body).then((data:any) => {
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

**ACLToken**

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

# **getACLTokens**
> Array<ACLTokenListStub> getACLTokens()


### Example


```typescript
import { nomad-client } from '';
import * as fs from 'fs';

const configuration = nomad-client.createConfiguration();
const apiInstance = new nomad-client.ACLApi(configuration);

let body:nomad-client.ACLApiGetACLTokensRequest = {
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

apiInstance.getACLTokens(body).then((data:any) => {
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

**Array<ACLTokenListStub>**

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

# **postACLBootstrap**
> ACLToken postACLBootstrap()


### Example


```typescript
import { nomad-client } from '';
import * as fs from 'fs';

const configuration = nomad-client.createConfiguration();
const apiInstance = new nomad-client.ACLApi(configuration);

let body:nomad-client.ACLApiPostACLBootstrapRequest = {
  // string | Filters results based on the specified region. (optional)
  region: "region_example",
  // string | Filters results based on the specified namespace. (optional)
  namespace: "namespace_example",
  // string | A Nomad ACL token. (optional)
  xNomadToken: "X-Nomad-Token_example",
  // string | Can be used to ensure operations are only run once. (optional)
  idempotencyToken: "idempotency_token_example",
};

apiInstance.postACLBootstrap(body).then((data:any) => {
  console.log('API called successfully. Returned data: ' + data);
}).catch((error:any) => console.error(error));
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **region** | [**string**] | Filters results based on the specified region. | (optional) defaults to undefined
 **namespace** | [**string**] | Filters results based on the specified namespace. | (optional) defaults to undefined
 **xNomadToken** | [**string**] | A Nomad ACL token. | (optional) defaults to undefined
 **idempotencyToken** | [**string**] | Can be used to ensure operations are only run once. | (optional) defaults to undefined


### Return type

**ACLToken**

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

# **postACLPolicy**
> void postACLPolicy(aCLPolicy)


### Example


```typescript
import { nomad-client } from '';
import * as fs from 'fs';

const configuration = nomad-client.createConfiguration();
const apiInstance = new nomad-client.ACLApi(configuration);

let body:nomad-client.ACLApiPostACLPolicyRequest = {
  // string | The ACL policy name.
  policyName: "policyName_example",
  // ACLPolicy
  aCLPolicy: {
    createIndex: 0,
    description: "description_example",
    modifyIndex: 0,
    name: "name_example",
    rules: "rules_example",
  },
  // string | Filters results based on the specified region. (optional)
  region: "region_example",
  // string | Filters results based on the specified namespace. (optional)
  namespace: "namespace_example",
  // string | A Nomad ACL token. (optional)
  xNomadToken: "X-Nomad-Token_example",
  // string | Can be used to ensure operations are only run once. (optional)
  idempotencyToken: "idempotency_token_example",
};

apiInstance.postACLPolicy(body).then((data:any) => {
  console.log('API called successfully. Returned data: ' + data);
}).catch((error:any) => console.error(error));
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aCLPolicy** | **ACLPolicy**|  |
 **policyName** | [**string**] | The ACL policy name. | defaults to undefined
 **region** | [**string**] | Filters results based on the specified region. | (optional) defaults to undefined
 **namespace** | [**string**] | Filters results based on the specified namespace. | (optional) defaults to undefined
 **xNomadToken** | [**string**] | A Nomad ACL token. | (optional) defaults to undefined
 **idempotencyToken** | [**string**] | Can be used to ensure operations are only run once. | (optional) defaults to undefined


### Return type

**void**

### Authorization

[X-Nomad-Token](README.md#X-Nomad-Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** |  |  * X-Nomad-Index - A unique identifier representing the current state of the requested resource. On a new Nomad cluster the value of this index starts at 1. <br>  |
**400** | Bad request |  -  |
**403** | Forbidden |  -  |
**405** | Method not allowed |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](README.md#documentation-for-api-endpoints) [[Back to Model list]](README.md#documentation-for-models) [[Back to README]](README.md)

# **postACLToken**
> ACLToken postACLToken(aCLToken)


### Example


```typescript
import { nomad-client } from '';
import * as fs from 'fs';

const configuration = nomad-client.createConfiguration();
const apiInstance = new nomad-client.ACLApi(configuration);

let body:nomad-client.ACLApiPostACLTokenRequest = {
  // string | The token accessor ID.
  tokenAccessor: "tokenAccessor_example",
  // ACLToken
  aCLToken: {
    accessorID: "accessorID_example",
    createIndex: 0,
    createTime: new Date('1970-01-01T00:00:00.00Z'),
    global: true,
    modifyIndex: 0,
    name: "name_example",
    policies: [
      "policies_example",
    ],
    secretID: "secretID_example",
    type: "type_example",
  },
  // string | Filters results based on the specified region. (optional)
  region: "region_example",
  // string | Filters results based on the specified namespace. (optional)
  namespace: "namespace_example",
  // string | A Nomad ACL token. (optional)
  xNomadToken: "X-Nomad-Token_example",
  // string | Can be used to ensure operations are only run once. (optional)
  idempotencyToken: "idempotency_token_example",
};

apiInstance.postACLToken(body).then((data:any) => {
  console.log('API called successfully. Returned data: ' + data);
}).catch((error:any) => console.error(error));
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aCLToken** | **ACLToken**|  |
 **tokenAccessor** | [**string**] | The token accessor ID. | defaults to undefined
 **region** | [**string**] | Filters results based on the specified region. | (optional) defaults to undefined
 **namespace** | [**string**] | Filters results based on the specified namespace. | (optional) defaults to undefined
 **xNomadToken** | [**string**] | A Nomad ACL token. | (optional) defaults to undefined
 **idempotencyToken** | [**string**] | Can be used to ensure operations are only run once. | (optional) defaults to undefined


### Return type

**ACLToken**

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

# **postACLTokenOnetime**
> OneTimeToken postACLTokenOnetime()


### Example


```typescript
import { nomad-client } from '';
import * as fs from 'fs';

const configuration = nomad-client.createConfiguration();
const apiInstance = new nomad-client.ACLApi(configuration);

let body:nomad-client.ACLApiPostACLTokenOnetimeRequest = {
  // string | Filters results based on the specified region. (optional)
  region: "region_example",
  // string | Filters results based on the specified namespace. (optional)
  namespace: "namespace_example",
  // string | A Nomad ACL token. (optional)
  xNomadToken: "X-Nomad-Token_example",
  // string | Can be used to ensure operations are only run once. (optional)
  idempotencyToken: "idempotency_token_example",
};

apiInstance.postACLTokenOnetime(body).then((data:any) => {
  console.log('API called successfully. Returned data: ' + data);
}).catch((error:any) => console.error(error));
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **region** | [**string**] | Filters results based on the specified region. | (optional) defaults to undefined
 **namespace** | [**string**] | Filters results based on the specified namespace. | (optional) defaults to undefined
 **xNomadToken** | [**string**] | A Nomad ACL token. | (optional) defaults to undefined
 **idempotencyToken** | [**string**] | Can be used to ensure operations are only run once. | (optional) defaults to undefined


### Return type

**OneTimeToken**

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

# **postACLTokenOnetimeExchange**
> ACLToken postACLTokenOnetimeExchange(oneTimeTokenExchangeRequest)


### Example


```typescript
import { nomad-client } from '';
import * as fs from 'fs';

const configuration = nomad-client.createConfiguration();
const apiInstance = new nomad-client.ACLApi(configuration);

let body:nomad-client.ACLApiPostACLTokenOnetimeExchangeRequest = {
  // OneTimeTokenExchangeRequest
  oneTimeTokenExchangeRequest: {
    oneTimeSecretID: "oneTimeSecretID_example",
  },
  // string | Filters results based on the specified region. (optional)
  region: "region_example",
  // string | Filters results based on the specified namespace. (optional)
  namespace: "namespace_example",
  // string | A Nomad ACL token. (optional)
  xNomadToken: "X-Nomad-Token_example",
  // string | Can be used to ensure operations are only run once. (optional)
  idempotencyToken: "idempotency_token_example",
};

apiInstance.postACLTokenOnetimeExchange(body).then((data:any) => {
  console.log('API called successfully. Returned data: ' + data);
}).catch((error:any) => console.error(error));
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oneTimeTokenExchangeRequest** | **OneTimeTokenExchangeRequest**|  |
 **region** | [**string**] | Filters results based on the specified region. | (optional) defaults to undefined
 **namespace** | [**string**] | Filters results based on the specified namespace. | (optional) defaults to undefined
 **xNomadToken** | [**string**] | A Nomad ACL token. | (optional) defaults to undefined
 **idempotencyToken** | [**string**] | Can be used to ensure operations are only run once. | (optional) defaults to undefined


### Return type

**ACLToken**

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


