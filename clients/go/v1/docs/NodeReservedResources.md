# NodeReservedResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpu** | Pointer to [**NodeReservedCpuResources**](NodeReservedCpuResources.md) |  | [optional] 
**Disk** | Pointer to [**NodeReservedDiskResources**](NodeReservedDiskResources.md) |  | [optional] 
**Memory** | Pointer to [**NodeReservedMemoryResources**](NodeReservedMemoryResources.md) |  | [optional] 
**Networks** | Pointer to [**NodeReservedNetworkResources**](NodeReservedNetworkResources.md) |  | [optional] 

## Methods

### NewNodeReservedResources

`func NewNodeReservedResources() *NodeReservedResources`

NewNodeReservedResources instantiates a new NodeReservedResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeReservedResourcesWithDefaults

`func NewNodeReservedResourcesWithDefaults() *NodeReservedResources`

NewNodeReservedResourcesWithDefaults instantiates a new NodeReservedResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpu

`func (o *NodeReservedResources) GetCpu() NodeReservedCpuResources`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *NodeReservedResources) GetCpuOk() (*NodeReservedCpuResources, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *NodeReservedResources) SetCpu(v NodeReservedCpuResources)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *NodeReservedResources) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetDisk

`func (o *NodeReservedResources) GetDisk() NodeReservedDiskResources`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *NodeReservedResources) GetDiskOk() (*NodeReservedDiskResources, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *NodeReservedResources) SetDisk(v NodeReservedDiskResources)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *NodeReservedResources) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### GetMemory

`func (o *NodeReservedResources) GetMemory() NodeReservedMemoryResources`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *NodeReservedResources) GetMemoryOk() (*NodeReservedMemoryResources, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *NodeReservedResources) SetMemory(v NodeReservedMemoryResources)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *NodeReservedResources) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetNetworks

`func (o *NodeReservedResources) GetNetworks() NodeReservedNetworkResources`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *NodeReservedResources) GetNetworksOk() (*NodeReservedNetworkResources, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *NodeReservedResources) SetNetworks(v NodeReservedNetworkResources)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *NodeReservedResources) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


