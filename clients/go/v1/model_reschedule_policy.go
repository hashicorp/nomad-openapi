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

// ReschedulePolicy struct for ReschedulePolicy
type ReschedulePolicy struct {
	Attempts *int32 `json:"Attempts,omitempty"`
	Delay *int64 `json:"Delay,omitempty"`
	DelayFunction *string `json:"DelayFunction,omitempty"`
	Interval *int64 `json:"Interval,omitempty"`
	MaxDelay *int64 `json:"MaxDelay,omitempty"`
	Unlimited *bool `json:"Unlimited,omitempty"`
}

// NewReschedulePolicy instantiates a new ReschedulePolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReschedulePolicy() *ReschedulePolicy {
	this := ReschedulePolicy{}
	return &this
}

// NewReschedulePolicyWithDefaults instantiates a new ReschedulePolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReschedulePolicyWithDefaults() *ReschedulePolicy {
	this := ReschedulePolicy{}
	return &this
}

// GetAttempts returns the Attempts field value if set, zero value otherwise.
func (o *ReschedulePolicy) GetAttempts() int32 {
	if o == nil || o.Attempts == nil {
		var ret int32
		return ret
	}
	return *o.Attempts
}

// GetAttemptsOk returns a tuple with the Attempts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReschedulePolicy) GetAttemptsOk() (*int32, bool) {
	if o == nil || o.Attempts == nil {
		return nil, false
	}
	return o.Attempts, true
}

// HasAttempts returns a boolean if a field has been set.
func (o *ReschedulePolicy) HasAttempts() bool {
	if o != nil && o.Attempts != nil {
		return true
	}

	return false
}

// SetAttempts gets a reference to the given int32 and assigns it to the Attempts field.
func (o *ReschedulePolicy) SetAttempts(v int32) {
	o.Attempts = &v
}

// GetDelay returns the Delay field value if set, zero value otherwise.
func (o *ReschedulePolicy) GetDelay() int64 {
	if o == nil || o.Delay == nil {
		var ret int64
		return ret
	}
	return *o.Delay
}

// GetDelayOk returns a tuple with the Delay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReschedulePolicy) GetDelayOk() (*int64, bool) {
	if o == nil || o.Delay == nil {
		return nil, false
	}
	return o.Delay, true
}

// HasDelay returns a boolean if a field has been set.
func (o *ReschedulePolicy) HasDelay() bool {
	if o != nil && o.Delay != nil {
		return true
	}

	return false
}

// SetDelay gets a reference to the given int64 and assigns it to the Delay field.
func (o *ReschedulePolicy) SetDelay(v int64) {
	o.Delay = &v
}

// GetDelayFunction returns the DelayFunction field value if set, zero value otherwise.
func (o *ReschedulePolicy) GetDelayFunction() string {
	if o == nil || o.DelayFunction == nil {
		var ret string
		return ret
	}
	return *o.DelayFunction
}

// GetDelayFunctionOk returns a tuple with the DelayFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReschedulePolicy) GetDelayFunctionOk() (*string, bool) {
	if o == nil || o.DelayFunction == nil {
		return nil, false
	}
	return o.DelayFunction, true
}

// HasDelayFunction returns a boolean if a field has been set.
func (o *ReschedulePolicy) HasDelayFunction() bool {
	if o != nil && o.DelayFunction != nil {
		return true
	}

	return false
}

// SetDelayFunction gets a reference to the given string and assigns it to the DelayFunction field.
func (o *ReschedulePolicy) SetDelayFunction(v string) {
	o.DelayFunction = &v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *ReschedulePolicy) GetInterval() int64 {
	if o == nil || o.Interval == nil {
		var ret int64
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReschedulePolicy) GetIntervalOk() (*int64, bool) {
	if o == nil || o.Interval == nil {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *ReschedulePolicy) HasInterval() bool {
	if o != nil && o.Interval != nil {
		return true
	}

	return false
}

// SetInterval gets a reference to the given int64 and assigns it to the Interval field.
func (o *ReschedulePolicy) SetInterval(v int64) {
	o.Interval = &v
}

// GetMaxDelay returns the MaxDelay field value if set, zero value otherwise.
func (o *ReschedulePolicy) GetMaxDelay() int64 {
	if o == nil || o.MaxDelay == nil {
		var ret int64
		return ret
	}
	return *o.MaxDelay
}

// GetMaxDelayOk returns a tuple with the MaxDelay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReschedulePolicy) GetMaxDelayOk() (*int64, bool) {
	if o == nil || o.MaxDelay == nil {
		return nil, false
	}
	return o.MaxDelay, true
}

// HasMaxDelay returns a boolean if a field has been set.
func (o *ReschedulePolicy) HasMaxDelay() bool {
	if o != nil && o.MaxDelay != nil {
		return true
	}

	return false
}

// SetMaxDelay gets a reference to the given int64 and assigns it to the MaxDelay field.
func (o *ReschedulePolicy) SetMaxDelay(v int64) {
	o.MaxDelay = &v
}

// GetUnlimited returns the Unlimited field value if set, zero value otherwise.
func (o *ReschedulePolicy) GetUnlimited() bool {
	if o == nil || o.Unlimited == nil {
		var ret bool
		return ret
	}
	return *o.Unlimited
}

// GetUnlimitedOk returns a tuple with the Unlimited field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReschedulePolicy) GetUnlimitedOk() (*bool, bool) {
	if o == nil || o.Unlimited == nil {
		return nil, false
	}
	return o.Unlimited, true
}

// HasUnlimited returns a boolean if a field has been set.
func (o *ReschedulePolicy) HasUnlimited() bool {
	if o != nil && o.Unlimited != nil {
		return true
	}

	return false
}

// SetUnlimited gets a reference to the given bool and assigns it to the Unlimited field.
func (o *ReschedulePolicy) SetUnlimited(v bool) {
	o.Unlimited = &v
}

func (o ReschedulePolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Attempts != nil {
		toSerialize["Attempts"] = o.Attempts
	}
	if o.Delay != nil {
		toSerialize["Delay"] = o.Delay
	}
	if o.DelayFunction != nil {
		toSerialize["DelayFunction"] = o.DelayFunction
	}
	if o.Interval != nil {
		toSerialize["Interval"] = o.Interval
	}
	if o.MaxDelay != nil {
		toSerialize["MaxDelay"] = o.MaxDelay
	}
	if o.Unlimited != nil {
		toSerialize["Unlimited"] = o.Unlimited
	}
	return json.Marshal(toSerialize)
}

type NullableReschedulePolicy struct {
	value *ReschedulePolicy
	isSet bool
}

func (v NullableReschedulePolicy) Get() *ReschedulePolicy {
	return v.value
}

func (v *NullableReschedulePolicy) Set(val *ReschedulePolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableReschedulePolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableReschedulePolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReschedulePolicy(val *ReschedulePolicy) *NullableReschedulePolicy {
	return &NullableReschedulePolicy{value: val, isSet: true}
}

func (v NullableReschedulePolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReschedulePolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


