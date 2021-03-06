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

// JobPlanResponse struct for JobPlanResponse
type JobPlanResponse struct {
	Annotations *PlanAnnotations `json:"Annotations,omitempty"`
	CreatedEvals *[]Evaluation `json:"CreatedEvals,omitempty"`
	Diff *JobDiff `json:"Diff,omitempty"`
	FailedTGAllocs *map[string]AllocationMetric `json:"FailedTGAllocs,omitempty"`
	JobModifyIndex *int32 `json:"JobModifyIndex,omitempty"`
	NextPeriodicLaunch *time.Time `json:"NextPeriodicLaunch,omitempty"`
	Warnings *string `json:"Warnings,omitempty"`
}

// NewJobPlanResponse instantiates a new JobPlanResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobPlanResponse() *JobPlanResponse {
	this := JobPlanResponse{}
	return &this
}

// NewJobPlanResponseWithDefaults instantiates a new JobPlanResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobPlanResponseWithDefaults() *JobPlanResponse {
	this := JobPlanResponse{}
	return &this
}

// GetAnnotations returns the Annotations field value if set, zero value otherwise.
func (o *JobPlanResponse) GetAnnotations() PlanAnnotations {
	if o == nil || o.Annotations == nil {
		var ret PlanAnnotations
		return ret
	}
	return *o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobPlanResponse) GetAnnotationsOk() (*PlanAnnotations, bool) {
	if o == nil || o.Annotations == nil {
		return nil, false
	}
	return o.Annotations, true
}

// HasAnnotations returns a boolean if a field has been set.
func (o *JobPlanResponse) HasAnnotations() bool {
	if o != nil && o.Annotations != nil {
		return true
	}

	return false
}

// SetAnnotations gets a reference to the given PlanAnnotations and assigns it to the Annotations field.
func (o *JobPlanResponse) SetAnnotations(v PlanAnnotations) {
	o.Annotations = &v
}

// GetCreatedEvals returns the CreatedEvals field value if set, zero value otherwise.
func (o *JobPlanResponse) GetCreatedEvals() []Evaluation {
	if o == nil || o.CreatedEvals == nil {
		var ret []Evaluation
		return ret
	}
	return *o.CreatedEvals
}

// GetCreatedEvalsOk returns a tuple with the CreatedEvals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobPlanResponse) GetCreatedEvalsOk() (*[]Evaluation, bool) {
	if o == nil || o.CreatedEvals == nil {
		return nil, false
	}
	return o.CreatedEvals, true
}

// HasCreatedEvals returns a boolean if a field has been set.
func (o *JobPlanResponse) HasCreatedEvals() bool {
	if o != nil && o.CreatedEvals != nil {
		return true
	}

	return false
}

// SetCreatedEvals gets a reference to the given []Evaluation and assigns it to the CreatedEvals field.
func (o *JobPlanResponse) SetCreatedEvals(v []Evaluation) {
	o.CreatedEvals = &v
}

// GetDiff returns the Diff field value if set, zero value otherwise.
func (o *JobPlanResponse) GetDiff() JobDiff {
	if o == nil || o.Diff == nil {
		var ret JobDiff
		return ret
	}
	return *o.Diff
}

// GetDiffOk returns a tuple with the Diff field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobPlanResponse) GetDiffOk() (*JobDiff, bool) {
	if o == nil || o.Diff == nil {
		return nil, false
	}
	return o.Diff, true
}

// HasDiff returns a boolean if a field has been set.
func (o *JobPlanResponse) HasDiff() bool {
	if o != nil && o.Diff != nil {
		return true
	}

	return false
}

// SetDiff gets a reference to the given JobDiff and assigns it to the Diff field.
func (o *JobPlanResponse) SetDiff(v JobDiff) {
	o.Diff = &v
}

// GetFailedTGAllocs returns the FailedTGAllocs field value if set, zero value otherwise.
func (o *JobPlanResponse) GetFailedTGAllocs() map[string]AllocationMetric {
	if o == nil || o.FailedTGAllocs == nil {
		var ret map[string]AllocationMetric
		return ret
	}
	return *o.FailedTGAllocs
}

// GetFailedTGAllocsOk returns a tuple with the FailedTGAllocs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobPlanResponse) GetFailedTGAllocsOk() (*map[string]AllocationMetric, bool) {
	if o == nil || o.FailedTGAllocs == nil {
		return nil, false
	}
	return o.FailedTGAllocs, true
}

