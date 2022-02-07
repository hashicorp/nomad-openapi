# OperatorHealthReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailureTolerance** | Pointer to **int32** |  | [optional] 
**Healthy** | Pointer to **bool** |  | [optional] 
**Servers** | Pointer to [**[]ServerHealth**](ServerHealth.md) |  | [optional] 

## Methods

### NewOperatorHealthReply

`func NewOperatorHealthReply() *OperatorHealthReply`

NewOperatorHealthReply instantiates a new OperatorHealthReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperatorHealthReplyWithDefaults

`func NewOperatorHealthReplyWithDefaults() *OperatorHealthReply`

NewOperatorHealthReplyWithDefaults instantiates a new OperatorHealthReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailureTolerance

`func (o *OperatorHealthReply) GetFailureTolerance() int32`

GetFailureTolerance returns the FailureTolerance field if non-nil, zero value otherwise.

### GetFailureToleranceOk

`func (o *OperatorHealthReply) GetFailureToleranceOk() (*int32, bool)`

GetFailureToleranceOk returns a tuple with the FailureTolerance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureTolerance

`func (o *OperatorHealthReply) SetFailureTolerance(v int32)`

SetFailureTolerance sets FailureTolerance field to given value.

### HasFailureTolerance

`func (o *OperatorHealthReply) HasFailureTolerance() bool`

HasFailureTolerance returns a boolean if a field has been set.

### GetHealthy

`func (o *OperatorHealthReply) GetHealthy() bool`

GetHealthy returns the Healthy field if non-nil, zero value otherwise.

### GetHealthyOk

`func (o *OperatorHealthReply) GetHealthyOk() (*bool, bool)`

GetHealthyOk returns a tuple with the Healthy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthy

`func (o *OperatorHealthReply) SetHealthy(v bool)`

SetHealthy sets Healthy field to given value.

### HasHealthy

`func (o *OperatorHealthReply) HasHealthy() bool`

HasHealthy returns a boolean if a field has been set.

### GetServers

`func (o *OperatorHealthReply) GetServers() []ServerHealth`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *OperatorHealthReply) GetServersOk() (*[]ServerHealth, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *OperatorHealthReply) SetServers(v []ServerHealth)`

SetServers sets Servers field to given value.

### HasServers

`func (o *OperatorHealthReply) HasServers() bool`

HasServers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


