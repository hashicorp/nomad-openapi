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

// OperatorHealthReply struct for OperatorHealthReply
type OperatorHealthReply struct {
	FailureTolerance *int32 `json:"FailureTolerance,omitempty"`
	Healthy *bool `json:"Healthy,omitempty"`
	Servers *[]ServerHealth `json:"Servers,omitempty"`
}

// NewOperatorHealthReply instantiates a new OperatorHealthReply object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperatorHealthReply() *OperatorHealthReply {
	this := OperatorHealthReply{}
	return &this
}

// NewOperatorHealthReplyWithDefaults instantiates a new OperatorHealthReply object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperatorHealthReplyWithDefaults() *OperatorHealthReply {
	this := OperatorHealthReply{}
	return &this
}

// GetFailureTolerance returns the FailureTolerance field value if set, zero value otherwise.
func (o *OperatorHealthReply) GetFailureTolerance() int32 {
	if o == nil || o.FailureTolerance == nil {
		var ret int32
		return ret
	}
	return *o.FailureTolerance
}

// GetFailureToleranceOk returns a tuple with the FailureTolerance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatorHealthReply) GetFailureToleranceOk() (*int32, bool) {
	if o == nil || o.FailureTolerance == nil {
		return nil, false
	}
	return o.FailureTolerance, true
}

// HasFailureTolerance returns a boolean if a field has been set.
func (o *OperatorHealthReply) HasFailureTolerance() bool {
	if o != nil && o.FailureTolerance != nil {
		return true
	}

	return false
}

// SetFailureTolerance gets a reference to the given int32 and assigns it to the FailureTolerance field.
func (o *OperatorHealthReply) SetFailureTolerance(v int32) {
	o.FailureTolerance = &v
}

// GetHealthy returns the Healthy field value if set, zero value otherwise.
func (o *OperatorHealthReply) GetHealthy() bool {
	if o == nil || o.Healthy == nil {
		var ret bool
		return ret
	}
	return *o.Healthy
}

// GetHealthyOk returns a tuple with the Healthy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatorHealthReply) GetHealthyOk() (*bool, bool) {
	if o == nil || o.Healthy == nil {
		return nil, false
	}
	return o.Healthy, true
}

// HasHealthy returns a boolean if a field has been set.
func (o *OperatorHealthReply) HasHealthy() bool {
	if o != nil && o.Healthy != nil {
		return true
	}

	return false
}

// SetHealthy gets a reference to the given bool and assigns it to the Healthy field.
func (o *OperatorHealthReply) SetHealthy(v bool) {
	o.Healthy = &v
}

// GetServers returns the Servers field value if set, zero value otherwise.
func (o *OperatorHealthReply) GetServers() []ServerHealth {
	if o == nil || o.Servers == nil {
		var ret []ServerHealth
		return ret
	}
	return *o.Servers
}

// GetServersOk returns a tuple with the Servers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperatorHealthReply) GetServersOk() (*[]ServerHealth, bool) {
	if o == nil || o.Servers == nil {
		return nil, false
	}
	return o.Servers, true
}

// HasServers returns a boolean if a field has been set.
func (o *OperatorHealthReply) HasServers() bool {
	if o != nil && o.Servers != nil {
		return true
	}

	return false
}

// SetServers gets a reference to the given []ServerHealth and assigns it to the Servers field.
func (o *OperatorHealthReply) SetServers(v []ServerHealth) {
	o.Servers = &v
}

func (o OperatorHealthReply) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FailureTolerance != nil {
		toSerialize["FailureTolerance"] = o.FailureTolerance
	}
	if o.Healthy != nil {
		toSerialize["Healthy"] = o.Healthy
	}
	if o.Servers != nil {
		toSerialize["Servers"] = o.Servers
	}
	return json.Marshal(toSerialize)
}

type NullableOperatorHealthReply struct {
	value *OperatorHealthReply
	isSet bool
}

func (v NullableOperatorHealthReply) Get() *OperatorHealthReply {
	return v.value
}

func (v *NullableOperatorHealthReply) Set(val *OperatorHealthReply) {
	v.value = val
	v.isSet = true
}

func (v NullableOperatorHealthReply) IsSet() bool {
	return v.isSet
}

func (v *NullableOperatorHealthReply) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperatorHealthReply(val *OperatorHealthReply) *NullableOperatorHealthReply {
	return &NullableOperatorHealthReply{value: val, isSet: true}
}

func (v NullableOperatorHealthReply) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperatorHealthReply) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


