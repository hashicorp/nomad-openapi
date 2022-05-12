# ServiceRegistration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** |  | [optional] 
**AllocID** | Pointer to **string** |  | [optional] 
**CreateIndex** | Pointer to **int32** |  | [optional] 
**Datacenter** | Pointer to **string** |  | [optional] 
**ID** | Pointer to **string** |  | [optional] 
**JobID** | Pointer to **string** |  | [optional] 
**ModifyIndex** | Pointer to **int32** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**NodeID** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**ServiceName** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewServiceRegistration

`func NewServiceRegistration() *ServiceRegistration`

NewServiceRegistration instantiates a new ServiceRegistration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceRegistrationWithDefaults

`func NewServiceRegistrationWithDefaults() *ServiceRegistration`

NewServiceRegistrationWithDefaults instantiates a new ServiceRegistration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *ServiceRegistration) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ServiceRegistration) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ServiceRegistration) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ServiceRegistration) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAllocID

`func (o *ServiceRegistration) GetAllocID() string`

GetAllocID returns the AllocID field if non-nil, zero value otherwise.

### GetAllocIDOk

`func (o *ServiceRegistration) GetAllocIDOk() (*string, bool)`

GetAllocIDOk returns a tuple with the AllocID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocID

`func (o *ServiceRegistration) SetAllocID(v string)`

SetAllocID sets AllocID field to given value.

### HasAllocID

`func (o *ServiceRegistration) HasAllocID() bool`

HasAllocID returns a boolean if a field has been set.

### GetCreateIndex

`func (o *ServiceRegistration) GetCreateIndex() int32`

GetCreateIndex returns the CreateIndex field if non-nil, zero value otherwise.

### GetCreateIndexOk

`func (o *ServiceRegistration) GetCreateIndexOk() (*int32, bool)`

GetCreateIndexOk returns a tuple with the CreateIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateIndex

`func (o *ServiceRegistration) SetCreateIndex(v int32)`

SetCreateIndex sets CreateIndex field to given value.

### HasCreateIndex

`func (o *ServiceRegistration) HasCreateIndex() bool`

HasCreateIndex returns a boolean if a field has been set.

### GetDatacenter

`func (o *ServiceRegistration) GetDatacenter() string`

GetDatacenter returns the Datacenter field if non-nil, zero value otherwise.

### GetDatacenterOk

`func (o *ServiceRegistration) GetDatacenterOk() (*string, bool)`

GetDatacenterOk returns a tuple with the Datacenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenter

`func (o *ServiceRegistration) SetDatacenter(v string)`

SetDatacenter sets Datacenter field to given value.

### HasDatacenter

`func (o *ServiceRegistration) HasDatacenter() bool`

HasDatacenter returns a boolean if a field has been set.

### GetID

`func (o *ServiceRegistration) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *ServiceRegistration) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *ServiceRegistration) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *ServiceRegistration) HasID() bool`

HasID returns a boolean if a field has been set.

### GetJobID

`func (o *ServiceRegistration) GetJobID() string`

GetJobID returns the JobID field if non-nil, zero value otherwise.

### GetJobIDOk

`func (o *ServiceRegistration) GetJobIDOk() (*string, bool)`

GetJobIDOk returns a tuple with the JobID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobID

`func (o *ServiceRegistration) SetJobID(v string)`

SetJobID sets JobID field to given value.

### HasJobID

`func (o *ServiceRegistration) HasJobID() bool`

HasJobID returns a boolean if a field has been set.

### GetModifyIndex

`func (o *ServiceRegistration) GetModifyIndex() int32`

GetModifyIndex returns the ModifyIndex field if non-nil, zero value otherwise.

### GetModifyIndexOk

`func (o *ServiceRegistration) GetModifyIndexOk() (*int32, bool)`

GetModifyIndexOk returns a tuple with the ModifyIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifyIndex

`func (o *ServiceRegistration) SetModifyIndex(v int32)`

SetModifyIndex sets ModifyIndex field to given value.

### HasModifyIndex

`func (o *ServiceRegistration) HasModifyIndex() bool`

HasModifyIndex returns a boolean if a field has been set.

### GetNamespace

`func (o *ServiceRegistration) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ServiceRegistration) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ServiceRegistration) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *ServiceRegistration) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetNodeID

`func (o *ServiceRegistration) GetNodeID() string`

GetNodeID returns the NodeID field if non-nil, zero value otherwise.

### GetNodeIDOk

`func (o *ServiceRegistration) GetNodeIDOk() (*string, bool)`

GetNodeIDOk returns a tuple with the NodeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeID

`func (o *ServiceRegistration) SetNodeID(v string)`

SetNodeID sets NodeID field to given value.

### HasNodeID

`func (o *ServiceRegistration) HasNodeID() bool`

HasNodeID returns a boolean if a field has been set.

### GetPort

`func (o *ServiceRegistration) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ServiceRegistration) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ServiceRegistration) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *ServiceRegistration) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetServiceName

`func (o *ServiceRegistration) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *ServiceRegistration) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *ServiceRegistration) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *ServiceRegistration) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetTags

`func (o *ServiceRegistration) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ServiceRegistration) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ServiceRegistration) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ServiceRegistration) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


