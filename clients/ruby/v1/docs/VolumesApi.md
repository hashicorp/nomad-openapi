# NomadClient::VolumesApi

All URIs are relative to *http://127.0.0.1:4646/v1*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**create_volume**](VolumesApi.md#create_volume) | **POST** /volume/csi/{volumeId}/{action} |  |
| [**delete_snapshot**](VolumesApi.md#delete_snapshot) | **DELETE** /volumes/snapshot |  |
| [**delete_volume_registration**](VolumesApi.md#delete_volume_registration) | **DELETE** /volume/csi/{volumeId} |  |
| [**detach_or_delete_volume**](VolumesApi.md#detach_or_delete_volume) | **DELETE** /volume/csi/{volumeId}/{action} |  |
| [**get_external_volumes**](VolumesApi.md#get_external_volumes) | **GET** /volumes/external |  |
| [**get_snapshots**](VolumesApi.md#get_snapshots) | **GET** /volumes/snapshot |  |
| [**get_volume**](VolumesApi.md#get_volume) | **GET** /volume/csi/{volumeId} |  |
| [**get_volumes**](VolumesApi.md#get_volumes) | **GET** /volumes |  |
| [**post_snapshot**](VolumesApi.md#post_snapshot) | **POST** /volumes/snapshot |  |
| [**post_volume**](VolumesApi.md#post_volume) | **POST** /volumes |  |
| [**post_volume_registration**](VolumesApi.md#post_volume_registration) | **POST** /volume/csi/{volumeId} |  |


## create_volume

> create_volume(volume_id, action, csi_volume_create_request, opts)



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

api_instance = NomadClient::VolumesApi.new
volume_id = 'volume_id_example' # String | Volume unique identifier.
action = 'action_example' # String | The action to perform on the Volume (create, detach, delete).
csi_volume_create_request = NomadClient::CSIVolumeCreateRequest.new # CSIVolumeCreateRequest | 
opts = {
  region: 'region_example', # String | Filters results based on the specified region.
  namespace: 'namespace_example', # String | Filters results based on the specified namespace.
  x_nomad_token: 'x_nomad_token_example', # String | A Nomad ACL token.
  idempotency_token: 'idempotency_token_example' # String | Can be used to ensure operations are only run once.
}

begin
  
  api_instance.create_volume(volume_id, action, csi_volume_create_request, opts)
rescue NomadClient::ApiError => e
  puts "Error when calling VolumesApi->create_volume: #{e}"
end
```

#### Using the create_volume_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> create_volume_with_http_info(volume_id, action, csi_volume_create_request, opts)

```ruby
begin
  
  data, status_code, headers = api_instance.create_volume_with_http_info(volume_id, action, csi_volume_create_request, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue NomadClient::ApiError => e
  puts "Error when calling VolumesApi->create_volume_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **volume_id** | **String** | Volume unique identifier. |  |
| **action** | **String** | The action to perform on the Volume (create, detach, delete). |  |
| **csi_volume_create_request** | [**CSIVolumeCreateRequest**](CSIVolumeCreateRequest.md) |  |  |
| **region** | **String** | Filters results based on the specified region. | [optional] |
| **namespace** | **String** | Filters results based on the specified namespace. | [optional] |
| **x_nomad_token** | **String** | A Nomad ACL token. | [optional] |
| **idempotency_token** | **String** | Can be used to ensure operations are only run once. | [optional] |

### Return type

nil (empty response body)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined


## delete_snapshot

> delete_snapshot(opts)



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

api_instance = NomadClient::VolumesApi.new
opts = {
  region: 'region_example', # String | Filters results based on the specified region.
  namespace: 'namespace_example', # String | Filters results based on the specified namespace.
  x_nomad_token: 'x_nomad_token_example', # String | A Nomad ACL token.
  idempotency_token: 'idempotency_token_example', # String | Can be used to ensure operations are only run once.
  plugin_id: 'plugin_id_example', # String | Filters volume lists by plugin ID.
  snapshot_id: 'snapshot_id_example' # String | The ID of the snapshot to target.
}

begin
  
  api_instance.delete_snapshot(opts)
rescue NomadClient::ApiError => e
  puts "Error when calling VolumesApi->delete_snapshot: #{e}"
end
```

#### Using the delete_snapshot_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> delete_snapshot_with_http_info(opts)

```ruby
begin
  
  data, status_code, headers = api_instance.delete_snapshot_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue NomadClient::ApiError => e
  puts "Error when calling VolumesApi->delete_snapshot_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **region** | **String** | Filters results based on the specified region. | [optional] |
| **namespace** | **String** | Filters results based on the specified namespace. | [optional] |
| **x_nomad_token** | **String** | A Nomad ACL token. | [optional] |
| **idempotency_token** | **String** | Can be used to ensure operations are only run once. | [optional] |
| **plugin_id** | **String** | Filters volume lists by plugin ID. | [optional] |
| **snapshot_id** | **String** | The ID of the snapshot to target. | [optional] |

### Return type

nil (empty response body)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined


## delete_volume_registration

> delete_volume_registration(volume_id, opts)



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

api_instance = NomadClient::VolumesApi.new
volume_id = 'volume_id_example' # String | Volume unique identifier.
opts = {
  region: 'region_example', # String | Filters results based on the specified region.
  namespace: 'namespace_example', # String | Filters results based on the specified namespace.
  x_nomad_token: 'x_nomad_token_example', # String | A Nomad ACL token.
  idempotency_token: 'idempotency_token_example', # String | Can be used to ensure operations are only run once.
  force: 'force_example' # String | Used to force the de-registration of a volume.
}

begin
  
  api_instance.delete_volume_registration(volume_id, opts)
rescue NomadClient::ApiError => e
  puts "Error when calling VolumesApi->delete_volume_registration: #{e}"
end
```

#### Using the delete_volume_registration_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> delete_volume_registration_with_http_info(volume_id, opts)

```ruby
begin
  
  data, status_code, headers = api_instance.delete_volume_registration_with_http_info(volume_id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue NomadClient::ApiError => e
  puts "Error when calling VolumesApi->delete_volume_registration_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **volume_id** | **String** | Volume unique identifier. |  |
| **region** | **String** | Filters results based on the specified region. | [optional] |
| **namespace** | **String** | Filters results based on the specified namespace. | [optional] |
| **x_nomad_token** | **String** | A Nomad ACL token. | [optional] |
| **idempotency_token** | **String** | Can be used to ensure operations are only run once. | [optional] |
| **force** | **String** | Used to force the de-registration of a volume. | [optional] |

### Return type

nil (empty response body)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined


## detach_or_delete_volume

> detach_or_delete_volume(volume_id, action, opts)



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

api_instance = NomadClient::VolumesApi.new
volume_id = 'volume_id_example' # String | Volume unique identifier.
action = 'action_example' # String | The action to perform on the Volume (create, detach, delete).
opts = {
  region: 'region_example', # String | Filters results based on the specified region.
  namespace: 'namespace_example', # String | Filters results based on the specified namespace.
  x_nomad_token: 'x_nomad_token_example', # String | A Nomad ACL token.
  idempotency_token: 'idempotency_token_example', # String | Can be used to ensure operations are only run once.
  node: 'node_example' # String | Specifies node to target volume operation for.
}

begin
  
  api_instance.detach_or_delete_volume(volume_id, action, opts)
rescue NomadClient::ApiError => e
  puts "Error when calling VolumesApi->detach_or_delete_volume: #{e}"
end
```

#### Using the detach_or_delete_volume_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> detach_or_delete_volume_with_http_info(volume_id, action, opts)

```ruby
begin
  
  data, status_code, headers = api_instance.detach_or_delete_volume_with_http_info(volume_id, action, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue NomadClient::ApiError => e
  puts "Error when calling VolumesApi->detach_or_delete_volume_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **volume_id** | **String** | Volume unique identifier. |  |
| **action** | **String** | The action to perform on the Volume (create, detach, delete). |  |
| **region** | **String** | Filters results based on the specified region. | [optional] |
| **namespace** | **String** | Filters results based on the specified namespace. | [optional] |
| **x_nomad_token** | **String** | A Nomad ACL token. | [optional] |
| **idempotency_token** | **String** | Can be used to ensure operations are only run once. | [optional] |
| **node** | **String** | Specifies node to target volume operation for. | [optional] |

### Return type

nil (empty response body)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined


## get_external_volumes

> <CSIVolumeListExternalResponse> get_external_volumes(opts)



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

api_instance = NomadClient::VolumesApi.new
opts = {
  region: 'region_example', # String | Filters results based on the specified region.
  namespace: 'namespace_example', # String | Filters results based on the specified namespace.
  index: 56, # Integer | If set, wait until query exceeds given index. Must be provided with WaitParam.
  wait: 'wait_example', # String | Provided with IndexParam to wait for change.
  stale: 'stale_example', # String | If present, results will include stale reads.
  prefix: 'prefix_example', # String | Constrains results to jobs that start with the defined prefix
  x_nomad_token: 'x_nomad_token_example', # String | A Nomad ACL token.
  per_page: 56, # Integer | Maximum number of results to return.
  next_token: 'next_token_example', # String | Indicates where to start paging for queries that support pagination.
  plugin_id: 'plugin_id_example' # String | Filters volume lists by plugin ID.
}

begin
  
  result = api_instance.get_external_volumes(opts)
  p result
rescue NomadClient::ApiError => e
  puts "Error when calling VolumesApi->get_external_volumes: #{e}"
end
```

#### Using the get_external_volumes_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<CSIVolumeListExternalResponse>, Integer, Hash)> get_external_volumes_with_http_info(opts)

```ruby
begin
  
  data, status_code, headers = api_instance.get_external_volumes_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <CSIVolumeListExternalResponse>
rescue NomadClient::ApiError => e
  puts "Error when calling VolumesApi->get_external_volumes_with_http_info: #{e}"
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
| **plugin_id** | **String** | Filters volume lists by plugin ID. | [optional] |

### Return type

[**CSIVolumeListExternalResponse**](CSIVolumeListExternalResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## get_snapshots

> <CSISnapshotListResponse> get_snapshots(opts)



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

api_instance = NomadClient::VolumesApi.new
opts = {
  region: 'region_example', # String | Filters results based on the specified region.
  namespace: 'namespace_example', # String | Filters results based on the specified namespace.
  index: 56, # Integer | If set, wait until query exceeds given index. Must be provided with WaitParam.
  wait: 'wait_example', # String | Provided with IndexParam to wait for change.
  stale: 'stale_example', # String | If present, results will include stale reads.
  prefix: 'prefix_example', # String | Constrains results to jobs that start with the defined prefix
  x_nomad_token: 'x_nomad_token_example', # String | A Nomad ACL token.
  per_page: 56, # Integer | Maximum number of results to return.
  next_token: 'next_token_example', # String | Indicates where to start paging for queries that support pagination.
  plugin_id: 'plugin_id_example' # String | Filters volume lists by plugin ID.
}

begin
  
  result = api_instance.get_snapshots(opts)
  p result
rescue NomadClient::ApiError => e
  puts "Error when calling VolumesApi->get_snapshots: #{e}"
end
```

#### Using the get_snapshots_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<CSISnapshotListResponse>, Integer, Hash)> get_snapshots_with_http_info(opts)

```ruby
begin
  
  data, status_code, headers = api_instance.get_snapshots_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <CSISnapshotListResponse>
rescue NomadClient::ApiError => e
  puts "Error when calling VolumesApi->get_snapshots_with_http_info: #{e}"
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
| **plugin_id** | **String** | Filters volume lists by plugin ID. | [optional] |

### Return type

[**CSISnapshotListResponse**](CSISnapshotListResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## get_volume

> <CSIVolume> get_volume(volume_id, opts)



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

api_instance = NomadClient::VolumesApi.new
volume_id = 'volume_id_example' # String | Volume unique identifier.
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
  
  result = api_instance.get_volume(volume_id, opts)
  p result
rescue NomadClient::ApiError => e
  puts "Error when calling VolumesApi->get_volume: #{e}"
end
```

#### Using the get_volume_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<CSIVolume>, Integer, Hash)> get_volume_with_http_info(volume_id, opts)

```ruby
begin
  
  data, status_code, headers = api_instance.get_volume_with_http_info(volume_id, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <CSIVolume>
rescue NomadClient::ApiError => e
  puts "Error when calling VolumesApi->get_volume_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **volume_id** | **String** | Volume unique identifier. |  |
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

[**CSIVolume**](CSIVolume.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## get_volumes

> <Array<CSIVolumeListStub>> get_volumes(opts)



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

api_instance = NomadClient::VolumesApi.new
opts = {
  region: 'region_example', # String | Filters results based on the specified region.
  namespace: 'namespace_example', # String | Filters results based on the specified namespace.
  index: 56, # Integer | If set, wait until query exceeds given index. Must be provided with WaitParam.
  wait: 'wait_example', # String | Provided with IndexParam to wait for change.
  stale: 'stale_example', # String | If present, results will include stale reads.
  prefix: 'prefix_example', # String | Constrains results to jobs that start with the defined prefix
  x_nomad_token: 'x_nomad_token_example', # String | A Nomad ACL token.
  per_page: 56, # Integer | Maximum number of results to return.
  next_token: 'next_token_example', # String | Indicates where to start paging for queries that support pagination.
  node_id: 'node_id_example', # String | Filters volume lists by node ID.
  plugin_id: 'plugin_id_example', # String | Filters volume lists by plugin ID.
  type: 'type_example' # String | Filters volume lists to a specific type.
}

begin
  
  result = api_instance.get_volumes(opts)
  p result
rescue NomadClient::ApiError => e
  puts "Error when calling VolumesApi->get_volumes: #{e}"
end
```

#### Using the get_volumes_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Array<CSIVolumeListStub>>, Integer, Hash)> get_volumes_with_http_info(opts)

```ruby
begin
  
  data, status_code, headers = api_instance.get_volumes_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Array<CSIVolumeListStub>>
rescue NomadClient::ApiError => e
  puts "Error when calling VolumesApi->get_volumes_with_http_info: #{e}"
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
| **node_id** | **String** | Filters volume lists by node ID. | [optional] |
| **plugin_id** | **String** | Filters volume lists by plugin ID. | [optional] |
| **type** | **String** | Filters volume lists to a specific type. | [optional] |

### Return type

[**Array&lt;CSIVolumeListStub&gt;**](CSIVolumeListStub.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## post_snapshot

> <CSISnapshotCreateResponse> post_snapshot(csi_snapshot_create_request, opts)



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

api_instance = NomadClient::VolumesApi.new
csi_snapshot_create_request = NomadClient::CSISnapshotCreateRequest.new # CSISnapshotCreateRequest | 
opts = {
  region: 'region_example', # String | Filters results based on the specified region.
  namespace: 'namespace_example', # String | Filters results based on the specified namespace.
  x_nomad_token: 'x_nomad_token_example', # String | A Nomad ACL token.
  idempotency_token: 'idempotency_token_example' # String | Can be used to ensure operations are only run once.
}

begin
  
  result = api_instance.post_snapshot(csi_snapshot_create_request, opts)
  p result
rescue NomadClient::ApiError => e
  puts "Error when calling VolumesApi->post_snapshot: #{e}"
end
```

#### Using the post_snapshot_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<CSISnapshotCreateResponse>, Integer, Hash)> post_snapshot_with_http_info(csi_snapshot_create_request, opts)

```ruby
begin
  
  data, status_code, headers = api_instance.post_snapshot_with_http_info(csi_snapshot_create_request, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <CSISnapshotCreateResponse>
rescue NomadClient::ApiError => e
  puts "Error when calling VolumesApi->post_snapshot_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **csi_snapshot_create_request** | [**CSISnapshotCreateRequest**](CSISnapshotCreateRequest.md) |  |  |
| **region** | **String** | Filters results based on the specified region. | [optional] |
| **namespace** | **String** | Filters results based on the specified namespace. | [optional] |
| **x_nomad_token** | **String** | A Nomad ACL token. | [optional] |
| **idempotency_token** | **String** | Can be used to ensure operations are only run once. | [optional] |

### Return type

[**CSISnapshotCreateResponse**](CSISnapshotCreateResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## post_volume

> post_volume(csi_volume_register_request, opts)



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

api_instance = NomadClient::VolumesApi.new
csi_volume_register_request = NomadClient::CSIVolumeRegisterRequest.new # CSIVolumeRegisterRequest | 
opts = {
  region: 'region_example', # String | Filters results based on the specified region.
  namespace: 'namespace_example', # String | Filters results based on the specified namespace.
  x_nomad_token: 'x_nomad_token_example', # String | A Nomad ACL token.
  idempotency_token: 'idempotency_token_example' # String | Can be used to ensure operations are only run once.
}

begin
  
  api_instance.post_volume(csi_volume_register_request, opts)
rescue NomadClient::ApiError => e
  puts "Error when calling VolumesApi->post_volume: #{e}"
end
```

#### Using the post_volume_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> post_volume_with_http_info(csi_volume_register_request, opts)

```ruby
begin
  
  data, status_code, headers = api_instance.post_volume_with_http_info(csi_volume_register_request, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue NomadClient::ApiError => e
  puts "Error when calling VolumesApi->post_volume_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **csi_volume_register_request** | [**CSIVolumeRegisterRequest**](CSIVolumeRegisterRequest.md) |  |  |
| **region** | **String** | Filters results based on the specified region. | [optional] |
| **namespace** | **String** | Filters results based on the specified namespace. | [optional] |
| **x_nomad_token** | **String** | A Nomad ACL token. | [optional] |
| **idempotency_token** | **String** | Can be used to ensure operations are only run once. | [optional] |

### Return type

nil (empty response body)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined


## post_volume_registration

> post_volume_registration(volume_id, csi_volume_register_request, opts)



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

api_instance = NomadClient::VolumesApi.new
volume_id = 'volume_id_example' # String | Volume unique identifier.
csi_volume_register_request = NomadClient::CSIVolumeRegisterRequest.new # CSIVolumeRegisterRequest | 
opts = {
  region: 'region_example', # String | Filters results based on the specified region.
  namespace: 'namespace_example', # String | Filters results based on the specified namespace.
  x_nomad_token: 'x_nomad_token_example', # String | A Nomad ACL token.
  idempotency_token: 'idempotency_token_example' # String | Can be used to ensure operations are only run once.
}

begin
  
  api_instance.post_volume_registration(volume_id, csi_volume_register_request, opts)
rescue NomadClient::ApiError => e
  puts "Error when calling VolumesApi->post_volume_registration: #{e}"
end
```

#### Using the post_volume_registration_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> post_volume_registration_with_http_info(volume_id, csi_volume_register_request, opts)

```ruby
begin
  
  data, status_code, headers = api_instance.post_volume_registration_with_http_info(volume_id, csi_volume_register_request, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue NomadClient::ApiError => e
  puts "Error when calling VolumesApi->post_volume_registration_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **volume_id** | **String** | Volume unique identifier. |  |
| **csi_volume_register_request** | [**CSIVolumeRegisterRequest**](CSIVolumeRegisterRequest.md) |  |  |
| **region** | **String** | Filters results based on the specified region. | [optional] |
| **namespace** | **String** | Filters results based on the specified namespace. | [optional] |
| **x_nomad_token** | **String** | A Nomad ACL token. | [optional] |
| **idempotency_token** | **String** | Can be used to ensure operations are only run once. | [optional] |

### Return type

nil (empty response body)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

