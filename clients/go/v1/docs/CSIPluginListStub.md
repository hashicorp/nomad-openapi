# CSIPluginListStub

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ControllerRequired** | Pointer to **bool** |  | [optional] 
**ControllersExpected** | Pointer to **int32** |  | [optional] 
**ControllersHealthy** | Pointer to **int32** |  | [optional] 
**CreateIndex** | Pointer to **int32** |  | [optional] 
**ID** | Pointer to **string** |  | [optional] 
**ModifyIndex** | Pointer to **int32** |  | [optional] 
**NodesExpected** | Pointer to **int32** |  | [optional] 
**NodesHealthy** | Pointer to **int32** |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 

## Methods

### NewCSIPluginListStub

`func NewCSIPluginListStub() *CSIPluginListStub`

NewCSIPluginListStub instantiates a new CSIPluginListStub object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSIPluginListStubWithDefaults

`func NewCSIPluginListStubWithDefaults() *CSIPluginListStub`

NewCSIPluginListStubWithDefaults instantiates a new CSIPluginListStub object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetControllerRequired

`func (o *CSIPluginListStub) GetControllerRequired() bool`

GetControllerRequired returns the ControllerRequired field if non-nil, zero value otherwise.

### GetControllerRequiredOk

`func (o *CSIPluginListStub) GetControllerRequiredOk() (*bool, bool)`

GetControllerRequiredOk returns a tuple with the ControllerRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerRequired

`func (o *CSIPluginListStub) SetControllerRequired(v bool)`

SetControllerRequired sets ControllerRequired field to given value.

### HasControllerRequired

`func (o *CSIPluginListStub) HasControllerRequired() bool`

HasControllerRequired returns a boolean if a field has been set.

### GetControllersExpected

`func (o *CSIPluginListStub) GetControllersExpected() int32`

GetControllersExpected returns the ControllersExpected field if non-nil, zero value otherwise.

### GetControllersExpectedOk

`func (o *CSIPluginListStub) GetControllersExpectedOk() (*int32, bool)`

GetControllersExpectedOk returns a tuple with the ControllersExpected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllersExpected

`func (o *CSIPluginListStub) SetControllersExpected(v int32)`

SetControllersExpected sets ControllersExpected field to given value.

### HasControllersExpected

`func (o *CSIPluginListStub) HasControllersExpected() bool`

HasControllersExpected returns a boolean if a field has been set.

### GetControllersHealthy

`func (o *CSIPluginListStub) GetControllersHealthy() int32`

GetControllersHealthy returns the ControllersHealthy field if non-nil, zero value otherwise.

### GetControllersHealthyOk

`func (o *CSIPluginListStub) GetControllersHealthyOk() (*int32, bool)`

GetControllersHealthyOk returns a tuple with the ControllersHealthy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllersHealthy

`func (o *CSIPluginListStub) SetControllersHealthy(v int32)`

SetControllersHealthy sets ControllersHealthy field to given value.

### HasControllersHealthy

`func (o *CSIPluginListStub) HasControllersHealthy() bool`

HasControllersHealthy returns a boolean if a field has been set.

### GetCreateIndex

`func (o *CSIPluginListStub) GetCreateIndex() int32`

GetCreateIndex returns the CreateIndex field if non-nil, zero value otherwise.

### GetCreateIndexOk

`func (o *CSIPluginListStub) GetCreateIndexOk() (*int32, bool)`

GetCreateIndexOk returns a tuple with the CreateIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateIndex

`func (o *CSIPluginListStub) SetCreateIndex(v int32)`

SetCreateIndex sets CreateIndex field to given value.

### HasCreateIndex

`func (o *CSIPluginListStub) HasCreateIndex() bool`

HasCreateIndex returns a boolean if a field has been set.

### GetID

`func (o *CSIPluginListStub) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *CSIPluginListStub) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *CSIPluginListStub) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *CSIPluginListStub) HasID() bool`

HasID returns a boolean if a field has been set.

### GetModifyIndex

`func (o *CSIPluginListStub) GetModifyIndex() int32`

GetModifyIndex returns the ModifyIndex field if non-nil, zero value otherwise.

### GetModifyIndexOk

`func (o *CSIPluginListStub) GetModifyIndexOk() (*int32, bool)`

GetModifyIndexOk returns a tuple with the ModifyIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifyIndex

`func (o *CSIPluginListStub) SetModifyIndex(v int32)`

SetModifyIndex sets ModifyIndex field to given value.

### HasModifyIndex

`func (o *CSIPluginListStub) HasModifyIndex() bool`

HasModifyIndex returns a boolean if a field has been set.

### GetNodesExpected

`func (o *CSIPluginListStub) GetNodesExpected() int32`

GetNodesExpected returns the NodesExpected field if non-nil, zero value otherwise.

### GetNodesExpectedOk

`func (o *CSIPluginListStub) GetNodesExpectedOk() (*int32, bool)`

GetNodesExpectedOk returns a tuple with the NodesExpected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodesExpected

`func (o *CSIPluginListStub) SetNodesExpected(v int32)`

SetNodesExpected sets NodesExpected field to given value.

### HasNodesExpected

`func (o *CSIPluginListStub) HasNodesExpected() bool`

HasNodesExpected returns a boolean if a field has been set.

### GetNodesHealthy

`func (o *CSIPluginListStub) GetNodesHealthy() int32`

GetNodesHealthy returns the NodesHealthy field if non-nil, zero value otherwise.

### GetNodesHealthyOk

`func (o *CSIPluginListStub) GetNodesHealthyOk() (*int32, bool)`

GetNodesHealthyOk returns a tuple with the NodesHealthy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodesHealthy

`func (o *CSIPluginListStub) SetNodesHealthy(v int32)`

SetNodesHealthy sets NodesHealthy field to given value.

### HasNodesHealthy

`func (o *CSIPluginListStub) HasNodesHealthy() bool`

HasNodesHealthy returns a boolean if a field has been set.

### GetProvider

`func (o *CSIPluginListStub) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CSIPluginListStub) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CSIPluginListStub) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *CSIPluginListStub) HasProvider() bool`

HasProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


