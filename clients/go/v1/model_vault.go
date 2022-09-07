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

// Vault struct for Vault
type Vault struct {
	ChangeMode *string `json:"ChangeMode,omitempty"`
	ChangeSignal *string `json:"ChangeSignal,omitempty"`
	Env *bool `json:"Env,omitempty"`
	Namespace *string `json:"Namespace,omitempty"`
	Policies []string `json:"Policies,omitempty"`
}

// NewVault instantiates a new Vault object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVault() *Vault {
	this := Vault{}
	return &this
}

// NewVaultWithDefaults instantiates a new Vault object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVaultWithDefaults() *Vault {
	this := Vault{}
	return &this
}

// GetChangeMode returns the ChangeMode field value if set, zero value otherwise.
func (o *Vault) GetChangeMode() string {
	if o == nil || o.ChangeMode == nil {
		var ret string
		return ret
	}
	return *o.ChangeMode
}

// GetChangeModeOk returns a tuple with the ChangeMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vault) GetChangeModeOk() (*string, bool) {
	if o == nil || o.ChangeMode == nil {
		return nil, false
	}
	return o.ChangeMode, true
}

// HasChangeMode returns a boolean if a field has been set.
func (o *Vault) HasChangeMode() bool {
	if o != nil && o.ChangeMode != nil {
		return true
	}

	return false
}

// SetChangeMode gets a reference to the given string and assigns it to the ChangeMode field.
func (o *Vault) SetChangeMode(v string) {
	o.ChangeMode = &v
}

// GetChangeSignal returns the ChangeSignal field value if set, zero value otherwise.
func (o *Vault) GetChangeSignal() string {
	if o == nil || o.ChangeSignal == nil {
		var ret string
		return ret
	}
	return *o.ChangeSignal
}

// GetChangeSignalOk returns a tuple with the ChangeSignal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vault) GetChangeSignalOk() (*string, bool) {
	if o == nil || o.ChangeSignal == nil {
		return nil, false
	}
	return o.ChangeSignal, true
}

// HasChangeSignal returns a boolean if a field has been set.
func (o *Vault) HasChangeSignal() bool {
	if o != nil && o.ChangeSignal != nil {
		return true
	}

	return false
}

// SetChangeSignal gets a reference to the given string and assigns it to the ChangeSignal field.
func (o *Vault) SetChangeSignal(v string) {
	o.ChangeSignal = &v
}

// GetEnv returns the Env field value if set, zero value otherwise.
func (o *Vault) GetEnv() bool {
	if o == nil || o.Env == nil {
		var ret bool
		return ret
	}
	return *o.Env
}

// GetEnvOk returns a tuple with the Env field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vault) GetEnvOk() (*bool, bool) {
	if o == nil || o.Env == nil {
		return nil, false
	}
	return o.Env, true
}

// HasEnv returns a boolean if a field has been set.
func (o *Vault) HasEnv() bool {
	if o != nil && o.Env != nil {
		return true
	}

	return false
}

// SetEnv gets a reference to the given bool and assigns it to the Env field.
func (o *Vault) SetEnv(v bool) {
	o.Env = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *Vault) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vault) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *Vault) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *Vault) SetNamespace(v string) {
	o.Namespace = &v
}

// GetPolicies returns the Policies field value if set, zero value otherwise.
func (o *Vault) GetPolicies() []string {
	if o == nil || o.Policies == nil {
		var ret []string
		return ret
	}
	return o.Policies
}

// GetPoliciesOk returns a tuple with the Policies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vault) GetPoliciesOk() ([]string, bool) {
	if o == nil || o.Policies == nil {
		return nil, false
	}
	return o.Policies, true
}

// HasPolicies returns a boolean if a field has been set.
func (o *Vault) HasPolicies() bool {
	if o != nil && o.Policies != nil {
		return true
	}

	return false
}

// SetPolicies gets a reference to the given []string and assigns it to the Policies field.
func (o *Vault) SetPolicies(v []string) {
	o.Policies = v
}

func (o Vault) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ChangeMode != nil {
		toSerialize["ChangeMode"] = o.ChangeMode
	}
	if o.ChangeSignal != nil {
		toSerialize["ChangeSignal"] = o.ChangeSignal
	}
	if o.Env != nil {
		toSerialize["Env"] = o.Env
	}
	if o.Namespace != nil {
		toSerialize["Namespace"] = o.Namespace
	}
	if o.Policies != nil {
		toSerialize["Policies"] = o.Policies
	}
	return json.Marshal(toSerialize)
}

type NullableVault struct {
	value *Vault
	isSet bool
}

func (v NullableVault) Get() *Vault {
	return v.value
}

func (v *NullableVault) Set(val *Vault) {
	v.value = val
	v.isSet = true
}

func (v NullableVault) IsSet() bool {
	return v.isSet
}

func (v *NullableVault) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVault(val *Vault) *NullableVault {
	return &NullableVault{value: val, isSet: true}
}

func (v NullableVault) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVault) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


