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
	"time"
)

// ACLToken struct for ACLToken
type ACLToken struct {
	AccessorID *string `json:"AccessorID,omitempty"`
	CreateIndex *int32 `json:"CreateIndex,omitempty"`
	CreateTime *time.Time `json:"CreateTime,omitempty"`
	Global *bool `json:"Global,omitempty"`
	ModifyIndex *int32 `json:"ModifyIndex,omitempty"`
	Name *string `json:"Name,omitempty"`
	Policies *[]string `json:"Policies,omitempty"`
	SecretID *string `json:"SecretID,omitempty"`
	Type *string `json:"Type,omitempty"`
}

// NewACLToken instantiates a new ACLToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewACLToken() *ACLToken {
	this := ACLToken{}
	return &this
}

// NewACLTokenWithDefaults instantiates a new ACLToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewACLTokenWithDefaults() *ACLToken {
	this := ACLToken{}
	return &this
}

// GetAccessorID returns the AccessorID field value if set, zero value otherwise.
func (o *ACLToken) GetAccessorID() string {
	if o == nil || o.AccessorID == nil {
		var ret string
		return ret
	}
	return *o.AccessorID
}

// GetAccessorIDOk returns a tuple with the AccessorID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACLToken) GetAccessorIDOk() (*string, bool) {
	if o == nil || o.AccessorID == nil {
		return nil, false
	}
	return o.AccessorID, true
}

// HasAccessorID returns a boolean if a field has been set.
func (o *ACLToken) HasAccessorID() bool {
	if o != nil && o.AccessorID != nil {
		return true
	}

	return false
}

// SetAccessorID gets a reference to the given string and assigns it to the AccessorID field.
func (o *ACLToken) SetAccessorID(v string) {
	o.AccessorID = &v
}

// GetCreateIndex returns the CreateIndex field value if set, zero value otherwise.
func (o *ACLToken) GetCreateIndex() int32 {
	if o == nil || o.CreateIndex == nil {
		var ret int32
		return ret
	}
	return *o.CreateIndex
}

// GetCreateIndexOk returns a tuple with the CreateIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACLToken) GetCreateIndexOk() (*int32, bool) {
	if o == nil || o.CreateIndex == nil {
		return nil, false
	}
	return o.CreateIndex, true
}

// HasCreateIndex returns a boolean if a field has been set.
func (o *ACLToken) HasCreateIndex() bool {
	if o != nil && o.CreateIndex != nil {
		return true
	}

	return false
}

// SetCreateIndex gets a reference to the given int32 and assigns it to the CreateIndex field.
func (o *ACLToken) SetCreateIndex(v int32) {
	o.CreateIndex = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *ACLToken) GetCreateTime() time.Time {
	if o == nil || o.CreateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACLToken) GetCreateTimeOk() (*time.Time, bool) {
	if o == nil || o.CreateTime == nil {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *ACLToken) HasCreateTime() bool {
	if o != nil && o.CreateTime != nil {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given time.Time and assigns it to the CreateTime field.
func (o *ACLToken) SetCreateTime(v time.Time) {
	o.CreateTime = &v
}

// GetGlobal returns the Global field value if set, zero value otherwise.
func (o *ACLToken) GetGlobal() bool {
	if o == nil || o.Global == nil {
		var ret bool
		return ret
	}
	return *o.Global
}

// GetGlobalOk returns a tuple with the Global field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACLToken) GetGlobalOk() (*bool, bool) {
	if o == nil || o.Global == nil {
		return nil, false
	}
	return o.Global, true
}

// HasGlobal returns a boolean if a field has been set.
func (o *ACLToken) HasGlobal() bool {
	if o != nil && o.Global != nil {
		return true
	}

	return false
}

// SetGlobal gets a reference to the given bool and assigns it to the Global field.
func (o *ACLToken) SetGlobal(v bool) {
	o.Global = &v
}

// GetModifyIndex returns the ModifyIndex field value if set, zero value otherwise.
func (o *ACLToken) GetModifyIndex() int32 {
	if o == nil || o.ModifyIndex == nil {
		var ret int32
		return ret
	}
	return *o.ModifyIndex
}

// GetModifyIndexOk returns a tuple with the ModifyIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACLToken) GetModifyIndexOk() (*int32, bool) {
	if o == nil || o.ModifyIndex == nil {
		return nil, false
	}
	return o.ModifyIndex, true
}

// HasModifyIndex returns a boolean if a field has been set.
func (o *ACLToken) HasModifyIndex() bool {
	if o != nil && o.ModifyIndex != nil {
		return true
	}

	return false
}

// SetModifyIndex gets a reference to the given int32 and assigns it to the ModifyIndex field.
func (o *ACLToken) SetModifyIndex(v int32) {
	o.ModifyIndex = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ACLToken) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACLToken) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ACLToken) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ACLToken) SetName(v string) {
	o.Name = &v
}

