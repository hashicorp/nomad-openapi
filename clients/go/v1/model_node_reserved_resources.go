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

// NodeReservedResources struct for NodeReservedResources
type NodeReservedResources struct {
	Cpu *NodeReservedCpuResources `json:"Cpu,omitempty"`
	Disk *NodeReservedDiskResources `json:"Disk,omitempty"`
	Memory *NodeReservedMemoryResources `json:"Memory,omitempty"`
	Networks *NodeReservedNetworkResources `json:"Networks,omitempty"`
}

// NewNodeReservedResources instantiates a new NodeReservedResources object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNodeReservedResources() *NodeReservedResources {
	this := NodeReservedResources{}
	return &this
}

// NewNodeReservedResourcesWithDefaults instantiates a new NodeReservedResources object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNodeReservedResourcesWithDefaults() *NodeReservedResources {
	this := NodeReservedResources{}
	return &this
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *NodeReservedResources) GetCpu() NodeReservedCpuResources {
	if o == nil || o.Cpu == nil {
		var ret NodeReservedCpuResources
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeReservedResources) GetCpuOk() (*NodeReservedCpuResources, bool) {
	if o == nil || o.Cpu == nil {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *NodeReservedResources) HasCpu() bool {
	if o != nil && o.Cpu != nil {
		return true
	}

	return false
}

// SetCpu gets a reference to the given NodeReservedCpuResources and assigns it to the Cpu field.
func (o *NodeReservedResources) SetCpu(v NodeReservedCpuResources) {
	o.Cpu = &v
}

// GetDisk returns the Disk field value if set, zero value otherwise.
func (o *NodeReservedResources) GetDisk() NodeReservedDiskResources {
	if o == nil || o.Disk == nil {
		var ret NodeReservedDiskResources
		return ret
	}
	return *o.Disk
}

// GetDiskOk returns a tuple with the Disk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeReservedResources) GetDiskOk() (*NodeReservedDiskResources, bool) {
	if o == nil || o.Disk == nil {
		return nil, false
	}
	return o.Disk, true
}

// HasDisk returns a boolean if a field has been set.
func (o *NodeReservedResources) HasDisk() bool {
	if o != nil && o.Disk != nil {
		return true
	}

	return false
}

// SetDisk gets a reference to the given NodeReservedDiskResources and assigns it to the Disk field.
func (o *NodeReservedResources) SetDisk(v NodeReservedDiskResources) {
	o.Disk = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *NodeReservedResources) GetMemory() NodeReservedMemoryResources {
	if o == nil || o.Memory == nil {
		var ret NodeReservedMemoryResources
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeReservedResources) GetMemoryOk() (*NodeReservedMemoryResources, bool) {
	if o == nil || o.Memory == nil {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *NodeReservedResources) HasMemory() bool {
	if o != nil && o.Memory != nil {
		return true
	}

	return false
}

// SetMemory gets a reference to the given NodeReservedMemoryResources and assigns it to the Memory field.
func (o *NodeReservedResources) SetMemory(v NodeReservedMemoryResources) {
	o.Memory = &v
}

// GetNetworks returns the Networks field value if set, zero value otherwise.
func (o *NodeReservedResources) GetNetworks() NodeReservedNetworkResources {
	if o == nil || o.Networks == nil {
		var ret NodeReservedNetworkResources
		return ret
	}
	return *o.Networks
}

// GetNetworksOk returns a tuple with the Networks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeReservedResources) GetNetworksOk() (*NodeReservedNetworkResources, bool) {
	if o == nil || o.Networks == nil {
		return nil, false
	}
	return o.Networks, true
}

// HasNetworks returns a boolean if a field has been set.
func (o *NodeReservedResources) HasNetworks() bool {
	if o != nil && o.Networks != nil {
		return true
	}

	return false
}

// SetNetworks gets a reference to the given NodeReservedNetworkResources and assigns it to the Networks field.
func (o *NodeReservedResources) SetNetworks(v NodeReservedNetworkResources) {
	o.Networks = &v
}

func (o NodeReservedResources) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Cpu != nil {
		toSerialize["Cpu"] = o.Cpu
	}
	if o.Disk != nil {
		toSerialize["Disk"] = o.Disk
	}
	if o.Memory != nil {
		toSerialize["Memory"] = o.Memory
	}
	if o.Networks != nil {
		toSerialize["Networks"] = o.Networks
	}
	return json.Marshal(toSerialize)
}

type NullableNodeReservedResources struct {
	value *NodeReservedResources
	isSet bool
}

func (v NullableNodeReservedResources) Get() *NodeReservedResources {
	return v.value
}

func (v *NullableNodeReservedResources) Set(val *NodeReservedResources) {
	v.value = val
	v.isSet = true
}

func (v NullableNodeReservedResources) IsSet() bool {
	return v.isSet
}

func (v *NullableNodeReservedResources) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNodeReservedResources(val *NodeReservedResources) *NullableNodeReservedResources {
	return &NullableNodeReservedResources{value: val, isSet: true}
}

func (v NullableNodeReservedResources) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNodeReservedResources) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


