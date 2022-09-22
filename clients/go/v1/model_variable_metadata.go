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

// VariableMetadata struct for VariableMetadata
type VariableMetadata struct {
	CreateIndex *int32 `json:"CreateIndex,omitempty"`
	CreateTime *int64 `json:"CreateTime,omitempty"`
	ModifyIndex *int32 `json:"ModifyIndex,omitempty"`
	ModifyTime *int64 `json:"ModifyTime,omitempty"`
	Namespace *string `json:"Namespace,omitempty"`
	Path *string `json:"Path,omitempty"`
}

// NewVariableMetadata instantiates a new VariableMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVariableMetadata() *VariableMetadata {
	this := VariableMetadata{}
	return &this
}

// NewVariableMetadataWithDefaults instantiates a new VariableMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVariableMetadataWithDefaults() *VariableMetadata {
	this := VariableMetadata{}
	return &this
}

// GetCreateIndex returns the CreateIndex field value if set, zero value otherwise.
func (o *VariableMetadata) GetCreateIndex() int32 {
	if o == nil || o.CreateIndex == nil {
		var ret int32
		return ret
	}
	return *o.CreateIndex
}

// GetCreateIndexOk returns a tuple with the CreateIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariableMetadata) GetCreateIndexOk() (*int32, bool) {
	if o == nil || o.CreateIndex == nil {
		return nil, false
	}
	return o.CreateIndex, true
}

// HasCreateIndex returns a boolean if a field has been set.
func (o *VariableMetadata) HasCreateIndex() bool {
	if o != nil && o.CreateIndex != nil {
		return true
	}

	return false
}

// SetCreateIndex gets a reference to the given int32 and assigns it to the CreateIndex field.
func (o *VariableMetadata) SetCreateIndex(v int32) {
	o.CreateIndex = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *VariableMetadata) GetCreateTime() int64 {
	if o == nil || o.CreateTime == nil {
		var ret int64
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariableMetadata) GetCreateTimeOk() (*int64, bool) {
	if o == nil || o.CreateTime == nil {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *VariableMetadata) HasCreateTime() bool {
	if o != nil && o.CreateTime != nil {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given int64 and assigns it to the CreateTime field.
func (o *VariableMetadata) SetCreateTime(v int64) {
	o.CreateTime = &v
}

// GetModifyIndex returns the ModifyIndex field value if set, zero value otherwise.
func (o *VariableMetadata) GetModifyIndex() int32 {
	if o == nil || o.ModifyIndex == nil {
		var ret int32
		return ret
	}
	return *o.ModifyIndex
}

// GetModifyIndexOk returns a tuple with the ModifyIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariableMetadata) GetModifyIndexOk() (*int32, bool) {
	if o == nil || o.ModifyIndex == nil {
		return nil, false
	}
	return o.ModifyIndex, true
}

// HasModifyIndex returns a boolean if a field has been set.
func (o *VariableMetadata) HasModifyIndex() bool {
	if o != nil && o.ModifyIndex != nil {
		return true
	}

	return false
}

// SetModifyIndex gets a reference to the given int32 and assigns it to the ModifyIndex field.
func (o *VariableMetadata) SetModifyIndex(v int32) {
	o.ModifyIndex = &v
}

// GetModifyTime returns the ModifyTime field value if set, zero value otherwise.
func (o *VariableMetadata) GetModifyTime() int64 {
	if o == nil || o.ModifyTime == nil {
		var ret int64
		return ret
	}
	return *o.ModifyTime
}

// GetModifyTimeOk returns a tuple with the ModifyTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariableMetadata) GetModifyTimeOk() (*int64, bool) {
	if o == nil || o.ModifyTime == nil {
		return nil, false
	}
	return o.ModifyTime, true
}

// HasModifyTime returns a boolean if a field has been set.
func (o *VariableMetadata) HasModifyTime() bool {
	if o != nil && o.ModifyTime != nil {
		return true
	}

	return false
}

// SetModifyTime gets a reference to the given int64 and assigns it to the ModifyTime field.
func (o *VariableMetadata) SetModifyTime(v int64) {
	o.ModifyTime = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *VariableMetadata) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariableMetadata) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *VariableMetadata) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *VariableMetadata) SetNamespace(v string) {
	o.Namespace = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *VariableMetadata) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariableMetadata) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *VariableMetadata) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *VariableMetadata) SetPath(v string) {
	o.Path = &v
}

func (o VariableMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreateIndex != nil {
		toSerialize["CreateIndex"] = o.CreateIndex
	}
	if o.CreateTime != nil {
		toSerialize["CreateTime"] = o.CreateTime
	}
	if o.ModifyIndex != nil {
		toSerialize["ModifyIndex"] = o.ModifyIndex
	}
	if o.ModifyTime != nil {
		toSerialize["ModifyTime"] = o.ModifyTime
	}
	if o.Namespace != nil {
		toSerialize["Namespace"] = o.Namespace
	}
	if o.Path != nil {
		toSerialize["Path"] = o.Path
	}
	return json.Marshal(toSerialize)
}

type NullableVariableMetadata struct {
	value *VariableMetadata
	isSet bool
}

func (v NullableVariableMetadata) Get() *VariableMetadata {
	return v.value
}

func (v *NullableVariableMetadata) Set(val *VariableMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableVariableMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableVariableMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVariableMetadata(val *VariableMetadata) *NullableVariableMetadata {
	return &NullableVariableMetadata{value: val, isSet: true}
}

func (v NullableVariableMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVariableMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

