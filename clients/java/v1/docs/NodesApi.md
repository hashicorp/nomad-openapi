# NodesApi

All URIs are relative to *http://127.0.0.1:4646/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**getNode**](NodesApi.md#getNode) | **GET** /node/{nodeId} | 
[**getNodeAllocations**](NodesApi.md#getNodeAllocations) | **GET** /node/{nodeId}/allocations | 
[**getNodes**](NodesApi.md#getNodes) | **GET** /nodes | 
[**updateNodeDrain**](NodesApi.md#updateNodeDrain) | **POST** /node/{nodeId}/drain | 
[**updateNodeEligibility**](NodesApi.md#updateNodeEligibility) | **POST** /node/{nodeId}/eligibility | 
[**updateNodePurge**](NodesApi.md#updateNodePurge) | **POST** /node/{nodeId}/purge | 


<a name="getNode"></a>
# **getNode**
> Node getNode(nodeId, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken)



### Example
```java
// Import classes:
import io.nomadproject.client.ApiClient;
import io.nomadproject.client.ApiException;
import io.nomadproject.client.Configuration;
import io.nomadproject.client.auth.*;
import io.nomadproject.client.models.*;
import io.nomadproject.client.api.NodesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://127.0.0.1:4646/v1");
    
    // Configure API key authorization: X-Nomad-Token
    ApiKeyAuth X-Nomad-Token = (ApiKeyAuth) defaultClient.getAuthentication("X-Nomad-Token");
    X-Nomad-Token.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //X-Nomad-Token.setApiKeyPrefix("Token");

    NodesApi apiInstance = new NodesApi(defaultClient);
    String nodeId = "nodeId_example"; // String | The ID of the node.
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
      Node result = apiInstance.getNode(nodeId, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NodesApi#getNode");
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
 **nodeId** | **String**| The ID of the node. |
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

[**Node**](Node.md)

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

<a name="getNodeAllocations"></a>
# **getNodeAllocations**
> List&lt;AllocationListStub&gt; getNodeAllocations(nodeId, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken)



### Example
```java
// Import classes:
import io.nomadproject.client.ApiClient;
import io.nomadproject.client.ApiException;
import io.nomadproject.client.Configuration;
import io.nomadproject.client.auth.*;
import io.nomadproject.client.models.*;
import io.nomadproject.client.api.NodesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://127.0.0.1:4646/v1");
    
    // Configure API key authorization: X-Nomad-Token
    ApiKeyAuth X-Nomad-Token = (ApiKeyAuth) defaultClient.getAuthentication("X-Nomad-Token");
    X-Nomad-Token.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //X-Nomad-Token.setApiKeyPrefix("Token");

    NodesApi apiInstance = new NodesApi(defaultClient);
    String nodeId = "nodeId_example"; // String | The ID of the node.
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
      List<AllocationListStub> result = apiInstance.getNodeAllocations(nodeId, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NodesApi#getNodeAllocations");
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
 **nodeId** | **String**| The ID of the node. |
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

<a name="getNodes"></a>
# **getNodes**
> List&lt;NodeListStub&gt; getNodes(region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken, resources)



### Example
```java
// Import classes:
import io.nomadproject.client.ApiClient;
import io.nomadproject.client.ApiException;
import io.nomadproject.client.Configuration;
import io.nomadproject.client.auth.*;
import io.nomadproject.client.models.*;
import io.nomadproject.client.api.NodesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://127.0.0.1:4646/v1");
    
    // Configure API key authorization: X-Nomad-Token
    ApiKeyAuth X-Nomad-Token = (ApiKeyAuth) defaultClient.getAuthentication("X-Nomad-Token");
    X-Nomad-Token.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //X-Nomad-Token.setApiKeyPrefix("Token");

    NodesApi apiInstance = new NodesApi(defaultClient);
    String region = "region_example"; // String | Filters results based on the specified region.
    String namespace = "namespace_example"; // String | Filters results based on the specified namespace.
    Integer index = 56; // Integer | If set, wait until query exceeds given index. Must be provided with WaitParam.
    String wait = "wait_example"; // String | Provided with IndexParam to wait for change.
    String stale = "stale_example"; // String | If present, results will include stale reads.
    String prefix = "prefix_example"; // String | Constrains results to jobs that start with the defined prefix
    String xNomadToken = "xNomadToken_example"; // String | A Nomad ACL token.
    Integer perPage = 56; // Integer | Maximum number of results to return.
    String nextToken = "nextToken_example"; // String | Indicates where to start paging for queries that support pagination.
    Boolean resources = true; // Boolean | Whether or not to include the NodeResources and ReservedResources fields in the response.
    try {
      List<NodeListStub> result = apiInstance.getNodes(region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken, resources);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NodesApi#getNodes");
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
 **resources** | **Boolean**| Whether or not to include the NodeResources and ReservedResources fields in the response. | [optional]

### Return type

[**List&lt;NodeListStub&gt;**](NodeListStub.md)

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

<a name="updateNodeDrain"></a>
# **updateNodeDrain**
> NodeDrainUpdateResponse updateNodeDrain(nodeId, nodeUpdateDrainRequest, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken)



### Example
```java
// Import classes:
import io.nomadproject.client.ApiClient;
import io.nomadproject.client.ApiException;
import io.nomadproject.client.Configuration;
import io.nomadproject.client.auth.*;
import io.nomadproject.client.models.*;
import io.nomadproject.client.api.NodesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://127.0.0.1:4646/v1");
    
    // Configure API key authorization: X-Nomad-Token
    ApiKeyAuth X-Nomad-Token = (ApiKeyAuth) defaultClient.getAuthentication("X-Nomad-Token");
    X-Nomad-Token.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //X-Nomad-Token.setApiKeyPrefix("Token");

    NodesApi apiInstance = new NodesApi(defaultClient);
    String nodeId = "nodeId_example"; // String | The ID of the node.
    NodeUpdateDrainRequest nodeUpdateDrainRequest = new NodeUpdateDrainRequest(); // NodeUpdateDrainRequest | 
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
      NodeDrainUpdateResponse result = apiInstance.updateNodeDrain(nodeId, nodeUpdateDrainRequest, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NodesApi#updateNodeDrain");
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
 **nodeId** | **String**| The ID of the node. |
 **nodeUpdateDrainRequest** | [**NodeUpdateDrainRequest**](NodeUpdateDrainRequest.md)|  |
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

[**NodeDrainUpdateResponse**](NodeDrainUpdateResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

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

<a name="updateNodeEligibility"></a>
# **updateNodeEligibility**
> NodeEligibilityUpdateResponse updateNodeEligibility(nodeId, nodeUpdateEligibilityRequest, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken)



### Example
```java
// Import classes:
import io.nomadproject.client.ApiClient;
import io.nomadproject.client.ApiException;
import io.nomadproject.client.Configuration;
import io.nomadproject.client.auth.*;
import io.nomadproject.client.models.*;
import io.nomadproject.client.api.NodesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://127.0.0.1:4646/v1");
    
    // Configure API key authorization: X-Nomad-Token
    ApiKeyAuth X-Nomad-Token = (ApiKeyAuth) defaultClient.getAuthentication("X-Nomad-Token");
    X-Nomad-Token.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //X-Nomad-Token.setApiKeyPrefix("Token");

    NodesApi apiInstance = new NodesApi(defaultClient);
    String nodeId = "nodeId_example"; // String | The ID of the node.
    NodeUpdateEligibilityRequest nodeUpdateEligibilityRequest = new NodeUpdateEligibilityRequest(); // NodeUpdateEligibilityRequest | 
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
      NodeEligibilityUpdateResponse result = apiInstance.updateNodeEligibility(nodeId, nodeUpdateEligibilityRequest, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NodesApi#updateNodeEligibility");
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
 **nodeId** | **String**| The ID of the node. |
 **nodeUpdateEligibilityRequest** | [**NodeUpdateEligibilityRequest**](NodeUpdateEligibilityRequest.md)|  |
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

[**NodeEligibilityUpdateResponse**](NodeEligibilityUpdateResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

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

<a name="updateNodePurge"></a>
# **updateNodePurge**
> NodePurgeResponse updateNodePurge(nodeId, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken)



### Example
```java
// Import classes:
import io.nomadproject.client.ApiClient;
import io.nomadproject.client.ApiException;
import io.nomadproject.client.Configuration;
import io.nomadproject.client.auth.*;
import io.nomadproject.client.models.*;
import io.nomadproject.client.api.NodesApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://127.0.0.1:4646/v1");
    
    // Configure API key authorization: X-Nomad-Token
    ApiKeyAuth X-Nomad-Token = (ApiKeyAuth) defaultClient.getAuthentication("X-Nomad-Token");
    X-Nomad-Token.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //X-Nomad-Token.setApiKeyPrefix("Token");

    NodesApi apiInstance = new NodesApi(defaultClient);
    String nodeId = "nodeId_example"; // String | The ID of the node.
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
      NodePurgeResponse result = apiInstance.updateNodePurge(nodeId, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling NodesApi#updateNodePurge");
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
 **nodeId** | **String**| The ID of the node. |
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

[**NodePurgeResponse**](NodePurgeResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

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

