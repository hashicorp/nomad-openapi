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

// UpdateStrategy struct for UpdateStrategy
type UpdateStrategy struct {
	AutoPromote *bool `json:"AutoPromote,omitempty"`
	AutoRevert *bool `json:"AutoRevert,omitempty"`
	Canary *int32 `json:"Canary,omitempty"`
	HealthCheck *string `json:"HealthCheck,omitempty"`
	HealthyDeadline *int64 `json:"HealthyDeadline,omitempty"`
	MaxParallel *int32 `json:"MaxParallel,omitempty"`
	MinHealthyTime *int64 `json:"MinHealthyTime,omitempty"`
	ProgressDeadline *int64 `json:"ProgressDeadline,omitempty"`
	Stagger *int64 `json:"Stagger,omitempty"`
}

// NewUpdateStrategy instantiates a new UpdateStrategy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateStrategy() *UpdateStrategy {
	this := UpdateStrategy{}
	return &this
}

// NewUpdateStrategyWithDefaults instantiates a new UpdateStrategy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateStrategyWithDefaults() *UpdateStrategy {
	this := UpdateStrategy{}
	return &this
}

// GetAutoPromote returns the AutoPromote field value if set, zero value otherwise.
func (o *UpdateStrategy) GetAutoPromote() bool {
	if o == nil || o.AutoPromote == nil {
		var ret bool
		return ret
	}
	return *o.AutoPromote
}

// GetAutoPromoteOk returns a tuple with the AutoPromote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStrategy) GetAutoPromoteOk() (*bool, bool) {
	if o == nil || o.AutoPromote == nil {
		return nil, false
	}
	return o.AutoPromote, true
}

// HasAutoPromote returns a boolean if a field has been set.
func (o *UpdateStrategy) HasAutoPromote() bool {
	if o != nil && o.AutoPromote != nil {
		return true
	}

	return false
}

// SetAutoPromote gets a reference to the given bool and assigns it to the AutoPromote field.
func (o *UpdateStrategy) SetAutoPromote(v bool) {
	o.AutoPromote = &v
}

// GetAutoRevert returns the AutoRevert field value if set, zero value otherwise.
func (o *UpdateStrategy) GetAutoRevert() bool {
	if o == nil || o.AutoRevert == nil {
		var ret bool
		return ret
	}
	return *o.AutoRevert
}

// GetAutoRevertOk returns a tuple with the AutoRevert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStrategy) GetAutoRevertOk() (*bool, bool) {
	if o == nil || o.AutoRevert == nil {
		return nil, false
	}
	return o.AutoRevert, true
}

// HasAutoRevert returns a boolean if a field has been set.
func (o *UpdateStrategy) HasAutoRevert() bool {
	if o != nil && o.AutoRevert != nil {
		return true
	}

	return false
}

// SetAutoRevert gets a reference to the given bool and assigns it to the AutoRevert field.
func (o *UpdateStrategy) SetAutoRevert(v bool) {
	o.AutoRevert = &v
}

// GetCanary returns the Canary field value if set, zero value otherwise.
func (o *UpdateStrategy) GetCanary() int32 {
	if o == nil || o.Canary == nil {
		var ret int32
		return ret
	}
	return *o.Canary
}

// GetCanaryOk returns a tuple with the Canary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStrategy) GetCanaryOk() (*int32, bool) {
	if o == nil || o.Canary == nil {
		return nil, false
	}
	return o.Canary, true
}

// HasCanary returns a boolean if a field has been set.
func (o *UpdateStrategy) HasCanary() bool {
	if o != nil && o.Canary != nil {
		return true
	}

	return false
}

// SetCanary gets a reference to the given int32 and assigns it to the Canary field.
func (o *UpdateStrategy) SetCanary(v int32) {
	o.Canary = &v
}

// GetHealthCheck returns the HealthCheck field value if set, zero value otherwise.
func (o *UpdateStrategy) GetHealthCheck() string {
	if o == nil || o.HealthCheck == nil {
		var ret string
		return ret
	}
	return *o.HealthCheck
}

// GetHealthCheckOk returns a tuple with the HealthCheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStrategy) GetHealthCheckOk() (*string, bool) {
	if o == nil || o.HealthCheck == nil {
		return nil, false
	}
	return o.HealthCheck, true
}

// HasHealthCheck returns a boolean if a field has been set.
func (o *UpdateStrategy) HasHealthCheck() bool {
	if o != nil && o.HealthCheck != nil {
		return true
	}

	return false
}

// SetHealthCheck gets a reference to the given string and assigns it to the HealthCheck field.
func (o *UpdateStrategy) SetHealthCheck(v string) {
	o.HealthCheck = &v
}

// GetHealthyDeadline returns the HealthyDeadline field value if set, zero value otherwise.
func (o *UpdateStrategy) GetHealthyDeadline() int64 {
	if o == nil || o.HealthyDeadline == nil {
		var ret int64
		return ret
	}
	return *o.HealthyDeadline
}

// GetHealthyDeadlineOk returns a tuple with the HealthyDeadline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStrategy) GetHealthyDeadlineOk() (*int64, bool) {
	if o == nil || o.HealthyDeadline == nil {
		return nil, false
	}
	return o.HealthyDeadline, true
}

// HasHealthyDeadline returns a boolean if a field has been set.
func (o *UpdateStrategy) HasHealthyDeadline() bool {
	if o != nil && o.HealthyDeadline != nil {
		return true
	}

	return false
}

