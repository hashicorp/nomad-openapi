# \OperatorApi

All URIs are relative to *https://127.0.0.1:4646/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**delete_operator_raft**](OperatorApi.md#delete_operator_raft) | **Delete** /operator/raft/ | 
[**get_operator_autopilot_configuration**](OperatorApi.md#get_operator_autopilot_configuration) | **Get** /operator/autopilot/configuration | 
[**get_operator_autopilot_health**](OperatorApi.md#get_operator_autopilot_health) | **Get** /operator/autopilot/health | 
[**get_operator_raft**](OperatorApi.md#get_operator_raft) | **Get** /operator/raft/ | 
[**get_operator_scheduler_configuration**](OperatorApi.md#get_operator_scheduler_configuration) | **Get** /operator/scheduler/configuration | 
[**post_operator_scheduler_configuration**](OperatorApi.md#post_operator_scheduler_configuration) | **Post** /operator/scheduler/configuration | 
[**put_operator_autopilot_configuration**](OperatorApi.md#put_operator_autopilot_configuration) | **Put** /operator/autopilot/configuration | 



## delete_operator_raft

> delete_operator_raft(region, namespace, x_nomad_token, idempotency_token)


### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**region** | Option<**String**> | Filters results based on the specified region. |  |
**namespace** | Option<**String**> | Filters results based on the specified namespace. |  |
**x_nomad_token** | Option<**String**> | A Nomad ACL token. |  |
**idempotency_token** | Option<**String**> | Can be used to ensure operations are only run once. |  |

### Return type

 (empty response body)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## get_operator_autopilot_configuration

> crate::models::AutopilotConfiguration get_operator_autopilot_configuration(region, namespace, index, wait, stale, prefix, x_nomad_token, per_page, next_token)


### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**region** | Option<**String**> | Filters results based on the specified region. |  |
**namespace** | Option<**String**> | Filters results based on the specified namespace. |  |
**index** | Option<**i32**> | If set, wait until query exceeds given index. Must be provided with WaitParam. |  |
**wait** | Option<**String**> | Provided with IndexParam to wait for change. |  |
**stale** | Option<**String**> | If present, results will include stale reads. |  |
**prefix** | Option<**String**> | Constrains results to jobs that start with the defined prefix |  |
**x_nomad_token** | Option<**String**> | A Nomad ACL token. |  |
**per_page** | Option<**i32**> | Maximum number of results to return. |  |
**next_token** | Option<**String**> | Indicates where to start paging for queries that support pagination. |  |

### Return type

[**crate::models::AutopilotConfiguration**](AutopilotConfiguration.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## get_operator_autopilot_health

> crate::models::OperatorHealthReply get_operator_autopilot_health(region, namespace, index, wait, stale, prefix, x_nomad_token, per_page, next_token)


### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**region** | Option<**String**> | Filters results based on the specified region. |  |
**namespace** | Option<**String**> | Filters results based on the specified namespace. |  |
**index** | Option<**i32**> | If set, wait until query exceeds given index. Must be provided with WaitParam. |  |
**wait** | Option<**String**> | Provided with IndexParam to wait for change. |  |
**stale** | Option<**String**> | If present, results will include stale reads. |  |
**prefix** | Option<**String**> | Constrains results to jobs that start with the defined prefix |  |
**x_nomad_token** | Option<**String**> | A Nomad ACL token. |  |
**per_page** | Option<**i32**> | Maximum number of results to return. |  |
**next_token** | Option<**String**> | Indicates where to start paging for queries that support pagination. |  |

### Return type

[**crate::models::OperatorHealthReply**](OperatorHealthReply.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## get_operator_raft

> Vec<crate::models::RaftServer> get_operator_raft(region, namespace, index, wait, stale, prefix, x_nomad_token, per_page, next_token)


### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**region** | Option<**String**> | Filters results based on the specified region. |  |
**namespace** | Option<**String**> | Filters results based on the specified namespace. |  |
**index** | Option<**i32**> | If set, wait until query exceeds given index. Must be provided with WaitParam. |  |
**wait** | Option<**String**> | Provided with IndexParam to wait for change. |  |
**stale** | Option<**String**> | If present, results will include stale reads. |  |
**prefix** | Option<**String**> | Constrains results to jobs that start with the defined prefix |  |
**x_nomad_token** | Option<**String**> | A Nomad ACL token. |  |
**per_page** | Option<**i32**> | Maximum number of results to return. |  |
**next_token** | Option<**String**> | Indicates where to start paging for queries that support pagination. |  |

### Return type

[**Vec<crate::models::RaftServer>**](RaftServer.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## get_operator_scheduler_configuration

> crate::models::SchedulerConfigurationResponse get_operator_scheduler_configuration(region, namespace, index, wait, stale, prefix, x_nomad_token, per_page, next_token)


### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**region** | Option<**String**> | Filters results based on the specified region. |  |
**namespace** | Option<**String**> | Filters results based on the specified namespace. |  |
**index** | Option<**i32**> | If set, wait until query exceeds given index. Must be provided with WaitParam. |  |
**wait** | Option<**String**> | Provided with IndexParam to wait for change. |  |
**stale** | Option<**String**> | If present, results will include stale reads. |  |
**prefix** | Option<**String**> | Constrains results to jobs that start with the defined prefix |  |
**x_nomad_token** | Option<**String**> | A Nomad ACL token. |  |
**per_page** | Option<**i32**> | Maximum number of results to return. |  |
**next_token** | Option<**String**> | Indicates where to start paging for queries that support pagination. |  |

### Return type

[**crate::models::SchedulerConfigurationResponse**](SchedulerConfigurationResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## post_operator_scheduler_configuration

> crate::models::SchedulerSetConfigurationResponse post_operator_scheduler_configuration(scheduler_configuration, region, namespace, x_nomad_token, idempotency_token)


### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**scheduler_configuration** | [**SchedulerConfiguration**](SchedulerConfiguration.md) |  | [required] |
**region** | Option<**String**> | Filters results based on the specified region. |  |
**namespace** | Option<**String**> | Filters results based on the specified namespace. |  |
**x_nomad_token** | Option<**String**> | A Nomad ACL token. |  |
**idempotency_token** | Option<**String**> | Can be used to ensure operations are only run once. |  |

### Return type

[**crate::models::SchedulerSetConfigurationResponse**](SchedulerSetConfigurationResponse.md)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## put_operator_autopilot_configuration

> put_operator_autopilot_configuration(autopilot_configuration, region, namespace, x_nomad_token, idempotency_token)


### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**autopilot_configuration** | [**AutopilotConfiguration**](AutopilotConfiguration.md) |  | [required] |
**region** | Option<**String**> | Filters results based on the specified region. |  |
**namespace** | Option<**String**> | Filters results based on the specified namespace. |  |
**x_nomad_token** | Option<**String**> | A Nomad ACL token. |  |
**idempotency_token** | Option<**String**> | Can be used to ensure operations are only run once. |  |

### Return type

 (empty response body)

### Authorization

[X-Nomad-Token](../README.md#X-Nomad-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

