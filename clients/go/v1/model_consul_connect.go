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

// ConsulConnect struct for ConsulConnect
type ConsulConnect struct {
	Gateway *ConsulGateway `json:"Gateway,omitempty"`
	Native *bool `json:"Native,omitempty"`
	SidecarService NullableConsulSidecarService `json:"SidecarService,omitempty"`
	SidecarTask NullableSidecarTask `json:"SidecarTask,omitempty"`
}

// NewConsulConnect instantiates a new ConsulConnect object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsulConnect() *ConsulConnect {
	this := ConsulConnect{}
	return &this
}

// NewConsulConnectWithDefaults instantiates a new ConsulConnect object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsulConnectWithDefaults() *ConsulConnect {
	this := ConsulConnect{}
	return &this
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *ConsulConnect) GetGateway() ConsulGateway {
	if o == nil || o.Gateway == nil {
		var ret ConsulGateway
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsulConnect) GetGatewayOk() (*ConsulGateway, bool) {
	if o == nil || o.Gateway == nil {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *ConsulConnect) HasGateway() bool {
	if o != nil && o.Gateway != nil {
		return true
	}

	return false
}

// SetGateway gets a reference to the given ConsulGateway and assigns it to the Gateway field.
func (o *ConsulConnect) SetGateway(v ConsulGateway) {
	o.Gateway = &v
}

// GetNative returns the Native field value if set, zero value otherwise.
func (o *ConsulConnect) GetNative() bool {
	if o == nil || o.Native == nil {
		var ret bool
		return ret
	}
	return *o.Native
}

// GetNativeOk returns a tuple with the Native field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsulConnect) GetNativeOk() (*bool, bool) {
	if o == nil || o.Native == nil {
		return nil, false
	}
	return o.Native, true
}

// HasNative returns a boolean if a field has been set.
func (o *ConsulConnect) HasNative() bool {
	if o != nil && o.Native != nil {
		return true
	}

	return false
}

// SetNative gets a reference to the given bool and assigns it to the Native field.
func (o *ConsulConnect) SetNative(v bool) {
	o.Native = &v
}

// GetSidecarService returns the SidecarService field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConsulConnect) GetSidecarService() ConsulSidecarService {
	if o == nil || o.SidecarService.Get() == nil {
		var ret ConsulSidecarService
		return ret
	}
	return *o.SidecarService.Get()
}

// GetSidecarServiceOk returns a tuple with the SidecarService field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConsulConnect) GetSidecarServiceOk() (*ConsulSidecarService, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SidecarService.Get(), o.SidecarService.IsSet()
}

// HasSidecarService returns a boolean if a field has been set.
func (o *ConsulConnect) HasSidecarService() bool {
	if o != nil && o.SidecarService.IsSet() {
		return true
	}

	return false
}

// SetSidecarService gets a reference to the given NullableConsulSidecarService and assigns it to the SidecarService field.
func (o *ConsulConnect) SetSidecarService(v ConsulSidecarService) {
	o.SidecarService.Set(&v)
}
// SetSidecarServiceNil sets the value for SidecarService to be an explicit nil
func (o *ConsulConnect) SetSidecarServiceNil() {
	o.SidecarService.Set(nil)
}

// UnsetSidecarService ensures that no value is present for SidecarService, not even an explicit nil
func (o *ConsulConnect) UnsetSidecarService() {
	o.SidecarService.Unset()
}

// GetSidecarTask returns the SidecarTask field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConsulConnect) GetSidecarTask() SidecarTask {
	if o == nil || o.SidecarTask.Get() == nil {
		var ret SidecarTask
		return ret
	}
	return *o.SidecarTask.Get()
}

// GetSidecarTaskOk returns a tuple with the SidecarTask field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConsulConnect) GetSidecarTaskOk() (*SidecarTask, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SidecarTask.Get(), o.SidecarTask.IsSet()
}

// HasSidecarTask returns a boolean if a field has been set.
func (o *ConsulConnect) HasSidecarTask() bool {
	if o != nil && o.SidecarTask.IsSet() {
		return true
	}

	return false
}

// SetSidecarTask gets a reference to the given NullableSidecarTask and assigns it to the SidecarTask field.
func (o *ConsulConnect) SetSidecarTask(v SidecarTask) {
	o.SidecarTask.Set(&v)
}
// SetSidecarTaskNil sets the value for SidecarTask to be an explicit nil
func (o *ConsulConnect) SetSidecarTaskNil() {
	o.SidecarTask.Set(nil)
}

// UnsetSidecarTask ensures that no value is present for SidecarTask, not even an explicit nil
func (o *ConsulConnect) UnsetSidecarTask() {
	o.SidecarTask.Unset()
}

func (o ConsulConnect) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Gateway != nil {
		toSerialize["Gateway"] = o.Gateway
	}
	if o.Native != nil {
		toSerialize["Native"] = o.Native
	}
	if o.SidecarService.IsSet() {
		toSerialize["SidecarService"] = o.SidecarService.Get()
	}
	if o.SidecarTask.IsSet() {
		toSerialize["SidecarTask"] = o.SidecarTask.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableConsulConnect struct {
	value *ConsulConnect
	isSet bool
}

func (v NullableConsulConnect) Get() *ConsulConnect {
	return v.value
}

func (v *NullableConsulConnect) Set(val *ConsulConnect) {
	v.value = val
	v.isSet = true
}

func (v NullableConsulConnect) IsSet() bool {
	return v.isSet
}

func (v *NullableConsulConnect) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsulConnect(val *ConsulConnect) *NullableConsulConnect {
	return &NullableConsulConnect{value: val, isSet: true}
}

func (v NullableConsulConnect) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsulConnect) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


