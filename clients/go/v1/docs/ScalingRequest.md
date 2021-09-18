# ScalingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int64** |  | [optional] 
**Error** | Pointer to **bool** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to **map[string]interface{}** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**PolicyOverride** | Pointer to **bool** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**SecretID** | Pointer to **string** |  | [optional] 
**Target** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewScalingRequest

`func NewScalingRequest() *ScalingRequest`

NewScalingRequest instantiates a new ScalingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScalingRequestWithDefaults

`func NewScalingRequestWithDefaults() *ScalingRequest`

NewScalingRequestWithDefaults instantiates a new ScalingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *ScalingRequest) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ScalingRequest) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ScalingRequest) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *ScalingRequest) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetError

`func (o *ScalingRequest) GetError() bool`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ScalingRequest) GetErrorOk() (*bool, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ScalingRequest) SetError(v bool)`

SetError sets Error field to given value.

### HasError

`func (o *ScalingRequest) HasError() bool`

HasError returns a boolean if a field has been set.

### GetMessage

`func (o *ScalingRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ScalingRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ScalingRequest) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ScalingRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetMeta

`func (o *ScalingRequest) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ScalingRequest) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ScalingRequest) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ScalingRequest) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetNamespace

`func (o *ScalingRequest) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ScalingRequest) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ScalingRequest) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *ScalingRequest) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetPolicyOverride

`func (o *ScalingRequest) GetPolicyOverride() bool`

GetPolicyOverride returns the PolicyOverride field if non-nil, zero value otherwise.

### GetPolicyOverrideOk

`func (o *ScalingRequest) GetPolicyOverrideOk() (*bool, bool)`

GetPolicyOverrideOk returns a tuple with the PolicyOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyOverride

`func (o *ScalingRequest) SetPolicyOverride(v bool)`

SetPolicyOverride sets PolicyOverride field to given value.

### HasPolicyOverride

`func (o *ScalingRequest) HasPolicyOverride() bool`

HasPolicyOverride returns a boolean if a field has been set.

### GetRegion

`func (o *ScalingRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ScalingRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ScalingRequest) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ScalingRequest) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSecretID

`func (o *ScalingRequest) GetSecretID() string`

GetSecretID returns the SecretID field if non-nil, zero value otherwise.

### GetSecretIDOk

`func (o *ScalingRequest) GetSecretIDOk() (*string, bool)`

GetSecretIDOk returns a tuple with the SecretID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretID

`func (o *ScalingRequest) SetSecretID(v string)`

SetSecretID sets SecretID field to given value.

### HasSecretID

`func (o *ScalingRequest) HasSecretID() bool`

HasSecretID returns a boolean if a field has been set.

### GetTarget

`func (o *ScalingRequest) GetTarget() map[string]string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *ScalingRequest) GetTargetOk() (*map[string]string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *ScalingRequest) SetTarget(v map[string]string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *ScalingRequest) HasTarget() bool`

HasTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


