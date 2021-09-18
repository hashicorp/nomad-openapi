# QuotaSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateIndex** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Limits** | Pointer to [**[]QuotaLimit**](QuotaLimit.md) |  | [optional] 
**ModifyIndex** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewQuotaSpec

`func NewQuotaSpec() *QuotaSpec`

NewQuotaSpec instantiates a new QuotaSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuotaSpecWithDefaults

`func NewQuotaSpecWithDefaults() *QuotaSpec`

NewQuotaSpecWithDefaults instantiates a new QuotaSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateIndex

`func (o *QuotaSpec) GetCreateIndex() int32`

GetCreateIndex returns the CreateIndex field if non-nil, zero value otherwise.

### GetCreateIndexOk

`func (o *QuotaSpec) GetCreateIndexOk() (*int32, bool)`

GetCreateIndexOk returns a tuple with the CreateIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateIndex

`func (o *QuotaSpec) SetCreateIndex(v int32)`

SetCreateIndex sets CreateIndex field to given value.

### HasCreateIndex

`func (o *QuotaSpec) HasCreateIndex() bool`

HasCreateIndex returns a boolean if a field has been set.

### GetDescription

`func (o *QuotaSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *QuotaSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *QuotaSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *QuotaSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLimits

`func (o *QuotaSpec) GetLimits() []QuotaLimit`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *QuotaSpec) GetLimitsOk() (*[]QuotaLimit, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *QuotaSpec) SetLimits(v []QuotaLimit)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *QuotaSpec) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### GetModifyIndex

`func (o *QuotaSpec) GetModifyIndex() int32`

GetModifyIndex returns the ModifyIndex field if non-nil, zero value otherwise.

### GetModifyIndexOk

`func (o *QuotaSpec) GetModifyIndexOk() (*int32, bool)`

GetModifyIndexOk returns a tuple with the ModifyIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifyIndex

`func (o *QuotaSpec) SetModifyIndex(v int32)`

SetModifyIndex sets ModifyIndex field to given value.

### HasModifyIndex

`func (o *QuotaSpec) HasModifyIndex() bool`

HasModifyIndex returns a boolean if a field has been set.

### GetName

`func (o *QuotaSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QuotaSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QuotaSpec) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *QuotaSpec) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


