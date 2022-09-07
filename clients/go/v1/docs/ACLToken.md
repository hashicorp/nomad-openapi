# ACLToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessorID** | Pointer to **string** |  | [optional] 
**CreateIndex** | Pointer to **int32** |  | [optional] 
**CreateTime** | Pointer to **NullableTime** |  | [optional] 
**Global** | Pointer to **bool** |  | [optional] 
**ModifyIndex** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Policies** | Pointer to **[]string** |  | [optional] 
**SecretID** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewACLToken

`func NewACLToken() *ACLToken`

NewACLToken instantiates a new ACLToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewACLTokenWithDefaults

`func NewACLTokenWithDefaults() *ACLToken`

NewACLTokenWithDefaults instantiates a new ACLToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessorID

`func (o *ACLToken) GetAccessorID() string`

GetAccessorID returns the AccessorID field if non-nil, zero value otherwise.

### GetAccessorIDOk

`func (o *ACLToken) GetAccessorIDOk() (*string, bool)`

GetAccessorIDOk returns a tuple with the AccessorID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessorID

`func (o *ACLToken) SetAccessorID(v string)`

SetAccessorID sets AccessorID field to given value.

### HasAccessorID

`func (o *ACLToken) HasAccessorID() bool`

HasAccessorID returns a boolean if a field has been set.

### GetCreateIndex

`func (o *ACLToken) GetCreateIndex() int32`

GetCreateIndex returns the CreateIndex field if non-nil, zero value otherwise.

### GetCreateIndexOk

`func (o *ACLToken) GetCreateIndexOk() (*int32, bool)`

GetCreateIndexOk returns a tuple with the CreateIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateIndex

`func (o *ACLToken) SetCreateIndex(v int32)`

SetCreateIndex sets CreateIndex field to given value.

### HasCreateIndex

`func (o *ACLToken) HasCreateIndex() bool`

HasCreateIndex returns a boolean if a field has been set.

### GetCreateTime

`func (o *ACLToken) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *ACLToken) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *ACLToken) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *ACLToken) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### SetCreateTimeNil

`func (o *ACLToken) SetCreateTimeNil(b bool)`

 SetCreateTimeNil sets the value for CreateTime to be an explicit nil

### UnsetCreateTime
`func (o *ACLToken) UnsetCreateTime()`

UnsetCreateTime ensures that no value is present for CreateTime, not even an explicit nil
### GetGlobal

`func (o *ACLToken) GetGlobal() bool`

GetGlobal returns the Global field if non-nil, zero value otherwise.

### GetGlobalOk

`func (o *ACLToken) GetGlobalOk() (*bool, bool)`

GetGlobalOk returns a tuple with the Global field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobal

`func (o *ACLToken) SetGlobal(v bool)`

SetGlobal sets Global field to given value.

### HasGlobal

`func (o *ACLToken) HasGlobal() bool`

HasGlobal returns a boolean if a field has been set.

### GetModifyIndex

`func (o *ACLToken) GetModifyIndex() int32`

GetModifyIndex returns the ModifyIndex field if non-nil, zero value otherwise.

### GetModifyIndexOk

`func (o *ACLToken) GetModifyIndexOk() (*int32, bool)`

GetModifyIndexOk returns a tuple with the ModifyIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifyIndex

`func (o *ACLToken) SetModifyIndex(v int32)`

SetModifyIndex sets ModifyIndex field to given value.

### HasModifyIndex

`func (o *ACLToken) HasModifyIndex() bool`

HasModifyIndex returns a boolean if a field has been set.

### GetName

`func (o *ACLToken) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ACLToken) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ACLToken) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ACLToken) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPolicies

`func (o *ACLToken) GetPolicies() []string`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *ACLToken) GetPoliciesOk() (*[]string, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *ACLToken) SetPolicies(v []string)`

SetPolicies sets Policies field to given value.

### HasPolicies

`func (o *ACLToken) HasPolicies() bool`

HasPolicies returns a boolean if a field has been set.

### GetSecretID

`func (o *ACLToken) GetSecretID() string`

GetSecretID returns the SecretID field if non-nil, zero value otherwise.

### GetSecretIDOk

`func (o *ACLToken) GetSecretIDOk() (*string, bool)`

GetSecretIDOk returns a tuple with the SecretID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretID

`func (o *ACLToken) SetSecretID(v string)`

SetSecretID sets SecretID field to given value.

### HasSecretID

`func (o *ACLToken) HasSecretID() bool`

HasSecretID returns a boolean if a field has been set.

### GetType

`func (o *ACLToken) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ACLToken) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ACLToken) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ACLToken) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


