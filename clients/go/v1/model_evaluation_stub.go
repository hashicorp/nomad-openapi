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
	"time"
)

// EvaluationStub struct for EvaluationStub
type EvaluationStub struct {
	BlockedEval *string `json:"BlockedEval,omitempty"`
	CreateIndex *int32 `json:"CreateIndex,omitempty"`
	CreateTime *int64 `json:"CreateTime,omitempty"`
	DeploymentID *string `json:"DeploymentID,omitempty"`
	ID *string `json:"ID,omitempty"`
	JobID *string `json:"JobID,omitempty"`
	ModifyIndex *int32 `json:"ModifyIndex,omitempty"`
	ModifyTime *int64 `json:"ModifyTime,omitempty"`
	Namespace *string `json:"Namespace,omitempty"`
	NextEval *string `json:"NextEval,omitempty"`
	NodeID *string `json:"NodeID,omitempty"`
	PreviousEval *string `json:"PreviousEval,omitempty"`
	Priority *int32 `json:"Priority,omitempty"`
	Status *string `json:"Status,omitempty"`
	StatusDescription *string `json:"StatusDescription,omitempty"`
	TriggeredBy *string `json:"TriggeredBy,omitempty"`
	Type *string `json:"Type,omitempty"`
	WaitUntil NullableTime `json:"WaitUntil,omitempty"`
}

// NewEvaluationStub instantiates a new EvaluationStub object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEvaluationStub() *EvaluationStub {
	this := EvaluationStub{}
	return &this
}

// NewEvaluationStubWithDefaults instantiates a new EvaluationStub object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEvaluationStubWithDefaults() *EvaluationStub {
	this := EvaluationStub{}
	return &this
}

// GetBlockedEval returns the BlockedEval field value if set, zero value otherwise.
func (o *EvaluationStub) GetBlockedEval() string {
	if o == nil || o.BlockedEval == nil {
		var ret string
		return ret
	}
	return *o.BlockedEval
}

// GetBlockedEvalOk returns a tuple with the BlockedEval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EvaluationStub) GetBlockedEvalOk() (*string, bool) {
	if o == nil || o.BlockedEval == nil {
		return nil, false
	}
	return o.BlockedEval, true
}

// HasBlockedEval returns a boolean if a field has been set.
func (o *EvaluationStub) HasBlockedEval() bool {
	if o != nil && o.BlockedEval != nil {
		return true
	}

	return false
}

// SetBlockedEval gets a reference to the given string and assigns it to the BlockedEval field.
func (o *EvaluationStub) SetBlockedEval(v string) {
	o.BlockedEval = &v
}

// GetCreateIndex returns the CreateIndex field value if set, zero value otherwise.
func (o *EvaluationStub) GetCreateIndex() int32 {
	if o == nil || o.CreateIndex == nil {
		var ret int32
		return ret
	}
	return *o.CreateIndex
}

// GetCreateIndexOk returns a tuple with the CreateIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EvaluationStub) GetCreateIndexOk() (*int32, bool) {
	if o == nil || o.CreateIndex == nil {
		return nil, false
	}
	return o.CreateIndex, true
}

// HasCreateIndex returns a boolean if a field has been set.
func (o *EvaluationStub) HasCreateIndex() bool {
	if o != nil && o.CreateIndex != nil {
		return true
	}

	return false
}

