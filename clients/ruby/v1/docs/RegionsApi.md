# NomadClient::RegionsApi

All URIs are relative to *https://127.0.0.1:4646/v1*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**get_regions**](RegionsApi.md#get_regions) | **GET** /regions |  |


## get_regions

> Array&lt;String&gt; get_regions



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

api_instance = NomadClient::RegionsApi.new

begin
  
  result = api_instance.get_regions
  p result
rescue NomadClient::ApiError => e
  puts "Error when calling RegionsApi->get_regions: #{e}"
end
```

#### Using the get_regions_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(Array&lt;String&gt;, Integer, Hash)> get_regions_with_http_info

```ruby
begin
  
  data, status_code, headers = api_instance.get_regions_with_http_info
  p status_code # => 2xx
  p headers # => { ... }
  p data # => Array&lt;String&gt;
rescue NomadClient::ApiError => e
  puts "Error when calling RegionsApi->get_regions_with_http_info: #{e}"
end
```

### Parameters

This endpoint does not need any parameter.

### Return type

**Array&lt;String&gt;**

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