// SetHealthyDeadline gets a reference to the given int64 and assigns it to the HealthyDeadline field.
func (o *UpdateStrategy) SetHealthyDeadline(v int64) {
	o.HealthyDeadline = &v
}

// GetMaxParallel returns the MaxParallel field value if set, zero value otherwise.
func (o *UpdateStrategy) GetMaxParallel() int32 {
	if o == nil || o.MaxParallel == nil {
		var ret int32
		return ret
	}
	return *o.MaxParallel
}

// GetMaxParallelOk returns a tuple with the MaxParallel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStrategy) GetMaxParallelOk() (*int32, bool) {
	if o == nil || o.MaxParallel == nil {
		return nil, false
	}
	return o.MaxParallel, true
}

// HasMaxParallel returns a boolean if a field has been set.
func (o *UpdateStrategy) HasMaxParallel() bool {
	if o != nil && o.MaxParallel != nil {
		return true
	}

	return false
}

// SetMaxParallel gets a reference to the given int32 and assigns it to the MaxParallel field.
func (o *UpdateStrategy) SetMaxParallel(v int32) {
	o.MaxParallel = &v
}

// GetMinHealthyTime returns the MinHealthyTime field value if set, zero value otherwise.
func (o *UpdateStrategy) GetMinHealthyTime() int64 {
	if o == nil || o.MinHealthyTime == nil {
		var ret int64
		return ret
	}
	return *o.MinHealthyTime
}

// GetMinHealthyTimeOk returns a tuple with the MinHealthyTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStrategy) GetMinHealthyTimeOk() (*int64, bool) {
	if o == nil || o.MinHealthyTime == nil {
		return nil, false
	}
	return o.MinHealthyTime, true
}

// HasMinHealthyTime returns a boolean if a field has been set.
func (o *UpdateStrategy) HasMinHealthyTime() bool {
	if o != nil && o.MinHealthyTime != nil {
		return true
	}

	return false
}

// SetMinHealthyTime gets a reference to the given int64 and assigns it to the MinHealthyTime field.
func (o *UpdateStrategy) SetMinHealthyTime(v int64) {
	o.MinHealthyTime = &v
}

// GetProgressDeadline returns the ProgressDeadline field value if set, zero value otherwise.
func (o *UpdateStrategy) GetProgressDeadline() int64 {
	if o == nil || o.ProgressDeadline == nil {
		var ret int64
		return ret
	}
	return *o.ProgressDeadline
}

// GetProgressDeadlineOk returns a tuple with the ProgressDeadline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStrategy) GetProgressDeadlineOk() (*int64, bool) {
	if o == nil || o.ProgressDeadline == nil {
		return nil, false
	}
	return o.ProgressDeadline, true
}

// HasProgressDeadline returns a boolean if a field has been set.
func (o *UpdateStrategy) HasProgressDeadline() bool {
	if o != nil && o.ProgressDeadline != nil {
		return true
	}

	return false
}

// SetProgressDeadline gets a reference to the given int64 and assigns it to the ProgressDeadline field.
func (o *UpdateStrategy) SetProgressDeadline(v int64) {
	o.ProgressDeadline = &v
}

// GetStagger returns the Stagger field value if set, zero value otherwise.
func (o *UpdateStrategy) GetStagger() int64 {
	if o == nil || o.Stagger == nil {
		var ret int64
		return ret
	}
	return *o.Stagger
}

// GetStaggerOk returns a tuple with the Stagger field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStrategy) GetStaggerOk() (*int64, bool) {
	if o == nil || o.Stagger == nil {
		return nil, false
	}
	return o.Stagger, true
}

// HasStagger returns a boolean if a field has been set.
func (o *UpdateStrategy) HasStagger() bool {
	if o != nil && o.Stagger != nil {
		return true
	}

	return false
}

// SetStagger gets a reference to the given int64 and assigns it to the Stagger field.
func (o *UpdateStrategy) SetStagger(v int64) {
	o.Stagger = &v
}

func (o UpdateStrategy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AutoPromote != nil {
		toSerialize["AutoPromote"] = o.AutoPromote
	}
	if o.AutoRevert != nil {
		toSerialize["AutoRevert"] = o.AutoRevert
	}
	if o.Canary != nil {
		toSerialize["Canary"] = o.Canary
	}
	if o.HealthCheck != nil {
		toSerialize["HealthCheck"] = o.HealthCheck
	}
	if o.HealthyDeadline != nil {
		toSerialize["HealthyDeadline"] = o.HealthyDeadline
	}
	if o.MaxParallel != nil {
		toSerialize["MaxParallel"] = o.MaxParallel
	}
	if o.MinHealthyTime != nil {
		toSerialize["MinHealthyTime"] = o.MinHealthyTime
	}
	if o.ProgressDeadline != nil {
		toSerialize["ProgressDeadline"] = o.ProgressDeadline
	}
	if o.Stagger != nil {
		toSerialize["Stagger"] = o.Stagger
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateStrategy struct {
	value *UpdateStrategy
	isSet bool
}

func (v NullableUpdateStrategy) Get() *UpdateStrategy {
	return v.value
}

func (v *NullableUpdateStrategy) Set(val *UpdateStrategy) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateStrategy) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateStrategy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateStrategy(val *UpdateStrategy) *NullableUpdateStrategy {
	return &NullableUpdateStrategy{value: val, isSet: true}
}

func (v NullableUpdateStrategy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateStrategy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


