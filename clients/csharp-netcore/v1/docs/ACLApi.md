# Nomad.Client.Api.ACLApi

All URIs are relative to *https://127.0.0.1:4646/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteACLPolicy**](ACLApi.md#deleteaclpolicy) | **DELETE** /acl/policy/{policyName} | 
[**DeleteACLToken**](ACLApi.md#deleteacltoken) | **DELETE** /acl/token/{tokenAccessor} | 
[**GetACLPolicies**](ACLApi.md#getaclpolicies) | **GET** /acl/policies | 
[**GetACLPolicy**](ACLApi.md#getaclpolicy) | **GET** /acl/policy/{policyName} | 
[**GetACLToken**](ACLApi.md#getacltoken) | **GET** /acl/token/{tokenAccessor} | 
[**GetACLTokenSelf**](ACLApi.md#getacltokenself) | **GET** /acl/token | 
[**GetACLTokens**](ACLApi.md#getacltokens) | **GET** /acl/tokens | 
[**PostACLBootstrap**](ACLApi.md#postaclbootstrap) | **POST** /acl/bootstrap | 
[**PostACLPolicy**](ACLApi.md#postaclpolicy) | **POST** /acl/policy/{policyName} | 
[**PostACLToken**](ACLApi.md#postacltoken) | **POST** /acl/token/{tokenAccessor} | 
[**PostACLTokenOnetime**](ACLApi.md#postacltokenonetime) | **POST** /acl/token/onetime | 
[**PostACLTokenOnetimeExchange**](ACLApi.md#postacltokenonetimeexchange) | **POST** /acl/token/onetime/exchange | 


<a name="deleteaclpolicy"></a>
# **DeleteACLPolicy**
> void DeleteACLPolicy (string policyName, string region = null, string _namespace = null, string xNomadToken = null, string idempotencyToken = null)



### Example
```csharp
using System.Collections.Generic;
using System.Diagnostics;
using Nomad.Client.Api;
using Nomad.Client.Client;
using Nomad.Client.Model;

namespace Example
{
    public class DeleteACLPolicyExample
    {
        public static void Main()
        {
            Configuration config = new Configuration();
            config.BasePath = "https://127.0.0.1:4646/v1";
            // Configure API key authorization: X-Nomad-Token
            config.AddApiKey("X-Nomad-Token", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // config.AddApiKeyPrefix("X-Nomad-Token", "Bearer");

            var apiInstance = new ACLApi(config);
            var policyName = policyName_example;  // string | The ACL policy name.
            var region = region_example;  // string | Filters results based on the specified region. (optional) 
            var _namespace = _namespace_example;  // string | Filters results based on the specified namespace. (optional) 
            var xNomadToken = xNomadToken_example;  // string | A Nomad ACL token. (optional) 
            var idempotencyToken = idempotencyToken_example;  // string | Can be used to ensure operations are only run once. (optional) 

            try
            {
                apiInstance.DeleteACLPolicy(policyName, region, _namespace, xNomadToken, idempotencyToken);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling ACLApi.DeleteACLPolicy: " + e.Message );
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
 **policyName** | **string**| The ACL policy name. | 
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

<a name="deleteacltoken"></a>
# **DeleteACLToken**
> void DeleteACLToken (string tokenAccessor, string region = null, string _namespace = null, string xNomadToken = null, string idempotencyToken = null)



### Example
```csharp
using System.Collections.Generic;
using System.Diagnostics;
using Nomad.Client.Api;
using Nomad.Client.Client;
using Nomad.Client.Model;

namespace Example
{
    public class DeleteACLTokenExample
    {
        public static void Main()
        {
            Configuration config = new Configuration();
            config.BasePath = "https://127.0.0.1:4646/v1";
            // Configure API key authorization: X-Nomad-Token
            config.AddApiKey("X-Nomad-Token", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // config.AddApiKeyPrefix("X-Nomad-Token", "Bearer");

            var apiInstance = new ACLApi(config);
            var tokenAccessor = tokenAccessor_example;  // string | The token accessor ID.
            var region = region_example;  // string | Filters results based on the specified region. (optional) 
            var _namespace = _namespace_example;  // string | Filters results based on the specified namespace. (optional) 
            var xNomadToken = xNomadToken_example;  // string | A Nomad ACL token. (optional) 
            var idempotencyToken = idempotencyToken_example;  // string | Can be used to ensure operations are only run once. (optional) 

            try
            {
                apiInstance.DeleteACLToken(tokenAccessor, region, _namespace, xNomadToken, idempotencyToken);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling ACLApi.DeleteACLToken: " + e.Message );
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
 **tokenAccessor** | **string**| The token accessor ID. | 
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

<a name="getaclpolicies"></a>
# **GetACLPolicies**
> List&lt;ACLPolicyListStub&gt; GetACLPolicies (string region = null, string _namespace = null, int? index = null, string wait = null, string stale = null, string prefix = null, string xNomadToken = null, int? perPage = null, string nextToken = null)



### Example
```csharp
using System.Collections.Generic;
using System.Diagnostics;
using Nomad.Client.Api;
using Nomad.Client.Client;
using Nomad.Client.Model;

namespace Example
{
    public class GetACLPoliciesExample
    {
        public static void Main()
        {
            Configuration config = new Configuration();
            config.BasePath = "https://127.0.0.1:4646/v1";
            // Configure API key authorization: X-Nomad-Token
            config.AddApiKey("X-Nomad-Token", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // config.AddApiKeyPrefix("X-Nomad-Token", "Bearer");

            var apiInstance = new ACLApi(config);
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
                List<ACLPolicyListStub> result = apiInstance.GetACLPolicies(region, _namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling ACLApi.GetACLPolicies: " + e.Message );
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

[**List&lt;ACLPolicyListStub&gt;**](ACLPolicyListStub.md)

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

<a name="getaclpolicy"></a>
# **GetACLPolicy**
> ACLPolicy GetACLPolicy (string policyName, string region = null, string _namespace = null, int? index = null, string wait = null, string stale = null, string prefix = null, string xNomadToken = null, int? perPage = null, string nextToken = null)



### Example
```csharp
using System.Collections.Generic;
using System.Diagnostics;
using Nomad.Client.Api;
using Nomad.Client.Client;
using Nomad.Client.Model;

namespace Example
{
    public class GetACLPolicyExample
    {
        public static void Main()
        {
            Configuration config = new Configuration();
            config.BasePath = "https://127.0.0.1:4646/v1";
            // Configure API key authorization: X-Nomad-Token
            config.AddApiKey("X-Nomad-Token", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // config.AddApiKeyPrefix("X-Nomad-Token", "Bearer");

            var apiInstance = new ACLApi(config);
            var policyName = policyName_example;  // string | The ACL policy name.
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
                ACLPolicy result = apiInstance.GetACLPolicy(policyName, region, _namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling ACLApi.GetACLPolicy: " + e.Message );
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
 **policyName** | **string**| The ACL policy name. | 
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

[**ACLPolicy**](ACLPolicy.md)

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

<a name="getacltoken"></a>
# **GetACLToken**
> ACLToken GetACLToken (string tokenAccessor, string region = null, string _namespace = null, int? index = null, string wait = null, string stale = null, string prefix = null, string xNomadToken = null, int? perPage = null, string nextToken = null)



### Example
```csharp
using System.Collections.Generic;
using System.Diagnostics;
using Nomad.Client.Api;
using Nomad.Client.Client;
using Nomad.Client.Model;

namespace Example
{
    public class GetACLTokenExample
    {
        public static void Main()
        {
            Configuration config = new Configuration();
            config.BasePath = "https://127.0.0.1:4646/v1";
            // Configure API key authorization: X-Nomad-Token
            config.AddApiKey("X-Nomad-Token", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // config.AddApiKeyPrefix("X-Nomad-Token", "Bearer");

            var apiInstance = new ACLApi(config);
            var tokenAccessor = tokenAccessor_example;  // string | The token accessor ID.
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
                ACLToken result = apiInstance.GetACLToken(tokenAccessor, region, _namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling ACLApi.GetACLToken: " + e.Message );
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
 **tokenAccessor** | **string**| The token accessor ID. | 
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

[**ACLToken**](ACLToken.md)

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

<a name="getacltokenself"></a>
# **GetACLTokenSelf**
> ACLToken GetACLTokenSelf (string region = null, string _namespace = null, int? index = null, string wait = null, string stale = null, string prefix = null, string xNomadToken = null, int? perPage = null, string nextToken = null)



### Example
```csharp
using System.Collections.Generic;
using System.Diagnostics;
using Nomad.Client.Api;
using Nomad.Client.Client;
using Nomad.Client.Model;

namespace Example
{
    public class GetACLTokenSelfExample
    {
        public static void Main()
        {
            Configuration config = new Configuration();
            config.BasePath = "https://127.0.0.1:4646/v1";
            // Configure API key authorization: X-Nomad-Token
            config.AddApiKey("X-Nomad-Token", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // config.AddApiKeyPrefix("X-Nomad-Token", "Bearer");

            var apiInstance = new ACLApi(config);
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
                ACLToken result = apiInstance.GetACLTokenSelf(region, _namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling ACLApi.GetACLTokenSelf: " + e.Message );
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

[**ACLToken**](ACLToken.md)

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

<a name="getacltokens"></a>
# **GetACLTokens**
> List&lt;ACLTokenListStub&gt; GetACLTokens (string region = null, string _namespace = null, int? index = null, string wait = null, string stale = null, string prefix = null, string xNomadToken = null, int? perPage = null, string nextToken = null)



### Example
```csharp
using System.Collections.Generic;
using System.Diagnostics;
using Nomad.Client.Api;
using Nomad.Client.Client;
using Nomad.Client.Model;

namespace Example
{
    public class GetACLTokensExample
    {
        public static void Main()
        {
            Configuration config = new Configuration();
            config.BasePath = "https://127.0.0.1:4646/v1";
            // Configure API key authorization: X-Nomad-Token
            config.AddApiKey("X-Nomad-Token", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // config.AddApiKeyPrefix("X-Nomad-Token", "Bearer");

            var apiInstance = new ACLApi(config);
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
                List<ACLTokenListStub> result = apiInstance.GetACLTokens(region, _namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling ACLApi.GetACLTokens: " + e.Message );
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

[**List&lt;ACLTokenListStub&gt;**](ACLTokenListStub.md)

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

<a name="postaclbootstrap"></a>
# **PostACLBootstrap**
> ACLToken PostACLBootstrap (string region = null, string _namespace = null, string xNomadToken = null, string idempotencyToken = null)



### Example
```csharp
using System.Collections.Generic;
using System.Diagnostics;
using Nomad.Client.Api;
using Nomad.Client.Client;
using Nomad.Client.Model;

namespace Example
{
    public class PostACLBootstrapExample
    {
        public static void Main()
        {
            Configuration config = new Configuration();
            config.BasePath = "https://127.0.0.1:4646/v1";
            // Configure API key authorization: X-Nomad-Token
            config.AddApiKey("X-Nomad-Token", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // config.AddApiKeyPrefix("X-Nomad-Token", "Bearer");

            var apiInstance = new ACLApi(config);
            var region = region_example;  // string | Filters results based on the specified region. (optional) 
            var _namespace = _namespace_example;  // string | Filters results based on the specified namespace. (optional) 
            var xNomadToken = xNomadToken_example;  // string | A Nomad ACL token. (optional) 
            var idempotencyToken = idempotencyToken_example;  // string | Can be used to ensure operations are only run once. (optional) 

            try
            {
                ACLToken result = apiInstance.PostACLBootstrap(region, _namespace, xNomadToken, idempotencyToken);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling ACLApi.PostACLBootstrap: " + e.Message );
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

[**ACLToken**](ACLToken.md)

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

<a name="postaclpolicy"></a>
# **PostACLPolicy**
> void PostACLPolicy (string policyName, ACLPolicy aCLPolicy, string region = null, string _namespace = null, string xNomadToken = null, string idempotencyToken = null)



### Example
```csharp
using System.Collections.Generic;
using System.Diagnostics;
using Nomad.Client.Api;
using Nomad.Client.Client;
using Nomad.Client.Model;

namespace Example
{
    public class PostACLPolicyExample
    {
        public static void Main()
        {
            Configuration config = new Configuration();
            config.BasePath = "https://127.0.0.1:4646/v1";
            // Configure API key authorization: X-Nomad-Token
            config.AddApiKey("X-Nomad-Token", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // config.AddApiKeyPrefix("X-Nomad-Token", "Bearer");

            var apiInstance = new ACLApi(config);
            var policyName = policyName_example;  // string | The ACL policy name.
            var aCLPolicy = new ACLPolicy(); // ACLPolicy | 
            var region = region_example;  // string | Filters results based on the specified region. (optional) 
            var _namespace = _namespace_example;  // string | Filters results based on the specified namespace. (optional) 
            var xNomadToken = xNomadToken_example;  // string | A Nomad ACL token. (optional) 
            var idempotencyToken = idempotencyToken_example;  // string | Can be used to ensure operations are only run once. (optional) 

            try
            {
                apiInstance.PostACLPolicy(policyName, aCLPolicy, region, _namespace, xNomadToken, idempotencyToken);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling ACLApi.PostACLPolicy: " + e.Message );
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
 **policyName** | **string**| The ACL policy name. | 
 **aCLPolicy** | [**ACLPolicy**](ACLPolicy.md)|  | 
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

<a name="postacltoken"></a>
# **PostACLToken**
> ACLToken PostACLToken (string tokenAccessor, ACLToken aCLToken, string region = null, string _namespace = null, string xNomadToken = null, string idempotencyToken = null)



### Example
```csharp
using System.Collections.Generic;
using System.Diagnostics;
using Nomad.Client.Api;
using Nomad.Client.Client;
using Nomad.Client.Model;

namespace Example
{
    public class PostACLTokenExample
    {
        public static void Main()
        {
            Configuration config = new Configuration();
            config.BasePath = "https://127.0.0.1:4646/v1";
            // Configure API key authorization: X-Nomad-Token
            config.AddApiKey("X-Nomad-Token", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // config.AddApiKeyPrefix("X-Nomad-Token", "Bearer");

            var apiInstance = new ACLApi(config);
            var tokenAccessor = tokenAccessor_example;  // string | The token accessor ID.
            var aCLToken = new ACLToken(); // ACLToken | 
            var region = region_example;  // string | Filters results based on the specified region. (optional) 
            var _namespace = _namespace_example;  // string | Filters results based on the specified namespace. (optional) 
            var xNomadToken = xNomadToken_example;  // string | A Nomad ACL token. (optional) 
            var idempotencyToken = idempotencyToken_example;  // string | Can be used to ensure operations are only run once. (optional) 

            try
            {
                ACLToken result = apiInstance.PostACLToken(tokenAccessor, aCLToken, region, _namespace, xNomadToken, idempotencyToken);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling ACLApi.PostACLToken: " + e.Message );
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
 **tokenAccessor** | **string**| The token accessor ID. | 
 **aCLToken** | [**ACLToken**](ACLToken.md)|  | 
 **region** | **string**| Filters results based on the specified region. | [optional] 
 **_namespace** | **string**| Filters results based on the specified namespace. | [optional] 
 **xNomadToken** | **string**| A Nomad ACL token. | [optional] 
 **idempotencyToken** | **string**| Can be used to ensure operations are only run once. | [optional] 

### Return type

[**ACLToken**](ACLToken.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

 - **Content-Type**: application/json
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

<a name="postacltokenonetime"></a>
# **PostACLTokenOnetime**
> OneTimeToken PostACLTokenOnetime (string region = null, string _namespace = null, string xNomadToken = null, string idempotencyToken = null)



### Example
```csharp
using System.Collections.Generic;
using System.Diagnostics;
using Nomad.Client.Api;
using Nomad.Client.Client;
using Nomad.Client.Model;

namespace Example
{
    public class PostACLTokenOnetimeExample
    {
        public static void Main()
        {
            Configuration config = new Configuration();
            config.BasePath = "https://127.0.0.1:4646/v1";
            // Configure API key authorization: X-Nomad-Token
            config.AddApiKey("X-Nomad-Token", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // config.AddApiKeyPrefix("X-Nomad-Token", "Bearer");

            var apiInstance = new ACLApi(config);
            var region = region_example;  // string | Filters results based on the specified region. (optional) 
            var _namespace = _namespace_example;  // string | Filters results based on the specified namespace. (optional) 
            var xNomadToken = xNomadToken_example;  // string | A Nomad ACL token. (optional) 
            var idempotencyToken = idempotencyToken_example;  // string | Can be used to ensure operations are only run once. (optional) 

            try
            {
                OneTimeToken result = apiInstance.PostACLTokenOnetime(region, _namespace, xNomadToken, idempotencyToken);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling ACLApi.PostACLTokenOnetime: " + e.Message );
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

[**OneTimeToken**](OneTimeToken.md)

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

<a name="postacltokenonetimeexchange"></a>
# **PostACLTokenOnetimeExchange**
> ACLToken PostACLTokenOnetimeExchange (OneTimeTokenExchangeRequest oneTimeTokenExchangeRequest, string region = null, string _namespace = null, string xNomadToken = null, string idempotencyToken = null)



### Example
```csharp
using System.Collections.Generic;
using System.Diagnostics;
using Nomad.Client.Api;
using Nomad.Client.Client;
using Nomad.Client.Model;

namespace Example
{
    public class PostACLTokenOnetimeExchangeExample
    {
        public static void Main()
        {
            Configuration config = new Configuration();
            config.BasePath = "https://127.0.0.1:4646/v1";
            // Configure API key authorization: X-Nomad-Token
            config.AddApiKey("X-Nomad-Token", "YOUR_API_KEY");
            // Uncomment below to setup prefix (e.g. Bearer) for API key, if needed
            // config.AddApiKeyPrefix("X-Nomad-Token", "Bearer");

            var apiInstance = new ACLApi(config);
            var oneTimeTokenExchangeRequest = new OneTimeTokenExchangeRequest(); // OneTimeTokenExchangeRequest | 
            var region = region_example;  // string | Filters results based on the specified region. (optional) 
            var _namespace = _namespace_example;  // string | Filters results based on the specified namespace. (optional) 
            var xNomadToken = xNomadToken_example;  // string | A Nomad ACL token. (optional) 
            var idempotencyToken = idempotencyToken_example;  // string | Can be used to ensure operations are only run once. (optional) 

            try
            {
                ACLToken result = apiInstance.PostACLTokenOnetimeExchange(oneTimeTokenExchangeRequest, region, _namespace, xNomadToken, idempotencyToken);
                Debug.WriteLine(result);
            }
            catch (ApiException  e)
            {
                Debug.Print("Exception when calling ACLApi.PostACLTokenOnetimeExchange: " + e.Message );
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
 **oneTimeTokenExchangeRequest** | [**OneTimeTokenExchangeRequest**](OneTimeTokenExchangeRequest.md)|  | 
 **region** | **string**| Filters results based on the specified region. | [optional] 
 **_namespace** | **string**| Filters results based on the specified namespace. | [optional] 
 **xNomadToken** | **string**| A Nomad ACL token. | [optional] 
 **idempotencyToken** | **string**| Can be used to ensure operations are only run once. | [optional] 

### Return type

[**ACLToken**](ACLToken.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

 - **Content-Type**: application/json
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

