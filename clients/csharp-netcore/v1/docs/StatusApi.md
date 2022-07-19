# Nomad.Client.Api.StatusApi

All URIs are relative to *https://127.0.0.1:4646/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetStatusLeader**](StatusApi.md#getstatusleader) | **GET** /status/leader | 
[**GetStatusPeers**](StatusApi.md#getstatuspeers) | **GET** /status/peers | 


<a name="getstatusleader"></a>
# **GetStatusLeader**
> string GetStatusLeader (string region = null, string _namespace = null, int? index = null, string wait = null, string stale = null, string prefix = null, string xNomadToken = null, int? perPage = null, string nextToken = null)



### Example
```csharp
using System.Collections.Generic;
using System.Diagnostics;
using Nomad.Client.Api;
using Nomad.Client.Client;
using Nomad.Client.Model;

namespace Example
{
    public class GetStatusLeaderExample
    {
        public static void Main()
        {
            Configuration config = new Configuration();
            config.BasePath = "https://127.0.0.1:4646/v1";
            // Configure API key authorization: X-Nomad-Token
            config.AddApiKey("X-Nomad-Token", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // config.AddApiKeyPrefix("X-Nomad-Token", "Bearer");

            var apiInstance = new StatusApi(config);
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
                string result = apiInstance.GetStatusLeader(region, _namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling StatusApi.GetStatusLeader: " + e.Message );
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

**string**

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

<a name="getstatuspeers"></a>
# **GetStatusPeers**
> List&lt;string&gt; GetStatusPeers (string region = null, string _namespace = null, int? index = null, string wait = null, string stale = null, string prefix = null, string xNomadToken = null, int? perPage = null, string nextToken = null)



### Example
```csharp
using System.Collections.Generic;
using System.Diagnostics;
using Nomad.Client.Api;
using Nomad.Client.Client;
using Nomad.Client.Model;

namespace Example
{
    public class GetStatusPeersExample
    {
        public static void Main()
        {
            Configuration config = new Configuration();
            config.BasePath = "https://127.0.0.1:4646/v1";
            // Configure API key authorization: X-Nomad-Token
            config.AddApiKey("X-Nomad-Token", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // config.AddApiKeyPrefix("X-Nomad-Token", "Bearer");

            var apiInstance = new StatusApi(config);
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
                List<string> result = apiInstance.GetStatusPeers(region, _namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling StatusApi.GetStatusPeers: " + e.Message );
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

**List<string>**

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

