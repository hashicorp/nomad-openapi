# SystemApi

All URIs are relative to *http://127.0.0.1:4646/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**putSystemGC**](SystemApi.md#putSystemGC) | **PUT** /system/gc | 
[**putSystemReconcileSummaries**](SystemApi.md#putSystemReconcileSummaries) | **PUT** /system/reconcile/summaries | 


<a name="putSystemGC"></a>
# **putSystemGC**
> putSystemGC(region, namespace, xNomadToken, idempotencyToken)



### Example
```java
// Import classes:
import io.nomadproject.client.ApiClient;
import io.nomadproject.client.ApiException;
import io.nomadproject.client.Configuration;
import io.nomadproject.client.auth.*;
import io.nomadproject.client.models.*;
import io.nomadproject.client.api.SystemApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://127.0.0.1:4646/v1");
    
    // Configure API key authorization: X-Nomad-Token
    ApiKeyAuth X-Nomad-Token = (ApiKeyAuth) defaultClient.getAuthentication("X-Nomad-Token");
    X-Nomad-Token.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //X-Nomad-Token.setApiKeyPrefix("Token");

    SystemApi apiInstance = new SystemApi(defaultClient);
    String region = "region_example"; // String | Filters results based on the specified region.
    String namespace = "namespace_example"; // String | Filters results based on the specified namespace.
    String xNomadToken = "xNomadToken_example"; // String | A Nomad ACL token.
    String idempotencyToken = "idempotencyToken_example"; // String | Can be used to ensure operations are only run once.
    try {
      apiInstance.putSystemGC(region, namespace, xNomadToken, idempotencyToken);
    } catch (ApiException e) {
      System.err.println("Exception when calling SystemApi#putSystemGC");
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
**200** |  |  -  |
**400** | Bad request |  -  |
**403** | Forbidden |  -  |
**405** | Method not allowed |  -  |
**500** | Internal server error |  -  |

<a name="putSystemReconcileSummaries"></a>
# **putSystemReconcileSummaries**
> putSystemReconcileSummaries(region, namespace, xNomadToken, idempotencyToken)



### Example
```java
// Import classes:
import io.nomadproject.client.ApiClient;
import io.nomadproject.client.ApiException;
import io.nomadproject.client.Configuration;
import io.nomadproject.client.auth.*;
import io.nomadproject.client.models.*;
import io.nomadproject.client.api.SystemApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://127.0.0.1:4646/v1");
    
    // Configure API key authorization: X-Nomad-Token
    ApiKeyAuth X-Nomad-Token = (ApiKeyAuth) defaultClient.getAuthentication("X-Nomad-Token");
    X-Nomad-Token.setApiKey("YOUR API KEY");
    // Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
    //X-Nomad-Token.setApiKeyPrefix("Token");

    SystemApi apiInstance = new SystemApi(defaultClient);
    String region = "region_example"; // String | Filters results based on the specified region.
    String namespace = "namespace_example"; // String | Filters results based on the specified namespace.
    String xNomadToken = "xNomadToken_example"; // String | A Nomad ACL token.
    String idempotencyToken = "idempotencyToken_example"; // String | Can be used to ensure operations are only run once.
    try {
      apiInstance.putSystemReconcileSummaries(region, namespace, xNomadToken, idempotencyToken);
    } catch (ApiException e) {
      System.err.println("Exception when calling SystemApi#putSystemReconcileSummaries");
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
**200** |  |  -  |
**400** | Bad request |  -  |
**403** | Forbidden |  -  |
**405** | Method not allowed |  -  |
**500** | Internal server error |  -  |

