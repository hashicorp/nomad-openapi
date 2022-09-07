# ServerHealth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** |  | [optional] 
**Healthy** | Pointer to **bool** |  | [optional] 
**ID** | Pointer to **string** |  | [optional] 
**LastContact** | Pointer to **int64** |  | [optional] 
**LastIndex** | Pointer to **int32** |  | [optional] 
**LastTerm** | Pointer to **int32** |  | [optional] 
**Leader** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**SerfStatus** | Pointer to **string** |  | [optional] 
**StableSince** | Pointer to **NullableTime** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Voter** | Pointer to **bool** |  | [optional] 

## Methods

### NewServerHealth

`func NewServerHealth() *ServerHealth`

NewServerHealth instantiates a new ServerHealth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerHealthWithDefaults

`func NewServerHealthWithDefaults() *ServerHealth`

NewServerHealthWithDefaults instantiates a new ServerHealth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *ServerHealth) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ServerHealth) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ServerHealth) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ServerHealth) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetHealthy

`func (o *ServerHealth) GetHealthy() bool`

GetHealthy returns the Healthy field if non-nil, zero value otherwise.

### GetHealthyOk

`func (o *ServerHealth) GetHealthyOk() (*bool, bool)`

GetHealthyOk returns a tuple with the Healthy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthy

`func (o *ServerHealth) SetHealthy(v bool)`

SetHealthy sets Healthy field to given value.

### HasHealthy

`func (o *ServerHealth) HasHealthy() bool`

HasHealthy returns a boolean if a field has been set.

### GetID

`func (o *ServerHealth) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *ServerHealth) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *ServerHealth) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *ServerHealth) HasID() bool`

HasID returns a boolean if a field has been set.

### GetLastContact

`func (o *ServerHealth) GetLastContact() int64`

GetLastContact returns the LastContact field if non-nil, zero value otherwise.

### GetLastContactOk

`func (o *ServerHealth) GetLastContactOk() (*int64, bool)`

GetLastContactOk returns a tuple with the LastContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastContact

`func (o *ServerHealth) SetLastContact(v int64)`

SetLastContact sets LastContact field to given value.

### HasLastContact

`func (o *ServerHealth) HasLastContact() bool`

HasLastContact returns a boolean if a field has been set.

### GetLastIndex

`func (o *ServerHealth) GetLastIndex() int32`

GetLastIndex returns the LastIndex field if non-nil, zero value otherwise.

### GetLastIndexOk

`func (o *ServerHealth) GetLastIndexOk() (*int32, bool)`

GetLastIndexOk returns a tuple with the LastIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastIndex

`func (o *ServerHealth) SetLastIndex(v int32)`

SetLastIndex sets LastIndex field to given value.

### HasLastIndex

`func (o *ServerHealth) HasLastIndex() bool`

HasLastIndex returns a boolean if a field has been set.

### GetLastTerm

`func (o *ServerHealth) GetLastTerm() int32`

GetLastTerm returns the LastTerm field if non-nil, zero value otherwise.

### GetLastTermOk

`func (o *ServerHealth) GetLastTermOk() (*int32, bool)`

GetLastTermOk returns a tuple with the LastTerm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTerm

`func (o *ServerHealth) SetLastTerm(v int32)`

SetLastTerm sets LastTerm field to given value.

### HasLastTerm

`func (o *ServerHealth) HasLastTerm() bool`

HasLastTerm returns a boolean if a field has been set.

### GetLeader

`func (o *ServerHealth) GetLeader() bool`

GetLeader returns the Leader field if non-nil, zero value otherwise.

### GetLeaderOk

`func (o *ServerHealth) GetLeaderOk() (*bool, bool)`

GetLeaderOk returns a tuple with the Leader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeader

`func (o *ServerHealth) SetLeader(v bool)`

SetLeader sets Leader field to given value.

### HasLeader

`func (o *ServerHealth) HasLeader() bool`

HasLeader returns a boolean if a field has been set.

### GetName

`func (o *ServerHealth) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServerHealth) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServerHealth) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServerHealth) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSerfStatus

`func (o *ServerHealth) GetSerfStatus() string`

GetSerfStatus returns the SerfStatus field if non-nil, zero value otherwise.

### GetSerfStatusOk

`func (o *ServerHealth) GetSerfStatusOk() (*string, bool)`

GetSerfStatusOk returns a tuple with the SerfStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerfStatus

`func (o *ServerHealth) SetSerfStatus(v string)`

SetSerfStatus sets SerfStatus field to given value.

### HasSerfStatus

`func (o *ServerHealth) HasSerfStatus() bool`

HasSerfStatus returns a boolean if a field has been set.

### GetStableSince

`func (o *ServerHealth) GetStableSince() time.Time`

GetStableSince returns the StableSince field if non-nil, zero value otherwise.

### GetStableSinceOk

`func (o *ServerHealth) GetStableSinceOk() (*time.Time, bool)`

GetStableSinceOk returns a tuple with the StableSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStableSince

`func (o *ServerHealth) SetStableSince(v time.Time)`

SetStableSince sets StableSince field to given value.

### HasStableSince

`func (o *ServerHealth) HasStableSince() bool`

HasStableSince returns a boolean if a field has been set.

### SetStableSinceNil

`func (o *ServerHealth) SetStableSinceNil(b bool)`

 SetStableSinceNil sets the value for StableSince to be an explicit nil

### UnsetStableSince
`func (o *ServerHealth) UnsetStableSince()`

UnsetStableSince ensures that no value is present for StableSince, not even an explicit nil
### GetVersion

`func (o *ServerHealth) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ServerHealth) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ServerHealth) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ServerHealth) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVoter

`func (o *ServerHealth) GetVoter() bool`

GetVoter returns the Voter field if non-nil, zero value otherwise.

### GetVoterOk

`func (o *ServerHealth) GetVoterOk() (*bool, bool)`

GetVoterOk returns a tuple with the Voter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoter

`func (o *ServerHealth) SetVoter(v bool)`

SetVoter sets Voter field to given value.

### HasVoter

`func (o *ServerHealth) HasVoter() bool`

HasVoter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


