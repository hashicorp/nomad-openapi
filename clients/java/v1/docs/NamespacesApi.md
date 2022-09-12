# NamespacesApi

All URIs are relative to *http://127.0.0.1:4646/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createNamespace**](NamespacesApi.md#createNamespace) | **POST** /namespace | 
[**deleteNamespace**](NamespacesApi.md#deleteNamespace) | **DELETE** /namespace/{namespaceName} | 
[**getNamespace**](NamespacesApi.md#getNamespace) | **GET** /namespace/{namespaceName} | 
[**getNamespaces**](NamespacesApi.md#getNamespaces) | **GET** /namespaces | 
[**postNamespace**](NamespacesApi.md#postNamespace) | **POST** /namespace/{namespaceName} | 


<a name="createNamespace"></a>
# **createNamespace**
> createNamespace(region, namespace, xNomadToken, idempotencyToken)



### Example
```java
// Import classes:
import io.nomadproject.client.ApiClient;
import io.nomadproject.client.ApiException;
import io.nomadproject.client.Configuration;
import io.nomadproject.client.auth.*;
import io.nomadproject.client.models.*;
import io.nomadproject.client.api.NamespacesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://127.0.0.1:4646/v1");
    
    // Configure API key authorization: X-Nomad-Token
    ApiKeyAuth X-Nomad-Token = (ApiKeyAuth) defaultClient.getAuthentication("X-Nomad-Token");
    X-Nomad-Token.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //X-Nomad-Token.setApiKeyPrefix("Token");

    NamespacesApi apiInstance = new NamespacesApi(defaultClient);
    String region = "region_example"; // String | Filters results based on the specified region.
    String namespace = "namespace_example"; // String | Filters results based on the specified namespace.
    String xNomadToken = "xNomadToken_example"; // String | A Nomad ACL token.
    String idempotencyToken = "idempotencyToken_example"; // String | Can be used to ensure operations are only run once.
    try {
      apiInstance.createNamespace(region, namespace, xNomadToken, idempotencyToken);
    } catch (ApiException e) {
      System.err.println("Exception when calling NamespacesApi#createNamespace");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **region** | **String**| Filters results based on the specified region. | [optional]
 **namespace** | **String**| Filters results based on the specified namespace. | [optional]
 **xNomadToken** | **String**| A Nomad ACL token. | [optional]
 **idempotencyToken** | **String**| Can be used to ensure operations are only run once. | [optional]

### Return type

null (empty response body)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

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

<a name="deleteNamespace"></a>
# **deleteNamespace**
> deleteNamespace(namespaceName, region, namespace, xNomadToken, idempotencyToken)



### Example
```java
// Import classes:
import io.nomadproject.client.ApiClient;
import io.nomadproject.client.ApiException;
import io.nomadproject.client.Configuration;
import io.nomadproject.client.auth.*;
import io.nomadproject.client.models.*;
import io.nomadproject.client.api.NamespacesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://127.0.0.1:4646/v1");
    
    // Configure API key authorization: X-Nomad-Token
    ApiKeyAuth X-Nomad-Token = (ApiKeyAuth) defaultClient.getAuthentication("X-Nomad-Token");
    X-Nomad-Token.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //X-Nomad-Token.setApiKeyPrefix("Token");

    NamespacesApi apiInstance = new NamespacesApi(defaultClient);
    String namespaceName = "namespaceName_example"; // String | The namespace identifier.
    String region = "region_example"; // String | Filters results based on the specified region.
    String namespace = "namespace_example"; // String | Filters results based on the specified namespace.
    String xNomadToken = "xNomadToken_example"; // String | A Nomad ACL token.
    String idempotencyToken = "idempotencyToken_example"; // String | Can be used to ensure operations are only run once.
    try {
      apiInstance.deleteNamespace(namespaceName, region, namespace, xNomadToken, idempotencyToken);
    } catch (ApiException e) {
      System.err.println("Exception when calling NamespacesApi#deleteNamespace");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespaceName** | **String**| The namespace identifier. |
 **region** | **String**| Filters results based on the specified region. | [optional]
 **namespace** | **String**| Filters results based on the specified namespace. | [optional]
 **xNomadToken** | **String**| A Nomad ACL token. | [optional]
 **idempotencyToken** | **String**| Can be used to ensure operations are only run once. | [optional]

### Return type

null (empty response body)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

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

<a name="getNamespace"></a>
# **getNamespace**
> Namespace getNamespace(namespaceName, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken)



### Example
```java
// Import classes:
import io.nomadproject.client.ApiClient;
import io.nomadproject.client.ApiException;
import io.nomadproject.client.Configuration;
import io.nomadproject.client.auth.*;
import io.nomadproject.client.models.*;
import io.nomadproject.client.api.NamespacesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://127.0.0.1:4646/v1");
    
    // Configure API key authorization: X-Nomad-Token
    ApiKeyAuth X-Nomad-Token = (ApiKeyAuth) defaultClient.getAuthentication("X-Nomad-Token");
    X-Nomad-Token.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //X-Nomad-Token.setApiKeyPrefix("Token");

    NamespacesApi apiInstance = new NamespacesApi(defaultClient);
    String namespaceName = "namespaceName_example"; // String | The namespace identifier.
    String region = "region_example"; // String | Filters results based on the specified region.
    String namespace = "namespace_example"; // String | Filters results based on the specified namespace.
    Integer index = 56; // Integer | If set, wait until query exceeds given index. Must be provided with WaitParam.
    String wait = "wait_example"; // String | Provided with IndexParam to wait for change.
    String stale = "stale_example"; // String | If present, results will include stale reads.
    String prefix = "prefix_example"; // String | Constrains results to jobs that start with the defined prefix
    String xNomadToken = "xNomadToken_example"; // String | A Nomad ACL token.
    Integer perPage = 56; // Integer | Maximum number of results to return.
    String nextToken = "nextToken_example"; // String | Indicates where to start paging for queries that support pagination.
    try {
      Namespace result = apiInstance.getNamespace(namespaceName, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NamespacesApi#getNamespace");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespaceName** | **String**| The namespace identifier. |
 **region** | **String**| Filters results based on the specified region. | [optional]
 **namespace** | **String**| Filters results based on the specified namespace. | [optional]
 **index** | **Integer**| If set, wait until query exceeds given index. Must be provided with WaitParam. | [optional]
 **wait** | **String**| Provided with IndexParam to wait for change. | [optional]
 **stale** | **String**| If present, results will include stale reads. | [optional]
 **prefix** | **String**| Constrains results to jobs that start with the defined prefix | [optional]
 **xNomadToken** | **String**| A Nomad ACL token. | [optional]
 **perPage** | **Integer**| Maximum number of results to return. | [optional]
 **nextToken** | **String**| Indicates where to start paging for queries that support pagination. | [optional]

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
**200** |  |  * X-Nomad-Index - A unique identifier representing the current state of the requested resource. On a new Nomad cluster the value of this index starts at 1. <br>  * X-Nomad-KnownLeader - Boolean indicating if there is a known cluster leader. <br>  * X-Nomad-LastContact - The time in milliseconds that a server was last contacted by the leader node. <br>  |
**400** | Bad request |  -  |
**403** | Forbidden |  -  |
**405** | Method not allowed |  -  |
**500** | Internal server error |  -  |

<a name="getNamespaces"></a>
# **getNamespaces**
> List&lt;Namespace&gt; getNamespaces(region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken)



### Example
```java
// Import classes:
import io.nomadproject.client.ApiClient;
import io.nomadproject.client.ApiException;
import io.nomadproject.client.Configuration;
import io.nomadproject.client.auth.*;
import io.nomadproject.client.models.*;
import io.nomadproject.client.api.NamespacesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://127.0.0.1:4646/v1");
    
    // Configure API key authorization: X-Nomad-Token
    ApiKeyAuth X-Nomad-Token = (ApiKeyAuth) defaultClient.getAuthentication("X-Nomad-Token");
    X-Nomad-Token.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //X-Nomad-Token.setApiKeyPrefix("Token");

    NamespacesApi apiInstance = new NamespacesApi(defaultClient);
    String region = "region_example"; // String | Filters results based on the specified region.
    String namespace = "namespace_example"; // String | Filters results based on the specified namespace.
    Integer index = 56; // Integer | If set, wait until query exceeds given index. Must be provided with WaitParam.
    String wait = "wait_example"; // String | Provided with IndexParam to wait for change.
    String stale = "stale_example"; // String | If present, results will include stale reads.
    String prefix = "prefix_example"; // String | Constrains results to jobs that start with the defined prefix
    String xNomadToken = "xNomadToken_example"; // String | A Nomad ACL token.
    Integer perPage = 56; // Integer | Maximum number of results to return.
    String nextToken = "nextToken_example"; // String | Indicates where to start paging for queries that support pagination.
    try {
      List<Namespace> result = apiInstance.getNamespaces(region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NamespacesApi#getNamespaces");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **region** | **String**| Filters results based on the specified region. | [optional]
 **namespace** | **String**| Filters results based on the specified namespace. | [optional]
 **index** | **Integer**| If set, wait until query exceeds given index. Must be provided with WaitParam. | [optional]
 **wait** | **String**| Provided with IndexParam to wait for change. | [optional]
 **stale** | **String**| If present, results will include stale reads. | [optional]
 **prefix** | **String**| Constrains results to jobs that start with the defined prefix | [optional]
 **xNomadToken** | **String**| A Nomad ACL token. | [optional]
 **perPage** | **Integer**| Maximum number of results to return. | [optional]
 **nextToken** | **String**| Indicates where to start paging for queries that support pagination. | [optional]

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
**200** |  |  * X-Nomad-Index - A unique identifier representing the current state of the requested resource. On a new Nomad cluster the value of this index starts at 1. <br>  * X-Nomad-KnownLeader - Boolean indicating if there is a known cluster leader. <br>  * X-Nomad-LastContact - The time in milliseconds that a server was last contacted by the leader node. <br>  |
**400** | Bad request |  -  |
**403** | Forbidden |  -  |
**405** | Method not allowed |  -  |
**500** | Internal server error |  -  |

<a name="postNamespace"></a>
# **postNamespace**
> postNamespace(namespaceName, namespace2, region, namespace, xNomadToken, idempotencyToken)



### Example
```java
// Import classes:
import io.nomadproject.client.ApiClient;
import io.nomadproject.client.ApiException;
import io.nomadproject.client.Configuration;
import io.nomadproject.client.auth.*;
import io.nomadproject.client.models.*;
import io.nomadproject.client.api.NamespacesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://127.0.0.1:4646/v1");
    
    // Configure API key authorization: X-Nomad-Token
    ApiKeyAuth X-Nomad-Token = (ApiKeyAuth) defaultClient.getAuthentication("X-Nomad-Token");
    X-Nomad-Token.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //X-Nomad-Token.setApiKeyPrefix("Token");

    NamespacesApi apiInstance = new NamespacesApi(defaultClient);
    String namespaceName = "namespaceName_example"; // String | The namespace identifier.
    Namespace namespace2 = new Namespace(); // Namespace | 
    String region = "region_example"; // String | Filters results based on the specified region.
    String namespace = "namespace_example"; // String | Filters results based on the specified namespace.
    String xNomadToken = "xNomadToken_example"; // String | A Nomad ACL token.
    String idempotencyToken = "idempotencyToken_example"; // String | Can be used to ensure operations are only run once.
    try {
      apiInstance.postNamespace(namespaceName, namespace2, region, namespace, xNomadToken, idempotencyToken);
    } catch (ApiException e) {
      System.err.println("Exception when calling NamespacesApi#postNamespace");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespaceName** | **String**| The namespace identifier. |
 **namespace2** | [**Namespace**](Namespace.md)|  |
 **region** | **String**| Filters results based on the specified region. | [optional]
 **namespace** | **String**| Filters results based on the specified namespace. | [optional]
 **xNomadToken** | **String**| A Nomad ACL token. | [optional]
 **idempotencyToken** | **String**| Can be used to ensure operations are only run once. | [optional]

### Return type

null (empty response body)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

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

