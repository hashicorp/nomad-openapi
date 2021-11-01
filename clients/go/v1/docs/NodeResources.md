# NodeResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpu** | Pointer to [**NodeCpuResources**](NodeCpuResources.md) |  | [optional] 
**Devices** | Pointer to [**[]NodeDeviceResource**](NodeDeviceResource.md) |  | [optional] 
**Disk** | Pointer to [**NodeDiskResources**](NodeDiskResources.md) |  | [optional] 
**Memory** | Pointer to [**NodeMemoryResources**](NodeMemoryResources.md) |  | [optional] 
**Networks** | Pointer to [**[]NetworkResource**](NetworkResource.md) |  | [optional] 

## Methods

### NewNodeResources

`func NewNodeResources() *NodeResources`

NewNodeResources instantiates a new NodeResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeResourcesWithDefaults

`func NewNodeResourcesWithDefaults() *NodeResources`

NewNodeResourcesWithDefaults instantiates a new NodeResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpu

`func (o *NodeResources) GetCpu() NodeCpuResources`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *NodeResources) GetCpuOk() (*NodeCpuResources, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *NodeResources) SetCpu(v NodeCpuResources)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *NodeResources) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetDevices

`func (o *NodeResources) GetDevices() []NodeDeviceResource`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *NodeResources) GetDevicesOk() (*[]NodeDeviceResource, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *NodeResources) SetDevices(v []NodeDeviceResource)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *NodeResources) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### GetDisk

`func (o *NodeResources) GetDisk() NodeDiskResources`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *NodeResources) GetDiskOk() (*NodeDiskResources, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *NodeResources) SetDisk(v NodeDiskResources)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *NodeResources) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### GetMemory

`func (o *NodeResources) GetMemory() NodeMemoryResources`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *NodeResources) GetMemoryOk() (*NodeMemoryResources, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *NodeResources) SetMemory(v NodeMemoryResources)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *NodeResources) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetNetworks

`func (o *NodeResources) GetNetworks() []NetworkResource`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *NodeResources) GetNetworksOk() (*[]NetworkResource, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *NodeResources) SetNetworks(v []NetworkResource)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *NodeResources) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


