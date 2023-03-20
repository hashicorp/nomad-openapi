# JobDispatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdPrefixTemplate** | Pointer to **string** |  | [optional] 
**JobID** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to **map[string]string** |  | [optional] 
**Payload** | Pointer to **string** |  | [optional] 

## Methods

### NewJobDispatchRequest

`func NewJobDispatchRequest() *JobDispatchRequest`

NewJobDispatchRequest instantiates a new JobDispatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobDispatchRequestWithDefaults

`func NewJobDispatchRequestWithDefaults() *JobDispatchRequest`

NewJobDispatchRequestWithDefaults instantiates a new JobDispatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdPrefixTemplate

`func (o *JobDispatchRequest) GetIdPrefixTemplate() string`

GetIdPrefixTemplate returns the IdPrefixTemplate field if non-nil, zero value otherwise.

### GetIdPrefixTemplateOk

`func (o *JobDispatchRequest) GetIdPrefixTemplateOk() (*string, bool)`

GetIdPrefixTemplateOk returns a tuple with the IdPrefixTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdPrefixTemplate

`func (o *JobDispatchRequest) SetIdPrefixTemplate(v string)`

SetIdPrefixTemplate sets IdPrefixTemplate field to given value.

### HasIdPrefixTemplate

`func (o *JobDispatchRequest) HasIdPrefixTemplate() bool`

HasIdPrefixTemplate returns a boolean if a field has been set.

### GetJobID

`func (o *JobDispatchRequest) GetJobID() string`

GetJobID returns the JobID field if non-nil, zero value otherwise.

### GetJobIDOk

`func (o *JobDispatchRequest) GetJobIDOk() (*string, bool)`

GetJobIDOk returns a tuple with the JobID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobID

`func (o *JobDispatchRequest) SetJobID(v string)`

SetJobID sets JobID field to given value.

### HasJobID

`func (o *JobDispatchRequest) HasJobID() bool`

HasJobID returns a boolean if a field has been set.

### GetMeta

`func (o *JobDispatchRequest) GetMeta() map[string]string`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *JobDispatchRequest) GetMetaOk() (*map[string]string, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *JobDispatchRequest) SetMeta(v map[string]string)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *JobDispatchRequest) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetPayload

`func (o *JobDispatchRequest) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *JobDispatchRequest) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *JobDispatchRequest) SetPayload(v string)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *JobDispatchRequest) HasPayload() bool`

HasPayload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


