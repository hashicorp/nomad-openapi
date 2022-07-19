# Nomad.Client.Api.NamespacesApi

All URIs are relative to *https://127.0.0.1:4646/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNamespace**](NamespacesApi.md#createnamespace) | **POST** /namespace | 
[**DeleteNamespace**](NamespacesApi.md#deletenamespace) | **DELETE** /namespace/{namespaceName} | 
[**GetNamespace**](NamespacesApi.md#getnamespace) | **GET** /namespace/{namespaceName} | 
[**GetNamespaces**](NamespacesApi.md#getnamespaces) | **GET** /namespaces | 
[**PostNamespace**](NamespacesApi.md#postnamespace) | **POST** /namespace/{namespaceName} | 


<a name="createnamespace"></a>
# **CreateNamespace**
> void CreateNamespace (string region = null, string _namespace = null, string xNomadToken = null, string idempotencyToken = null)



### Example
```csharp
using System.Collections.Generic;
using System.Diagnostics;
using Nomad.Client.Api;
using Nomad.Client.Client;
using Nomad.Client.Model;

namespace Example
{
    public class CreateNamespaceExample
    {
        public static void Main()
        {
            Configuration config = new Configuration();
            config.BasePath = "https://127.0.0.1:4646/v1";
            // Configure API key authorization: X-Nomad-Token
            config.AddApiKey("X-Nomad-Token", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // config.AddApiKeyPrefix("X-Nomad-Token", "Bearer");

            var apiInstance = new NamespacesApi(config);
            var region = region_example;  // string | Filters results based on the specified region. (optional) 
            var _namespace = _namespace_example;  // string | Filters results based on the specified namespace. (optional) 
            var xNomadToken = xNomadToken_example;  // string | A Nomad ACL token. (optional) 
            var idempotencyToken = idempotencyToken_example;  // string | Can be used to ensure operations are only run once. (optional) 

            try
            {
                apiInstance.CreateNamespace(region, _namespace, xNomadToken, idempotencyToken);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling NamespacesApi.CreateNamespace: " + e.Message );
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
 **xNomadToken** | **string**| A Nomad ACL token. | [optional] 
 **idempotencyToken** | **string**| Can be used to ensure operations are only run once. | [optional] 

### Return type

void (empty response body)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** |  |  * X-Nomad-Index - A unique identifier representing the current state of the requested resource. On a new Nomad cluster the value of this index starts at 1. <br>  |
| **400** | Bad request |  -  |
| **403** | Forbidden |  -  |
| **405** | Method not allowed |  -  |
| **500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="deletenamespace"></a>
# **DeleteNamespace**
> void DeleteNamespace (string namespaceName, string region = null, string _namespace = null, string xNomadToken = null, string idempotencyToken = null)



### Example
```csharp
using System.Collections.Generic;
using System.Diagnostics;
using Nomad.Client.Api;
using Nomad.Client.Client;
using Nomad.Client.Model;

namespace Example
{
    public class DeleteNamespaceExample
    {
        public static void Main()
        {
            Configuration config = new Configuration();
            config.BasePath = "https://127.0.0.1:4646/v1";
            // Configure API key authorization: X-Nomad-Token
            config.AddApiKey("X-Nomad-Token", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // config.AddApiKeyPrefix("X-Nomad-Token", "Bearer");

            var apiInstance = new NamespacesApi(config);
            var namespaceName = namespaceName_example;  // string | The namespace identifier.
            var region = region_example;  // string | Filters results based on the specified region. (optional) 
            var _namespace = _namespace_example;  // string | Filters results based on the specified namespace. (optional) 
            var xNomadToken = xNomadToken_example;  // string | A Nomad ACL token. (optional) 
            var idempotencyToken = idempotencyToken_example;  // string | Can be used to ensure operations are only run once. (optional) 

            try
            {
                apiInstance.DeleteNamespace(namespaceName, region, _namespace, xNomadToken, idempotencyToken);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling NamespacesApi.DeleteNamespace: " + e.Message );
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
 **namespaceName** | **string**| The namespace identifier. | 
 **region** | **string**| Filters results based on the specified region. | [optional] 
 **_namespace** | **string**| Filters results based on the specified namespace. | [optional] 
 **xNomadToken** | **string**| A Nomad ACL token. | [optional] 
 **idempotencyToken** | **string**| Can be used to ensure operations are only run once. | [optional] 

### Return type

void (empty response body)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** |  |  * X-Nomad-Index - A unique identifier representing the current state of the requested resource. On a new Nomad cluster the value of this index starts at 1. <br>  |
| **400** | Bad request |  -  |
| **403** | Forbidden |  -  |
| **405** | Method not allowed |  -  |
| **500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

<a name="getnamespace"></a>
# **GetNamespace**
> Namespace GetNamespace (string namespaceName, string region = null, string _namespace = null, int? index = null, string wait = null, string stale = null, string prefix = null, string xNomadToken = null, int? perPage = null, string nextToken = null)



### Example
```csharp
using System.Collections.Generic;
using System.Diagnostics;
using Nomad.Client.Api;
using Nomad.Client.Client;
using Nomad.Client.Model;

namespace Example
{
    public class GetNamespaceExample
    {
        public static void Main()
        {
            Configuration config = new Configuration();
            config.BasePath = "https://127.0.0.1:4646/v1";
            // Configure API key authorization: X-Nomad-Token
            config.AddApiKey("X-Nomad-Token", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // config.AddApiKeyPrefix("X-Nomad-Token", "Bearer");

            var apiInstance = new NamespacesApi(config);
            var namespaceName = namespaceName_example;  // string | The namespace identifier.
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
                Namespace result = apiInstance.GetNamespace(namespaceName, region, _namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling NamespacesApi.GetNamespace: " + e.Message );
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
 **namespaceName** | **string**| The namespace identifier. | 
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

[**Namespace**](Namespace.md)

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

<a name="getnamespaces"></a>
# **GetNamespaces**
> List&lt;Namespace&gt; GetNamespaces (string region = null, string _namespace = null, int? index = null, string wait = null, string stale = null, string prefix = null, string xNomadToken = null, int? perPage = null, string nextToken = null)



### Example
```csharp
using System.Collections.Generic;
using System.Diagnostics;
using Nomad.Client.Api;
using Nomad.Client.Client;
using Nomad.Client.Model;

namespace Example
{
    public class GetNamespacesExample
    {
        public static void Main()
        {
            Configuration config = new Configuration();
            config.BasePath = "https://127.0.0.1:4646/v1";
            // Configure API key authorization: X-Nomad-Token
            config.AddApiKey("X-Nomad-Token", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // config.AddApiKeyPrefix("X-Nomad-Token", "Bearer");

            var apiInstance = new NamespacesApi(config);
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
                List<Namespace> result = apiInstance.GetNamespaces(region, _namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling NamespacesApi.GetNamespaces: " + e.Message );
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

[**List&lt;Namespace&gt;**](Namespace.md)

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

<a name="postnamespace"></a>
# **PostNamespace**
> void PostNamespace (string namespaceName, Namespace _namespace2, string region = null, string _namespace = null, string xNomadToken = null, string idempotencyToken = null)



### Example
```csharp
using System.Collections.Generic;
using System.Diagnostics;
using Nomad.Client.Api;
using Nomad.Client.Client;
using Nomad.Client.Model;

namespace Example
{
    public class PostNamespaceExample
    {
        public static void Main()
        {
            Configuration config = new Configuration();
            config.BasePath = "https://127.0.0.1:4646/v1";
            // Configure API key authorization: X-Nomad-Token
            config.AddApiKey("X-Nomad-Token", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // config.AddApiKeyPrefix("X-Nomad-Token", "Bearer");

            var apiInstance = new NamespacesApi(config);
            var namespaceName = namespaceName_example;  // string | The namespace identifier.
            var _namespace2 = new Namespace(); // Namespace | 
            var region = region_example;  // string | Filters results based on the specified region. (optional) 
            var _namespace = _namespace_example;  // string | Filters results based on the specified namespace. (optional) 
            var xNomadToken = xNomadToken_example;  // string | A Nomad ACL token. (optional) 
            var idempotencyToken = idempotencyToken_example;  // string | Can be used to ensure operations are only run once. (optional) 

            try
            {
                apiInstance.PostNamespace(namespaceName, _namespace2, region, _namespace, xNomadToken, idempotencyToken);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling NamespacesApi.PostNamespace: " + e.Message );
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
 **namespaceName** | **string**| The namespace identifier. | 
 **_namespace2** | [**Namespace**](Namespace.md)|  | 
 **region** | **string**| Filters results based on the specified region. | [optional] 
 **_namespace** | **string**| Filters results based on the specified namespace. | [optional] 
 **xNomadToken** | **string**| A Nomad ACL token. | [optional] 
 **idempotencyToken** | **string**| Can be used to ensure operations are only run once. | [optional] 

### Return type

void (empty response body)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined


### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
| **200** |  |  * X-Nomad-Index - A unique identifier representing the current state of the requested resource. On a new Nomad cluster the value of this index starts at 1. <br>  |
| **400** | Bad request |  -  |
| **403** | Forbidden |  -  |
| **405** | Method not allowed |  -  |
| **500** | Internal server error |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

