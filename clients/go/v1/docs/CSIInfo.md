# CSIInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllocID** | Pointer to **string** |  | [optional] 
**ControllerInfo** | Pointer to [**NullableCSIControllerInfo**](CSIControllerInfo.md) |  | [optional] 
**HealthDescription** | Pointer to **string** |  | [optional] 
**Healthy** | Pointer to **bool** |  | [optional] 
**NodeInfo** | Pointer to [**NullableCSINodeInfo**](CSINodeInfo.md) |  | [optional] 
**PluginID** | Pointer to **string** |  | [optional] 
**RequiresControllerPlugin** | Pointer to **bool** |  | [optional] 
**RequiresTopologies** | Pointer to **bool** |  | [optional] 
**UpdateTime** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewCSIInfo

`func NewCSIInfo() *CSIInfo`

NewCSIInfo instantiates a new CSIInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSIInfoWithDefaults

`func NewCSIInfoWithDefaults() *CSIInfo`

NewCSIInfoWithDefaults instantiates a new CSIInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllocID

`func (o *CSIInfo) GetAllocID() string`

GetAllocID returns the AllocID field if non-nil, zero value otherwise.

### GetAllocIDOk

`func (o *CSIInfo) GetAllocIDOk() (*string, bool)`

GetAllocIDOk returns a tuple with the AllocID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocID

`func (o *CSIInfo) SetAllocID(v string)`

SetAllocID sets AllocID field to given value.

### HasAllocID

`func (o *CSIInfo) HasAllocID() bool`

HasAllocID returns a boolean if a field has been set.

### GetControllerInfo

`func (o *CSIInfo) GetControllerInfo() CSIControllerInfo`

GetControllerInfo returns the ControllerInfo field if non-nil, zero value otherwise.

### GetControllerInfoOk

`func (o *CSIInfo) GetControllerInfoOk() (*CSIControllerInfo, bool)`

GetControllerInfoOk returns a tuple with the ControllerInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerInfo

`func (o *CSIInfo) SetControllerInfo(v CSIControllerInfo)`

SetControllerInfo sets ControllerInfo field to given value.

### HasControllerInfo

`func (o *CSIInfo) HasControllerInfo() bool`

HasControllerInfo returns a boolean if a field has been set.

### SetControllerInfoNil

`func (o *CSIInfo) SetControllerInfoNil(b bool)`

 SetControllerInfoNil sets the value for ControllerInfo to be an explicit nil

### UnsetControllerInfo
`func (o *CSIInfo) UnsetControllerInfo()`

UnsetControllerInfo ensures that no value is present for ControllerInfo, not even an explicit nil
### GetHealthDescription

`func (o *CSIInfo) GetHealthDescription() string`

GetHealthDescription returns the HealthDescription field if non-nil, zero value otherwise.

### GetHealthDescriptionOk

`func (o *CSIInfo) GetHealthDescriptionOk() (*string, bool)`

GetHealthDescriptionOk returns a tuple with the HealthDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthDescription

`func (o *CSIInfo) SetHealthDescription(v string)`

SetHealthDescription sets HealthDescription field to given value.

### HasHealthDescription

`func (o *CSIInfo) HasHealthDescription() bool`

HasHealthDescription returns a boolean if a field has been set.

### GetHealthy

`func (o *CSIInfo) GetHealthy() bool`

GetHealthy returns the Healthy field if non-nil, zero value otherwise.

### GetHealthyOk

`func (o *CSIInfo) GetHealthyOk() (*bool, bool)`

GetHealthyOk returns a tuple with the Healthy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthy

`func (o *CSIInfo) SetHealthy(v bool)`

SetHealthy sets Healthy field to given value.

### HasHealthy

`func (o *CSIInfo) HasHealthy() bool`

HasHealthy returns a boolean if a field has been set.

### GetNodeInfo

`func (o *CSIInfo) GetNodeInfo() CSINodeInfo`

GetNodeInfo returns the NodeInfo field if non-nil, zero value otherwise.

### GetNodeInfoOk

`func (o *CSIInfo) GetNodeInfoOk() (*CSINodeInfo, bool)`

GetNodeInfoOk returns a tuple with the NodeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeInfo

`func (o *CSIInfo) SetNodeInfo(v CSINodeInfo)`

SetNodeInfo sets NodeInfo field to given value.

### HasNodeInfo

`func (o *CSIInfo) HasNodeInfo() bool`

HasNodeInfo returns a boolean if a field has been set.

### SetNodeInfoNil

`func (o *CSIInfo) SetNodeInfoNil(b bool)`

 SetNodeInfoNil sets the value for NodeInfo to be an explicit nil

### UnsetNodeInfo
`func (o *CSIInfo) UnsetNodeInfo()`

UnsetNodeInfo ensures that no value is present for NodeInfo, not even an explicit nil
### GetPluginID

`func (o *CSIInfo) GetPluginID() string`

GetPluginID returns the PluginID field if non-nil, zero value otherwise.

### GetPluginIDOk

`func (o *CSIInfo) GetPluginIDOk() (*string, bool)`

GetPluginIDOk returns a tuple with the PluginID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginID

`func (o *CSIInfo) SetPluginID(v string)`

SetPluginID sets PluginID field to given value.

### HasPluginID

`func (o *CSIInfo) HasPluginID() bool`

HasPluginID returns a boolean if a field has been set.

### GetRequiresControllerPlugin

`func (o *CSIInfo) GetRequiresControllerPlugin() bool`

GetRequiresControllerPlugin returns the RequiresControllerPlugin field if non-nil, zero value otherwise.

### GetRequiresControllerPluginOk

`func (o *CSIInfo) GetRequiresControllerPluginOk() (*bool, bool)`

GetRequiresControllerPluginOk returns a tuple with the RequiresControllerPlugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresControllerPlugin

`func (o *CSIInfo) SetRequiresControllerPlugin(v bool)`

SetRequiresControllerPlugin sets RequiresControllerPlugin field to given value.

### HasRequiresControllerPlugin

`func (o *CSIInfo) HasRequiresControllerPlugin() bool`

HasRequiresControllerPlugin returns a boolean if a field has been set.

### GetRequiresTopologies

`func (o *CSIInfo) GetRequiresTopologies() bool`

GetRequiresTopologies returns the RequiresTopologies field if non-nil, zero value otherwise.

### GetRequiresTopologiesOk

`func (o *CSIInfo) GetRequiresTopologiesOk() (*bool, bool)`

GetRequiresTopologiesOk returns a tuple with the RequiresTopologies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresTopologies

`func (o *CSIInfo) SetRequiresTopologies(v bool)`

SetRequiresTopologies sets RequiresTopologies field to given value.

### HasRequiresTopologies

`func (o *CSIInfo) HasRequiresTopologies() bool`

HasRequiresTopologies returns a boolean if a field has been set.

### GetUpdateTime

`func (o *CSIInfo) GetUpdateTime() time.Time`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *CSIInfo) GetUpdateTimeOk() (*time.Time, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *CSIInfo) SetUpdateTime(v time.Time)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *CSIInfo) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### SetUpdateTimeNil

`func (o *CSIInfo) SetUpdateTimeNil(b bool)`

 SetUpdateTimeNil sets the value for UpdateTime to be an explicit nil

### UnsetUpdateTime
`func (o *CSIInfo) UnsetUpdateTime()`

UnsetUpdateTime ensures that no value is present for UpdateTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


