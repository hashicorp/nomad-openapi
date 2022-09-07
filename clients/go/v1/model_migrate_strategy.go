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

// MigrateStrategy struct for MigrateStrategy
type MigrateStrategy struct {
	HealthCheck *string `json:"HealthCheck,omitempty"`
	HealthyDeadline *int64 `json:"HealthyDeadline,omitempty"`
	MaxParallel *int32 `json:"MaxParallel,omitempty"`
	MinHealthyTime *int64 `json:"MinHealthyTime,omitempty"`
}

// NewMigrateStrategy instantiates a new MigrateStrategy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMigrateStrategy() *MigrateStrategy {
	this := MigrateStrategy{}
	return &this
}

// NewMigrateStrategyWithDefaults instantiates a new MigrateStrategy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMigrateStrategyWithDefaults() *MigrateStrategy {
	this := MigrateStrategy{}
	return &this
}

// GetHealthCheck returns the HealthCheck field value if set, zero value otherwise.
func (o *MigrateStrategy) GetHealthCheck() string {
	if o == nil || o.HealthCheck == nil {
		var ret string
		return ret
	}
	return *o.HealthCheck
}

// GetHealthCheckOk returns a tuple with the HealthCheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MigrateStrategy) GetHealthCheckOk() (*string, bool) {
	if o == nil || o.HealthCheck == nil {
		return nil, false
	}
	return o.HealthCheck, true
}

// HasHealthCheck returns a boolean if a field has been set.
func (o *MigrateStrategy) HasHealthCheck() bool {
	if o != nil && o.HealthCheck != nil {
		return true
	}

	return false
}

// SetHealthCheck gets a reference to the given string and assigns it to the HealthCheck field.
func (o *MigrateStrategy) SetHealthCheck(v string) {
	o.HealthCheck = &v
}

// GetHealthyDeadline returns the HealthyDeadline field value if set, zero value otherwise.
func (o *MigrateStrategy) GetHealthyDeadline() int64 {
	if o == nil || o.HealthyDeadline == nil {
		var ret int64
		return ret
	}
	return *o.HealthyDeadline
}

// GetHealthyDeadlineOk returns a tuple with the HealthyDeadline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MigrateStrategy) GetHealthyDeadlineOk() (*int64, bool) {
	if o == nil || o.HealthyDeadline == nil {
		return nil, false
	}
	return o.HealthyDeadline, true
}

// HasHealthyDeadline returns a boolean if a field has been set.
func (o *MigrateStrategy) HasHealthyDeadline() bool {
	if o != nil && o.HealthyDeadline != nil {
		return true
	}

	return false
}

// SetHealthyDeadline gets a reference to the given int64 and assigns it to the HealthyDeadline field.
func (o *MigrateStrategy) SetHealthyDeadline(v int64) {
	o.HealthyDeadline = &v
}

// GetMaxParallel returns the MaxParallel field value if set, zero value otherwise.
func (o *MigrateStrategy) GetMaxParallel() int32 {
	if o == nil || o.MaxParallel == nil {
		var ret int32
		return ret
	}
	return *o.MaxParallel
}

// GetMaxParallelOk returns a tuple with the MaxParallel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MigrateStrategy) GetMaxParallelOk() (*int32, bool) {
	if o == nil || o.MaxParallel == nil {
		return nil, false
	}
	return o.MaxParallel, true
}

// HasMaxParallel returns a boolean if a field has been set.
func (o *MigrateStrategy) HasMaxParallel() bool {
	if o != nil && o.MaxParallel != nil {
		return true
	}

	return false
}

// SetMaxParallel gets a reference to the given int32 and assigns it to the MaxParallel field.
func (o *MigrateStrategy) SetMaxParallel(v int32) {
	o.MaxParallel = &v
}

// GetMinHealthyTime returns the MinHealthyTime field value if set, zero value otherwise.
func (o *MigrateStrategy) GetMinHealthyTime() int64 {
	if o == nil || o.MinHealthyTime == nil {
		var ret int64
		return ret
	}
	return *o.MinHealthyTime
}

// GetMinHealthyTimeOk returns a tuple with the MinHealthyTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MigrateStrategy) GetMinHealthyTimeOk() (*int64, bool) {
	if o == nil || o.MinHealthyTime == nil {
		return nil, false
	}
	return o.MinHealthyTime, true
}

// HasMinHealthyTime returns a boolean if a field has been set.
func (o *MigrateStrategy) HasMinHealthyTime() bool {
	if o != nil && o.MinHealthyTime != nil {
		return true
	}

	return false
}

// SetMinHealthyTime gets a reference to the given int64 and assigns it to the MinHealthyTime field.
func (o *MigrateStrategy) SetMinHealthyTime(v int64) {
	o.MinHealthyTime = &v
}

func (o MigrateStrategy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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
	return json.Marshal(toSerialize)
}

type NullableMigrateStrategy struct {
	value *MigrateStrategy
	isSet bool
}

func (v NullableMigrateStrategy) Get() *MigrateStrategy {
	return v.value
}

func (v *NullableMigrateStrategy) Set(val *MigrateStrategy) {
	v.value = val
	v.isSet = true
}

func (v NullableMigrateStrategy) IsSet() bool {
	return v.isSet
}

func (v *NullableMigrateStrategy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMigrateStrategy(val *MigrateStrategy) *NullableMigrateStrategy {
	return &NullableMigrateStrategy{value: val, isSet: true}
}

func (v NullableMigrateStrategy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMigrateStrategy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


