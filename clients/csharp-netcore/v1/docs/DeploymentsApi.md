# Nomad.Client.Api.DeploymentsApi

All URIs are relative to *https://127.0.0.1:4646/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDeployment**](DeploymentsApi.md#getdeployment) | **GET** /deployment/{deploymentID} | 
[**GetDeploymentAllocations**](DeploymentsApi.md#getdeploymentallocations) | **GET** /deployment/allocations/{deploymentID} | 
[**GetDeployments**](DeploymentsApi.md#getdeployments) | **GET** /deployments | 
[**PostDeploymentAllocationHealth**](DeploymentsApi.md#postdeploymentallocationhealth) | **POST** /deployment/allocation-health/{deploymentID} | 
[**PostDeploymentFail**](DeploymentsApi.md#postdeploymentfail) | **POST** /deployment/fail/{deploymentID} | 
[**PostDeploymentPause**](DeploymentsApi.md#postdeploymentpause) | **POST** /deployment/pause/{deploymentID} | 
[**PostDeploymentPromote**](DeploymentsApi.md#postdeploymentpromote) | **POST** /deployment/promote/{deploymentID} | 
[**PostDeploymentUnblock**](DeploymentsApi.md#postdeploymentunblock) | **POST** /deployment/unblock/{deploymentID} | 


<a name="getdeployment"></a>
# **GetDeployment**
> Deployment GetDeployment (string deploymentID, string region = null, string _namespace = null, int? index = null, string wait = null, string stale = null, string prefix = null, string xNomadToken = null, int? perPage = null, string nextToken = null)



### Example
```csharp
using System.Collections.Generic;
using System.Diagnostics;
using Nomad.Client.Api;
using Nomad.Client.Client;
using Nomad.Client.Model;

namespace Example
{
    public class GetDeploymentExample
    {
        public static void Main()
        {
            Configuration config = new Configuration();
            config.BasePath = "https://127.0.0.1:4646/v1";
            // Configure API key authorization: X-Nomad-Token
            config.AddApiKey("X-Nomad-Token", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // config.AddApiKeyPrefix("X-Nomad-Token", "Bearer");

            var apiInstance = new DeploymentsApi(config);
            var deploymentID = deploymentID_example;  // string | Deployment ID.
            var region = region_example;  // string | Filters results based on the specified region. (optional) 
            var _namespace = _namespace_example;  // string | Filters results based on the specified namespace. (optional) 
            var index = 56;  // int? | If set, wait until query exceeds given index. Must be provided with WaitParam. (optional) 
            var wait = wait_example;  // string | Provided with IndexParam to wait for change. (optional) 
            var stale = stale_example;  // string | If present, results will include stale reads. (optional) 
            var prefix = prefix_example;  // string | Constrains results to jobs that start with the defined prefix (optional) 
            var xNomadToken = xNomadToken_example;  // string | A Nomad ACL token. (optional) 
            var perPage = 56;  // int? | Maximum number of results to return. (optional) 
            var nextToken = nextToken_example;  // string | Indicates where to start paging for queries that support pagination. (optional) 

            try
            {
                Deployment result = apiInstance.GetDeployment(deploymentID, region, _namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling DeploymentsApi.GetDeployment: " + e.Message );
                Debug.Print("Status Code: "+ e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deploymentID** | **string**| Deployment ID. | 
 **region** | **string**| Filters results based on the specified region. | [optional] 
 **_namespace** | **string**| Filters results based on the specified namespace. | [optional] 
 **index** | **int?**| If set, wait until query exceeds given index. Must be provided with WaitParam. | [optional] 
 **wait** | **string**| Provided with IndexParam to wait for change. | [optional] 
 **stale** | **string**| If present, results will include stale reads. | [optional] 
 **prefix** | **string**| Constrains results to jobs that start with the defined prefix | [optional] 
 **xNomadToken** | **string**| A Nomad ACL token. | [optional] 
 **perPage** | **int?**| Maximum number of results to return. | [optional] 
 **nextToken** | **string**| Indicates where to start paging for queries that support pagination. | [optional] 

### Return type

[**Deployment**](Deployment.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** |  |  * X-Nomad-Index - A unique identifier representing the current state of the requested resource. On a new Nomad cluster the value of this index starts at 1. <br>  * X-Nomad-KnownLeader - Boolean indicating if there is a known cluster leader. <br>  * X-Nomad-LastContact - The time in milliseconds that a server was last contacted by the leader node. <br>  |
| **400** | Bad request |  -  |
| **403** | Forbidden |  -  |
| **405** | Method not allowed |  -  |
| **500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="getdeploymentallocations"></a>
# **GetDeploymentAllocations**
> List&lt;AllocationListStub&gt; GetDeploymentAllocations (string deploymentID, string region = null, string _namespace = null, int? index = null, string wait = null, string stale = null, string prefix = null, string xNomadToken = null, int? perPage = null, string nextToken = null)



### Example
```csharp
using System.Collections.Generic;
using System.Diagnostics;
using Nomad.Client.Api;
using Nomad.Client.Client;
using Nomad.Client.Model;

namespace Example
{
    public class GetDeploymentAllocationsExample
    {
        public static void Main()
        {
            Configuration config = new Configuration();
            config.BasePath = "https://127.0.0.1:4646/v1";
            // Configure API key authorization: X-Nomad-Token
            config.AddApiKey("X-Nomad-Token", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // config.AddApiKeyPrefix("X-Nomad-Token", "Bearer");

            var apiInstance = new DeploymentsApi(config);
            var deploymentID = deploymentID_example;  // string | Deployment ID.
            var region = region_example;  // string | Filters results based on the specified region. (optional) 
            var _namespace = _namespace_example;  // string | Filters results based on the specified namespace. (optional) 
            var index = 56;  // int? | If set, wait until query exceeds given index. Must be provided with WaitParam. (optional) 
            var wait = wait_example;  // string | Provided with IndexParam to wait for change. (optional) 
            var stale = stale_example;  // string | If present, results will include stale reads. (optional) 
            var prefix = prefix_example;  // string | Constrains results to jobs that start with the defined prefix (optional) 
            var xNomadToken = xNomadToken_example;  // string | A Nomad ACL token. (optional) 
            var perPage = 56;  // int? | Maximum number of results to return. (optional) 
            var nextToken = nextToken_example;  // string | Indicates where to start paging for queries that support pagination. (optional) 

            try
            {
                List<AllocationListStub> result = apiInstance.GetDeploymentAllocations(deploymentID, region, _namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling DeploymentsApi.GetDeploymentAllocations: " + e.Message );
                Debug.Print("Status Code: "+ e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deploymentID** | **string**| Deployment ID. | 
 **region** | **string**| Filters results based on the specified region. | [optional] 
 **_namespace** | **string**| Filters results based on the specified namespace. | [optional] 
 **index** | **int?**| If set, wait until query exceeds given index. Must be provided with WaitParam. | [optional] 
 **wait** | **string**| Provided with IndexParam to wait for change. | [optional] 
 **stale** | **string**| If present, results will include stale reads. | [optional] 
 **prefix** | **string**| Constrains results to jobs that start with the defined prefix | [optional] 
 **xNomadToken** | **string**| A Nomad ACL token. | [optional] 
 **perPage** | **int?**| Maximum number of results to return. | [optional] 
 **nextToken** | **string**| Indicates where to start paging for queries that support pagination. | [optional] 

### Return type

[**List&lt;AllocationListStub&gt;**](AllocationListStub.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** |  |  * X-Nomad-Index - A unique identifier representing the current state of the requested resource. On a new Nomad cluster the value of this index starts at 1. <br>  * X-Nomad-KnownLeader - Boolean indicating if there is a known cluster leader. <br>  * X-Nomad-LastContact - The time in milliseconds that a server was last contacted by the leader node. <br>  |
| **400** | Bad request |  -  |
| **403** | Forbidden |  -  |
| **405** | Method not allowed |  -  |
| **500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="getdeployments"></a>
# **GetDeployments**
> List&lt;Deployment&gt; GetDeployments (string region = null, string _namespace = null, int? index = null, string wait = null, string stale = null, string prefix = null, string xNomadToken = null, int? perPage = null, string nextToken = null)



### Example
```csharp
using System.Collections.Generic;
using System.Diagnostics;
using Nomad.Client.Api;
using Nomad.Client.Client;
using Nomad.Client.Model;

namespace Example
{
    public class GetDeploymentsExample
    {
        public static void Main()
        {
            Configuration config = new Configuration();
            config.BasePath = "https://127.0.0.1:4646/v1";
            // Configure API key authorization: X-Nomad-Token
            config.AddApiKey("X-Nomad-Token", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // config.AddApiKeyPrefix("X-Nomad-Token", "Bearer");

            var apiInstance = new DeploymentsApi(config);
            var region = region_example;  // string | Filters results based on the specified region. (optional) 
            var _namespace = _namespace_example;  // string | Filters results based on the specified namespace. (optional) 
            var index = 56;  // int? | If set, wait until query exceeds given index. Must be provided with WaitParam. (optional) 
            var wait = wait_example;  // string | Provided with IndexParam to wait for change. (optional) 
            var stale = stale_example;  // string | If present, results will include stale reads. (optional) 
            var prefix = prefix_example;  // string | Constrains results to jobs that start with the defined prefix (optional) 
            var xNomadToken = xNomadToken_example;  // string | A Nomad ACL token. (optional) 
            var perPage = 56;  // int? | Maximum number of results to return. (optional) 
            var nextToken = nextToken_example;  // string | Indicates where to start paging for queries that support pagination. (optional) 

            try
            {
                List<Deployment> result = apiInstance.GetDeployments(region, _namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling DeploymentsApi.GetDeployments: " + e.Message );
                Debug.Print("Status Code: "+ e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **region** | **string**| Filters results based on the specified region. | [optional] 
 **_namespace** | **string**| Filters results based on the specified namespace. | [optional] 
 **index** | **int?**| If set, wait until query exceeds given index. Must be provided with WaitParam. | [optional] 
 **wait** | **string**| Provided with IndexParam to wait for change. | [optional] 
 **stale** | **string**| If present, results will include stale reads. | [optional] 
 **prefix** | **string**| Constrains results to jobs that start with the defined prefix | [optional] 
 **xNomadToken** | **string**| A Nomad ACL token. | [optional] 
 **perPage** | **int?**| Maximum number of results to return. | [optional] 
 **nextToken** | **string**| Indicates where to start paging for queries that support pagination. | [optional] 

### Return type

[**List&lt;Deployment&gt;**](Deployment.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** |  |  * X-Nomad-Index - A unique identifier representing the current state of the requested resource. On a new Nomad cluster the value of this index starts at 1. <br>  * X-Nomad-KnownLeader - Boolean indicating if there is a known cluster leader. <br>  * X-Nomad-LastContact - The time in milliseconds that a server was last contacted by the leader node. <br>  |
| **400** | Bad request |  -  |
| **403** | Forbidden |  -  |
| **405** | Method not allowed |  -  |
| **500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="postdeploymentallocationhealth"></a>
# **PostDeploymentAllocationHealth**
> DeploymentUpdateResponse PostDeploymentAllocationHealth (string deploymentID, DeploymentAllocHealthRequest deploymentAllocHealthRequest, string region = null, string _namespace = null, string xNomadToken = null, string idempotencyToken = null)



### Example
```csharp
using System.Collections.Generic;
using System.Diagnostics;
using Nomad.Client.Api;
using Nomad.Client.Client;
using Nomad.Client.Model;

namespace Example
{
    public class PostDeploymentAllocationHealthExample
    {
        public static void Main()
        {
            Configuration config = new Configuration();
            config.BasePath = "https://127.0.0.1:4646/v1";
            // Configure API key authorization: X-Nomad-Token
            config.AddApiKey("X-Nomad-Token", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // config.AddApiKeyPrefix("X-Nomad-Token", "Bearer");

            var apiInstance = new DeploymentsApi(config);
            var deploymentID = deploymentID_example;  // string | Deployment ID.
            var deploymentAllocHealthRequest = new DeploymentAllocHealthRequest(); // DeploymentAllocHealthRequest | 
            var region = region_example;  // string | Filters results based on the specified region. (optional) 
            var _namespace = _namespace_example;  // string | Filters results based on the specified namespace. (optional) 
            var xNomadToken = xNomadToken_example;  // string | A Nomad ACL token. (optional) 
            var idempotencyToken = idempotencyToken_example;  // string | Can be used to ensure operations are only run once. (optional) 

            try
            {
                DeploymentUpdateResponse result = apiInstance.PostDeploymentAllocationHealth(deploymentID, deploymentAllocHealthRequest, region, _namespace, xNomadToken, idempotencyToken);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling DeploymentsApi.PostDeploymentAllocationHealth: " + e.Message );
                Debug.Print("Status Code: "+ e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deploymentID** | **string**| Deployment ID. | 
 **deploymentAllocHealthRequest** | [**DeploymentAllocHealthRequest**](DeploymentAllocHealthRequest.md)|  | 
 **region** | **string**| Filters results based on the specified region. | [optional] 
 **_namespace** | **string**| Filters results based on the specified namespace. | [optional] 
 **xNomadToken** | **string**| A Nomad ACL token. | [optional] 
 **idempotencyToken** | **string**| Can be used to ensure operations are only run once. | [optional] 

### Return type

[**DeploymentUpdateResponse**](DeploymentUpdateResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** |  |  -  |
| **400** | Bad request |  -  |
| **403** | Forbidden |  -  |
| **405** | Method not allowed |  -  |
| **500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="postdeploymentfail"></a>
# **PostDeploymentFail**
> DeploymentUpdateResponse PostDeploymentFail (string deploymentID, string region = null, string _namespace = null, string xNomadToken = null, string idempotencyToken = null)



### Example
```csharp
using System.Collections.Generic;
using System.Diagnostics;
using Nomad.Client.Api;
using Nomad.Client.Client;
using Nomad.Client.Model;

namespace Example
{
    public class PostDeploymentFailExample
    {
        public static void Main()
        {
            Configuration config = new Configuration();
            config.BasePath = "https://127.0.0.1:4646/v1";
            // Configure API key authorization: X-Nomad-Token
            config.AddApiKey("X-Nomad-Token", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // config.AddApiKeyPrefix("X-Nomad-Token", "Bearer");

            var apiInstance = new DeploymentsApi(config);
            var deploymentID = deploymentID_example;  // string | Deployment ID.
            var region = region_example;  // string | Filters results based on the specified region. (optional) 
            var _namespace = _namespace_example;  // string | Filters results based on the specified namespace. (optional) 
            var xNomadToken = xNomadToken_example;  // string | A Nomad ACL token. (optional) 
            var idempotencyToken = idempotencyToken_example;  // string | Can be used to ensure operations are only run once. (optional) 

            try
            {
                DeploymentUpdateResponse result = apiInstance.PostDeploymentFail(deploymentID, region, _namespace, xNomadToken, idempotencyToken);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling DeploymentsApi.PostDeploymentFail: " + e.Message );
                Debug.Print("Status Code: "+ e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deploymentID** | **string**| Deployment ID. | 
 **region** | **string**| Filters results based on the specified region. | [optional] 
 **_namespace** | **string**| Filters results based on the specified namespace. | [optional] 
 **xNomadToken** | **string**| A Nomad ACL token. | [optional] 
 **idempotencyToken** | **string**| Can be used to ensure operations are only run once. | [optional] 

### Return type

[**DeploymentUpdateResponse**](DeploymentUpdateResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** |  |  -  |
| **400** | Bad request |  -  |
| **403** | Forbidden |  -  |
| **405** | Method not allowed |  -  |
| **500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="postdeploymentpause"></a>
# **PostDeploymentPause**
> DeploymentUpdateResponse PostDeploymentPause (string deploymentID, DeploymentPauseRequest deploymentPauseRequest, string region = null, string _namespace = null, string xNomadToken = null, string idempotencyToken = null)



### Example
```csharp
using System.Collections.Generic;
using System.Diagnostics;
using Nomad.Client.Api;
using Nomad.Client.Client;
using Nomad.Client.Model;

namespace Example
{
    public class PostDeploymentPauseExample
    {
        public static void Main()
        {
            Configuration config = new Configuration();
            config.BasePath = "https://127.0.0.1:4646/v1";
            // Configure API key authorization: X-Nomad-Token
            config.AddApiKey("X-Nomad-Token", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // config.AddApiKeyPrefix("X-Nomad-Token", "Bearer");

            var apiInstance = new DeploymentsApi(config);
            var deploymentID = deploymentID_example;  // string | Deployment ID.
            var deploymentPauseRequest = new DeploymentPauseRequest(); // DeploymentPauseRequest | 
            var region = region_example;  // string | Filters results based on the specified region. (optional) 
            var _namespace = _namespace_example;  // string | Filters results based on the specified namespace. (optional) 
            var xNomadToken = xNomadToken_example;  // string | A Nomad ACL token. (optional) 
            var idempotencyToken = idempotencyToken_example;  // string | Can be used to ensure operations are only run once. (optional) 

            try
            {
                DeploymentUpdateResponse result = apiInstance.PostDeploymentPause(deploymentID, deploymentPauseRequest, region, _namespace, xNomadToken, idempotencyToken);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling DeploymentsApi.PostDeploymentPause: " + e.Message );
                Debug.Print("Status Code: "+ e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deploymentID** | **string**| Deployment ID. | 
 **deploymentPauseRequest** | [**DeploymentPauseRequest**](DeploymentPauseRequest.md)|  | 
 **region** | **string**| Filters results based on the specified region. | [optional] 
 **_namespace** | **string**| Filters results based on the specified namespace. | [optional] 
 **xNomadToken** | **string**| A Nomad ACL token. | [optional] 
 **idempotencyToken** | **string**| Can be used to ensure operations are only run once. | [optional] 

### Return type

[**DeploymentUpdateResponse**](DeploymentUpdateResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** |  |  -  |
| **400** | Bad request |  -  |
| **403** | Forbidden |  -  |
| **405** | Method not allowed |  -  |
| **500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="postdeploymentpromote"></a>
# **PostDeploymentPromote**
> DeploymentUpdateResponse PostDeploymentPromote (string deploymentID, DeploymentPromoteRequest deploymentPromoteRequest, string region = null, string _namespace = null, string xNomadToken = null, string idempotencyToken = null)



### Example
```csharp
using System.Collections.Generic;
using System.Diagnostics;
using Nomad.Client.Api;
using Nomad.Client.Client;
using Nomad.Client.Model;

namespace Example
{
    public class PostDeploymentPromoteExample
    {
        public static void Main()
        {
            Configuration config = new Configuration();
            config.BasePath = "https://127.0.0.1:4646/v1";
            // Configure API key authorization: X-Nomad-Token
            config.AddApiKey("X-Nomad-Token", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // config.AddApiKeyPrefix("X-Nomad-Token", "Bearer");

            var apiInstance = new DeploymentsApi(config);
            var deploymentID = deploymentID_example;  // string | Deployment ID.
            var deploymentPromoteRequest = new DeploymentPromoteRequest(); // DeploymentPromoteRequest | 
            var region = region_example;  // string | Filters results based on the specified region. (optional) 
            var _namespace = _namespace_example;  // string | Filters results based on the specified namespace. (optional) 
            var xNomadToken = xNomadToken_example;  // string | A Nomad ACL token. (optional) 
            var idempotencyToken = idempotencyToken_example;  // string | Can be used to ensure operations are only run once. (optional) 

            try
            {
                DeploymentUpdateResponse result = apiInstance.PostDeploymentPromote(deploymentID, deploymentPromoteRequest, region, _namespace, xNomadToken, idempotencyToken);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling DeploymentsApi.PostDeploymentPromote: " + e.Message );
                Debug.Print("Status Code: "+ e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deploymentID** | **string**| Deployment ID. | 
 **deploymentPromoteRequest** | [**DeploymentPromoteRequest**](DeploymentPromoteRequest.md)|  | 
 **region** | **string**| Filters results based on the specified region. | [optional] 
 **_namespace** | **string**| Filters results based on the specified namespace. | [optional] 
 **xNomadToken** | **string**| A Nomad ACL token. | [optional] 
 **idempotencyToken** | **string**| Can be used to ensure operations are only run once. | [optional] 

### Return type

[**DeploymentUpdateResponse**](DeploymentUpdateResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** |  |  -  |
| **400** | Bad request |  -  |
| **403** | Forbidden |  -  |
| **405** | Method not allowed |  -  |
| **500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="postdeploymentunblock"></a>
# **PostDeploymentUnblock**
> DeploymentUpdateResponse PostDeploymentUnblock (string deploymentID, DeploymentUnblockRequest deploymentUnblockRequest, string region = null, string _namespace = null, string xNomadToken = null, string idempotencyToken = null)



### Example
```csharp
using System.Collections.Generic;
using System.Diagnostics;
using Nomad.Client.Api;
using Nomad.Client.Client;
using Nomad.Client.Model;

namespace Example
{
    public class PostDeploymentUnblockExample
    {
        public static void Main()
        {
            Configuration config = new Configuration();
            config.BasePath = "https://127.0.0.1:4646/v1";
            // Configure API key authorization: X-Nomad-Token
            config.AddApiKey("X-Nomad-Token", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // config.AddApiKeyPrefix("X-Nomad-Token", "Bearer");

            var apiInstance = new DeploymentsApi(config);
            var deploymentID = deploymentID_example;  // string | Deployment ID.
            var deploymentUnblockRequest = new DeploymentUnblockRequest(); // DeploymentUnblockRequest | 
            var region = region_example;  // string | Filters results based on the specified region. (optional) 
            var _namespace = _namespace_example;  // string | Filters results based on the specified namespace. (optional) 
            var xNomadToken = xNomadToken_example;  // string | A Nomad ACL token. (optional) 
            var idempotencyToken = idempotencyToken_example;  // string | Can be used to ensure operations are only run once. (optional) 

            try
            {
                DeploymentUpdateResponse result = apiInstance.PostDeploymentUnblock(deploymentID, deploymentUnblockRequest, region, _namespace, xNomadToken, idempotencyToken);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling DeploymentsApi.PostDeploymentUnblock: " + e.Message );
                Debug.Print("Status Code: "+ e.ErrorCode);
                Debug.Print(e.StackTrace);
            }
        }
    }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deploymentID** | **string**| Deployment ID. | 
 **deploymentUnblockRequest** | [**DeploymentUnblockRequest**](DeploymentUnblockRequest.md)|  | 
 **region** | **string**| Filters results based on the specified region. | [optional] 
 **_namespace** | **string**| Filters results based on the specified namespace. | [optional] 
 **xNomadToken** | **string**| A Nomad ACL token. | [optional] 
 **idempotencyToken** | **string**| Can be used to ensure operations are only run once. | [optional] 

### Return type

[**DeploymentUpdateResponse**](DeploymentUpdateResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** |  |  -  |
| **400** | Bad request |  -  |
| **403** | Forbidden |  -  |
| **405** | Method not allowed |  -  |
| **500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

