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

// CSIControllerInfo struct for CSIControllerInfo
type CSIControllerInfo struct {
	SupportsAttachDetach *bool `json:"SupportsAttachDetach,omitempty"`
	SupportsClone *bool `json:"SupportsClone,omitempty"`
	SupportsCondition *bool `json:"SupportsCondition,omitempty"`
	SupportsCreateDelete *bool `json:"SupportsCreateDelete,omitempty"`
	SupportsCreateDeleteSnapshot *bool `json:"SupportsCreateDeleteSnapshot,omitempty"`
	SupportsExpand *bool `json:"SupportsExpand,omitempty"`
	SupportsGet *bool `json:"SupportsGet,omitempty"`
	SupportsGetCapacity *bool `json:"SupportsGetCapacity,omitempty"`
	SupportsListSnapshots *bool `json:"SupportsListSnapshots,omitempty"`
	SupportsListVolumes *bool `json:"SupportsListVolumes,omitempty"`
	SupportsListVolumesAttachedNodes *bool `json:"SupportsListVolumesAttachedNodes,omitempty"`
	SupportsReadOnlyAttach *bool `json:"SupportsReadOnlyAttach,omitempty"`
}

// NewCSIControllerInfo instantiates a new CSIControllerInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCSIControllerInfo() *CSIControllerInfo {
	this := CSIControllerInfo{}
	return &this
}

// NewCSIControllerInfoWithDefaults instantiates a new CSIControllerInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCSIControllerInfoWithDefaults() *CSIControllerInfo {
	this := CSIControllerInfo{}
	return &this
}

// GetSupportsAttachDetach returns the SupportsAttachDetach field value if set, zero value otherwise.
func (o *CSIControllerInfo) GetSupportsAttachDetach() bool {
	if o == nil || o.SupportsAttachDetach == nil {
		var ret bool
		return ret
	}
	return *o.SupportsAttachDetach
}

// GetSupportsAttachDetachOk returns a tuple with the SupportsAttachDetach field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSIControllerInfo) GetSupportsAttachDetachOk() (*bool, bool) {
	if o == nil || o.SupportsAttachDetach == nil {
		return nil, false
	}
	return o.SupportsAttachDetach, true
}

// HasSupportsAttachDetach returns a boolean if a field has been set.
func (o *CSIControllerInfo) HasSupportsAttachDetach() bool {
	if o != nil && o.SupportsAttachDetach != nil {
		return true
	}

	return false
}

// SetSupportsAttachDetach gets a reference to the given bool and assigns it to the SupportsAttachDetach field.
func (o *CSIControllerInfo) SetSupportsAttachDetach(v bool) {
	o.SupportsAttachDetach = &v
}

// GetSupportsClone returns the SupportsClone field value if set, zero value otherwise.
func (o *CSIControllerInfo) GetSupportsClone() bool {
	if o == nil || o.SupportsClone == nil {
		var ret bool
		return ret
	}
	return *o.SupportsClone
}

// GetSupportsCloneOk returns a tuple with the SupportsClone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSIControllerInfo) GetSupportsCloneOk() (*bool, bool) {
	if o == nil || o.SupportsClone == nil {
		return nil, false
	}
	return o.SupportsClone, true
}

// HasSupportsClone returns a boolean if a field has been set.
func (o *CSIControllerInfo) HasSupportsClone() bool {
	if o != nil && o.SupportsClone != nil {
		return true
	}

	return false
}

// SetSupportsClone gets a reference to the given bool and assigns it to the SupportsClone field.
func (o *CSIControllerInfo) SetSupportsClone(v bool) {
	o.SupportsClone = &v
}

// GetSupportsCondition returns the SupportsCondition field value if set, zero value otherwise.
func (o *CSIControllerInfo) GetSupportsCondition() bool {
	if o == nil || o.SupportsCondition == nil {
		var ret bool
		return ret
	}
	return *o.SupportsCondition
}

// GetSupportsConditionOk returns a tuple with the SupportsCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSIControllerInfo) GetSupportsConditionOk() (*bool, bool) {
	if o == nil || o.SupportsCondition == nil {
		return nil, false
	}
	return o.SupportsCondition, true
}

// HasSupportsCondition returns a boolean if a field has been set.
func (o *CSIControllerInfo) HasSupportsCondition() bool {
	if o != nil && o.SupportsCondition != nil {
		return true
	}

	return false
}