// SetCreateIndex gets a reference to the given int32 and assigns it to the CreateIndex field.
func (o *EvaluationStub) SetCreateIndex(v int32) {
	o.CreateIndex = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *EvaluationStub) GetCreateTime() int64 {
	if o == nil || o.CreateTime == nil {
		var ret int64
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EvaluationStub) GetCreateTimeOk() (*int64, bool) {
	if o == nil || o.CreateTime == nil {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *EvaluationStub) HasCreateTime() bool {
	if o != nil && o.CreateTime != nil {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given int64 and assigns it to the CreateTime field.
func (o *EvaluationStub) SetCreateTime(v int64) {
	o.CreateTime = &v
}

// GetDeploymentID returns the DeploymentID field value if set, zero value otherwise.
func (o *EvaluationStub) GetDeploymentID() string {
	if o == nil || o.DeploymentID == nil {
		var ret string
		return ret
	}
	return *o.DeploymentID
}

// GetDeploymentIDOk returns a tuple with the DeploymentID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EvaluationStub) GetDeploymentIDOk() (*string, bool) {
	if o == nil || o.DeploymentID == nil {
		return nil, false
	}
	return o.DeploymentID, true
}

// HasDeploymentID returns a boolean if a field has been set.
func (o *EvaluationStub) HasDeploymentID() bool {
	if o != nil && o.DeploymentID != nil {
		return true
	}

	return false
}

// SetDeploymentID gets a reference to the given string and assigns it to the DeploymentID field.
func (o *EvaluationStub) SetDeploymentID(v string) {
	o.DeploymentID = &v
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *EvaluationStub) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EvaluationStub) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *EvaluationStub) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *EvaluationStub) SetID(v string) {
	o.ID = &v
}

// GetJobID returns the JobID field value if set, zero value otherwise.
func (o *EvaluationStub) GetJobID() string {
	if o == nil || o.JobID == nil {
		var ret string
		return ret
	}
	return *o.JobID
}

// GetJobIDOk returns a tuple with the JobID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EvaluationStub) GetJobIDOk() (*string, bool) {
	if o == nil || o.JobID == nil {
		return nil, false
	}
	return o.JobID, true
}

// HasJobID returns a boolean if a field has been set.
func (o *EvaluationStub) HasJobID() bool {
	if o != nil && o.JobID != nil {
		return true
	}

	return false
}

// SetJobID gets a reference to the given string and assigns it to the JobID field.
func (o *EvaluationStub) SetJobID(v string) {
	o.JobID = &v
}

// GetModifyIndex returns the ModifyIndex field value if set, zero value otherwise.
func (o *EvaluationStub) GetModifyIndex() int32 {
	if o == nil || o.ModifyIndex == nil {
		var ret int32
		return ret
	}
	return *o.ModifyIndex
}

// GetModifyIndexOk returns a tuple with the ModifyIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EvaluationStub) GetModifyIndexOk() (*int32, bool) {
	if o == nil || o.ModifyIndex == nil {
		return nil, false
	}
	return o.ModifyIndex, true
}

// HasModifyIndex returns a boolean if a field has been set.
func (o *EvaluationStub) HasModifyIndex() bool {
	if o != nil && o.ModifyIndex != nil {
		return true
	}

	return false
}

// SetModifyIndex gets a reference to the given int32 and assigns it to the ModifyIndex field.
func (o *EvaluationStub) SetModifyIndex(v int32) {
	o.ModifyIndex = &v
}

// GetModifyTime returns the ModifyTime field value if set, zero value otherwise.
func (o *EvaluationStub) GetModifyTime() int64 {
	if o == nil || o.ModifyTime == nil {
		var ret int64
		return ret
	}
	return *o.ModifyTime
}

// GetModifyTimeOk returns a tuple with the ModifyTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EvaluationStub) GetModifyTimeOk() (*int64, bool) {
	if o == nil || o.ModifyTime == nil {
		return nil, false
	}
	return o.ModifyTime, true
}

// HasModifyTime returns a boolean if a field has been set.
func (o *EvaluationStub) HasModifyTime() bool {
	if o != nil && o.ModifyTime != nil {
		return true
	}

	return false
}

// SetModifyTime gets a reference to the given int64 and assigns it to the ModifyTime field.
func (o *EvaluationStub) SetModifyTime(v int64) {
	o.ModifyTime = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *EvaluationStub) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EvaluationStub) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *EvaluationStub) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *EvaluationStub) SetNamespace(v string) {
	o.Namespace = &v
}

// GetNextEval returns the NextEval field value if set, zero value otherwise.
func (o *EvaluationStub) GetNextEval() string {
	if o == nil || o.NextEval == nil {
		var ret string
		return ret
	}
	return *o.NextEval
}

