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

// RaftServer struct for RaftServer
type RaftServer struct {
	Address *string `json:"Address,omitempty"`
	ID *string `json:"ID,omitempty"`
	Leader *bool `json:"Leader,omitempty"`
	Node *string `json:"Node,omitempty"`
	RaftProtocol *string `json:"RaftProtocol,omitempty"`
	Voter *bool `json:"Voter,omitempty"`
}

// NewRaftServer instantiates a new RaftServer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRaftServer() *RaftServer {
	this := RaftServer{}
	return &this
}

// NewRaftServerWithDefaults instantiates a new RaftServer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRaftServerWithDefaults() *RaftServer {
	this := RaftServer{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *RaftServer) GetAddress() string {
	if o == nil || o.Address == nil {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RaftServer) GetAddressOk() (*string, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *RaftServer) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *RaftServer) SetAddress(v string) {
	o.Address = &v
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *RaftServer) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RaftServer) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *RaftServer) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *RaftServer) SetID(v string) {
	o.ID = &v
}

// GetLeader returns the Leader field value if set, zero value otherwise.
func (o *RaftServer) GetLeader() bool {
	if o == nil || o.Leader == nil {
		var ret bool
		return ret
	}
	return *o.Leader
}

// GetLeaderOk returns a tuple with the Leader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RaftServer) GetLeaderOk() (*bool, bool) {
	if o == nil || o.Leader == nil {
		return nil, false
	}
	return o.Leader, true
}

// HasLeader returns a boolean if a field has been set.
func (o *RaftServer) HasLeader() bool {
	if o != nil && o.Leader != nil {
		return true
	}

	return false
}

// SetLeader gets a reference to the given bool and assigns it to the Leader field.
func (o *RaftServer) SetLeader(v bool) {
	o.Leader = &v
}

// GetNode returns the Node field value if set, zero value otherwise.
func (o *RaftServer) GetNode() string {
	if o == nil || o.Node == nil {
		var ret string
		return ret
	}
	return *o.Node
}

// GetNodeOk returns a tuple with the Node field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RaftServer) GetNodeOk() (*string, bool) {
	if o == nil || o.Node == nil {
		return nil, false
	}
	return o.Node, true
}

// HasNode returns a boolean if a field has been set.
func (o *RaftServer) HasNode() bool {
	if o != nil && o.Node != nil {
		return true
	}

	return false
}

// SetNode gets a reference to the given string and assigns it to the Node field.
func (o *RaftServer) SetNode(v string) {
	o.Node = &v
}

// GetRaftProtocol returns the RaftProtocol field value if set, zero value otherwise.
func (o *RaftServer) GetRaftProtocol() string {
	if o == nil || o.RaftProtocol == nil {
		var ret string
		return ret
	}
	return *o.RaftProtocol
}

// GetRaftProtocolOk returns a tuple with the RaftProtocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RaftServer) GetRaftProtocolOk() (*string, bool) {
	if o == nil || o.RaftProtocol == nil {
		return nil, false
	}
	return o.RaftProtocol, true
}

// HasRaftProtocol returns a boolean if a field has been set.
func (o *RaftServer) HasRaftProtocol() bool {
	if o != nil && o.RaftProtocol != nil {
		return true
	}

	return false
}

// SetRaftProtocol gets a reference to the given string and assigns it to the RaftProtocol field.
func (o *RaftServer) SetRaftProtocol(v string) {
	o.RaftProtocol = &v
}

// GetVoter returns the Voter field value if set, zero value otherwise.
func (o *RaftServer) GetVoter() bool {
	if o == nil || o.Voter == nil {
		var ret bool
		return ret
	}
	return *o.Voter
}

// GetVoterOk returns a tuple with the Voter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RaftServer) GetVoterOk() (*bool, bool) {
	if o == nil || o.Voter == nil {
		return nil, false
	}
	return o.Voter, true
}

// HasVoter returns a boolean if a field has been set.
func (o *RaftServer) HasVoter() bool {
	if o != nil && o.Voter != nil {
		return true
	}

	return false
}

// SetVoter gets a reference to the given bool and assigns it to the Voter field.
func (o *RaftServer) SetVoter(v bool) {
	o.Voter = &v
}

func (o RaftServer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Address != nil {
		toSerialize["Address"] = o.Address
	}
	if o.ID != nil {
		toSerialize["ID"] = o.ID
	}
	if o.Leader != nil {
		toSerialize["Leader"] = o.Leader
	}
	if o.Node != nil {
		toSerialize["Node"] = o.Node
	}
	if o.RaftProtocol != nil {
		toSerialize["RaftProtocol"] = o.RaftProtocol
	}
	if o.Voter != nil {
		toSerialize["Voter"] = o.Voter
	}
	return json.Marshal(toSerialize)
}

type NullableRaftServer struct {
	value *RaftServer
	isSet bool
}

func (v NullableRaftServer) Get() *RaftServer {
	return v.value
}

func (v *NullableRaftServer) Set(val *RaftServer) {
	v.value = val
	v.isSet = true
}

func (v NullableRaftServer) IsSet() bool {
	return v.isSet
}

func (v *NullableRaftServer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRaftServer(val *RaftServer) *NullableRaftServer {
	return &NullableRaftServer{value: val, isSet: true}
}

func (v NullableRaftServer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRaftServer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


