# DeploymentsApi

All URIs are relative to *http://127.0.0.1:4646/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**getDeployment**](DeploymentsApi.md#getDeployment) | **GET** /deployment/{deploymentID} | 
[**getDeploymentAllocations**](DeploymentsApi.md#getDeploymentAllocations) | **GET** /deployment/allocations/{deploymentID} | 
[**getDeployments**](DeploymentsApi.md#getDeployments) | **GET** /deployments | 
[**postDeploymentAllocationHealth**](DeploymentsApi.md#postDeploymentAllocationHealth) | **POST** /deployment/allocation-health/{deploymentID} | 
[**postDeploymentFail**](DeploymentsApi.md#postDeploymentFail) | **POST** /deployment/fail/{deploymentID} | 
[**postDeploymentPause**](DeploymentsApi.md#postDeploymentPause) | **POST** /deployment/pause/{deploymentID} | 
[**postDeploymentPromote**](DeploymentsApi.md#postDeploymentPromote) | **POST** /deployment/promote/{deploymentID} | 
[**postDeploymentUnblock**](DeploymentsApi.md#postDeploymentUnblock) | **POST** /deployment/unblock/{deploymentID} | 


<a name="getDeployment"></a>
# **getDeployment**
> Deployment getDeployment(deploymentID, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken)



### Example
```java
// Import classes:
import io.nomadproject.client.ApiClient;
import io.nomadproject.client.ApiException;
import io.nomadproject.client.Configuration;
import io.nomadproject.client.auth.*;
import io.nomadproject.client.models.*;
import io.nomadproject.client.api.DeploymentsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://127.0.0.1:4646/v1");
    
    // Configure API key authorization: X-Nomad-Token
    ApiKeyAuth X-Nomad-Token = (ApiKeyAuth) defaultClient.getAuthentication("X-Nomad-Token");
    X-Nomad-Token.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //X-Nomad-Token.setApiKeyPrefix("Token");

    DeploymentsApi apiInstance = new DeploymentsApi(defaultClient);
    String deploymentID = "deploymentID_example"; // String | Deployment ID.
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
      Deployment result = apiInstance.getDeployment(deploymentID, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DeploymentsApi#getDeployment");
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
 **deploymentID** | **String**| Deployment ID. |
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

[**Deployment**](Deployment.md)

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

<a name="getDeploymentAllocations"></a>
# **getDeploymentAllocations**
> List&lt;AllocationListStub&gt; getDeploymentAllocations(deploymentID, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken)



### Example
```java
// Import classes:
import io.nomadproject.client.ApiClient;
import io.nomadproject.client.ApiException;
import io.nomadproject.client.Configuration;
import io.nomadproject.client.auth.*;
import io.nomadproject.client.models.*;
import io.nomadproject.client.api.DeploymentsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://127.0.0.1:4646/v1");
    
    // Configure API key authorization: X-Nomad-Token
    ApiKeyAuth X-Nomad-Token = (ApiKeyAuth) defaultClient.getAuthentication("X-Nomad-Token");
    X-Nomad-Token.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //X-Nomad-Token.setApiKeyPrefix("Token");

    DeploymentsApi apiInstance = new DeploymentsApi(defaultClient);
    String deploymentID = "deploymentID_example"; // String | Deployment ID.
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
      List<AllocationListStub> result = apiInstance.getDeploymentAllocations(deploymentID, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DeploymentsApi#getDeploymentAllocations");
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
 **deploymentID** | **String**| Deployment ID. |
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

[**List&lt;AllocationListStub&gt;**](AllocationListStub.md)

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

<a name="getDeployments"></a>
# **getDeployments**
> List&lt;Deployment&gt; getDeployments(region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken)



### Example
```java
// Import classes:
import io.nomadproject.client.ApiClient;
import io.nomadproject.client.ApiException;
import io.nomadproject.client.Configuration;
import io.nomadproject.client.auth.*;
import io.nomadproject.client.models.*;
import io.nomadproject.client.api.DeploymentsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://127.0.0.1:4646/v1");
    
    // Configure API key authorization: X-Nomad-Token
    ApiKeyAuth X-Nomad-Token = (ApiKeyAuth) defaultClient.getAuthentication("X-Nomad-Token");
    X-Nomad-Token.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //X-Nomad-Token.setApiKeyPrefix("Token");

    DeploymentsApi apiInstance = new DeploymentsApi(defaultClient);
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
      List<Deployment> result = apiInstance.getDeployments(region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DeploymentsApi#getDeployments");
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

[**List&lt;Deployment&gt;**](Deployment.md)

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

<a name="postDeploymentAllocationHealth"></a>
# **postDeploymentAllocationHealth**
> DeploymentUpdateResponse postDeploymentAllocationHealth(deploymentID, deploymentAllocHealthRequest, region, namespace, xNomadToken, idempotencyToken)



### Example
```java
// Import classes:
import io.nomadproject.client.ApiClient;
import io.nomadproject.client.ApiException;
import io.nomadproject.client.Configuration;
import io.nomadproject.client.auth.*;
import io.nomadproject.client.models.*;
import io.nomadproject.client.api.DeploymentsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://127.0.0.1:4646/v1");
    
    // Configure API key authorization: X-Nomad-Token
    ApiKeyAuth X-Nomad-Token = (ApiKeyAuth) defaultClient.getAuthentication("X-Nomad-Token");
    X-Nomad-Token.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //X-Nomad-Token.setApiKeyPrefix("Token");

    DeploymentsApi apiInstance = new DeploymentsApi(defaultClient);
    String deploymentID = "deploymentID_example"; // String | Deployment ID.
    DeploymentAllocHealthRequest deploymentAllocHealthRequest = new DeploymentAllocHealthRequest(); // DeploymentAllocHealthRequest | 
    String region = "region_example"; // String | Filters results based on the specified region.
    String namespace = "namespace_example"; // String | Filters results based on the specified namespace.
    String xNomadToken = "xNomadToken_example"; // String | A Nomad ACL token.
    String idempotencyToken = "idempotencyToken_example"; // String | Can be used to ensure operations are only run once.
    try {
      DeploymentUpdateResponse result = apiInstance.postDeploymentAllocationHealth(deploymentID, deploymentAllocHealthRequest, region, namespace, xNomadToken, idempotencyToken);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DeploymentsApi#postDeploymentAllocationHealth");
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
 **deploymentID** | **String**| Deployment ID. |
 **deploymentAllocHealthRequest** | [**DeploymentAllocHealthRequest**](DeploymentAllocHealthRequest.md)|  |
 **region** | **String**| Filters results based on the specified region. | [optional]
 **namespace** | **String**| Filters results based on the specified namespace. | [optional]
 **xNomadToken** | **String**| A Nomad ACL token. | [optional]
 **idempotencyToken** | **String**| Can be used to ensure operations are only run once. | [optional]

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
**200** |  |  -  |
**400** | Bad request |  -  |
**403** | Forbidden |  -  |
**405** | Method not allowed |  -  |
**500** | Internal server error |  -  |

<a name="postDeploymentFail"></a>
# **postDeploymentFail**
> DeploymentUpdateResponse postDeploymentFail(deploymentID, region, namespace, xNomadToken, idempotencyToken)



### Example
```java
// Import classes:
import io.nomadproject.client.ApiClient;
import io.nomadproject.client.ApiException;
import io.nomadproject.client.Configuration;
import io.nomadproject.client.auth.*;
import io.nomadproject.client.models.*;
import io.nomadproject.client.api.DeploymentsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://127.0.0.1:4646/v1");
    
    // Configure API key authorization: X-Nomad-Token
    ApiKeyAuth X-Nomad-Token = (ApiKeyAuth) defaultClient.getAuthentication("X-Nomad-Token");
    X-Nomad-Token.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //X-Nomad-Token.setApiKeyPrefix("Token");

    DeploymentsApi apiInstance = new DeploymentsApi(defaultClient);
    String deploymentID = "deploymentID_example"; // String | Deployment ID.
    String region = "region_example"; // String | Filters results based on the specified region.
    String namespace = "namespace_example"; // String | Filters results based on the specified namespace.
    String xNomadToken = "xNomadToken_example"; // String | A Nomad ACL token.
    String idempotencyToken = "idempotencyToken_example"; // String | Can be used to ensure operations are only run once.
    try {
      DeploymentUpdateResponse result = apiInstance.postDeploymentFail(deploymentID, region, namespace, xNomadToken, idempotencyToken);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DeploymentsApi#postDeploymentFail");
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
 **deploymentID** | **String**| Deployment ID. |
 **region** | **String**| Filters results based on the specified region. | [optional]
 **namespace** | **String**| Filters results based on the specified namespace. | [optional]
 **xNomadToken** | **String**| A Nomad ACL token. | [optional]
 **idempotencyToken** | **String**| Can be used to ensure operations are only run once. | [optional]

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
**200** |  |  -  |
**400** | Bad request |  -  |
**403** | Forbidden |  -  |
**405** | Method not allowed |  -  |
**500** | Internal server error |  -  |

<a name="postDeploymentPause"></a>
# **postDeploymentPause**
> DeploymentUpdateResponse postDeploymentPause(deploymentID, deploymentPauseRequest, region, namespace, xNomadToken, idempotencyToken)



### Example
```java
// Import classes:
import io.nomadproject.client.ApiClient;
import io.nomadproject.client.ApiException;
import io.nomadproject.client.Configuration;
import io.nomadproject.client.auth.*;
import io.nomadproject.client.models.*;
import io.nomadproject.client.api.DeploymentsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://127.0.0.1:4646/v1");
    
    // Configure API key authorization: X-Nomad-Token
    ApiKeyAuth X-Nomad-Token = (ApiKeyAuth) defaultClient.getAuthentication("X-Nomad-Token");
    X-Nomad-Token.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //X-Nomad-Token.setApiKeyPrefix("Token");

    DeploymentsApi apiInstance = new DeploymentsApi(defaultClient);
    String deploymentID = "deploymentID_example"; // String | Deployment ID.
    DeploymentPauseRequest deploymentPauseRequest = new DeploymentPauseRequest(); // DeploymentPauseRequest | 
    String region = "region_example"; // String | Filters results based on the specified region.
    String namespace = "namespace_example"; // String | Filters results based on the specified namespace.
    String xNomadToken = "xNomadToken_example"; // String | A Nomad ACL token.
    String idempotencyToken = "idempotencyToken_example"; // String | Can be used to ensure operations are only run once.
    try {
      DeploymentUpdateResponse result = apiInstance.postDeploymentPause(deploymentID, deploymentPauseRequest, region, namespace, xNomadToken, idempotencyToken);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DeploymentsApi#postDeploymentPause");
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
 **deploymentID** | **String**| Deployment ID. |
 **deploymentPauseRequest** | [**DeploymentPauseRequest**](DeploymentPauseRequest.md)|  |
 **region** | **String**| Filters results based on the specified region. | [optional]
 **namespace** | **String**| Filters results based on the specified namespace. | [optional]
 **xNomadToken** | **String**| A Nomad ACL token. | [optional]
 **idempotencyToken** | **String**| Can be used to ensure operations are only run once. | [optional]

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
**200** |  |  -  |
**400** | Bad request |  -  |
**403** | Forbidden |  -  |
**405** | Method not allowed |  -  |
**500** | Internal server error |  -  |

<a name="postDeploymentPromote"></a>
# **postDeploymentPromote**
> DeploymentUpdateResponse postDeploymentPromote(deploymentID, deploymentPromoteRequest, region, namespace, xNomadToken, idempotencyToken)



### Example
```java
// Import classes:
import io.nomadproject.client.ApiClient;
import io.nomadproject.client.ApiException;
import io.nomadproject.client.Configuration;
import io.nomadproject.client.auth.*;
import io.nomadproject.client.models.*;
import io.nomadproject.client.api.DeploymentsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://127.0.0.1:4646/v1");
    
    // Configure API key authorization: X-Nomad-Token
    ApiKeyAuth X-Nomad-Token = (ApiKeyAuth) defaultClient.getAuthentication("X-Nomad-Token");
    X-Nomad-Token.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //X-Nomad-Token.setApiKeyPrefix("Token");

    DeploymentsApi apiInstance = new DeploymentsApi(defaultClient);
    String deploymentID = "deploymentID_example"; // String | Deployment ID.
    DeploymentPromoteRequest deploymentPromoteRequest = new DeploymentPromoteRequest(); // DeploymentPromoteRequest | 
    String region = "region_example"; // String | Filters results based on the specified region.
    String namespace = "namespace_example"; // String | Filters results based on the specified namespace.
    String xNomadToken = "xNomadToken_example"; // String | A Nomad ACL token.
    String idempotencyToken = "idempotencyToken_example"; // String | Can be used to ensure operations are only run once.
    try {
      DeploymentUpdateResponse result = apiInstance.postDeploymentPromote(deploymentID, deploymentPromoteRequest, region, namespace, xNomadToken, idempotencyToken);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DeploymentsApi#postDeploymentPromote");
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
 **deploymentID** | **String**| Deployment ID. |
 **deploymentPromoteRequest** | [**DeploymentPromoteRequest**](DeploymentPromoteRequest.md)|  |
 **region** | **String**| Filters results based on the specified region. | [optional]
 **namespace** | **String**| Filters results based on the specified namespace. | [optional]
 **xNomadToken** | **String**| A Nomad ACL token. | [optional]
 **idempotencyToken** | **String**| Can be used to ensure operations are only run once. | [optional]

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
**200** |  |  -  |
**400** | Bad request |  -  |
**403** | Forbidden |  -  |
**405** | Method not allowed |  -  |
**500** | Internal server error |  -  |

<a name="postDeploymentUnblock"></a>
# **postDeploymentUnblock**
> DeploymentUpdateResponse postDeploymentUnblock(deploymentID, deploymentUnblockRequest, region, namespace, xNomadToken, idempotencyToken)



### Example
```java
// Import classes:
import io.nomadproject.client.ApiClient;
import io.nomadproject.client.ApiException;
import io.nomadproject.client.Configuration;
import io.nomadproject.client.auth.*;
import io.nomadproject.client.models.*;
import io.nomadproject.client.api.DeploymentsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://127.0.0.1:4646/v1");
    
    // Configure API key authorization: X-Nomad-Token
    ApiKeyAuth X-Nomad-Token = (ApiKeyAuth) defaultClient.getAuthentication("X-Nomad-Token");
    X-Nomad-Token.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //X-Nomad-Token.setApiKeyPrefix("Token");

    DeploymentsApi apiInstance = new DeploymentsApi(defaultClient);
    String deploymentID = "deploymentID_example"; // String | Deployment ID.
    DeploymentUnblockRequest deploymentUnblockRequest = new DeploymentUnblockRequest(); // DeploymentUnblockRequest | 
    String region = "region_example"; // String | Filters results based on the specified region.
    String namespace = "namespace_example"; // String | Filters results based on the specified namespace.
    String xNomadToken = "xNomadToken_example"; // String | A Nomad ACL token.
    String idempotencyToken = "idempotencyToken_example"; // String | Can be used to ensure operations are only run once.
    try {
      DeploymentUpdateResponse result = apiInstance.postDeploymentUnblock(deploymentID, deploymentUnblockRequest, region, namespace, xNomadToken, idempotencyToken);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling DeploymentsApi#postDeploymentUnblock");
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
 **deploymentID** | **String**| Deployment ID. |
 **deploymentUnblockRequest** | [**DeploymentUnblockRequest**](DeploymentUnblockRequest.md)|  |
 **region** | **String**| Filters results based on the specified region. | [optional]
 **namespace** | **String**| Filters results based on the specified namespace. | [optional]
 **xNomadToken** | **String**| A Nomad ACL token. | [optional]
 **idempotencyToken** | **String**| Can be used to ensure operations are only run once. | [optional]

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
**200** |  |  -  |
**400** | Bad request |  -  |
**403** | Forbidden |  -  |
**405** | Method not allowed |  -  |
**500** | Internal server error |  -  |

