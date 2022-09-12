# NomadClient::SystemApi

All URIs are relative to *http://127.0.0.1:4646/v1*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**put_system_gc**](SystemApi.md#put_system_gc) | **PUT** /system/gc |  |
| [**put_system_reconcile_summaries**](SystemApi.md#put_system_reconcile_summaries) | **PUT** /system/reconcile/summaries |  |


## put_system_gc

> put_system_gc(opts)



### Examples

```ruby
require 'time'
require 'nomad_client'
# setup authorization
NomadClient.configure do |config|
  # Configure API key authorization: X-Nomad-Token
  config.api_key['X-Nomad-Token'] = 'YOUR API KEY'
  # Uncomment the following line to set a prefix for the API key, e.g. 'Bearer' (defaults to nil)
  # config.api_key_prefix['X-Nomad-Token'] = 'Bearer'
end

api_instance = NomadClient::SystemApi.new
opts = {
  region: 'region_example', # String | Filters results based on the specified region.
  namespace: 'namespace_example', # String | Filters results based on the specified namespace.
  x_nomad_token: 'x_nomad_token_example', # String | A Nomad ACL token.
  idempotency_token: 'idempotency_token_example' # String | Can be used to ensure operations are only run once.
}

begin
  
  api_instance.put_system_gc(opts)
rescue NomadClient::ApiError => e
  puts "Error when calling SystemApi->put_system_gc: #{e}"
end
```

#### Using the put_system_gc_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> put_system_gc_with_http_info(opts)

```ruby
begin
  
  data, status_code, headers = api_instance.put_system_gc_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue NomadClient::ApiError => e
  puts "Error when calling SystemApi->put_system_gc_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **region** | **String** | Filters results based on the specified region. | [optional] |
| **namespace** | **String** | Filters results based on the specified namespace. | [optional] |
| **x_nomad_token** | **String** | A Nomad ACL token. | [optional] |
| **idempotency_token** | **String** | Can be used to ensure operations are only run once. | [optional] |

### Return type

nil (empty response body)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined


## put_system_reconcile_summaries

> put_system_reconcile_summaries(opts)



### Examples

```ruby
require 'time'
require 'nomad_client'
# setup authorization
NomadClient.configure do |config|
  # Configure API key authorization: X-Nomad-Token
  config.api_key['X-Nomad-Token'] = 'YOUR API KEY'
  # Uncomment the following line to set a prefix for the API key, e.g. 'Bearer' (defaults to nil)
  # config.api_key_prefix['X-Nomad-Token'] = 'Bearer'
end

api_instance = NomadClient::SystemApi.new
opts = {
  region: 'region_example', # String | Filters results based on the specified region.
  namespace: 'namespace_example', # String | Filters results based on the specified namespace.
  x_nomad_token: 'x_nomad_token_example', # String | A Nomad ACL token.
  idempotency_token: 'idempotency_token_example' # String | Can be used to ensure operations are only run once.
}

begin
  
  api_instance.put_system_reconcile_summaries(opts)
rescue NomadClient::ApiError => e
  puts "Error when calling SystemApi->put_system_reconcile_summaries: #{e}"
end
```

#### Using the put_system_reconcile_summaries_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> put_system_reconcile_summaries_with_http_info(opts)

```ruby
begin
  
  data, status_code, headers = api_instance.put_system_reconcile_summaries_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue NomadClient::ApiError => e
  puts "Error when calling SystemApi->put_system_reconcile_summaries_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **region** | **String** | Filters results based on the specified region. | [optional] |
| **namespace** | **String** | Filters results based on the specified namespace. | [optional] |
| **x_nomad_token** | **String** | A Nomad ACL token. | [optional] |
| **idempotency_token** | **String** | Can be used to ensure operations are only run once. | [optional] |

### Return type

nil (empty response body)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

