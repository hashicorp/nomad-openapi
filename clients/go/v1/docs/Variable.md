# Variable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateIndex** | Pointer to **int32** |  | [optional] 
**CreateTime** | Pointer to **int64** |  | [optional] 
**Items** | Pointer to **map[string]string** |  | [optional] 
**ModifyIndex** | Pointer to **int32** |  | [optional] 
**ModifyTime** | Pointer to **int64** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 

## Methods

### NewVariable

`func NewVariable() *Variable`

NewVariable instantiates a new Variable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariableWithDefaults

`func NewVariableWithDefaults() *Variable`

NewVariableWithDefaults instantiates a new Variable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateIndex

`func (o *Variable) GetCreateIndex() int32`

GetCreateIndex returns the CreateIndex field if non-nil, zero value otherwise.

### GetCreateIndexOk

`func (o *Variable) GetCreateIndexOk() (*int32, bool)`

GetCreateIndexOk returns a tuple with the CreateIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateIndex

`func (o *Variable) SetCreateIndex(v int32)`

SetCreateIndex sets CreateIndex field to given value.

### HasCreateIndex

`func (o *Variable) HasCreateIndex() bool`

HasCreateIndex returns a boolean if a field has been set.

### GetCreateTime

`func (o *Variable) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *Variable) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *Variable) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *Variable) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetItems

`func (o *Variable) GetItems() map[string]string`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *Variable) GetItemsOk() (*map[string]string, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *Variable) SetItems(v map[string]string)`

SetItems sets Items field to given value.

### HasItems

`func (o *Variable) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetModifyIndex

`func (o *Variable) GetModifyIndex() int32`

GetModifyIndex returns the ModifyIndex field if non-nil, zero value otherwise.

### GetModifyIndexOk

`func (o *Variable) GetModifyIndexOk() (*int32, bool)`

GetModifyIndexOk returns a tuple with the ModifyIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifyIndex

`func (o *Variable) SetModifyIndex(v int32)`

SetModifyIndex sets ModifyIndex field to given value.

### HasModifyIndex

`func (o *Variable) HasModifyIndex() bool`

HasModifyIndex returns a boolean if a field has been set.

### GetModifyTime

`func (o *Variable) GetModifyTime() int64`

GetModifyTime returns the ModifyTime field if non-nil, zero value otherwise.

### GetModifyTimeOk

`func (o *Variable) GetModifyTimeOk() (*int64, bool)`

GetModifyTimeOk returns a tuple with the ModifyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifyTime

`func (o *Variable) SetModifyTime(v int64)`

SetModifyTime sets ModifyTime field to given value.

### HasModifyTime

`func (o *Variable) HasModifyTime() bool`

HasModifyTime returns a boolean if a field has been set.

### GetNamespace

`func (o *Variable) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *Variable) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *Variable) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *Variable) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetPath

`func (o *Variable) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *Variable) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *Variable) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *Variable) HasPath() bool`

HasPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


