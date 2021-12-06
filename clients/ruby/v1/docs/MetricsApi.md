# NomadClient::MetricsApi

All URIs are relative to *https://127.0.0.1:4646/v1*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**get_metrics_summary**](MetricsApi.md#get_metrics_summary) | **GET** /metrics |  |


## get_metrics_summary

> <MetricsSummary> get_metrics_summary(opts)



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

api_instance = NomadClient::MetricsApi.new
opts = {
  format: 'format_example' # String | The format the user requested for the metrics summary (e.g. prometheus)
}

begin
  
  result = api_instance.get_metrics_summary(opts)
  p result
rescue NomadClient::ApiError => e
  puts "Error when calling MetricsApi->get_metrics_summary: #{e}"
end
```

#### Using the get_metrics_summary_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<MetricsSummary>, Integer, Hash)> get_metrics_summary_with_http_info(opts)

```ruby
begin
  
  data, status_code, headers = api_instance.get_metrics_summary_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <MetricsSummary>
rescue NomadClient::ApiError => e
  puts "Error when calling MetricsApi->get_metrics_summary_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **format** | **String** | The format the user requested for the metrics summary (e.g. prometheus) | [optional] |

### Return type

[**MetricsSummary**](MetricsSummary.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

