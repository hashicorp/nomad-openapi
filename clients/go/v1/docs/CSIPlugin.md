# CSIPlugin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Allocations** | Pointer to [**[]AllocationListStub**](AllocationListStub.md) |  | [optional] 
**ControllerRequired** | Pointer to **bool** |  | [optional] 
**Controllers** | Pointer to [**map[string]CSIInfo**](CSIInfo.md) |  | [optional] 
**ControllersExpected** | Pointer to **int32** |  | [optional] 
**ControllersHealthy** | Pointer to **int32** |  | [optional] 
**CreateIndex** | Pointer to **int32** |  | [optional] 
**ID** | Pointer to **string** |  | [optional] 
**ModifyIndex** | Pointer to **int32** |  | [optional] 
**Nodes** | Pointer to [**map[string]CSIInfo**](CSIInfo.md) |  | [optional] 
**NodesExpected** | Pointer to **int32** |  | [optional] 
**NodesHealthy** | Pointer to **int32** |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewCSIPlugin

`func NewCSIPlugin() *CSIPlugin`

NewCSIPlugin instantiates a new CSIPlugin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSIPluginWithDefaults

`func NewCSIPluginWithDefaults() *CSIPlugin`

NewCSIPluginWithDefaults instantiates a new CSIPlugin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllocations

`func (o *CSIPlugin) GetAllocations() []AllocationListStub`

GetAllocations returns the Allocations field if non-nil, zero value otherwise.

### GetAllocationsOk

`func (o *CSIPlugin) GetAllocationsOk() (*[]AllocationListStub, bool)`

GetAllocationsOk returns a tuple with the Allocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocations

`func (o *CSIPlugin) SetAllocations(v []AllocationListStub)`

SetAllocations sets Allocations field to given value.

### HasAllocations

`func (o *CSIPlugin) HasAllocations() bool`

HasAllocations returns a boolean if a field has been set.

### GetControllerRequired

`func (o *CSIPlugin) GetControllerRequired() bool`

GetControllerRequired returns the ControllerRequired field if non-nil, zero value otherwise.

### GetControllerRequiredOk

`func (o *CSIPlugin) GetControllerRequiredOk() (*bool, bool)`

GetControllerRequiredOk returns a tuple with the ControllerRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerRequired

`func (o *CSIPlugin) SetControllerRequired(v bool)`

SetControllerRequired sets ControllerRequired field to given value.

### HasControllerRequired

`func (o *CSIPlugin) HasControllerRequired() bool`

HasControllerRequired returns a boolean if a field has been set.

### GetControllers

`func (o *CSIPlugin) GetControllers() map[string]CSIInfo`

GetControllers returns the Controllers field if non-nil, zero value otherwise.

### GetControllersOk

`func (o *CSIPlugin) GetControllersOk() (*map[string]CSIInfo, bool)`

GetControllersOk returns a tuple with the Controllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllers

`func (o *CSIPlugin) SetControllers(v map[string]CSIInfo)`

SetControllers sets Controllers field to given value.

### HasControllers

`func (o *CSIPlugin) HasControllers() bool`

HasControllers returns a boolean if a field has been set.

### GetControllersExpected

`func (o *CSIPlugin) GetControllersExpected() int32`

GetControllersExpected returns the ControllersExpected field if non-nil, zero value otherwise.

### GetControllersExpectedOk

`func (o *CSIPlugin) GetControllersExpectedOk() (*int32, bool)`

GetControllersExpectedOk returns a tuple with the ControllersExpected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllersExpected

`func (o *CSIPlugin) SetControllersExpected(v int32)`

SetControllersExpected sets ControllersExpected field to given value.

### HasControllersExpected

`func (o *CSIPlugin) HasControllersExpected() bool`

HasControllersExpected returns a boolean if a field has been set.

### GetControllersHealthy

`func (o *CSIPlugin) GetControllersHealthy() int32`

