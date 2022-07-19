# Nomad.Client.Api.AllocationsApi

All URIs are relative to *https://127.0.0.1:4646/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllocation**](AllocationsApi.md#getallocation) | **GET** /allocation/{allocID} | 
[**GetAllocationServices**](AllocationsApi.md#getallocationservices) | **GET** /allocation/{allocID}/services | 
[**GetAllocations**](AllocationsApi.md#getallocations) | **GET** /allocations | 
[**PostAllocationStop**](AllocationsApi.md#postallocationstop) | **POST** /allocation/{allocID}/stop | 


<a name="getallocation"></a>
# **GetAllocation**
> Allocation GetAllocation (string allocID, string region = null, string _namespace = null, int? index = null, string wait = null, string stale = null, string prefix = null, string xNomadToken = null, int? perPage = null, string nextToken = null)



### Example
```csharp
using System.Collections.Generic;
using System.Diagnostics;
using Nomad.Client.Api;
using Nomad.Client.Client;
using Nomad.Client.Model;

namespace Example
{
    public class GetAllocationExample
    {
        public static void Main()
        {
            Configuration config = new Configuration();
            config.BasePath = "https://127.0.0.1:4646/v1";
            // Configure API key authorization: X-Nomad-Token
            config.AddApiKey("X-Nomad-Token", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // config.AddApiKeyPrefix("X-Nomad-Token", "Bearer");

            var apiInstance = new AllocationsApi(config);
            var allocID = allocID_example;  // string | Allocation ID.
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
                Allocation result = apiInstance.GetAllocation(allocID, region, _namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling AllocationsApi.GetAllocation: " + e.Message );
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
 **allocID** | **string**| Allocation ID. | 
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

[**Allocation**](Allocation.md)

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

<a name="getallocationservices"></a>
# **GetAllocationServices**
> List&lt;ServiceRegistration&gt; GetAllocationServices (string allocID, string region = null, string _namespace = null, int? index = null, string wait = null, string stale = null, string prefix = null, string xNomadToken = null, int? perPage = null, string nextToken = null)



### Example
```csharp
using System.Collections.Generic;
using System.Diagnostics;
using Nomad.Client.Api;
using Nomad.Client.Client;
using Nomad.Client.Model;

namespace Example
{
    public class GetAllocationServicesExample
    {
        public static void Main()
        {
            Configuration config = new Configuration();
            config.BasePath = "https://127.0.0.1:4646/v1";
            // Configure API key authorization: X-Nomad-Token
            config.AddApiKey("X-Nomad-Token", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // config.AddApiKeyPrefix("X-Nomad-Token", "Bearer");

            var apiInstance = new AllocationsApi(config);
            var allocID = allocID_example;  // string | Allocation ID.
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
                List<ServiceRegistration> result = apiInstance.GetAllocationServices(allocID, region, _namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling AllocationsApi.GetAllocationServices: " + e.Message );
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
 **allocID** | **string**| Allocation ID. | 
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

[**List&lt;ServiceRegistration&gt;**](ServiceRegistration.md)

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

<a name="getallocations"></a>
# **GetAllocations**
> List&lt;AllocationListStub&gt; GetAllocations (string region = null, string _namespace = null, int? index = null, string wait = null, string stale = null, string prefix = null, string xNomadToken = null, int? perPage = null, string nextToken = null, bool? resources = null, bool? taskStates = null)



### Example
```csharp
using System.Collections.Generic;
using System.Diagnostics;
using Nomad.Client.Api;
using Nomad.Client.Client;
using Nomad.Client.Model;

namespace Example
{
    public class GetAllocationsExample
    {
        public static void Main()
        {
            Configuration config = new Configuration();
            config.BasePath = "https://127.0.0.1:4646/v1";
            // Configure API key authorization: X-Nomad-Token
            config.AddApiKey("X-Nomad-Token", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // config.AddApiKeyPrefix("X-Nomad-Token", "Bearer");

            var apiInstance = new AllocationsApi(config);
            var region = region_example;  // string | Filters results based on the specified region. (optional) 
            var _namespace = _namespace_example;  // string | Filters results based on the specified namespace. (optional) 
            var index = 56;  // int? | If set, wait until query exceeds given index. Must be provided with WaitParam. (optional) 
            var wait = wait_example;  // string | Provided with IndexParam to wait for change. (optional) 
            var stale = stale_example;  // string | If present, results will include stale reads. (optional) 
            var prefix = prefix_example;  // string | Constrains results to jobs that start with the defined prefix (optional) 
            var xNomadToken = xNomadToken_example;  // string | A Nomad ACL token. (optional) 
            var perPage = 56;  // int? | Maximum number of results to return. (optional) 
            var nextToken = nextToken_example;  // string | Indicates where to start paging for queries that support pagination. (optional) 
            var resources = true;  // bool? | Flag indicating whether to include resources in response. (optional) 
            var taskStates = true;  // bool? | Flag indicating whether to include task states in response. (optional) 

            try
            {
                List<AllocationListStub> result = apiInstance.GetAllocations(region, _namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken, resources, taskStates);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling AllocationsApi.GetAllocations: " + e.Message );
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
 **resources** | **bool?**| Flag indicating whether to include resources in response. | [optional] 
 **taskStates** | **bool?**| Flag indicating whether to include task states in response. | [optional] 

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

<a name="postallocationstop"></a>
# **PostAllocationStop**
> AllocStopResponse PostAllocationStop (string allocID, string region = null, string _namespace = null, int? index = null, string wait = null, string stale = null, string prefix = null, string xNomadToken = null, int? perPage = null, string nextToken = null, bool? noShutdownDelay = null)



### Example
```csharp
using System.Collections.Generic;
using System.Diagnostics;
using Nomad.Client.Api;
using Nomad.Client.Client;
using Nomad.Client.Model;

namespace Example
{
    public class PostAllocationStopExample
    {
        public static void Main()
        {
            Configuration config = new Configuration();
            config.BasePath = "https://127.0.0.1:4646/v1";
            // Configure API key authorization: X-Nomad-Token
            config.AddApiKey("X-Nomad-Token", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // config.AddApiKeyPrefix("X-Nomad-Token", "Bearer");

            var apiInstance = new AllocationsApi(config);
            var allocID = allocID_example;  // string | Allocation ID.
            var region = region_example;  // string | Filters results based on the specified region. (optional) 
            var _namespace = _namespace_example;  // string | Filters results based on the specified namespace. (optional) 
            var index = 56;  // int? | If set, wait until query exceeds given index. Must be provided with WaitParam. (optional) 
            var wait = wait_example;  // string | Provided with IndexParam to wait for change. (optional) 
            var stale = stale_example;  // string | If present, results will include stale reads. (optional) 
            var prefix = prefix_example;  // string | Constrains results to jobs that start with the defined prefix (optional) 
            var xNomadToken = xNomadToken_example;  // string | A Nomad ACL token. (optional) 
            var perPage = 56;  // int? | Maximum number of results to return. (optional) 
            var nextToken = nextToken_example;  // string | Indicates where to start paging for queries that support pagination. (optional) 
            var noShutdownDelay = true;  // bool? | Flag indicating whether to delay shutdown when requesting an allocation stop. (optional) 

            try
            {
                AllocStopResponse result = apiInstance.PostAllocationStop(allocID, region, _namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken, noShutdownDelay);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling AllocationsApi.PostAllocationStop: " + e.Message );
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
 **allocID** | **string**| Allocation ID. | 
 **region** | **string**| Filters results based on the specified region. | [optional] 
 **_namespace** | **string**| Filters results based on the specified namespace. | [optional] 
 **index** | **int?**| If set, wait until query exceeds given index. Must be provided with WaitParam. | [optional] 
 **wait** | **string**| Provided with IndexParam to wait for change. | [optional] 
 **stale** | **string**| If present, results will include stale reads. | [optional] 
 **prefix** | **string**| Constrains results to jobs that start with the defined prefix | [optional] 
 **xNomadToken** | **string**| A Nomad ACL token. | [optional] 
 **perPage** | **int?**| Maximum number of results to return. | [optional] 
 **nextToken** | **string**| Indicates where to start paging for queries that support pagination. | [optional] 
 **noShutdownDelay** | **bool?**| Flag indicating whether to delay shutdown when requesting an allocation stop. | [optional] 

### Return type

[**AllocStopResponse**](AllocStopResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** |  |  * X-Nomad-Index - A unique identifier representing the current state of the requested resource. On a new Nomad cluster the value of this index starts at 1. <br>  |
| **400** | Bad request |  -  |
| **403** | Forbidden |  -  |
| **405** | Method not allowed |  -  |
| **500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

