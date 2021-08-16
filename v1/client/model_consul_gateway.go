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

// ConsulGateway struct for ConsulGateway
type ConsulGateway struct {
	Ingress *ConsulIngressConfigEntry `json:"Ingress,omitempty"`
	Mesh interface{} `json:"Mesh,omitempty"`
	Proxy *ConsulGatewayProxy `json:"Proxy,omitempty"`
	Terminating *ConsulTerminatingConfigEntry `json:"Terminating,omitempty"`
}

// NewConsulGateway instantiates a new ConsulGateway object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsulGateway() *ConsulGateway {
	this := ConsulGateway{}
	return &this
}

// NewConsulGatewayWithDefaults instantiates a new ConsulGateway object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsulGatewayWithDefaults() *ConsulGateway {
	this := ConsulGateway{}
	return &this
}

// GetIngress returns the Ingress field value if set, zero value otherwise.
func (o *ConsulGateway) GetIngress() ConsulIngressConfigEntry {
	if o == nil || o.Ingress == nil {
		var ret ConsulIngressConfigEntry
		return ret
	}
	return *o.Ingress
}

// GetIngressOk returns a tuple with the Ingress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsulGateway) GetIngressOk() (*ConsulIngressConfigEntry, bool) {
	if o == nil || o.Ingress == nil {
		return nil, false
	}
	return o.Ingress, true
}

// HasIngress returns a boolean if a field has been set.
func (o *ConsulGateway) HasIngress() bool {
	if o != nil && o.Ingress != nil {
		return true
	}

	return false
}

// SetIngress gets a reference to the given ConsulIngressConfigEntry and assigns it to the Ingress field.
func (o *ConsulGateway) SetIngress(v ConsulIngressConfigEntry) {
	o.Ingress = &v
}

// GetMesh returns the Mesh field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConsulGateway) GetMesh() interface{} {
	if o == nil  {
		var ret interface{}
		return ret
	}
	return o.Mesh
}

// GetMeshOk returns a tuple with the Mesh field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConsulGateway) GetMeshOk() (*interface{}, bool) {
	if o == nil || o.Mesh == nil {
		return nil, false
	}
	return &o.Mesh, true
}

// HasMesh returns a boolean if a field has been set.
func (o *ConsulGateway) HasMesh() bool {
	if o != nil && o.Mesh != nil {
		return true
	}

	return false
}

// SetMesh gets a reference to the given interface{} and assigns it to the Mesh field.
func (o *ConsulGateway) SetMesh(v interface{}) {
	o.Mesh = v
}

// GetProxy returns the Proxy field value if set, zero value otherwise.
func (o *ConsulGateway) GetProxy() ConsulGatewayProxy {
	if o == nil || o.Proxy == nil {
		var ret ConsulGatewayProxy
		return ret
	}
	return *o.Proxy
}

// GetProxyOk returns a tuple with the Proxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsulGateway) GetProxyOk() (*ConsulGatewayProxy, bool) {
	if o == nil || o.Proxy == nil {
		return nil, false
	}
	return o.Proxy, true
}

// HasProxy returns a boolean if a field has been set.
func (o *ConsulGateway) HasProxy() bool {
	if o != nil && o.Proxy != nil {
		return true
	}

	return false
}

// SetProxy gets a reference to the given ConsulGatewayProxy and assigns it to the Proxy field.
func (o *ConsulGateway) SetProxy(v ConsulGatewayProxy) {
	o.Proxy = &v
}

// GetTerminating returns the Terminating field value if set, zero value otherwise.
func (o *ConsulGateway) GetTerminating() ConsulTerminatingConfigEntry {
	if o == nil || o.Terminating == nil {
		var ret ConsulTerminatingConfigEntry
		return ret
	}
	return *o.Terminating
}

// GetTerminatingOk returns a tuple with the Terminating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsulGateway) GetTerminatingOk() (*ConsulTerminatingConfigEntry, bool) {
	if o == nil || o.Terminating == nil {
		return nil, false
	}
	return o.Terminating, true
}

// HasTerminating returns a boolean if a field has been set.
func (o *ConsulGateway) HasTerminating() bool {
	if o != nil && o.Terminating != nil {
		return true
	}

	return false
}

// SetTerminating gets a reference to the given ConsulTerminatingConfigEntry and assigns it to the Terminating field.
func (o *ConsulGateway) SetTerminating(v ConsulTerminatingConfigEntry) {
	o.Terminating = &v
}

func (o ConsulGateway) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Ingress != nil {
		toSerialize["Ingress"] = o.Ingress
	}
	if o.Mesh != nil {
		toSerialize["Mesh"] = o.Mesh
	}
	if o.Proxy != nil {
		toSerialize["Proxy"] = o.Proxy
	}
	if o.Terminating != nil {
		toSerialize["Terminating"] = o.Terminating
	}
	return json.Marshal(toSerialize)
}

type NullableConsulGateway struct {
	value *ConsulGateway
	isSet bool
}

func (v NullableConsulGateway) Get() *ConsulGateway {
	return v.value
}

func (v *NullableConsulGateway) Set(val *ConsulGateway) {
	v.value = val
	v.isSet = true
}

func (v NullableConsulGateway) IsSet() bool {
	return v.isSet
}

func (v *NullableConsulGateway) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsulGateway(val *ConsulGateway) *NullableConsulGateway {
	return &NullableConsulGateway{value: val, isSet: true}
}

func (v NullableConsulGateway) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsulGateway) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


