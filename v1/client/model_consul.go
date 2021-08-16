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
)

// Consul struct for Consul
type Consul struct {
	Namespace *string `json:"Namespace,omitempty"`
}

// NewConsul instantiates a new Consul object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsul() *Consul {
	this := Consul{}
	return &this
}

// NewConsulWithDefaults instantiates a new Consul object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsulWithDefaults() *Consul {
	this := Consul{}
	return &this
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *Consul) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Consul) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *Consul) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *Consul) SetNamespace(v string) {
	o.Namespace = &v
}

func (o Consul) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Namespace != nil {
		toSerialize["Namespace"] = o.Namespace
	}
	return json.Marshal(toSerialize)
}

type NullableConsul struct {
	value *Consul
	isSet bool
}

func (v NullableConsul) Get() *Consul {
	return v.value
}

func (v *NullableConsul) Set(val *Consul) {
	v.value = val
	v.isSet = true
}

func (v NullableConsul) IsSet() bool {
	return v.isSet
}

func (v *NullableConsul) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsul(val *Consul) *NullableConsul {
	return &NullableConsul{value: val, isSet: true}
}

func (v NullableConsul) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsul) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


