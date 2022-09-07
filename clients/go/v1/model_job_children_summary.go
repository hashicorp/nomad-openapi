/*
Nomad

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.1.4
Contact: support@hashicorp.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// JobChildrenSummary struct for JobChildrenSummary
type JobChildrenSummary struct {
	Dead *int64 `json:"Dead,omitempty"`
	Pending *int64 `json:"Pending,omitempty"`
	Running *int64 `json:"Running,omitempty"`
}

// NewJobChildrenSummary instantiates a new JobChildrenSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobChildrenSummary() *JobChildrenSummary {
	this := JobChildrenSummary{}
	return &this
}

// NewJobChildrenSummaryWithDefaults instantiates a new JobChildrenSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobChildrenSummaryWithDefaults() *JobChildrenSummary {
	this := JobChildrenSummary{}
	return &this
}

// GetDead returns the Dead field value if set, zero value otherwise.
func (o *JobChildrenSummary) GetDead() int64 {
	if o == nil || o.Dead == nil {
		var ret int64
		return ret
	}
	return *o.Dead
}

// GetDeadOk returns a tuple with the Dead field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobChildrenSummary) GetDeadOk() (*int64, bool) {
	if o == nil || o.Dead == nil {
		return nil, false
	}
	return o.Dead, true
}

// HasDead returns a boolean if a field has been set.
func (o *JobChildrenSummary) HasDead() bool {
	if o != nil && o.Dead != nil {
		return true
	}

	return false
}

// SetDead gets a reference to the given int64 and assigns it to the Dead field.
func (o *JobChildrenSummary) SetDead(v int64) {
	o.Dead = &v
}

// GetPending returns the Pending field value if set, zero value otherwise.
func (o *JobChildrenSummary) GetPending() int64 {
	if o == nil || o.Pending == nil {
		var ret int64
		return ret
	}
	return *o.Pending
}

// GetPendingOk returns a tuple with the Pending field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobChildrenSummary) GetPendingOk() (*int64, bool) {
	if o == nil || o.Pending == nil {
		return nil, false
	}
	return o.Pending, true
}

// HasPending returns a boolean if a field has been set.
func (o *JobChildrenSummary) HasPending() bool {
	if o != nil && o.Pending != nil {
		return true
	}

	return false
}

// SetPending gets a reference to the given int64 and assigns it to the Pending field.
func (o *JobChildrenSummary) SetPending(v int64) {
	o.Pending = &v
}

// GetRunning returns the Running field value if set, zero value otherwise.
func (o *JobChildrenSummary) GetRunning() int64 {
	if o == nil || o.Running == nil {
		var ret int64
		return ret
	}
	return *o.Running
}

// GetRunningOk returns a tuple with the Running field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobChildrenSummary) GetRunningOk() (*int64, bool) {
	if o == nil || o.Running == nil {
		return nil, false
	}
	return o.Running, true
}

// HasRunning returns a boolean if a field has been set.
func (o *JobChildrenSummary) HasRunning() bool {
	if o != nil && o.Running != nil {
		return true
	}

	return false
}

// SetRunning gets a reference to the given int64 and assigns it to the Running field.
func (o *JobChildrenSummary) SetRunning(v int64) {
	o.Running = &v
}

func (o JobChildrenSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Dead != nil {
		toSerialize["Dead"] = o.Dead
	}
	if o.Pending != nil {
		toSerialize["Pending"] = o.Pending
	}
	if o.Running != nil {
		toSerialize["Running"] = o.Running
	}
	return json.Marshal(toSerialize)
}

type NullableJobChildrenSummary struct {
	value *JobChildrenSummary
	isSet bool
}

func (v NullableJobChildrenSummary) Get() *JobChildrenSummary {
	return v.value
}

func (v *NullableJobChildrenSummary) Set(val *JobChildrenSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableJobChildrenSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableJobChildrenSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobChildrenSummary(val *JobChildrenSummary) *NullableJobChildrenSummary {
	return &NullableJobChildrenSummary{value: val, isSet: true}
}

func (v NullableJobChildrenSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobChildrenSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


