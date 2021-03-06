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

// AllocatedDeviceResource struct for AllocatedDeviceResource
type AllocatedDeviceResource struct {
	DeviceIDs *[]string `json:"DeviceIDs,omitempty"`
	Name *string `json:"Name,omitempty"`
	Type *string `json:"Type,omitempty"`
	Vendor *string `json:"Vendor,omitempty"`
}

// NewAllocatedDeviceResource instantiates a new AllocatedDeviceResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllocatedDeviceResource() *AllocatedDeviceResource {
	this := AllocatedDeviceResource{}
	return &this
}

// NewAllocatedDeviceResourceWithDefaults instantiates a new AllocatedDeviceResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllocatedDeviceResourceWithDefaults() *AllocatedDeviceResource {
	this := AllocatedDeviceResource{}
	return &this
}

// GetDeviceIDs returns the DeviceIDs field value if set, zero value otherwise.
func (o *AllocatedDeviceResource) GetDeviceIDs() []string {
	if o == nil || o.DeviceIDs == nil {
		var ret []string
		return ret
	}
	return *o.DeviceIDs
}

// GetDeviceIDsOk returns a tuple with the DeviceIDs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllocatedDeviceResource) GetDeviceIDsOk() (*[]string, bool) {
	if o == nil || o.DeviceIDs == nil {
		return nil, false
	}
	return o.DeviceIDs, true
}

// HasDeviceIDs returns a boolean if a field has been set.
func (o *AllocatedDeviceResource) HasDeviceIDs() bool {
	if o != nil && o.DeviceIDs != nil {
		return true
	}

	return false
}

// SetDeviceIDs gets a reference to the given []string and assigns it to the DeviceIDs field.
func (o *AllocatedDeviceResource) SetDeviceIDs(v []string) {
	o.DeviceIDs = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AllocatedDeviceResource) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllocatedDeviceResource) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AllocatedDeviceResource) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AllocatedDeviceResource) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AllocatedDeviceResource) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllocatedDeviceResource) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AllocatedDeviceResource) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AllocatedDeviceResource) SetType(v string) {
	o.Type = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *AllocatedDeviceResource) GetVendor() string {
	if o == nil || o.Vendor == nil {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllocatedDeviceResource) GetVendorOk() (*string, bool) {
	if o == nil || o.Vendor == nil {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *AllocatedDeviceResource) HasVendor() bool {
	if o != nil && o.Vendor != nil {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *AllocatedDeviceResource) SetVendor(v string) {
	o.Vendor = &v
}

func (o AllocatedDeviceResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeviceIDs != nil {
		toSerialize["DeviceIDs"] = o.DeviceIDs
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	if o.Vendor != nil {
		toSerialize["Vendor"] = o.Vendor
	}
	return json.Marshal(toSerialize)
}

type NullableAllocatedDeviceResource struct {
	value *AllocatedDeviceResource
	isSet bool
}

func (v NullableAllocatedDeviceResource) Get() *AllocatedDeviceResource {
	return v.value
}

func (v *NullableAllocatedDeviceResource) Set(val *AllocatedDeviceResource) {
	v.value = val
	v.isSet = true
}

func (v NullableAllocatedDeviceResource) IsSet() bool {
	return v.isSet
}

func (v *NullableAllocatedDeviceResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllocatedDeviceResource(val *AllocatedDeviceResource) *NullableAllocatedDeviceResource {
	return &NullableAllocatedDeviceResource{value: val, isSet: true}
}

func (v NullableAllocatedDeviceResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllocatedDeviceResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


