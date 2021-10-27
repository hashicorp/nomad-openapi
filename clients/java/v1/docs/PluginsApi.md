# PluginsApi

All URIs are relative to *https://127.0.0.1:4646/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**getPluginCSI**](PluginsApi.md#getPluginCSI) | **GET** /plugin/csi/{pluginID} | 
[**getPlugins**](PluginsApi.md#getPlugins) | **GET** /plugins | 


<a name="getPluginCSI"></a>
# **getPluginCSI**
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
> List&lt;CSIPlugin&gt; getPluginCSI(pluginID, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken)
=======
> List&lt;CSIPlugin&gt; getPluginCSI(pluginId, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken)
>>>>>>> d99c3e2 (populated generator/plugins.go, added v1/plugins.go)
=======
> List&lt;CSIPlugin&gt; getPluginCSI(pluginID, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken)
>>>>>>> 13a3eee (added tests for plugins)
=======
> List&lt;CSIPlugin&gt; getPluginCSI(pluginID, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken)
>>>>>>> 6f570d317a34c315cff4c0923431310f4315843b



### Example
```java
// Import classes:
import io.nomadproject.client.ApiClient;
import io.nomadproject.client.ApiException;
import io.nomadproject.client.Configuration;
import io.nomadproject.client.auth.*;
import io.nomadproject.client.models.*;
import io.nomadproject.client.api.PluginsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://127.0.0.1:4646/v1");
    
    // Configure API key authorization: X-Nomad-Token
    ApiKeyAuth X-Nomad-Token = (ApiKeyAuth) defaultClient.getAuthentication("X-Nomad-Token");
    X-Nomad-Token.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //X-Nomad-Token.setApiKeyPrefix("Token");

    PluginsApi apiInstance = new PluginsApi(defaultClient);
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
    String pluginID = "pluginID_example"; // String | The CSI plugin identifier.
=======
    String pluginId = "pluginId_example"; // String | The CSI plugin identifier.
>>>>>>> d99c3e2 (populated generator/plugins.go, added v1/plugins.go)
=======
    String pluginID = "pluginID_example"; // String | The CSI plugin identifier.
>>>>>>> 13a3eee (added tests for plugins)
=======
    String pluginID = "pluginID_example"; // String | The CSI plugin identifier.
>>>>>>> 6f570d317a34c315cff4c0923431310f4315843b
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
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
      List<CSIPlugin> result = apiInstance.getPluginCSI(pluginID, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken);
=======
      List<CSIPlugin> result = apiInstance.getPluginCSI(pluginId, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken);
>>>>>>> d99c3e2 (populated generator/plugins.go, added v1/plugins.go)
=======
      List<CSIPlugin> result = apiInstance.getPluginCSI(pluginID, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken);
>>>>>>> 13a3eee (added tests for plugins)
=======
      List<CSIPlugin> result = apiInstance.getPluginCSI(pluginID, region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken);
>>>>>>> 6f570d317a34c315cff4c0923431310f4315843b
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling PluginsApi#getPluginCSI");
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
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
 **pluginID** | **String**| The CSI plugin identifier. |
=======
 **pluginId** | **String**| The CSI plugin identifier. |
>>>>>>> d99c3e2 (populated generator/plugins.go, added v1/plugins.go)
=======
 **pluginID** | **String**| The CSI plugin identifier. |
>>>>>>> 13a3eee (added tests for plugins)
=======
 **pluginID** | **String**| The CSI plugin identifier. |
>>>>>>> 6f570d317a34c315cff4c0923431310f4315843b
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

[**List&lt;CSIPlugin&gt;**](CSIPlugin.md)

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

<a name="getPlugins"></a>
# **getPlugins**
> List&lt;CSIPluginListStub&gt; getPlugins(region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken)



### Example
```java
// Import classes:
import io.nomadproject.client.ApiClient;
import io.nomadproject.client.ApiException;
import io.nomadproject.client.Configuration;
import io.nomadproject.client.auth.*;
import io.nomadproject.client.models.*;
import io.nomadproject.client.api.PluginsApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("https://127.0.0.1:4646/v1");
    
    // Configure API key authorization: X-Nomad-Token
    ApiKeyAuth X-Nomad-Token = (ApiKeyAuth) defaultClient.getAuthentication("X-Nomad-Token");
    X-Nomad-Token.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //X-Nomad-Token.setApiKeyPrefix("Token");

    PluginsApi apiInstance = new PluginsApi(defaultClient);
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
      List<CSIPluginListStub> result = apiInstance.getPlugins(region, namespace, index, wait, stale, prefix, xNomadToken, perPage, nextToken);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling PluginsApi#getPlugins");
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

[**List&lt;CSIPluginListStub&gt;**](CSIPluginListStub.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
**200** |  |  -  |
=======
**200** |  |  * X-Nomad-Index - A unique identifier representing the current state of the requested resource. On a new Nomad cluster the value of this index starts at 1. <br>  * X-Nomad-KnownLeader - Boolean indicating if there is a known cluster leader. <br>  * X-Nomad-LastContact - The time in milliseconds that a server was last contacted by the leader node. <br>  |
>>>>>>> d99c3e2 (populated generator/plugins.go, added v1/plugins.go)
=======
**200** |  |  -  |
>>>>>>> 13a3eee (added tests for plugins)
=======
**200** |  |  -  |
>>>>>>> 6f570d317a34c315cff4c0923431310f4315843b
**400** | Bad request |  -  |
**403** | Forbidden |  -  |
**405** | Method not allowed |  -  |
**500** | Internal server error |  -  |

