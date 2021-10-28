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

// CSIPlugin struct for CSIPlugin
type CSIPlugin struct {
	Allocations *[]AllocationListStub `json:"Allocations,omitempty"`
	ControllerRequired *bool `json:"ControllerRequired,omitempty"`
	Controllers *map[string]CSIInfo `json:"Controllers,omitempty"`
	ControllersExpected *int32 `json:"ControllersExpected,omitempty"`
	ControllersHealthy *int32 `json:"ControllersHealthy,omitempty"`
	CreateIndex *int32 `json:"CreateIndex,omitempty"`
	ID *string `json:"ID,omitempty"`
	ModifyIndex *int32 `json:"ModifyIndex,omitempty"`
	Nodes *map[string]CSIInfo `json:"Nodes,omitempty"`
	NodesExpected *int32 `json:"NodesExpected,omitempty"`
	NodesHealthy *int32 `json:"NodesHealthy,omitempty"`
	Provider *string `json:"Provider,omitempty"`
	Version *string `json:"Version,omitempty"`
}

// NewCSIPlugin instantiates a new CSIPlugin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCSIPlugin() *CSIPlugin {
	this := CSIPlugin{}
	return &this
}

// NewCSIPluginWithDefaults instantiates a new CSIPlugin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCSIPluginWithDefaults() *CSIPlugin {
	this := CSIPlugin{}
	return &this
}

// GetAllocations returns the Allocations field value if set, zero value otherwise.
func (o *CSIPlugin) GetAllocations() []AllocationListStub {
	if o == nil || o.Allocations == nil {
		var ret []AllocationListStub
		return ret
	}
	return *o.Allocations
}

// GetAllocationsOk returns a tuple with the Allocations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSIPlugin) GetAllocationsOk() (*[]AllocationListStub, bool) {
	if o == nil || o.Allocations == nil {
		return nil, false
	}
	return o.Allocations, true
}

// HasAllocations returns a boolean if a field has been set.
func (o *CSIPlugin) HasAllocations() bool {
	if o != nil && o.Allocations != nil {
		return true
	}

	return false
}

// SetAllocations gets a reference to the given []AllocationListStub and assigns it to the Allocations field.
func (o *CSIPlugin) SetAllocations(v []AllocationListStub) {
	o.Allocations = &v
}

// GetControllerRequired returns the ControllerRequired field value if set, zero value otherwise.
func (o *CSIPlugin) GetControllerRequired() bool {
	if o == nil || o.ControllerRequired == nil {
		var ret bool
		return ret
	}
	return *o.ControllerRequired
}

// GetControllerRequiredOk returns a tuple with the ControllerRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSIPlugin) GetControllerRequiredOk() (*bool, bool) {
	if o == nil || o.ControllerRequired == nil {
		return nil, false
	}
	return o.ControllerRequired, true
}

// HasControllerRequired returns a boolean if a field has been set.
func (o *CSIPlugin) HasControllerRequired() bool {
	if o != nil && o.ControllerRequired != nil {
		return true
	}

	return false
}

// SetControllerRequired gets a reference to the given bool and assigns it to the ControllerRequired field.
func (o *CSIPlugin) SetControllerRequired(v bool) {
	o.ControllerRequired = &v
}

// GetControllers returns the Controllers field value if set, zero value otherwise.
func (o *CSIPlugin) GetControllers() map[string]CSIInfo {
	if o == nil || o.Controllers == nil {
		var ret map[string]CSIInfo
		return ret
	}
	return *o.Controllers
}

// GetControllersOk returns a tuple with the Controllers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSIPlugin) GetControllersOk() (*map[string]CSIInfo, bool) {
	if o == nil || o.Controllers == nil {
		return nil, false
	}
	return o.Controllers, true
}

// HasControllers returns a boolean if a field has been set.
func (o *CSIPlugin) HasControllers() bool {
	if o != nil && o.Controllers != nil {
		return true
	}

	return false
}

// SetControllers gets a reference to the given map[string]CSIInfo and assigns it to the Controllers field.
func (o *CSIPlugin) SetControllers(v map[string]CSIInfo) {
	o.Controllers = &v
}

