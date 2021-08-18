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

// CSIVolumeCreateRequest struct for CSIVolumeCreateRequest
type CSIVolumeCreateRequest struct {
	Namespace *string `json:"Namespace,omitempty"`
	Region *string `json:"Region,omitempty"`
	SecretID *string `json:"SecretID,omitempty"`
	Volumes *[]CSIVolume `json:"Volumes,omitempty"`
}

// NewCSIVolumeCreateRequest instantiates a new CSIVolumeCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCSIVolumeCreateRequest() *CSIVolumeCreateRequest {
	this := CSIVolumeCreateRequest{}
	return &this
}

// NewCSIVolumeCreateRequestWithDefaults instantiates a new CSIVolumeCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCSIVolumeCreateRequestWithDefaults() *CSIVolumeCreateRequest {
	this := CSIVolumeCreateRequest{}
	return &this
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *CSIVolumeCreateRequest) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSIVolumeCreateRequest) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *CSIVolumeCreateRequest) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *CSIVolumeCreateRequest) SetNamespace(v string) {
	o.Namespace = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *CSIVolumeCreateRequest) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSIVolumeCreateRequest) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *CSIVolumeCreateRequest) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *CSIVolumeCreateRequest) SetRegion(v string) {
	o.Region = &v
}

// GetSecretID returns the SecretID field value if set, zero value otherwise.
func (o *CSIVolumeCreateRequest) GetSecretID() string {
	if o == nil || o.SecretID == nil {
		var ret string
		return ret
	}
	return *o.SecretID
}

// GetSecretIDOk returns a tuple with the SecretID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSIVolumeCreateRequest) GetSecretIDOk() (*string, bool) {
	if o == nil || o.SecretID == nil {
		return nil, false
	}
	return o.SecretID, true
}

// HasSecretID returns a boolean if a field has been set.
func (o *CSIVolumeCreateRequest) HasSecretID() bool {
	if o != nil && o.SecretID != nil {
		return true
	}

	return false
}

// SetSecretID gets a reference to the given string and assigns it to the SecretID field.
func (o *CSIVolumeCreateRequest) SetSecretID(v string) {
	o.SecretID = &v
}

// GetVolumes returns the Volumes field value if set, zero value otherwise.
func (o *CSIVolumeCreateRequest) GetVolumes() []CSIVolume {
	if o == nil || o.Volumes == nil {
		var ret []CSIVolume
		return ret
	}
	return *o.Volumes
}

// GetVolumesOk returns a tuple with the Volumes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSIVolumeCreateRequest) GetVolumesOk() (*[]CSIVolume, bool) {
	if o == nil || o.Volumes == nil {
		return nil, false
	}
	return o.Volumes, true
}

// HasVolumes returns a boolean if a field has been set.
func (o *CSIVolumeCreateRequest) HasVolumes() bool {
	if o != nil && o.Volumes != nil {
		return true
	}

	return false
}

// SetVolumes gets a reference to the given []CSIVolume and assigns it to the Volumes field.
func (o *CSIVolumeCreateRequest) SetVolumes(v []CSIVolume) {
	o.Volumes = &v
}

func (o CSIVolumeCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Namespace != nil {
		toSerialize["Namespace"] = o.Namespace
	}
	if o.Region != nil {
		toSerialize["Region"] = o.Region
	}
	if o.SecretID != nil {
		toSerialize["SecretID"] = o.SecretID
	}
	if o.Volumes != nil {
		toSerialize["Volumes"] = o.Volumes
	}
	return json.Marshal(toSerialize)
}

type NullableCSIVolumeCreateRequest struct {
	value *CSIVolumeCreateRequest
	isSet bool
}

func (v NullableCSIVolumeCreateRequest) Get() *CSIVolumeCreateRequest {
	return v.value
}

func (v *NullableCSIVolumeCreateRequest) Set(val *CSIVolumeCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCSIVolumeCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCSIVolumeCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCSIVolumeCreateRequest(val *CSIVolumeCreateRequest) *NullableCSIVolumeCreateRequest {
	return &NullableCSIVolumeCreateRequest{value: val, isSet: true}
}

func (v NullableCSIVolumeCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCSIVolumeCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

