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

// RestartPolicy struct for RestartPolicy
type RestartPolicy struct {
	Attempts *int32 `json:"Attempts,omitempty"`
	Delay *int64 `json:"Delay,omitempty"`
	Interval *int64 `json:"Interval,omitempty"`
	Mode *string `json:"Mode,omitempty"`
}

// NewRestartPolicy instantiates a new RestartPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestartPolicy() *RestartPolicy {
	this := RestartPolicy{}
	return &this
}

// NewRestartPolicyWithDefaults instantiates a new RestartPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestartPolicyWithDefaults() *RestartPolicy {
	this := RestartPolicy{}
	return &this
}

// GetAttempts returns the Attempts field value if set, zero value otherwise.
func (o *RestartPolicy) GetAttempts() int32 {
	if o == nil || o.Attempts == nil {
		var ret int32
		return ret
	}
	return *o.Attempts
}

// GetAttemptsOk returns a tuple with the Attempts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestartPolicy) GetAttemptsOk() (*int32, bool) {
	if o == nil || o.Attempts == nil {
		return nil, false
	}
	return o.Attempts, true
}

// HasAttempts returns a boolean if a field has been set.
func (o *RestartPolicy) HasAttempts() bool {
	if o != nil && o.Attempts != nil {
		return true
	}

	return false
}

// SetAttempts gets a reference to the given int32 and assigns it to the Attempts field.
func (o *RestartPolicy) SetAttempts(v int32) {
	o.Attempts = &v
}

// GetDelay returns the Delay field value if set, zero value otherwise.
func (o *RestartPolicy) GetDelay() int64 {
	if o == nil || o.Delay == nil {
		var ret int64
		return ret
	}
	return *o.Delay
}

// GetDelayOk returns a tuple with the Delay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestartPolicy) GetDelayOk() (*int64, bool) {
	if o == nil || o.Delay == nil {
		return nil, false
	}
	return o.Delay, true
}

// HasDelay returns a boolean if a field has been set.
func (o *RestartPolicy) HasDelay() bool {
	if o != nil && o.Delay != nil {
		return true
	}

	return false
}

// SetDelay gets a reference to the given int64 and assigns it to the Delay field.
func (o *RestartPolicy) SetDelay(v int64) {
	o.Delay = &v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *RestartPolicy) GetInterval() int64 {
	if o == nil || o.Interval == nil {
		var ret int64
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestartPolicy) GetIntervalOk() (*int64, bool) {
	if o == nil || o.Interval == nil {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *RestartPolicy) HasInterval() bool {
	if o != nil && o.Interval != nil {
		return true
	}

	return false
}

// SetInterval gets a reference to the given int64 and assigns it to the Interval field.
func (o *RestartPolicy) SetInterval(v int64) {
	o.Interval = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *RestartPolicy) GetMode() string {
	if o == nil || o.Mode == nil {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestartPolicy) GetModeOk() (*string, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *RestartPolicy) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *RestartPolicy) SetMode(v string) {
	o.Mode = &v
}

func (o RestartPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Attempts != nil {
		toSerialize["Attempts"] = o.Attempts
	}
	if o.Delay != nil {
		toSerialize["Delay"] = o.Delay
	}
	if o.Interval != nil {
		toSerialize["Interval"] = o.Interval
	}
	if o.Mode != nil {
		toSerialize["Mode"] = o.Mode
	}
	return json.Marshal(toSerialize)
}

type NullableRestartPolicy struct {
	value *RestartPolicy
	isSet bool
}

func (v NullableRestartPolicy) Get() *RestartPolicy {
	return v.value
}

func (v *NullableRestartPolicy) Set(val *RestartPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableRestartPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableRestartPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestartPolicy(val *RestartPolicy) *NullableRestartPolicy {
	return &NullableRestartPolicy{value: val, isSet: true}
}

func (v NullableRestartPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestartPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


