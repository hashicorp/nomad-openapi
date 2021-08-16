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

// PortMapping struct for PortMapping
type PortMapping struct {
	HostIP *string `json:"HostIP,omitempty"`
	Label *string `json:"Label,omitempty"`
	To *int32 `json:"To,omitempty"`
	Value *int32 `json:"Value,omitempty"`
}

// NewPortMapping instantiates a new PortMapping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortMapping() *PortMapping {
	this := PortMapping{}
	return &this
}

// NewPortMappingWithDefaults instantiates a new PortMapping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortMappingWithDefaults() *PortMapping {
	this := PortMapping{}
	return &this
}

// GetHostIP returns the HostIP field value if set, zero value otherwise.
func (o *PortMapping) GetHostIP() string {
	if o == nil || o.HostIP == nil {
		var ret string
		return ret
	}
	return *o.HostIP
}

// GetHostIPOk returns a tuple with the HostIP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortMapping) GetHostIPOk() (*string, bool) {
	if o == nil || o.HostIP == nil {
		return nil, false
	}
	return o.HostIP, true
}

// HasHostIP returns a boolean if a field has been set.
func (o *PortMapping) HasHostIP() bool {
	if o != nil && o.HostIP != nil {
		return true
	}

	return false
}

// SetHostIP gets a reference to the given string and assigns it to the HostIP field.
func (o *PortMapping) SetHostIP(v string) {
	o.HostIP = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *PortMapping) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortMapping) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *PortMapping) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *PortMapping) SetLabel(v string) {
	o.Label = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *PortMapping) GetTo() int32 {
	if o == nil || o.To == nil {
		var ret int32
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortMapping) GetToOk() (*int32, bool) {
	if o == nil || o.To == nil {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *PortMapping) HasTo() bool {
	if o != nil && o.To != nil {
		return true
	}

	return false
}

// SetTo gets a reference to the given int32 and assigns it to the To field.
func (o *PortMapping) SetTo(v int32) {
	o.To = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *PortMapping) GetValue() int32 {
	if o == nil || o.Value == nil {
		var ret int32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortMapping) GetValueOk() (*int32, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *PortMapping) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given int32 and assigns it to the Value field.
func (o *PortMapping) SetValue(v int32) {
	o.Value = &v
}

func (o PortMapping) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.HostIP != nil {
		toSerialize["HostIP"] = o.HostIP
	}
	if o.Label != nil {
		toSerialize["Label"] = o.Label
	}
	if o.To != nil {
		toSerialize["To"] = o.To
	}
	if o.Value != nil {
		toSerialize["Value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullablePortMapping struct {
	value *PortMapping
	isSet bool
}

func (v NullablePortMapping) Get() *PortMapping {
	return v.value
}

func (v *NullablePortMapping) Set(val *PortMapping) {
	v.value = val
	v.isSet = true
}

func (v NullablePortMapping) IsSet() bool {
	return v.isSet
}

func (v *NullablePortMapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortMapping(val *PortMapping) *NullablePortMapping {
	return &NullablePortMapping{value: val, isSet: true}
}

func (v NullablePortMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


