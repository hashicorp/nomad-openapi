# nomad-client.EnterpriseApi

All URIs are relative to *https://127.0.0.1:4646/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createQuotaSpec**](EnterpriseApi.md#createQuotaSpec) | **POST** /quota | 
[**deleteQuotaSpec**](EnterpriseApi.md#deleteQuotaSpec) | **DELETE** /quota/{specName} | 
[**getQuotaSpec**](EnterpriseApi.md#getQuotaSpec) | **GET** /quota/{specName} | 
[**getQuotas**](EnterpriseApi.md#getQuotas) | **GET** /quotas | 
[**postQuotaSpec**](EnterpriseApi.md#postQuotaSpec) | **POST** /quota/{specName} | 


# **createQuotaSpec**
> void createQuotaSpec(quotaSpec)


### Example


```typescript
import { nomad-client } from '';
import * as fs from 'fs';

const configuration = nomad-client.createConfiguration();
const apiInstance = new nomad-client.EnterpriseApi(configuration);

let body:nomad-client.EnterpriseApiCreateQuotaSpecRequest = {
  // QuotaSpec
  quotaSpec: {
    createIndex: 0,
    description: "description_example",
    limits: [
      {
        hash: 'YQ==',
        region: "region_example",
        regionLimit: {
          CPU: 1,
          cores: 1,
          devices: [
            {
              affinities: [
                {
                  lTarget: "lTarget_example",
                  operand: "operand_example",
                  rTarget: "rTarget_example",
                  weight: -128,
                },
              ],
              constraints: [
                {
                  lTarget: "lTarget_example",
                  operand: "operand_example",
                  rTarget: "rTarget_example",
                },
              ],
              count: 0,
              name: "name_example",
            },
          ],
          diskMB: 1,
          IOPS: 1,
          memoryMB: 1,
          memoryMaxMB: 1,
          networks: [
            {
              CIDR: "CIDR_example",
              DNS: {
                options: [
                  "options_example",
                ],
                searches: [
                  "searches_example",
                ],
                servers: [
                  "servers_example",
                ],
              },
              device: "device_example",
              dynamicPorts: [
                {
                  hostNetwork: "hostNetwork_example",
                  label: "label_example",
                  to: 1,
                  value: 1,
                },
              ],
              hostname: "hostname_example",
              IP: "IP_example",
              mBits: 1,
              mode: "mode_example",
              reservedPorts: [
                {
                  hostNetwork: "hostNetwork_example",
                  label: "label_example",
                  to: 1,
                  value: 1,
                },
              ],
            },
          ],
        },
      },
    ],
    modifyIndex: 0,
    name: "name_example",
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

apiInstance.createQuotaSpec(body).then((data:any) => {
  console.log('API called successfully. Returned data: ' + data);
}).catch((error:any) => console.error(error));
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **quotaSpec** | **QuotaSpec**|  |
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

# **deleteQuotaSpec**
> void deleteQuotaSpec()


### Example


```typescript
import { nomad-client } from '';
import * as fs from 'fs';

const configuration = nomad-client.createConfiguration();
const apiInstance = new nomad-client.EnterpriseApi(configuration);

let body:nomad-client.EnterpriseApiDeleteQuotaSpecRequest = {
  // string | The quota spec identifier.
  specName: "specName_example",
  // string | Filters results based on the specified region. (optional)
  region: "region_example",
  // string | Filters results based on the specified namespace. (optional)
  namespace: "namespace_example",
  // string | A Nomad ACL token. (optional)
  xNomadToken: "X-Nomad-Token_example",
  // string | Can be used to ensure operations are only run once. (optional)
  idempotencyToken: "idempotency_token_example",
};

apiInstance.deleteQuotaSpec(body).then((data:any) => {
  console.log('API called successfully. Returned data: ' + data);
}).catch((error:any) => console.error(error));
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **specName** | [**string**] | The quota spec identifier. | defaults to undefined
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

# **getQuotaSpec**
> QuotaSpec getQuotaSpec()


### Example


```typescript
import { nomad-client } from '';
import * as fs from 'fs';

const configuration = nomad-client.createConfiguration();
const apiInstance = new nomad-client.EnterpriseApi(configuration);

let body:nomad-client.EnterpriseApiGetQuotaSpecRequest = {
  // string | The quota spec identifier.
  specName: "specName_example",
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

apiInstance.getQuotaSpec(body).then((data:any) => {
  console.log('API called successfully. Returned data: ' + data);
}).catch((error:any) => console.error(error));
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **specName** | [**string**] | The quota spec identifier. | defaults to undefined
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

**QuotaSpec**

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

# **getQuotas**
> Array<any> getQuotas()


### Example


```typescript
import { nomad-client } from '';
import * as fs from 'fs';

const configuration = nomad-client.createConfiguration();
const apiInstance = new nomad-client.EnterpriseApi(configuration);

let body:nomad-client.EnterpriseApiGetQuotasRequest = {
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

apiInstance.getQuotas(body).then((data:any) => {
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

**Array<any>**

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

# **postQuotaSpec**
> void postQuotaSpec(quotaSpec)


### Example


```typescript
import { nomad-client } from '';
import * as fs from 'fs';

const configuration = nomad-client.createConfiguration();
const apiInstance = new nomad-client.EnterpriseApi(configuration);

let body:nomad-client.EnterpriseApiPostQuotaSpecRequest = {
  // string | The quota spec identifier.
  specName: "specName_example",
  // QuotaSpec
  quotaSpec: {
    createIndex: 0,
    description: "description_example",
    limits: [
      {
        hash: 'YQ==',
        region: "region_example",
        regionLimit: {
          CPU: 1,
          cores: 1,
          devices: [
            {
              affinities: [
                {
                  lTarget: "lTarget_example",
                  operand: "operand_example",
                  rTarget: "rTarget_example",
                  weight: -128,
                },
              ],
              constraints: [
                {
                  lTarget: "lTarget_example",
                  operand: "operand_example",
                  rTarget: "rTarget_example",
                },
              ],
              count: 0,
              name: "name_example",
            },
          ],
          diskMB: 1,
          IOPS: 1,
          memoryMB: 1,
          memoryMaxMB: 1,
          networks: [
            {
              CIDR: "CIDR_example",
              DNS: {
                options: [
                  "options_example",
                ],
                searches: [
                  "searches_example",
                ],
                servers: [
                  "servers_example",
                ],
              },
              device: "device_example",
              dynamicPorts: [
                {
                  hostNetwork: "hostNetwork_example",
                  label: "label_example",
                  to: 1,
                  value: 1,
                },
              ],
              hostname: "hostname_example",
              IP: "IP_example",
              mBits: 1,
              mode: "mode_example",
              reservedPorts: [
                {
                  hostNetwork: "hostNetwork_example",
                  label: "label_example",
                  to: 1,
                  value: 1,
                },
              ],
            },
          ],
        },
      },
    ],
    modifyIndex: 0,
    name: "name_example",
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

apiInstance.postQuotaSpec(body).then((data:any) => {
  console.log('API called successfully. Returned data: ' + data);
}).catch((error:any) => console.error(error));
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **quotaSpec** | **QuotaSpec**|  |
 **specName** | [**string**] | The quota spec identifier. | defaults to undefined
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


