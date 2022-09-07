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

// CSINodeInfo struct for CSINodeInfo
type CSINodeInfo struct {
	AccessibleTopology NullableCSITopology `json:"AccessibleTopology,omitempty"`
	ID *string `json:"ID,omitempty"`
	MaxVolumes *int64 `json:"MaxVolumes,omitempty"`
	RequiresNodeStageVolume *bool `json:"RequiresNodeStageVolume,omitempty"`
	SupportsCondition *bool `json:"SupportsCondition,omitempty"`
	SupportsExpand *bool `json:"SupportsExpand,omitempty"`
	SupportsStats *bool `json:"SupportsStats,omitempty"`
}

// NewCSINodeInfo instantiates a new CSINodeInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCSINodeInfo() *CSINodeInfo {
	this := CSINodeInfo{}
	return &this
}

// NewCSINodeInfoWithDefaults instantiates a new CSINodeInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCSINodeInfoWithDefaults() *CSINodeInfo {
	this := CSINodeInfo{}
	return &this
}

// GetAccessibleTopology returns the AccessibleTopology field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CSINodeInfo) GetAccessibleTopology() CSITopology {
	if o == nil || o.AccessibleTopology.Get() == nil {
		var ret CSITopology
		return ret
	}
	return *o.AccessibleTopology.Get()
}

// GetAccessibleTopologyOk returns a tuple with the AccessibleTopology field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CSINodeInfo) GetAccessibleTopologyOk() (*CSITopology, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AccessibleTopology.Get(), o.AccessibleTopology.IsSet()
}

// HasAccessibleTopology returns a boolean if a field has been set.
func (o *CSINodeInfo) HasAccessibleTopology() bool {
	if o != nil && o.AccessibleTopology.IsSet() {
		return true
	}

	return false
}

// SetAccessibleTopology gets a reference to the given NullableCSITopology and assigns it to the AccessibleTopology field.
func (o *CSINodeInfo) SetAccessibleTopology(v CSITopology) {
	o.AccessibleTopology.Set(&v)
}
// SetAccessibleTopologyNil sets the value for AccessibleTopology to be an explicit nil
func (o *CSINodeInfo) SetAccessibleTopologyNil() {
	o.AccessibleTopology.Set(nil)
}

// UnsetAccessibleTopology ensures that no value is present for AccessibleTopology, not even an explicit nil
func (o *CSINodeInfo) UnsetAccessibleTopology() {
	o.AccessibleTopology.Unset()
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *CSINodeInfo) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSINodeInfo) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *CSINodeInfo) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *CSINodeInfo) SetID(v string) {
	o.ID = &v
}

// GetMaxVolumes returns the MaxVolumes field value if set, zero value otherwise.
func (o *CSINodeInfo) GetMaxVolumes() int64 {
	if o == nil || o.MaxVolumes == nil {
		var ret int64
		return ret
	}
	return *o.MaxVolumes
}

// GetMaxVolumesOk returns a tuple with the MaxVolumes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSINodeInfo) GetMaxVolumesOk() (*int64, bool) {
	if o == nil || o.MaxVolumes == nil {
		return nil, false
	}
	return o.MaxVolumes, true
}

// HasMaxVolumes returns a boolean if a field has been set.
func (o *CSINodeInfo) HasMaxVolumes() bool {
	if o != nil && o.MaxVolumes != nil {
		return true
	}

	return false
}

// SetMaxVolumes gets a reference to the given int64 and assigns it to the MaxVolumes field.
func (o *CSINodeInfo) SetMaxVolumes(v int64) {
	o.MaxVolumes = &v
}

// GetRequiresNodeStageVolume returns the RequiresNodeStageVolume field value if set, zero value otherwise.
func (o *CSINodeInfo) GetRequiresNodeStageVolume() bool {
	if o == nil || o.RequiresNodeStageVolume == nil {
		var ret bool
		return ret
	}
	return *o.RequiresNodeStageVolume
}

// GetRequiresNodeStageVolumeOk returns a tuple with the RequiresNodeStageVolume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSINodeInfo) GetRequiresNodeStageVolumeOk() (*bool, bool) {
	if o == nil || o.RequiresNodeStageVolume == nil {
		return nil, false
	}
	return o.RequiresNodeStageVolume, true
}

// HasRequiresNodeStageVolume returns a boolean if a field has been set.
func (o *CSINodeInfo) HasRequiresNodeStageVolume() bool {
	if o != nil && o.RequiresNodeStageVolume != nil {
		return true
	}

	return false
}

