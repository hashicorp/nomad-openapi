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
	"time"
)

// AllocDeploymentStatus struct for AllocDeploymentStatus
type AllocDeploymentStatus struct {
	Canary *bool `json:"Canary,omitempty"`
	Healthy *bool `json:"Healthy,omitempty"`
	ModifyIndex *int32 `json:"ModifyIndex,omitempty"`
	Timestamp *time.Time `json:"Timestamp,omitempty"`
}

// NewAllocDeploymentStatus instantiates a new AllocDeploymentStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllocDeploymentStatus() *AllocDeploymentStatus {
	this := AllocDeploymentStatus{}
	return &this
}

// NewAllocDeploymentStatusWithDefaults instantiates a new AllocDeploymentStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllocDeploymentStatusWithDefaults() *AllocDeploymentStatus {
	this := AllocDeploymentStatus{}
	return &this
}

// GetCanary returns the Canary field value if set, zero value otherwise.
func (o *AllocDeploymentStatus) GetCanary() bool {
	if o == nil || o.Canary == nil {
		var ret bool
		return ret
	}
	return *o.Canary
}

// GetCanaryOk returns a tuple with the Canary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllocDeploymentStatus) GetCanaryOk() (*bool, bool) {
	if o == nil || o.Canary == nil {
		return nil, false
	}
	return o.Canary, true
}

// HasCanary returns a boolean if a field has been set.
func (o *AllocDeploymentStatus) HasCanary() bool {
	if o != nil && o.Canary != nil {
		return true
	}

	return false
}

// SetCanary gets a reference to the given bool and assigns it to the Canary field.
func (o *AllocDeploymentStatus) SetCanary(v bool) {
	o.Canary = &v
}

// GetHealthy returns the Healthy field value if set, zero value otherwise.
func (o *AllocDeploymentStatus) GetHealthy() bool {
	if o == nil || o.Healthy == nil {
		var ret bool
		return ret
	}
	return *o.Healthy
}

// GetHealthyOk returns a tuple with the Healthy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllocDeploymentStatus) GetHealthyOk() (*bool, bool) {
	if o == nil || o.Healthy == nil {
		return nil, false
	}
	return o.Healthy, true
}

// HasHealthy returns a boolean if a field has been set.
func (o *AllocDeploymentStatus) HasHealthy() bool {
	if o != nil && o.Healthy != nil {
		return true
	}

	return false
}

// SetHealthy gets a reference to the given bool and assigns it to the Healthy field.
func (o *AllocDeploymentStatus) SetHealthy(v bool) {
	o.Healthy = &v
}

// GetModifyIndex returns the ModifyIndex field value if set, zero value otherwise.
func (o *AllocDeploymentStatus) GetModifyIndex() int32 {
	if o == nil || o.ModifyIndex == nil {
		var ret int32
		return ret
	}
	return *o.ModifyIndex
}

// GetModifyIndexOk returns a tuple with the ModifyIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllocDeploymentStatus) GetModifyIndexOk() (*int32, bool) {
	if o == nil || o.ModifyIndex == nil {
		return nil, false
	}
	return o.ModifyIndex, true
}

// HasModifyIndex returns a boolean if a field has been set.
func (o *AllocDeploymentStatus) HasModifyIndex() bool {
	if o != nil && o.ModifyIndex != nil {
		return true
	}

	return false
}

// SetModifyIndex gets a reference to the given int32 and assigns it to the ModifyIndex field.
func (o *AllocDeploymentStatus) SetModifyIndex(v int32) {
	o.ModifyIndex = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *AllocDeploymentStatus) GetTimestamp() time.Time {
	if o == nil || o.Timestamp == nil {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllocDeploymentStatus) GetTimestampOk() (*time.Time, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *AllocDeploymentStatus) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *AllocDeploymentStatus) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

func (o AllocDeploymentStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Canary != nil {
		toSerialize["Canary"] = o.Canary
	}
	if o.Healthy != nil {
		toSerialize["Healthy"] = o.Healthy
	}
	if o.ModifyIndex != nil {
		toSerialize["ModifyIndex"] = o.ModifyIndex
	}
	if o.Timestamp != nil {
		toSerialize["Timestamp"] = o.Timestamp
	}
	return json.Marshal(toSerialize)
}

type NullableAllocDeploymentStatus struct {
	value *AllocDeploymentStatus
	isSet bool
}

func (v NullableAllocDeploymentStatus) Get() *AllocDeploymentStatus {
	return v.value
}

func (v *NullableAllocDeploymentStatus) Set(val *AllocDeploymentStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableAllocDeploymentStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableAllocDeploymentStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllocDeploymentStatus(val *AllocDeploymentStatus) *NullableAllocDeploymentStatus {
	return &NullableAllocDeploymentStatus{value: val, isSet: true}
}

func (v NullableAllocDeploymentStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllocDeploymentStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