// GetNextEvalOk returns a tuple with the NextEval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EvaluationStub) GetNextEvalOk() (*string, bool) {
	if o == nil || o.NextEval == nil {
		return nil, false
	}
	return o.NextEval, true
}

// HasNextEval returns a boolean if a field has been set.
func (o *EvaluationStub) HasNextEval() bool {
	if o != nil && o.NextEval != nil {
		return true
	}

	return false
}

// SetNextEval gets a reference to the given string and assigns it to the NextEval field.
func (o *EvaluationStub) SetNextEval(v string) {
	o.NextEval = &v
}

// GetNodeID returns the NodeID field value if set, zero value otherwise.
func (o *EvaluationStub) GetNodeID() string {
	if o == nil || o.NodeID == nil {
		var ret string
		return ret
	}
	return *o.NodeID
}

// GetNodeIDOk returns a tuple with the NodeID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EvaluationStub) GetNodeIDOk() (*string, bool) {
	if o == nil || o.NodeID == nil {
		return nil, false
	}
	return o.NodeID, true
}

// HasNodeID returns a boolean if a field has been set.
func (o *EvaluationStub) HasNodeID() bool {
	if o != nil && o.NodeID != nil {
		return true
	}

	return false
}

// SetNodeID gets a reference to the given string and assigns it to the NodeID field.
func (o *EvaluationStub) SetNodeID(v string) {
	o.NodeID = &v
}

// GetPreviousEval returns the PreviousEval field value if set, zero value otherwise.
func (o *EvaluationStub) GetPreviousEval() string {
	if o == nil || o.PreviousEval == nil {
		var ret string
		return ret
	}
	return *o.PreviousEval
}

// GetPreviousEvalOk returns a tuple with the PreviousEval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EvaluationStub) GetPreviousEvalOk() (*string, bool) {
	if o == nil || o.PreviousEval == nil {
		return nil, false
	}
	return o.PreviousEval, true
}

// HasPreviousEval returns a boolean if a field has been set.
func (o *EvaluationStub) HasPreviousEval() bool {
	if o != nil && o.PreviousEval != nil {
		return true
	}

	return false
}

// SetPreviousEval gets a reference to the given string and assigns it to the PreviousEval field.
func (o *EvaluationStub) SetPreviousEval(v string) {
	o.PreviousEval = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *EvaluationStub) GetPriority() int32 {
	if o == nil || o.Priority == nil {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EvaluationStub) GetPriorityOk() (*int32, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *EvaluationStub) HasPriority() bool {
	if o != nil && o.Priority != nil {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *EvaluationStub) SetPriority(v int32) {
	o.Priority = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *EvaluationStub) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EvaluationStub) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *EvaluationStub) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *EvaluationStub) SetStatus(v string) {
	o.Status = &v
}

// GetStatusDescription returns the StatusDescription field value if set, zero value otherwise.
func (o *EvaluationStub) GetStatusDescription() string {
	if o == nil || o.StatusDescription == nil {
		var ret string
		return ret
	}
	return *o.StatusDescription
}

// GetStatusDescriptionOk returns a tuple with the StatusDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EvaluationStub) GetStatusDescriptionOk() (*string, bool) {
	if o == nil || o.StatusDescription == nil {
		return nil, false
	}
	return o.StatusDescription, true
}

// HasStatusDescription returns a boolean if a field has been set.
func (o *EvaluationStub) HasStatusDescription() bool {
	if o != nil && o.StatusDescription != nil {
		return true
	}

	return false
}

// SetStatusDescription gets a reference to the given string and assigns it to the StatusDescription field.
func (o *EvaluationStub) SetStatusDescription(v string) {
	o.StatusDescription = &v
}

// GetTriggeredBy returns the TriggeredBy field value if set, zero value otherwise.
func (o *EvaluationStub) GetTriggeredBy() string {
	if o == nil || o.TriggeredBy == nil {
		var ret string
		return ret
	}
	return *o.TriggeredBy
}

// GetTriggeredByOk returns a tuple with the TriggeredBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EvaluationStub) GetTriggeredByOk() (*string, bool) {
	if o == nil || o.TriggeredBy == nil {
		return nil, false
	}
	return o.TriggeredBy, true
}

