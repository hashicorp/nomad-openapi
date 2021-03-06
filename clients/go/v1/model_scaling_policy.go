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

// ScalingPolicy struct for ScalingPolicy
type ScalingPolicy struct {
	CreateIndex *int32 `json:"CreateIndex,omitempty"`
	Enabled *bool `json:"Enabled,omitempty"`
	ID *string `json:"ID,omitempty"`
	Max *int64 `json:"Max,omitempty"`
	Min *int64 `json:"Min,omitempty"`
	ModifyIndex *int32 `json:"ModifyIndex,omitempty"`
	Namespace *string `json:"Namespace,omitempty"`
	Policy *map[string]interface{} `json:"Policy,omitempty"`
	Target *map[string]string `json:"Target,omitempty"`
	Type *string `json:"Type,omitempty"`
}

// NewScalingPolicy instantiates a new ScalingPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScalingPolicy() *ScalingPolicy {
	this := ScalingPolicy{}
	return &this
}

// NewScalingPolicyWithDefaults instantiates a new ScalingPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScalingPolicyWithDefaults() *ScalingPolicy {
	this := ScalingPolicy{}
	return &this
}

// GetCreateIndex returns the CreateIndex field value if set, zero value otherwise.
func (o *ScalingPolicy) GetCreateIndex() int32 {
	if o == nil || o.CreateIndex == nil {
		var ret int32
		return ret
	}
	return *o.CreateIndex
}

// GetCreateIndexOk returns a tuple with the CreateIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScalingPolicy) GetCreateIndexOk() (*int32, bool) {
	if o == nil || o.CreateIndex == nil {
		return nil, false
	}
	return o.CreateIndex, true
}

// HasCreateIndex returns a boolean if a field has been set.
func (o *ScalingPolicy) HasCreateIndex() bool {
	if o != nil && o.CreateIndex != nil {
		return true
	}

	return false
}

// SetCreateIndex gets a reference to the given int32 and assigns it to the CreateIndex field.
func (o *ScalingPolicy) SetCreateIndex(v int32) {
	o.CreateIndex = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ScalingPolicy) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScalingPolicy) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ScalingPolicy) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ScalingPolicy) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *ScalingPolicy) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScalingPolicy) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *ScalingPolicy) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *ScalingPolicy) SetID(v string) {
	o.ID = &v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *ScalingPolicy) GetMax() int64 {
	if o == nil || o.Max == nil {
		var ret int64
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScalingPolicy) GetMaxOk() (*int64, bool) {
	if o == nil || o.Max == nil {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *ScalingPolicy) HasMax() bool {
	if o != nil && o.Max != nil {
		return true
	}

	return false
}

// SetMax gets a reference to the given int64 and assigns it to the Max field.
func (o *ScalingPolicy) SetMax(v int64) {
	o.Max = &v
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *ScalingPolicy) GetMin() int64 {
	if o == nil || o.Min == nil {
		var ret int64
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScalingPolicy) GetMinOk() (*int64, bool) {
	if o == nil || o.Min == nil {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *ScalingPolicy) HasMin() bool {
	if o != nil && o.Min != nil {
		return true
	}

	return false
}

// SetMin gets a reference to the given int64 and assigns it to the Min field.
func (o *ScalingPolicy) SetMin(v int64) {
	o.Min = &v
}

// GetModifyIndex returns the ModifyIndex field value if set, zero value otherwise.
func (o *ScalingPolicy) GetModifyIndex() int32 {
	if o == nil || o.ModifyIndex == nil {
		var ret int32
		return ret
	}
	return *o.ModifyIndex
}

// GetModifyIndexOk returns a tuple with the ModifyIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScalingPolicy) GetModifyIndexOk() (*int32, bool) {
	if o == nil || o.ModifyIndex == nil {
		return nil, false
	}
	return o.ModifyIndex, true
}

// HasModifyIndex returns a boolean if a field has been set.
func (o *ScalingPolicy) HasModifyIndex() bool {
	if o != nil && o.ModifyIndex != nil {
		return true
	}

	return false
}

// SetModifyIndex gets a reference to the given int32 and assigns it to the ModifyIndex field.
func (o *ScalingPolicy) SetModifyIndex(v int32) {
	o.ModifyIndex = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *ScalingPolicy) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScalingPolicy) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *ScalingPolicy) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *ScalingPolicy) SetNamespace(v string) {
	o.Namespace = &v
}

// GetPolicy returns the Policy field value if set, zero value otherwise.
func (o *ScalingPolicy) GetPolicy() map[string]interface{} {
	if o == nil || o.Policy == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScalingPolicy) GetPolicyOk() (*map[string]interface{}, bool) {
	if o == nil || o.Policy == nil {
		return nil, false
	}
	return o.Policy, true
}

// HasPolicy returns a boolean if a field has been set.
func (o *ScalingPolicy) HasPolicy() bool {
	if o != nil && o.Policy != nil {
		return true
	}

	return false
}

// SetPolicy gets a reference to the given map[string]interface{} and assigns it to the Policy field.
func (o *ScalingPolicy) SetPolicy(v map[string]interface{}) {
	o.Policy = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *ScalingPolicy) GetTarget() map[string]string {
	if o == nil || o.Target == nil {
		var ret map[string]string
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScalingPolicy) GetTargetOk() (*map[string]string, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *ScalingPolicy) HasTarget() bool {
	if o != nil && o.Target != nil {
		return true
	}

	return false
}

// SetTarget gets a reference to the given map[string]string and assigns it to the Target field.
func (o *ScalingPolicy) SetTarget(v map[string]string) {
	o.Target = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ScalingPolicy) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScalingPolicy) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ScalingPolicy) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ScalingPolicy) SetType(v string) {
	o.Type = &v
}

func (o ScalingPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreateIndex != nil {
		toSerialize["CreateIndex"] = o.CreateIndex
	}
	if o.Enabled != nil {
		toSerialize["Enabled"] = o.Enabled
	}
	if o.ID != nil {
		toSerialize["ID"] = o.ID
	}
	if o.Max != nil {
		toSerialize["Max"] = o.Max
	}
	if o.Min != nil {
		toSerialize["Min"] = o.Min
	}
	if o.ModifyIndex != nil {
		toSerialize["ModifyIndex"] = o.ModifyIndex
	}
	if o.Namespace != nil {
		toSerialize["Namespace"] = o.Namespace
	}
	if o.Policy != nil {
		toSerialize["Policy"] = o.Policy
	}
	if o.Target != nil {
		toSerialize["Target"] = o.Target
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableScalingPolicy struct {
	value *ScalingPolicy
	isSet bool
}

func (v NullableScalingPolicy) Get() *ScalingPolicy {
	return v.value
}

func (v *NullableScalingPolicy) Set(val *ScalingPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableScalingPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableScalingPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScalingPolicy(val *ScalingPolicy) *NullableScalingPolicy {
	return &NullableScalingPolicy{value: val, isSet: true}
}

func (v NullableScalingPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScalingPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


