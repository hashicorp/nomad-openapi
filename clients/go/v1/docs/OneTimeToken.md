# OneTimeToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessorID** | Pointer to **string** |  | [optional] 
**CreateIndex** | Pointer to **int32** |  | [optional] 
**ExpiresAt** | Pointer to **NullableTime** |  | [optional] 
**ModifyIndex** | Pointer to **int32** |  | [optional] 
**OneTimeSecretID** | Pointer to **string** |  | [optional] 

## Methods

### NewOneTimeToken

`func NewOneTimeToken() *OneTimeToken`

NewOneTimeToken instantiates a new OneTimeToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOneTimeTokenWithDefaults

`func NewOneTimeTokenWithDefaults() *OneTimeToken`

NewOneTimeTokenWithDefaults instantiates a new OneTimeToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessorID

`func (o *OneTimeToken) GetAccessorID() string`

GetAccessorID returns the AccessorID field if non-nil, zero value otherwise.

### GetAccessorIDOk

`func (o *OneTimeToken) GetAccessorIDOk() (*string, bool)`

GetAccessorIDOk returns a tuple with the AccessorID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessorID

`func (o *OneTimeToken) SetAccessorID(v string)`

SetAccessorID sets AccessorID field to given value.

### HasAccessorID

`func (o *OneTimeToken) HasAccessorID() bool`

HasAccessorID returns a boolean if a field has been set.

### GetCreateIndex

`func (o *OneTimeToken) GetCreateIndex() int32`

GetCreateIndex returns the CreateIndex field if non-nil, zero value otherwise.

### GetCreateIndexOk

`func (o *OneTimeToken) GetCreateIndexOk() (*int32, bool)`

GetCreateIndexOk returns a tuple with the CreateIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateIndex

`func (o *OneTimeToken) SetCreateIndex(v int32)`

SetCreateIndex sets CreateIndex field to given value.

### HasCreateIndex

`func (o *OneTimeToken) HasCreateIndex() bool`

HasCreateIndex returns a boolean if a field has been set.

### GetExpiresAt

`func (o *OneTimeToken) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *OneTimeToken) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *OneTimeToken) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *OneTimeToken) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *OneTimeToken) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *OneTimeToken) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
### GetModifyIndex

`func (o *OneTimeToken) GetModifyIndex() int32`

GetModifyIndex returns the ModifyIndex field if non-nil, zero value otherwise.

### GetModifyIndexOk

`func (o *OneTimeToken) GetModifyIndexOk() (*int32, bool)`

GetModifyIndexOk returns a tuple with the ModifyIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifyIndex

`func (o *OneTimeToken) SetModifyIndex(v int32)`

SetModifyIndex sets ModifyIndex field to given value.

### HasModifyIndex

`func (o *OneTimeToken) HasModifyIndex() bool`

HasModifyIndex returns a boolean if a field has been set.

### GetOneTimeSecretID

`func (o *OneTimeToken) GetOneTimeSecretID() string`

GetOneTimeSecretID returns the OneTimeSecretID field if non-nil, zero value otherwise.

### GetOneTimeSecretIDOk

`func (o *OneTimeToken) GetOneTimeSecretIDOk() (*string, bool)`

GetOneTimeSecretIDOk returns a tuple with the OneTimeSecretID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOneTimeSecretID

`func (o *OneTimeToken) SetOneTimeSecretID(v string)`

SetOneTimeSecretID sets OneTimeSecretID field to given value.

### HasOneTimeSecretID

`func (o *OneTimeToken) HasOneTimeSecretID() bool`

HasOneTimeSecretID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