// SetSupportsCondition gets a reference to the given bool and assigns it to the SupportsCondition field.
func (o *CSIControllerInfo) SetSupportsCondition(v bool) {
	o.SupportsCondition = &v
}

// GetSupportsCreateDelete returns the SupportsCreateDelete field value if set, zero value otherwise.
func (o *CSIControllerInfo) GetSupportsCreateDelete() bool {
	if o == nil || o.SupportsCreateDelete == nil {
		var ret bool
		return ret
	}
	return *o.SupportsCreateDelete
}

// GetSupportsCreateDeleteOk returns a tuple with the SupportsCreateDelete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSIControllerInfo) GetSupportsCreateDeleteOk() (*bool, bool) {
	if o == nil || o.SupportsCreateDelete == nil {
		return nil, false
	}
	return o.SupportsCreateDelete, true
}

// HasSupportsCreateDelete returns a boolean if a field has been set.
func (o *CSIControllerInfo) HasSupportsCreateDelete() bool {
	if o != nil && o.SupportsCreateDelete != nil {
		return true
	}

	return false
}

// SetSupportsCreateDelete gets a reference to the given bool and assigns it to the SupportsCreateDelete field.
func (o *CSIControllerInfo) SetSupportsCreateDelete(v bool) {
	o.SupportsCreateDelete = &v
}

// GetSupportsCreateDeleteSnapshot returns the SupportsCreateDeleteSnapshot field value if set, zero value otherwise.
func (o *CSIControllerInfo) GetSupportsCreateDeleteSnapshot() bool {
	if o == nil || o.SupportsCreateDeleteSnapshot == nil {
		var ret bool
		return ret
	}
	return *o.SupportsCreateDeleteSnapshot
}

// GetSupportsCreateDeleteSnapshotOk returns a tuple with the SupportsCreateDeleteSnapshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSIControllerInfo) GetSupportsCreateDeleteSnapshotOk() (*bool, bool) {
	if o == nil || o.SupportsCreateDeleteSnapshot == nil {
		return nil, false
	}
	return o.SupportsCreateDeleteSnapshot, true
}

// HasSupportsCreateDeleteSnapshot returns a boolean if a field has been set.
func (o *CSIControllerInfo) HasSupportsCreateDeleteSnapshot() bool {
	if o != nil && o.SupportsCreateDeleteSnapshot != nil {
		return true
	}

	return false
}

// SetSupportsCreateDeleteSnapshot gets a reference to the given bool and assigns it to the SupportsCreateDeleteSnapshot field.
func (o *CSIControllerInfo) SetSupportsCreateDeleteSnapshot(v bool) {
	o.SupportsCreateDeleteSnapshot = &v
}

// GetSupportsExpand returns the SupportsExpand field value if set, zero value otherwise.
func (o *CSIControllerInfo) GetSupportsExpand() bool {
	if o == nil || o.SupportsExpand == nil {
		var ret bool
		return ret
	}
	return *o.SupportsExpand
}

// GetSupportsExpandOk returns a tuple with the SupportsExpand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSIControllerInfo) GetSupportsExpandOk() (*bool, bool) {
	if o == nil || o.SupportsExpand == nil {
		return nil, false
	}
	return o.SupportsExpand, true
}

// HasSupportsExpand returns a boolean if a field has been set.
func (o *CSIControllerInfo) HasSupportsExpand() bool {
	if o != nil && o.SupportsExpand != nil {
		return true
	}

	return false
}

// SetSupportsExpand gets a reference to the given bool and assigns it to the SupportsExpand field.
func (o *CSIControllerInfo) SetSupportsExpand(v bool) {
	o.SupportsExpand = &v
}

// GetSupportsGet returns the SupportsGet field value if set, zero value otherwise.
func (o *CSIControllerInfo) GetSupportsGet() bool {
	if o == nil || o.SupportsGet == nil {
		var ret bool
		return ret
	}
	return *o.SupportsGet
}

// GetSupportsGetOk returns a tuple with the SupportsGet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSIControllerInfo) GetSupportsGetOk() (*bool, bool) {
	if o == nil || o.SupportsGet == nil {
		return nil, false
	}
	return o.SupportsGet, true
}

