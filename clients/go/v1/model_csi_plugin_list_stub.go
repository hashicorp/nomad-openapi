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

// CSIPluginListStub struct for CSIPluginListStub
type CSIPluginListStub struct {
	ControllerRequired *bool `json:"ControllerRequired,omitempty"`
	ControllersExpected *int32 `json:"ControllersExpected,omitempty"`
	ControllersHealthy *int32 `json:"ControllersHealthy,omitempty"`
	CreateIndex *int32 `json:"CreateIndex,omitempty"`
	ID *string `json:"ID,omitempty"`
	ModifyIndex *int32 `json:"ModifyIndex,omitempty"`
	NodesExpected *int32 `json:"NodesExpected,omitempty"`
	NodesHealthy *int32 `json:"NodesHealthy,omitempty"`
	Provider *string `json:"Provider,omitempty"`
}

// NewCSIPluginListStub instantiates a new CSIPluginListStub object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCSIPluginListStub() *CSIPluginListStub {
	this := CSIPluginListStub{}
	return &this
}

// NewCSIPluginListStubWithDefaults instantiates a new CSIPluginListStub object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCSIPluginListStubWithDefaults() *CSIPluginListStub {
	this := CSIPluginListStub{}
	return &this
}

// GetControllerRequired returns the ControllerRequired field value if set, zero value otherwise.
func (o *CSIPluginListStub) GetControllerRequired() bool {
	if o == nil || o.ControllerRequired == nil {
		var ret bool
		return ret
	}
	return *o.ControllerRequired
}

// GetControllerRequiredOk returns a tuple with the ControllerRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSIPluginListStub) GetControllerRequiredOk() (*bool, bool) {
	if o == nil || o.ControllerRequired == nil {
		return nil, false
	}
	return o.ControllerRequired, true
}

// HasControllerRequired returns a boolean if a field has been set.
func (o *CSIPluginListStub) HasControllerRequired() bool {
	if o != nil && o.ControllerRequired != nil {
		return true
	}

	return false
}

// SetControllerRequired gets a reference to the given bool and assigns it to the ControllerRequired field.
func (o *CSIPluginListStub) SetControllerRequired(v bool) {
	o.ControllerRequired = &v
}

// GetControllersExpected returns the ControllersExpected field value if set, zero value otherwise.
func (o *CSIPluginListStub) GetControllersExpected() int32 {
	if o == nil || o.ControllersExpected == nil {
		var ret int32
		return ret
	}
	return *o.ControllersExpected
}

// GetControllersExpectedOk returns a tuple with the ControllersExpected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSIPluginListStub) GetControllersExpectedOk() (*int32, bool) {
	if o == nil || o.ControllersExpected == nil {
		return nil, false
	}
	return o.ControllersExpected, true
}

// HasControllersExpected returns a boolean if a field has been set.
func (o *CSIPluginListStub) HasControllersExpected() bool {
	if o != nil && o.ControllersExpected != nil {
		return true
	}

	return false
}

// SetControllersExpected gets a reference to the given int32 and assigns it to the ControllersExpected field.
func (o *CSIPluginListStub) SetControllersExpected(v int32) {
	o.ControllersExpected = &v
}

// GetControllersHealthy returns the ControllersHealthy field value if set, zero value otherwise.
func (o *CSIPluginListStub) GetControllersHealthy() int32 {
	if o == nil || o.ControllersHealthy == nil {
		var ret int32
		return ret
	}
	return *o.ControllersHealthy
}

// GetControllersHealthyOk returns a tuple with the ControllersHealthy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSIPluginListStub) GetControllersHealthyOk() (*int32, bool) {
	if o == nil || o.ControllersHealthy == nil {
		return nil, false
	}
	return o.ControllersHealthy, true
}

// HasControllersHealthy returns a boolean if a field has been set.
func (o *CSIPluginListStub) HasControllersHealthy() bool {
	if o != nil && o.ControllersHealthy != nil {
		return true
	}

	return false
}

// SetControllersHealthy gets a reference to the given int32 and assigns it to the ControllersHealthy field.
func (o *CSIPluginListStub) SetControllersHealthy(v int32) {
	o.ControllersHealthy = &v
}

// GetCreateIndex returns the CreateIndex field value if set, zero value otherwise.
func (o *CSIPluginListStub) GetCreateIndex() int32 {
	if o == nil || o.CreateIndex == nil {
		var ret int32
		return ret
	}
	return *o.CreateIndex
}

// GetCreateIndexOk returns a tuple with the CreateIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSIPluginListStub) GetCreateIndexOk() (*int32, bool) {
	if o == nil || o.CreateIndex == nil {
		return nil, false
	}
	return o.CreateIndex, true
}

// HasCreateIndex returns a boolean if a field has been set.
func (o *CSIPluginListStub) HasCreateIndex() bool {
	if o != nil && o.CreateIndex != nil {
		return true
	}

	return false
}

// SetCreateIndex gets a reference to the given int32 and assigns it to the CreateIndex field.
func (o *CSIPluginListStub) SetCreateIndex(v int32) {
	o.CreateIndex = &v
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *CSIPluginListStub) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSIPluginListStub) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *CSIPluginListStub) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *CSIPluginListStub) SetID(v string) {
	o.ID = &v
}

