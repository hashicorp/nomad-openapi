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

// HostVolumeInfo struct for HostVolumeInfo
type HostVolumeInfo struct {
	Path *string `json:"Path,omitempty"`
	ReadOnly *bool `json:"ReadOnly,omitempty"`
}

// NewHostVolumeInfo instantiates a new HostVolumeInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHostVolumeInfo() *HostVolumeInfo {
	this := HostVolumeInfo{}
	return &this
}

// NewHostVolumeInfoWithDefaults instantiates a new HostVolumeInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHostVolumeInfoWithDefaults() *HostVolumeInfo {
	this := HostVolumeInfo{}
	return &this
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *HostVolumeInfo) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostVolumeInfo) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *HostVolumeInfo) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *HostVolumeInfo) SetPath(v string) {
	o.Path = &v
}

// GetReadOnly returns the ReadOnly field value if set, zero value otherwise.
func (o *HostVolumeInfo) GetReadOnly() bool {
	if o == nil || o.ReadOnly == nil {
		var ret bool
		return ret
	}
	return *o.ReadOnly
}

// GetReadOnlyOk returns a tuple with the ReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostVolumeInfo) GetReadOnlyOk() (*bool, bool) {
	if o == nil || o.ReadOnly == nil {
		return nil, false
	}
	return o.ReadOnly, true
}

// HasReadOnly returns a boolean if a field has been set.
func (o *HostVolumeInfo) HasReadOnly() bool {
	if o != nil && o.ReadOnly != nil {
		return true
	}

	return false
}

// SetReadOnly gets a reference to the given bool and assigns it to the ReadOnly field.
func (o *HostVolumeInfo) SetReadOnly(v bool) {
	o.ReadOnly = &v
}

func (o HostVolumeInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Path != nil {
		toSerialize["Path"] = o.Path
	}
	if o.ReadOnly != nil {
		toSerialize["ReadOnly"] = o.ReadOnly
	}
	return json.Marshal(toSerialize)
}

type NullableHostVolumeInfo struct {
	value *HostVolumeInfo
	isSet bool
}

func (v NullableHostVolumeInfo) Get() *HostVolumeInfo {
	return v.value
}

func (v *NullableHostVolumeInfo) Set(val *HostVolumeInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableHostVolumeInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableHostVolumeInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHostVolumeInfo(val *HostVolumeInfo) *NullableHostVolumeInfo {
	return &NullableHostVolumeInfo{value: val, isSet: true}
}

func (v NullableHostVolumeInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHostVolumeInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


