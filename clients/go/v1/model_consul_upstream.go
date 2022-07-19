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

// ConsulUpstream struct for ConsulUpstream
type ConsulUpstream struct {
	Datacenter *string `json:"Datacenter,omitempty"`
	DestinationName *string `json:"DestinationName,omitempty"`
	DestinationNamespace *string `json:"DestinationNamespace,omitempty"`
	LocalBindAddress *string `json:"LocalBindAddress,omitempty"`
	LocalBindPort *int32 `json:"LocalBindPort,omitempty"`
	MeshGateway *ConsulMeshGateway `json:"MeshGateway,omitempty"`
}

// NewConsulUpstream instantiates a new ConsulUpstream object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsulUpstream() *ConsulUpstream {
	this := ConsulUpstream{}
	return &this
}

// NewConsulUpstreamWithDefaults instantiates a new ConsulUpstream object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsulUpstreamWithDefaults() *ConsulUpstream {
	this := ConsulUpstream{}
	return &this
}

// GetDatacenter returns the Datacenter field value if set, zero value otherwise.
func (o *ConsulUpstream) GetDatacenter() string {
	if o == nil || o.Datacenter == nil {
		var ret string
		return ret
	}
	return *o.Datacenter
}

// GetDatacenterOk returns a tuple with the Datacenter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsulUpstream) GetDatacenterOk() (*string, bool) {
	if o == nil || o.Datacenter == nil {
		return nil, false
	}
	return o.Datacenter, true
}

// HasDatacenter returns a boolean if a field has been set.
func (o *ConsulUpstream) HasDatacenter() bool {
	if o != nil && o.Datacenter != nil {
		return true
	}

	return false
}

// SetDatacenter gets a reference to the given string and assigns it to the Datacenter field.
func (o *ConsulUpstream) SetDatacenter(v string) {
	o.Datacenter = &v
}

// GetDestinationName returns the DestinationName field value if set, zero value otherwise.
func (o *ConsulUpstream) GetDestinationName() string {
	if o == nil || o.DestinationName == nil {
		var ret string
		return ret
	}
	return *o.DestinationName
}

// GetDestinationNameOk returns a tuple with the DestinationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsulUpstream) GetDestinationNameOk() (*string, bool) {
	if o == nil || o.DestinationName == nil {
		return nil, false
	}
	return o.DestinationName, true
}

// HasDestinationName returns a boolean if a field has been set.
func (o *ConsulUpstream) HasDestinationName() bool {
	if o != nil && o.DestinationName != nil {
		return true
	}

	return false
}

// SetDestinationName gets a reference to the given string and assigns it to the DestinationName field.
func (o *ConsulUpstream) SetDestinationName(v string) {
	o.DestinationName = &v
}

// GetDestinationNamespace returns the DestinationNamespace field value if set, zero value otherwise.
func (o *ConsulUpstream) GetDestinationNamespace() string {
	if o == nil || o.DestinationNamespace == nil {
		var ret string
		return ret
	}
	return *o.DestinationNamespace
}

// GetDestinationNamespaceOk returns a tuple with the DestinationNamespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsulUpstream) GetDestinationNamespaceOk() (*string, bool) {
	if o == nil || o.DestinationNamespace == nil {
		return nil, false
	}
	return o.DestinationNamespace, true
}

// HasDestinationNamespace returns a boolean if a field has been set.
func (o *ConsulUpstream) HasDestinationNamespace() bool {
	if o != nil && o.DestinationNamespace != nil {
		return true
	}

	return false
}

// SetDestinationNamespace gets a reference to the given string and assigns it to the DestinationNamespace field.
func (o *ConsulUpstream) SetDestinationNamespace(v string) {
	o.DestinationNamespace = &v
}

// GetLocalBindAddress returns the LocalBindAddress field value if set, zero value otherwise.
func (o *ConsulUpstream) GetLocalBindAddress() string {
	if o == nil || o.LocalBindAddress == nil {
		var ret string
		return ret
	}
	return *o.LocalBindAddress
}

