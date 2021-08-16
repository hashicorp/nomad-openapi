/*
 * Nomad
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.1.3
 * Contact: support@hashicorp.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// AllocatedMemoryResources struct for AllocatedMemoryResources
type AllocatedMemoryResources struct {
	MemoryMB *int64 `json:"MemoryMB,omitempty"`
	MemoryMaxMB *int64 `json:"MemoryMaxMB,omitempty"`
}

// NewAllocatedMemoryResources instantiates a new AllocatedMemoryResources object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllocatedMemoryResources() *AllocatedMemoryResources {
	this := AllocatedMemoryResources{}
	return &this
}

// NewAllocatedMemoryResourcesWithDefaults instantiates a new AllocatedMemoryResources object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllocatedMemoryResourcesWithDefaults() *AllocatedMemoryResources {
	this := AllocatedMemoryResources{}
	return &this
}

// GetMemoryMB returns the MemoryMB field value if set, zero value otherwise.
func (o *AllocatedMemoryResources) GetMemoryMB() int64 {
	if o == nil || o.MemoryMB == nil {
		var ret int64
		return ret
	}
	return *o.MemoryMB
}

// GetMemoryMBOk returns a tuple with the MemoryMB field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllocatedMemoryResources) GetMemoryMBOk() (*int64, bool) {
	if o == nil || o.MemoryMB == nil {
		return nil, false
	}
	return o.MemoryMB, true
}

// HasMemoryMB returns a boolean if a field has been set.
func (o *AllocatedMemoryResources) HasMemoryMB() bool {
	if o != nil && o.MemoryMB != nil {
		return true
	}

	return false
}

// SetMemoryMB gets a reference to the given int64 and assigns it to the MemoryMB field.
func (o *AllocatedMemoryResources) SetMemoryMB(v int64) {
	o.MemoryMB = &v
}

// GetMemoryMaxMB returns the MemoryMaxMB field value if set, zero value otherwise.
func (o *AllocatedMemoryResources) GetMemoryMaxMB() int64 {
	if o == nil || o.MemoryMaxMB == nil {
		var ret int64
		return ret
	}
	return *o.MemoryMaxMB
}

// GetMemoryMaxMBOk returns a tuple with the MemoryMaxMB field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllocatedMemoryResources) GetMemoryMaxMBOk() (*int64, bool) {
	if o == nil || o.MemoryMaxMB == nil {
		return nil, false
	}
	return o.MemoryMaxMB, true
}

// HasMemoryMaxMB returns a boolean if a field has been set.
func (o *AllocatedMemoryResources) HasMemoryMaxMB() bool {
	if o != nil && o.MemoryMaxMB != nil {
		return true
	}

	return false
}

// SetMemoryMaxMB gets a reference to the given int64 and assigns it to the MemoryMaxMB field.
func (o *AllocatedMemoryResources) SetMemoryMaxMB(v int64) {
	o.MemoryMaxMB = &v
}

func (o AllocatedMemoryResources) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MemoryMB != nil {
		toSerialize["MemoryMB"] = o.MemoryMB
	}
	if o.MemoryMaxMB != nil {
		toSerialize["MemoryMaxMB"] = o.MemoryMaxMB
	}
	return json.Marshal(toSerialize)
}

type NullableAllocatedMemoryResources struct {
	value *AllocatedMemoryResources
	isSet bool
}

func (v NullableAllocatedMemoryResources) Get() *AllocatedMemoryResources {
	return v.value
}

func (v *NullableAllocatedMemoryResources) Set(val *AllocatedMemoryResources) {
	v.value = val
	v.isSet = true
}

func (v NullableAllocatedMemoryResources) IsSet() bool {
	return v.isSet
}

func (v *NullableAllocatedMemoryResources) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllocatedMemoryResources(val *AllocatedMemoryResources) *NullableAllocatedMemoryResources {
	return &NullableAllocatedMemoryResources{value: val, isSet: true}
}

func (v NullableAllocatedMemoryResources) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllocatedMemoryResources) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