// GetControllersExpected returns the ControllersExpected field value if set, zero value otherwise.
func (o *CSIPlugin) GetControllersExpected() int32 {
	if o == nil || o.ControllersExpected == nil {
		var ret int32
		return ret
	}
	return *o.ControllersExpected
}

// GetControllersExpectedOk returns a tuple with the ControllersExpected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSIPlugin) GetControllersExpectedOk() (*int32, bool) {
	if o == nil || o.ControllersExpected == nil {
		return nil, false
	}
	return o.ControllersExpected, true
}

// HasControllersExpected returns a boolean if a field has been set.
func (o *CSIPlugin) HasControllersExpected() bool {
	if o != nil && o.ControllersExpected != nil {
		return true
	}

	return false
}

// SetControllersExpected gets a reference to the given int32 and assigns it to the ControllersExpected field.
func (o *CSIPlugin) SetControllersExpected(v int32) {
	o.ControllersExpected = &v
}

// GetControllersHealthy returns the ControllersHealthy field value if set, zero value otherwise.
func (o *CSIPlugin) GetControllersHealthy() int32 {
	if o == nil || o.ControllersHealthy == nil {
		var ret int32
		return ret
	}
	return *o.ControllersHealthy
}

// GetControllersHealthyOk returns a tuple with the ControllersHealthy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSIPlugin) GetControllersHealthyOk() (*int32, bool) {
	if o == nil || o.ControllersHealthy == nil {
		return nil, false
	}
	return o.ControllersHealthy, true
}

// HasControllersHealthy returns a boolean if a field has been set.
func (o *CSIPlugin) HasControllersHealthy() bool {
	if o != nil && o.ControllersHealthy != nil {
		return true
	}

	return false
}

// SetControllersHealthy gets a reference to the given int32 and assigns it to the ControllersHealthy field.
func (o *CSIPlugin) SetControllersHealthy(v int32) {
	o.ControllersHealthy = &v
}

// GetCreateIndex returns the CreateIndex field value if set, zero value otherwise.
func (o *CSIPlugin) GetCreateIndex() int32 {
	if o == nil || o.CreateIndex == nil {
		var ret int32
		return ret
	}
	return *o.CreateIndex
}

// GetCreateIndexOk returns a tuple with the CreateIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSIPlugin) GetCreateIndexOk() (*int32, bool) {
	if o == nil || o.CreateIndex == nil {
		return nil, false
	}
	return o.CreateIndex, true
}

// HasCreateIndex returns a boolean if a field has been set.
func (o *CSIPlugin) HasCreateIndex() bool {
	if o != nil && o.CreateIndex != nil {
		return true
	}

	return false
}

// SetCreateIndex gets a reference to the given int32 and assigns it to the CreateIndex field.
func (o *CSIPlugin) SetCreateIndex(v int32) {
	o.CreateIndex = &v
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *CSIPlugin) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSIPlugin) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *CSIPlugin) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *CSIPlugin) SetID(v string) {
	o.ID = &v
}

// GetModifyIndex returns the ModifyIndex field value if set, zero value otherwise.
func (o *CSIPlugin) GetModifyIndex() int32 {
	if o == nil || o.ModifyIndex == nil {
		var ret int32
		return ret
	}
	return *o.ModifyIndex
}

// GetModifyIndexOk returns a tuple with the ModifyIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSIPlugin) GetModifyIndexOk() (*int32, bool) {
	if o == nil || o.ModifyIndex == nil {
		return nil, false
	}
	return o.ModifyIndex, true
}

// HasModifyIndex returns a boolean if a field has been set.
func (o *CSIPlugin) HasModifyIndex() bool {
	if o != nil && o.ModifyIndex != nil {
		return true
	}

	return false
}

// SetModifyIndex gets a reference to the given int32 and assigns it to the ModifyIndex field.
func (o *CSIPlugin) SetModifyIndex(v int32) {
	o.ModifyIndex = &v
}