// HasSupportsGet returns a boolean if a field has been set.
func (o *CSIControllerInfo) HasSupportsGet() bool {
	if o != nil && o.SupportsGet != nil {
		return true
	}

	return false
}

// SetSupportsGet gets a reference to the given bool and assigns it to the SupportsGet field.
func (o *CSIControllerInfo) SetSupportsGet(v bool) {
	o.SupportsGet = &v
}

// GetSupportsGetCapacity returns the SupportsGetCapacity field value if set, zero value otherwise.
func (o *CSIControllerInfo) GetSupportsGetCapacity() bool {
	if o == nil || o.SupportsGetCapacity == nil {
		var ret bool
		return ret
	}
	return *o.SupportsGetCapacity
}

// GetSupportsGetCapacityOk returns a tuple with the SupportsGetCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSIControllerInfo) GetSupportsGetCapacityOk() (*bool, bool) {
	if o == nil || o.SupportsGetCapacity == nil {
		return nil, false
	}
	return o.SupportsGetCapacity, true
}

// HasSupportsGetCapacity returns a boolean if a field has been set.
func (o *CSIControllerInfo) HasSupportsGetCapacity() bool {
	if o != nil && o.SupportsGetCapacity != nil {
		return true
	}

	return false
}

// SetSupportsGetCapacity gets a reference to the given bool and assigns it to the SupportsGetCapacity field.
func (o *CSIControllerInfo) SetSupportsGetCapacity(v bool) {
	o.SupportsGetCapacity = &v
}

// GetSupportsListSnapshots returns the SupportsListSnapshots field value if set, zero value otherwise.
func (o *CSIControllerInfo) GetSupportsListSnapshots() bool {
	if o == nil || o.SupportsListSnapshots == nil {
		var ret bool
		return ret
	}
	return *o.SupportsListSnapshots
}

// GetSupportsListSnapshotsOk returns a tuple with the SupportsListSnapshots field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSIControllerInfo) GetSupportsListSnapshotsOk() (*bool, bool) {
	if o == nil || o.SupportsListSnapshots == nil {
		return nil, false
	}
	return o.SupportsListSnapshots, true
}

// HasSupportsListSnapshots returns a boolean if a field has been set.
func (o *CSIControllerInfo) HasSupportsListSnapshots() bool {
	if o != nil && o.SupportsListSnapshots != nil {
		return true
	}

	return false
}

// SetSupportsListSnapshots gets a reference to the given bool and assigns it to the SupportsListSnapshots field.
func (o *CSIControllerInfo) SetSupportsListSnapshots(v bool) {
	o.SupportsListSnapshots = &v
}

// GetSupportsListVolumes returns the SupportsListVolumes field value if set, zero value otherwise.
func (o *CSIControllerInfo) GetSupportsListVolumes() bool {
	if o == nil || o.SupportsListVolumes == nil {
		var ret bool
		return ret
	}
	return *o.SupportsListVolumes
}

// GetSupportsListVolumesOk returns a tuple with the SupportsListVolumes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSIControllerInfo) GetSupportsListVolumesOk() (*bool, bool) {
	if o == nil || o.SupportsListVolumes == nil {
		return nil, false
	}
	return o.SupportsListVolumes, true
}

// HasSupportsListVolumes returns a boolean if a field has been set.
func (o *CSIControllerInfo) HasSupportsListVolumes() bool {
	if o != nil && o.SupportsListVolumes != nil {
		return true
	}

	return false
}

// SetSupportsListVolumes gets a reference to the given bool and assigns it to the SupportsListVolumes field.
func (o *CSIControllerInfo) SetSupportsListVolumes(v bool) {
	o.SupportsListVolumes = &v
}

// GetSupportsListVolumesAttachedNodes returns the SupportsListVolumesAttachedNodes field value if set, zero value otherwise.
func (o *CSIControllerInfo) GetSupportsListVolumesAttachedNodes() bool {
	if o == nil || o.SupportsListVolumesAttachedNodes == nil {
		var ret bool
		return ret
	}
	return *o.SupportsListVolumesAttachedNodes
}

// GetSupportsListVolumesAttachedNodesOk returns a tuple with the SupportsListVolumesAttachedNodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSIControllerInfo) GetSupportsListVolumesAttachedNodesOk() (*bool, bool) {
	if o == nil || o.SupportsListVolumesAttachedNodes == nil {
		return nil, false
	}
	return o.SupportsListVolumesAttachedNodes, true
}

