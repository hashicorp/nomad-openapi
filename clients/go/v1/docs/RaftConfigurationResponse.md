# RaftConfigurationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | Pointer to **int32** |  | [optional] 
**Servers** | Pointer to [**[]RaftServer**](RaftServer.md) |  | [optional] 

## Methods

### NewRaftConfigurationResponse

`func NewRaftConfigurationResponse() *RaftConfigurationResponse`

NewRaftConfigurationResponse instantiates a new RaftConfigurationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRaftConfigurationResponseWithDefaults

`func NewRaftConfigurationResponseWithDefaults() *RaftConfigurationResponse`

NewRaftConfigurationResponseWithDefaults instantiates a new RaftConfigurationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *RaftConfigurationResponse) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *RaftConfigurationResponse) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *RaftConfigurationResponse) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *RaftConfigurationResponse) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetServers

`func (o *RaftConfigurationResponse) GetServers() []RaftServer`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *RaftConfigurationResponse) GetServersOk() (*[]RaftServer, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *RaftConfigurationResponse) SetServers(v []RaftServer)`

SetServers sets Servers field to given value.

### HasServers

`func (o *RaftConfigurationResponse) HasServers() bool`

HasServers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


