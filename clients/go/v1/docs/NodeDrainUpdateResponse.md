# NodeDrainUpdateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EvalCreateIndex** | Pointer to **int32** |  | [optional] 
**EvalIDs** | Pointer to **[]string** |  | [optional] 
**LastIndex** | Pointer to **int32** |  | [optional] 
**NodeModifyIndex** | Pointer to **int32** |  | [optional] 
**RequestTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewNodeDrainUpdateResponse

`func NewNodeDrainUpdateResponse() *NodeDrainUpdateResponse`

NewNodeDrainUpdateResponse instantiates a new NodeDrainUpdateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeDrainUpdateResponseWithDefaults

`func NewNodeDrainUpdateResponseWithDefaults() *NodeDrainUpdateResponse`

NewNodeDrainUpdateResponseWithDefaults instantiates a new NodeDrainUpdateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvalCreateIndex

`func (o *NodeDrainUpdateResponse) GetEvalCreateIndex() int32`

GetEvalCreateIndex returns the EvalCreateIndex field if non-nil, zero value otherwise.

### GetEvalCreateIndexOk

`func (o *NodeDrainUpdateResponse) GetEvalCreateIndexOk() (*int32, bool)`

GetEvalCreateIndexOk returns a tuple with the EvalCreateIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvalCreateIndex

`func (o *NodeDrainUpdateResponse) SetEvalCreateIndex(v int32)`

SetEvalCreateIndex sets EvalCreateIndex field to given value.

### HasEvalCreateIndex

`func (o *NodeDrainUpdateResponse) HasEvalCreateIndex() bool`

HasEvalCreateIndex returns a boolean if a field has been set.

### GetEvalIDs

`func (o *NodeDrainUpdateResponse) GetEvalIDs() []string`

GetEvalIDs returns the EvalIDs field if non-nil, zero value otherwise.

### GetEvalIDsOk

`func (o *NodeDrainUpdateResponse) GetEvalIDsOk() (*[]string, bool)`

GetEvalIDsOk returns a tuple with the EvalIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvalIDs

`func (o *NodeDrainUpdateResponse) SetEvalIDs(v []string)`

SetEvalIDs sets EvalIDs field to given value.

### HasEvalIDs

`func (o *NodeDrainUpdateResponse) HasEvalIDs() bool`

HasEvalIDs returns a boolean if a field has been set.

### GetLastIndex

`func (o *NodeDrainUpdateResponse) GetLastIndex() int32`

GetLastIndex returns the LastIndex field if non-nil, zero value otherwise.

### GetLastIndexOk

`func (o *NodeDrainUpdateResponse) GetLastIndexOk() (*int32, bool)`

GetLastIndexOk returns a tuple with the LastIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastIndex

`func (o *NodeDrainUpdateResponse) SetLastIndex(v int32)`

SetLastIndex sets LastIndex field to given value.

### HasLastIndex

`func (o *NodeDrainUpdateResponse) HasLastIndex() bool`

HasLastIndex returns a boolean if a field has been set.

### GetNodeModifyIndex

`func (o *NodeDrainUpdateResponse) GetNodeModifyIndex() int32`

GetNodeModifyIndex returns the NodeModifyIndex field if non-nil, zero value otherwise.

### GetNodeModifyIndexOk

`func (o *NodeDrainUpdateResponse) GetNodeModifyIndexOk() (*int32, bool)`

GetNodeModifyIndexOk returns a tuple with the NodeModifyIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeModifyIndex

`func (o *NodeDrainUpdateResponse) SetNodeModifyIndex(v int32)`

SetNodeModifyIndex sets NodeModifyIndex field to given value.

### HasNodeModifyIndex

`func (o *NodeDrainUpdateResponse) HasNodeModifyIndex() bool`

HasNodeModifyIndex returns a boolean if a field has been set.

### GetRequestTime

`func (o *NodeDrainUpdateResponse) GetRequestTime() int64`

GetRequestTime returns the RequestTime field if non-nil, zero value otherwise.

### GetRequestTimeOk

`func (o *NodeDrainUpdateResponse) GetRequestTimeOk() (*int64, bool)`

GetRequestTimeOk returns a tuple with the RequestTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTime

`func (o *NodeDrainUpdateResponse) SetRequestTime(v int64)`

SetRequestTime sets RequestTime field to given value.

### HasRequestTime

`func (o *NodeDrainUpdateResponse) HasRequestTime() bool`

HasRequestTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


