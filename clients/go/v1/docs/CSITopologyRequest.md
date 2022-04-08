# CSITopologyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Preferred** | Pointer to [**[]CSITopology**](CSITopology.md) |  | [optional] 
**Required** | Pointer to [**[]CSITopology**](CSITopology.md) |  | [optional] 

## Methods

### NewCSITopologyRequest

`func NewCSITopologyRequest() *CSITopologyRequest`

NewCSITopologyRequest instantiates a new CSITopologyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSITopologyRequestWithDefaults

`func NewCSITopologyRequestWithDefaults() *CSITopologyRequest`

NewCSITopologyRequestWithDefaults instantiates a new CSITopologyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPreferred

`func (o *CSITopologyRequest) GetPreferred() []CSITopology`

GetPreferred returns the Preferred field if non-nil, zero value otherwise.

### GetPreferredOk

`func (o *CSITopologyRequest) GetPreferredOk() (*[]CSITopology, bool)`

GetPreferredOk returns a tuple with the Preferred field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferred

`func (o *CSITopologyRequest) SetPreferred(v []CSITopology)`

SetPreferred sets Preferred field to given value.

### HasPreferred

`func (o *CSITopologyRequest) HasPreferred() bool`

HasPreferred returns a boolean if a field has been set.

### GetRequired

`func (o *CSITopologyRequest) GetRequired() []CSITopology`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *CSITopologyRequest) GetRequiredOk() (*[]CSITopology, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *CSITopologyRequest) SetRequired(v []CSITopology)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *CSITopologyRequest) HasRequired() bool`

HasRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


