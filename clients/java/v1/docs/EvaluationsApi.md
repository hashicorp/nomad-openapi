# EvaluationsApi

All URIs are relative to *http://127.0.0.1:4646/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**getEvaluation**](EvaluationsApi.md#getEvaluation) | **GET** /evaluation/{evalID} | 
[**getEvaluationAllocations**](EvaluationsApi.md#getEvaluationAllocations) | **GET** /evaluation/{evalID}/allocations | 
[**getEvaluations**](EvaluationsApi.md#getEvaluations) | **GET** /evaluations | 


<a name="getEvaluation"></a>
# **getEvaluation**
> Evaluation getEvaluation(evalID, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken)



### Example
```java
// Import classes:
import io.nomadproject.client.ApiClient;
import io.nomadproject.client.ApiException;
import io.nomadproject.client.Configuration;
import io.nomadproject.client.auth.*;
import io.nomadproject.client.models.*;
import io.nomadproject.client.api.EvaluationsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://127.0.0.1:4646/v1");
    
    // Configure API key authorization: X-Nomad-Token
    ApiKeyAuth X-Nomad-Token = (ApiKeyAuth) defaultClient.getAuthentication("X-Nomad-Token");
    X-Nomad-Token.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //X-Nomad-Token.setApiKeyPrefix("Token");

    EvaluationsApi apiInstance = new EvaluationsApi(defaultClient);
    String evalID = "evalID_example"; // String | Evaluation ID.
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
      Evaluation result = apiInstance.getEvaluation(evalID, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling EvaluationsApi#getEvaluation");
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
 **evalID** | **String**| Evaluation ID. |
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

[**Evaluation**](Evaluation.md)

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

<a name="getEvaluationAllocations"></a>
# **getEvaluationAllocations**
> List&lt;AllocationListStub&gt; getEvaluationAllocations(evalID, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken)



### Example
```java
// Import classes:
import io.nomadproject.client.ApiClient;
import io.nomadproject.client.ApiException;
import io.nomadproject.client.Configuration;
import io.nomadproject.client.auth.*;
import io.nomadproject.client.models.*;
import io.nomadproject.client.api.EvaluationsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://127.0.0.1:4646/v1");
    
    // Configure API key authorization: X-Nomad-Token
    ApiKeyAuth X-Nomad-Token = (ApiKeyAuth) defaultClient.getAuthentication("X-Nomad-Token");
    X-Nomad-Token.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //X-Nomad-Token.setApiKeyPrefix("Token");

    EvaluationsApi apiInstance = new EvaluationsApi(defaultClient);
    String evalID = "evalID_example"; // String | Evaluation ID.
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
      List<AllocationListStub> result = apiInstance.getEvaluationAllocations(evalID, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling EvaluationsApi#getEvaluationAllocations");
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
 **evalID** | **String**| Evaluation ID. |
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

<a name="getEvaluations"></a>
# **getEvaluations**
> List&lt;Evaluation&gt; getEvaluations(region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken)



### Example
```java
// Import classes:
import io.nomadproject.client.ApiClient;
import io.nomadproject.client.ApiException;
import io.nomadproject.client.Configuration;
import io.nomadproject.client.auth.*;
import io.nomadproject.client.models.*;
import io.nomadproject.client.api.EvaluationsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://127.0.0.1:4646/v1");
    
    // Configure API key authorization: X-Nomad-Token
    ApiKeyAuth X-Nomad-Token = (ApiKeyAuth) defaultClient.getAuthentication("X-Nomad-Token");
    X-Nomad-Token.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //X-Nomad-Token.setApiKeyPrefix("Token");

    EvaluationsApi apiInstance = new EvaluationsApi(defaultClient);
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
      List<Evaluation> result = apiInstance.getEvaluations(region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling EvaluationsApi#getEvaluations");
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

[**List&lt;Evaluation&gt;**](Evaluation.md)

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