// HasSupportsListVolumesAttachedNodes returns a boolean if a field has been set.
func (o *CSIControllerInfo) HasSupportsListVolumesAttachedNodes() bool {
	if o != nil && o.SupportsListVolumesAttachedNodes != nil {
		return true
	}

	return false
}

// SetSupportsListVolumesAttachedNodes gets a reference to the given bool and assigns it to the SupportsListVolumesAttachedNodes field.
func (o *CSIControllerInfo) SetSupportsListVolumesAttachedNodes(v bool) {
	o.SupportsListVolumesAttachedNodes = &v
}

// GetSupportsReadOnlyAttach returns the SupportsReadOnlyAttach field value if set, zero value otherwise.
func (o *CSIControllerInfo) GetSupportsReadOnlyAttach() bool {
	if o == nil || o.SupportsReadOnlyAttach == nil {
		var ret bool
		return ret
	}
	return *o.SupportsReadOnlyAttach
}

// GetSupportsReadOnlyAttachOk returns a tuple with the SupportsReadOnlyAttach field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSIControllerInfo) GetSupportsReadOnlyAttachOk() (*bool, bool) {
	if o == nil || o.SupportsReadOnlyAttach == nil {
		return nil, false
	}
	return o.SupportsReadOnlyAttach, true
}

// HasSupportsReadOnlyAttach returns a boolean if a field has been set.
func (o *CSIControllerInfo) HasSupportsReadOnlyAttach() bool {
	if o != nil && o.SupportsReadOnlyAttach != nil {
		return true
	}

	return false
}

// SetSupportsReadOnlyAttach gets a reference to the given bool and assigns it to the SupportsReadOnlyAttach field.
func (o *CSIControllerInfo) SetSupportsReadOnlyAttach(v bool) {
	o.SupportsReadOnlyAttach = &v
}

func (o CSIControllerInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SupportsAttachDetach != nil {
		toSerialize["SupportsAttachDetach"] = o.SupportsAttachDetach
	}
	if o.SupportsClone != nil {
		toSerialize["SupportsClone"] = o.SupportsClone
	}
	if o.SupportsCondition != nil {
		toSerialize["SupportsCondition"] = o.SupportsCondition
	}
	if o.SupportsCreateDelete != nil {
		toSerialize["SupportsCreateDelete"] = o.SupportsCreateDelete
	}
	if o.SupportsCreateDeleteSnapshot != nil {
		toSerialize["SupportsCreateDeleteSnapshot"] = o.SupportsCreateDeleteSnapshot
	}
	if o.SupportsExpand != nil {
		toSerialize["SupportsExpand"] = o.SupportsExpand
	}
	if o.SupportsGet != nil {
		toSerialize["SupportsGet"] = o.SupportsGet
	}
	if o.SupportsGetCapacity != nil {
		toSerialize["SupportsGetCapacity"] = o.SupportsGetCapacity
	}
	if o.SupportsListSnapshots != nil {
		toSerialize["SupportsListSnapshots"] = o.SupportsListSnapshots
	}
	if o.SupportsListVolumes != nil {
		toSerialize["SupportsListVolumes"] = o.SupportsListVolumes
	}
	if o.SupportsListVolumesAttachedNodes != nil {
		toSerialize["SupportsListVolumesAttachedNodes"] = o.SupportsListVolumesAttachedNodes
	}
	if o.SupportsReadOnlyAttach != nil {
		toSerialize["SupportsReadOnlyAttach"] = o.SupportsReadOnlyAttach
	}
	return json.Marshal(toSerialize)
}

type NullableCSIControllerInfo struct {
	value *CSIControllerInfo
	isSet bool
}

func (v NullableCSIControllerInfo) Get() *CSIControllerInfo {
	return v.value
}

func (v *NullableCSIControllerInfo) Set(val *CSIControllerInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableCSIControllerInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableCSIControllerInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCSIControllerInfo(val *CSIControllerInfo) *NullableCSIControllerInfo {
	return &NullableCSIControllerInfo{value: val, isSet: true}
}

func (v NullableCSIControllerInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCSIControllerInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