// GetNodes returns the Nodes field value if set, zero value otherwise.
func (o *CSIPlugin) GetNodes() map[string]CSIInfo {
	if o == nil || o.Nodes == nil {
		var ret map[string]CSIInfo
		return ret
	}
	return *o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSIPlugin) GetNodesOk() (*map[string]CSIInfo, bool) {
	if o == nil || o.Nodes == nil {
		return nil, false
	}
	return o.Nodes, true
}

// HasNodes returns a boolean if a field has been set.
func (o *CSIPlugin) HasNodes() bool {
	if o != nil && o.Nodes != nil {
		return true
	}

	return false
}

// SetNodes gets a reference to the given map[string]CSIInfo and assigns it to the Nodes field.
func (o *CSIPlugin) SetNodes(v map[string]CSIInfo) {
	o.Nodes = &v
}

// GetNodesExpected returns the NodesExpected field value if set, zero value otherwise.
func (o *CSIPlugin) GetNodesExpected() int32 {
	if o == nil || o.NodesExpected == nil {
		var ret int32
		return ret
	}
	return *o.NodesExpected
}

// GetNodesExpectedOk returns a tuple with the NodesExpected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSIPlugin) GetNodesExpectedOk() (*int32, bool) {
	if o == nil || o.NodesExpected == nil {
		return nil, false
	}
	return o.NodesExpected, true
}

// HasNodesExpected returns a boolean if a field has been set.
func (o *CSIPlugin) HasNodesExpected() bool {
	if o != nil && o.NodesExpected != nil {
		return true
	}

	return false
}

// SetNodesExpected gets a reference to the given int32 and assigns it to the NodesExpected field.
func (o *CSIPlugin) SetNodesExpected(v int32) {
	o.NodesExpected = &v
}

// GetNodesHealthy returns the NodesHealthy field value if set, zero value otherwise.
func (o *CSIPlugin) GetNodesHealthy() int32 {
	if o == nil || o.NodesHealthy == nil {
		var ret int32
		return ret
	}
	return *o.NodesHealthy
}

// GetNodesHealthyOk returns a tuple with the NodesHealthy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSIPlugin) GetNodesHealthyOk() (*int32, bool) {
	if o == nil || o.NodesHealthy == nil {
		return nil, false
	}
	return o.NodesHealthy, true
}

// HasNodesHealthy returns a boolean if a field has been set.
func (o *CSIPlugin) HasNodesHealthy() bool {
	if o != nil && o.NodesHealthy != nil {
		return true
	}

	return false
}

// SetNodesHealthy gets a reference to the given int32 and assigns it to the NodesHealthy field.
func (o *CSIPlugin) SetNodesHealthy(v int32) {
	o.NodesHealthy = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *CSIPlugin) GetProvider() string {
	if o == nil || o.Provider == nil {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSIPlugin) GetProviderOk() (*string, bool) {
	if o == nil || o.Provider == nil {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *CSIPlugin) HasProvider() bool {
	if o != nil && o.Provider != nil {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *CSIPlugin) SetProvider(v string) {
	o.Provider = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *CSIPlugin) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSIPlugin) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *CSIPlugin) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *CSIPlugin) SetVersion(v string) {
	o.Version = &v
}

func (o CSIPlugin) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Allocations != nil {
		toSerialize["Allocations"] = o.Allocations
	}
	if o.ControllerRequired != nil {
		toSerialize["ControllerRequired"] = o.ControllerRequired
	}
	if o.Controllers != nil {
		toSerialize["Controllers"] = o.Controllers
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
	if o.Nodes != nil {
		toSerialize["Nodes"] = o.Nodes
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
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableCSIPlugin struct {
	value *CSIPlugin
	isSet bool
}

func (v NullableCSIPlugin) Get() *CSIPlugin {
	return v.value
}

func (v *NullableCSIPlugin) Set(val *CSIPlugin) {
	v.value = val
	v.isSet = true
}

func (v NullableCSIPlugin) IsSet() bool {
	return v.isSet
}

func (v *NullableCSIPlugin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCSIPlugin(val *CSIPlugin) *NullableCSIPlugin {
	return &NullableCSIPlugin{value: val, isSet: true}
}

func (v NullableCSIPlugin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCSIPlugin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


