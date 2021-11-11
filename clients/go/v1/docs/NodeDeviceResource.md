# NodeDeviceResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**map[string]Attribute**](Attribute.md) |  | [optional] 
**Instances** | Pointer to [**[]NodeDevice**](NodeDevice.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Vendor** | Pointer to **string** |  | [optional] 

## Methods

### NewNodeDeviceResource

`func NewNodeDeviceResource() *NodeDeviceResource`

NewNodeDeviceResource instantiates a new NodeDeviceResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeDeviceResourceWithDefaults

`func NewNodeDeviceResourceWithDefaults() *NodeDeviceResource`

NewNodeDeviceResourceWithDefaults instantiates a new NodeDeviceResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *NodeDeviceResource) GetAttributes() map[string]Attribute`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *NodeDeviceResource) GetAttributesOk() (*map[string]Attribute, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *NodeDeviceResource) SetAttributes(v map[string]Attribute)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *NodeDeviceResource) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetInstances

`func (o *NodeDeviceResource) GetInstances() []NodeDevice`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *NodeDeviceResource) GetInstancesOk() (*[]NodeDevice, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *NodeDeviceResource) SetInstances(v []NodeDevice)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *NodeDeviceResource) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### GetName

`func (o *NodeDeviceResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NodeDeviceResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NodeDeviceResource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NodeDeviceResource) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *NodeDeviceResource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NodeDeviceResource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NodeDeviceResource) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NodeDeviceResource) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVendor

`func (o *NodeDeviceResource) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *NodeDeviceResource) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *NodeDeviceResource) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *NodeDeviceResource) HasVendor() bool`

HasVendor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


