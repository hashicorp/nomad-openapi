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

// DeploymentUnblockRequest struct for DeploymentUnblockRequest
type DeploymentUnblockRequest struct {
	DeploymentID *string `json:"DeploymentID,omitempty"`
	Namespace *string `json:"Namespace,omitempty"`
	Region *string `json:"Region,omitempty"`
	SecretID *string `json:"SecretID,omitempty"`
}

// NewDeploymentUnblockRequest instantiates a new DeploymentUnblockRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeploymentUnblockRequest() *DeploymentUnblockRequest {
	this := DeploymentUnblockRequest{}
	return &this
}

// NewDeploymentUnblockRequestWithDefaults instantiates a new DeploymentUnblockRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentUnblockRequestWithDefaults() *DeploymentUnblockRequest {
	this := DeploymentUnblockRequest{}
	return &this
}

// GetDeploymentID returns the DeploymentID field value if set, zero value otherwise.
func (o *DeploymentUnblockRequest) GetDeploymentID() string {
	if o == nil || o.DeploymentID == nil {
		var ret string
		return ret
	}
	return *o.DeploymentID
}

// GetDeploymentIDOk returns a tuple with the DeploymentID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentUnblockRequest) GetDeploymentIDOk() (*string, bool) {
	if o == nil || o.DeploymentID == nil {
		return nil, false
	}
	return o.DeploymentID, true
}

// HasDeploymentID returns a boolean if a field has been set.
func (o *DeploymentUnblockRequest) HasDeploymentID() bool {
	if o != nil && o.DeploymentID != nil {
		return true
	}

	return false
}

// SetDeploymentID gets a reference to the given string and assigns it to the DeploymentID field.
func (o *DeploymentUnblockRequest) SetDeploymentID(v string) {
	o.DeploymentID = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *DeploymentUnblockRequest) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentUnblockRequest) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *DeploymentUnblockRequest) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *DeploymentUnblockRequest) SetNamespace(v string) {
	o.Namespace = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *DeploymentUnblockRequest) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentUnblockRequest) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *DeploymentUnblockRequest) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *DeploymentUnblockRequest) SetRegion(v string) {
	o.Region = &v
}

// GetSecretID returns the SecretID field value if set, zero value otherwise.
func (o *DeploymentUnblockRequest) GetSecretID() string {
	if o == nil || o.SecretID == nil {
		var ret string
		return ret
	}
	return *o.SecretID
}

// GetSecretIDOk returns a tuple with the SecretID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentUnblockRequest) GetSecretIDOk() (*string, bool) {
	if o == nil || o.SecretID == nil {
		return nil, false
	}
	return o.SecretID, true
}

// HasSecretID returns a boolean if a field has been set.
func (o *DeploymentUnblockRequest) HasSecretID() bool {
	if o != nil && o.SecretID != nil {
		return true
	}

	return false
}

// SetSecretID gets a reference to the given string and assigns it to the SecretID field.
func (o *DeploymentUnblockRequest) SetSecretID(v string) {
	o.SecretID = &v
}

func (o DeploymentUnblockRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeploymentID != nil {
		toSerialize["DeploymentID"] = o.DeploymentID
	}
	if o.Namespace != nil {
		toSerialize["Namespace"] = o.Namespace
	}
	if o.Region != nil {
		toSerialize["Region"] = o.Region
	}
	if o.SecretID != nil {
		toSerialize["SecretID"] = o.SecretID
	}
	return json.Marshal(toSerialize)
}

type NullableDeploymentUnblockRequest struct {
	value *DeploymentUnblockRequest
	isSet bool
}

func (v NullableDeploymentUnblockRequest) Get() *DeploymentUnblockRequest {
	return v.value
}

func (v *NullableDeploymentUnblockRequest) Set(val *DeploymentUnblockRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeploymentUnblockRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploymentUnblockRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploymentUnblockRequest(val *DeploymentUnblockRequest) *NullableDeploymentUnblockRequest {
	return &NullableDeploymentUnblockRequest{value: val, isSet: true}
}

func (v NullableDeploymentUnblockRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploymentUnblockRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