// GetLocalBindAddressOk returns a tuple with the LocalBindAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsulUpstream) GetLocalBindAddressOk() (*string, bool) {
	if o == nil || o.LocalBindAddress == nil {
		return nil, false
	}
	return o.LocalBindAddress, true
}

// HasLocalBindAddress returns a boolean if a field has been set.
func (o *ConsulUpstream) HasLocalBindAddress() bool {
	if o != nil && o.LocalBindAddress != nil {
		return true
	}

	return false
}

// SetLocalBindAddress gets a reference to the given string and assigns it to the LocalBindAddress field.
func (o *ConsulUpstream) SetLocalBindAddress(v string) {
	o.LocalBindAddress = &v
}

// GetLocalBindPort returns the LocalBindPort field value if set, zero value otherwise.
func (o *ConsulUpstream) GetLocalBindPort() int32 {
	if o == nil || o.LocalBindPort == nil {
		var ret int32
		return ret
	}
	return *o.LocalBindPort
}

// GetLocalBindPortOk returns a tuple with the LocalBindPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsulUpstream) GetLocalBindPortOk() (*int32, bool) {
	if o == nil || o.LocalBindPort == nil {
		return nil, false
	}
	return o.LocalBindPort, true
}

// HasLocalBindPort returns a boolean if a field has been set.
func (o *ConsulUpstream) HasLocalBindPort() bool {
	if o != nil && o.LocalBindPort != nil {
		return true
	}

	return false
}

// SetLocalBindPort gets a reference to the given int32 and assigns it to the LocalBindPort field.
func (o *ConsulUpstream) SetLocalBindPort(v int32) {
	o.LocalBindPort = &v
}

// GetMeshGateway returns the MeshGateway field value if set, zero value otherwise.
func (o *ConsulUpstream) GetMeshGateway() ConsulMeshGateway {
	if o == nil || o.MeshGateway == nil {
		var ret ConsulMeshGateway
		return ret
	}
	return *o.MeshGateway
}

// GetMeshGatewayOk returns a tuple with the MeshGateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsulUpstream) GetMeshGatewayOk() (*ConsulMeshGateway, bool) {
	if o == nil || o.MeshGateway == nil {
		return nil, false
	}
	return o.MeshGateway, true
}

// HasMeshGateway returns a boolean if a field has been set.
func (o *ConsulUpstream) HasMeshGateway() bool {
	if o != nil && o.MeshGateway != nil {
		return true
	}

	return false
}

// SetMeshGateway gets a reference to the given ConsulMeshGateway and assigns it to the MeshGateway field.
func (o *ConsulUpstream) SetMeshGateway(v ConsulMeshGateway) {
	o.MeshGateway = &v
}

func (o ConsulUpstream) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Datacenter != nil {
		toSerialize["Datacenter"] = o.Datacenter
	}
	if o.DestinationName != nil {
		toSerialize["DestinationName"] = o.DestinationName
	}
	if o.DestinationNamespace != nil {
		toSerialize["DestinationNamespace"] = o.DestinationNamespace
	}
	if o.LocalBindAddress != nil {
		toSerialize["LocalBindAddress"] = o.LocalBindAddress
	}
	if o.LocalBindPort != nil {
		toSerialize["LocalBindPort"] = o.LocalBindPort
	}
	if o.MeshGateway != nil {
		toSerialize["MeshGateway"] = o.MeshGateway
	}
	return json.Marshal(toSerialize)
}

type NullableConsulUpstream struct {
	value *ConsulUpstream
	isSet bool
}

func (v NullableConsulUpstream) Get() *ConsulUpstream {
	return v.value
}

func (v *NullableConsulUpstream) Set(val *ConsulUpstream) {
	v.value = val
	v.isSet = true
}

func (v NullableConsulUpstream) IsSet() bool {
	return v.isSet
}

func (v *NullableConsulUpstream) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsulUpstream(val *ConsulUpstream) *NullableConsulUpstream {
	return &NullableConsulUpstream{value: val, isSet: true}
}

func (v NullableConsulUpstream) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsulUpstream) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


