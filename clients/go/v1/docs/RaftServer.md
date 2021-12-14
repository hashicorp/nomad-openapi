# RaftServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** |  | [optional] 
**ID** | Pointer to **string** |  | [optional] 
**Leader** | Pointer to **bool** |  | [optional] 
**Node** | Pointer to **string** |  | [optional] 
**RaftProtocol** | Pointer to **string** |  | [optional] 
**Voter** | Pointer to **bool** |  | [optional] 

## Methods

### NewRaftServer

`func NewRaftServer() *RaftServer`

NewRaftServer instantiates a new RaftServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRaftServerWithDefaults

`func NewRaftServerWithDefaults() *RaftServer`

NewRaftServerWithDefaults instantiates a new RaftServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *RaftServer) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *RaftServer) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *RaftServer) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *RaftServer) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetID

`func (o *RaftServer) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *RaftServer) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *RaftServer) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *RaftServer) HasID() bool`

HasID returns a boolean if a field has been set.

### GetLeader

`func (o *RaftServer) GetLeader() bool`

GetLeader returns the Leader field if non-nil, zero value otherwise.

### GetLeaderOk

`func (o *RaftServer) GetLeaderOk() (*bool, bool)`

GetLeaderOk returns a tuple with the Leader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeader

`func (o *RaftServer) SetLeader(v bool)`

SetLeader sets Leader field to given value.

### HasLeader

`func (o *RaftServer) HasLeader() bool`

HasLeader returns a boolean if a field has been set.

### GetNode

`func (o *RaftServer) GetNode() string`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *RaftServer) GetNodeOk() (*string, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *RaftServer) SetNode(v string)`

SetNode sets Node field to given value.

### HasNode

`func (o *RaftServer) HasNode() bool`

HasNode returns a boolean if a field has been set.

### GetRaftProtocol

`func (o *RaftServer) GetRaftProtocol() string`

GetRaftProtocol returns the RaftProtocol field if non-nil, zero value otherwise.

### GetRaftProtocolOk

`func (o *RaftServer) GetRaftProtocolOk() (*string, bool)`

GetRaftProtocolOk returns a tuple with the RaftProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaftProtocol

`func (o *RaftServer) SetRaftProtocol(v string)`

SetRaftProtocol sets RaftProtocol field to given value.

### HasRaftProtocol

`func (o *RaftServer) HasRaftProtocol() bool`

HasRaftProtocol returns a boolean if a field has been set.

### GetVoter

`func (o *RaftServer) GetVoter() bool`

GetVoter returns the Voter field if non-nil, zero value otherwise.

### GetVoterOk

`func (o *RaftServer) GetVoterOk() (*bool, bool)`

GetVoterOk returns a tuple with the Voter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoter

`func (o *RaftServer) SetVoter(v bool)`

SetVoter sets Voter field to given value.

### HasVoter

`func (o *RaftServer) HasVoter() bool`

HasVoter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