// HasFailedTGAllocs returns a boolean if a field has been set.
func (o *JobPlanResponse) HasFailedTGAllocs() bool {
	if o != nil && o.FailedTGAllocs != nil {
		return true
	}

	return false
}

// SetFailedTGAllocs gets a reference to the given map[string]AllocationMetric and assigns it to the FailedTGAllocs field.
func (o *JobPlanResponse) SetFailedTGAllocs(v map[string]AllocationMetric) {
	o.FailedTGAllocs = &v
}

// GetJobModifyIndex returns the JobModifyIndex field value if set, zero value otherwise.
func (o *JobPlanResponse) GetJobModifyIndex() int32 {
	if o == nil || o.JobModifyIndex == nil {
		var ret int32
		return ret
	}
	return *o.JobModifyIndex
}

// GetJobModifyIndexOk returns a tuple with the JobModifyIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobPlanResponse) GetJobModifyIndexOk() (*int32, bool) {
	if o == nil || o.JobModifyIndex == nil {
		return nil, false
	}
	return o.JobModifyIndex, true
}

// HasJobModifyIndex returns a boolean if a field has been set.
func (o *JobPlanResponse) HasJobModifyIndex() bool {
	if o != nil && o.JobModifyIndex != nil {
		return true
	}

	return false
}

// SetJobModifyIndex gets a reference to the given int32 and assigns it to the JobModifyIndex field.
func (o *JobPlanResponse) SetJobModifyIndex(v int32) {
	o.JobModifyIndex = &v
}

// GetNextPeriodicLaunch returns the NextPeriodicLaunch field value if set, zero value otherwise.
func (o *JobPlanResponse) GetNextPeriodicLaunch() time.Time {
	if o == nil || o.NextPeriodicLaunch == nil {
		var ret time.Time
		return ret
	}
	return *o.NextPeriodicLaunch
}

// GetNextPeriodicLaunchOk returns a tuple with the NextPeriodicLaunch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobPlanResponse) GetNextPeriodicLaunchOk() (*time.Time, bool) {
	if o == nil || o.NextPeriodicLaunch == nil {
		return nil, false
	}
	return o.NextPeriodicLaunch, true
}

// HasNextPeriodicLaunch returns a boolean if a field has been set.
func (o *JobPlanResponse) HasNextPeriodicLaunch() bool {
	if o != nil && o.NextPeriodicLaunch != nil {
		return true
	}

	return false
}

// SetNextPeriodicLaunch gets a reference to the given time.Time and assigns it to the NextPeriodicLaunch field.
func (o *JobPlanResponse) SetNextPeriodicLaunch(v time.Time) {
	o.NextPeriodicLaunch = &v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *JobPlanResponse) GetWarnings() string {
	if o == nil || o.Warnings == nil {
		var ret string
		return ret
	}
	return *o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobPlanResponse) GetWarningsOk() (*string, bool) {
	if o == nil || o.Warnings == nil {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *JobPlanResponse) HasWarnings() bool {
	if o != nil && o.Warnings != nil {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given string and assigns it to the Warnings field.
func (o *JobPlanResponse) SetWarnings(v string) {
	o.Warnings = &v
}

func (o JobPlanResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Annotations != nil {
		toSerialize["Annotations"] = o.Annotations
	}
	if o.CreatedEvals != nil {
		toSerialize["CreatedEvals"] = o.CreatedEvals
	}
	if o.Diff != nil {
		toSerialize["Diff"] = o.Diff
	}
	if o.FailedTGAllocs != nil {
		toSerialize["FailedTGAllocs"] = o.FailedTGAllocs
	}
	if o.JobModifyIndex != nil {
		toSerialize["JobModifyIndex"] = o.JobModifyIndex
	}
	if o.NextPeriodicLaunch != nil {
		toSerialize["NextPeriodicLaunch"] = o.NextPeriodicLaunch
	}
	if o.Warnings != nil {
		toSerialize["Warnings"] = o.Warnings
	}
	return json.Marshal(toSerialize)
}

type NullableJobPlanResponse struct {
	value *JobPlanResponse
	isSet bool
}

func (v NullableJobPlanResponse) Get() *JobPlanResponse {
	return v.value
}

func (v *NullableJobPlanResponse) Set(val *JobPlanResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableJobPlanResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableJobPlanResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobPlanResponse(val *JobPlanResponse) *NullableJobPlanResponse {
	return &NullableJobPlanResponse{value: val, isSet: true}
}

func (v NullableJobPlanResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobPlanResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