// SetRequiresNodeStageVolume gets a reference to the given bool and assigns it to the RequiresNodeStageVolume field.
func (o *CSINodeInfo) SetRequiresNodeStageVolume(v bool) {
	o.RequiresNodeStageVolume = &v
}

// GetSupportsCondition returns the SupportsCondition field value if set, zero value otherwise.
func (o *CSINodeInfo) GetSupportsCondition() bool {
	if o == nil || o.SupportsCondition == nil {
		var ret bool
		return ret
	}
	return *o.SupportsCondition
}

// GetSupportsConditionOk returns a tuple with the SupportsCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSINodeInfo) GetSupportsConditionOk() (*bool, bool) {
	if o == nil || o.SupportsCondition == nil {
		return nil, false
	}
	return o.SupportsCondition, true
}

// HasSupportsCondition returns a boolean if a field has been set.
func (o *CSINodeInfo) HasSupportsCondition() bool {
	if o != nil && o.SupportsCondition != nil {
		return true
	}

	return false
}

// SetSupportsCondition gets a reference to the given bool and assigns it to the SupportsCondition field.
func (o *CSINodeInfo) SetSupportsCondition(v bool) {
	o.SupportsCondition = &v
}

// GetSupportsExpand returns the SupportsExpand field value if set, zero value otherwise.
func (o *CSINodeInfo) GetSupportsExpand() bool {
	if o == nil || o.SupportsExpand == nil {
		var ret bool
		return ret
	}
	return *o.SupportsExpand
}

// GetSupportsExpandOk returns a tuple with the SupportsExpand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSINodeInfo) GetSupportsExpandOk() (*bool, bool) {
	if o == nil || o.SupportsExpand == nil {
		return nil, false
	}
	return o.SupportsExpand, true
}

// HasSupportsExpand returns a boolean if a field has been set.
func (o *CSINodeInfo) HasSupportsExpand() bool {
	if o != nil && o.SupportsExpand != nil {
		return true
	}

	return false
}

// SetSupportsExpand gets a reference to the given bool and assigns it to the SupportsExpand field.
func (o *CSINodeInfo) SetSupportsExpand(v bool) {
	o.SupportsExpand = &v
}

// GetSupportsStats returns the SupportsStats field value if set, zero value otherwise.
func (o *CSINodeInfo) GetSupportsStats() bool {
	if o == nil || o.SupportsStats == nil {
		var ret bool
		return ret
	}
	return *o.SupportsStats
}

// GetSupportsStatsOk returns a tuple with the SupportsStats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSINodeInfo) GetSupportsStatsOk() (*bool, bool) {
	if o == nil || o.SupportsStats == nil {
		return nil, false
	}
	return o.SupportsStats, true
}

// HasSupportsStats returns a boolean if a field has been set.
func (o *CSINodeInfo) HasSupportsStats() bool {
	if o != nil && o.SupportsStats != nil {
		return true
	}

	return false
}

// SetSupportsStats gets a reference to the given bool and assigns it to the SupportsStats field.
func (o *CSINodeInfo) SetSupportsStats(v bool) {
	o.SupportsStats = &v
}

func (o CSINodeInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessibleTopology.IsSet() {
		toSerialize["AccessibleTopology"] = o.AccessibleTopology.Get()
	}
	if o.ID != nil {
		toSerialize["ID"] = o.ID
	}
	if o.MaxVolumes != nil {
		toSerialize["MaxVolumes"] = o.MaxVolumes
	}
	if o.RequiresNodeStageVolume != nil {
		toSerialize["RequiresNodeStageVolume"] = o.RequiresNodeStageVolume
	}
	if o.SupportsCondition != nil {
		toSerialize["SupportsCondition"] = o.SupportsCondition
	}
	if o.SupportsExpand != nil {
		toSerialize["SupportsExpand"] = o.SupportsExpand
	}
	if o.SupportsStats != nil {
		toSerialize["SupportsStats"] = o.SupportsStats
	}
	return json.Marshal(toSerialize)
}

type NullableCSINodeInfo struct {
	value *CSINodeInfo
	isSet bool
}

func (v NullableCSINodeInfo) Get() *CSINodeInfo {
	return v.value
}

func (v *NullableCSINodeInfo) Set(val *CSINodeInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableCSINodeInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableCSINodeInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCSINodeInfo(val *CSINodeInfo) *NullableCSINodeInfo {
	return &NullableCSINodeInfo{value: val, isSet: true}
}

func (v NullableCSINodeInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCSINodeInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


