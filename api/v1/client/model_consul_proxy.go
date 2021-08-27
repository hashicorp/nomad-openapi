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

// ConsulProxy struct for ConsulProxy
type ConsulProxy struct {
	Config *map[string]interface{} `json:"Config,omitempty"`
	ExposeConfig *ConsulExposeConfig `json:"ExposeConfig,omitempty"`
	LocalServiceAddress *string `json:"LocalServiceAddress,omitempty"`
	LocalServicePort *int32 `json:"LocalServicePort,omitempty"`
	Upstreams *[]ConsulUpstream `json:"Upstreams,omitempty"`
}

// NewConsulProxy instantiates a new ConsulProxy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsulProxy() *ConsulProxy {
	this := ConsulProxy{}
	return &this
}

// NewConsulProxyWithDefaults instantiates a new ConsulProxy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsulProxyWithDefaults() *ConsulProxy {
	this := ConsulProxy{}
	return &this
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *ConsulProxy) GetConfig() map[string]interface{} {
	if o == nil || o.Config == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsulProxy) GetConfigOk() (*map[string]interface{}, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *ConsulProxy) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given map[string]interface{} and assigns it to the Config field.
func (o *ConsulProxy) SetConfig(v map[string]interface{}) {
	o.Config = &v
}

// GetExposeConfig returns the ExposeConfig field value if set, zero value otherwise.
func (o *ConsulProxy) GetExposeConfig() ConsulExposeConfig {
	if o == nil || o.ExposeConfig == nil {
		var ret ConsulExposeConfig
		return ret
	}
	return *o.ExposeConfig
}

// GetExposeConfigOk returns a tuple with the ExposeConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsulProxy) GetExposeConfigOk() (*ConsulExposeConfig, bool) {
	if o == nil || o.ExposeConfig == nil {
		return nil, false
	}
	return o.ExposeConfig, true
}

// HasExposeConfig returns a boolean if a field has been set.
func (o *ConsulProxy) HasExposeConfig() bool {
	if o != nil && o.ExposeConfig != nil {
		return true
	}

	return false
}

// SetExposeConfig gets a reference to the given ConsulExposeConfig and assigns it to the ExposeConfig field.
func (o *ConsulProxy) SetExposeConfig(v ConsulExposeConfig) {
	o.ExposeConfig = &v
}

// GetLocalServiceAddress returns the LocalServiceAddress field value if set, zero value otherwise.
func (o *ConsulProxy) GetLocalServiceAddress() string {
	if o == nil || o.LocalServiceAddress == nil {
		var ret string
		return ret
	}
	return *o.LocalServiceAddress
}

// GetLocalServiceAddressOk returns a tuple with the LocalServiceAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsulProxy) GetLocalServiceAddressOk() (*string, bool) {
	if o == nil || o.LocalServiceAddress == nil {
		return nil, false
	}
	return o.LocalServiceAddress, true
}

// HasLocalServiceAddress returns a boolean if a field has been set.
func (o *ConsulProxy) HasLocalServiceAddress() bool {
	if o != nil && o.LocalServiceAddress != nil {
		return true
	}

	return false
}

// SetLocalServiceAddress gets a reference to the given string and assigns it to the LocalServiceAddress field.
func (o *ConsulProxy) SetLocalServiceAddress(v string) {
	o.LocalServiceAddress = &v
}

// GetLocalServicePort returns the LocalServicePort field value if set, zero value otherwise.
func (o *ConsulProxy) GetLocalServicePort() int32 {
	if o == nil || o.LocalServicePort == nil {
		var ret int32
		return ret
	}
	return *o.LocalServicePort
}

// GetLocalServicePortOk returns a tuple with the LocalServicePort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsulProxy) GetLocalServicePortOk() (*int32, bool) {
	if o == nil || o.LocalServicePort == nil {
		return nil, false
	}
	return o.LocalServicePort, true
}

// HasLocalServicePort returns a boolean if a field has been set.
func (o *ConsulProxy) HasLocalServicePort() bool {
	if o != nil && o.LocalServicePort != nil {
		return true
	}

	return false
}

// SetLocalServicePort gets a reference to the given int32 and assigns it to the LocalServicePort field.
func (o *ConsulProxy) SetLocalServicePort(v int32) {
	o.LocalServicePort = &v
}

// GetUpstreams returns the Upstreams field value if set, zero value otherwise.
func (o *ConsulProxy) GetUpstreams() []ConsulUpstream {
	if o == nil || o.Upstreams == nil {
		var ret []ConsulUpstream
		return ret
	}
	return *o.Upstreams
}

// GetUpstreamsOk returns a tuple with the Upstreams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsulProxy) GetUpstreamsOk() (*[]ConsulUpstream, bool) {
	if o == nil || o.Upstreams == nil {
		return nil, false
	}
	return o.Upstreams, true
}

// HasUpstreams returns a boolean if a field has been set.
func (o *ConsulProxy) HasUpstreams() bool {
	if o != nil && o.Upstreams != nil {
		return true
	}

	return false
}

// SetUpstreams gets a reference to the given []ConsulUpstream and assigns it to the Upstreams field.
func (o *ConsulProxy) SetUpstreams(v []ConsulUpstream) {
	o.Upstreams = &v
}

func (o ConsulProxy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Config != nil {
		toSerialize["Config"] = o.Config
	}
	if o.ExposeConfig != nil {
		toSerialize["ExposeConfig"] = o.ExposeConfig
	}
	if o.LocalServiceAddress != nil {
		toSerialize["LocalServiceAddress"] = o.LocalServiceAddress
	}
	if o.LocalServicePort != nil {
		toSerialize["LocalServicePort"] = o.LocalServicePort
	}
	if o.Upstreams != nil {
		toSerialize["Upstreams"] = o.Upstreams
	}
	return json.Marshal(toSerialize)
}

type NullableConsulProxy struct {
	value *ConsulProxy
	isSet bool
}

func (v NullableConsulProxy) Get() *ConsulProxy {
	return v.value
}

func (v *NullableConsulProxy) Set(val *ConsulProxy) {
	v.value = val
	v.isSet = true
}

func (v NullableConsulProxy) IsSet() bool {
	return v.isSet
}

func (v *NullableConsulProxy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsulProxy(val *ConsulProxy) *NullableConsulProxy {
	return &NullableConsulProxy{value: val, isSet: true}
}

func (v NullableConsulProxy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsulProxy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

