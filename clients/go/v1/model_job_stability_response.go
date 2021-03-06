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

// JobStabilityResponse struct for JobStabilityResponse
type JobStabilityResponse struct {
	Index *int32 `json:"Index,omitempty"`
}

// NewJobStabilityResponse instantiates a new JobStabilityResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobStabilityResponse() *JobStabilityResponse {
	this := JobStabilityResponse{}
	return &this
}

// NewJobStabilityResponseWithDefaults instantiates a new JobStabilityResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobStabilityResponseWithDefaults() *JobStabilityResponse {
	this := JobStabilityResponse{}
	return &this
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *JobStabilityResponse) GetIndex() int32 {
	if o == nil || o.Index == nil {
		var ret int32
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobStabilityResponse) GetIndexOk() (*int32, bool) {
	if o == nil || o.Index == nil {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *JobStabilityResponse) HasIndex() bool {
	if o != nil && o.Index != nil {
		return true
	}

	return false
}

// SetIndex gets a reference to the given int32 and assigns it to the Index field.
func (o *JobStabilityResponse) SetIndex(v int32) {
	o.Index = &v
}

func (o JobStabilityResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Index != nil {
		toSerialize["Index"] = o.Index
	}
	return json.Marshal(toSerialize)
}

type NullableJobStabilityResponse struct {
	value *JobStabilityResponse
	isSet bool
}

func (v NullableJobStabilityResponse) Get() *JobStabilityResponse {
	return v.value
}

func (v *NullableJobStabilityResponse) Set(val *JobStabilityResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableJobStabilityResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableJobStabilityResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobStabilityResponse(val *JobStabilityResponse) *NullableJobStabilityResponse {
	return &NullableJobStabilityResponse{value: val, isSet: true}
}

func (v NullableJobStabilityResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobStabilityResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


