/*
 * Nomad
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.1.4
 * Contact: support@hashicorp.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// NodeResources struct for NodeResources
type NodeResources struct {
	Cpu *NodeCpuResources `json:"Cpu,omitempty"`
	Devices *[]NodeDeviceResource `json:"Devices,omitempty"`
	Disk *NodeDiskResources `json:"Disk,omitempty"`
	MaxDynamicPort *int32 `json:"MaxDynamicPort,omitempty"`
	Memory *NodeMemoryResources `json:"Memory,omitempty"`
	MinDynamicPort *int32 `json:"MinDynamicPort,omitempty"`
	Networks *[]NetworkResource `json:"Networks,omitempty"`
}

// NewNodeResources instantiates a new NodeResources object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNodeResources() *NodeResources {
	this := NodeResources{}
	return &this
}

// NewNodeResourcesWithDefaults instantiates a new NodeResources object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNodeResourcesWithDefaults() *NodeResources {
	this := NodeResources{}
	return &this
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *NodeResources) GetCpu() NodeCpuResources {
	if o == nil || o.Cpu == nil {
		var ret NodeCpuResources
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeResources) GetCpuOk() (*NodeCpuResources, bool) {
	if o == nil || o.Cpu == nil {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *NodeResources) HasCpu() bool {
	if o != nil && o.Cpu != nil {
		return true
	}

	return false
}

// SetCpu gets a reference to the given NodeCpuResources and assigns it to the Cpu field.
func (o *NodeResources) SetCpu(v NodeCpuResources) {
	o.Cpu = &v
}

// GetDevices returns the Devices field value if set, zero value otherwise.
func (o *NodeResources) GetDevices() []NodeDeviceResource {
	if o == nil || o.Devices == nil {
		var ret []NodeDeviceResource
		return ret
	}
	return *o.Devices
}

// GetDevicesOk returns a tuple with the Devices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeResources) GetDevicesOk() (*[]NodeDeviceResource, bool) {
	if o == nil || o.Devices == nil {
		return nil, false
	}
	return o.Devices, true
}

// HasDevices returns a boolean if a field has been set.
func (o *NodeResources) HasDevices() bool {
	if o != nil && o.Devices != nil {
		return true
	}

	return false
}

// SetDevices gets a reference to the given []NodeDeviceResource and assigns it to the Devices field.
func (o *NodeResources) SetDevices(v []NodeDeviceResource) {
	o.Devices = &v
}

// GetDisk returns the Disk field value if set, zero value otherwise.
func (o *NodeResources) GetDisk() NodeDiskResources {
	if o == nil || o.Disk == nil {
		var ret NodeDiskResources
		return ret
	}
	return *o.Disk
}

// GetDiskOk returns a tuple with the Disk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeResources) GetDiskOk() (*NodeDiskResources, bool) {
	if o == nil || o.Disk == nil {
		return nil, false
	}
	return o.Disk, true
}

// HasDisk returns a boolean if a field has been set.
func (o *NodeResources) HasDisk() bool {
	if o != nil && o.Disk != nil {
		return true
	}

	return false
}

// SetDisk gets a reference to the given NodeDiskResources and assigns it to the Disk field.
func (o *NodeResources) SetDisk(v NodeDiskResources) {
	o.Disk = &v
}

// GetMaxDynamicPort returns the MaxDynamicPort field value if set, zero value otherwise.
func (o *NodeResources) GetMaxDynamicPort() int32 {
	if o == nil || o.MaxDynamicPort == nil {
		var ret int32
		return ret
	}
	return *o.MaxDynamicPort
}

// GetMaxDynamicPortOk returns a tuple with the MaxDynamicPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeResources) GetMaxDynamicPortOk() (*int32, bool) {
	if o == nil || o.MaxDynamicPort == nil {
		return nil, false
	}
	return o.MaxDynamicPort, true
}

// HasMaxDynamicPort returns a boolean if a field has been set.
func (o *NodeResources) HasMaxDynamicPort() bool {
	if o != nil && o.MaxDynamicPort != nil {
		return true
	}

	return false
}

// SetMaxDynamicPort gets a reference to the given int32 and assigns it to the MaxDynamicPort field.
func (o *NodeResources) SetMaxDynamicPort(v int32) {
	o.MaxDynamicPort = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *NodeResources) GetMemory() NodeMemoryResources {
	if o == nil || o.Memory == nil {
		var ret NodeMemoryResources
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeResources) GetMemoryOk() (*NodeMemoryResources, bool) {
	if o == nil || o.Memory == nil {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *NodeResources) HasMemory() bool {
	if o != nil && o.Memory != nil {
		return true
	}

	return false
}

// SetMemory gets a reference to the given NodeMemoryResources and assigns it to the Memory field.
func (o *NodeResources) SetMemory(v NodeMemoryResources) {
	o.Memory = &v
}

// GetMinDynamicPort returns the MinDynamicPort field value if set, zero value otherwise.
func (o *NodeResources) GetMinDynamicPort() int32 {
	if o == nil || o.MinDynamicPort == nil {
		var ret int32
		return ret
	}
	return *o.MinDynamicPort
}

// GetMinDynamicPortOk returns a tuple with the MinDynamicPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeResources) GetMinDynamicPortOk() (*int32, bool) {
	if o == nil || o.MinDynamicPort == nil {
		return nil, false
	}
	return o.MinDynamicPort, true
}

// HasMinDynamicPort returns a boolean if a field has been set.
func (o *NodeResources) HasMinDynamicPort() bool {
	if o != nil && o.MinDynamicPort != nil {
		return true
	}

	return false
}

// SetMinDynamicPort gets a reference to the given int32 and assigns it to the MinDynamicPort field.
func (o *NodeResources) SetMinDynamicPort(v int32) {
	o.MinDynamicPort = &v
}

// GetNetworks returns the Networks field value if set, zero value otherwise.
func (o *NodeResources) GetNetworks() []NetworkResource {
	if o == nil || o.Networks == nil {
		var ret []NetworkResource
		return ret
	}
	return *o.Networks
}

// GetNetworksOk returns a tuple with the Networks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeResources) GetNetworksOk() (*[]NetworkResource, bool) {
	if o == nil || o.Networks == nil {
		return nil, false
	}
	return o.Networks, true
}

// HasNetworks returns a boolean if a field has been set.
func (o *NodeResources) HasNetworks() bool {
	if o != nil && o.Networks != nil {
		return true
	}

	return false
}

// SetNetworks gets a reference to the given []NetworkResource and assigns it to the Networks field.
func (o *NodeResources) SetNetworks(v []NetworkResource) {
	o.Networks = &v
}

func (o NodeResources) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Cpu != nil {
		toSerialize["Cpu"] = o.Cpu
	}
	if o.Devices != nil {
		toSerialize["Devices"] = o.Devices
	}
	if o.Disk != nil {
		toSerialize["Disk"] = o.Disk
	}
	if o.MaxDynamicPort != nil {
		toSerialize["MaxDynamicPort"] = o.MaxDynamicPort
	}
	if o.Memory != nil {
		toSerialize["Memory"] = o.Memory
	}
	if o.MinDynamicPort != nil {
		toSerialize["MinDynamicPort"] = o.MinDynamicPort
	}
	if o.Networks != nil {
		toSerialize["Networks"] = o.Networks
	}
	return json.Marshal(toSerialize)
}

type NullableNodeResources struct {
	value *NodeResources
	isSet bool
}

func (v NullableNodeResources) Get() *NodeResources {
	return v.value
}

func (v *NullableNodeResources) Set(val *NodeResources) {
	v.value = val
	v.isSet = true
}

func (v NullableNodeResources) IsSet() bool {
	return v.isSet
}

func (v *NullableNodeResources) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNodeResources(val *NodeResources) *NullableNodeResources {
	return &NullableNodeResources{value: val, isSet: true}
}

func (v NullableNodeResources) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNodeResources) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


