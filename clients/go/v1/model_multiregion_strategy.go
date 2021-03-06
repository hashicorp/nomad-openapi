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

// MultiregionStrategy struct for MultiregionStrategy
type MultiregionStrategy struct {
	MaxParallel *int32 `json:"MaxParallel,omitempty"`
	OnFailure *string `json:"OnFailure,omitempty"`
}

// NewMultiregionStrategy instantiates a new MultiregionStrategy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultiregionStrategy() *MultiregionStrategy {
	this := MultiregionStrategy{}
	return &this
}

// NewMultiregionStrategyWithDefaults instantiates a new MultiregionStrategy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultiregionStrategyWithDefaults() *MultiregionStrategy {
	this := MultiregionStrategy{}
	return &this
}

// GetMaxParallel returns the MaxParallel field value if set, zero value otherwise.
func (o *MultiregionStrategy) GetMaxParallel() int32 {
	if o == nil || o.MaxParallel == nil {
		var ret int32
		return ret
	}
	return *o.MaxParallel
}

// GetMaxParallelOk returns a tuple with the MaxParallel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiregionStrategy) GetMaxParallelOk() (*int32, bool) {
	if o == nil || o.MaxParallel == nil {
		return nil, false
	}
	return o.MaxParallel, true
}

// HasMaxParallel returns a boolean if a field has been set.
func (o *MultiregionStrategy) HasMaxParallel() bool {
	if o != nil && o.MaxParallel != nil {
		return true
	}

	return false
}

// SetMaxParallel gets a reference to the given int32 and assigns it to the MaxParallel field.
func (o *MultiregionStrategy) SetMaxParallel(v int32) {
	o.MaxParallel = &v
}

// GetOnFailure returns the OnFailure field value if set, zero value otherwise.
func (o *MultiregionStrategy) GetOnFailure() string {
	if o == nil || o.OnFailure == nil {
		var ret string
		return ret
	}
	return *o.OnFailure
}

// GetOnFailureOk returns a tuple with the OnFailure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiregionStrategy) GetOnFailureOk() (*string, bool) {
	if o == nil || o.OnFailure == nil {
		return nil, false
	}
	return o.OnFailure, true
}

// HasOnFailure returns a boolean if a field has been set.
func (o *MultiregionStrategy) HasOnFailure() bool {
	if o != nil && o.OnFailure != nil {
		return true
	}

	return false
}

// SetOnFailure gets a reference to the given string and assigns it to the OnFailure field.
func (o *MultiregionStrategy) SetOnFailure(v string) {
	o.OnFailure = &v
}

func (o MultiregionStrategy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MaxParallel != nil {
		toSerialize["MaxParallel"] = o.MaxParallel
	}
	if o.OnFailure != nil {
		toSerialize["OnFailure"] = o.OnFailure
	}
	return json.Marshal(toSerialize)
}

type NullableMultiregionStrategy struct {
	value *MultiregionStrategy
	isSet bool
}

func (v NullableMultiregionStrategy) Get() *MultiregionStrategy {
	return v.value
}

func (v *NullableMultiregionStrategy) Set(val *MultiregionStrategy) {
	v.value = val
	v.isSet = true
}

func (v NullableMultiregionStrategy) IsSet() bool {
	return v.isSet
}

func (v *NullableMultiregionStrategy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultiregionStrategy(val *MultiregionStrategy) *NullableMultiregionStrategy {
	return &NullableMultiregionStrategy{value: val, isSet: true}
}

func (v NullableMultiregionStrategy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultiregionStrategy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


