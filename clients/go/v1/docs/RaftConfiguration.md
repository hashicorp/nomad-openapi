# RaftConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | Pointer to **int32** |  | [optional] 
**Servers** | Pointer to [**[]RaftServer**](RaftServer.md) |  | [optional] 

## Methods

### NewRaftConfiguration

`func NewRaftConfiguration() *RaftConfiguration`

NewRaftConfiguration instantiates a new RaftConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRaftConfigurationWithDefaults

`func NewRaftConfigurationWithDefaults() *RaftConfiguration`

NewRaftConfigurationWithDefaults instantiates a new RaftConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *RaftConfiguration) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *RaftConfiguration) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *RaftConfiguration) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *RaftConfiguration) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetServers

`func (o *RaftConfiguration) GetServers() []RaftServer`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *RaftConfiguration) GetServersOk() (*[]RaftServer, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *RaftConfiguration) SetServers(v []RaftServer)`

SetServers sets Servers field to given value.

### HasServers

`func (o *RaftConfiguration) HasServers() bool`

HasServers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


