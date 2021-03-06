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

// ConsulExposePath struct for ConsulExposePath
type ConsulExposePath struct {
	ListenerPort *string `json:"ListenerPort,omitempty"`
	LocalPathPort *int32 `json:"LocalPathPort,omitempty"`
	Path *string `json:"Path,omitempty"`
	Protocol *string `json:"Protocol,omitempty"`
}

// NewConsulExposePath instantiates a new ConsulExposePath object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsulExposePath() *ConsulExposePath {
	this := ConsulExposePath{}
	return &this
}

// NewConsulExposePathWithDefaults instantiates a new ConsulExposePath object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsulExposePathWithDefaults() *ConsulExposePath {
	this := ConsulExposePath{}
	return &this
}

// GetListenerPort returns the ListenerPort field value if set, zero value otherwise.
func (o *ConsulExposePath) GetListenerPort() string {
	if o == nil || o.ListenerPort == nil {
		var ret string
		return ret
	}
	return *o.ListenerPort
}

// GetListenerPortOk returns a tuple with the ListenerPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsulExposePath) GetListenerPortOk() (*string, bool) {
	if o == nil || o.ListenerPort == nil {
		return nil, false
	}
	return o.ListenerPort, true
}

// HasListenerPort returns a boolean if a field has been set.
func (o *ConsulExposePath) HasListenerPort() bool {
	if o != nil && o.ListenerPort != nil {
		return true
	}

	return false
}

// SetListenerPort gets a reference to the given string and assigns it to the ListenerPort field.
func (o *ConsulExposePath) SetListenerPort(v string) {
	o.ListenerPort = &v
}

// GetLocalPathPort returns the LocalPathPort field value if set, zero value otherwise.
func (o *ConsulExposePath) GetLocalPathPort() int32 {
	if o == nil || o.LocalPathPort == nil {
		var ret int32
		return ret
	}
	return *o.LocalPathPort
}

// GetLocalPathPortOk returns a tuple with the LocalPathPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsulExposePath) GetLocalPathPortOk() (*int32, bool) {
	if o == nil || o.LocalPathPort == nil {
		return nil, false
	}
	return o.LocalPathPort, true
}

// HasLocalPathPort returns a boolean if a field has been set.
func (o *ConsulExposePath) HasLocalPathPort() bool {
	if o != nil && o.LocalPathPort != nil {
		return true
	}

	return false
}

// SetLocalPathPort gets a reference to the given int32 and assigns it to the LocalPathPort field.
func (o *ConsulExposePath) SetLocalPathPort(v int32) {
	o.LocalPathPort = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *ConsulExposePath) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsulExposePath) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *ConsulExposePath) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *ConsulExposePath) SetPath(v string) {
	o.Path = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *ConsulExposePath) GetProtocol() string {
	if o == nil || o.Protocol == nil {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsulExposePath) GetProtocolOk() (*string, bool) {
	if o == nil || o.Protocol == nil {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *ConsulExposePath) HasProtocol() bool {
	if o != nil && o.Protocol != nil {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *ConsulExposePath) SetProtocol(v string) {
	o.Protocol = &v
}

func (o ConsulExposePath) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ListenerPort != nil {
		toSerialize["ListenerPort"] = o.ListenerPort
	}
	if o.LocalPathPort != nil {
		toSerialize["LocalPathPort"] = o.LocalPathPort
	}
	if o.Path != nil {
		toSerialize["Path"] = o.Path
	}
	if o.Protocol != nil {
		toSerialize["Protocol"] = o.Protocol
	}
	return json.Marshal(toSerialize)
}

type NullableConsulExposePath struct {
	value *ConsulExposePath
	isSet bool
}

func (v NullableConsulExposePath) Get() *ConsulExposePath {
	return v.value
}

func (v *NullableConsulExposePath) Set(val *ConsulExposePath) {
	v.value = val
	v.isSet = true
}

func (v NullableConsulExposePath) IsSet() bool {
	return v.isSet
}

func (v *NullableConsulExposePath) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsulExposePath(val *ConsulExposePath) *NullableConsulExposePath {
	return &NullableConsulExposePath{value: val, isSet: true}
}

func (v NullableConsulExposePath) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsulExposePath) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


