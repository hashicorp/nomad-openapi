# nomad-client.JobsApi

All URIs are relative to *https://127.0.0.1:4646/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**deleteJob**](JobsApi.md#deleteJob) | **DELETE** /job/{jobName} | 
[**getJob**](JobsApi.md#getJob) | **GET** /job/{jobName} | 
[**getJobAllocations**](JobsApi.md#getJobAllocations) | **GET** /job/{jobName}/allocations | 
[**getJobDeployment**](JobsApi.md#getJobDeployment) | **GET** /job/{jobName}/deployment | 
[**getJobDeployments**](JobsApi.md#getJobDeployments) | **GET** /job/{jobName}/deployments | 
[**getJobEvaluations**](JobsApi.md#getJobEvaluations) | **GET** /job/{jobName}/evaluations | 
[**getJobScaleStatus**](JobsApi.md#getJobScaleStatus) | **GET** /job/{jobName}/scale | 
[**getJobSummary**](JobsApi.md#getJobSummary) | **GET** /job/{jobName}/summary | 
[**getJobVersions**](JobsApi.md#getJobVersions) | **GET** /job/{jobName}/versions | 
[**getJobs**](JobsApi.md#getJobs) | **GET** /jobs | 
[**postJob**](JobsApi.md#postJob) | **POST** /job/{jobName} | 
[**postJobDispatch**](JobsApi.md#postJobDispatch) | **POST** /job/{jobName}/dispatch | 
[**postJobEvaluate**](JobsApi.md#postJobEvaluate) | **POST** /job/{jobName}/evaluate | 
[**postJobParse**](JobsApi.md#postJobParse) | **POST** /jobs/parse | 
[**postJobPeriodicForce**](JobsApi.md#postJobPeriodicForce) | **POST** /job/{jobName}/periodic/force | 
[**postJobPlan**](JobsApi.md#postJobPlan) | **POST** /job/{jobName}/plan | 
[**postJobRevert**](JobsApi.md#postJobRevert) | **POST** /job/{jobName}/revert | 
[**postJobScalingRequest**](JobsApi.md#postJobScalingRequest) | **POST** /job/{jobName}/scale | 
[**postJobStability**](JobsApi.md#postJobStability) | **POST** /job/{jobName}/stable | 
[**postJobValidateRequest**](JobsApi.md#postJobValidateRequest) | **POST** /validate/job | 
[**registerJob**](JobsApi.md#registerJob) | **POST** /jobs | 


# **deleteJob**
> JobDeregisterResponse deleteJob()


### Example


```typescript
import { nomad-client } from '';
import * as fs from 'fs';

const configuration = nomad-client.createConfiguration();
const apiInstance = new nomad-client.JobsApi(configuration);

let body:nomad-client.JobsApiDeleteJobRequest = {
  // string | The job identifier.
  jobName: "jobName_example",
  // string | Filters results based on the specified region. (optional)
  region: "region_example",
  // string | Filters results based on the specified namespace. (optional)
  namespace: "namespace_example",
  // string | A Nomad ACL token. (optional)
  xNomadToken: "X-Nomad-Token_example",
  // string | Can be used to ensure operations are only run once. (optional)
  idempotencyToken: "idempotency_token_example",
  // boolean | Boolean flag indicating whether to purge allocations of the job after deleting. (optional)
  purge: true,
  // boolean | Boolean flag indicating whether the operation should apply to all instances of the job globally. (optional)
  global: true,
};

apiInstance.deleteJob(body).then((data:any) => {
  console.log('API called successfully. Returned data: ' + data);
}).catch((error:any) => console.error(error));
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jobName** | [**string**] | The job identifier. | defaults to undefined
 **region** | [**string**] | Filters results based on the specified region. | (optional) defaults to undefined
 **namespace** | [**string**] | Filters results based on the specified namespace. | (optional) defaults to undefined
 **xNomadToken** | [**string**] | A Nomad ACL token. | (optional) defaults to undefined
 **idempotencyToken** | [**string**] | Can be used to ensure operations are only run once. | (optional) defaults to undefined
 **purge** | [**boolean**] | Boolean flag indicating whether to purge allocations of the job after deleting. | (optional) defaults to undefined
 **global** | [**boolean**] | Boolean flag indicating whether the operation should apply to all instances of the job globally. | (optional) defaults to undefined


### Return type

**JobDeregisterResponse**

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

# **getJob**
> Job getJob()


### Example


```typescript
import { nomad-client } from '';
import * as fs from 'fs';

const configuration = nomad-client.createConfiguration();
const apiInstance = new nomad-client.JobsApi(configuration);

let body:nomad-client.JobsApiGetJobRequest = {
  // string | The job identifier.
  jobName: "jobName_example",
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

apiInstance.getJob(body).then((data:any) => {
  console.log('API called successfully. Returned data: ' + data);
}).catch((error:any) => console.error(error));
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jobName** | [**string**] | The job identifier. | defaults to undefined
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

**Job**

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

# **getJobAllocations**
> Array<AllocationListStub> getJobAllocations()


### Example


```typescript
import { nomad-client } from '';
import * as fs from 'fs';

const configuration = nomad-client.createConfiguration();
const apiInstance = new nomad-client.JobsApi(configuration);

let body:nomad-client.JobsApiGetJobAllocationsRequest = {
  // string | The job identifier.
  jobName: "jobName_example",
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
  // boolean | Specifies whether the list of allocations should include allocations from a previously registered job with the same ID. This is possible if the job is deregistered and reregistered. (optional)
  all: true,
};

apiInstance.getJobAllocations(body).then((data:any) => {
  console.log('API called successfully. Returned data: ' + data);
}).catch((error:any) => console.error(error));
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jobName** | [**string**] | The job identifier. | defaults to undefined
 **region** | [**string**] | Filters results based on the specified region. | (optional) defaults to undefined
 **namespace** | [**string**] | Filters results based on the specified namespace. | (optional) defaults to undefined
 **index** | [**number**] | If set, wait until query exceeds given index. Must be provided with WaitParam. | (optional) defaults to undefined
 **wait** | [**string**] | Provided with IndexParam to wait for change. | (optional) defaults to undefined
 **stale** | [**string**] | If present, results will include stale reads. | (optional) defaults to undefined
 **prefix** | [**string**] | Constrains results to jobs that start with the defined prefix | (optional) defaults to undefined
 **xNomadToken** | [**string**] | A Nomad ACL token. | (optional) defaults to undefined
 **perPage** | [**number**] | Maximum number of results to return. | (optional) defaults to undefined
 **nextToken** | [**string**] | Indicates where to start paging for queries that support pagination. | (optional) defaults to undefined
 **all** | [**boolean**] | Specifies whether the list of allocations should include allocations from a previously registered job with the same ID. This is possible if the job is deregistered and reregistered. | (optional) defaults to undefined


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

# **getJobDeployment**
> Deployment getJobDeployment()


### Example


```typescript
import { nomad-client } from '';
import * as fs from 'fs';

const configuration = nomad-client.createConfiguration();
const apiInstance = new nomad-client.JobsApi(configuration);

let body:nomad-client.JobsApiGetJobDeploymentRequest = {
  // string | The job identifier.
  jobName: "jobName_example",
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

apiInstance.getJobDeployment(body).then((data:any) => {
  console.log('API called successfully. Returned data: ' + data);
}).catch((error:any) => console.error(error));
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jobName** | [**string**] | The job identifier. | defaults to undefined
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

**Deployment**

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

# **getJobDeployments**
> Array<Deployment> getJobDeployments()


### Example


```typescript
import { nomad-client } from '';
import * as fs from 'fs';

const configuration = nomad-client.createConfiguration();
const apiInstance = new nomad-client.JobsApi(configuration);

let body:nomad-client.JobsApiGetJobDeploymentsRequest = {
  // string | The job identifier.
  jobName: "jobName_example",
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
  // number | Flag indicating whether to constrain by job creation index or not. (optional)
  all: 1,
};

apiInstance.getJobDeployments(body).then((data:any) => {
  console.log('API called successfully. Returned data: ' + data);
}).catch((error:any) => console.error(error));
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jobName** | [**string**] | The job identifier. | defaults to undefined
 **region** | [**string**] | Filters results based on the specified region. | (optional) defaults to undefined
 **namespace** | [**string**] | Filters results based on the specified namespace. | (optional) defaults to undefined
 **index** | [**number**] | If set, wait until query exceeds given index. Must be provided with WaitParam. | (optional) defaults to undefined
 **wait** | [**string**] | Provided with IndexParam to wait for change. | (optional) defaults to undefined
 **stale** | [**string**] | If present, results will include stale reads. | (optional) defaults to undefined
 **prefix** | [**string**] | Constrains results to jobs that start with the defined prefix | (optional) defaults to undefined
 **xNomadToken** | [**string**] | A Nomad ACL token. | (optional) defaults to undefined
 **perPage** | [**number**] | Maximum number of results to return. | (optional) defaults to undefined
 **nextToken** | [**string**] | Indicates where to start paging for queries that support pagination. | (optional) defaults to undefined
 **all** | [**number**] | Flag indicating whether to constrain by job creation index or not. | (optional) defaults to undefined


### Return type

**Array<Deployment>**

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

# **getJobEvaluations**
> Array<Evaluation> getJobEvaluations()


### Example


```typescript
import { nomad-client } from '';
import * as fs from 'fs';

const configuration = nomad-client.createConfiguration();
const apiInstance = new nomad-client.JobsApi(configuration);

let body:nomad-client.JobsApiGetJobEvaluationsRequest = {
  // string | The job identifier.
  jobName: "jobName_example",
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

apiInstance.getJobEvaluations(body).then((data:any) => {
  console.log('API called successfully. Returned data: ' + data);
}).catch((error:any) => console.error(error));
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jobName** | [**string**] | The job identifier. | defaults to undefined
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

# **getJobScaleStatus**
> JobScaleStatusResponse getJobScaleStatus()


### Example


```typescript
import { nomad-client } from '';
import * as fs from 'fs';

const configuration = nomad-client.createConfiguration();
const apiInstance = new nomad-client.JobsApi(configuration);

let body:nomad-client.JobsApiGetJobScaleStatusRequest = {
  // string | The job identifier.
  jobName: "jobName_example",
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

apiInstance.getJobScaleStatus(body).then((data:any) => {
  console.log('API called successfully. Returned data: ' + data);
}).catch((error:any) => console.error(error));
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jobName** | [**string**] | The job identifier. | defaults to undefined
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

**JobScaleStatusResponse**

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

# **getJobSummary**
> JobSummary getJobSummary()


### Example


```typescript
import { nomad-client } from '';
import * as fs from 'fs';

const configuration = nomad-client.createConfiguration();
const apiInstance = new nomad-client.JobsApi(configuration);

let body:nomad-client.JobsApiGetJobSummaryRequest = {
  // string | The job identifier.
  jobName: "jobName_example",
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

apiInstance.getJobSummary(body).then((data:any) => {
  console.log('API called successfully. Returned data: ' + data);
}).catch((error:any) => console.error(error));
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jobName** | [**string**] | The job identifier. | defaults to undefined
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

**JobSummary**

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

# **getJobVersions**
> JobVersionsResponse getJobVersions()


### Example


```typescript
import { nomad-client } from '';
import * as fs from 'fs';

const configuration = nomad-client.createConfiguration();
const apiInstance = new nomad-client.JobsApi(configuration);

let body:nomad-client.JobsApiGetJobVersionsRequest = {
  // string | The job identifier.
  jobName: "jobName_example",
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
  // boolean | Boolean flag indicating whether to compute job diffs. (optional)
  diffs: true,
};

apiInstance.getJobVersions(body).then((data:any) => {
  console.log('API called successfully. Returned data: ' + data);
}).catch((error:any) => console.error(error));
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jobName** | [**string**] | The job identifier. | defaults to undefined
 **region** | [**string**] | Filters results based on the specified region. | (optional) defaults to undefined
 **namespace** | [**string**] | Filters results based on the specified namespace. | (optional) defaults to undefined
 **index** | [**number**] | If set, wait until query exceeds given index. Must be provided with WaitParam. | (optional) defaults to undefined
 **wait** | [**string**] | Provided with IndexParam to wait for change. | (optional) defaults to undefined
 **stale** | [**string**] | If present, results will include stale reads. | (optional) defaults to undefined
 **prefix** | [**string**] | Constrains results to jobs that start with the defined prefix | (optional) defaults to undefined
 **xNomadToken** | [**string**] | A Nomad ACL token. | (optional) defaults to undefined
 **perPage** | [**number**] | Maximum number of results to return. | (optional) defaults to undefined
 **nextToken** | [**string**] | Indicates where to start paging for queries that support pagination. | (optional) defaults to undefined
 **diffs** | [**boolean**] | Boolean flag indicating whether to compute job diffs. | (optional) defaults to undefined


### Return type

**JobVersionsResponse**

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

# **getJobs**
> Array<JobListStub> getJobs()


### Example


```typescript
import { nomad-client } from '';
import * as fs from 'fs';

const configuration = nomad-client.createConfiguration();
const apiInstance = new nomad-client.JobsApi(configuration);

let body:nomad-client.JobsApiGetJobsRequest = {
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

apiInstance.getJobs(body).then((data:any) => {
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

**Array<JobListStub>**

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

# **postJob**
> JobRegisterResponse postJob(jobRegisterRequest)


### Example


```typescript
import { nomad-client } from '';
import * as fs from 'fs';

const configuration = nomad-client.createConfiguration();
const apiInstance = new nomad-client.JobsApi(configuration);

let body:nomad-client.JobsApiPostJobRequest = {
  // string | The job identifier.
  jobName: "jobName_example",
  // JobRegisterRequest
  jobRegisterRequest: {
    enforceIndex: true,
    evalPriority: 1,
    job: {
      affinities: [
        {
          lTarget: "lTarget_example",
          operand: "operand_example",
          rTarget: "rTarget_example",
          weight: -128,
        },
      ],
      allAtOnce: true,
      constraints: [
        {
          lTarget: "lTarget_example",
          operand: "operand_example",
          rTarget: "rTarget_example",
        },
      ],
      consulNamespace: "consulNamespace_example",
      consulToken: "consulToken_example",
      createIndex: 0,
      datacenters: [
        "datacenters_example",
      ],
      dispatchIdempotencyToken: "dispatchIdempotencyToken_example",
      dispatched: true,
      ID: "ID_example",
      jobModifyIndex: 0,
      meta: {
        "key": "key_example",
      },
      migrate: {
        healthCheck: "healthCheck_example",
        healthyDeadline: 1,
        maxParallel: 1,
        minHealthyTime: 1,
      },
      modifyIndex: 0,
      multiregion: {
        regions: [
          {
            count: 1,
            datacenters: [
              "datacenters_example",
            ],
            meta: {
              "key": "key_example",
            },
            name: "name_example",
          },
        ],
        strategy: {
          maxParallel: 1,
          onFailure: "onFailure_example",
        },
      },
      name: "name_example",
      namespace: "namespace_example",
      nomadTokenID: "nomadTokenID_example",
      parameterizedJob: {
        metaOptional: [
          "metaOptional_example",
        ],
        metaRequired: [
          "metaRequired_example",
        ],
        payload: "payload_example",
      },
      parentID: "parentID_example",
      payload: 'YQ==',
      periodic: {
        enabled: true,
        prohibitOverlap: true,
        spec: "spec_example",
        specType: "specType_example",
        timeZone: "timeZone_example",
      },
      priority: 1,
      region: "region_example",
      reschedule: {
        attempts: 1,
        delay: 1,
        delayFunction: "delayFunction_example",
        interval: 1,
        maxDelay: 1,
        unlimited: true,
      },
      spreads: [
        {
          attribute: "attribute_example",
          spreadTarget: [
            {
              percent: 0,
              value: "value_example",
            },
          ],
          weight: -128,
        },
      ],
      stable: true,
      status: "status_example",
      statusDescription: "statusDescription_example",
      stop: true,
      submitTime: 1,
      taskGroups: [
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
          consul: {
            namespace: "namespace_example",
          },
          count: 1,
          ephemeralDisk: {
            migrate: true,
            sizeMB: 1,
            sticky: true,
          },
          maxClientDisconnect: 1,
          meta: {
            "key": "key_example",
          },
          migrate: {
            healthCheck: "healthCheck_example",
            healthyDeadline: 1,
            maxParallel: 1,
            minHealthyTime: 1,
          },
          name: "name_example",
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
          reschedulePolicy: {
            attempts: 1,
            delay: 1,
            delayFunction: "delayFunction_example",
            interval: 1,
            maxDelay: 1,
            unlimited: true,
          },
          restartPolicy: {
            attempts: 1,
            delay: 1,
            interval: 1,
            mode: "mode_example",
          },
          scaling: {
            createIndex: 0,
            enabled: true,
            ID: "ID_example",
            max: 1,
            min: 1,
            modifyIndex: 0,
            namespace: "namespace_example",
            policy: {
              "key": null,
            },
            target: {
              "key": "key_example",
            },
            type: "type_example",
          },
          services: [
            {
              address: "address_example",
              addressMode: "addressMode_example",
              canaryMeta: {
                "key": "key_example",
              },
              canaryTags: [
                "canaryTags_example",
              ],
              checkRestart: {
                grace: 1,
                ignoreWarnings: true,
                limit: 1,
              },
              checks: [
                {
                  addressMode: "addressMode_example",
                  advertise: "advertise_example",
                  args: [
                    "args_example",
                  ],
                  body: "body_example",
                  checkRestart: {
                    grace: 1,
                    ignoreWarnings: true,
                    limit: 1,
                  },
                  command: "command_example",
                  expose: true,
                  failuresBeforeCritical: 1,
                  gRPCService: "gRPCService_example",
                  gRPCUseTLS: true,
                  header: {
                    "key": [
                      "key_example",
                    ],
                  },
                  initialStatus: "initialStatus_example",
                  interval: 1,
                  method: "method_example",
                  name: "name_example",
                  onUpdate: "onUpdate_example",
                  path: "path_example",
                  portLabel: "portLabel_example",
                  protocol: "protocol_example",
                  successBeforePassing: 1,
                  tLSSkipVerify: true,
                  taskName: "taskName_example",
                  timeout: 1,
                  type: "type_example",
                },
              ],
              connect: {
                gateway: {
                  ingress: {
                    listeners: [
                      {
                        port: 1,
                        protocol: "protocol_example",
                        services: [
                          {
                            hosts: [
                              "hosts_example",
                            ],
                            name: "name_example",
                          },
                        ],
                      },
                    ],
                    TLS: {
                      enabled: true,
                    },
                  },
                  mesh: null,
                  proxy: {
                    config: {
                      "key": null,
                    },
                    connectTimeout: 1,
                    envoyDNSDiscoveryType: "envoyDNSDiscoveryType_example",
                    envoyGatewayBindAddresses: {
                      "key": {
                        address: "address_example",
                        name: "name_example",
                        port: 1,
                      },
                    },
                    envoyGatewayBindTaggedAddresses: true,
                    envoyGatewayNoDefaultBind: true,
                  },
                  terminating: {
                    services: [
                      {
                        cAFile: "cAFile_example",
                        certFile: "certFile_example",
                        keyFile: "keyFile_example",
                        name: "name_example",
                        SNI: "SNI_example",
                      },
                    ],
                  },
                },
                _native: true,
                sidecarService: {
                  disableDefaultTCPCheck: true,
                  port: "port_example",
                  proxy: {
                    config: {
                      "key": null,
                    },
                    exposeConfig: {
                      path: [
                        {
                          listenerPort: "listenerPort_example",
                          localPathPort: 1,
                          path: "path_example",
                          protocol: "protocol_example",
                        },
                      ],
                    },
                    localServiceAddress: "localServiceAddress_example",
                    localServicePort: 1,
                    upstreams: [
                      {
                        datacenter: "datacenter_example",
                        destinationName: "destinationName_example",
                        localBindAddress: "localBindAddress_example",
                        localBindPort: 1,
                        meshGateway: {
                          mode: "mode_example",
                        },
                      },
                    ],
                  },
                  tags: [
                    "tags_example",
                  ],
                },
                sidecarTask: {
                  config: {
                    "key": null,
                  },
                  driver: "driver_example",
                  env: {
                    "key": "key_example",
                  },
                  killSignal: "killSignal_example",
                  killTimeout: 1,
                  logConfig: {
                    maxFileSizeMB: 1,
                    maxFiles: 1,
                  },
                  meta: {
                    "key": "key_example",
                  },
                  name: "name_example",
                  resources: {
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
                  shutdownDelay: 1,
                  user: "user_example",
                },
              },
              enableTagOverride: true,
              meta: {
                "key": "key_example",
              },
              name: "name_example",
              onUpdate: "onUpdate_example",
              portLabel: "portLabel_example",
              provider: "provider_example",
              tags: [
                "tags_example",
              ],
              taskName: "taskName_example",
            },
          ],
          shutdownDelay: 1,
          spreads: [
            {
              attribute: "attribute_example",
              spreadTarget: [
                {
                  percent: 0,
                  value: "value_example",
                },
              ],
              weight: -128,
            },
          ],
          stopAfterClientDisconnect: 1,
          tasks: [
            {
              affinities: [
                {
                  lTarget: "lTarget_example",
                  operand: "operand_example",
                  rTarget: "rTarget_example",
                  weight: -128,
                },
              ],
              artifacts: [
                {
                  getterHeaders: {
                    "key": "key_example",
                  },
                  getterMode: "getterMode_example",
                  getterOptions: {
                    "key": "key_example",
                  },
                  getterSource: "getterSource_example",
                  relativeDest: "relativeDest_example",
                },
              ],
              cSIPluginConfig: {
                ID: "ID_example",
                mountDir: "mountDir_example",
                type: "type_example",
              },
              config: {
                "key": null,
              },
              constraints: [
                {
                  lTarget: "lTarget_example",
                  operand: "operand_example",
                  rTarget: "rTarget_example",
                },
              ],
              dispatchPayload: {
                file: "file_example",
              },
              driver: "driver_example",
              env: {
                "key": "key_example",
              },
              killSignal: "killSignal_example",
              killTimeout: 1,
              kind: "kind_example",
              leader: true,
              lifecycle: {
                hook: "hook_example",
                sidecar: true,
              },
              logConfig: {
                maxFileSizeMB: 1,
                maxFiles: 1,
              },
              meta: {
                "key": "key_example",
              },
              name: "name_example",
              resources: {
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
              restartPolicy: {
                attempts: 1,
                delay: 1,
                interval: 1,
                mode: "mode_example",
              },
              scalingPolicies: [
                {
                  createIndex: 0,
                  enabled: true,
                  ID: "ID_example",
                  max: 1,
                  min: 1,
                  modifyIndex: 0,
                  namespace: "namespace_example",
                  policy: {
                    "key": null,
                  },
                  target: {
                    "key": "key_example",
                  },
                  type: "type_example",
                },
              ],
              services: [
                {
                  address: "address_example",
                  addressMode: "addressMode_example",
                  canaryMeta: {
                    "key": "key_example",
                  },
                  canaryTags: [
                    "canaryTags_example",
                  ],
                  checkRestart: {
                    grace: 1,
                    ignoreWarnings: true,
                    limit: 1,
                  },
                  checks: [
                    {
                      addressMode: "addressMode_example",
                      advertise: "advertise_example",
                      args: [
                        "args_example",
                      ],
                      body: "body_example",
                      checkRestart: {
                        grace: 1,
                        ignoreWarnings: true,
                        limit: 1,
                      },
                      command: "command_example",
                      expose: true,
                      failuresBeforeCritical: 1,
                      gRPCService: "gRPCService_example",
                      gRPCUseTLS: true,
                      header: {
                        "key": [
                          "key_example",
                        ],
                      },
                      initialStatus: "initialStatus_example",
                      interval: 1,
                      method: "method_example",
                      name: "name_example",
                      onUpdate: "onUpdate_example",
                      path: "path_example",
                      portLabel: "portLabel_example",
                      protocol: "protocol_example",
                      successBeforePassing: 1,
                      tLSSkipVerify: true,
                      taskName: "taskName_example",
                      timeout: 1,
                      type: "type_example",
                    },
                  ],
                  connect: {
                    gateway: {
                      ingress: {
                        listeners: [
                          {
                            port: 1,
                            protocol: "protocol_example",
                            services: [
                              {
                                hosts: [
                                  "hosts_example",
                                ],
                                name: "name_example",
                              },
                            ],
                          },
                        ],
                        TLS: {
                          enabled: true,
                        },
                      },
                      mesh: null,
                      proxy: {
                        config: {
                          "key": null,
                        },
                        connectTimeout: 1,
                        envoyDNSDiscoveryType: "envoyDNSDiscoveryType_example",
                        envoyGatewayBindAddresses: {
                          "key": {
                            address: "address_example",
                            name: "name_example",
                            port: 1,
                          },
                        },
                        envoyGatewayBindTaggedAddresses: true,
                        envoyGatewayNoDefaultBind: true,
                      },
                      terminating: {
                        services: [
                          {
                            cAFile: "cAFile_example",
                            certFile: "certFile_example",
                            keyFile: "keyFile_example",
                            name: "name_example",
                            SNI: "SNI_example",
                          },
                        ],
                      },
                    },
                    _native: true,
                    sidecarService: {
                      disableDefaultTCPCheck: true,
                      port: "port_example",
                      proxy: {
                        config: {
                          "key": null,
                        },
                        exposeConfig: {
                          path: [
                            {
                              listenerPort: "listenerPort_example",
                              localPathPort: 1,
                              path: "path_example",
                              protocol: "protocol_example",
                            },
                          ],
                        },
                        localServiceAddress: "localServiceAddress_example",
                        localServicePort: 1,
                        upstreams: [
                          {
                            datacenter: "datacenter_example",
                            destinationName: "destinationName_example",
                            localBindAddress: "localBindAddress_example",
                            localBindPort: 1,
                            meshGateway: {
                              mode: "mode_example",
                            },
                          },
                        ],
                      },
                      tags: [
                        "tags_example",
                      ],
                    },
                    sidecarTask: {
                      config: {
                        "key": null,
                      },
                      driver: "driver_example",
                      env: {
                        "key": "key_example",
                      },
                      killSignal: "killSignal_example",
                      killTimeout: 1,
                      logConfig: {
                        maxFileSizeMB: 1,
                        maxFiles: 1,
                      },
                      meta: {
                        "key": "key_example",
                      },
                      name: "name_example",
                      resources: {
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
                      shutdownDelay: 1,
                      user: "user_example",
                    },
                  },
                  enableTagOverride: true,
                  meta: {
                    "key": "key_example",
                  },
                  name: "name_example",
                  onUpdate: "onUpdate_example",
                  portLabel: "portLabel_example",
                  provider: "provider_example",
                  tags: [
                    "tags_example",
                  ],
                  taskName: "taskName_example",
                },
              ],
              shutdownDelay: 1,
              templates: [
                {
                  changeMode: "changeMode_example",
                  changeSignal: "changeSignal_example",
                  destPath: "destPath_example",
                  embeddedTmpl: "embeddedTmpl_example",
                  envvars: true,
                  leftDelim: "leftDelim_example",
                  perms: "perms_example",
                  rightDelim: "rightDelim_example",
                  sourcePath: "sourcePath_example",
                  splay: 1,
                  vaultGrace: 1,
                  wait: {
                    max: 1,
                    min: 1,
                  },
                },
              ],
              user: "user_example",
              vault: {
                changeMode: "changeMode_example",
                changeSignal: "changeSignal_example",
                env: true,
                namespace: "namespace_example",
                policies: [
                  "policies_example",
                ],
              },
              volumeMounts: [
                {
                  destination: "destination_example",
                  propagationMode: "propagationMode_example",
                  readOnly: true,
                  volume: "volume_example",
                },
              ],
            },
          ],
          update: {
            autoPromote: true,
            autoRevert: true,
            canary: 1,
            healthCheck: "healthCheck_example",
            healthyDeadline: 1,
            maxParallel: 1,
            minHealthyTime: 1,
            progressDeadline: 1,
            stagger: 1,
          },
          volumes: {
            "key": {
              accessMode: "accessMode_example",
              attachmentMode: "attachmentMode_example",
              mountOptions: {
                fSType: "fSType_example",
                mountFlags: [
                  "mountFlags_example",
                ],
              },
              name: "name_example",
              perAlloc: true,
              readOnly: true,
              source: "source_example",
              type: "type_example",
            },
          },
        },
      ],
      type: "type_example",
      update: {
        autoPromote: true,
        autoRevert: true,
        canary: 1,
        healthCheck: "healthCheck_example",
        healthyDeadline: 1,
        maxParallel: 1,
        minHealthyTime: 1,
        progressDeadline: 1,
        stagger: 1,
      },
      vaultNamespace: "vaultNamespace_example",
      vaultToken: "vaultToken_example",
      version: 0,
    },
    jobModifyIndex: 0,
    namespace: "namespace_example",
    policyOverride: true,
    preserveCounts: true,
    region: "region_example",
    secretID: "secretID_example",
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

apiInstance.postJob(body).then((data:any) => {
  console.log('API called successfully. Returned data: ' + data);
}).catch((error:any) => console.error(error));
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jobRegisterRequest** | **JobRegisterRequest**|  |
 **jobName** | [**string**] | The job identifier. | defaults to undefined
 **region** | [**string**] | Filters results based on the specified region. | (optional) defaults to undefined
 **namespace** | [**string**] | Filters results based on the specified namespace. | (optional) defaults to undefined
 **xNomadToken** | [**string**] | A Nomad ACL token. | (optional) defaults to undefined
 **idempotencyToken** | [**string**] | Can be used to ensure operations are only run once. | (optional) defaults to undefined


### Return type

**JobRegisterResponse**

### Authorization

[X-Nomad-Token](README.md#X-Nomad-Token)

### HTTP request headers

 - **Content-Type**: application/json
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

# **postJobDispatch**
> JobDispatchResponse postJobDispatch(jobDispatchRequest)


### Example


```typescript
import { nomad-client } from '';
import * as fs from 'fs';

const configuration = nomad-client.createConfiguration();
const apiInstance = new nomad-client.JobsApi(configuration);

let body:nomad-client.JobsApiPostJobDispatchRequest = {
  // string | The job identifier.
  jobName: "jobName_example",
  // JobDispatchRequest
  jobDispatchRequest: {
    jobID: "jobID_example",
    meta: {
      "key": "key_example",
    },
    payload: 'YQ==',
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

apiInstance.postJobDispatch(body).then((data:any) => {
  console.log('API called successfully. Returned data: ' + data);
}).catch((error:any) => console.error(error));
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jobDispatchRequest** | **JobDispatchRequest**|  |
 **jobName** | [**string**] | The job identifier. | defaults to undefined
 **region** | [**string**] | Filters results based on the specified region. | (optional) defaults to undefined
 **namespace** | [**string**] | Filters results based on the specified namespace. | (optional) defaults to undefined
 **xNomadToken** | [**string**] | A Nomad ACL token. | (optional) defaults to undefined
 **idempotencyToken** | [**string**] | Can be used to ensure operations are only run once. | (optional) defaults to undefined


### Return type

**JobDispatchResponse**

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

# **postJobEvaluate**
> JobRegisterResponse postJobEvaluate(jobEvaluateRequest)


### Example


```typescript
import { nomad-client } from '';
import * as fs from 'fs';

const configuration = nomad-client.createConfiguration();
const apiInstance = new nomad-client.JobsApi(configuration);

let body:nomad-client.JobsApiPostJobEvaluateRequest = {
  // string | The job identifier.
  jobName: "jobName_example",
  // JobEvaluateRequest
  jobEvaluateRequest: {
    evalOptions: {
      forceReschedule: true,
    },
    jobID: "jobID_example",
    namespace: "namespace_example",
    region: "region_example",
    secretID: "secretID_example",
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

apiInstance.postJobEvaluate(body).then((data:any) => {
  console.log('API called successfully. Returned data: ' + data);
}).catch((error:any) => console.error(error));
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jobEvaluateRequest** | **JobEvaluateRequest**|  |
 **jobName** | [**string**] | The job identifier. | defaults to undefined
 **region** | [**string**] | Filters results based on the specified region. | (optional) defaults to undefined
 **namespace** | [**string**] | Filters results based on the specified namespace. | (optional) defaults to undefined
 **xNomadToken** | [**string**] | A Nomad ACL token. | (optional) defaults to undefined
 **idempotencyToken** | [**string**] | Can be used to ensure operations are only run once. | (optional) defaults to undefined


### Return type

**JobRegisterResponse**

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

# **postJobParse**
> Job postJobParse(jobsParseRequest)


### Example


```typescript
import { nomad-client } from '';
import * as fs from 'fs';

const configuration = nomad-client.createConfiguration();
const apiInstance = new nomad-client.JobsApi(configuration);

let body:nomad-client.JobsApiPostJobParseRequest = {
  // JobsParseRequest
  jobsParseRequest: {
    canonicalize: true,
    jobHCL: "jobHCL_example",
    hclv1: true,
  },
};

apiInstance.postJobParse(body).then((data:any) => {
  console.log('API called successfully. Returned data: ' + data);
}).catch((error:any) => console.error(error));
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jobsParseRequest** | **JobsParseRequest**|  |


### Return type

**Job**

### Authorization

[X-Nomad-Token](README.md#X-Nomad-Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** |  |  -  |
**400** | Bad request |  -  |
**403** | Forbidden |  -  |
**405** | Method not allowed |  -  |
**500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](README.md#documentation-for-api-endpoints) [[Back to Model list]](README.md#documentation-for-models) [[Back to README]](README.md)

# **postJobPeriodicForce**
> PeriodicForceResponse postJobPeriodicForce()


### Example


```typescript
import { nomad-client } from '';
import * as fs from 'fs';

const configuration = nomad-client.createConfiguration();
const apiInstance = new nomad-client.JobsApi(configuration);

let body:nomad-client.JobsApiPostJobPeriodicForceRequest = {
  // string | The job identifier.
  jobName: "jobName_example",
  // string | Filters results based on the specified region. (optional)
  region: "region_example",
  // string | Filters results based on the specified namespace. (optional)
  namespace: "namespace_example",
  // string | A Nomad ACL token. (optional)
  xNomadToken: "X-Nomad-Token_example",
  // string | Can be used to ensure operations are only run once. (optional)
  idempotencyToken: "idempotency_token_example",
};

apiInstance.postJobPeriodicForce(body).then((data:any) => {
  console.log('API called successfully. Returned data: ' + data);
}).catch((error:any) => console.error(error));
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jobName** | [**string**] | The job identifier. | defaults to undefined
 **region** | [**string**] | Filters results based on the specified region. | (optional) defaults to undefined
 **namespace** | [**string**] | Filters results based on the specified namespace. | (optional) defaults to undefined
 **xNomadToken** | [**string**] | A Nomad ACL token. | (optional) defaults to undefined
 **idempotencyToken** | [**string**] | Can be used to ensure operations are only run once. | (optional) defaults to undefined


### Return type

**PeriodicForceResponse**

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

# **postJobPlan**
> JobPlanResponse postJobPlan(jobPlanRequest)


### Example


```typescript
import { nomad-client } from '';
import * as fs from 'fs';

const configuration = nomad-client.createConfiguration();
const apiInstance = new nomad-client.JobsApi(configuration);

let body:nomad-client.JobsApiPostJobPlanRequest = {
  // string | The job identifier.
  jobName: "jobName_example",
  // JobPlanRequest
  jobPlanRequest: {
    diff: true,
    job: {
      affinities: [
        {
          lTarget: "lTarget_example",
          operand: "operand_example",
          rTarget: "rTarget_example",
          weight: -128,
        },
      ],
      allAtOnce: true,
      constraints: [
        {
          lTarget: "lTarget_example",
          operand: "operand_example",
          rTarget: "rTarget_example",
        },
      ],
      consulNamespace: "consulNamespace_example",
      consulToken: "consulToken_example",
      createIndex: 0,
      datacenters: [
        "datacenters_example",
      ],
      dispatchIdempotencyToken: "dispatchIdempotencyToken_example",
      dispatched: true,
      ID: "ID_example",
      jobModifyIndex: 0,
      meta: {
        "key": "key_example",
      },
      migrate: {
        healthCheck: "healthCheck_example",
        healthyDeadline: 1,
        maxParallel: 1,
        minHealthyTime: 1,
      },
      modifyIndex: 0,
      multiregion: {
        regions: [
          {
            count: 1,
            datacenters: [
              "datacenters_example",
            ],
            meta: {
              "key": "key_example",
            },
            name: "name_example",
          },
        ],
        strategy: {
          maxParallel: 1,
          onFailure: "onFailure_example",
        },
      },
      name: "name_example",
      namespace: "namespace_example",
      nomadTokenID: "nomadTokenID_example",
      parameterizedJob: {
        metaOptional: [
          "metaOptional_example",
        ],
        metaRequired: [
          "metaRequired_example",
        ],
        payload: "payload_example",
      },
      parentID: "parentID_example",
      payload: 'YQ==',
      periodic: {
        enabled: true,
        prohibitOverlap: true,
        spec: "spec_example",
        specType: "specType_example",
        timeZone: "timeZone_example",
      },
      priority: 1,
      region: "region_example",
      reschedule: {
        attempts: 1,
        delay: 1,
        delayFunction: "delayFunction_example",
        interval: 1,
        maxDelay: 1,
        unlimited: true,
      },
      spreads: [
        {
          attribute: "attribute_example",
          spreadTarget: [
            {
              percent: 0,
              value: "value_example",
            },
          ],
          weight: -128,
        },
      ],
      stable: true,
      status: "status_example",
      statusDescription: "statusDescription_example",
      stop: true,
      submitTime: 1,
      taskGroups: [
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
          consul: {
            namespace: "namespace_example",
          },
          count: 1,
          ephemeralDisk: {
            migrate: true,
            sizeMB: 1,
            sticky: true,
          },
          maxClientDisconnect: 1,
          meta: {
            "key": "key_example",
          },
          migrate: {
            healthCheck: "healthCheck_example",
            healthyDeadline: 1,
            maxParallel: 1,
            minHealthyTime: 1,
          },
          name: "name_example",
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
          reschedulePolicy: {
            attempts: 1,
            delay: 1,
            delayFunction: "delayFunction_example",
            interval: 1,
            maxDelay: 1,
            unlimited: true,
          },
          restartPolicy: {
            attempts: 1,
            delay: 1,
            interval: 1,
            mode: "mode_example",
          },
          scaling: {
            createIndex: 0,
            enabled: true,
            ID: "ID_example",
            max: 1,
            min: 1,
            modifyIndex: 0,
            namespace: "namespace_example",
            policy: {
              "key": null,
            },
            target: {
              "key": "key_example",
            },
            type: "type_example",
          },
          services: [
            {
              address: "address_example",
              addressMode: "addressMode_example",
              canaryMeta: {
                "key": "key_example",
              },
              canaryTags: [
                "canaryTags_example",
              ],
              checkRestart: {
                grace: 1,
                ignoreWarnings: true,
                limit: 1,
              },
              checks: [
                {
                  addressMode: "addressMode_example",
                  advertise: "advertise_example",
                  args: [
                    "args_example",
                  ],
                  body: "body_example",
                  checkRestart: {
                    grace: 1,
                    ignoreWarnings: true,
                    limit: 1,
                  },
                  command: "command_example",
                  expose: true,
                  failuresBeforeCritical: 1,
                  gRPCService: "gRPCService_example",
                  gRPCUseTLS: true,
                  header: {
                    "key": [
                      "key_example",
                    ],
                  },
                  initialStatus: "initialStatus_example",
                  interval: 1,
                  method: "method_example",
                  name: "name_example",
                  onUpdate: "onUpdate_example",
                  path: "path_example",
                  portLabel: "portLabel_example",
                  protocol: "protocol_example",
                  successBeforePassing: 1,
                  tLSSkipVerify: true,
                  taskName: "taskName_example",
                  timeout: 1,
                  type: "type_example",
                },
              ],
              connect: {
                gateway: {
                  ingress: {
                    listeners: [
                      {
                        port: 1,
                        protocol: "protocol_example",
                        services: [
                          {
                            hosts: [
                              "hosts_example",
                            ],
                            name: "name_example",
                          },
                        ],
                      },
                    ],
                    TLS: {
                      enabled: true,
                    },
                  },
                  mesh: null,
                  proxy: {
                    config: {
                      "key": null,
                    },
                    connectTimeout: 1,
                    envoyDNSDiscoveryType: "envoyDNSDiscoveryType_example",
                    envoyGatewayBindAddresses: {
                      "key": {
                        address: "address_example",
                        name: "name_example",
                        port: 1,
                      },
                    },
                    envoyGatewayBindTaggedAddresses: true,
                    envoyGatewayNoDefaultBind: true,
                  },
                  terminating: {
                    services: [
                      {
                        cAFile: "cAFile_example",
                        certFile: "certFile_example",
                        keyFile: "keyFile_example",
                        name: "name_example",
                        SNI: "SNI_example",
                      },
                    ],
                  },
                },
                _native: true,
                sidecarService: {
                  disableDefaultTCPCheck: true,
                  port: "port_example",
                  proxy: {
                    config: {
                      "key": null,
                    },
                    exposeConfig: {
                      path: [
                        {
                          listenerPort: "listenerPort_example",
                          localPathPort: 1,
                          path: "path_example",
                          protocol: "protocol_example",
                        },
                      ],
                    },
                    localServiceAddress: "localServiceAddress_example",
                    localServicePort: 1,
                    upstreams: [
                      {
                        datacenter: "datacenter_example",
                        destinationName: "destinationName_example",
                        localBindAddress: "localBindAddress_example",
                        localBindPort: 1,
                        meshGateway: {
                          mode: "mode_example",
                        },
                      },
                    ],
                  },
                  tags: [
                    "tags_example",
                  ],
                },
                sidecarTask: {
                  config: {
                    "key": null,
                  },
                  driver: "driver_example",
                  env: {
                    "key": "key_example",
                  },
                  killSignal: "killSignal_example",
                  killTimeout: 1,
                  logConfig: {
                    maxFileSizeMB: 1,
                    maxFiles: 1,
                  },
                  meta: {
                    "key": "key_example",
                  },
                  name: "name_example",
                  resources: {
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
                  shutdownDelay: 1,
                  user: "user_example",
                },
              },
              enableTagOverride: true,
              meta: {
                "key": "key_example",
              },
              name: "name_example",
              onUpdate: "onUpdate_example",
              portLabel: "portLabel_example",
              provider: "provider_example",
              tags: [
                "tags_example",
              ],
              taskName: "taskName_example",
            },
          ],
          shutdownDelay: 1,
          spreads: [
            {
              attribute: "attribute_example",
              spreadTarget: [
                {
                  percent: 0,
                  value: "value_example",
                },
              ],
              weight: -128,
            },
          ],
          stopAfterClientDisconnect: 1,
          tasks: [
            {
              affinities: [
                {
                  lTarget: "lTarget_example",
                  operand: "operand_example",
                  rTarget: "rTarget_example",
                  weight: -128,
                },
              ],
              artifacts: [
                {
                  getterHeaders: {
                    "key": "key_example",
                  },
                  getterMode: "getterMode_example",
                  getterOptions: {
                    "key": "key_example",
                  },
                  getterSource: "getterSource_example",
                  relativeDest: "relativeDest_example",
                },
              ],
              cSIPluginConfig: {
                ID: "ID_example",
                mountDir: "mountDir_example",
                type: "type_example",
              },
              config: {
                "key": null,
              },
              constraints: [
                {
                  lTarget: "lTarget_example",
                  operand: "operand_example",
                  rTarget: "rTarget_example",
                },
              ],
              dispatchPayload: {
                file: "file_example",
              },
              driver: "driver_example",
              env: {
                "key": "key_example",
              },
              killSignal: "killSignal_example",
              killTimeout: 1,
              kind: "kind_example",
              leader: true,
              lifecycle: {
                hook: "hook_example",
                sidecar: true,
              },
              logConfig: {
                maxFileSizeMB: 1,
                maxFiles: 1,
              },
              meta: {
                "key": "key_example",
              },
              name: "name_example",
              resources: {
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
              restartPolicy: {
                attempts: 1,
                delay: 1,
                interval: 1,
                mode: "mode_example",
              },
              scalingPolicies: [
                {
                  createIndex: 0,
                  enabled: true,
                  ID: "ID_example",
                  max: 1,
                  min: 1,
                  modifyIndex: 0,
                  namespace: "namespace_example",
                  policy: {
                    "key": null,
                  },
                  target: {
                    "key": "key_example",
                  },
                  type: "type_example",
                },
              ],
              services: [
                {
                  address: "address_example",
                  addressMode: "addressMode_example",
                  canaryMeta: {
                    "key": "key_example",
                  },
                  canaryTags: [
                    "canaryTags_example",
                  ],
                  checkRestart: {
                    grace: 1,
                    ignoreWarnings: true,
                    limit: 1,
                  },
                  checks: [
                    {
                      addressMode: "addressMode_example",
                      advertise: "advertise_example",
                      args: [
                        "args_example",
                      ],
                      body: "body_example",
                      checkRestart: {
                        grace: 1,
                        ignoreWarnings: true,
                        limit: 1,
                      },
                      command: "command_example",
                      expose: true,
                      failuresBeforeCritical: 1,
                      gRPCService: "gRPCService_example",
                      gRPCUseTLS: true,
                      header: {
                        "key": [
                          "key_example",
                        ],
                      },
                      initialStatus: "initialStatus_example",
                      interval: 1,
                      method: "method_example",
                      name: "name_example",
                      onUpdate: "onUpdate_example",
                      path: "path_example",
                      portLabel: "portLabel_example",
                      protocol: "protocol_example",
                      successBeforePassing: 1,
                      tLSSkipVerify: true,
                      taskName: "taskName_example",
                      timeout: 1,
                      type: "type_example",
                    },
                  ],
                  connect: {
                    gateway: {
                      ingress: {
                        listeners: [
                          {
                            port: 1,
                            protocol: "protocol_example",
                            services: [
                              {
                                hosts: [
                                  "hosts_example",
                                ],
                                name: "name_example",
                              },
                            ],
                          },
                        ],
                        TLS: {
                          enabled: true,
                        },
                      },
                      mesh: null,
                      proxy: {
                        config: {
                          "key": null,
                        },
                        connectTimeout: 1,
                        envoyDNSDiscoveryType: "envoyDNSDiscoveryType_example",
                        envoyGatewayBindAddresses: {
                          "key": {
                            address: "address_example",
                            name: "name_example",
                            port: 1,
                          },
                        },
                        envoyGatewayBindTaggedAddresses: true,
                        envoyGatewayNoDefaultBind: true,
                      },
                      terminating: {
                        services: [
                          {
                            cAFile: "cAFile_example",
                            certFile: "certFile_example",
                            keyFile: "keyFile_example",
                            name: "name_example",
                            SNI: "SNI_example",
                          },
                        ],
                      },
                    },
                    _native: true,
                    sidecarService: {
                      disableDefaultTCPCheck: true,
                      port: "port_example",
                      proxy: {
                        config: {
                          "key": null,
                        },
                        exposeConfig: {
                          path: [
                            {
                              listenerPort: "listenerPort_example",
                              localPathPort: 1,
                              path: "path_example",
                              protocol: "protocol_example",
                            },
                          ],
                        },
                        localServiceAddress: "localServiceAddress_example",
                        localServicePort: 1,
                        upstreams: [
                          {
                            datacenter: "datacenter_example",
                            destinationName: "destinationName_example",
                            localBindAddress: "localBindAddress_example",
                            localBindPort: 1,
                            meshGateway: {
                              mode: "mode_example",
                            },
                          },
                        ],
                      },
                      tags: [
                        "tags_example",
                      ],
                    },
                    sidecarTask: {
                      config: {
                        "key": null,
                      },
                      driver: "driver_example",
                      env: {
                        "key": "key_example",
                      },
                      killSignal: "killSignal_example",
                      killTimeout: 1,
                      logConfig: {
                        maxFileSizeMB: 1,
                        maxFiles: 1,
                      },
                      meta: {
                        "key": "key_example",
                      },
                      name: "name_example",
                      resources: {
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
                      shutdownDelay: 1,
                      user: "user_example",
                    },
                  },
                  enableTagOverride: true,
                  meta: {
                    "key": "key_example",
                  },
                  name: "name_example",
                  onUpdate: "onUpdate_example",
                  portLabel: "portLabel_example",
                  provider: "provider_example",
                  tags: [
                    "tags_example",
                  ],
                  taskName: "taskName_example",
                },
              ],
              shutdownDelay: 1,
              templates: [
                {
                  changeMode: "changeMode_example",
                  changeSignal: "changeSignal_example",
                  destPath: "destPath_example",
                  embeddedTmpl: "embeddedTmpl_example",
                  envvars: true,
                  leftDelim: "leftDelim_example",
                  perms: "perms_example",
                  rightDelim: "rightDelim_example",
                  sourcePath: "sourcePath_example",
                  splay: 1,
                  vaultGrace: 1,
                  wait: {
                    max: 1,
                    min: 1,
                  },
                },
              ],
              user: "user_example",
              vault: {
                changeMode: "changeMode_example",
                changeSignal: "changeSignal_example",
                env: true,
                namespace: "namespace_example",
                policies: [
                  "policies_example",
                ],
              },
              volumeMounts: [
                {
                  destination: "destination_example",
                  propagationMode: "propagationMode_example",
                  readOnly: true,
                  volume: "volume_example",
                },
              ],
            },
          ],
          update: {
            autoPromote: true,
            autoRevert: true,
            canary: 1,
            healthCheck: "healthCheck_example",
            healthyDeadline: 1,
            maxParallel: 1,
            minHealthyTime: 1,
            progressDeadline: 1,
            stagger: 1,
          },
          volumes: {
            "key": {
              accessMode: "accessMode_example",
              attachmentMode: "attachmentMode_example",
              mountOptions: {
                fSType: "fSType_example",
                mountFlags: [
                  "mountFlags_example",
                ],
              },
              name: "name_example",
              perAlloc: true,
              readOnly: true,
              source: "source_example",
              type: "type_example",
            },
          },
        },
      ],
      type: "type_example",
      update: {
        autoPromote: true,
        autoRevert: true,
        canary: 1,
        healthCheck: "healthCheck_example",
        healthyDeadline: 1,
        maxParallel: 1,
        minHealthyTime: 1,
        progressDeadline: 1,
        stagger: 1,
      },
      vaultNamespace: "vaultNamespace_example",
      vaultToken: "vaultToken_example",
      version: 0,
    },
    namespace: "namespace_example",
    policyOverride: true,
    region: "region_example",
    secretID: "secretID_example",
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

apiInstance.postJobPlan(body).then((data:any) => {
  console.log('API called successfully. Returned data: ' + data);
}).catch((error:any) => console.error(error));
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jobPlanRequest** | **JobPlanRequest**|  |
 **jobName** | [**string**] | The job identifier. | defaults to undefined
 **region** | [**string**] | Filters results based on the specified region. | (optional) defaults to undefined
 **namespace** | [**string**] | Filters results based on the specified namespace. | (optional) defaults to undefined
 **xNomadToken** | [**string**] | A Nomad ACL token. | (optional) defaults to undefined
 **idempotencyToken** | [**string**] | Can be used to ensure operations are only run once. | (optional) defaults to undefined


### Return type

**JobPlanResponse**

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

# **postJobRevert**
> JobRegisterResponse postJobRevert(jobRevertRequest)


### Example


```typescript
import { nomad-client } from '';
import * as fs from 'fs';

const configuration = nomad-client.createConfiguration();
const apiInstance = new nomad-client.JobsApi(configuration);

let body:nomad-client.JobsApiPostJobRevertRequest = {
  // string | The job identifier.
  jobName: "jobName_example",
  // JobRevertRequest
  jobRevertRequest: {
    consulToken: "consulToken_example",
    enforcePriorVersion: 0,
    jobID: "jobID_example",
    jobVersion: 0,
    namespace: "namespace_example",
    region: "region_example",
    secretID: "secretID_example",
    vaultToken: "vaultToken_example",
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

apiInstance.postJobRevert(body).then((data:any) => {
  console.log('API called successfully. Returned data: ' + data);
}).catch((error:any) => console.error(error));
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jobRevertRequest** | **JobRevertRequest**|  |
 **jobName** | [**string**] | The job identifier. | defaults to undefined
 **region** | [**string**] | Filters results based on the specified region. | (optional) defaults to undefined
 **namespace** | [**string**] | Filters results based on the specified namespace. | (optional) defaults to undefined
 **xNomadToken** | [**string**] | A Nomad ACL token. | (optional) defaults to undefined
 **idempotencyToken** | [**string**] | Can be used to ensure operations are only run once. | (optional) defaults to undefined


### Return type

**JobRegisterResponse**

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

# **postJobScalingRequest**
> JobRegisterResponse postJobScalingRequest(scalingRequest)


### Example


```typescript
import { nomad-client } from '';
import * as fs from 'fs';

const configuration = nomad-client.createConfiguration();
const apiInstance = new nomad-client.JobsApi(configuration);

let body:nomad-client.JobsApiPostJobScalingRequestRequest = {
  // string | The job identifier.
  jobName: "jobName_example",
  // ScalingRequest
  scalingRequest: {
    count: 1,
    error: true,
    message: "message_example",
    meta: {
      "key": null,
    },
    namespace: "namespace_example",
    policyOverride: true,
    region: "region_example",
    secretID: "secretID_example",
    target: {
      "key": "key_example",
    },
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

apiInstance.postJobScalingRequest(body).then((data:any) => {
  console.log('API called successfully. Returned data: ' + data);
}).catch((error:any) => console.error(error));
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scalingRequest** | **ScalingRequest**|  |
 **jobName** | [**string**] | The job identifier. | defaults to undefined
 **region** | [**string**] | Filters results based on the specified region. | (optional) defaults to undefined
 **namespace** | [**string**] | Filters results based on the specified namespace. | (optional) defaults to undefined
 **xNomadToken** | [**string**] | A Nomad ACL token. | (optional) defaults to undefined
 **idempotencyToken** | [**string**] | Can be used to ensure operations are only run once. | (optional) defaults to undefined


### Return type

**JobRegisterResponse**

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

# **postJobStability**
> JobStabilityResponse postJobStability(jobStabilityRequest)


### Example


```typescript
import { nomad-client } from '';
import * as fs from 'fs';

const configuration = nomad-client.createConfiguration();
const apiInstance = new nomad-client.JobsApi(configuration);

let body:nomad-client.JobsApiPostJobStabilityRequest = {
  // string | The job identifier.
  jobName: "jobName_example",
  // JobStabilityRequest
  jobStabilityRequest: {
    jobID: "jobID_example",
    jobVersion: 0,
    namespace: "namespace_example",
    region: "region_example",
    secretID: "secretID_example",
    stable: true,
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

apiInstance.postJobStability(body).then((data:any) => {
  console.log('API called successfully. Returned data: ' + data);
}).catch((error:any) => console.error(error));
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jobStabilityRequest** | **JobStabilityRequest**|  |
 **jobName** | [**string**] | The job identifier. | defaults to undefined
 **region** | [**string**] | Filters results based on the specified region. | (optional) defaults to undefined
 **namespace** | [**string**] | Filters results based on the specified namespace. | (optional) defaults to undefined
 **xNomadToken** | [**string**] | A Nomad ACL token. | (optional) defaults to undefined
 **idempotencyToken** | [**string**] | Can be used to ensure operations are only run once. | (optional) defaults to undefined


### Return type

**JobStabilityResponse**

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

# **postJobValidateRequest**
> JobValidateResponse postJobValidateRequest(jobValidateRequest)


### Example


```typescript
import { nomad-client } from '';
import * as fs from 'fs';

const configuration = nomad-client.createConfiguration();
const apiInstance = new nomad-client.JobsApi(configuration);

let body:nomad-client.JobsApiPostJobValidateRequestRequest = {
  // JobValidateRequest
  jobValidateRequest: {
    job: {
      affinities: [
        {
          lTarget: "lTarget_example",
          operand: "operand_example",
          rTarget: "rTarget_example",
          weight: -128,
        },
      ],
      allAtOnce: true,
      constraints: [
        {
          lTarget: "lTarget_example",
          operand: "operand_example",
          rTarget: "rTarget_example",
        },
      ],
      consulNamespace: "consulNamespace_example",
      consulToken: "consulToken_example",
      createIndex: 0,
      datacenters: [
        "datacenters_example",
      ],
      dispatchIdempotencyToken: "dispatchIdempotencyToken_example",
      dispatched: true,
      ID: "ID_example",
      jobModifyIndex: 0,
      meta: {
        "key": "key_example",
      },
      migrate: {
        healthCheck: "healthCheck_example",
        healthyDeadline: 1,
        maxParallel: 1,
        minHealthyTime: 1,
      },
      modifyIndex: 0,
      multiregion: {
        regions: [
          {
            count: 1,
            datacenters: [
              "datacenters_example",
            ],
            meta: {
              "key": "key_example",
            },
            name: "name_example",
          },
        ],
        strategy: {
          maxParallel: 1,
          onFailure: "onFailure_example",
        },
      },
      name: "name_example",
      namespace: "namespace_example",
      nomadTokenID: "nomadTokenID_example",
      parameterizedJob: {
        metaOptional: [
          "metaOptional_example",
        ],
        metaRequired: [
          "metaRequired_example",
        ],
        payload: "payload_example",
      },
      parentID: "parentID_example",
      payload: 'YQ==',
      periodic: {
        enabled: true,
        prohibitOverlap: true,
        spec: "spec_example",
        specType: "specType_example",
        timeZone: "timeZone_example",
      },
      priority: 1,
      region: "region_example",
      reschedule: {
        attempts: 1,
        delay: 1,
        delayFunction: "delayFunction_example",
        interval: 1,
        maxDelay: 1,
        unlimited: true,
      },
      spreads: [
        {
          attribute: "attribute_example",
          spreadTarget: [
            {
              percent: 0,
              value: "value_example",
            },
          ],
          weight: -128,
        },
      ],
      stable: true,
      status: "status_example",
      statusDescription: "statusDescription_example",
      stop: true,
      submitTime: 1,
      taskGroups: [
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
          consul: {
            namespace: "namespace_example",
          },
          count: 1,
          ephemeralDisk: {
            migrate: true,
            sizeMB: 1,
            sticky: true,
          },
          maxClientDisconnect: 1,
          meta: {
            "key": "key_example",
          },
          migrate: {
            healthCheck: "healthCheck_example",
            healthyDeadline: 1,
            maxParallel: 1,
            minHealthyTime: 1,
          },
          name: "name_example",
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
          reschedulePolicy: {
            attempts: 1,
            delay: 1,
            delayFunction: "delayFunction_example",
            interval: 1,
            maxDelay: 1,
            unlimited: true,
          },
          restartPolicy: {
            attempts: 1,
            delay: 1,
            interval: 1,
            mode: "mode_example",
          },
          scaling: {
            createIndex: 0,
            enabled: true,
            ID: "ID_example",
            max: 1,
            min: 1,
            modifyIndex: 0,
            namespace: "namespace_example",
            policy: {
              "key": null,
            },
            target: {
              "key": "key_example",
            },
            type: "type_example",
          },
          services: [
            {
              address: "address_example",
              addressMode: "addressMode_example",
              canaryMeta: {
                "key": "key_example",
              },
              canaryTags: [
                "canaryTags_example",
              ],
              checkRestart: {
                grace: 1,
                ignoreWarnings: true,
                limit: 1,
              },
              checks: [
                {
                  addressMode: "addressMode_example",
                  advertise: "advertise_example",
                  args: [
                    "args_example",
                  ],
                  body: "body_example",
                  checkRestart: {
                    grace: 1,
                    ignoreWarnings: true,
                    limit: 1,
                  },
                  command: "command_example",
                  expose: true,
                  failuresBeforeCritical: 1,
                  gRPCService: "gRPCService_example",
                  gRPCUseTLS: true,
                  header: {
                    "key": [
                      "key_example",
                    ],
                  },
                  initialStatus: "initialStatus_example",
                  interval: 1,
                  method: "method_example",
                  name: "name_example",
                  onUpdate: "onUpdate_example",
                  path: "path_example",
                  portLabel: "portLabel_example",
                  protocol: "protocol_example",
                  successBeforePassing: 1,
                  tLSSkipVerify: true,
                  taskName: "taskName_example",
                  timeout: 1,
                  type: "type_example",
                },
              ],
              connect: {
                gateway: {
                  ingress: {
                    listeners: [
                      {
                        port: 1,
                        protocol: "protocol_example",
                        services: [
                          {
                            hosts: [
                              "hosts_example",
                            ],
                            name: "name_example",
                          },
                        ],
                      },
                    ],
                    TLS: {
                      enabled: true,
                    },
                  },
                  mesh: null,
                  proxy: {
                    config: {
                      "key": null,
                    },
                    connectTimeout: 1,
                    envoyDNSDiscoveryType: "envoyDNSDiscoveryType_example",
                    envoyGatewayBindAddresses: {
                      "key": {
                        address: "address_example",
                        name: "name_example",
                        port: 1,
                      },
                    },
                    envoyGatewayBindTaggedAddresses: true,
                    envoyGatewayNoDefaultBind: true,
                  },
                  terminating: {
                    services: [
                      {
                        cAFile: "cAFile_example",
                        certFile: "certFile_example",
                        keyFile: "keyFile_example",
                        name: "name_example",
                        SNI: "SNI_example",
                      },
                    ],
                  },
                },
                _native: true,
                sidecarService: {
                  disableDefaultTCPCheck: true,
                  port: "port_example",
                  proxy: {
                    config: {
                      "key": null,
                    },
                    exposeConfig: {
                      path: [
                        {
                          listenerPort: "listenerPort_example",
                          localPathPort: 1,
                          path: "path_example",
                          protocol: "protocol_example",
                        },
                      ],
                    },
                    localServiceAddress: "localServiceAddress_example",
                    localServicePort: 1,
                    upstreams: [
                      {
                        datacenter: "datacenter_example",
                        destinationName: "destinationName_example",
                        localBindAddress: "localBindAddress_example",
                        localBindPort: 1,
                        meshGateway: {
                          mode: "mode_example",
                        },
                      },
                    ],
                  },
                  tags: [
                    "tags_example",
                  ],
                },
                sidecarTask: {
                  config: {
                    "key": null,
                  },
                  driver: "driver_example",
                  env: {
                    "key": "key_example",
                  },
                  killSignal: "killSignal_example",
                  killTimeout: 1,
                  logConfig: {
                    maxFileSizeMB: 1,
                    maxFiles: 1,
                  },
                  meta: {
                    "key": "key_example",
                  },
                  name: "name_example",
                  resources: {
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
                  shutdownDelay: 1,
                  user: "user_example",
                },
              },
              enableTagOverride: true,
              meta: {
                "key": "key_example",
              },
              name: "name_example",
              onUpdate: "onUpdate_example",
              portLabel: "portLabel_example",
              provider: "provider_example",
              tags: [
                "tags_example",
              ],
              taskName: "taskName_example",
            },
          ],
          shutdownDelay: 1,
          spreads: [
            {
              attribute: "attribute_example",
              spreadTarget: [
                {
                  percent: 0,
                  value: "value_example",
                },
              ],
              weight: -128,
            },
          ],
          stopAfterClientDisconnect: 1,
          tasks: [
            {
              affinities: [
                {
                  lTarget: "lTarget_example",
                  operand: "operand_example",
                  rTarget: "rTarget_example",
                  weight: -128,
                },
              ],
              artifacts: [
                {
                  getterHeaders: {
                    "key": "key_example",
                  },
                  getterMode: "getterMode_example",
                  getterOptions: {
                    "key": "key_example",
                  },
                  getterSource: "getterSource_example",
                  relativeDest: "relativeDest_example",
                },
              ],
              cSIPluginConfig: {
                ID: "ID_example",
                mountDir: "mountDir_example",
                type: "type_example",
              },
              config: {
                "key": null,
              },
              constraints: [
                {
                  lTarget: "lTarget_example",
                  operand: "operand_example",
                  rTarget: "rTarget_example",
                },
              ],
              dispatchPayload: {
                file: "file_example",
              },
              driver: "driver_example",
              env: {
                "key": "key_example",
              },
              killSignal: "killSignal_example",
              killTimeout: 1,
              kind: "kind_example",
              leader: true,
              lifecycle: {
                hook: "hook_example",
                sidecar: true,
              },
              logConfig: {
                maxFileSizeMB: 1,
                maxFiles: 1,
              },
              meta: {
                "key": "key_example",
              },
              name: "name_example",
              resources: {
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
              restartPolicy: {
                attempts: 1,
                delay: 1,
                interval: 1,
                mode: "mode_example",
              },
              scalingPolicies: [
                {
                  createIndex: 0,
                  enabled: true,
                  ID: "ID_example",
                  max: 1,
                  min: 1,
                  modifyIndex: 0,
                  namespace: "namespace_example",
                  policy: {
                    "key": null,
                  },
                  target: {
                    "key": "key_example",
                  },
                  type: "type_example",
                },
              ],
              services: [
                {
                  address: "address_example",
                  addressMode: "addressMode_example",
                  canaryMeta: {
                    "key": "key_example",
                  },
                  canaryTags: [
                    "canaryTags_example",
                  ],
                  checkRestart: {
                    grace: 1,
                    ignoreWarnings: true,
                    limit: 1,
                  },
                  checks: [
                    {
                      addressMode: "addressMode_example",
                      advertise: "advertise_example",
                      args: [
                        "args_example",
                      ],
                      body: "body_example",
                      checkRestart: {
                        grace: 1,
                        ignoreWarnings: true,
                        limit: 1,
                      },
                      command: "command_example",
                      expose: true,
                      failuresBeforeCritical: 1,
                      gRPCService: "gRPCService_example",
                      gRPCUseTLS: true,
                      header: {
                        "key": [
                          "key_example",
                        ],
                      },
                      initialStatus: "initialStatus_example",
                      interval: 1,
                      method: "method_example",
                      name: "name_example",
                      onUpdate: "onUpdate_example",
                      path: "path_example",
                      portLabel: "portLabel_example",
                      protocol: "protocol_example",
                      successBeforePassing: 1,
                      tLSSkipVerify: true,
                      taskName: "taskName_example",
                      timeout: 1,
                      type: "type_example",
                    },
                  ],
                  connect: {
                    gateway: {
                      ingress: {
                        listeners: [
                          {
                            port: 1,
                            protocol: "protocol_example",
                            services: [
                              {
                                hosts: [
                                  "hosts_example",
                                ],
                                name: "name_example",
                              },
                            ],
                          },
                        ],
                        TLS: {
                          enabled: true,
                        },
                      },
                      mesh: null,
                      proxy: {
                        config: {
                          "key": null,
                        },
                        connectTimeout: 1,
                        envoyDNSDiscoveryType: "envoyDNSDiscoveryType_example",
                        envoyGatewayBindAddresses: {
                          "key": {
                            address: "address_example",
                            name: "name_example",
                            port: 1,
                          },
                        },
                        envoyGatewayBindTaggedAddresses: true,
                        envoyGatewayNoDefaultBind: true,
                      },
                      terminating: {
                        services: [
                          {
                            cAFile: "cAFile_example",
                            certFile: "certFile_example",
                            keyFile: "keyFile_example",
                            name: "name_example",
                            SNI: "SNI_example",
                          },
                        ],
                      },
                    },
                    _native: true,
                    sidecarService: {
                      disableDefaultTCPCheck: true,
                      port: "port_example",
                      proxy: {
                        config: {
                          "key": null,
                        },
                        exposeConfig: {
                          path: [
                            {
                              listenerPort: "listenerPort_example",
                              localPathPort: 1,
                              path: "path_example",
                              protocol: "protocol_example",
                            },
                          ],
                        },
                        localServiceAddress: "localServiceAddress_example",
                        localServicePort: 1,
                        upstreams: [
                          {
                            datacenter: "datacenter_example",
                            destinationName: "destinationName_example",
                            localBindAddress: "localBindAddress_example",
                            localBindPort: 1,
                            meshGateway: {
                              mode: "mode_example",
                            },
                          },
                        ],
                      },
                      tags: [
                        "tags_example",
                      ],
                    },
                    sidecarTask: {
                      config: {
                        "key": null,
                      },
                      driver: "driver_example",
                      env: {
                        "key": "key_example",
                      },
                      killSignal: "killSignal_example",
                      killTimeout: 1,
                      logConfig: {
                        maxFileSizeMB: 1,
                        maxFiles: 1,
                      },
                      meta: {
                        "key": "key_example",
                      },
                      name: "name_example",
                      resources: {
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
                      shutdownDelay: 1,
                      user: "user_example",
                    },
                  },
                  enableTagOverride: true,
                  meta: {
                    "key": "key_example",
                  },
                  name: "name_example",
                  onUpdate: "onUpdate_example",
                  portLabel: "portLabel_example",
                  provider: "provider_example",
                  tags: [
                    "tags_example",
                  ],
                  taskName: "taskName_example",
                },
              ],
              shutdownDelay: 1,
              templates: [
                {
                  changeMode: "changeMode_example",
                  changeSignal: "changeSignal_example",
                  destPath: "destPath_example",
                  embeddedTmpl: "embeddedTmpl_example",
                  envvars: true,
                  leftDelim: "leftDelim_example",
                  perms: "perms_example",
                  rightDelim: "rightDelim_example",
                  sourcePath: "sourcePath_example",
                  splay: 1,
                  vaultGrace: 1,
                  wait: {
                    max: 1,
                    min: 1,
                  },
                },
              ],
              user: "user_example",
              vault: {
                changeMode: "changeMode_example",
                changeSignal: "changeSignal_example",
                env: true,
                namespace: "namespace_example",
                policies: [
                  "policies_example",
                ],
              },
              volumeMounts: [
                {
                  destination: "destination_example",
                  propagationMode: "propagationMode_example",
                  readOnly: true,
                  volume: "volume_example",
                },
              ],
            },
          ],
          update: {
            autoPromote: true,
            autoRevert: true,
            canary: 1,
            healthCheck: "healthCheck_example",
            healthyDeadline: 1,
            maxParallel: 1,
            minHealthyTime: 1,
            progressDeadline: 1,
            stagger: 1,
          },
          volumes: {
            "key": {
              accessMode: "accessMode_example",
              attachmentMode: "attachmentMode_example",
              mountOptions: {
                fSType: "fSType_example",
                mountFlags: [
                  "mountFlags_example",
                ],
              },
              name: "name_example",
              perAlloc: true,
              readOnly: true,
              source: "source_example",
              type: "type_example",
            },
          },
        },
      ],
      type: "type_example",
      update: {
        autoPromote: true,
        autoRevert: true,
        canary: 1,
        healthCheck: "healthCheck_example",
        healthyDeadline: 1,
        maxParallel: 1,
        minHealthyTime: 1,
        progressDeadline: 1,
        stagger: 1,
      },
      vaultNamespace: "vaultNamespace_example",
      vaultToken: "vaultToken_example",
      version: 0,
    },
    namespace: "namespace_example",
    region: "region_example",
    secretID: "secretID_example",
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

apiInstance.postJobValidateRequest(body).then((data:any) => {
  console.log('API called successfully. Returned data: ' + data);
}).catch((error:any) => console.error(error));
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jobValidateRequest** | **JobValidateRequest**|  |
 **region** | [**string**] | Filters results based on the specified region. | (optional) defaults to undefined
 **namespace** | [**string**] | Filters results based on the specified namespace. | (optional) defaults to undefined
 **xNomadToken** | [**string**] | A Nomad ACL token. | (optional) defaults to undefined
 **idempotencyToken** | [**string**] | Can be used to ensure operations are only run once. | (optional) defaults to undefined


### Return type

**JobValidateResponse**

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

# **registerJob**
> JobRegisterResponse registerJob(jobRegisterRequest)


### Example


```typescript
import { nomad-client } from '';
import * as fs from 'fs';

const configuration = nomad-client.createConfiguration();
const apiInstance = new nomad-client.JobsApi(configuration);

let body:nomad-client.JobsApiRegisterJobRequest = {
  // JobRegisterRequest
  jobRegisterRequest: {
    enforceIndex: true,
    evalPriority: 1,
    job: {
      affinities: [
        {
          lTarget: "lTarget_example",
          operand: "operand_example",
          rTarget: "rTarget_example",
          weight: -128,
        },
      ],
      allAtOnce: true,
      constraints: [
        {
          lTarget: "lTarget_example",
          operand: "operand_example",
          rTarget: "rTarget_example",
        },
      ],
      consulNamespace: "consulNamespace_example",
      consulToken: "consulToken_example",
      createIndex: 0,
      datacenters: [
        "datacenters_example",
      ],
      dispatchIdempotencyToken: "dispatchIdempotencyToken_example",
      dispatched: true,
      ID: "ID_example",
      jobModifyIndex: 0,
      meta: {
        "key": "key_example",
      },
      migrate: {
        healthCheck: "healthCheck_example",
        healthyDeadline: 1,
        maxParallel: 1,
        minHealthyTime: 1,
      },
      modifyIndex: 0,
      multiregion: {
        regions: [
          {
            count: 1,
            datacenters: [
              "datacenters_example",
            ],
            meta: {
              "key": "key_example",
            },
            name: "name_example",
          },
        ],
        strategy: {
          maxParallel: 1,
          onFailure: "onFailure_example",
        },
      },
      name: "name_example",
      namespace: "namespace_example",
      nomadTokenID: "nomadTokenID_example",
      parameterizedJob: {
        metaOptional: [
          "metaOptional_example",
        ],
        metaRequired: [
          "metaRequired_example",
        ],
        payload: "payload_example",
      },
      parentID: "parentID_example",
      payload: 'YQ==',
      periodic: {
        enabled: true,
        prohibitOverlap: true,
        spec: "spec_example",
        specType: "specType_example",
        timeZone: "timeZone_example",
      },
      priority: 1,
      region: "region_example",
      reschedule: {
        attempts: 1,
        delay: 1,
        delayFunction: "delayFunction_example",
        interval: 1,
        maxDelay: 1,
        unlimited: true,
      },
      spreads: [
        {
          attribute: "attribute_example",
          spreadTarget: [
            {
              percent: 0,
              value: "value_example",
            },
          ],
          weight: -128,
        },
      ],
      stable: true,
      status: "status_example",
      statusDescription: "statusDescription_example",
      stop: true,
      submitTime: 1,
      taskGroups: [
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
          consul: {
            namespace: "namespace_example",
          },
          count: 1,
          ephemeralDisk: {
            migrate: true,
            sizeMB: 1,
            sticky: true,
          },
          maxClientDisconnect: 1,
          meta: {
            "key": "key_example",
          },
          migrate: {
            healthCheck: "healthCheck_example",
            healthyDeadline: 1,
            maxParallel: 1,
            minHealthyTime: 1,
          },
          name: "name_example",
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
          reschedulePolicy: {
            attempts: 1,
            delay: 1,
            delayFunction: "delayFunction_example",
            interval: 1,
            maxDelay: 1,
            unlimited: true,
          },
          restartPolicy: {
            attempts: 1,
            delay: 1,
            interval: 1,
            mode: "mode_example",
          },
          scaling: {
            createIndex: 0,
            enabled: true,
            ID: "ID_example",
            max: 1,
            min: 1,
            modifyIndex: 0,
            namespace: "namespace_example",
            policy: {
              "key": null,
            },
            target: {
              "key": "key_example",
            },
            type: "type_example",
          },
          services: [
            {
              address: "address_example",
              addressMode: "addressMode_example",
              canaryMeta: {
                "key": "key_example",
              },
              canaryTags: [
                "canaryTags_example",
              ],
              checkRestart: {
                grace: 1,
                ignoreWarnings: true,
                limit: 1,
              },
              checks: [
                {
                  addressMode: "addressMode_example",
                  advertise: "advertise_example",
                  args: [
                    "args_example",
                  ],
                  body: "body_example",
                  checkRestart: {
                    grace: 1,
                    ignoreWarnings: true,
                    limit: 1,
                  },
                  command: "command_example",
                  expose: true,
                  failuresBeforeCritical: 1,
                  gRPCService: "gRPCService_example",
                  gRPCUseTLS: true,
                  header: {
                    "key": [
                      "key_example",
                    ],
                  },
                  initialStatus: "initialStatus_example",
                  interval: 1,
                  method: "method_example",
                  name: "name_example",
                  onUpdate: "onUpdate_example",
                  path: "path_example",
                  portLabel: "portLabel_example",
                  protocol: "protocol_example",
                  successBeforePassing: 1,
                  tLSSkipVerify: true,
                  taskName: "taskName_example",
                  timeout: 1,
                  type: "type_example",
                },
              ],
              connect: {
                gateway: {
                  ingress: {
                    listeners: [
                      {
                        port: 1,
                        protocol: "protocol_example",
                        services: [
                          {
                            hosts: [
                              "hosts_example",
                            ],
                            name: "name_example",
                          },
                        ],
                      },
                    ],
                    TLS: {
                      enabled: true,
                    },
                  },
                  mesh: null,
                  proxy: {
                    config: {
                      "key": null,
                    },
                    connectTimeout: 1,
                    envoyDNSDiscoveryType: "envoyDNSDiscoveryType_example",
                    envoyGatewayBindAddresses: {
                      "key": {
                        address: "address_example",
                        name: "name_example",
                        port: 1,
                      },
                    },
                    envoyGatewayBindTaggedAddresses: true,
                    envoyGatewayNoDefaultBind: true,
                  },
                  terminating: {
                    services: [
                      {
                        cAFile: "cAFile_example",
                        certFile: "certFile_example",
                        keyFile: "keyFile_example",
                        name: "name_example",
                        SNI: "SNI_example",
                      },
                    ],
                  },
                },
                _native: true,
                sidecarService: {
                  disableDefaultTCPCheck: true,
                  port: "port_example",
                  proxy: {
                    config: {
                      "key": null,
                    },
                    exposeConfig: {
                      path: [
                        {
                          listenerPort: "listenerPort_example",
                          localPathPort: 1,
                          path: "path_example",
                          protocol: "protocol_example",
                        },
                      ],
                    },
                    localServiceAddress: "localServiceAddress_example",
                    localServicePort: 1,
                    upstreams: [
                      {
                        datacenter: "datacenter_example",
                        destinationName: "destinationName_example",
                        localBindAddress: "localBindAddress_example",
                        localBindPort: 1,
                        meshGateway: {
                          mode: "mode_example",
                        },
                      },
                    ],
                  },
                  tags: [
                    "tags_example",
                  ],
                },
                sidecarTask: {
                  config: {
                    "key": null,
                  },
                  driver: "driver_example",
                  env: {
                    "key": "key_example",
                  },
                  killSignal: "killSignal_example",
                  killTimeout: 1,
                  logConfig: {
                    maxFileSizeMB: 1,
                    maxFiles: 1,
                  },
                  meta: {
                    "key": "key_example",
                  },
                  name: "name_example",
                  resources: {
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
                  shutdownDelay: 1,
                  user: "user_example",
                },
              },
              enableTagOverride: true,
              meta: {
                "key": "key_example",
              },
              name: "name_example",
              onUpdate: "onUpdate_example",
              portLabel: "portLabel_example",
              provider: "provider_example",
              tags: [
                "tags_example",
              ],
              taskName: "taskName_example",
            },
          ],
          shutdownDelay: 1,
          spreads: [
            {
              attribute: "attribute_example",
              spreadTarget: [
                {
                  percent: 0,
                  value: "value_example",
                },
              ],
              weight: -128,
            },
          ],
          stopAfterClientDisconnect: 1,
          tasks: [
            {
              affinities: [
                {
                  lTarget: "lTarget_example",
                  operand: "operand_example",
                  rTarget: "rTarget_example",
                  weight: -128,
                },
              ],
              artifacts: [
                {
                  getterHeaders: {
                    "key": "key_example",
                  },
                  getterMode: "getterMode_example",
                  getterOptions: {
                    "key": "key_example",
                  },
                  getterSource: "getterSource_example",
                  relativeDest: "relativeDest_example",
                },
              ],
              cSIPluginConfig: {
                ID: "ID_example",
                mountDir: "mountDir_example",
                type: "type_example",
              },
              config: {
                "key": null,
              },
              constraints: [
                {
                  lTarget: "lTarget_example",
                  operand: "operand_example",
                  rTarget: "rTarget_example",
                },
              ],
              dispatchPayload: {
                file: "file_example",
              },
              driver: "driver_example",
              env: {
                "key": "key_example",
              },
              killSignal: "killSignal_example",
              killTimeout: 1,
              kind: "kind_example",
              leader: true,
              lifecycle: {
                hook: "hook_example",
                sidecar: true,
              },
              logConfig: {
                maxFileSizeMB: 1,
                maxFiles: 1,
              },
              meta: {
                "key": "key_example",
              },
              name: "name_example",
              resources: {
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
              restartPolicy: {
                attempts: 1,
                delay: 1,
                interval: 1,
                mode: "mode_example",
              },
              scalingPolicies: [
                {
                  createIndex: 0,
                  enabled: true,
                  ID: "ID_example",
                  max: 1,
                  min: 1,
                  modifyIndex: 0,
                  namespace: "namespace_example",
                  policy: {
                    "key": null,
                  },
                  target: {
                    "key": "key_example",
                  },
                  type: "type_example",
                },
              ],
              services: [
                {
                  address: "address_example",
                  addressMode: "addressMode_example",
                  canaryMeta: {
                    "key": "key_example",
                  },
                  canaryTags: [
                    "canaryTags_example",
                  ],
                  checkRestart: {
                    grace: 1,
                    ignoreWarnings: true,
                    limit: 1,
                  },
                  checks: [
                    {
                      addressMode: "addressMode_example",
                      advertise: "advertise_example",
                      args: [
                        "args_example",
                      ],
                      body: "body_example",
                      checkRestart: {
                        grace: 1,
                        ignoreWarnings: true,
                        limit: 1,
                      },
                      command: "command_example",
                      expose: true,
                      failuresBeforeCritical: 1,
                      gRPCService: "gRPCService_example",
                      gRPCUseTLS: true,
                      header: {
                        "key": [
                          "key_example",
                        ],
                      },
                      initialStatus: "initialStatus_example",
                      interval: 1,
                      method: "method_example",
                      name: "name_example",
                      onUpdate: "onUpdate_example",
                      path: "path_example",
                      portLabel: "portLabel_example",
                      protocol: "protocol_example",
                      successBeforePassing: 1,
                      tLSSkipVerify: true,
                      taskName: "taskName_example",
                      timeout: 1,
                      type: "type_example",
                    },
                  ],
                  connect: {
                    gateway: {
                      ingress: {
                        listeners: [
                          {
                            port: 1,
                            protocol: "protocol_example",
                            services: [
                              {
                                hosts: [
                                  "hosts_example",
                                ],
                                name: "name_example",
                              },
                            ],
                          },
                        ],
                        TLS: {
                          enabled: true,
                        },
                      },
                      mesh: null,
                      proxy: {
                        config: {
                          "key": null,
                        },
                        connectTimeout: 1,
                        envoyDNSDiscoveryType: "envoyDNSDiscoveryType_example",
                        envoyGatewayBindAddresses: {
                          "key": {
                            address: "address_example",
                            name: "name_example",
                            port: 1,
                          },
                        },
                        envoyGatewayBindTaggedAddresses: true,
                        envoyGatewayNoDefaultBind: true,
                      },
                      terminating: {
                        services: [
                          {
                            cAFile: "cAFile_example",
                            certFile: "certFile_example",
                            keyFile: "keyFile_example",
                            name: "name_example",
                            SNI: "SNI_example",
                          },
                        ],
                      },
                    },
                    _native: true,
                    sidecarService: {
                      disableDefaultTCPCheck: true,
                      port: "port_example",
                      proxy: {
                        config: {
                          "key": null,
                        },
                        exposeConfig: {
                          path: [
                            {
                              listenerPort: "listenerPort_example",
                              localPathPort: 1,
                              path: "path_example",
                              protocol: "protocol_example",
                            },
                          ],
                        },
                        localServiceAddress: "localServiceAddress_example",
                        localServicePort: 1,
                        upstreams: [
                          {
                            datacenter: "datacenter_example",
                            destinationName: "destinationName_example",
                            localBindAddress: "localBindAddress_example",
                            localBindPort: 1,
                            meshGateway: {
                              mode: "mode_example",
                            },
                          },
                        ],
                      },
                      tags: [
                        "tags_example",
                      ],
                    },
                    sidecarTask: {
                      config: {
                        "key": null,
                      },
                      driver: "driver_example",
                      env: {
                        "key": "key_example",
                      },
                      killSignal: "killSignal_example",
                      killTimeout: 1,
                      logConfig: {
                        maxFileSizeMB: 1,
                        maxFiles: 1,
                      },
                      meta: {
                        "key": "key_example",
                      },
                      name: "name_example",
                      resources: {
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
                      shutdownDelay: 1,
                      user: "user_example",
                    },
                  },
                  enableTagOverride: true,
                  meta: {
                    "key": "key_example",
                  },
                  name: "name_example",
                  onUpdate: "onUpdate_example",
                  portLabel: "portLabel_example",
                  provider: "provider_example",
                  tags: [
                    "tags_example",
                  ],
                  taskName: "taskName_example",
                },
              ],
              shutdownDelay: 1,
              templates: [
                {
                  changeMode: "changeMode_example",
                  changeSignal: "changeSignal_example",
                  destPath: "destPath_example",
                  embeddedTmpl: "embeddedTmpl_example",
                  envvars: true,
                  leftDelim: "leftDelim_example",
                  perms: "perms_example",
                  rightDelim: "rightDelim_example",
                  sourcePath: "sourcePath_example",
                  splay: 1,
                  vaultGrace: 1,
                  wait: {
                    max: 1,
                    min: 1,
                  },
                },
              ],
              user: "user_example",
              vault: {
                changeMode: "changeMode_example",
                changeSignal: "changeSignal_example",
                env: true,
                namespace: "namespace_example",
                policies: [
                  "policies_example",
                ],
              },
              volumeMounts: [
                {
                  destination: "destination_example",
                  propagationMode: "propagationMode_example",
                  readOnly: true,
                  volume: "volume_example",
                },
              ],
            },
          ],
          update: {
            autoPromote: true,
            autoRevert: true,
            canary: 1,
            healthCheck: "healthCheck_example",
            healthyDeadline: 1,
            maxParallel: 1,
            minHealthyTime: 1,
            progressDeadline: 1,
            stagger: 1,
          },
          volumes: {
            "key": {
              accessMode: "accessMode_example",
              attachmentMode: "attachmentMode_example",
              mountOptions: {
                fSType: "fSType_example",
                mountFlags: [
                  "mountFlags_example",
                ],
              },
              name: "name_example",
              perAlloc: true,
              readOnly: true,
              source: "source_example",
              type: "type_example",
            },
          },
        },
      ],
      type: "type_example",
      update: {
        autoPromote: true,
        autoRevert: true,
        canary: 1,
        healthCheck: "healthCheck_example",
        healthyDeadline: 1,
        maxParallel: 1,
        minHealthyTime: 1,
        progressDeadline: 1,
        stagger: 1,
      },
      vaultNamespace: "vaultNamespace_example",
      vaultToken: "vaultToken_example",
      version: 0,
    },
    jobModifyIndex: 0,
    namespace: "namespace_example",
    policyOverride: true,
    preserveCounts: true,
    region: "region_example",
    secretID: "secretID_example",
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

apiInstance.registerJob(body).then((data:any) => {
  console.log('API called successfully. Returned data: ' + data);
}).catch((error:any) => console.error(error));
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jobRegisterRequest** | **JobRegisterRequest**|  |
 **region** | [**string**] | Filters results based on the specified region. | (optional) defaults to undefined
 **namespace** | [**string**] | Filters results based on the specified namespace. | (optional) defaults to undefined
 **xNomadToken** | [**string**] | A Nomad ACL token. | (optional) defaults to undefined
 **idempotencyToken** | [**string**] | Can be used to ensure operations are only run once. | (optional) defaults to undefined


### Return type

**JobRegisterResponse**

### Authorization

[X-Nomad-Token](README.md#X-Nomad-Token)

### HTTP request headers

 - **Content-Type**: application/json
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


