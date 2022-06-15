# NodeDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HealthDescription** | Pointer to **string** |  | [optional] 
**Healthy** | Pointer to **bool** |  | [optional] 
**ID** | Pointer to **string** |  | [optional] 
**Locality** | Pointer to [**NullableNodeDeviceLocality**](NodeDeviceLocality.md) |  | [optional] 

## Methods

### NewNodeDevice

`func NewNodeDevice() *NodeDevice`

NewNodeDevice instantiates a new NodeDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeDeviceWithDefaults

`func NewNodeDeviceWithDefaults() *NodeDevice`

NewNodeDeviceWithDefaults instantiates a new NodeDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHealthDescription

`func (o *NodeDevice) GetHealthDescription() string`

GetHealthDescription returns the HealthDescription field if non-nil, zero value otherwise.

### GetHealthDescriptionOk

`func (o *NodeDevice) GetHealthDescriptionOk() (*string, bool)`

GetHealthDescriptionOk returns a tuple with the HealthDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthDescription

`func (o *NodeDevice) SetHealthDescription(v string)`

SetHealthDescription sets HealthDescription field to given value.

### HasHealthDescription

`func (o *NodeDevice) HasHealthDescription() bool`

HasHealthDescription returns a boolean if a field has been set.

### GetHealthy

`func (o *NodeDevice) GetHealthy() bool`

GetHealthy returns the Healthy field if non-nil, zero value otherwise.

### GetHealthyOk

`func (o *NodeDevice) GetHealthyOk() (*bool, bool)`

GetHealthyOk returns a tuple with the Healthy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthy

`func (o *NodeDevice) SetHealthy(v bool)`

SetHealthy sets Healthy field to given value.

### HasHealthy

`func (o *NodeDevice) HasHealthy() bool`

HasHealthy returns a boolean if a field has been set.

### GetID

`func (o *NodeDevice) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *NodeDevice) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *NodeDevice) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *NodeDevice) HasID() bool`

HasID returns a boolean if a field has been set.

### GetLocality

`func (o *NodeDevice) GetLocality() NodeDeviceLocality`

GetLocality returns the Locality field if non-nil, zero value otherwise.

### GetLocalityOk

`func (o *NodeDevice) GetLocalityOk() (*NodeDeviceLocality, bool)`

GetLocalityOk returns a tuple with the Locality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocality

`func (o *NodeDevice) SetLocality(v NodeDeviceLocality)`

SetLocality sets Locality field to given value.

### HasLocality

`func (o *NodeDevice) HasLocality() bool`

HasLocality returns a boolean if a field has been set.

### SetLocalityNil

`func (o *NodeDevice) SetLocalityNil(b bool)`

 SetLocalityNil sets the value for Locality to be an explicit nil

### UnsetLocality
`func (o *NodeDevice) UnsetLocality()`

UnsetLocality ensures that no value is present for Locality, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


