# NomadClient::PluginsApi

All URIs are relative to *http://127.0.0.1:4646/v1*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**get_plugin_csi**](PluginsApi.md#get_plugin_csi) | **GET** /plugin/csi/{pluginID} |  |
| [**get_plugins**](PluginsApi.md#get_plugins) | **GET** /plugins |  |


## get_plugin_csi

> <Array<CSIPlugin>> get_plugin_csi(plugin_id, opts)



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

api_instance = NomadClient::PluginsApi.new
plugin_id = 'plugin_id_example' # String | The CSI plugin identifier.
opts = {
  region: 'region_example', # String | Filters results based on the specified region.
  namespace: 'namespace_example', # String | Filters results based on the specified namespace.
  index: 56, # Integer | If set, wait until query exceeds given index. Must be provided with WaitParam.
  wait: 'wait_example', # String | Provided with IndexParam to wait for change.
  stale: 'stale_example', # String | If present, results will include stale reads.
  prefix: 'prefix_example', # String | Constrains results to jobs that start with the defined prefix
  x_nomad_token: 'x_nomad_token_example', # String | A Nomad ACL token.
  per_page: 56, # Integer | Maximum number of results to return.
  next_token: 'next_token_example' # String | Indicates where to start paging for queries that support pagination.
}

begin
  
  result = api_instance.get_plugin_csi(plugin_id, opts)
  p result
rescue NomadClient::ApiError => e
  puts "Error when calling PluginsApi->get_plugin_csi: #{e}"
end
```

#### Using the get_plugin_csi_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Array<CSIPlugin>>, Integer, Hash)> get_plugin_csi_with_http_info(plugin_id, opts)

```ruby
begin
  
  data, status_code, headers = api_instance.get_plugin_csi_with_http_info(plugin_id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Array<CSIPlugin>>
rescue NomadClient::ApiError => e
  puts "Error when calling PluginsApi->get_plugin_csi_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **plugin_id** | **String** | The CSI plugin identifier. |  |
| **region** | **String** | Filters results based on the specified region. | [optional] |
| **namespace** | **String** | Filters results based on the specified namespace. | [optional] |
| **index** | **Integer** | If set, wait until query exceeds given index. Must be provided with WaitParam. | [optional] |
| **wait** | **String** | Provided with IndexParam to wait for change. | [optional] |
| **stale** | **String** | If present, results will include stale reads. | [optional] |
| **prefix** | **String** | Constrains results to jobs that start with the defined prefix | [optional] |
| **x_nomad_token** | **String** | A Nomad ACL token. | [optional] |
| **per_page** | **Integer** | Maximum number of results to return. | [optional] |
| **next_token** | **String** | Indicates where to start paging for queries that support pagination. | [optional] |

### Return type

[**Array&lt;CSIPlugin&gt;**](CSIPlugin.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## get_plugins

> <Array<CSIPluginListStub>> get_plugins(opts)



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

api_instance = NomadClient::PluginsApi.new
opts = {
  region: 'region_example', # String | Filters results based on the specified region.
  namespace: 'namespace_example', # String | Filters results based on the specified namespace.
  index: 56, # Integer | If set, wait until query exceeds given index. Must be provided with WaitParam.
  wait: 'wait_example', # String | Provided with IndexParam to wait for change.
  stale: 'stale_example', # String | If present, results will include stale reads.
  prefix: 'prefix_example', # String | Constrains results to jobs that start with the defined prefix
  x_nomad_token: 'x_nomad_token_example', # String | A Nomad ACL token.
  per_page: 56, # Integer | Maximum number of results to return.
  next_token: 'next_token_example' # String | Indicates where to start paging for queries that support pagination.
}

begin
  
  result = api_instance.get_plugins(opts)
  p result
rescue NomadClient::ApiError => e
  puts "Error when calling PluginsApi->get_plugins: #{e}"
end
```

#### Using the get_plugins_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Array<CSIPluginListStub>>, Integer, Hash)> get_plugins_with_http_info(opts)

```ruby
begin
  
  data, status_code, headers = api_instance.get_plugins_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Array<CSIPluginListStub>>
rescue NomadClient::ApiError => e
  puts "Error when calling PluginsApi->get_plugins_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **region** | **String** | Filters results based on the specified region. | [optional] |
| **namespace** | **String** | Filters results based on the specified namespace. | [optional] |
| **index** | **Integer** | If set, wait until query exceeds given index. Must be provided with WaitParam. | [optional] |
| **wait** | **String** | Provided with IndexParam to wait for change. | [optional] |
| **stale** | **String** | If present, results will include stale reads. | [optional] |
| **prefix** | **String** | Constrains results to jobs that start with the defined prefix | [optional] |
| **x_nomad_token** | **String** | A Nomad ACL token. | [optional] |
| **per_page** | **Integer** | Maximum number of results to return. | [optional] |
| **next_token** | **String** | Indicates where to start paging for queries that support pagination. | [optional] |

### Return type

[**Array&lt;CSIPluginListStub&gt;**](CSIPluginListStub.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