// GetPolicies returns the Policies field value if set, zero value otherwise.
func (o *ACLToken) GetPolicies() []string {
	if o == nil || o.Policies == nil {
		var ret []string
		return ret
	}
	return *o.Policies
}

// GetPoliciesOk returns a tuple with the Policies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACLToken) GetPoliciesOk() (*[]string, bool) {
	if o == nil || o.Policies == nil {
		return nil, false
	}
	return o.Policies, true
}

// HasPolicies returns a boolean if a field has been set.
func (o *ACLToken) HasPolicies() bool {
	if o != nil && o.Policies != nil {
		return true
	}

	return false
}

// SetPolicies gets a reference to the given []string and assigns it to the Policies field.
func (o *ACLToken) SetPolicies(v []string) {
	o.Policies = &v
}

// GetSecretID returns the SecretID field value if set, zero value otherwise.
func (o *ACLToken) GetSecretID() string {
	if o == nil || o.SecretID == nil {
		var ret string
		return ret
	}
	return *o.SecretID
}

// GetSecretIDOk returns a tuple with the SecretID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACLToken) GetSecretIDOk() (*string, bool) {
	if o == nil || o.SecretID == nil {
		return nil, false
	}
	return o.SecretID, true
}

// HasSecretID returns a boolean if a field has been set.
func (o *ACLToken) HasSecretID() bool {
	if o != nil && o.SecretID != nil {
		return true
	}

	return false
}

// SetSecretID gets a reference to the given string and assigns it to the SecretID field.
func (o *ACLToken) SetSecretID(v string) {
	o.SecretID = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ACLToken) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACLToken) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ACLToken) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ACLToken) SetType(v string) {
	o.Type = &v
}

func (o ACLToken) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessorID != nil {
		toSerialize["AccessorID"] = o.AccessorID
	}
	if o.CreateIndex != nil {
		toSerialize["CreateIndex"] = o.CreateIndex
	}
	if o.CreateTime != nil {
		toSerialize["CreateTime"] = o.CreateTime
	}
	if o.Global != nil {
		toSerialize["Global"] = o.Global
	}
	if o.ModifyIndex != nil {
		toSerialize["ModifyIndex"] = o.ModifyIndex
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Policies != nil {
		toSerialize["Policies"] = o.Policies
	}
	if o.SecretID != nil {
		toSerialize["SecretID"] = o.SecretID
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableACLToken struct {
	value *ACLToken
	isSet bool
}

func (v NullableACLToken) Get() *ACLToken {
	return v.value
}

func (v *NullableACLToken) Set(val *ACLToken) {
	v.value = val
	v.isSet = true
}

func (v NullableACLToken) IsSet() bool {
	return v.isSet
}

func (v *NullableACLToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableACLToken(val *ACLToken) *NullableACLToken {
	return &NullableACLToken{value: val, isSet: true}
}

func (v NullableACLToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableACLToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


