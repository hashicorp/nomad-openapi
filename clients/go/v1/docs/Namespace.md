# Namespace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capabilities** | Pointer to [**NullableNamespaceCapabilities**](NamespaceCapabilities.md) |  | [optional] 
**CreateIndex** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to **map[string]string** |  | [optional] 
**ModifyIndex** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Quota** | Pointer to **string** |  | [optional] 

## Methods

### NewNamespace

`func NewNamespace() *Namespace`

NewNamespace instantiates a new Namespace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNamespaceWithDefaults

`func NewNamespaceWithDefaults() *Namespace`

NewNamespaceWithDefaults instantiates a new Namespace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapabilities

`func (o *Namespace) GetCapabilities() NamespaceCapabilities`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *Namespace) GetCapabilitiesOk() (*NamespaceCapabilities, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *Namespace) SetCapabilities(v NamespaceCapabilities)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *Namespace) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### SetCapabilitiesNil

`func (o *Namespace) SetCapabilitiesNil(b bool)`

 SetCapabilitiesNil sets the value for Capabilities to be an explicit nil

### UnsetCapabilities
`func (o *Namespace) UnsetCapabilities()`

UnsetCapabilities ensures that no value is present for Capabilities, not even an explicit nil
### GetCreateIndex

`func (o *Namespace) GetCreateIndex() int32`

GetCreateIndex returns the CreateIndex field if non-nil, zero value otherwise.

### GetCreateIndexOk

`func (o *Namespace) GetCreateIndexOk() (*int32, bool)`

GetCreateIndexOk returns a tuple with the CreateIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateIndex

`func (o *Namespace) SetCreateIndex(v int32)`

SetCreateIndex sets CreateIndex field to given value.

### HasCreateIndex

`func (o *Namespace) HasCreateIndex() bool`

HasCreateIndex returns a boolean if a field has been set.

### GetDescription

`func (o *Namespace) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Namespace) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Namespace) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Namespace) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMeta

`func (o *Namespace) GetMeta() map[string]string`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *Namespace) GetMetaOk() (*map[string]string, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *Namespace) SetMeta(v map[string]string)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *Namespace) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetModifyIndex

`func (o *Namespace) GetModifyIndex() int32`

GetModifyIndex returns the ModifyIndex field if non-nil, zero value otherwise.

### GetModifyIndexOk

`func (o *Namespace) GetModifyIndexOk() (*int32, bool)`

GetModifyIndexOk returns a tuple with the ModifyIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifyIndex

`func (o *Namespace) SetModifyIndex(v int32)`

SetModifyIndex sets ModifyIndex field to given value.

### HasModifyIndex

`func (o *Namespace) HasModifyIndex() bool`

HasModifyIndex returns a boolean if a field has been set.

### GetName

`func (o *Namespace) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Namespace) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Namespace) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Namespace) HasName() bool`

HasName returns a boolean if a field has been set.

### GetQuota

`func (o *Namespace) GetQuota() string`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *Namespace) GetQuotaOk() (*string, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *Namespace) SetQuota(v string)`

SetQuota sets Quota field to given value.

### HasQuota

`func (o *Namespace) HasQuota() bool`

HasQuota returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


