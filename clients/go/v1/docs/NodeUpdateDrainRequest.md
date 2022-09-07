# NodeUpdateDrainRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DrainSpec** | Pointer to [**NullableDrainSpec**](DrainSpec.md) |  | [optional] 
**MarkEligible** | Pointer to **bool** |  | [optional] 
**Meta** | Pointer to **map[string]string** |  | [optional] 
**NodeID** | Pointer to **string** |  | [optional] 

## Methods

### NewNodeUpdateDrainRequest

`func NewNodeUpdateDrainRequest() *NodeUpdateDrainRequest`

NewNodeUpdateDrainRequest instantiates a new NodeUpdateDrainRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeUpdateDrainRequestWithDefaults

`func NewNodeUpdateDrainRequestWithDefaults() *NodeUpdateDrainRequest`

NewNodeUpdateDrainRequestWithDefaults instantiates a new NodeUpdateDrainRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDrainSpec

`func (o *NodeUpdateDrainRequest) GetDrainSpec() DrainSpec`

GetDrainSpec returns the DrainSpec field if non-nil, zero value otherwise.

### GetDrainSpecOk

`func (o *NodeUpdateDrainRequest) GetDrainSpecOk() (*DrainSpec, bool)`

GetDrainSpecOk returns a tuple with the DrainSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrainSpec

`func (o *NodeUpdateDrainRequest) SetDrainSpec(v DrainSpec)`

SetDrainSpec sets DrainSpec field to given value.

### HasDrainSpec

`func (o *NodeUpdateDrainRequest) HasDrainSpec() bool`

HasDrainSpec returns a boolean if a field has been set.

### SetDrainSpecNil

`func (o *NodeUpdateDrainRequest) SetDrainSpecNil(b bool)`

 SetDrainSpecNil sets the value for DrainSpec to be an explicit nil

### UnsetDrainSpec
`func (o *NodeUpdateDrainRequest) UnsetDrainSpec()`

UnsetDrainSpec ensures that no value is present for DrainSpec, not even an explicit nil
### GetMarkEligible

`func (o *NodeUpdateDrainRequest) GetMarkEligible() bool`

GetMarkEligible returns the MarkEligible field if non-nil, zero value otherwise.

### GetMarkEligibleOk

`func (o *NodeUpdateDrainRequest) GetMarkEligibleOk() (*bool, bool)`

GetMarkEligibleOk returns a tuple with the MarkEligible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkEligible

`func (o *NodeUpdateDrainRequest) SetMarkEligible(v bool)`

SetMarkEligible sets MarkEligible field to given value.

### HasMarkEligible

`func (o *NodeUpdateDrainRequest) HasMarkEligible() bool`

HasMarkEligible returns a boolean if a field has been set.

### GetMeta

`func (o *NodeUpdateDrainRequest) GetMeta() map[string]string`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *NodeUpdateDrainRequest) GetMetaOk() (*map[string]string, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *NodeUpdateDrainRequest) SetMeta(v map[string]string)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *NodeUpdateDrainRequest) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetNodeID

`func (o *NodeUpdateDrainRequest) GetNodeID() string`

GetNodeID returns the NodeID field if non-nil, zero value otherwise.

### GetNodeIDOk

`func (o *NodeUpdateDrainRequest) GetNodeIDOk() (*string, bool)`

GetNodeIDOk returns a tuple with the NodeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeID

`func (o *NodeUpdateDrainRequest) SetNodeID(v string)`

SetNodeID sets NodeID field to given value.

### HasNodeID

`func (o *NodeUpdateDrainRequest) HasNodeID() bool`

HasNodeID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