GetControllersHealthy returns the ControllersHealthy field if non-nil, zero value otherwise.

### GetControllersHealthyOk

`func (o *CSIPlugin) GetControllersHealthyOk() (*int32, bool)`

GetControllersHealthyOk returns a tuple with the ControllersHealthy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllersHealthy

`func (o *CSIPlugin) SetControllersHealthy(v int32)`

SetControllersHealthy sets ControllersHealthy field to given value.

### HasControllersHealthy

`func (o *CSIPlugin) HasControllersHealthy() bool`

HasControllersHealthy returns a boolean if a field has been set.

### GetCreateIndex

`func (o *CSIPlugin) GetCreateIndex() int32`

GetCreateIndex returns the CreateIndex field if non-nil, zero value otherwise.

### GetCreateIndexOk

`func (o *CSIPlugin) GetCreateIndexOk() (*int32, bool)`

GetCreateIndexOk returns a tuple with the CreateIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateIndex

`func (o *CSIPlugin) SetCreateIndex(v int32)`

SetCreateIndex sets CreateIndex field to given value.

### HasCreateIndex

`func (o *CSIPlugin) HasCreateIndex() bool`

HasCreateIndex returns a boolean if a field has been set.

### GetID

`func (o *CSIPlugin) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *CSIPlugin) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *CSIPlugin) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *CSIPlugin) HasID() bool`

HasID returns a boolean if a field has been set.

### GetModifyIndex

`func (o *CSIPlugin) GetModifyIndex() int32`

GetModifyIndex returns the ModifyIndex field if non-nil, zero value otherwise.

### GetModifyIndexOk

`func (o *CSIPlugin) GetModifyIndexOk() (*int32, bool)`

GetModifyIndexOk returns a tuple with the ModifyIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifyIndex

`func (o *CSIPlugin) SetModifyIndex(v int32)`

SetModifyIndex sets ModifyIndex field to given value.

### HasModifyIndex

`func (o *CSIPlugin) HasModifyIndex() bool`

HasModifyIndex returns a boolean if a field has been set.

### GetNodes

`func (o *CSIPlugin) GetNodes() map[string]CSIInfo`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *CSIPlugin) GetNodesOk() (*map[string]CSIInfo, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *CSIPlugin) SetNodes(v map[string]CSIInfo)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *CSIPlugin) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetNodesExpected

`func (o *CSIPlugin) GetNodesExpected() int32`

GetNodesExpected returns the NodesExpected field if non-nil, zero value otherwise.

### GetNodesExpectedOk

`func (o *CSIPlugin) GetNodesExpectedOk() (*int32, bool)`

GetNodesExpectedOk returns a tuple with the NodesExpected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodesExpected

`func (o *CSIPlugin) SetNodesExpected(v int32)`

SetNodesExpected sets NodesExpected field to given value.

### HasNodesExpected

`func (o *CSIPlugin) HasNodesExpected() bool`

HasNodesExpected returns a boolean if a field has been set.

### GetNodesHealthy

`func (o *CSIPlugin) GetNodesHealthy() int32`

GetNodesHealthy returns the NodesHealthy field if non-nil, zero value otherwise.

### GetNodesHealthyOk

`func (o *CSIPlugin) GetNodesHealthyOk() (*int32, bool)`

GetNodesHealthyOk returns a tuple with the NodesHealthy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodesHealthy

`func (o *CSIPlugin) SetNodesHealthy(v int32)`

SetNodesHealthy sets NodesHealthy field to given value.

### HasNodesHealthy

`func (o *CSIPlugin) HasNodesHealthy() bool`

HasNodesHealthy returns a boolean if a field has been set.

### GetProvider

`func (o *CSIPlugin) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CSIPlugin) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CSIPlugin) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *CSIPlugin) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetVersion

`func (o *CSIPlugin) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CSIPlugin) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CSIPlugin) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CSIPlugin) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