// HasTriggeredBy returns a boolean if a field has been set.
func (o *EvaluationStub) HasTriggeredBy() bool {
	if o != nil && o.TriggeredBy != nil {
		return true
	}

	return false
}

// SetTriggeredBy gets a reference to the given string and assigns it to the TriggeredBy field.
func (o *EvaluationStub) SetTriggeredBy(v string) {
	o.TriggeredBy = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *EvaluationStub) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EvaluationStub) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *EvaluationStub) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *EvaluationStub) SetType(v string) {
	o.Type = &v
}

// GetWaitUntil returns the WaitUntil field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EvaluationStub) GetWaitUntil() time.Time {
	if o == nil || o.WaitUntil.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.WaitUntil.Get()
}

// GetWaitUntilOk returns a tuple with the WaitUntil field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EvaluationStub) GetWaitUntilOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.WaitUntil.Get(), o.WaitUntil.IsSet()
}

// HasWaitUntil returns a boolean if a field has been set.
func (o *EvaluationStub) HasWaitUntil() bool {
	if o != nil && o.WaitUntil.IsSet() {
		return true
	}

	return false
}

// SetWaitUntil gets a reference to the given NullableTime and assigns it to the WaitUntil field.
func (o *EvaluationStub) SetWaitUntil(v time.Time) {
	o.WaitUntil.Set(&v)
}
// SetWaitUntilNil sets the value for WaitUntil to be an explicit nil
func (o *EvaluationStub) SetWaitUntilNil() {
	o.WaitUntil.Set(nil)
}

// UnsetWaitUntil ensures that no value is present for WaitUntil, not even an explicit nil
func (o *EvaluationStub) UnsetWaitUntil() {
	o.WaitUntil.Unset()
}

func (o EvaluationStub) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BlockedEval != nil {
		toSerialize["BlockedEval"] = o.BlockedEval
	}
	if o.CreateIndex != nil {
		toSerialize["CreateIndex"] = o.CreateIndex
	}
	if o.CreateTime != nil {
		toSerialize["CreateTime"] = o.CreateTime
	}
	if o.DeploymentID != nil {
		toSerialize["DeploymentID"] = o.DeploymentID
	}
	if o.ID != nil {
		toSerialize["ID"] = o.ID
	}
	if o.JobID != nil {
		toSerialize["JobID"] = o.JobID
	}
	if o.ModifyIndex != nil {
		toSerialize["ModifyIndex"] = o.ModifyIndex
	}
	if o.ModifyTime != nil {
		toSerialize["ModifyTime"] = o.ModifyTime
	}
	if o.Namespace != nil {
		toSerialize["Namespace"] = o.Namespace
	}
	if o.NextEval != nil {
		toSerialize["NextEval"] = o.NextEval
	}
	if o.NodeID != nil {
		toSerialize["NodeID"] = o.NodeID
	}
	if o.PreviousEval != nil {
		toSerialize["PreviousEval"] = o.PreviousEval
	}
	if o.Priority != nil {
		toSerialize["Priority"] = o.Priority
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.StatusDescription != nil {
		toSerialize["StatusDescription"] = o.StatusDescription
	}
	if o.TriggeredBy != nil {
		toSerialize["TriggeredBy"] = o.TriggeredBy
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	if o.WaitUntil.IsSet() {
		toSerialize["WaitUntil"] = o.WaitUntil.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableEvaluationStub struct {
	value *EvaluationStub
	isSet bool
}

func (v NullableEvaluationStub) Get() *EvaluationStub {
	return v.value
}

func (v *NullableEvaluationStub) Set(val *EvaluationStub) {
	v.value = val
	v.isSet = true
}

func (v NullableEvaluationStub) IsSet() bool {
	return v.isSet
}

func (v *NullableEvaluationStub) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEvaluationStub(val *EvaluationStub) *NullableEvaluationStub {
	return &NullableEvaluationStub{value: val, isSet: true}
}

func (v NullableEvaluationStub) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEvaluationStub) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