// GetModifyIndex returns the ModifyIndex field value if set, zero value otherwise.
func (o *CSIPluginListStub) GetModifyIndex() int32 {
	if o == nil || o.ModifyIndex == nil {
		var ret int32
		return ret
	}
	return *o.ModifyIndex
}

// GetModifyIndexOk returns a tuple with the ModifyIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSIPluginListStub) GetModifyIndexOk() (*int32, bool) {
	if o == nil || o.ModifyIndex == nil {
		return nil, false
	}
	return o.ModifyIndex, true
}

// HasModifyIndex returns a boolean if a field has been set.
func (o *CSIPluginListStub) HasModifyIndex() bool {
	if o != nil && o.ModifyIndex != nil {
		return true
	}

	return false
}

// SetModifyIndex gets a reference to the given int32 and assigns it to the ModifyIndex field.
func (o *CSIPluginListStub) SetModifyIndex(v int32) {
	o.ModifyIndex = &v
}

// GetNodesExpected returns the NodesExpected field value if set, zero value otherwise.
func (o *CSIPluginListStub) GetNodesExpected() int32 {
	if o == nil || o.NodesExpected == nil {
		var ret int32
		return ret
	}
	return *o.NodesExpected
}

// GetNodesExpectedOk returns a tuple with the NodesExpected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSIPluginListStub) GetNodesExpectedOk() (*int32, bool) {
	if o == nil || o.NodesExpected == nil {
		return nil, false
	}
	return o.NodesExpected, true
}

// HasNodesExpected returns a boolean if a field has been set.
func (o *CSIPluginListStub) HasNodesExpected() bool {
	if o != nil && o.NodesExpected != nil {
		return true
	}

	return false
}

// SetNodesExpected gets a reference to the given int32 and assigns it to the NodesExpected field.
func (o *CSIPluginListStub) SetNodesExpected(v int32) {
	o.NodesExpected = &v
}

// GetNodesHealthy returns the NodesHealthy field value if set, zero value otherwise.
func (o *CSIPluginListStub) GetNodesHealthy() int32 {
	if o == nil || o.NodesHealthy == nil {
		var ret int32
		return ret
	}
	return *o.NodesHealthy
}

// GetNodesHealthyOk returns a tuple with the NodesHealthy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSIPluginListStub) GetNodesHealthyOk() (*int32, bool) {
	if o == nil || o.NodesHealthy == nil {
		return nil, false
	}
	return o.NodesHealthy, true
}

// HasNodesHealthy returns a boolean if a field has been set.
func (o *CSIPluginListStub) HasNodesHealthy() bool {
	if o != nil && o.NodesHealthy != nil {
		return true
	}

	return false
}

// SetNodesHealthy gets a reference to the given int32 and assigns it to the NodesHealthy field.
func (o *CSIPluginListStub) SetNodesHealthy(v int32) {
	o.NodesHealthy = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *CSIPluginListStub) GetProvider() string {
	if o == nil || o.Provider == nil {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSIPluginListStub) GetProviderOk() (*string, bool) {
	if o == nil || o.Provider == nil {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *CSIPluginListStub) HasProvider() bool {
	if o != nil && o.Provider != nil {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *CSIPluginListStub) SetProvider(v string) {
	o.Provider = &v
}

func (o CSIPluginListStub) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ControllerRequired != nil {
		toSerialize["ControllerRequired"] = o.ControllerRequired
	}
	if o.ControllersExpected != nil {
		toSerialize["ControllersExpected"] = o.ControllersExpected
	}
	if o.ControllersHealthy != nil {
		toSerialize["ControllersHealthy"] = o.ControllersHealthy
	}
	if o.CreateIndex != nil {
		toSerialize["CreateIndex"] = o.CreateIndex
	}
	if o.ID != nil {
		toSerialize["ID"] = o.ID
	}
	if o.ModifyIndex != nil {
		toSerialize["ModifyIndex"] = o.ModifyIndex
	}
	if o.NodesExpected != nil {
		toSerialize["NodesExpected"] = o.NodesExpected
	}
	if o.NodesHealthy != nil {
		toSerialize["NodesHealthy"] = o.NodesHealthy
	}
	if o.Provider != nil {
		toSerialize["Provider"] = o.Provider
	}
	return json.Marshal(toSerialize)
}

type NullableCSIPluginListStub struct {
	value *CSIPluginListStub
	isSet bool
}

func (v NullableCSIPluginListStub) Get() *CSIPluginListStub {
	return v.value
}

func (v *NullableCSIPluginListStub) Set(val *CSIPluginListStub) {
	v.value = val
	v.isSet = true
}

func (v NullableCSIPluginListStub) IsSet() bool {
	return v.isSet
}

func (v *NullableCSIPluginListStub) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCSIPluginListStub(val *CSIPluginListStub) *NullableCSIPluginListStub {
	return &NullableCSIPluginListStub{value: val, isSet: true}
}

func (v NullableCSIPluginListStub) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCSIPluginListStub) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

